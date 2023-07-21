// Package css provides all the css colors.
package css

import (
	"math/rand"
)

/* ANSI Color Value Constants */

const (
	AliceBlue            = "\x1b[38;2;240;248;255m"
	AntiqueWhite         = "\x1b[38;2;250;235;215m"
	Aqua                 = "\x1b[38;2;0;255;255m"
	Aquamarine           = "\x1b[38;2;127;255;212m"
	Azure                = "\x1b[38;2;1240;255;255m"
	Beige                = "\x1b[38;2;245;245;220m"
	Bisque               = "\x1b[38;2;255;228;196m"
	Black                = "\x1b[38;2;0;0;0m"
	BlanchedAlmond       = "\x1b[38;2;255;235;205m"
	Blue                 = "\x1b[38;2;0;0;255m"
	BlueViolet           = "\x1b[38;2;138;43;226m"
	Brown                = "\x1b[38;2;165;42;42m"
	BurlyWood            = "\x1b[38;2;222;184;135m"
	CadetBlue            = "\x1b[38;2;95;158;160m"
	Chartreuse           = "\x1b[38;2;95;158;160m"
	Chocolate            = "\x1b[38;2;210;105;30m"
	Coral                = "\x1b[38;2;255;127;80m"
	CornflowerBlue       = "\x1b[38;2;100;149;237m"
	Cornsilk             = "\x1b[38;2;255;248;220m"
	Crimson              = "\x1b[38;2;220;20;60m"
	Cyan                 = "\x1b[38;2;0;255;255m"
	DarkBlue             = "\x1b[38;2;0;0;139m"
	DarkCyan             = "\x1b[38;2;0;139;139m"
	DarkGoldenrod        = "\x1b[38;2;184;134;11m"
	DarkGray             = "\x1b[38;2;169;169;169m"
	DarkGreen            = "\x1b[38;2;0;100;0m"
	DarkKhaki            = "\x1b[38;2;189;183;107m"
	DarkMagenta          = "\x1b[38;2;139;0;139m"
	DarkOliveGreen       = "\x1b[38;2;85;107;47m"
	DarkOrange           = "\x1b[38;2;255;140;0m"
	DarkOrchid           = "\x1b[38;2;153;50;204m"
	DarkRed              = "\x1b[38;2;139;0;0m"
	DarkSalmon           = "\x1b[38;2;233;150;122m"
	DarkSeagreen         = "\x1b[38;2;143;188;143m"
	DarkSlateBlue        = "\x1b[38;2;72;61;139m"
	DarkSlateGray        = "\x1b[38;2;47;79;79m"
	DarkTurquoise        = "\x1b[38;2;0;206;209m"
	DarkViolet           = "\x1b[38;2;148;0;211m"
	DeepPink             = "\x1b[38;2;255;20;147m"
	DeepSkyBlue          = "\x1b[38;2;0;191;255m"
	DimGray              = "\x1b[38;2;0;191;255m"
	DodgerBlue           = "\x1b[38;2;30;144;255m"
	FireBrick            = "\x1b[38;2;178;34;34m"
	FloralWhite          = "\x1b[38;2;255;250;240m"
	ForestGreen          = "\x1b[38;2;34;139;34m"
	Fuchsia              = "\x1b[38;2;255;0;255m"
	Gainsboro            = "\x1b[38;2;220;220;220m"
	GhostWhite           = "\x1b[38;2;248;248;255m"
	Gold                 = "\x1b[38;2;255;215;0m"
	GoldenRod            = "\x1b[38;2;218;165;32m"
	Gray                 = "\x1b[38;2;127;127;127m"
	Green                = "\x1b[38;2;0;128;0m"
	GreenYellow          = "\x1b[38;2;173;255;47m"
	Honeydew             = "\x1b[38;2;240;255;240m"
	HotPink              = "\x1b[38;2;255;105;180m"
	IndianRed            = "\x1b[38;2;205;92;92m"
	Indigo               = "\x1b[38;2;75;0;130m"
	Ivory                = "\x1b[38;2;255;255;240m"
	Khaki                = "\x1b[38;2;240;230;140m"
	Lavender             = "\x1b[38;2;230;230;250m"
	LavenderBlush        = "\x1b[38;2;255;240;245m"
	LawnGreen            = "\x1b[38;2;124;252;0m"
	LemonChiffon         = "\x1b[38;2;255;250;205m"
	LightBlue            = "\x1b[38;2;173;216;230m"
	LightCoral           = "\x1b[38;2;240;128;128m"
	LightCyan            = "\x1b[38;2;224;255;255m"
	LightGoldenRodYellow = "\x1b[38;2;250;250;210m"
	LightGreen           = "\x1b[38;2;144;238;144m"
	LightGray            = "\x1b[38;2;211;211;211m"
	LightPink            = "\x1b[38;2;255;182;193m"
	LightSalmon          = "\x1b[38;2;255;160;122m"
	LightSeaGreen        = "\x1b[38;2;32;178;170m"
	LightSkyBlue         = "\x1b[38;2;135;206;250m"
	LightSlateGray       = "\x1b[38;2;119;136;153m"
	LightSteelBlue       = "\x1b[38;2;176;196;222m"
	LightYellow          = "\x1b[38;2;255;255;224m"
	Lime                 = "\x1b[38;2;0;255;0m"
	LimeGreen            = "\x1b[38;2;50;205;50m"
	Linen                = "\x1b[38;2;250;240;230m"
	Magenta              = "\x1b[38;2;255;0;255m"
	Maroon               = "\x1b[38;2;128;0;0m"
	MediumAquamarine     = "\x1b[38;2;102;205;170m"
	MediumBlue           = "\x1b[38;2;0;0;205m"
	MediumOrchid         = "\x1b[38;2;186;85;211m"
	MediumPurple         = "\x1b[38;2;147;112;219m"
	MediumSeaGreen       = "\x1b[38;2;60;179;113m"
	MediumSlateBlue      = "\x1b[38;2;123;104;238m"
	MediumSpringGreen    = "\x1b[38;2;0;250;154m"
	MediumTurquoise      = "\x1b[38;2;72;209;204m"
	MediumVioletRed      = "\x1b[38;2;199;21;133m"
	MidnightBlue         = "\x1b[38;2;25;25;112m"
	MintCream            = "\x1b[38;2;245;255;250m"
	MistyRose            = "\x1b[38;2;255;228;225m"
	Moccasin             = "\x1b[38;2;255;228;181m"
	NavajoWhite          = "\x1b[38;2;255;222;173m"
	Navy                 = "\x1b[38;2;0;0;128m"
	NavyBlue             = "\x1b[38;2;159;175;223m"
	OldLace              = "\x1b[38;2;253;245;230m"
	Olive                = "\x1b[38;2;128;128;0m"
	OliveDrab            = "\x1b[38;2;107;142;35m"
	Orange               = "\x1b[38;2;255;165;0m"
	OrangeRed            = "\x1b[38;2;255;69;0m"
	Orchid               = "\x1b[38;2;218;112;214m"
	PaleGoldenRod        = "\x1b[38;2;238;232;170m"
	PaleGreen            = "\x1b[38;2;152;251;152m"
	PaleTurquoise        = "\x1b[38;2;175;238;238m"
	PaleVioletRed        = "\x1b[38;2;219;112;147m"
	PapayaWhip           = "\x1b[38;2;255;239;213m"
	PeachPuff            = "\x1b[38;2;255;218;185m"
	Peru                 = "\x1b[38;2;205;133;63m"
	Pink                 = "\x1b[38;2;255;192;203m"
	Plum                 = "\x1b[38;2;221;160;221m"
	PowderBlue           = "\x1b[38;2;176;224;230m"
	Purple               = "\x1b[38;2;128;0;128m"
	RebeccaPurple        = "\x1b[38:2;102;51;153m"
	Red                  = "\x1b[38;2;255;0;0m"
	RosyBrown            = "\x1b[38;2;188;143;143m"
	RoyalBlue            = "\x1b[38;2;65;105;225m"
	SaddleBrown          = "\x1b[38;2;139;69;19m"
	Salmon               = "\x1b[38;2;250;128;114m"
	SandyBrown           = "\x1b[38;2;244;164;96m"
	SeaGreen             = "\x1b[38;2;46;139;87m"
	Seashell             = "\x1b[38;2;255;245;238m"
	Sienna               = "\x1b[38;2;160;82;45m"
	Silver               = "\x1b[38;2;192;192;192m"
	SkyBlue              = "\x1b[38;2;135;206;235m"
	SlateBlue            = "\x1b[38;2;106;90;205m"
	SlateGray            = "\x1b[38;2;112;128;144m"
	Snow                 = "\x1b[38;2;255;250;250m"
	SpringGreen          = "\x1b[38;2;0;255;127m"
	SteelBlue            = "\x1b[38;2;70;130;180m"
	Tan                  = "\x1b[38;2;210;180;140m"
	Teal                 = "\x1b[38;2;0;128;128m"
	Thistle              = "\x1b[38;2;216;191;216m"
	Tomato               = "\x1b[38;2;255;99;71m"
	Turquoise            = "\x1b[38;2;64;224;208m"
	Violet               = "\x1b[38;2;238;130;238m"
	Wheat                = "\x1b[38;2;245;222;179m"
	White                = "\x1b[38;2;255;255;255m"
	WhiteSmoke           = "\x1b[38;2;245;245;245m"
	Yellow               = "\x1b[38;2;255;255;0m"
	YellowGreen          = "\x1b[38;2;139;205;50m"

	AliceBlueBackground            = "\x1b[48;2;240;248;255m"
	AntiqueWhiteBackground         = "\x1b[48;2;250;235;215m"
	AquaBackground                 = "\x1b[48;2;0;255;255m"
	AquamarineBackground           = "\x1b[48;2;127;255;212m"
	AzureBackground                = "\x1b[48;2;1240;255;255m"
	BeigeBackground                = "\x1b[48;2;245;245;220m"
	BisqueBackground               = "\x1b[48;2;255;228;196m"
	BlackBackground                = "\x1b[48;2;0;0;0m"
	BlanchedAlmondBackground       = "\x1b[48;2;255;235;205m"
	BlueBackground                 = "\x1b[48;2;0;0;255m"
	BlueVioletBackground           = "\x1b[48;2;138;43;226m"
	BrownBackground                = "\x1b[48;2;165;42;42m"
	BurlyWoodBackground            = "\x1b[48;2;222;184;135m"
	CadetBlueBackground            = "\x1b[48;2;95;158;160m"
	ChartreuseBackground           = "\x1b[48;2;95;158;160m"
	ChocolateBackground            = "\x1b[48;2;210;105;30m"
	CoralBackground                = "\x1b[48;2;255;127;80m"
	CornflowerBlueBackground       = "\x1b[48;2;100;149;237m"
	CornsilkBackground             = "\x1b[48;2;255;248;220m"
	CrimsonBackground              = "\x1b[48;2;220;20;60m"
	CyanBackground                 = "\x1b[48;2;0;255;255m"
	DarkBlueBackground             = "\x1b[48;2;0;0;139m"
	DarkCyanBackground             = "\x1b[48;2;0;139;139m"
	DarkGoldenrodBackground        = "\x1b[48;2;184;134;11m"
	DarkGrayBackground             = "\x1b[48;2;169;169;169m"
	DarkGreenBackground            = "\x1b[48;2;0;100;0m"
	DarkKhakiBackground            = "\x1b[48;2;189;183;107m"
	DarkMagentaBackground          = "\x1b[48;2;139;0;139m"
	DarkOliveGreenBackground       = "\x1b[48;2;85;107;47m"
	DarkOrangeBackground           = "\x1b[48;2;255;140;0m"
	DarkOrchidBackground           = "\x1b[48;2;153;50;204m"
	DarkRedBackground              = "\x1b[48;2;139;0;0m"
	DarkSalmonBackground           = "\x1b[48;2;233;150;122m"
	DarkSeagreenBackground         = "\x1b[48;2;143;188;143m"
	DarkSlateBlueBackground        = "\x1b[48;2;72;61;139m"
	DarkSlateGrayBackground        = "\x1b[48;2;47;79;79m"
	DarkTurquoiseBackground        = "\x1b[48;2;0;206;209m"
	DarkVioletBackground           = "\x1b[48;2;148;0;211m"
	DeepPinkBackground             = "\x1b[48;2;255;20;147m"
	DeepSkyBlueBackground          = "\x1b[48;2;0;191;255m"
	DimGrayBackground              = "\x1b[48;2;0;191;255m"
	DodgerBlueBackground           = "\x1b[48;2;30;144;255m"
	FireBrickBackground            = "\x1b[48;2;178;34;34m"
	FloralWhiteBackground          = "\x1b[48;2;255;250;240m"
	ForestGreenBackground          = "\x1b[48;2;34;139;34m"
	FuchsiaBackground              = "\x1b[48;2;255;0;255m"
	GainsboroBackground            = "\x1b[48;2;220;220;220m"
	GhostWhiteBackground           = "\x1b[48;2;248;248;255m"
	GoldBackground                 = "\x1b[48;2;255;215;0m"
	GoldenRodBackground            = "\x1b[48;2;218;165;32m"
	GrayBackground                 = "\x1b[48;2;127;127;127m"
	GreenBackground                = "\x1b[48;2;0;128;0m"
	GreenYellowBackground          = "\x1b[48;2;173;255;47m"
	HoneydewBackground             = "\x1b[48;2;240;255;240m"
	HotPinkBackground              = "\x1b[48;2;255;105;180m"
	IndianRedBackground            = "\x1b[48;2;205;92;92m"
	IndigoBackground               = "\x1b[48;2;75;0;130m"
	IvoryBackground                = "\x1b[48;2;255;255;240m"
	KhakiBackground                = "\x1b[48;2;240;230;140m"
	LavenderBackground             = "\x1b[48;2;230;230;250m"
	LavenderBlushBackground        = "\x1b[48;2;255;240;245m"
	LawnGreenBackground            = "\x1b[48;2;124;252;0m"
	LemonChiffonBackground         = "\x1b[48;2;255;250;205m"
	LightBlueBackground            = "\x1b[48;2;173;216;230m"
	LightCoralBackground           = "\x1b[48;2;240;128;128m"
	LightCyanBackground            = "\x1b[48;2;224;255;255m"
	LightGoldenRodYellowBackground = "\x1b[48;2;250;250;210m"
	LightGreenBackground           = "\x1b[48;2;144;238;144m"
	LightGrayBackground            = "\x1b[48;2;211;211;211m"
	LightPinkBackground            = "\x1b[48;2;255;182;193m"
	LightSalmonBackground          = "\x1b[48;2;255;160;122m"
	LightSeaGreenBackground        = "\x1b[48;2;32;178;170m"
	LightSkyBlueBackground         = "\x1b[48;2;135;206;250m"
	LightSlateGrayBackground       = "\x1b[48;2;119;136;153m"
	LightSteelBlueBackground       = "\x1b[48;2;176;196;222m"
	LightYellowBackground          = "\x1b[48;2;255;255;224m"
	LimeBackground                 = "\x1b[48;2;0;255;0m"
	LimeGreenBackground            = "\x1b[48;2;50;205;50m"
	LinenBackground                = "\x1b[48;2;250;240;230m"
	MagentaBackground              = "\x1b[48;2;255;0;255m"
	MaroonBackground               = "\x1b[48;2;128;0;0m"
	MediumAquamarineBackground     = "\x1b[48;2;102;205;170m"
	MediumBlueBackground           = "\x1b[48;2;0;0;205m"
	MediumOrchidBackground         = "\x1b[48;2;186;85;211m"
	MediumPurpleBackground         = "\x1b[48;2;147;112;219m"
	MediumSeaGreenBackground       = "\x1b[48;2;60;179;113m"
	MediumSlateBlueBackground      = "\x1b[48;2;123;104;238m"
	MediumSpringGreenBackground    = "\x1b[48;2;0;250;154m"
	MediumTurquoiseBackground      = "\x1b[48;2;72;209;204m"
	MediumVioletRedBackground      = "\x1b[48;2;199;21;133m"
	MidnightBlueBackground         = "\x1b[48;2;25;25;112m"
	MintCreamBackground            = "\x1b[48;2;245;255;250m"
	MistyRoseBackground            = "\x1b[48;2;255;228;225m"
	MoccasinBackground             = "\x1b[48;2;255;228;181m"
	NavajoWhiteBackground          = "\x1b[48;2;255;222;173m"
	NavyBackground                 = "\x1b[48;2;0;0;128m"
	NavyBlueBackground             = "\x1b[48;2;159;175;223m"
	OldLaceBackground              = "\x1b[48;2;253;245;230m"
	OliveBackground                = "\x1b[48;2;128;128;0m"
	OliveDrabBackground            = "\x1b[48;2;107;142;35m"
	OrangeBackground               = "\x1b[48;2;255;165;0m"
	OrangeRedBackground            = "\x1b[48;2;255;69;0m"
	OrchidBackground               = "\x1b[48;2;218;112;214m"
	PaleGoldenRodBackground        = "\x1b[48;2;238;232;170m"
	PaleGreenBackground            = "\x1b[48;2;152;251;152m"
	PaleTurquoiseBackground        = "\x1b[48;2;175;238;238m"
	PaleVioletRedBackground        = "\x1b[48;2;219;112;147m"
	PapayaWhipBackground           = "\x1b[48;2;255;239;213m"
	PeachPuffBackground            = "\x1b[48;2;255;218;185m"
	PeruBackground                 = "\x1b[48;2;205;133;63m"
	PinkBackground                 = "\x1b[48;2;255;192;203m"
	PlumBackground                 = "\x1b[48;2;221;160;221m"
	PowderBlueBackground           = "\x1b[48;2;176;224;230m"
	PurpleBackground               = "\x1b[48;2;128;0;128m"
	RebeccaPurpleBackground        = "\x1b[38:2;102;51;153m"
	RedBackground                  = "\x1b[48;2;255;0;0m"
	RosyBrownBackground            = "\x1b[48;2;188;143;143m"
	RoyalBlueBackground            = "\x1b[48;2;65;105;225m"
	SaddleBrownBackground          = "\x1b[48;2;139;69;19m"
	SalmonBackground               = "\x1b[48;2;250;128;114m"
	SandyBrownBackground           = "\x1b[48;2;244;164;96m"
	SeaGreenBackground             = "\x1b[48;2;46;139;87m"
	SeashellBackground             = "\x1b[48;2;255;245;238m"
	SiennaBackground               = "\x1b[48;2;160;82;45m"
	SilverBackground               = "\x1b[48;2;192;192;192m"
	SkyBlueBackground              = "\x1b[48;2;135;206;235m"
	SlateBlueBackground            = "\x1b[48;2;106;90;205m"
	SlateGrayBackground            = "\x1b[48;2;112;128;144m"
	SnowBackground                 = "\x1b[48;2;255;250;250m"
	SpringGreenBackground          = "\x1b[48;2;0;255;127m"
	SteelBlueBackground            = "\x1b[48;2;70;130;180m"
	TanBackground                  = "\x1b[48;2;210;180;140m"
	TealBackground                 = "\x1b[48;2;0;128;128m"
	ThistleBackground              = "\x1b[48;2;216;191;216m"
	TomatoBackground               = "\x1b[48;2;255;99;71m"
	TurquoiseBackground            = "\x1b[48;2;64;224;208m"
	VioletBackground               = "\x1b[48;2;238;130;238m"
	WheatBackground                = "\x1b[48;2;245;222;179m"
	WhiteBackground                = "\x1b[48;2;255;255;255m"
	WhiteSmokeBackground           = "\x1b[48;2;245;245;245m"
	YellowBackground               = "\x1b[48;2;255;255;0m"
	YellowGreenBackground          = "\x1b[48;2;139;205;50m"

	// Extra ANSI escape codes
	Reset = "\x1b[0m"
	X     = "\x1b[0m"
)

