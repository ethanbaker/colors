// Package xterm provides all of the xterm colors.
package xterm

import "math/rand"

/* ANSI Color Value Constants */

const (
	Black             = "\x1b[38;2;0;0;0m"
	Maroon            = "\x1b[38;2;128;0;0m"
	Green             = "\x1b[38;2;0;128;0m"
	Olive             = "\x1b[38;2;128;128;0m"
	Navy              = "\x1b[38;2;0;0;128m"
	Purple            = "\x1b[38;2;128;0;128m"
	Teal              = "\x1b[38;2;0;128;128m"
	Silver            = "\x1b[38;2;192;192;192m"
	Grey              = "\x1b[38;2;128;128;128m"
	Red               = "\x1b[38;2;255;0;0m"
	Lime              = "\x1b[38;2;0;255;0m"
	Yellow            = "\x1b[38;2;255;255;0m"
	Blue              = "\x1b[38;2;0;0;255m"
	Fuchsia           = "\x1b[38;2;255;0;255m"
	Aqua              = "\x1b[38;2;0;255;255m"
	White             = "\x1b[38;2;255;255;255m"
	NavajoWhite1      = "\x1b[38;2;255;215;175m"
	NavajoWhite3      = "\x1b[38;2;175;175;135m"
	Red1              = "\x1b[38;2;255;0;0m"
	Red2              = "\x1b[38;2;215;0;0m"
	Red3              = "\x1b[38;2;175;0;0m"
	DarkRed1          = "\x1b[38;2;135;0;0m"
	DarkRed2          = "\x1b[38;2;95;0;0m"
	IndianRed1        = "\x1b[38;2;255;95;135m"
	IndianRed2        = "\x1b[38;2;255;95;95m"
	IndianRed3        = "\x1b[38;2;215;95;95m"
	IndianRed4        = "\x1b[38;2;175;95;95m"
	OrangeRed1        = "\x1b[38;2;255;95;0m"
	Orange1           = "\x1b[38;2;255;175;0m"
	Orange2           = "\x1b[38;2;215;135;0m"
	Orange3           = "\x1b[38;2;135;95;0m"
	Orange4           = "\x1b[38;2;95;95;0m"
	DarkOrange1       = "\x1b[38;2;255;135;0m"
	DarkOrange2       = "\x1b[38;2;215;95;0m"
	DarkOrange3       = "\x1b[38;2;175;95;0m"
	Gold1             = "\x1b[38;2;255;215;0m"
	Gold2             = "\x1b[38;2;215;175;0m"
	Gold3             = "\x1b[38;2;175;175;0m"
	LightGoldenrod1   = "\x1b[38;2;255;255;95m"
	LightGoldenrod2   = "\x1b[38;2;255;215;135m"
	LightGoldenrod3   = "\x1b[38;2;255;215;95m"
	LightGoldenrod4   = "\x1b[38;2;215;215;135m"
	LightGoldenrod5   = "\x1b[38;2;215;175;95m"
	DarkGoldenrod     = "\x1b[38;2;175;135;0m"
	Yellow1           = "\x1b[38;2;255;255;0m"
	Yellow2           = "\x1b[38;2;215;255;0m"
	Yellow3           = "\x1b[38;2;215;215;0m"
	Yellow4           = "\x1b[38;2;175;215;0m"
	Yellow5           = "\x1b[38;2;135;175;0m"
	Yellow6           = "\x1b[38;2;135;135;0m"
	LightYellow1      = "\x1b[38;2;215;215;175m"
	GreenYellow       = "\x1b[38;2;175;255;0m"
	Green1            = "\x1b[38;2;0;255;0m"
	Green2            = "\x1b[38;2;0;215;0m"
	Green3            = "\x1b[38;2;0;175;0m"
	Green4            = "\x1b[38;2;0;135;0m"
	DarkGreen         = "\x1b[38;2;0;95;0m"
	LightGreen1       = "\x1b[38;2;135;255;135m"
	LightGreen2       = "\x1b[38;2;135;255;95m"
	PaleGreen1        = "\x1b[38;2;175;255;135m"
	PaleGreen2        = "\x1b[38;2;135;255;175m"
	PaleGreen3        = "\x1b[38;2;135;215;135m"
	PaleGreen4        = "\x1b[38;2;95;215;95m"
	DarkOliveGreen1   = "\x1b[38;2;215;255;135m"
	DarkOliveGreen2   = "\x1b[38;2;215;255;95m"
	DarkOliveGreen3   = "\x1b[38;2;175;255;95m"
	DarkOliveGreen4   = "\x1b[38;2;175;215;95m"
	DarkOliveGreen5   = "\x1b[38;2;135;215;95m"
	DarkOliveGreen6   = "\x1b[38;2;135;175;95m"
	SpringGreen1      = "\x1b[38;2;0;255;135m"
	SpringGreen2      = "\x1b[38;2;0;255;95m"
	SpringGreen3      = "\x1b[38;2;0;215;135m"
	SpringGreen4      = "\x1b[38;2;0;215;95m"
	SpringGreen5      = "\x1b[38;2;0;175;95m"
	SpringGreen6      = "\x1b[38;2;0;135;95m"
	MediumSpringGreen = "\x1b[38;2;0;255;175m"
	SeaGreen1         = "\x1b[38;2;95;255;175m"
	SeaGreen2         = "\x1b[38;2;95;255;135m"
	SeaGreen3         = "\x1b[38;2;95;255;95m"
	SeaGreen4         = "\x1b[38;2;95;215;135m"
	DarkSeaGreen1     = "\x1b[38;2;215;255;175m"
	DarkSeaGreen2     = "\x1b[38;2;175;255;215m"
	DarkSeaGreen3     = "\x1b[38;2;175;255;175m"
	DarkSeaGreen4     = "\x1b[38;2;175;215;175m"
	DarkSeaGreen5     = "\x1b[38;2;135;215;175m"
	DarkSeaGreen6     = "\x1b[38;2;175;215;135m"
	DarkSeaGreen7     = "\x1b[38;2;135;175;135m"
	DarkSeaGreen8     = "\x1b[38;2;95;175;95m"
	DarkSeaGreen9     = "\x1b[38;2;95;135;95m"
	LightSeaGreen     = "\x1b[38;2;0;175;175m"
	Chartreuse1       = "\x1b[38;2;135;255;0m"
	Chartreuse2       = "\x1b[38;2;135;215;0m"
	Chartreuse3       = "\x1b[38;2;95;255;0m"
	Chartreuse4       = "\x1b[38;2;95;215;0m"
	Chartreuse5       = "\x1b[38;2;95;175;0m"
	Chartreuse6       = "\x1b[38;2;95;135;0m"
	Blue1             = "\x1b[38;2;0;0;255m"
	Blue2             = "\x1b[38;2;0;0;175m"
	Blue3             = "\x1b[38;2;0;0;215m"
	DarkBlue          = "\x1b[38;2;0;0;135m"
	NavyBlue          = "\x1b[38;2;0;0;95m"
	LightCoral        = "\x1b[38;2;255;135;135m"
	RoyalBlue         = "\x1b[38;2;95;95;255m"
	CornflowerBlue    = "\x1b[38;2;95;135;255m"
	CadetBlue1        = "\x1b[38;2;95;175;175m"
	CadetBlue2        = "\x1b[38;2;95;175;135m"
	DeepSkyBlue1      = "\x1b[38;2;0;135;175m"
	DeepSkyBlue2      = "\x1b[38;2;0;135;215m"
	DeepSkyBlue3      = "\x1b[38;2;0;95;175m"
	DeepSkyBlue4      = "\x1b[38;2;0;95;135m"
	DeepSkyBlue5      = "\x1b[38;2;0;95;95m"
	DeepSkyBlue6      = "\x1b[38;2;0;175;215m"
	DeepSkyBlue7      = "\x1b[38;2;0;175;255m"
	DodgerBlue1       = "\x1b[38;2;0;95;215m"
	DodgerBlue2       = "\x1b[38;2;0;95;255m"
	DodgerBlue3       = "\x1b[38;2;0;135;255m"
	SlateBlue1        = "\x1b[38;2;135;95;255m"
	SlateBlue2        = "\x1b[38;2;95;95;215m"
	SlateBlue3        = "\x1b[38;2;95;95;175m"
	LightSlateBlue    = "\x1b[38;2;135;135;255m"
	SteelBlue1        = "\x1b[38;2;95;215;255m"
	SteelBlue2        = "\x1b[38;2;95;175;255m"
	SteelBlue3        = "\x1b[38;2;95;135;215m"
	SteelBlue4        = "\x1b[38;2;95;135;175m"
	LightSteelBlue1   = "\x1b[38;2;215;215;255m"
	LightSteelBlue2   = "\x1b[38;2;175;175;255m"
	LightSteelBlue3   = "\x1b[38;2;175;175;215m"
	SkyBlue1          = "\x1b[38;2;135;215;255m"
	SkyBlue2          = "\x1b[38;2;135;175;255m"
	SkyBlue3          = "\x1b[38;2;95;175;215m"
	LightSkyBlue1     = "\x1b[38;2;175;215;255m"
	LightSkyBlue2     = "\x1b[38;2;135;175;215m"
	LightSkyBlue3     = "\x1b[38;2;135;175;175m"
	Aquamarine1       = "\x1b[38;2;135;255;215m"
	Aquamarine2       = "\x1b[38;2;95;255;215m"
	Aquamarine3       = "\x1b[38;2;95;215;175m"
	Cyan1             = "\x1b[38;2;0;255;255m"
	Cyan2             = "\x1b[38;2;0;255;215m"
	Cyan3             = "\x1b[38;2;0;215;175m"
	DarkCyan          = "\x1b[38;2;0;175;135m"
	LightCyan1        = "\x1b[38;2;215;255;255m"
	LightCyan2        = "\x1b[38;2;175;215;215m"
	Turquoise1        = "\x1b[38;2;0;215;255m"
	Turquoise2        = "\x1b[38;2;0;135;135m"
	DarkTurquoise     = "\x1b[38;2;0;215;215m"
	MediumTurquoise   = "\x1b[38;2;95;215;215m"
	PaleTurquoise1    = "\x1b[38;2;175;255;255m"
	PaleTurquoise2    = "\x1b[38;2;95;135;135m"
	BlueViolet        = "\x1b[38;2;95;0;255m"
	Violet            = "\x1b[38;2;215;135;255m"
	DarkViolet1       = "\x1b[38;2;175;0;215m"
	DarkViolet2       = "\x1b[38;2;135;0;215m"
	PaleVioletRed     = "\x1b[38;2;255;135;175m"
	MediumVioletRed   = "\x1b[38;2;175;0;135m"
	Plum1             = "\x1b[38;2;255;175;255m"
	Plum2             = "\x1b[38;2;215;175;255m"
	Plum3             = "\x1b[38;2;215;135;215m"
	Plum4             = "\x1b[38;2;135;95;135m"
	Purple1           = "\x1b[38;2;175;0;255m"
	Purple2           = "\x1b[38;2;135;0;255m"
	Purple3           = "\x1b[38;2;95;0;215m"
	Purple4           = "\x1b[38;2;95;0;175m"
	Purple5           = "\x1b[38;2;95;0;135m"
	MediumPurple1     = "\x1b[38;2;175;135;255m"
	MediumPurple2     = "\x1b[38;2;175;95;255m"
	MediumPurple3     = "\x1b[38;2;175;135;215m"
	MediumPurple4     = "\x1b[38;2;135;135;215m"
	MediumPurple5     = "\x1b[38;2;135;95;215m"
	MediumPurple6     = "\x1b[38;2;135;95;175m"
	MediumPurple7     = "\x1b[38;2;95;95;135m"
	Pink1             = "\x1b[38;2;255;175;215m"
	Pink2             = "\x1b[38;2;215;135;175m"
	DeepPink1         = "\x1b[38;2;255;0;175m"
	DeepPink2         = "\x1b[38;2;255;0;135m"
	DeepPink3         = "\x1b[38;2;255;0;95m"
	DeepPink4         = "\x1b[38;2;215;0;135m"
	DeepPink5         = "\x1b[38;2;215;0;95m"
	DeepPink6         = "\x1b[38;2;135;0;95m"
	DeepPink7         = "\x1b[38;2;95;0;95m"
	DeepPink8         = "\x1b[38;2;175;0;95m"
	HotPink1          = "\x1b[38;2;255;95;215m"
	HotPink2          = "\x1b[38;2;255;95;175m"
	HotPink3          = "\x1b[38;2;215;95;175m"
	HotPink4          = "\x1b[38;2;215;95;135m"
	HotPink5          = "\x1b[38;2;175;95;135m"
	LightPink1        = "\x1b[38;2;255;175;175m"
	LightPink2        = "\x1b[38;2;215;135;135m"
	LightPink3        = "\x1b[38;2;135;95;95m"
	Salmon1           = "\x1b[38;2;255;135;95m"
	LightSalmon1      = "\x1b[38;2;255;175;135m"
	LightSalmon2      = "\x1b[38;2;215;135;95m"
	LightSalmon3      = "\x1b[38;2;175;135;95m"
	MistyRose1        = "\x1b[38;2;255;215;215m"
	MistyRose3        = "\x1b[38;2;215;175;175m"
	Magenta1          = "\x1b[38;2;255;0;255m"
	Magenta2          = "\x1b[38;2;255;0;215m"
	Magenta3          = "\x1b[38;2;215;0;255m"
	Magenta4          = "\x1b[38;2;215;0;215m"
	Magenta5          = "\x1b[38;2;215;0;175m"
	Magenta6          = "\x1b[38;2;175;0;175m"
	DarkMagenta1      = "\x1b[38;2;135;0;175m"
	DarkMagenta2      = "\x1b[38;2;135;0;135m"
	Orchid1           = "\x1b[38;2;215;95;215m"
	Orchid2           = "\x1b[38;2;255;135;215m"
	Orchid3           = "\x1b[38;2;255;135;255m"
	MediumOrchid1     = "\x1b[38;2;175;95;175m"
	MediumOrchid2     = "\x1b[38;2;175;95;215m"
	MediumOrchid3     = "\x1b[38;2;215;95;255m"
	MediumOrchid4     = "\x1b[38;2;255;95;255m"
	Wheat1            = "\x1b[38;2;255;255;175m"
	Wheat2            = "\x1b[38;2;135;135;95m"
	SandyBrown        = "\x1b[38;2;255;175;95m"
	RosyBrown         = "\x1b[38;2;175;135;135m"
	Khaki1            = "\x1b[38;2;255;255;135m"
	Khaki3            = "\x1b[38;2;215;215;95m"
	DarkKhaki         = "\x1b[38;2;175;175;95m"
	Tan               = "\x1b[38;2;215;175;135m"
	Thistle1          = "\x1b[38;2;255;215;255m"
	Thistle3          = "\x1b[38;2;215;175;215m"
	Honeydew2         = "\x1b[38;2;215;255;215m"
	Cornsilk1         = "\x1b[38;2;255;255;215m"
	DarkSlateGray1    = "\x1b[38;2;135;255;255m"
	DarkSlateGray2    = "\x1b[38;2;95;255;255m"
	DarkSlateGray3    = "\x1b[38;2;135;215;215m"
	LightSlateGrey    = "\x1b[38;2;135;135;175m"
	Grey0             = "\x1b[38;2;0;0;0m"
	Grey3             = "\x1b[38;2;8;8;8m"
	Grey7             = "\x1b[38;2;18;18;18m"
	Grey11            = "\x1b[38;2;28;28;28m"
	Grey15            = "\x1b[38;2;38;38;38m"
	Grey19            = "\x1b[38;2;48;48;48m"
	Grey23            = "\x1b[38;2;58;58;58m"
	Grey27            = "\x1b[38;2;68;68;68m"
	Grey30            = "\x1b[38;2;78;78;78m"
	Grey35            = "\x1b[38;2;88;88;88m"
	Grey37            = "\x1b[38;2;95;95;95m"
	Grey39            = "\x1b[38;2;96;96;96m"
	Grey42            = "\x1b[38;2;102;102;102m"
	Grey46            = "\x1b[38;2;118;118;118m"
	Grey50            = "\x1b[38;2;128;128;128m"
	Grey53            = "\x1b[38;2;135;135;135m"
	Grey54            = "\x1b[38;2;138;138;138m"
	Grey58            = "\x1b[38;2;148;148;148m"
	Grey62            = "\x1b[38;2;158;158;158m"
	Grey63            = "\x1b[38;2;175;135;175m"
	Grey66            = "\x1b[38;2;168;168;168m"
	Grey69            = "\x1b[38;2;175;175;175m"
	Grey70            = "\x1b[38;2;178;178;178m"
	Grey74            = "\x1b[38;2;188;188;188m"
	Grey78            = "\x1b[38;2;198;198;198m"
	Grey82            = "\x1b[38;2;208;208;208m"
	Grey84            = "\x1b[38;2;215;215;215m"
	Grey85            = "\x1b[38;2;218;218;218m"
	Grey89            = "\x1b[38;2;228;228;228m"
	Grey93            = "\x1b[38;2;238;238;238m"
	Grey100           = "\x1b[38;2;255;255;255m"

	BlackBackground             = "\x1b[48;2;0;0;0m"
	MaroonBackground            = "\x1b[48;2;128;0;0m"
	GreenBackground             = "\x1b[48;2;0;128;0m"
	OliveBackground             = "\x1b[48;2;128;128;0m"
	NavyBackground              = "\x1b[48;2;0;0;128m"
	PurpleBackground            = "\x1b[48;2;128;0;128m"
	TealBackground              = "\x1b[48;2;0;128;128m"
	SilverBackground            = "\x1b[48;2;192;192;192m"
	GreyBackground              = "\x1b[48;2;128;128;128m"
	RedBackground               = "\x1b[48;2;255;0;0m"
	LimeBackground              = "\x1b[48;2;0;255;0m"
	YellowBackground            = "\x1b[48;2;255;255;0m"
	BlueBackground              = "\x1b[48;2;0;0;255m"
	FuchsiaBackground           = "\x1b[48;2;255;0;255m"
	AquaBackground              = "\x1b[48;2;0;255;255m"
	WhiteBackground             = "\x1b[48;2;255;255;255m"
	NavajoWhite1Background      = "\x1b[48;2;255;215;175m"
	NavajoWhite3Background      = "\x1b[48;2;175;175;135m"
	Red1Background              = "\x1b[48;2;255;0;0m"
	Red2Background              = "\x1b[48;2;215;0;0m"
	Red3Background              = "\x1b[48;2;175;0;0m"
	DarkRed1Background          = "\x1b[48;2;135;0;0m"
	DarkRed2Background          = "\x1b[48;2;95;0;0m"
	IndianRed1Background        = "\x1b[48;2;255;95;135m"
	IndianRed2Background        = "\x1b[48;2;255;95;95m"
	IndianRed3Background        = "\x1b[48;2;215;95;95m"
	IndianRed4Background        = "\x1b[48;2;175;95;95m"
	OrangeRed1Background        = "\x1b[48;2;255;95;0m"
	Orange1Background           = "\x1b[48;2;255;175;0m"
	Orange2Background           = "\x1b[48;2;215;135;0m"
	Orange3Background           = "\x1b[48;2;135;95;0m"
	Orange4Background           = "\x1b[48;2;95;95;0m"
	DarkOrange1Background       = "\x1b[48;2;255;135;0m"
	DarkOrange2Background       = "\x1b[48;2;215;95;0m"
	DarkOrange3Background       = "\x1b[48;2;175;95;0m"
	Gold1Background             = "\x1b[48;2;255;215;0m"
	Gold2Background             = "\x1b[48;2;215;175;0m"
	Gold3Background             = "\x1b[48;2;175;175;0m"
	LightGoldenrod1Background   = "\x1b[48;2;255;255;95m"
	LightGoldenrod2Background   = "\x1b[48;2;255;215;135m"
	LightGoldenrod3Background   = "\x1b[48;2;255;215;95m"
	LightGoldenrod4Background   = "\x1b[48;2;215;215;135m"
	LightGoldenrod5Background   = "\x1b[48;2;215;175;95m"
	DarkGoldenrodBackground     = "\x1b[48;2;175;135;0m"
	Yellow1Background           = "\x1b[48;2;255;255;0m"
	Yellow2Background           = "\x1b[48;2;215;255;0m"
	Yellow3Background           = "\x1b[48;2;215;215;0m"
	Yellow4Background           = "\x1b[48;2;175;215;0m"
	Yellow5Background           = "\x1b[48;2;135;175;0m"
	Yellow6Background           = "\x1b[48;2;135;135;0m"
	LightYellow1Background      = "\x1b[48;2;215;215;175m"
	GreenYellowBackground       = "\x1b[48;2;175;255;0m"
	Green1Background            = "\x1b[48;2;0;255;0m"
	Green2Background            = "\x1b[48;2;0;215;0m"
	Green3Background            = "\x1b[48;2;0;175;0m"
	Green4Background            = "\x1b[48;2;0;135;0m"
	DarkGreenBackground         = "\x1b[48;2;0;95;0m"
	LightGreen1Background       = "\x1b[48;2;135;255;135m"
	LightGreen2Background       = "\x1b[48;2;135;255;95m"
	PaleGreen1Background        = "\x1b[48;2;175;255;135m"
	PaleGreen2Background        = "\x1b[48;2;135;255;175m"
	PaleGreen3Background        = "\x1b[48;2;135;215;135m"
	PaleGreen4Background        = "\x1b[48;2;95;215;95m"
	DarkOliveGreen1Background   = "\x1b[48;2;215;255;135m"
	DarkOliveGreen2Background   = "\x1b[48;2;215;255;95m"
	DarkOliveGreen3Background   = "\x1b[48;2;175;255;95m"
	DarkOliveGreen4Background   = "\x1b[48;2;175;215;95m"
	DarkOliveGreen5Background   = "\x1b[48;2;135;215;95m"
	DarkOliveGreen6Background   = "\x1b[48;2;135;175;95m"
	SpringGreen1Background      = "\x1b[48;2;0;255;135m"
	SpringGreen2Background      = "\x1b[48;2;0;255;95m"
	SpringGreen3Background      = "\x1b[48;2;0;215;135m"
	SpringGreen4Background      = "\x1b[48;2;0;215;95m"
	SpringGreen5Background      = "\x1b[48;2;0;175;95m"
	SpringGreen6Background      = "\x1b[48;2;0;135;95m"
	MediumSpringGreenBackground = "\x1b[48;2;0;255;175m"
	SeaGreen1Background         = "\x1b[48;2;95;255;175m"
	SeaGreen2Background         = "\x1b[48;2;95;255;135m"
	SeaGreen3Background         = "\x1b[48;2;95;255;95m"
	SeaGreen4Background         = "\x1b[48;2;95;215;135m"
	DarkSeaGreen1Background     = "\x1b[48;2;215;255;175m"
	DarkSeaGreen2Background     = "\x1b[48;2;175;255;215m"
	DarkSeaGreen3Background     = "\x1b[48;2;175;255;175m"
	DarkSeaGreen4Background     = "\x1b[48;2;175;215;175m"
	DarkSeaGreen5Background     = "\x1b[48;2;135;215;175m"
	DarkSeaGreen6Background     = "\x1b[48;2;175;215;135m"
	DarkSeaGreen7Background     = "\x1b[48;2;135;175;135m"
	DarkSeaGreen8Background     = "\x1b[48;2;95;175;95m"
	DarkSeaGreen9Background     = "\x1b[48;2;95;135;95m"
	LightSeaGreenBackground     = "\x1b[48;2;0;175;175m"
	Chartreuse1Background       = "\x1b[48;2;135;255;0m"
	Chartreuse2Background       = "\x1b[48;2;135;215;0m"
	Chartreuse3Background       = "\x1b[48;2;95;255;0m"
	Chartreuse4Background       = "\x1b[48;2;95;215;0m"
	Chartreuse5Background       = "\x1b[48;2;95;175;0m"
	Chartreuse6Background       = "\x1b[48;2;95;135;0m"
	Blue1Background             = "\x1b[48;2;0;0;255m"
	Blue2Background             = "\x1b[48;2;0;0;175m"
	Blue3Background             = "\x1b[48;2;0;0;215m"
	DarkBlueBackground          = "\x1b[48;2;0;0;135m"
	NavyBlueBackground          = "\x1b[48;2;0;0;95m"
	LightCoralBackground        = "\x1b[48;2;255;135;135m"
	RoyalBlueBackground         = "\x1b[48;2;95;95;255m"
	CornflowerBlueBackground    = "\x1b[48;2;95;135;255m"
	CadetBlue1Background        = "\x1b[48;2;95;175;175m"
	CadetBlue2Background        = "\x1b[48;2;95;175;135m"
	DeepSkyBlue1Background      = "\x1b[48;2;0;135;175m"
	DeepSkyBlue2Background      = "\x1b[48;2;0;135;215m"
	DeepSkyBlue3Background      = "\x1b[48;2;0;95;175m"
	DeepSkyBlue4Background      = "\x1b[48;2;0;95;135m"
	DeepSkyBlue5Background      = "\x1b[48;2;0;95;95m"
	DeepSkyBlue6Background      = "\x1b[48;2;0;175;215m"
	DeepSkyBlue7Background      = "\x1b[48;2;0;175;255m"
	DodgerBlue1Background       = "\x1b[48;2;0;95;215m"
	DodgerBlue2Background       = "\x1b[48;2;0;95;255m"
	DodgerBlue3Background       = "\x1b[48;2;0;135;255m"
	SlateBlue1Background        = "\x1b[48;2;135;95;255m"
	SlateBlue2Background        = "\x1b[48;2;95;95;215m"
	SlateBlue3Background        = "\x1b[48;2;95;95;175m"
	LightSlateBlueBackground    = "\x1b[48;2;135;135;255m"
	SteelBlue1Background        = "\x1b[48;2;95;215;255m"
	SteelBlue2Background        = "\x1b[48;2;95;175;255m"
	SteelBlue3Background        = "\x1b[48;2;95;135;215m"
	SteelBlue4Background        = "\x1b[48;2;95;135;175m"
	LightSteelBlue1Background   = "\x1b[48;2;215;215;255m"
	LightSteelBlue2Background   = "\x1b[48;2;175;175;255m"
	LightSteelBlue3Background   = "\x1b[48;2;175;175;215m"
	SkyBlue1Background          = "\x1b[48;2;135;215;255m"
	SkyBlue2Background          = "\x1b[48;2;135;175;255m"
	SkyBlue3Background          = "\x1b[48;2;95;175;215m"
	LightSkyBlue1Background     = "\x1b[48;2;175;215;255m"
	LightSkyBlue2Background     = "\x1b[48;2;135;175;215m"
	LightSkyBlue3Background     = "\x1b[48;2;135;175;175m"
	Aquamarine1Background       = "\x1b[48;2;135;255;215m"
	Aquamarine2Background       = "\x1b[48;2;95;255;215m"
	Aquamarine3Background       = "\x1b[48;2;95;215;175m"
	Cyan1Background             = "\x1b[48;2;0;255;255m"
	Cyan2Background             = "\x1b[48;2;0;255;215m"
	Cyan3Background             = "\x1b[48;2;0;215;175m"
	DarkCyanBackground          = "\x1b[48;2;0;175;135m"
	LightCyan1Background        = "\x1b[48;2;215;255;255m"
	LightCyan2Background        = "\x1b[48;2;175;215;215m"
	Turquoise1Background        = "\x1b[48;2;0;215;255m"
	Turquoise2Background        = "\x1b[48;2;0;135;135m"
	DarkTurquoiseBackground     = "\x1b[48;2;0;215;215m"
	MediumTurquoiseBackground   = "\x1b[48;2;95;215;215m"
	PaleTurquoise1Background    = "\x1b[48;2;175;255;255m"
	PaleTurquoise2Background    = "\x1b[48;2;95;135;135m"
	BlueVioletBackground        = "\x1b[48;2;95;0;255m"
	VioletBackground            = "\x1b[48;2;215;135;255m"
	DarkViolet1Background       = "\x1b[48;2;175;0;215m"
	DarkViolet2Background       = "\x1b[48;2;135;0;215m"
	PaleVioletRedBackground     = "\x1b[48;2;255;135;175m"
	MediumVioletRedBackground   = "\x1b[48;2;175;0;135m"
	Plum1Background             = "\x1b[48;2;255;175;255m"
	Plum2Background             = "\x1b[48;2;215;175;255m"
	Plum3Background             = "\x1b[48;2;215;135;215m"
	Plum4Background             = "\x1b[48;2;135;95;135m"
	Purple1Background           = "\x1b[48;2;175;0;255m"
	Purple2Background           = "\x1b[48;2;135;0;255m"
	Purple3Background           = "\x1b[48;2;95;0;215m"
	Purple4Background           = "\x1b[48;2;95;0;175m"
	Purple5Background           = "\x1b[48;2;95;0;135m"
	MediumPurple1Background     = "\x1b[48;2;175;135;255m"
	MediumPurple2Background     = "\x1b[48;2;175;95;255m"
	MediumPurple3Background     = "\x1b[48;2;175;135;215m"
	MediumPurple4Background     = "\x1b[48;2;135;135;215m"
	MediumPurple5Background     = "\x1b[48;2;135;95;215m"
	MediumPurple6Background     = "\x1b[48;2;135;95;175m"
	MediumPurple7Background     = "\x1b[48;2;95;95;135m"
	Pink1Background             = "\x1b[48;2;255;175;215m"
	Pink2Background             = "\x1b[48;2;215;135;175m"
	DeepPink1Background         = "\x1b[48;2;255;0;175m"
	DeepPink2Background         = "\x1b[48;2;255;0;135m"
	DeepPink3Background         = "\x1b[48;2;255;0;95m"
	DeepPink4Background         = "\x1b[48;2;215;0;135m"
	DeepPink5Background         = "\x1b[48;2;215;0;95m"
	DeepPink6Background         = "\x1b[48;2;135;0;95m"
	DeepPink7Background         = "\x1b[48;2;95;0;95m"
	DeepPink8Background         = "\x1b[48;2;175;0;95m"
	HotPink1Background          = "\x1b[48;2;255;95;215m"
	HotPink2Background          = "\x1b[48;2;255;95;175m"
	HotPink3Background          = "\x1b[48;2;215;95;175m"
	HotPink4Background          = "\x1b[48;2;215;95;135m"
	HotPink5Background          = "\x1b[48;2;175;95;135m"
	LightPink1Background        = "\x1b[48;2;255;175;175m"
	LightPink2Background        = "\x1b[48;2;215;135;135m"
	LightPink3Background        = "\x1b[48;2;135;95;95m"
	Salmon1Background           = "\x1b[48;2;255;135;95m"
	LightSalmon1Background      = "\x1b[48;2;255;175;135m"
	LightSalmon2Background      = "\x1b[48;2;215;135;95m"
	LightSalmon3Background      = "\x1b[48;2;175;135;95m"
	MistyRose1Background        = "\x1b[48;2;255;215;215m"
	MistyRose3Background        = "\x1b[48;2;215;175;175m"
	Magenta1Background          = "\x1b[48;2;255;0;255m"
	Magenta2Background          = "\x1b[48;2;255;0;215m"
	Magenta3Background          = "\x1b[48;2;215;0;255m"
	Magenta4Background          = "\x1b[48;2;215;0;215m"
	Magenta5Background          = "\x1b[48;2;215;0;175m"
	Magenta6Background          = "\x1b[48;2;175;0;175m"
	DarkMagenta1Background      = "\x1b[48;2;135;0;175m"
	DarkMagenta2Background      = "\x1b[48;2;135;0;135m"
	Orchid1Background           = "\x1b[48;2;215;95;215m"
	Orchid2Background           = "\x1b[48;2;255;135;215m"
	Orchid3Background           = "\x1b[48;2;255;135;255m"
	MediumOrchid1Background     = "\x1b[48;2;175;95;175m"
	MediumOrchid2Background     = "\x1b[48;2;175;95;215m"
	MediumOrchid3Background     = "\x1b[48;2;215;95;255m"
	MediumOrchid4Background     = "\x1b[48;2;255;95;255m"
	Wheat1Background            = "\x1b[48;2;255;255;175m"
	Wheat2Background            = "\x1b[48;2;135;135;95m"
	SandyBrownBackground        = "\x1b[48;2;255;175;95m"
	RosyBrownBackground         = "\x1b[48;2;175;135;135m"
	Khaki1Background            = "\x1b[48;2;255;255;135m"
	Khaki3Background            = "\x1b[48;2;215;215;95m"
	DarkKhakiBackground         = "\x1b[48;2;175;175;95m"
	TanBackground               = "\x1b[48;2;215;175;135m"
	Thistle1Background          = "\x1b[48;2;255;215;255m"
	Thistle3Background          = "\x1b[48;2;215;175;215m"
	Honeydew2Background         = "\x1b[48;2;215;255;215m"
	Cornsilk1Background         = "\x1b[48;2;255;255;215m"
	DarkSlateGray1Background    = "\x1b[48;2;135;255;255m"
	DarkSlateGray2Background    = "\x1b[48;2;95;255;255m"
	DarkSlateGray3Background    = "\x1b[48;2;135;215;215m"
	LightSlateGreyBackground    = "\x1b[48;2;135;135;175m"
	Grey0Background             = "\x1b[48;2;0;0;0m"
	Grey3Background             = "\x1b[48;2;8;8;8m"
	Grey7Background             = "\x1b[48;2;18;18;18m"
	Grey11Background            = "\x1b[48;2;28;28;28m"
	Grey15Background            = "\x1b[48;2;38;38;38m"
	Grey19Background            = "\x1b[48;2;48;48;48m"
	Grey23Background            = "\x1b[48;2;58;58;58m"
	Grey27Background            = "\x1b[48;2;68;68;68m"
	Grey30Background            = "\x1b[48;2;78;78;78m"
	Grey35Background            = "\x1b[48;2;88;88;88m"
	Grey37Background            = "\x1b[48;2;95;95;95m"
	Grey39Background            = "\x1b[48;2;96;96;96m"
	Grey42Background            = "\x1b[48;2;102;102;102m"
	Grey46Background            = "\x1b[48;2;118;118;118m"
	Grey50Background            = "\x1b[48;2;128;128;128m"
	Grey53Background            = "\x1b[48;2;135;135;135m"
	Grey54Background            = "\x1b[48;2;138;138;138m"
	Grey58Background            = "\x1b[48;2;148;148;148m"
	Grey62Background            = "\x1b[48;2;158;158;158m"
	Grey63Background            = "\x1b[48;2;175;135;175m"
	Grey66Background            = "\x1b[48;2;168;168;168m"
	Grey69Background            = "\x1b[48;2;175;175;175m"
	Grey70Background            = "\x1b[48;2;178;178;178m"
	Grey74Background            = "\x1b[48;2;188;188;188m"
	Grey78Background            = "\x1b[48;2;198;198;198m"
	Grey82Background            = "\x1b[48;2;208;208;208m"
	Grey84Background            = "\x1b[48;2;215;215;215m"
	Grey85Background            = "\x1b[48;2;218;218;218m"
	Grey89Background            = "\x1b[48;2;228;228;228m"
	Grey93Background            = "\x1b[48;2;238;238;238m"
	Grey100Background           = "\x1b[48;2;255;255;255m"

	// Extra ANSI escape codes
	Reset = "\x1b[0m"
	X     = "\x1b[0m"
)

