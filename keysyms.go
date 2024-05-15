// generated by go generate; DO NOT EDIT.
package xf86keysym

import "github.com/jezek/xgb/xproto"

const (
	XF86XK_ModeLock                = 0x1008ff01 /* Mode Switch Lock */
	XF86XK_MonBrightnessUp         = 0x1008ff02 /* Monitor/panel brightness */
	XF86XK_MonBrightnessDown       = 0x1008ff03 /* Monitor/panel brightness */
	XF86XK_KbdLightOnOff           = 0x1008ff04 /* Keyboards may be lit     */
	XF86XK_KbdBrightnessUp         = 0x1008ff05 /* Keyboards may be lit     */
	XF86XK_KbdBrightnessDown       = 0x1008ff06 /* Keyboards may be lit     */
	XF86XK_MonBrightnessCycle      = 0x1008ff07 /* Monitor/panel brightness */
	XF86XK_Standby                 = 0x1008ff10 /* System into standby mode   */
	XF86XK_AudioLowerVolume        = 0x1008ff11 /* Volume control down        */
	XF86XK_AudioMute               = 0x1008ff12 /* Mute sound from the system */
	XF86XK_AudioRaiseVolume        = 0x1008ff13 /* Volume control up          */
	XF86XK_AudioPlay               = 0x1008ff14 /* Start playing of audio >   */
	XF86XK_AudioStop               = 0x1008ff15 /* Stop playing audio         */
	XF86XK_AudioPrev               = 0x1008ff16 /* Previous track             */
	XF86XK_AudioNext               = 0x1008ff17 /* Next track                 */
	XF86XK_HomePage                = 0x1008ff18 /* Display user's home page   */
	XF86XK_Mail                    = 0x1008ff19 /* Invoke user's mail program */
	XF86XK_Start                   = 0x1008ff1a /* Start application          */
	XF86XK_Search                  = 0x1008ff1b /* Search                     */
	XF86XK_AudioRecord             = 0x1008ff1c /* Record audio application   */
	XF86XK_Calculator              = 0x1008ff1d /* Invoke calculator program  */
	XF86XK_Memo                    = 0x1008ff1e /* Invoke Memo taking program */
	XF86XK_ToDoList                = 0x1008ff1f /* Invoke To Do List program  */
	XF86XK_Calendar                = 0x1008ff20 /* Invoke Calendar program    */
	XF86XK_PowerDown               = 0x1008ff21 /* Deep sleep the system      */
	XF86XK_ContrastAdjust          = 0x1008ff22 /* Adjust screen contrast     */
	XF86XK_RockerUp                = 0x1008ff23 /* Rocker switches exist up   */
	XF86XK_RockerDown              = 0x1008ff24 /* and down                   */
	XF86XK_RockerEnter             = 0x1008ff25 /* and let you press them     */
	XF86XK_Back                    = 0x1008ff26 /* Like back on a browser     */
	XF86XK_Forward                 = 0x1008ff27 /* Like forward on a browser  */
	XF86XK_Stop                    = 0x1008ff28 /* Stop current operation     */
	XF86XK_Refresh                 = 0x1008ff29 /* Refresh the page           */
	XF86XK_PowerOff                = 0x1008ff2a /* Power off system entirely  */
	XF86XK_WakeUp                  = 0x1008ff2b /* Wake up system from sleep  */
	XF86XK_Eject                   = 0x1008ff2c /* Eject device (e.g. DVD)    */
	XF86XK_ScreenSaver             = 0x1008ff2d /* Invoke screensaver         */
	XF86XK_WWW                     = 0x1008ff2e /* Invoke web browser         */
	XF86XK_Sleep                   = 0x1008ff2f /* Put system to sleep        */
	XF86XK_Favorites               = 0x1008ff30 /* Show favorite locations    */
	XF86XK_AudioPause              = 0x1008ff31 /* Pause audio playing        */
	XF86XK_AudioMedia              = 0x1008ff32 /* Launch media collection app */
	XF86XK_MyComputer              = 0x1008ff33 /* Display "My Computer" window */
	XF86XK_VendorHome              = 0x1008ff34 /* Display vendor home web site */
	XF86XK_LightBulb               = 0x1008ff35 /* Light bulb keys exist       */
	XF86XK_Shop                    = 0x1008ff36 /* Display shopping web site   */
	XF86XK_History                 = 0x1008ff37 /* Show history of web surfing */
	XF86XK_OpenURL                 = 0x1008ff38 /* Open selected URL           */
	XF86XK_AddFavorite             = 0x1008ff39 /* Add URL to favorites list   */
	XF86XK_HotLinks                = 0x1008ff3a /* Show "hot" links            */
	XF86XK_BrightnessAdjust        = 0x1008ff3b /* Invoke brightness adj. UI   */
	XF86XK_Finance                 = 0x1008ff3c /* Display financial site      */
	XF86XK_Community               = 0x1008ff3d /* Display user's community    */
	XF86XK_AudioRewind             = 0x1008ff3e /* "rewind" audio track        */
	XF86XK_BackForward             = 0x1008ff3f /* ??? */
	XF86XK_Launch0                 = 0x1008ff40 /* Launch Application          */
	XF86XK_Launch1                 = 0x1008ff41 /* Launch Application          */
	XF86XK_Launch2                 = 0x1008ff42 /* Launch Application          */
	XF86XK_Launch3                 = 0x1008ff43 /* Launch Application          */
	XF86XK_Launch4                 = 0x1008ff44 /* Launch Application          */
	XF86XK_Launch5                 = 0x1008ff45 /* Launch Application          */
	XF86XK_Launch6                 = 0x1008ff46 /* Launch Application          */
	XF86XK_Launch7                 = 0x1008ff47 /* Launch Application          */
	XF86XK_Launch8                 = 0x1008ff48 /* Launch Application          */
	XF86XK_Launch9                 = 0x1008ff49 /* Launch Application          */
	XF86XK_LaunchA                 = 0x1008ff4a /* Launch Application          */
	XF86XK_LaunchB                 = 0x1008ff4b /* Launch Application          */
	XF86XK_LaunchC                 = 0x1008ff4c /* Launch Application          */
	XF86XK_LaunchD                 = 0x1008ff4d /* Launch Application          */
	XF86XK_LaunchE                 = 0x1008ff4e /* Launch Application          */
	XF86XK_LaunchF                 = 0x1008ff4f /* Launch Application          */
	XF86XK_ApplicationLeft         = 0x1008ff50 /* switch to application, left */
	XF86XK_ApplicationRight        = 0x1008ff51 /* switch to application, right*/
	XF86XK_Book                    = 0x1008ff52 /* Launch bookreader           */
	XF86XK_CD                      = 0x1008ff53 /* Launch CD/DVD player        */
	XF86XK_Calculater              = 0x1008ff54 /* Launch Calculater           */
	XF86XK_Clear                   = 0x1008ff55 /* Clear window, screen        */
	XF86XK_Close                   = 0x1008ff56 /* Close window                */
	XF86XK_Copy                    = 0x1008ff57 /* Copy selection              */
	XF86XK_Cut                     = 0x1008ff58 /* Cut selection               */
	XF86XK_Display                 = 0x1008ff59 /* Output switch key           */
	XF86XK_DOS                     = 0x1008ff5a /* Launch DOS (emulation)      */
	XF86XK_Documents               = 0x1008ff5b /* Open documents window       */
	XF86XK_Excel                   = 0x1008ff5c /* Launch spread sheet         */
	XF86XK_Explorer                = 0x1008ff5d /* Launch file explorer        */
	XF86XK_Game                    = 0x1008ff5e /* Launch game                 */
	XF86XK_Go                      = 0x1008ff5f /* Go to URL                   */
	XF86XK_iTouch                  = 0x1008ff60 /* Logitech iTouch- don't use  */
	XF86XK_LogOff                  = 0x1008ff61 /* Log off system              */
	XF86XK_Market                  = 0x1008ff62 /* ??                          */
	XF86XK_Meeting                 = 0x1008ff63 /* enter meeting in calendar   */
	XF86XK_MenuKB                  = 0x1008ff65 /* distinguish keyboard from PB */
	XF86XK_MenuPB                  = 0x1008ff66 /* distinguish PB from keyboard */
	XF86XK_MySites                 = 0x1008ff67 /* Favourites                  */
	XF86XK_New                     = 0x1008ff68 /* New (folder, document...    */
	XF86XK_News                    = 0x1008ff69 /* News                        */
	XF86XK_OfficeHome              = 0x1008ff6a /* Office home (old Staroffice)*/
	XF86XK_Open                    = 0x1008ff6b /* Open                        */
	XF86XK_Option                  = 0x1008ff6c /* ?? */
	XF86XK_Paste                   = 0x1008ff6d /* Paste                       */
	XF86XK_Phone                   = 0x1008ff6e /* Launch phone; dial number   */
	XF86XK_Q                       = 0x1008ff70 /* Compaq's Q - don't use      */
	XF86XK_Reply                   = 0x1008ff72 /* Reply e.g., mail            */
	XF86XK_Reload                  = 0x1008ff73 /* Reload web page, file, etc. */
	XF86XK_RotateWindows           = 0x1008ff74 /* Rotate windows e.g. xrandr  */
	XF86XK_RotationPB              = 0x1008ff75 /* don't use                   */
	XF86XK_RotationKB              = 0x1008ff76 /* don't use                   */
	XF86XK_Save                    = 0x1008ff77 /* Save (file, document, state */
	XF86XK_ScrollUp                = 0x1008ff78 /* Scroll window/contents up   */
	XF86XK_ScrollDown              = 0x1008ff79 /* Scrool window/contentd down */
	XF86XK_ScrollClick             = 0x1008ff7a /* Use XKB mousekeys instead   */
	XF86XK_Send                    = 0x1008ff7b /* Send mail, file, object     */
	XF86XK_Spell                   = 0x1008ff7c /* Spell checker               */
	XF86XK_SplitScreen             = 0x1008ff7d /* Split window or screen      */
	XF86XK_Support                 = 0x1008ff7e /* Get support (??)            */
	XF86XK_TaskPane                = 0x1008ff7f /* Show tasks */
	XF86XK_Terminal                = 0x1008ff80 /* Launch terminal emulator    */
	XF86XK_Tools                   = 0x1008ff81 /* toolbox of desktop/app.     */
	XF86XK_Travel                  = 0x1008ff82 /* ?? */
	XF86XK_UserPB                  = 0x1008ff84 /* ?? */
	XF86XK_User1KB                 = 0x1008ff85 /* ?? */
	XF86XK_User2KB                 = 0x1008ff86 /* ?? */
	XF86XK_Video                   = 0x1008ff87 /* Launch video player       */
	XF86XK_WheelButton             = 0x1008ff88 /* button from a mouse wheel */
	XF86XK_Word                    = 0x1008ff89 /* Launch word processor     */
	XF86XK_Xfer                    = 0x1008ff8a
	XF86XK_ZoomIn                  = 0x1008ff8b /* zoom in view, map, etc.   */
	XF86XK_ZoomOut                 = 0x1008ff8c /* zoom out view, map, etc.  */
	XF86XK_Away                    = 0x1008ff8d /* mark yourself as away     */
	XF86XK_Messenger               = 0x1008ff8e /* as in instant messaging   */
	XF86XK_WebCam                  = 0x1008ff8f /* Launch web camera app.    */
	XF86XK_MailForward             = 0x1008ff90 /* Forward in mail           */
	XF86XK_Pictures                = 0x1008ff91 /* Show pictures             */
	XF86XK_Music                   = 0x1008ff92 /* Launch music application  */
	XF86XK_Battery                 = 0x1008ff93 /* Display battery information */
	XF86XK_Bluetooth               = 0x1008ff94 /* Enable/disable Bluetooth    */
	XF86XK_WLAN                    = 0x1008ff95 /* Enable/disable WLAN         */
	XF86XK_UWB                     = 0x1008ff96 /* Enable/disable UWB	    */
	XF86XK_AudioForward            = 0x1008ff97 /* fast-forward audio track    */
	XF86XK_AudioRepeat             = 0x1008ff98 /* toggle repeat mode          */
	XF86XK_AudioRandomPlay         = 0x1008ff99 /* toggle shuffle mode         */
	XF86XK_Subtitle                = 0x1008ff9a /* cycle through subtitle      */
	XF86XK_AudioCycleTrack         = 0x1008ff9b /* cycle through audio tracks  */
	XF86XK_CycleAngle              = 0x1008ff9c /* cycle through angles        */
	XF86XK_FrameBack               = 0x1008ff9d /* video: go one frame back    */
	XF86XK_FrameForward            = 0x1008ff9e /* video: go one frame forward */
	XF86XK_Time                    = 0x1008ff9f /* display, or shows an entry for time seeking */
	XF86XK_Select                  = 0x1008ffa0 /* Select button on joypads and remotes */
	XF86XK_View                    = 0x1008ffa1 /* Show a view options/properties */
	XF86XK_TopMenu                 = 0x1008ffa2 /* Go to a top-level menu in a video */
	XF86XK_Red                     = 0x1008ffa3 /* Red button                  */
	XF86XK_Green                   = 0x1008ffa4 /* Green button                */
	XF86XK_Yellow                  = 0x1008ffa5 /* Yellow button               */
	XF86XK_Blue                    = 0x1008ffa6 /* Blue button                 */
	XF86XK_Suspend                 = 0x1008ffa7 /* Sleep to RAM                */
	XF86XK_Hibernate               = 0x1008ffa8 /* Sleep to disk               */
	XF86XK_TouchpadToggle          = 0x1008ffa9 /* Toggle between touchpad/trackstick */
	XF86XK_TouchpadOn              = 0x1008ffb0 /* The touchpad got switched on */
	XF86XK_TouchpadOff             = 0x1008ffb1 /* The touchpad got switched off */
	XF86XK_AudioMicMute            = 0x1008ffb2 /* Mute the Mic from the system */
	XF86XK_Keyboard                = 0x1008ffb3 /* User defined keyboard related action */
	XF86XK_WWAN                    = 0x1008ffb4 /* Toggle WWAN (LTE, UMTS, etc.) radio */
	XF86XK_RFKill                  = 0x1008ffb5 /* Toggle radios on/off */
	XF86XK_AudioPreset             = 0x1008ffb6 /* Select equalizer preset, e.g. theatre-mode */
	XF86XK_RotationLockToggle      = 0x1008ffb7 /* Toggle screen rotation lock on/off */
	XF86XK_FullScreen              = 0x1008ffb8 /* Toggle fullscreen */
	XF86XK_Switch_VT_1             = 0x1008fe01
	XF86XK_Switch_VT_2             = 0x1008fe02
	XF86XK_Switch_VT_3             = 0x1008fe03
	XF86XK_Switch_VT_4             = 0x1008fe04
	XF86XK_Switch_VT_5             = 0x1008fe05
	XF86XK_Switch_VT_6             = 0x1008fe06
	XF86XK_Switch_VT_7             = 0x1008fe07
	XF86XK_Switch_VT_8             = 0x1008fe08
	XF86XK_Switch_VT_9             = 0x1008fe09
	XF86XK_Switch_VT_10            = 0x1008fe0a
	XF86XK_Switch_VT_11            = 0x1008fe0b
	XF86XK_Switch_VT_12            = 0x1008fe0c
	XF86XK_Ungrab                  = 0x1008fe20 /* force ungrab               */
	XF86XK_ClearGrab               = 0x1008fe21 /* kill application with grab */
	XF86XK_Next_VMode              = 0x1008fe22 /* next video mode available  */
	XF86XK_Prev_VMode              = 0x1008fe23 /* prev. video mode available */
	XF86XK_LogWindowTree           = 0x1008fe24 /* print window tree to log   */
	XF86XK_LogGrabInfo             = 0x1008fe25 /* print all active grabs to log */
	XF86XK_BrightnessAuto          = 0x100810f4 /* v3.16 KEY_BRIGHTNESS_AUTO */
	XF86XK_DisplayOff              = 0x100810f5 /* v2.6.23 KEY_DISPLAY_OFF */
	XF86XK_Info                    = 0x10081166 /*       KEY_INFO */
	XF86XK_AspectRatio             = 0x10081177 /* v5.1  KEY_ASPECT_RATIO */
	XF86XK_DVD                     = 0x10081185 /*       KEY_DVD */
	XF86XK_Audio                   = 0x10081188 /*       KEY_AUDIO */
	XF86XK_ChannelUp               = 0x10081192 /*       KEY_CHANNELUP */
	XF86XK_ChannelDown             = 0x10081193 /*       KEY_CHANNELDOWN */
	XF86XK_Break                   = 0x1008119b /*       KEY_BREAK */
	XF86XK_VideoPhone              = 0x100811a0 /* v2.6.20 KEY_VIDEOPHONE */
	XF86XK_ZoomReset               = 0x100811a4 /* v2.6.20 KEY_ZOOMRESET */
	XF86XK_Editor                  = 0x100811a6 /* v2.6.20 KEY_EDITOR */
	XF86XK_GraphicsEditor          = 0x100811a8 /* v2.6.20 KEY_GRAPHICSEDITOR */
	XF86XK_Presentation            = 0x100811a9 /* v2.6.20 KEY_PRESENTATION */
	XF86XK_Database                = 0x100811aa /* v2.6.20 KEY_DATABASE */
	XF86XK_Voicemail               = 0x100811ac /* v2.6.20 KEY_VOICEMAIL */
	XF86XK_Addressbook             = 0x100811ad /* v2.6.20 KEY_ADDRESSBOOK */
	XF86XK_DisplayToggle           = 0x100811af /* v2.6.20 KEY_DISPLAYTOGGLE */
	XF86XK_SpellCheck              = 0x100811b0 /* v2.6.24 KEY_SPELLCHECK */
	XF86XK_ContextMenu             = 0x100811b6 /* v2.6.24 KEY_CONTEXT_MENU */
	XF86XK_MediaRepeat             = 0x100811b7 /* v2.6.26 KEY_MEDIA_REPEAT */
	XF86XK_10ChannelsUp            = 0x100811b8 /* v2.6.38 KEY_10CHANNELSUP */
	XF86XK_10ChannelsDown          = 0x100811b9 /* v2.6.38 KEY_10CHANNELSDOWN */
	XF86XK_Images                  = 0x100811ba /* v2.6.39 KEY_IMAGES */
	XF86XK_NotificationCenter      = 0x100811bc /* v5.10 KEY_NOTIFICATION_CENTER */
	XF86XK_PickupPhone             = 0x100811bd /* v5.10 KEY_PICKUP_PHONE */
	XF86XK_HangupPhone             = 0x100811be /* v5.10 KEY_HANGUP_PHONE */
	XF86XK_Fn                      = 0x100811d0 /*       KEY_FN */
	XF86XK_Fn_Esc                  = 0x100811d1 /*       KEY_FN_ESC */
	XF86XK_FnRightShift            = 0x100811e5 /* v5.10 KEY_FN_RIGHT_SHIFT */
	XF86XK_Numeric0                = 0x10081200 /* v2.6.28 KEY_NUMERIC_0 */
	XF86XK_Numeric1                = 0x10081201 /* v2.6.28 KEY_NUMERIC_1 */
	XF86XK_Numeric2                = 0x10081202 /* v2.6.28 KEY_NUMERIC_2 */
	XF86XK_Numeric3                = 0x10081203 /* v2.6.28 KEY_NUMERIC_3 */
	XF86XK_Numeric4                = 0x10081204 /* v2.6.28 KEY_NUMERIC_4 */
	XF86XK_Numeric5                = 0x10081205 /* v2.6.28 KEY_NUMERIC_5 */
	XF86XK_Numeric6                = 0x10081206 /* v2.6.28 KEY_NUMERIC_6 */
	XF86XK_Numeric7                = 0x10081207 /* v2.6.28 KEY_NUMERIC_7 */
	XF86XK_Numeric8                = 0x10081208 /* v2.6.28 KEY_NUMERIC_8 */
	XF86XK_Numeric9                = 0x10081209 /* v2.6.28 KEY_NUMERIC_9 */
	XF86XK_NumericStar             = 0x1008120a /* v2.6.28 KEY_NUMERIC_STAR */
	XF86XK_NumericPound            = 0x1008120b /* v2.6.28 KEY_NUMERIC_POUND */
	XF86XK_NumericA                = 0x1008120c /* v4.1  KEY_NUMERIC_A */
	XF86XK_NumericB                = 0x1008120d /* v4.1  KEY_NUMERIC_B */
	XF86XK_NumericC                = 0x1008120e /* v4.1  KEY_NUMERIC_C */
	XF86XK_NumericD                = 0x1008120f /* v4.1  KEY_NUMERIC_D */
	XF86XK_CameraFocus             = 0x10081210 /* v2.6.33 KEY_CAMERA_FOCUS */
	XF86XK_WPSButton               = 0x10081211 /* v2.6.34 KEY_WPS_BUTTON */
	XF86XK_CameraZoomIn            = 0x10081215 /* v2.6.39 KEY_CAMERA_ZOOMIN */
	XF86XK_CameraZoomOut           = 0x10081216 /* v2.6.39 KEY_CAMERA_ZOOMOUT */
	XF86XK_CameraUp                = 0x10081217 /* v2.6.39 KEY_CAMERA_UP */
	XF86XK_CameraDown              = 0x10081218 /* v2.6.39 KEY_CAMERA_DOWN */
	XF86XK_CameraLeft              = 0x10081219 /* v2.6.39 KEY_CAMERA_LEFT */
	XF86XK_CameraRight             = 0x1008121a /* v2.6.39 KEY_CAMERA_RIGHT */
	XF86XK_AttendantOn             = 0x1008121b /* v3.10 KEY_ATTENDANT_ON */
	XF86XK_AttendantOff            = 0x1008121c /* v3.10 KEY_ATTENDANT_OFF */
	XF86XK_AttendantToggle         = 0x1008121d /* v3.10 KEY_ATTENDANT_TOGGLE */
	XF86XK_LightsToggle            = 0x1008121e /* v3.10 KEY_LIGHTS_TOGGLE */
	XF86XK_ALSToggle               = 0x10081230 /* v3.13 KEY_ALS_TOGGLE */
	XF86XK_Buttonconfig            = 0x10081240 /* v3.16 KEY_BUTTONCONFIG */
	XF86XK_Taskmanager             = 0x10081241 /* v3.16 KEY_TASKMANAGER */
	XF86XK_Journal                 = 0x10081242 /* v3.16 KEY_JOURNAL */
	XF86XK_ControlPanel            = 0x10081243 /* v3.16 KEY_CONTROLPANEL */
	XF86XK_AppSelect               = 0x10081244 /* v3.16 KEY_APPSELECT */
	XF86XK_Screensaver             = 0x10081245 /* v3.16 KEY_SCREENSAVER */
	XF86XK_VoiceCommand            = 0x10081246 /* v3.16 KEY_VOICECOMMAND */
	XF86XK_Assistant               = 0x10081247 /* v4.13 KEY_ASSISTANT */
	XF86XK_EmojiPicker             = 0x10081249 /* v5.13 KEY_EMOJI_PICKER */
	XF86XK_Dictate                 = 0x1008124a /* v5.17 KEY_DICTATE */
	XF86XK_BrightnessMin           = 0x10081250 /* v3.16 KEY_BRIGHTNESS_MIN */
	XF86XK_BrightnessMax           = 0x10081251 /* v3.16 KEY_BRIGHTNESS_MAX */
	XF86XK_KbdInputAssistPrev      = 0x10081260 /* v3.18 KEY_KBDINPUTASSIST_PREV */
	XF86XK_KbdInputAssistNext      = 0x10081261 /* v3.18 KEY_KBDINPUTASSIST_NEXT */
	XF86XK_KbdInputAssistPrevgroup = 0x10081262 /* v3.18 KEY_KBDINPUTASSIST_PREVGROUP */
	XF86XK_KbdInputAssistNextgroup = 0x10081263 /* v3.18 KEY_KBDINPUTASSIST_NEXTGROUP */
	XF86XK_KbdInputAssistAccept    = 0x10081264 /* v3.18 KEY_KBDINPUTASSIST_ACCEPT */
	XF86XK_KbdInputAssistCancel    = 0x10081265 /* v3.18 KEY_KBDINPUTASSIST_CANCEL */
	XF86XK_RightUp                 = 0x10081266 /* v4.7  KEY_RIGHT_UP */
	XF86XK_RightDown               = 0x10081267 /* v4.7  KEY_RIGHT_DOWN */
	XF86XK_LeftUp                  = 0x10081268 /* v4.7  KEY_LEFT_UP */
	XF86XK_LeftDown                = 0x10081269 /* v4.7  KEY_LEFT_DOWN */
	XF86XK_RootMenu                = 0x1008126a /* v4.7  KEY_ROOT_MENU */
	XF86XK_MediaTopMenu            = 0x1008126b /* v4.7  KEY_MEDIA_TOP_MENU */
	XF86XK_Numeric11               = 0x1008126c /* v4.7  KEY_NUMERIC_11 */
	XF86XK_Numeric12               = 0x1008126d /* v4.7  KEY_NUMERIC_12 */
	XF86XK_AudioDesc               = 0x1008126e /* v4.7  KEY_AUDIO_DESC */
	XF86XK_3DMode                  = 0x1008126f /* v4.7  KEY_3D_MODE */
	XF86XK_NextFavorite            = 0x10081270 /* v4.7  KEY_NEXT_FAVORITE */
	XF86XK_StopRecord              = 0x10081271 /* v4.7  KEY_STOP_RECORD */
	XF86XK_PauseRecord             = 0x10081272 /* v4.7  KEY_PAUSE_RECORD */
	XF86XK_VOD                     = 0x10081273 /* v4.7  KEY_VOD */
	XF86XK_Unmute                  = 0x10081274 /* v4.7  KEY_UNMUTE */
	XF86XK_FastReverse             = 0x10081275 /* v4.7  KEY_FASTREVERSE */
	XF86XK_SlowReverse             = 0x10081276 /* v4.7  KEY_SLOWREVERSE */
	XF86XK_Data                    = 0x10081277 /* v4.7  KEY_DATA */
	XF86XK_OnScreenKeyboard        = 0x10081278 /* v4.12 KEY_ONSCREEN_KEYBOARD */
	XF86XK_PrivacyScreenToggle     = 0x10081279 /* v5.5  KEY_PRIVACY_SCREEN_TOGGLE */
	XF86XK_SelectiveScreenshot     = 0x1008127a /* v5.6  KEY_SELECTIVE_SCREENSHOT */
	XF86XK_Macro1                  = 0x10081290 /* v5.5  KEY_MACRO1 */
	XF86XK_Macro2                  = 0x10081291 /* v5.5  KEY_MACRO2 */
	XF86XK_Macro3                  = 0x10081292 /* v5.5  KEY_MACRO3 */
	XF86XK_Macro4                  = 0x10081293 /* v5.5  KEY_MACRO4 */
	XF86XK_Macro5                  = 0x10081294 /* v5.5  KEY_MACRO5 */
	XF86XK_Macro6                  = 0x10081295 /* v5.5  KEY_MACRO6 */
	XF86XK_Macro7                  = 0x10081296 /* v5.5  KEY_MACRO7 */
	XF86XK_Macro8                  = 0x10081297 /* v5.5  KEY_MACRO8 */
	XF86XK_Macro9                  = 0x10081298 /* v5.5  KEY_MACRO9 */
	XF86XK_Macro10                 = 0x10081299 /* v5.5  KEY_MACRO10 */
	XF86XK_Macro11                 = 0x1008129a /* v5.5  KEY_MACRO11 */
	XF86XK_Macro12                 = 0x1008129b /* v5.5  KEY_MACRO12 */
	XF86XK_Macro13                 = 0x1008129c /* v5.5  KEY_MACRO13 */
	XF86XK_Macro14                 = 0x1008129d /* v5.5  KEY_MACRO14 */
	XF86XK_Macro15                 = 0x1008129e /* v5.5  KEY_MACRO15 */
	XF86XK_Macro16                 = 0x1008129f /* v5.5  KEY_MACRO16 */
	XF86XK_Macro17                 = 0x100812a0 /* v5.5  KEY_MACRO17 */
	XF86XK_Macro18                 = 0x100812a1 /* v5.5  KEY_MACRO18 */
	XF86XK_Macro19                 = 0x100812a2 /* v5.5  KEY_MACRO19 */
	XF86XK_Macro20                 = 0x100812a3 /* v5.5  KEY_MACRO20 */
	XF86XK_Macro21                 = 0x100812a4 /* v5.5  KEY_MACRO21 */
	XF86XK_Macro22                 = 0x100812a5 /* v5.5  KEY_MACRO22 */
	XF86XK_Macro23                 = 0x100812a6 /* v5.5  KEY_MACRO23 */
	XF86XK_Macro24                 = 0x100812a7 /* v5.5  KEY_MACRO24 */
	XF86XK_Macro25                 = 0x100812a8 /* v5.5  KEY_MACRO25 */
	XF86XK_Macro26                 = 0x100812a9 /* v5.5  KEY_MACRO26 */
	XF86XK_Macro27                 = 0x100812aa /* v5.5  KEY_MACRO27 */
	XF86XK_Macro28                 = 0x100812ab /* v5.5  KEY_MACRO28 */
	XF86XK_Macro29                 = 0x100812ac /* v5.5  KEY_MACRO29 */
	XF86XK_Macro30                 = 0x100812ad /* v5.5  KEY_MACRO30 */
	XF86XK_MacroRecordStart        = 0x100812b0 /* v5.5  KEY_MACRO_RECORD_START */
	XF86XK_MacroRecordStop         = 0x100812b1 /* v5.5  KEY_MACRO_RECORD_STOP */
	XF86XK_MacroPresetCycle        = 0x100812b2 /* v5.5  KEY_MACRO_PRESET_CYCLE */
	XF86XK_MacroPreset1            = 0x100812b3 /* v5.5  KEY_MACRO_PRESET1 */
	XF86XK_MacroPreset2            = 0x100812b4 /* v5.5  KEY_MACRO_PRESET2 */
	XF86XK_MacroPreset3            = 0x100812b5 /* v5.5  KEY_MACRO_PRESET3 */
	XF86XK_KbdLcdMenu1             = 0x100812b8 /* v5.5  KEY_KBD_LCD_MENU1 */
	XF86XK_KbdLcdMenu2             = 0x100812b9 /* v5.5  KEY_KBD_LCD_MENU2 */
	XF86XK_KbdLcdMenu3             = 0x100812ba /* v5.5  KEY_KBD_LCD_MENU3 */
	XF86XK_KbdLcdMenu4             = 0x100812bb /* v5.5  KEY_KBD_LCD_MENU4 */
	XF86XK_KbdLcdMenu5             = 0x100812bc /* v5.5  KEY_KBD_LCD_MENU5 */

)