// HexMap is a map of all css colors matched with their corresponding
// decimal values
var HexMap = map[string]int32{
	"aliceblue":            0xF0F8FF,
	"antiquewhite":         0xFAEBD7,
	"aqua":                 0x00FFFF,
	"aquamarine":           0x7FFFD4,
	"azure":                0xF0FFFF,
	"beige":                0xF5F5DC,
	"black":                0x000000,
	"bisque":               0xFFE4C4,
	"blanchedalmond":       0xFFEBCD,
	"blue":                 0x0000FF,
	"blueviolet":           0x8A2BE2,
	"brown":                0xA52A2A,
	"burlywood":            0xDEB887,
	"cadetblue":            0x5F9EA0,
	"chartreuse":           0x7FFF00,
	"chocolate":            0xD2691E,
	"coral":                0xFF7F50,
	"cornflowerblue":       0x6495ED,
	"cornsilk":             0xFFF8DC,
	"crimson":              0xDC143C,
	"cyan":                 0x00FFFF,
	"darkblue":             0x00008B,
	"darkcyan":             0x008B8B,
	"darkgoldenrod":        0xB8860B,
	"darkgray":             0xA9A9A9,
	"darkgreen":            0x006400,
	"darkkhaki":            0xBDB76B,
	"darkmagenta":          0x8B008B,
	"darkolivegreen":       0x556B2F,
	"darkorange":           0xFF8C00,
	"darkorchid":           0x9932CC,
	"darkred":              0x8B0000,
	"darksalmon":           0xE9967A,
	"darkseagreen":         0x8FBC8F,
	"darkslateblue":        0x483D8B,
	"darkslategray":        0x2F4F4F,
	"darkturquoise":        0x00CED1,
	"darkviolet":           0x9400D3,
	"deeppink":             0xFF1493,
	"deepskyblue":          0x00BFFF,
	"dimgray":              0x696969,
	"dodgerblue":           0x1E90FF,
	"firebrick":            0xB22222,
	"floralwhite":          0xFFFAF0,
	"forestgreen":          0x228B22,
	"fuchsia":              0xFF00FF,
	"gainsboro":            0xDCDCDC,
	"ghostwhite":           0xF8F8FF,
	"gold":                 0xFFD700,
	"goldenrod":            0xDAA520,
	"gray":                 0x808080,
	"green":                0x008000,
	"greenyellow":          0xADFF2F,
	"honeydew":             0xF0FFF0,
	"hotpink":              0xFF69B4,
	"indianred":            0xCD5C5C,
	"indigo":               0x4B0082,
	"ivory":                0xFFFFF0,
	"khaki":                0xF0E68C,
	"lavender":             0xE6E6FA,
	"lavenderblush":        0xFFF0F5,
	"lawngreen":            0x7CFC00,
	"lemonchiffon":         0xFFFACD,
	"lightblue":            0xADD8E6,
	"lightcoral":           0xF08080,
	"lightcyan":            0xE0FFFF,
	"lightgoldenrodyellow": 0xFAFAD2,
	"lightgray":            0xD3D3D3,
	"lightgreen":           0x90EE90,
	"lightpink":            0xFFB6C1,
	"lightsalmon":          0xFFA07A,
	"lightseagreen":        0x20B2AA,
	"lightskyblue":         0x87CEFA,
	"lightslategray":       0x778899,
	"lightsteelblue":       0xB0C4DE,
	"lightyellow":          0xFFFFE0,
	"lime":                 0x00FF00,
	"limegreen":            0x32CD32,
	"linen":                0xFAF0E6,
	"magenta":              0xFF00FF,
	"maroon":               0x800000,
	"mediumaquamarine":     0x66CDAA,
	"mediumblue":           0x0000CD,
	"mediumorchid":         0xBA55D3,
	"mediumpurple":         0x9370DB,
	"mediumseagreen":       0x3CB371,
	"mediumslateblue":      0x7B68EE,
	"mediumspringgreen":    0x00FA9A,
	"mediumturquoise":      0x48D1CC,
	"mediumvioletred":      0xC71585,
	"midnightblue":         0x191970,
	"mintcream":            0xF5FFFA,
	"mistyrose":            0xFFE4E1,
	"moccasin":             0xFFE4B5,
	"navajowhite":          0xFFDEAD,
	"navy":                 0x000080,
	"navyblue":             0x9FAFDF,
	"oldlace":              0xFDF5E6,
	"olive":                0x808000,
	"olivedrab":            0x6B8E23,
	"orange":               0xFFA500,
	"orangered":            0xFF4500,
	"orchid":               0xDA70D6,
	"palegoldenrod":        0xEEE8AA,
	"palegreen":            0x98FB98,
	"paleturquoise":        0xAFEEEE,
	"palevioletred":        0xDB7093,
	"papayawhip":           0xFFEFD5,
	"peachpuff":            0xFFDAB9,
	"peru":                 0xCD853F,
	"pink":                 0xFFC0CB,
	"plum":                 0xDDA0DD,
	"powderblue":           0xB0E0E6,
	"purple":               0x800080,
	"red":                  0xFF0000,
	"rebeccapurple":        0x663399,
	"rosybrown":            0xBC8F8F,
	"royalblue":            0x4169E1,
	"saddlebrown":          0x8B4513,
	"salmon":               0xFA8072,
	"sandybrown":           0xF4A460,
	"seagreen":             0x2E8B57,
	"seashell":             0xFFF5EE,
	"sienna":               0xA0522D,
	"silver":               0xC0C0C0,
	"skyblue":              0x87CEEB,
	"slateblue":            0x6A5ACD,
	"slategray":            0x708090,
	"snow":                 0xFFFAFA,
	"springgreen":          0x00FF7F,
	"steelblue":            0x4682B4,
	"tan":                  0xD2B48C,
	"teal":                 0x008080,
	"thistle":              0xD8BFD8,
	"tomato":               0xFF6347,
	"turquoise":            0x40E0D0,
	"violet":               0xEE82EE,
	"wheat":                0xF5DEB3,
	"white":                0xFFFFFF,
	"whitesmoke":           0xF5F5F5,
	"yellow":               0xFFFF00,
	"yellowgreen":          0x9ACD32,
}