// HexMap is a map of all xterm colors matched with their corresponding
// decimal values
var HexMap = map[string]int32{
	"black":             0x000000,
	"maroon":            0x800000,
	"green":             0x008000,
	"olive":             0x808000,
	"navy":              0x000080,
	"purple":            0x800080,
	"teal":              0x008080,
	"silver":            0xc0c0c0,
	"grey":              0x808080,
	"red":               0xff0000,
	"lime":              0x00ff00,
	"yellow":            0xffff00,
	"blue":              0x0000ff,
	"fuchsia":           0xff00ff,
	"aqua":              0x00ffff,
	"white":             0xffffff,
	"navajowhite1":      0xffd7af,
	"navajowhite3":      0xafaf87,
	"red1":              0xff0000,
	"red2":              0xd70000,
	"red3":              0xaf0000,
	"darkred1":          0x870000,
	"darkred2":          0x5f0000,
	"indianred1":        0xff5f87,
	"indianred2":        0xff5f5f,
	"indianred3":        0xd75f5f,
	"indianred4":        0xaf5f5f,
	"orangered1":        0xff5f00,
	"orange1":           0xffaf00,
	"orange2":           0xd78700,
	"orange3":           0x875f00,
	"orange4":           0x5f5f00,
	"darkorange1":       0xff8700,
	"darkorange2":       0xd75f00,
	"darkorange3":       0xaf5f00,
	"gold1":             0xffd700,
	"gold2":             0xd7af00,
	"gold3":             0xafaf00,
	"lightgoldenrod1":   0xffff5f,
	"lightgoldenrod2":   0xffd787,
	"lightgoldenrod3":   0xffd75f,
	"lightgoldenrod4":   0xd7d787,
	"lightgoldenrod5":   0xd7af5f,
	"darkgoldenrod":     0xaf8700,
	"yellow1":           0xffff00,
	"yellow2":           0xd7ff00,
	"yellow3":           0xd7d700,
	"yellow4":           0xafd700,
	"yellow5":           0x87af00,
	"yellow6":           0x878700,
	"lightyellow1":      0xd7d7af,
	"greenyellow":       0xafff00,
	"green1":            0x00ff00,
	"green2":            0x00d700,
	"green3":            0x00af00,
	"green4":            0x008700,
	"darkgreen":         0x005f00,
	"lightgreen1":       0x87ff87,
	"lightgreen2":       0x87ff5f,
	"palegreen1":        0xafff87,
	"palegreen2":        0x87ffaf,
	"palegreen3":        0x87d787,
	"palegreen4":        0x5fd75f,
	"darkolivegreen1":   0xd7ff87,
	"darkolivegreen2":   0xd7ff5f,
	"darkolivegreen3":   0xafff5f,
	"darkolivegreen4":   0xafd75f,
	"darkolivegreen5":   0x87d75f,
	"darkolivegreen6":   0x87af5f,
	"springgreen1":      0x00ff87,
	"springgreen2":      0x00ff5f,
	"springgreen3":      0x00d787,
	"springgreen4":      0x00d75f,
	"springgreen5":      0x00af5f,
	"springgreen6":      0x00875f,
	"mediumspringgreen": 0x00ffaf,
	"seagreen1":         0x5fffaf,
	"seagreen2":         0x5fff87,
	"seagreen3":         0x5fff5f,
	"seagreen4":         0x5fd787,
	"darkseagreen1":     0xd7ffaf,
	"darkseagreen2":     0xafffd7,
	"darkseagreen3":     0xafffaf,
	"darkseagreen4":     0xafd7af,
	"darkseagreen5":     0x87d7af,
	"darkseagreen6":     0xafd787,
	"darkseagreen7":     0x87af87,
	"darkseagreen8":     0x5faf5f,
	"darkseagreen9":     0x5f875f,
	"lightseagreen":     0x00afaf,
	"chartreuse1":       0x87ff00,
	"chartreuse2":       0x87d700,
	"chartreuse3":       0x5fff00,
	"chartreuse4":       0x5fd700,
	"chartreuse5":       0x5faf00,
	"chartreuse6":       0x5f8700,
	"blue1":             0x0000ff,
	"blue2":             0x0000af,
	"blue3":             0x0000d7,
	"darkblue":          0x000087,
	"navyblue":          0x00005f,
	"lightcoral":        0xff8787,
	"royalblue":         0x5f5fff,
	"cornflowerblue":    0x5f87ff,
	"cadetblue1":        0x5fafaf,
	"cadetblue2":        0x5faf87,
	"deepskyblue1":      0x0087af,
	"deepskyblue2":      0x0087d7,
	"deepskyblue3":      0x005faf,
	"deepskyblue4":      0x005f87,
	"deepskyblue5":      0x005f5f,
	"deepskyblue6":      0x00afd7,
	"deepskyblue7":      0x00afff,
	"dodgerblue1":       0x005fd7,
	"dodgerblue2":       0x005fff,
	"dodgerblue3":       0x0087ff,
	"slateblue1":        0x875fff,
	"slateblue2":        0x5f5fd7,
	"slateblue3":        0x5f5faf,
	"lightslateblue":    0x8787ff,
	"steelblue1":        0x5fd7ff,
	"steelblue2":        0x5fafff,
	"steelblue3":        0x5f87d7,
	"steelblue4":        0x5f87af,
	"lightsteelblue1":   0xd7d7ff,
	"lightsteelblue2":   0xafafff,
	"lightsteelblue3":   0xafafd7,
	"skyblue1":          0x87d7ff,
	"skyblue2":          0x87afff,
	"skyblue3":          0x5fafd7,
	"lightskyblue1":     0xafd7ff,
	"lightskyblue2":     0x87afd7,
	"lightskyblue3":     0x87afaf,
	"aquamarine1":       0x87ffd7,
	"aquamarine2":       0x5fffd7,
	"aquamarine3":       0x5fd7af,
	"cyan1":             0x00ffff,
	"cyan2":             0x00ffd7,
	"cyan3":             0x00d7af,
	"darkcyan":          0x00af87,
	"lightcyan1":        0xd7ffff,
	"lightcyan2":        0xafd7d7,
	"turquoise1":        0x00d7ff,
	"turquoise2":        0x008787,
	"darkturquoise":     0x00d7d7,
	"mediumturquoise":   0x5fd7d7,
	"paleturquoise1":    0xafffff,
	"paleturquoise2":    0x5f8787,
	"blueviolet":        0x5f00ff,
	"violet":            0xd787ff,
	"darkviolet1":       0xaf00d7,
	"darkviolet2":       0x8700d7,
	"palevioletred":     0xff87af,
	"mediumvioletred":   0xaf0087,
	"plum1":             0xffafff,
	"plum2":             0xd7afff,
	"plum3":             0xd787d7,
	"plum4":             0x875f87,
	"purple1":           0xaf00ff,
	"purple2":           0x8700ff,
	"purple3":           0x5f00d7,
	"purple4":           0x5f00af,
	"purple5":           0x5f0087,
	"mediumpurple1":     0xaf87ff,
	"mediumpurple2":     0xaf5fff,
	"mediumpurple3":     0xaf87d7,
	"mediumpurple4":     0x8787d7,
	"mediumpurple5":     0x875fd7,
	"mediumpurple6":     0x875faf,
	"mediumpurple7":     0x5f5f87,
	"pink1":             0xffafd7,
	"pink2":             0xd787af,
	"deeppink1":         0xff00af,
	"deeppink2":         0xff0087,
	"deeppink3":         0xff005f,
	"deeppink4":         0xd70087,
	"deeppink5":         0xd7005f,
	"deeppink6":         0x87005f,
	"deeppink7":         0x5f005f,
	"deeppink8":         0xaf005f,
	"hotpink1":          0xff5fd7,
	"hotpink2":          0xff5faf,
	"hotpink3":          0xd75faf,
	"hotpink4":          0xd75f87,
	"hotpink5":          0xaf5f87,
	"lightpink1":        0xffafaf,
	"lightpink2":        0xd78787,
	"lightpink3":        0x875f5f,
	"salmon1":           0xff875f,
	"lightsalmon1":      0xffaf87,
	"lightsalmon2":      0xd7875f,
	"lightsalmon3":      0xaf875f,
	"mistyrose1":        0xffd7d7,
	"mistyrose3":        0xd7afaf,
	"magenta1":          0xff00ff,
	"magenta2":          0xff00d7,
	"magenta3":          0xd700ff,
	"magenta4":          0xd700d7,
	"magenta5":          0xd700af,
	"magenta6":          0xaf00af,
	"darkmagenta1":      0x8700af,
	"darkmagenta2":      0x870087,
	"orchid1":           0xd75fd7,
	"orchid2":           0xff87d7,
	"orchid3":           0xff87ff,
	"mediumorchid1":     0xaf5faf,
	"mediumorchid2":     0xaf5fd7,
	"mediumorchid3":     0xd75fff,
	"mediumorchid4":     0xff5fff,
	"wheat1":            0xffffaf,
	"wheat2":            0x87875f,
	"sandybrown":        0xffaf5f,
	"rosybrown":         0xaf8787,
	"khaki1":            0xffff87,
	"khaki3":            0xd7d75f,
	"darkkhaki":         0xafaf5f,
	"tan":               0xd7af87,
	"thistle1":          0xffd7ff,
	"thistle3":          0xd7afd7,
	"honeydew2":         0xd7ffd7,
	"cornsilk1":         0xffffd7,
	"darkslategray1":    0x87ffff,
	"darkslategray2":    0x5fffff,
	"darkslategray3":    0x87d7d7,
	"lightslategrey":    0x8787af,
	"grey0":             0x000000,
	"grey3":             0x080808,
	"grey7":             0x121212,
	"grey11":            0x1c1c1c,
	"grey15":            0x262626,
	"grey19":            0x303030,
	"grey23":            0x3a3a3a,
	"grey27":            0x444444,
	"grey30":            0x4e4e4e,
	"grey35":            0x585858,
	"grey37":            0x5f5f5f,
	"grey39":            0x606060,
	"grey42":            0x666666,
	"grey46":            0x767676,
	"grey50":            0x808080,
	"grey53":            0x878787,
	"grey54":            0x8a8a8a,
	"grey58":            0x949494,
	"grey62":            0x9e9e9e,
	"grey63":            0xaf87af,
	"grey66":            0xa8a8a8,
	"grey69":            0xafafaf,
	"grey70":            0xb2b2b2,
	"grey74":            0xbcbcbc,
	"grey78":            0xc6c6c6,
	"grey82":            0xd0d0d0,
	"grey84":            0xd7d7d7,
	"grey85":            0xdadada,
	"grey89":            0xe4e4e4,
	"grey93":            0xeeeeee,
	"grey100":           0xffffff,
}