var String2KeysymMap = map[string]xproto.Keysym{
	"XF86XK_ModeLock":                0x1008ff01,
	"XF86XK_MonBrightnessUp":         0x1008ff02,
	"XF86XK_MonBrightnessDown":       0x1008ff03,
	"XF86XK_KbdLightOnOff":           0x1008ff04,
	"XF86XK_KbdBrightnessUp":         0x1008ff05,
	"XF86XK_KbdBrightnessDown":       0x1008ff06,
	"XF86XK_MonBrightnessCycle":      0x1008ff07,
	"XF86XK_Standby":                 0x1008ff10,
	"XF86XK_AudioLowerVolume":        0x1008ff11,
	"XF86XK_AudioMute":               0x1008ff12,
	"XF86XK_AudioRaiseVolume":        0x1008ff13,
	"XF86XK_AudioPlay":               0x1008ff14,
	"XF86XK_AudioStop":               0x1008ff15,
	"XF86XK_AudioPrev":               0x1008ff16,
	"XF86XK_AudioNext":               0x1008ff17,
	"XF86XK_HomePage":                0x1008ff18,
	"XF86XK_Mail":                    0x1008ff19,
	"XF86XK_Start":                   0x1008ff1a,
	"XF86XK_Search":                  0x1008ff1b,
	"XF86XK_AudioRecord":             0x1008ff1c,
	"XF86XK_Calculator":              0x1008ff1d,
	"XF86XK_Memo":                    0x1008ff1e,
	"XF86XK_ToDoList":                0x1008ff1f,
	"XF86XK_Calendar":                0x1008ff20,
	"XF86XK_PowerDown":               0x1008ff21,
	"XF86XK_ContrastAdjust":          0x1008ff22,
	"XF86XK_RockerUp":                0x1008ff23,
	"XF86XK_RockerDown":              0x1008ff24,
	"XF86XK_RockerEnter":             0x1008ff25,
	"XF86XK_Back":                    0x1008ff26,
	"XF86XK_Forward":                 0x1008ff27,
	"XF86XK_Stop":                    0x1008ff28,
	"XF86XK_Refresh":                 0x1008ff29,
	"XF86XK_PowerOff":                0x1008ff2a,
	"XF86XK_WakeUp":                  0x1008ff2b,
	"XF86XK_Eject":                   0x1008ff2c,
	"XF86XK_ScreenSaver":             0x1008ff2d,
	"XF86XK_WWW":                     0x1008ff2e,
	"XF86XK_Sleep":                   0x1008ff2f,
	"XF86XK_Favorites":               0x1008ff30,
	"XF86XK_AudioPause":              0x1008ff31,
	"XF86XK_AudioMedia":              0x1008ff32,
	"XF86XK_MyComputer":              0x1008ff33,
	"XF86XK_VendorHome":              0x1008ff34,
	"XF86XK_LightBulb":               0x1008ff35,
	"XF86XK_Shop":                    0x1008ff36,
	"XF86XK_History":                 0x1008ff37,
	"XF86XK_OpenURL":                 0x1008ff38,
	"XF86XK_AddFavorite":             0x1008ff39,
	"XF86XK_HotLinks":                0x1008ff3a,
	"XF86XK_BrightnessAdjust":        0x1008ff3b,
	"XF86XK_Finance":                 0x1008ff3c,
	"XF86XK_Community":               0x1008ff3d,
	"XF86XK_AudioRewind":             0x1008ff3e,
	"XF86XK_BackForward":             0x1008ff3f,
	"XF86XK_Launch0":                 0x1008ff40,
	"XF86XK_Launch1":                 0x1008ff41,
	"XF86XK_Launch2":                 0x1008ff42,
	"XF86XK_Launch3":                 0x1008ff43,
	"XF86XK_Launch4":                 0x1008ff44,
	"XF86XK_Launch5":                 0x1008ff45,
	"XF86XK_Launch6":                 0x1008ff46,
	"XF86XK_Launch7":                 0x1008ff47,
	"XF86XK_Launch8":                 0x1008ff48,
	"XF86XK_Launch9":                 0x1008ff49,
	"XF86XK_LaunchA":                 0x1008ff4a,
	"XF86XK_LaunchB":                 0x1008ff4b,
	"XF86XK_LaunchC":                 0x1008ff4c,
	"XF86XK_LaunchD":                 0x1008ff4d,
	"XF86XK_LaunchE":                 0x1008ff4e,
	"XF86XK_LaunchF":                 0x1008ff4f,
	"XF86XK_ApplicationLeft":         0x1008ff50,
	"XF86XK_ApplicationRight":        0x1008ff51,
	"XF86XK_Book":                    0x1008ff52,
	"XF86XK_CD":                      0x1008ff53,
	"XF86XK_Calculater":              0x1008ff54,
	"XF86XK_Clear":                   0x1008ff55,
	"XF86XK_Close":                   0x1008ff56,
	"XF86XK_Copy":                    0x1008ff57,
	"XF86XK_Cut":                     0x1008ff58,
	"XF86XK_Display":                 0x1008ff59,
	"XF86XK_DOS":                     0x1008ff5a,
	"XF86XK_Documents":               0x1008ff5b,
	"XF86XK_Excel":                   0x1008ff5c,
	"XF86XK_Explorer":                0x1008ff5d,
	"XF86XK_Game":                    0x1008ff5e,
	"XF86XK_Go":                      0x1008ff5f,
	"XF86XK_iTouch":                  0x1008ff60,
	"XF86XK_LogOff":                  0x1008ff61,
	"XF86XK_Market":                  0x1008ff62,
	"XF86XK_Meeting":                 0x1008ff63,
	"XF86XK_MenuKB":                  0x1008ff65,
	"XF86XK_MenuPB":                  0x1008ff66,
	"XF86XK_MySites":                 0x1008ff67,
	"XF86XK_New":                     0x1008ff68,
	"XF86XK_News":                    0x1008ff69,
	"XF86XK_OfficeHome":              0x1008ff6a,
	"XF86XK_Open":                    0x1008ff6b,
	"XF86XK_Option":                  0x1008ff6c,
	"XF86XK_Paste":                   0x1008ff6d,
	"XF86XK_Phone":                   0x1008ff6e,
	"XF86XK_Q":                       0x1008ff70,
	"XF86XK_Reply":                   0x1008ff72,
	"XF86XK_Reload":                  0x1008ff73,
	"XF86XK_RotateWindows":           0x1008ff74,
	"XF86XK_RotationPB":              0x1008ff75,
	"XF86XK_RotationKB":              0x1008ff76,
	"XF86XK_Save":                    0x1008ff77,
	"XF86XK_ScrollUp":                0x1008ff78,
	"XF86XK_ScrollDown":              0x1008ff79,
	"XF86XK_ScrollClick":             0x1008ff7a,
	"XF86XK_Send":                    0x1008ff7b,
	"XF86XK_Spell":                   0x1008ff7c,
	"XF86XK_SplitScreen":             0x1008ff7d,
	"XF86XK_Support":                 0x1008ff7e,
	"XF86XK_TaskPane":                0x1008ff7f,
	"XF86XK_Terminal":                0x1008ff80,
	"XF86XK_Tools":                   0x1008ff81,
	"XF86XK_Travel":                  0x1008ff82,
	"XF86XK_UserPB":                  0x1008ff84,
	"XF86XK_User1KB":                 0x1008ff85,
	"XF86XK_User2KB":                 0x1008ff86,
	"XF86XK_Video":                   0x1008ff87,
	"XF86XK_WheelButton":             0x1008ff88,
	"XF86XK_Word":                    0x1008ff89,
	"XF86XK_Xfer":                    0x1008ff8a,
	"XF86XK_ZoomIn":                  0x1008ff8b,
	"XF86XK_ZoomOut":                 0x1008ff8c,
	"XF86XK_Away":                    0x1008ff8d,
	"XF86XK_Messenger":               0x1008ff8e,
	"XF86XK_WebCam":                  0x1008ff8f,
	"XF86XK_MailForward":             0x1008ff90,
	"XF86XK_Pictures":                0x1008ff91,
	"XF86XK_Music":                   0x1008ff92,
	"XF86XK_Battery":                 0x1008ff93,
	"XF86XK_Bluetooth":               0x1008ff94,
	"XF86XK_WLAN":                    0x1008ff95,
	"XF86XK_UWB":                     0x1008ff96,
	"XF86XK_AudioForward":            0x1008ff97,
	"XF86XK_AudioRepeat":             0x1008ff98,
	"XF86XK_AudioRandomPlay":         0x1008ff99,
	"XF86XK_Subtitle":                0x1008ff9a,
	"XF86XK_AudioCycleTrack":         0x1008ff9b,
	"XF86XK_CycleAngle":              0x1008ff9c,
	"XF86XK_FrameBack":               0x1008ff9d,
	"XF86XK_FrameForward":            0x1008ff9e,
	"XF86XK_Time":                    0x1008ff9f,
	"XF86XK_Select":                  0x1008ffa0,
	"XF86XK_View":                    0x1008ffa1,
	"XF86XK_TopMenu":                 0x1008ffa2,
	"XF86XK_Red":                     0x1008ffa3,
	"XF86XK_Green":                   0x1008ffa4,
	"XF86XK_Yellow":                  0x1008ffa5,
	"XF86XK_Blue":                    0x1008ffa6,
	"XF86XK_Suspend":                 0x1008ffa7,
	"XF86XK_Hibernate":               0x1008ffa8,
	"XF86XK_TouchpadToggle":          0x1008ffa9,
	"XF86XK_TouchpadOn":              0x1008ffb0,
	"XF86XK_TouchpadOff":             0x1008ffb1,
	"XF86XK_AudioMicMute":            0x1008ffb2,
	"XF86XK_Keyboard":                0x1008ffb3,
	"XF86XK_WWAN":                    0x1008ffb4,
	"XF86XK_RFKill":                  0x1008ffb5,
	"XF86XK_AudioPreset":             0x1008ffb6,
	"XF86XK_RotationLockToggle":      0x1008ffb7,
	"XF86XK_FullScreen":              0x1008ffb8,
	"XF86XK_Switch_VT_1":             0x1008fe01,
	"XF86XK_Switch_VT_2":             0x1008fe02,
	"XF86XK_Switch_VT_3":             0x1008fe03,
	"XF86XK_Switch_VT_4":             0x1008fe04,
	"XF86XK_Switch_VT_5":             0x1008fe05,
	"XF86XK_Switch_VT_6":             0x1008fe06,
	"XF86XK_Switch_VT_7":             0x1008fe07,
	"XF86XK_Switch_VT_8":             0x1008fe08,
	"XF86XK_Switch_VT_9":             0x1008fe09,
	"XF86XK_Switch_VT_10":            0x1008fe0a,
	"XF86XK_Switch_VT_11":            0x1008fe0b,
	"XF86XK_Switch_VT_12":            0x1008fe0c,
	"XF86XK_Ungrab":                  0x1008fe20,
	"XF86XK_ClearGrab":               0x1008fe21,
	"XF86XK_Next_VMode":              0x1008fe22,
	"XF86XK_Prev_VMode":              0x1008fe23,
	"XF86XK_LogWindowTree":           0x1008fe24,
	"XF86XK_LogGrabInfo":             0x1008fe25,
	"XF86XK_BrightnessAuto":          0x100810f4,
	"XF86XK_DisplayOff":              0x100810f5,
	"XF86XK_Info":                    0x10081166,
	"XF86XK_AspectRatio":             0x10081177,
	"XF86XK_DVD":                     0x10081185,
	"XF86XK_Audio":                   0x10081188,
	"XF86XK_ChannelUp":               0x10081192,
	"XF86XK_ChannelDown":             0x10081193,
	"XF86XK_Break":                   0x1008119b,
	"XF86XK_VideoPhone":              0x100811a0,
	"XF86XK_ZoomReset":               0x100811a4,
	"XF86XK_Editor":                  0x100811a6,
	"XF86XK_GraphicsEditor":          0x100811a8,
	"XF86XK_Presentation":            0x100811a9,
	"XF86XK_Database":                0x100811aa,
	"XF86XK_Voicemail":               0x100811ac,
	"XF86XK_Addressbook":             0x100811ad,
	"XF86XK_DisplayToggle":           0x100811af,
	"XF86XK_SpellCheck":              0x100811b0,
	"XF86XK_ContextMenu":             0x100811b6,
	"XF86XK_MediaRepeat":             0x100811b7,
	"XF86XK_10ChannelsUp":            0x100811b8,
	"XF86XK_10ChannelsDown":          0x100811b9,
	"XF86XK_Images":                  0x100811ba,
	"XF86XK_NotificationCenter":      0x100811bc,
	"XF86XK_PickupPhone":             0x100811bd,
	"XF86XK_HangupPhone":             0x100811be,
	"XF86XK_Fn":                      0x100811d0,
	"XF86XK_Fn_Esc":                  0x100811d1,
	"XF86XK_FnRightShift":            0x100811e5,
	"XF86XK_Numeric0":                0x10081200,
	"XF86XK_Numeric1":                0x10081201,
	"XF86XK_Numeric2":                0x10081202,
	"XF86XK_Numeric3":                0x10081203,
	"XF86XK_Numeric4":                0x10081204,
	"XF86XK_Numeric5":                0x10081205,
	"XF86XK_Numeric6":                0x10081206,
	"XF86XK_Numeric7":                0x10081207,
	"XF86XK_Numeric8":                0x10081208,
	"XF86XK_Numeric9":                0x10081209,
	"XF86XK_NumericStar":             0x1008120a,
	"XF86XK_NumericPound":            0x1008120b,
	"XF86XK_NumericA":                0x1008120c,
	"XF86XK_NumericB":                0x1008120d,
	"XF86XK_NumericC":                0x1008120e,
	"XF86XK_NumericD":                0x1008120f,
	"XF86XK_CameraFocus":             0x10081210,
	"XF86XK_WPSButton":               0x10081211,
	"XF86XK_CameraZoomIn":            0x10081215,
	"XF86XK_CameraZoomOut":           0x10081216,
	"XF86XK_CameraUp":                0x10081217,
	"XF86XK_CameraDown":              0x10081218,
	"XF86XK_CameraLeft":              0x10081219,
	"XF86XK_CameraRight":             0x1008121a,
	"XF86XK_AttendantOn":             0x1008121b,
	"XF86XK_AttendantOff":            0x1008121c,
	"XF86XK_AttendantToggle":         0x1008121d,
	"XF86XK_LightsToggle":            0x1008121e,
	"XF86XK_ALSToggle":               0x10081230,
	"XF86XK_Buttonconfig":            0x10081240,
	"XF86XK_Taskmanager":             0x10081241,
	"XF86XK_Journal":                 0x10081242,
	"XF86XK_ControlPanel":            0x10081243,
	"XF86XK_AppSelect":               0x10081244,
	"XF86XK_Screensaver":             0x10081245,
	"XF86XK_VoiceCommand":            0x10081246,
	"XF86XK_Assistant":               0x10081247,
	"XF86XK_EmojiPicker":             0x10081249,
	"XF86XK_Dictate":                 0x1008124a,
	"XF86XK_BrightnessMin":           0x10081250,
	"XF86XK_BrightnessMax":           0x10081251,
	"XF86XK_KbdInputAssistPrev":      0x10081260,
	"XF86XK_KbdInputAssistNext":      0x10081261,
	"XF86XK_KbdInputAssistPrevgroup": 0x10081262,
	"XF86XK_KbdInputAssistNextgroup": 0x10081263,
	"XF86XK_KbdInputAssistAccept":    0x10081264,
	"XF86XK_KbdInputAssistCancel":    0x10081265,
	"XF86XK_RightUp":                 0x10081266,
	"XF86XK_RightDown":               0x10081267,
	"XF86XK_LeftUp":                  0x10081268,
	"XF86XK_LeftDown":                0x10081269,
	"XF86XK_RootMenu":                0x1008126a,
	"XF86XK_MediaTopMenu":            0x1008126b,
	"XF86XK_Numeric11":               0x1008126c,
	"XF86XK_Numeric12":               0x1008126d,
	"XF86XK_AudioDesc":               0x1008126e,
	"XF86XK_3DMode":                  0x1008126f,
	"XF86XK_NextFavorite":            0x10081270,
	"XF86XK_StopRecord":              0x10081271,
	"XF86XK_PauseRecord":             0x10081272,
	"XF86XK_VOD":                     0x10081273,
	"XF86XK_Unmute":                  0x10081274,
	"XF86XK_FastReverse":             0x10081275,
	"XF86XK_SlowReverse":             0x10081276,
	"XF86XK_Data":                    0x10081277,
	"XF86XK_OnScreenKeyboard":        0x10081278,
	"XF86XK_PrivacyScreenToggle":     0x10081279,
	"XF86XK_SelectiveScreenshot":     0x1008127a,
	"XF86XK_Macro1":                  0x10081290,
	"XF86XK_Macro2":                  0x10081291,
	"XF86XK_Macro3":                  0x10081292,
	"XF86XK_Macro4":                  0x10081293,
	"XF86XK_Macro5":                  0x10081294,
	"XF86XK_Macro6":                  0x10081295,
	"XF86XK_Macro7":                  0x10081296,
	"XF86XK_Macro8":                  0x10081297,
	"XF86XK_Macro9":                  0x10081298,
	"XF86XK_Macro10":                 0x10081299,
	"XF86XK_Macro11":                 0x1008129a,
	"XF86XK_Macro12":                 0x1008129b,
	"XF86XK_Macro13":                 0x1008129c,
	"XF86XK_Macro14":                 0x1008129d,
	"XF86XK_Macro15":                 0x1008129e,
	"XF86XK_Macro16":                 0x1008129f,
	"XF86XK_Macro17":                 0x100812a0,
	"XF86XK_Macro18":                 0x100812a1,
	"XF86XK_Macro19":                 0x100812a2,
	"XF86XK_Macro20":                 0x100812a3,
	"XF86XK_Macro21":                 0x100812a4,
	"XF86XK_Macro22":                 0x100812a5,
	"XF86XK_Macro23":                 0x100812a6,
	"XF86XK_Macro24":                 0x100812a7,
	"XF86XK_Macro25":                 0x100812a8,
	"XF86XK_Macro26":                 0x100812a9,
	"XF86XK_Macro27":                 0x100812aa,
	"XF86XK_Macro28":                 0x100812ab,
	"XF86XK_Macro29":                 0x100812ac,
	"XF86XK_Macro30":                 0x100812ad,
	"XF86XK_MacroRecordStart":        0x100812b0,
	"XF86XK_MacroRecordStop":         0x100812b1,
	"XF86XK_MacroPresetCycle":        0x100812b2,
	"XF86XK_MacroPreset1":            0x100812b3,
	"XF86XK_MacroPreset2":            0x100812b4,
	"XF86XK_MacroPreset3":            0x100812b5,
	"XF86XK_KbdLcdMenu1":             0x100812b8,
	"XF86XK_KbdLcdMenu2":             0x100812b9,
	"XF86XK_KbdLcdMenu3":             0x100812ba,
	"XF86XK_KbdLcdMenu4":             0x100812bb,
	"XF86XK_KbdLcdMenu5":             0x100812bc,
}
