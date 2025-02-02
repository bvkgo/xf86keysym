// Copyright (c) 2024 BVK Chaitanya

//go:build ignore

// This program generates Go constants for X11 keysyms from the
// /usr/include/X11/XF86keysym.h file.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	type Entry struct {
		Keysym  string
		Value   uint32
		Comment string
	}
	var items []*Entry

	xf86keysym, err := os.ReadFile("/usr/include/X11/XF86keysym.h")
	if err != nil {
		log.Fatalf("reading XF86keysym.h: %v", err)
	}

	re := `#define \s+ (XF86XK_[a-zA-Z_0-9]+) \s+ (?: _EVDEVK\((0x[0-9a-fA-F]+)\) | (0x[0-9a-fA-F]+)) \s* (/\* .* \*/)? .*`
	xf86keysymRe := regexp.MustCompile(strings.Replace(re, " ", "", -1))

	s := bufio.NewScanner(bytes.NewBuffer(xf86keysym))
	for s.Scan() {
		text := strings.TrimSpace(s.Text())
		m := xf86keysymRe.FindStringSubmatch(text)
		if m == nil {
			continue
		}

		var value uint32
		if len(m[2]) != 0 {
			v, err := strconv.ParseUint(m[2], 0, 32)
			if err != nil {
				log.Fatalf("could not parse xf86 _EVDEVK keysym value: %v", err)
			}
			value = 0x10081000 + uint32(v)
		} else if len(m[3]) != 0 {
			v, err := strconv.ParseUint(m[3], 0, 32)
			if err != nil {
				log.Fatalf("could not parse xf86 keysym value: %v", err)
			}
			value = uint32(v)
		}

		item := &Entry{
			Keysym:  m[1],
			Value:   uint32(value),
			Comment: m[4],
		}
		items = append(items, item)
	}

	// Remove duplicates. If name/value duplicates are found, prefer the item
	// with a unicode codepoint or a comment over.
	pickOne := func(i, j int) int {
		if items[i] == nil && items[j] != nil {
			return j
		}
		if items[j] == nil && items[i] != nil {
			return i
		}
		if len(items[i].Comment) == 0 && len(items[j].Comment) != 0 {
			items[i] = nil
			return j
		}
		if len(items[j].Comment) == 0 && len(items[i].Comment) != 0 {
			items[j] = nil
			return i
		}
		items[i] = nil
		return j
	}

	nameDupMap := make(map[string]int)
	valueDupMap := make(map[uint32]int)
	for i, item := range items {
		if j, ok := nameDupMap[item.Keysym]; ok {
			i = pickOne(i, j)
		}
		if j, ok := valueDupMap[item.Value]; ok {
			i = pickOne(i, j)
		}
		nameDupMap[item.Keysym] = i
		valueDupMap[item.Value] = i
	}
	items = slices.DeleteFunc(items, func(v *Entry) bool { return v == nil })

	// Prepare a buffer with constant definitions.
	constsBuf := &bytes.Buffer{}
	for _, item := range items {
		if len(item.Comment) > 0 {
			fmt.Fprintf(constsBuf, "%s = 0x%x %s\n", item.Keysym, item.Value, item.Comment)
		} else {
			fmt.Fprintf(constsBuf, "%s = 0x%x\n", item.Keysym, item.Value)
		}
	}

	// Prepare a buffer with entries for keysym string to value map.
	symsBuf := &bytes.Buffer{}
	for _, item := range items {
		fmt.Fprintf(symsBuf, "%q: 0x%x,\n", item.Keysym, item.Value)
	}

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, `// generated by go generate; DO NOT EDIT.
package xf86keysym

import "github.com/jezek/xgb/xproto"

const (
`+constsBuf.String()+`
)

var String2KeysymMap = map[string]xproto.Keysym{
`+symsBuf.String()+`
}
`)

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s\n", buf.Bytes())
		log.Fatalf("formatting output: %v", err)
	}

	if err := os.WriteFile("keysyms.go", formatted, 0644); err != nil {
		log.Fatalf("writing keysyms.go: %v", err)
	}
}