// AnsiMap is a map that contains all of the xterm color names and their
// corresponding ansi escape codes for text color
var AnsiMap = map[string]string{
	"black":             "\x1b[38;2;0;0;0m",
	"maroon":            "\x1b[38;2;128;0;0m",
	"green":             "\x1b[38;2;0;128;0m",
	"olive":             "\x1b[38;2;128;128;0m",
	"navy":              "\x1b[38;2;0;0;128m",
	"purple":            "\x1b[38;2;128;0;128m",
	"teal":              "\x1b[38;2;0;128;128m",
	"silver":            "\x1b[38;2;192;192;192m",
	"grey":              "\x1b[38;2;128;128;128m",
	"red":               "\x1b[38;2;255;0;0m",
	"lime":              "\x1b[38;2;0;255;0m",
	"yellow":            "\x1b[38;2;255;255;0m",
	"blue":              "\x1b[38;2;0;0;255m",
	"fuchsia":           "\x1b[38;2;255;0;255m",
	"aqua":              "\x1b[38;2;0;255;255m",
	"white":             "\x1b[38;2;255;255;255m",
	"navajowhite1":      "\x1b[38;2;255;215;175m",
	"navajowhite3":      "\x1b[38;2;175;175;135m",
	"red1":              "\x1b[38;2;255;0;0m",
	"red2":              "\x1b[38;2;215;0;0m",
	"red3":              "\x1b[38;2;175;0;0m",
	"darkred1":          "\x1b[38;2;135;0;0m",
	"darkred2":          "\x1b[38;2;95;0;0m",
	"indianred1":        "\x1b[38;2;255;95;135m",
	"indianred2":        "\x1b[38;2;255;95;95m",
	"indianred3":        "\x1b[38;2;215;95;95m",
	"indianred4":        "\x1b[38;2;175;95;95m",
	"orangered1":        "\x1b[38;2;255;95;0m",
	"orange1":           "\x1b[38;2;255;175;0m",
	"orange2":           "\x1b[38;2;215;135;0m",
	"orange3":           "\x1b[38;2;135;95;0m",
	"orange4":           "\x1b[38;2;95;95;0m",
	"darkorange1":       "\x1b[38;2;255;135;0m",
	"darkorange2":       "\x1b[38;2;215;95;0m",
	"darkorange3":       "\x1b[38;2;175;95;0m",
	"gold1":             "\x1b[38;2;255;215;0m",
	"gold2":             "\x1b[38;2;215;175;0m",
	"gold3":             "\x1b[38;2;175;175;0m",
	"lightgoldenrod1":   "\x1b[38;2;255;255;95m",
	"lightgoldenrod2":   "\x1b[38;2;255;215;135m",
	"lightgoldenrod3":   "\x1b[38;2;255;215;95m",
	"lightgoldenrod4":   "\x1b[38;2;215;215;135m",
	"lightgoldenrod5":   "\x1b[38;2;215;175;95m",
	"darkgoldenrod":     "\x1b[38;2;175;135;0m",
	"yellow1":           "\x1b[38;2;255;255;0m",
	"yellow2":           "\x1b[38;2;215;255;0m",
	"yellow3":           "\x1b[38;2;215;215;0m",
	"yellow4":           "\x1b[38;2;175;215;0m",
	"yellow5":           "\x1b[38;2;135;175;0m",
	"yellow6":           "\x1b[38;2;135;135;0m",
	"lightyellow1":      "\x1b[38;2;215;215;175m",
	"greenyellow":       "\x1b[38;2;175;255;0m",
	"green1":            "\x1b[38;2;0;255;0m",
	"green2":            "\x1b[38;2;0;215;0m",
	"green3":            "\x1b[38;2;0;175;0m",
	"green4":            "\x1b[38;2;0;135;0m",
	"darkgreen":         "\x1b[38;2;0;95;0m",
	"lightgreen1":       "\x1b[38;2;135;255;135m",
	"lightgreen2":       "\x1b[38;2;135;255;95m",
	"palegreen1":        "\x1b[38;2;175;255;135m",
	"palegreen2":        "\x1b[38;2;135;255;175m",
	"palegreen3":        "\x1b[38;2;135;215;135m",
	"palegreen4":        "\x1b[38;2;95;215;95m",
	"darkolivegreen1":   "\x1b[38;2;215;255;135m",
	"darkolivegreen2":   "\x1b[38;2;215;255;95m",
	"darkolivegreen3":   "\x1b[38;2;175;255;95m",
	"darkolivegreen4":   "\x1b[38;2;175;215;95m",
	"darkolivegreen5":   "\x1b[38;2;135;215;95m",
	"darkolivegreen6":   "\x1b[38;2;135;175;95m",
	"springgreen1":      "\x1b[38;2;0;255;135m",
	"springgreen2":      "\x1b[38;2;0;255;95m",
	"springgreen3":      "\x1b[38;2;0;215;135m",
	"springgreen4":      "\x1b[38;2;0;215;95m",
	"springgreen5":      "\x1b[38;2;0;175;95m",
	"springgreen6":      "\x1b[38;2;0;135;95m",
	"mediumspringgreen": "\x1b[38;2;0;255;175m",
	"seagreen1":         "\x1b[38;2;95;255;175m",
	"seagreen2":         "\x1b[38;2;95;255;135m",
	"seagreen3":         "\x1b[38;2;95;255;95m",
	"seagreen4":         "\x1b[38;2;95;215;135m",
	"darkseagreen1":     "\x1b[38;2;215;255;175m",
	"darkseagreen2":     "\x1b[38;2;175;255;215m",
	"darkseagreen3":     "\x1b[38;2;175;255;175m",
	"darkseagreen4":     "\x1b[38;2;175;215;175m",
	"darkseagreen5":     "\x1b[38;2;135;215;175m",
	"darkseagreen6":     "\x1b[38;2;175;215;135m",
	"darkseagreen7":     "\x1b[38;2;135;175;135m",
	"darkseagreen8":     "\x1b[38;2;95;175;95m",
	"darkseagreen9":     "\x1b[38;2;95;135;95m",
	"lightseagreen":     "\x1b[38;2;0;175;175m",
	"chartreuse1":       "\x1b[38;2;135;255;0m",
	"chartreuse2":       "\x1b[38;2;135;215;0m",
	"chartreuse3":       "\x1b[38;2;95;255;0m",
	"chartreuse4":       "\x1b[38;2;95;215;0m",
	"chartreuse5":       "\x1b[38;2;95;175;0m",
	"chartreuse6":       "\x1b[38;2;95;135;0m",
	"blue1":             "\x1b[38;2;0;0;255m",
	"blue2":             "\x1b[38;2;0;0;175m",
	"blue3":             "\x1b[38;2;0;0;215m",
	"darkblue":          "\x1b[38;2;0;0;135m",
	"navyblue":          "\x1b[38;2;0;0;95m",
	"lightcoral":        "\x1b[38;2;255;135;135m",
	"royalblue":         "\x1b[38;2;95;95;255m",
	"cornflowerblue":    "\x1b[38;2;95;135;255m",
	"cadetblue1":        "\x1b[38;2;95;175;175m",
	"cadetblue2":        "\x1b[38;2;95;175;135m",
	"deepskyblue1":      "\x1b[38;2;0;135;175m",
	"deepskyblue2":      "\x1b[38;2;0;135;215m",
	"deepskyblue3":      "\x1b[38;2;0;95;175m",
	"deepskyblue4":      "\x1b[38;2;0;95;135m",
	"deepskyblue5":      "\x1b[38;2;0;95;95m",
	"deepskyblue6":      "\x1b[38;2;0;175;215m",
	"deepskyblue7":      "\x1b[38;2;0;175;255m",
	"dodgerblue1":       "\x1b[38;2;0;95;215m",
	"dodgerblue2":       "\x1b[38;2;0;95;255m",
	"dodgerblue3":       "\x1b[38;2;0;135;255m",
	"slateblue1":        "\x1b[38;2;135;95;255m",
	"slateblue2":        "\x1b[38;2;95;95;215m",
	"slateblue3":        "\x1b[38;2;95;95;175m",
	"lightslateblue":    "\x1b[38;2;135;135;255m",
	"steelblue1":        "\x1b[38;2;95;215;255m",
	"steelblue2":        "\x1b[38;2;95;175;255m",
	"steelblue3":        "\x1b[38;2;95;135;215m",
	"steelblue4":        "\x1b[38;2;95;135;175m",
	"lightsteelblue1":   "\x1b[38;2;215;215;255m",
	"lightsteelblue2":   "\x1b[38;2;175;175;255m",
	"lightsteelblue3":   "\x1b[38;2;175;175;215m",
	"skyblue1":          "\x1b[38;2;135;215;255m",
	"skyblue2":          "\x1b[38;2;135;175;255m",
	"skyblue3":          "\x1b[38;2;95;175;215m",
	"lightskyblue1":     "\x1b[38;2;175;215;255m",
	"lightskyblue2":     "\x1b[38;2;135;175;215m",
	"lightskyblue3":     "\x1b[38;2;135;175;175m",
	"aquamarine1":       "\x1b[38;2;135;255;215m",
	"aquamarine2":       "\x1b[38;2;95;255;215m",
	"aquamarine3":       "\x1b[38;2;95;215;175m",
	"cyan1":             "\x1b[38;2;0;255;255m",
	"cyan2":             "\x1b[38;2;0;255;215m",
	"cyan3":             "\x1b[38;2;0;215;175m",
	"darkcyan":          "\x1b[38;2;0;175;135m",
	"lightcyan1":        "\x1b[38;2;215;255;255m",
	"lightcyan2":        "\x1b[38;2;175;215;215m",
	"turquoise1":        "\x1b[38;2;0;215;255m",
	"turquoise2":        "\x1b[38;2;0;135;135m",
	"darkturquoise":     "\x1b[38;2;0;215;215m",
	"mediumturquoise":   "\x1b[38;2;95;215;215m",
	"paleturquoise1":    "\x1b[38;2;175;255;255m",
	"paleturquoise2":    "\x1b[38;2;95;135;135m",
	"blueviolet":        "\x1b[38;2;95;0;255m",
	"violet":            "\x1b[38;2;215;135;255m",
	"darkviolet1":       "\x1b[38;2;175;0;215m",
	"darkviolet2":       "\x1b[38;2;135;0;215m",
	"palevioletred":     "\x1b[38;2;255;135;175m",
	"mediumvioletred":   "\x1b[38;2;175;0;135m",
	"plum1":             "\x1b[38;2;255;175;255m",
	"plum2":             "\x1b[38;2;215;175;255m",
	"plum3":             "\x1b[38;2;215;135;215m",
	"plum4":             "\x1b[38;2;135;95;135m",
	"purple1":           "\x1b[38;2;175;0;255m",
	"purple2":           "\x1b[38;2;135;0;255m",
	"purple3":           "\x1b[38;2;95;0;215m",
	"purple4":           "\x1b[38;2;95;0;175m",
	"purple5":           "\x1b[38;2;95;0;135m",
	"mediumpurple1":     "\x1b[38;2;175;135;255m",
	"mediumpurple2":     "\x1b[38;2;175;95;255m",
	"mediumpurple3":     "\x1b[38;2;175;135;215m",
	"mediumpurple4":     "\x1b[38;2;135;135;215m",
	"mediumpurple5":     "\x1b[38;2;135;95;215m",
	"mediumpurple6":     "\x1b[38;2;135;95;175m",
	"mediumpurple7":     "\x1b[38;2;95;95;135m",
	"pink1":             "\x1b[38;2;255;175;215m",
	"pink2":             "\x1b[38;2;215;135;175m",
	"deeppink1":         "\x1b[38;2;255;0;175m",
	"deeppink2":         "\x1b[38;2;255;0;135m",
	"deeppink3":         "\x1b[38;2;255;0;95m",
	"deeppink4":         "\x1b[38;2;215;0;135m",
	"deeppink5":         "\x1b[38;2;215;0;95m",
	"deeppink6":         "\x1b[38;2;135;0;95m",
	"deeppink7":         "\x1b[38;2;95;0;95m",
	"deeppink8":         "\x1b[38;2;175;0;95m",
	"hotpink1":          "\x1b[38;2;255;95;215m",
	"hotpink2":          "\x1b[38;2;255;95;175m",
	"hotpink3":          "\x1b[38;2;215;95;175m",
	"hotpink4":          "\x1b[38;2;215;95;135m",
	"hotpink5":          "\x1b[38;2;175;95;135m",
	"lightpink1":        "\x1b[38;2;255;175;175m",
	"lightpink2":        "\x1b[38;2;215;135;135m",
	"lightpink3":        "\x1b[38;2;135;95;95m",
	"salmon1":           "\x1b[38;2;255;135;95m",
	"lightsalmon1":      "\x1b[38;2;255;175;135m",
	"lightsalmon2":      "\x1b[38;2;215;135;95m",
	"lightsalmon3":      "\x1b[38;2;175;135;95m",
	"mistyrose1":        "\x1b[38;2;255;215;215m",
	"mistyrose3":        "\x1b[38;2;215;175;175m",
	"magenta1":          "\x1b[38;2;255;0;255m",
	"magenta2":          "\x1b[38;2;255;0;215m",
	"magenta3":          "\x1b[38;2;215;0;255m",
	"magenta4":          "\x1b[38;2;215;0;215m",
	"magenta5":          "\x1b[38;2;215;0;175m",
	"magenta6":          "\x1b[38;2;175;0;175m",
	"darkmagenta1":      "\x1b[38;2;135;0;175m",
	"darkmagenta2":      "\x1b[38;2;135;0;135m",
	"orchid1":           "\x1b[38;2;215;95;215m",
	"orchid2":           "\x1b[38;2;255;135;215m",
	"orchid3":           "\x1b[38;2;255;135;255m",
	"mediumorchid1":     "\x1b[38;2;175;95;175m",
	"mediumorchid2":     "\x1b[38;2;175;95;215m",
	"mediumorchid3":     "\x1b[38;2;215;95;255m",
	"mediumorchid4":     "\x1b[38;2;255;95;255m",
	"wheat1":            "\x1b[38;2;255;255;175m",
	"wheat2":            "\x1b[38;2;135;135;95m",
	"sandybrown":        "\x1b[38;2;255;175;95m",
	"rosybrown":         "\x1b[38;2;175;135;135m",
	"khaki1":            "\x1b[38;2;255;255;135m",
	"khaki3":            "\x1b[38;2;215;215;95m",
	"darkkhaki":         "\x1b[38;2;175;175;95m",
	"tan":               "\x1b[38;2;215;175;135m",
	"thistle1":          "\x1b[38;2;255;215;255m",
	"thistle3":          "\x1b[38;2;215;175;215m",
	"honeydew2":         "\x1b[38;2;215;255;215m",
	"cornsilk1":         "\x1b[38;2;255;255;215m",
	"darkslategray1":    "\x1b[38;2;135;255;255m",
	"darkslategray2":    "\x1b[38;2;95;255;255m",
	"darkslategray3":    "\x1b[38;2;135;215;215m",
	"lightslategrey":    "\x1b[38;2;135;135;175m",
	"grey0":             "\x1b[38;2;0;0;0m",
	"grey3":             "\x1b[38;2;8;8;8m",
	"grey7":             "\x1b[38;2;18;18;18m",
	"grey11":            "\x1b[38;2;28;28;28m",
	"grey15":            "\x1b[38;2;38;38;38m",
	"grey19":            "\x1b[38;2;48;48;48m",
	"grey23":            "\x1b[38;2;58;58;58m",
	"grey27":            "\x1b[38;2;68;68;68m",
	"grey30":            "\x1b[38;2;78;78;78m",
	"grey35":            "\x1b[38;2;88;88;88m",
	"grey37":            "\x1b[38;2;95;95;95m",
	"grey39":            "\x1b[38;2;96;96;96m",
	"grey42":            "\x1b[38;2;102;102;102m",
	"grey46":            "\x1b[38;2;118;118;118m",
	"grey50":            "\x1b[38;2;128;128;128m",
	"grey53":            "\x1b[38;2;135;135;135m",
	"grey54":            "\x1b[38;2;138;138;138m",
	"grey58":            "\x1b[38;2;148;148;148m",
	"grey62":            "\x1b[38;2;158;158;158m",
	"grey63":            "\x1b[38;2;175;135;175m",
	"grey66":            "\x1b[38;2;168;168;168m",
	"grey69":            "\x1b[38;2;175;175;175m",
	"grey70":            "\x1b[38;2;178;178;178m",
	"grey74":            "\x1b[38;2;188;188;188m",
	"grey78":            "\x1b[38;2;198;198;198m",
	"grey82":            "\x1b[38;2;208;208;208m",
	"grey84":            "\x1b[38;2;215;215;215m",
	"grey85":            "\x1b[38;2;218;218;218m",
	"grey89":            "\x1b[38;2;228;228;228m",
	"grey93":            "\x1b[38;2;238;238;238m",
	"grey100":           "\x1b[38;2;255;255;255m",
}