// AnsiMap is a map that contains all of the css color names and their
// corresponding ansi escape codes for text color
var AnsiMap = map[string]string{
	"aliceblue":            "\x1b[38;2;240;248;255m",
	"antiquewhite":         "\x1b[38;2;250;235;215m",
	"aqua":                 "\x1b[38;2;0;255;255m",
	"aquamarine":           "\x1b[38;2;127;255;212m",
	"azure":                "\x1b[38;2;1240;255;255m",
	"beige":                "\x1b[38;2;245;245;220m",
	"bisque":               "\x1b[38;2;255;228;196m",
	"black":                "\x1b[38;2;0;0;0m",
	"blanchedalmond":       "\x1b[38;2;255;235;205m",
	"blue":                 "\x1b[38;2;0;0;255m",
	"blueviolet":           "\x1b[38;2;138;43;226m",
	"brown":                "\x1b[38;2;165;42;42m",
	"burlywood":            "\x1b[38;2;222;184;135m",
	"cadetblue":            "\x1b[38;2;95;158;160m",
	"chartreuse":           "\x1b[38;2;95;158;160m",
	"chocolate":            "\x1b[38;2;210;105;30m",
	"coral":                "\x1b[38;2;255;127;80m",
	"cornflowerblue":       "\x1b[38;2;100;149;237m",
	"cornsilk":             "\x1b[38;2;255;248;220m",
	"crimson":              "\x1b[38;2;220;20;60m",
	"cyan":                 "\x1b[38;2;0;255;255m",
	"darkblue":             "\x1b[38;2;0;0;139m",
	"darkcyan":             "\x1b[38;2;0;139;139m",
	"darkgoldenrod":        "\x1b[38;2;184;134;11m",
	"darkgray":             "\x1b[38;2;169;169;169m",
	"darkgreen":            "\x1b[38;2;0;100;0m",
	"darkkhaki":            "\x1b[38;2;189;183;107m",
	"darkmagenta":          "\x1b[38;2;139;0;139m",
	"darkolivegreen":       "\x1b[38;2;85;107;47m",
	"darkorange":           "\x1b[38;2;255;140;0m",
	"darkorchid":           "\x1b[38;2;153;50;204m",
	"darkred":              "\x1b[38;2;139;0;0m",
	"darksalmon":           "\x1b[38;2;233;150;122m",
	"darkseagreen":         "\x1b[38;2;143;188;143m",
	"darkslateblue":        "\x1b[38;2;72;61;139m",
	"darkslategray":        "\x1b[38;2;47;79;79m",
	"darkturquoise":        "\x1b[38;2;0;206;209m",
	"darkviolet":           "\x1b[38;2;148;0;211m",
	"deeppink":             "\x1b[38;2;255;20;147m",
	"deepskyblue":          "\x1b[38;2;0;191;255m",
	"dimgray":              "\x1b[38;2;0;191;255m",
	"dodgerblue":           "\x1b[38;2;30;144;255m",
	"firebrick":            "\x1b[38;2;178;34;34m",
	"floralwhite":          "\x1b[38;2;255;250;240m",
	"forestgreen":          "\x1b[38;2;34;139;34m",
	"fuchsia":              "\x1b[38;2;255;0;255m",
	"gainsboro":            "\x1b[38;2;220;220;220m",
	"ghostwhite":           "\x1b[38;2;248;248;255m",
	"gold":                 "\x1b[38;2;255;215;0m",
	"goldenrod":            "\x1b[38;2;218;165;32m",
	"gray":                 "\x1b[38;2;127;127;127m",
	"green":                "\x1b[38;2;0;128;0m",
	"greenyellow":          "\x1b[38;2;173;255;47m",
	"honeydew":             "\x1b[38;2;240;255;240m",
	"hotpink":              "\x1b[38;2;255;105;180m",
	"indianred":            "\x1b[38;2;205;92;92m",
	"indigo":               "\x1b[38;2;75;0;130m",
	"ivory":                "\x1b[38;2;255;255;240m",
	"khaki":                "\x1b[38;2;240;230;140m",
	"lavender":             "\x1b[38;2;230;230;250m",
	"lavenderblush":        "\x1b[38;2;255;240;245m",
	"lawngreen":            "\x1b[38;2;124;252;0m",
	"lemonchiffon":         "\x1b[38;2;255;250;205m",
	"lightblue":            "\x1b[38;2;173;216;230m",
	"lightcoral":           "\x1b[38;2;240;128;128m",
	"lightcyan":            "\x1b[38;2;224;255;255m",
	"lightgoldenrodyellow": "\x1b[38;2;250;250;210m",
	"lightgreen":           "\x1b[38;2;144;238;144m",
	"lightgray":            "\x1b[38;2;211;211;211m",
	"lightpink":            "\x1b[38;2;255;182;193m",
	"lightsalmon":          "\x1b[38;2;255;160;122m",
	"lightseagreen":        "\x1b[38;2;32;178;170m",
	"lightskyblue":         "\x1b[38;2;135;206;250m",
	"lightslategray":       "\x1b[38;2;119;136;153m",
	"lightsteelblue":       "\x1b[38;2;176;196;222m",
	"lightyellow":          "\x1b[38;2;255;255;224m",
	"lime":                 "\x1b[38;2;0;255;0m",
	"limegreen":            "\x1b[38;2;50;205;50m",
	"linen":                "\x1b[38;2;250;240;230m",
	"magenta":              "\x1b[38;2;255;0;255m",
	"maroon":               "\x1b[38;2;128;0;0m",
	"mediumaquamarine":     "\x1b[38;2;102;205;170m",
	"mediumblue":           "\x1b[38;2;0;0;205m",
	"mediumorchid":         "\x1b[38;2;186;85;211m",
	"mediumpurple":         "\x1b[38;2;147;112;219m",
	"mediumseagreen":       "\x1b[38;2;60;179;113m",
	"mediumslateblue":      "\x1b[38;2;123;104;238m",
	"mediumspringgreen":    "\x1b[38;2;0;250;154m",
	"mediumturquoise":      "\x1b[38;2;72;209;204m",
	"mediumvioletred":      "\x1b[38;2;199;21;133m",
	"midnightblue":         "\x1b[38;2;25;25;112m",
	"mintcream":            "\x1b[38;2;245;255;250m",
	"mistyrose":            "\x1b[38;2;255;228;225m",
	"moccasin":             "\x1b[38;2;255;228;181m",
	"navajowhite":          "\x1b[38;2;255;222;173m",
	"navy":                 "\x1b[38;2;0;0;128m",
	"navyblue":             "\x1b[38;2;159;175;223m",
	"oldlace":              "\x1b[38;2;253;245;230m",
	"olive":                "\x1b[38;2;128;128;0m",
	"olivedrab":            "\x1b[38;2;107;142;35m",
	"orange":               "\x1b[38;2;255;165;0m",
	"orangered":            "\x1b[38;2;255;69;0m",
	"orchid":               "\x1b[38;2;218;112;214m",
	"palegoldenrod":        "\x1b[38;2;238;232;170m",
	"palegreen":            "\x1b[38;2;152;251;152m",
	"paleturquoise":        "\x1b[38;2;175;238;238m",
	"palevioletred":        "\x1b[38;2;219;112;147m",
	"papayawhip":           "\x1b[38;2;255;239;213m",
	"peachpuff":            "\x1b[38;2;255;218;185m",
	"peru":                 "\x1b[38;2;205;133;63m",
	"pink":                 "\x1b[38;2;255;192;203m",
	"plum":                 "\x1b[38;2;221;160;221m",
	"powderblue":           "\x1b[38;2;176;224;230m",
	"purple":               "\x1b[38;2;128;0;128m",
	"red":                  "\x1b[38;2;255;0;0m",
	"rosybrown":            "\x1b[38;2;188;143;143m",
	"royalblue":            "\x1b[38;2;65;105;225m",
	"saddlebrown":          "\x1b[38;2;139;69;19m",
	"salmon":               "\x1b[38;2;250;128;114m",
	"sandybrown":           "\x1b[38;2;244;164;96m",
	"seagreen":             "\x1b[38;2;46;139;87m",
	"seashell":             "\x1b[38;2;255;245;238m",
	"sienna":               "\x1b[38;2;160;82;45m",
	"silver":               "\x1b[38;2;192;192;192m",
	"skyblue":              "\x1b[38;2;135;206;235m",
	"slateblue":            "\x1b[38;2;106;90;205m",
	"slategray":            "\x1b[38;2;112;128;144m",
	"snow":                 "\x1b[38;2;255;250;250m",
	"springgreen":          "\x1b[38;2;0;255;127m",
	"steelblue":            "\x1b[38;2;70;130;180m",
	"tan":                  "\x1b[38;2;210;180;140m",
	"teal":                 "\x1b[38;2;0;128;128m",
	"thistle":              "\x1b[38;2;216;191;216m",
	"tomato":               "\x1b[38;2;255;99;71m",
	"turquoise":            "\x1b[38;2;64;224;208m",
	"violet":               "\x1b[38;2;238;130;238m",
	"wheat":                "\x1b[38;2;245;222;179m",
	"white":                "\x1b[38;2;255;255;255m",
	"whitesmoke":           "\x1b[38;2;245;245;245m",
	"yellow":               "\x1b[38;2;255;255;0m",
	"yellowgreen":          "\x1b[38;2;139;205;50m",
}

