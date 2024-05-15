/*
Package xf86keysym defines Go constants and unicode codepoint mappings for the
X11 key symbols from /usr/include/X11/XF86keysym.h file.

Key symbol names are the same as their C language identifiers, including the
underscores, XF86XK_ prefix and capitalization.
*/
package xf86keysym

//go:generate go run gen.go