// AnsiBackgroundMap is a map that contains all of the xterm color names
// and their corresponding ansi escape codes for background color
var AnsiBackgroundMap = map[string]string{
	"blackbackground":             "\x1b[48;2;0;0;0m",
	"maroonbackground":            "\x1b[48;2;128;0;0m",
	"greenbackground":             "\x1b[48;2;0;128;0m",
	"olivebackground":             "\x1b[48;2;128;128;0m",
	"navybackground":              "\x1b[48;2;0;0;128m",
	"purplebackground":            "\x1b[48;2;128;0;128m",
	"tealbackground":              "\x1b[48;2;0;128;128m",
	"silverbackground":            "\x1b[48;2;192;192;192m",
	"greybackground":              "\x1b[48;2;128;128;128m",
	"redbackground":               "\x1b[48;2;255;0;0m",
	"limebackground":              "\x1b[48;2;0;255;0m",
	"yellowbackground":            "\x1b[48;2;255;255;0m",
	"bluebackground":              "\x1b[48;2;0;0;255m",
	"fuchsiabackground":           "\x1b[48;2;255;0;255m",
	"aquabackground":              "\x1b[48;2;0;255;255m",
	"whitebackground":             "\x1b[48;2;255;255;255m",
	"navajowhite1background":      "\x1b[48;2;255;215;175m",
	"navajowhite3background":      "\x1b[48;2;175;175;135m",
	"red1background":              "\x1b[48;2;255;0;0m",
	"red2background":              "\x1b[48;2;215;0;0m",
	"red3background":              "\x1b[48;2;175;0;0m",
	"darkred1background":          "\x1b[48;2;135;0;0m",
	"darkred2background":          "\x1b[48;2;95;0;0m",
	"indianred1background":        "\x1b[48;2;255;95;135m",
	"indianred2background":        "\x1b[48;2;255;95;95m",
	"indianred3background":        "\x1b[48;2;215;95;95m",
	"indianred4background":        "\x1b[48;2;175;95;95m",
	"orangered1background":        "\x1b[48;2;255;95;0m",
	"orange1background":           "\x1b[48;2;255;175;0m",
	"orange2background":           "\x1b[48;2;215;135;0m",
	"orange3background":           "\x1b[48;2;135;95;0m",
	"orange4background":           "\x1b[48;2;95;95;0m",
	"darkorange1background":       "\x1b[48;2;255;135;0m",
	"darkorange2background":       "\x1b[48;2;215;95;0m",
	"darkorange3background":       "\x1b[48;2;175;95;0m",
	"gold1background":             "\x1b[48;2;255;215;0m",
	"gold2background":             "\x1b[48;2;215;175;0m",
	"gold3background":             "\x1b[48;2;175;175;0m",
	"lightgoldenrod1background":   "\x1b[48;2;255;255;95m",
	"lightgoldenrod2background":   "\x1b[48;2;255;215;135m",
	"lightgoldenrod3background":   "\x1b[48;2;255;215;95m",
	"lightgoldenrod4background":   "\x1b[48;2;215;215;135m",
	"lightgoldenrod5background":   "\x1b[48;2;215;175;95m",
	"darkgoldenrodbackground":     "\x1b[48;2;175;135;0m",
	"yellow1background":           "\x1b[48;2;255;255;0m",
	"yellow2background":           "\x1b[48;2;215;255;0m",
	"yellow3background":           "\x1b[48;2;215;215;0m",
	"yellow4background":           "\x1b[48;2;175;215;0m",
	"yellow5background":           "\x1b[48;2;135;175;0m",
	"yellow6background":           "\x1b[48;2;135;135;0m",
	"lightyellow1background":      "\x1b[48;2;215;215;175m",
	"greenyellowbackground":       "\x1b[48;2;175;255;0m",
	"green1background":            "\x1b[48;2;0;255;0m",
	"green2background":            "\x1b[48;2;0;215;0m",
	"green3background":            "\x1b[48;2;0;175;0m",
	"green4background":            "\x1b[48;2;0;135;0m",
	"darkgreenbackground":         "\x1b[48;2;0;95;0m",
	"lightgreen1background":       "\x1b[48;2;135;255;135m",
	"lightgreen2background":       "\x1b[48;2;135;255;95m",
	"palegreen1background":        "\x1b[48;2;175;255;135m",
	"palegreen2background":        "\x1b[48;2;135;255;175m",
	"palegreen3background":        "\x1b[48;2;135;215;135m",
	"palegreen4background":        "\x1b[48;2;95;215;95m",
	"darkolivegreen1background":   "\x1b[48;2;215;255;135m",
	"darkolivegreen2background":   "\x1b[48;2;215;255;95m",
	"darkolivegreen3background":   "\x1b[48;2;175;255;95m",
	"darkolivegreen4background":   "\x1b[48;2;175;215;95m",
	"darkolivegreen5background":   "\x1b[48;2;135;215;95m",
	"darkolivegreen6background":   "\x1b[48;2;135;175;95m",
	"springgreen1background":      "\x1b[48;2;0;255;135m",
	"springgreen2background":      "\x1b[48;2;0;255;95m",
	"springgreen3background":      "\x1b[48;2;0;215;135m",
	"springgreen4background":      "\x1b[48;2;0;215;95m",
	"springgreen5background":      "\x1b[48;2;0;175;95m",
	"springgreen6background":      "\x1b[48;2;0;135;95m",
	"mediumspringgreenbackground": "\x1b[48;2;0;255;175m",
	"seagreen1background":         "\x1b[48;2;95;255;175m",
	"seagreen2background":         "\x1b[48;2;95;255;135m",
	"seagreen3background":         "\x1b[48;2;95;255;95m",
	"seagreen4background":         "\x1b[48;2;95;215;135m",
	"darkseagreen1background":     "\x1b[48;2;215;255;175m",
	"darkseagreen2background":     "\x1b[48;2;175;255;215m",
	"darkseagreen3background":     "\x1b[48;2;175;255;175m",
	"darkseagreen4background":     "\x1b[48;2;175;215;175m",
	"darkseagreen5background":     "\x1b[48;2;135;215;175m",
	"darkseagreen6background":     "\x1b[48;2;175;215;135m",
	"darkseagreen7background":     "\x1b[48;2;135;175;135m",
	"darkseagreen8background":     "\x1b[48;2;95;175;95m",
	"darkseagreen9background":     "\x1b[48;2;95;135;95m",
	"lightseagreenbackground":     "\x1b[48;2;0;175;175m",
	"chartreuse1background":       "\x1b[48;2;135;255;0m",
	"chartreuse2background":       "\x1b[48;2;135;215;0m",
	"chartreuse3background":       "\x1b[48;2;95;255;0m",
	"chartreuse4background":       "\x1b[48;2;95;215;0m",
	"chartreuse5background":       "\x1b[48;2;95;175;0m",
	"chartreuse6background":       "\x1b[48;2;95;135;0m",
	"blue1background":             "\x1b[48;2;0;0;255m",
	"blue2background":             "\x1b[48;2;0;0;175m",
	"blue3background":             "\x1b[48;2;0;0;215m",
	"darkbluebackground":          "\x1b[48;2;0;0;135m",
	"navybluebackground":          "\x1b[48;2;0;0;95m",
	"lightcoralbackground":        "\x1b[48;2;255;135;135m",
	"royalbluebackground":         "\x1b[48;2;95;95;255m",
	"cornflowerbluebackground":    "\x1b[48;2;95;135;255m",
	"cadetblue1background":        "\x1b[48;2;95;175;175m",
	"cadetblue2background":        "\x1b[48;2;95;175;135m",
	"deepskyblue1background":      "\x1b[48;2;0;135;175m",
	"deepskyblue2background":      "\x1b[48;2;0;135;215m",
	"deepskyblue3background":      "\x1b[48;2;0;95;175m",
	"deepskyblue4background":      "\x1b[48;2;0;95;135m",
	"deepskyblue5background":      "\x1b[48;2;0;95;95m",
	"deepskyblue6background":      "\x1b[48;2;0;175;215m",
	"deepskyblue7background":      "\x1b[48;2;0;175;255m",
	"dodgerblue1background":       "\x1b[48;2;0;95;215m",
	"dodgerblue2background":       "\x1b[48;2;0;95;255m",
	"dodgerblue3background":       "\x1b[48;2;0;135;255m",
	"slateblue1background":        "\x1b[48;2;135;95;255m",
	"slateblue2background":        "\x1b[48;2;95;95;215m",
	"slateblue3background":        "\x1b[48;2;95;95;175m",
	"lightslatebluebackground":    "\x1b[48;2;135;135;255m",
	"steelblue1background":        "\x1b[48;2;95;215;255m",
	"steelblue2background":        "\x1b[48;2;95;175;255m",
	"steelblue3background":        "\x1b[48;2;95;135;215m",
	"steelblue4background":        "\x1b[48;2;95;135;175m",
	"lightsteelblue1background":   "\x1b[48;2;215;215;255m",
	"lightsteelblue2background":   "\x1b[48;2;175;175;255m",
	"lightsteelblue3background":   "\x1b[48;2;175;175;215m",
	"skyblue1background":          "\x1b[48;2;135;215;255m",
	"skyblue2background":          "\x1b[48;2;135;175;255m",
	"skyblue3background":          "\x1b[48;2;95;175;215m",
	"lightskyblue1background":     "\x1b[48;2;175;215;255m",
	"lightskyblue2background":     "\x1b[48;2;135;175;215m",
	"lightskyblue3background":     "\x1b[48;2;135;175;175m",
	"aquamarine1background":       "\x1b[48;2;135;255;215m",
	"aquamarine2background":       "\x1b[48;2;95;255;215m",
	"aquamarine3background":       "\x1b[48;2;95;215;175m",
	"cyan1background":             "\x1b[48;2;0;255;255m",
	"cyan2background":             "\x1b[48;2;0;255;215m",
	"cyan3background":             "\x1b[48;2;0;215;175m",
	"darkcyanbackground":          "\x1b[48;2;0;175;135m",
	"lightcyan1background":        "\x1b[48;2;215;255;255m",
	"lightcyan2background":        "\x1b[48;2;175;215;215m",
	"turquoise1background":        "\x1b[48;2;0;215;255m",
	"turquoise2background":        "\x1b[48;2;0;135;135m",
	"darkturquoisebackground":     "\x1b[48;2;0;215;215m",
	"mediumturquoisebackground":   "\x1b[48;2;95;215;215m",
	"paleturquoise1background":    "\x1b[48;2;175;255;255m",
	"paleturquoise2background":    "\x1b[48;2;95;135;135m",
	"bluevioletbackground":        "\x1b[48;2;95;0;255m",
	"violetbackground":            "\x1b[48;2;215;135;255m",
	"darkviolet1background":       "\x1b[48;2;175;0;215m",
	"darkviolet2background":       "\x1b[48;2;135;0;215m",
	"palevioletredbackground":     "\x1b[48;2;255;135;175m",
	"mediumvioletredbackground":   "\x1b[48;2;175;0;135m",
	"plum1background":             "\x1b[48;2;255;175;255m",
	"plum2background":             "\x1b[48;2;215;175;255m",
	"plum3background":             "\x1b[48;2;215;135;215m",
	"plum4background":             "\x1b[48;2;135;95;135m",
	"purple1background":           "\x1b[48;2;175;0;255m",
	"purple2background":           "\x1b[48;2;135;0;255m",
	"purple3background":           "\x1b[48;2;95;0;215m",
	"purple4background":           "\x1b[48;2;95;0;175m",
	"purple5background":           "\x1b[48;2;95;0;135m",
	"mediumpurple1background":     "\x1b[48;2;175;135;255m",
	"mediumpurple2background":     "\x1b[48;2;175;95;255m",
	"mediumpurple3background":     "\x1b[48;2;175;135;215m",
	"mediumpurple4background":     "\x1b[48;2;135;135;215m",
	"mediumpurple5background":     "\x1b[48;2;135;95;215m",
	"mediumpurple6background":     "\x1b[48;2;135;95;175m",
	"mediumpurple7background":     "\x1b[48;2;95;95;135m",
	"pink1background":             "\x1b[48;2;255;175;215m",
	"pink2background":             "\x1b[48;2;215;135;175m",
	"deeppink1background":         "\x1b[48;2;255;0;175m",
	"deeppink2background":         "\x1b[48;2;255;0;135m",
	"deeppink3background":         "\x1b[48;2;255;0;95m",
	"deeppink4background":         "\x1b[48;2;215;0;135m",
	"deeppink5background":         "\x1b[48;2;215;0;95m",
	"deeppink6background":         "\x1b[48;2;135;0;95m",
	"deeppink7background":         "\x1b[48;2;95;0;95m",
	"deeppink8background":         "\x1b[48;2;175;0;95m",
	"hotpink1background":          "\x1b[48;2;255;95;215m",
	"hotpink2background":          "\x1b[48;2;255;95;175m",
	"hotpink3background":          "\x1b[48;2;215;95;175m",
	"hotpink4background":          "\x1b[48;2;215;95;135m",
	"hotpink5background":          "\x1b[48;2;175;95;135m",
	"lightpink1background":        "\x1b[48;2;255;175;175m",
	"lightpink2background":        "\x1b[48;2;215;135;135m",
	"lightpink3background":        "\x1b[48;2;135;95;95m",
	"salmon1background":           "\x1b[48;2;255;135;95m",
	"lightsalmon1background":      "\x1b[48;2;255;175;135m",
	"lightsalmon2background":      "\x1b[48;2;215;135;95m",
	"lightsalmon3background":      "\x1b[48;2;175;135;95m",
	"mistyrose1background":        "\x1b[48;2;255;215;215m",
	"mistyrose3background":        "\x1b[48;2;215;175;175m",
	"magenta1background":          "\x1b[48;2;255;0;255m",
	"magenta2background":          "\x1b[48;2;255;0;215m",
	"magenta3background":          "\x1b[48;2;215;0;255m",
	"magenta4background":          "\x1b[48;2;215;0;215m",
	"magenta5background":          "\x1b[48;2;215;0;175m",
	"magenta6background":          "\x1b[48;2;175;0;175m",
	"darkmagenta1background":      "\x1b[48;2;135;0;175m",
	"darkmagenta2background":      "\x1b[48;2;135;0;135m",
	"orchid1background":           "\x1b[48;2;215;95;215m",
	"orchid2background":           "\x1b[48;2;255;135;215m",
	"orchid3background":           "\x1b[48;2;255;135;255m",
	"mediumorchid1background":     "\x1b[48;2;175;95;175m",
	"mediumorchid2background":     "\x1b[48;2;175;95;215m",
	"mediumorchid3background":     "\x1b[48;2;215;95;255m",
	"mediumorchid4background":     "\x1b[48;2;255;95;255m",
	"wheat1background":            "\x1b[48;2;255;255;175m",
	"wheat2background":            "\x1b[48;2;135;135;95m",
	"sandybrownbackground":        "\x1b[48;2;255;175;95m",
	"rosybrownbackground":         "\x1b[48;2;175;135;135m",
	"khaki1background":            "\x1b[48;2;255;255;135m",
	"khaki3background":            "\x1b[48;2;215;215;95m",
	"darkkhakibackground":         "\x1b[48;2;175;175;95m",
	"tanbackground":               "\x1b[48;2;215;175;135m",
	"thistle1background":          "\x1b[48;2;255;215;255m",
	"thistle3background":          "\x1b[48;2;215;175;215m",
	"honeydew2background":         "\x1b[48;2;215;255;215m",
	"cornsilk1background":         "\x1b[48;2;255;255;215m",
	"darkslategray1background":    "\x1b[48;2;135;255;255m",
	"darkslategray2background":    "\x1b[48;2;95;255;255m",
	"darkslategray3background":    "\x1b[48;2;135;215;215m",
	"lightslategreybackground":    "\x1b[48;2;135;135;175m",
	"grey0background":             "\x1b[48;2;0;0;0m",
	"grey3background":             "\x1b[48;2;8;8;8m",
	"grey7background":             "\x1b[48;2;18;18;18m",
	"grey11background":            "\x1b[48;2;28;28;28m",
	"grey15background":            "\x1b[48;2;38;38;38m",
	"grey19background":            "\x1b[48;2;48;48;48m",
	"grey23background":            "\x1b[48;2;58;58;58m",
	"grey27background":            "\x1b[48;2;68;68;68m",
	"grey30background":            "\x1b[48;2;78;78;78m",
	"grey35background":            "\x1b[48;2;88;88;88m",
	"grey37background":            "\x1b[48;2;95;95;95m",
	"grey39background":            "\x1b[48;2;96;96;96m",
	"grey42background":            "\x1b[48;2;102;102;102m",
	"grey46background":            "\x1b[48;2;118;118;118m",
	"grey50background":            "\x1b[48;2;128;128;128m",
	"grey53background":            "\x1b[48;2;135;135;135m",
	"grey54background":            "\x1b[48;2;138;138;138m",
	"grey58background":            "\x1b[48;2;148;148;148m",
	"grey62background":            "\x1b[48;2;158;158;158m",
	"grey63background":            "\x1b[48;2;175;135;175m",
	"grey66background":            "\x1b[48;2;168;168;168m",
	"grey69background":            "\x1b[48;2;175;175;175m",
	"grey70background":            "\x1b[48;2;178;178;178m",
	"grey74background":            "\x1b[48;2;188;188;188m",
	"grey78background":            "\x1b[48;2;198;198;198m",
	"grey82background":            "\x1b[48;2;208;208;208m",
	"grey84background":            "\x1b[48;2;215;215;215m",
	"grey85background":            "\x1b[48;2;218;218;218m",
	"grey89background":            "\x1b[48;2;228;228;228m",
	"grey93background":            "\x1b[48;2;238;238;238m",
	"grey100background":           "\x1b[48;2;255;255;255m",
}

/* Color Functions */

// Random returns a random css color as an ANSI escape code
func Random() string {
	return pick(AnsiMap)
}

// RandomBackground returns a random css color as an ANSI background escape code
func RandomBackground() string {
	return pick(AnsiBackgroundMap)
}

// RandomHex returns a random css color as a hexadecimal value
func RandomHex() int32 {
	return pick(HexMap)
}

// Multiple colorizes every rune in a string randomly from the AnsiColors
func Multiple(s string) string {
	var m string
	for _, r := range s {
		m += Random() + string(r)
	}
	m += Reset
	return m
}

// MultipleBackground colorizes every rune's background in a string randomly from the AnsiColors
func MultipleBackground(s string) string {
	var m string
	for _, r := range s {
		m += RandomBackground() + string(r)
	}
	m += Reset
	return m
}

// Helper function to choose a random element from a map
func pick[K comparable, V any](m map[K]V) V {
	k := rand.Intn(len(m))
	i := 0

	for _, x := range m {
		if i == k {
			return x
		}
		i++
	}

	panic("unreachable")
}