// AnsiBackgroundMap is a map of all css color names matched with their
// corresponding ansi escape codes for backgrounds
var AnsiBackgroundMap = map[string]string{
	"alicebluebackground":            "\x1b[48;2;240;248;255m",
	"antiquewhitebackground":         "\x1b[48;2;250;235;215m",
	"aquabackground":                 "\x1b[48;2;0;255;255m",
	"aquamarinebackground":           "\x1b[48;2;127;255;212m",
	"azurebackground":                "\x1b[48;2;1240;255;255m",
	"beigebackground":                "\x1b[48;2;245;245;220m",
	"bisquebackground":               "\x1b[48;2;255;228;196m",
	"blackbackground":                "\x1b[48;2;0;0;0m",
	"blanchedalmondbackground":       "\x1b[48;2;255;235;205m",
	"bluebackground":                 "\x1b[48;2;0;0;255m",
	"bluevioletbackground":           "\x1b[48;2;138;43;226m",
	"brownbackground":                "\x1b[48;2;165;42;42m",
	"burlywoodbackground":            "\x1b[48;2;222;184;135m",
	"cadetbluebackground":            "\x1b[48;2;95;158;160m",
	"chartreusebackground":           "\x1b[48;2;95;158;160m",
	"chocolatebackground":            "\x1b[48;2;210;105;30m",
	"coralbackground":                "\x1b[48;2;255;127;80m",
	"cornflowerbluebackground":       "\x1b[48;2;100;149;237m",
	"cornsilkbackground":             "\x1b[48;2;255;248;220m",
	"crimsonbackground":              "\x1b[48;2;220;20;60m",
	"cyanbackground":                 "\x1b[48;2;0;255;255m",
	"darkbluebackground":             "\x1b[48;2;0;0;139m",
	"darkcyanbackground":             "\x1b[48;2;0;139;139m",
	"darkgoldenrodbackground":        "\x1b[48;2;184;134;11m",
	"darkgraybackground":             "\x1b[48;2;169;169;169m",
	"darkgreenbackground":            "\x1b[48;2;0;100;0m",
	"darkkhakibackground":            "\x1b[48;2;189;183;107m",
	"darkmagentabackground":          "\x1b[48;2;139;0;139m",
	"darkolivegreenbackground":       "\x1b[48;2;85;107;47m",
	"darkorangebackground":           "\x1b[48;2;255;140;0m",
	"darkorchidbackground":           "\x1b[48;2;153;50;204m",
	"darkredbackground":              "\x1b[48;2;139;0;0m",
	"darksalmonbackground":           "\x1b[48;2;233;150;122m",
	"darkseagreenbackground":         "\x1b[48;2;143;188;143m",
	"darkslatebluebackground":        "\x1b[48;2;72;61;139m",
	"darkslategraybackground":        "\x1b[48;2;47;79;79m",
	"darkturquoisebackground":        "\x1b[48;2;0;206;209m",
	"darkvioletbackground":           "\x1b[48;2;148;0;211m",
	"deeppinkbackground":             "\x1b[48;2;255;20;147m",
	"deepskybluebackground":          "\x1b[48;2;0;191;255m",
	"dimgraybackground":              "\x1b[48;2;0;191;255m",
	"dodgerbluebackground":           "\x1b[48;2;30;144;255m",
	"firebrickbackground":            "\x1b[48;2;178;34;34m",
	"floralwhitebackground":          "\x1b[48;2;255;250;240m",
	"forestgreenbackground":          "\x1b[48;2;34;139;34m",
	"fuchsiabackground":              "\x1b[48;2;255;0;255m",
	"gainsborobackground":            "\x1b[48;2;220;220;220m",
	"ghostwhitebackground":           "\x1b[48;2;248;248;255m",
	"goldbackground":                 "\x1b[48;2;255;215;0m",
	"goldenrodbackground":            "\x1b[48;2;218;165;32m",
	"graybackground":                 "\x1b[48;2;127;127;127m",
	"greenbackground":                "\x1b[48;2;0;128;0m",
	"greenyellowbackground":          "\x1b[48;2;173;255;47m",
	"honeydewbackground":             "\x1b[48;2;240;255;240m",
	"hotpinkbackground":              "\x1b[48;2;255;105;180m",
	"indianredbackground":            "\x1b[48;2;205;92;92m",
	"indigobackground":               "\x1b[48;2;75;0;130m",
	"ivorybackground":                "\x1b[48;2;255;255;240m",
	"khakibackground":                "\x1b[48;2;240;230;140m",
	"lavenderbackground":             "\x1b[48;2;230;230;250m",
	"lavenderblushbackground":        "\x1b[48;2;255;240;245m",
	"lawngreenbackground":            "\x1b[48;2;124;252;0m",
	"lemonchiffonbackground":         "\x1b[48;2;255;250;205m",
	"lightbluebackground":            "\x1b[48;2;173;216;230m",
	"lightcoralbackground":           "\x1b[48;2;240;128;128m",
	"lightcyanbackground":            "\x1b[48;2;224;255;255m",
	"lightgoldenrodyellowbackground": "\x1b[48;2;250;250;210m",
	"lightgreenbackground":           "\x1b[48;2;144;238;144m",
	"lightgraybackground":            "\x1b[48;2;211;211;211m",
	"lightpinkbackground":            "\x1b[48;2;255;182;193m",
	"lightsalmonbackground":          "\x1b[48;2;255;160;122m",
	"lightseagreenbackground":        "\x1b[48;2;32;178;170m",
	"lightskybluebackground":         "\x1b[48;2;135;206;250m",
	"lightslategraybackground":       "\x1b[48;2;119;136;153m",
	"lightsteelbluebackground":       "\x1b[48;2;176;196;222m",
	"lightyellowbackground":          "\x1b[48;2;255;255;224m",
	"limebackground":                 "\x1b[48;2;0;255;0m",
	"limegreenbackground":            "\x1b[48;2;50;205;50m",
	"linenbackground":                "\x1b[48;2;250;240;230m",
	"magentabackground":              "\x1b[48;2;255;0;255m",
	"maroonbackground":               "\x1b[48;2;128;0;0m",
	"mediumaquamarinebackground":     "\x1b[48;2;102;205;170m",
	"mediumbluebackground":           "\x1b[48;2;0;0;205m",
	"mediumorchidbackground":         "\x1b[48;2;186;85;211m",
	"mediumpurplebackground":         "\x1b[48;2;147;112;219m",
	"mediumseagreenbackground":       "\x1b[48;2;60;179;113m",
	"mediumslatebluebackground":      "\x1b[48;2;123;104;238m",
	"mediumspringgreenbackground":    "\x1b[48;2;0;250;154m",
	"mediumturquoisebackground":      "\x1b[48;2;72;209;204m",
	"mediumvioletredbackground":      "\x1b[48;2;199;21;133m",
	"midnightbluebackground":         "\x1b[48;2;25;25;112m",
	"mintcreambackground":            "\x1b[48;2;245;255;250m",
	"mistyrosebackground":            "\x1b[48;2;255;228;225m",
	"moccasinbackground":             "\x1b[48;2;255;228;181m",
	"navajowhitebackground":          "\x1b[48;2;255;222;173m",
	"navybackground":                 "\x1b[48;2;0;0;128m",
	"navybluebackground":             "\x1b[48;2;159;175;223m",
	"oldlacebackground":              "\x1b[48;2;253;245;230m",
	"olivebackground":                "\x1b[48;2;128;128;0m",
	"olivedrabbackground":            "\x1b[48;2;107;142;35m",
	"orangebackground":               "\x1b[48;2;255;165;0m",
	"orangeredbackground":            "\x1b[48;2;255;69;0m",
	"orchidbackground":               "\x1b[48;2;218;112;214m",
	"palegoldenrodbackground":        "\x1b[48;2;238;232;170m",
	"palegreenbackground":            "\x1b[48;2;152;251;152m",
	"paleturquoisebackground":        "\x1b[48;2;175;238;238m",
	"palevioletredbackground":        "\x1b[48;2;219;112;147m",
	"papayawhipbackground":           "\x1b[48;2;255;239;213m",
	"peachpuffbackground":            "\x1b[48;2;255;218;185m",
	"perubackground":                 "\x1b[48;2;205;133;63m",
	"pinkbackground":                 "\x1b[48;2;255;192;203m",
	"plumbackground":                 "\x1b[48;2;221;160;221m",
	"powderbluebackground":           "\x1b[48;2;176;224;230m",
	"purplebackground":               "\x1b[48;2;128;0;128m",
	"redbackground":                  "\x1b[48;2;255;0;0m",
	"rosybrownbackground":            "\x1b[48;2;188;143;143m",
	"royalbluebackground":            "\x1b[48;2;65;105;225m",
	"saddlebrownbackground":          "\x1b[48;2;139;69;19m",
	"salmonbackground":               "\x1b[48;2;250;128;114m",
	"sandybrownbackground":           "\x1b[48;2;244;164;96m",
	"seagreenbackground":             "\x1b[48;2;46;139;87m",
	"seashellbackground":             "\x1b[48;2;255;245;238m",
	"siennabackground":               "\x1b[48;2;160;82;45m",
	"silverbackground":               "\x1b[48;2;192;192;192m",
	"skybluebackground":              "\x1b[48;2;135;206;235m",
	"slatebluebackground":            "\x1b[48;2;106;90;205m",
	"slategraybackground":            "\x1b[48;2;112;128;144m",
	"snowbackground":                 "\x1b[48;2;255;250;250m",
	"springgreenbackground":          "\x1b[48;2;0;255;127m",
	"steelbluebackground":            "\x1b[48;2;70;130;180m",
	"tanbackground":                  "\x1b[48;2;210;180;140m",
	"tealbackground":                 "\x1b[48;2;0;128;128m",
	"thistlebackground":              "\x1b[48;2;216;191;216m",
	"tomatobackground":               "\x1b[48;2;255;99;71m",
	"turquoisebackground":            "\x1b[48;2;64;224;208m",
	"violetbackground":               "\x1b[48;2;238;130;238m",
	"wheatbackground":                "\x1b[48;2;245;222;179m",
	"whitebackground":                "\x1b[48;2;255;255;255m",
	"whitesmokebackground":           "\x1b[48;2;245;245;245m",
	"yellowbackground":               "\x1b[48;2;255;255;0m",
	"yellowgreenbackground":          "\x1b[48;2;139;205;50m",
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
