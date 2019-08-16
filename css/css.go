package css

import (
        "math/rand"
        "time"
)

const (
	AliceBlue            = "\x1b[38;2;240;248;255m"
	AntiqueWhite         = "\x1b[38;2;250;235;215m"
	Aqua                 = "\x1b[38;2;0;255;255m"
	AquaMarine           = "\x1b[38;2;127;255;212m"
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
	CornFlowerBlue       = "\x1b[38;2;100;149;237m"
	CornSilk             = "\x1b[38;2;255;248;220m"
	Crimson              = "\x1b[38;2;220;20;60m"
	Cyan                 = "\x1b[38;2;0;255;255m"
	DarkBlue             = "\x1b[38;2;0;0;139m"
	DarkCyan             = "\x1b[38;2;0;139;139m"
	DarkGoldenRod        = "\x1b[38;2;184;134;11m"
	DarkGray             = "\x1b[38;2;169;169;169m"
	DarkGreen            = "\x1b[38;2;0;100;0m"
	DarkKhaki            = "\x1b[38;2;189;183;107m"
	DarkMagenta          = "\x1b[38;2;139;0;139m"
	DarkOliveGreen       = "\x1b[38;2;85;107;47m"
	DarkOrange           = "\x1b[38;2;255;140;0m"
	DarkOrchid           = "\x1b[38;2;153;50;204m"
	DarkRed              = "\x1b[38;2;139;0;0m"
	DarkSalmon           = "\x1b[38;2;233;150;122m"
	DarkSeaGreen         = "\x1b[38;2;143;188;143m"
	DarkSlateBlue        = "\x1b[38;2;72;61;139m"
	DarkSlateGray        = "\x1b[38;2;47;79;79m"
	DarkTurquoise        = "\x1b[38;2;0;206;209m"
	DarkViolet           = "\x1b[38;2;148;0;211m"
	DeepPink             = "\x1b[38;2;255;20;147m"
	DeepSkyBlue          = "\x1b[38;2;0;191;255m"
	DimGray              = "\x1b[38;2;0;191;255m"
	DodgerBlue           = "\x1b[38;2;30;144;255m"
	Firebrick            = "\x1b[38;2;178;34;34m"
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
	LightGrey            = "\x1b[38;2;211;211;211m"
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
	MediumAquaMarine     = "\x1b[38;2;102;205;170m"
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
	Red                  = "\x1b[38;2;255;0;0m"
	RosyBrown            = "\x1b[38;2;188;143;143m"
	RoyalBlue            = "\x1b[38;2;65;105;225m"
	SaddleBrown          = "\x1b[38;2;139;69;19m"
	Salmon               = "\x1b[38;2;250;128;114m"
	SandyBrown           = "\x1b[38;2;244;164;96m"
	SeaGreen             = "\x1b[38;2;46;139;87m"
	SeaShell             = "\x1b[38;2;255;245;238m"
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

	Reset = "\033[0m"
	X     = "\033[0m"

	ClearScreen = "\033[2J\033[H"
	ClearLine   = "\033[2K\033[G"
	CursorOff   = "\033[?25l"
	CursorOn    = "\033[?25h"
	StrikeOut   = "\033[9m"
)

var Index = map[string]string{
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
	"lightgrey":            "\x1b[38;2;211;211;211m",
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

var Colors = [...]string{AliceBlue, AntiqueWhite, Aqua, AquaMarine, Azure, Beige, Bisque, Black, BlanchedAlmond, Blue, BlueViolet, Brown, BurlyWood, CadetBlue, Chartreuse, Chocolate, Coral, CornFlowerBlue, CornSilk, Crimson, Cyan, DarkBlue, DarkCyan, DarkGoldenRod, DarkGray, DarkGreen, DarkKhaki, DarkMagenta, DarkOliveGreen, DarkOrange, DarkOrchid, DarkRed, DarkSalmon, DarkSeaGreen, DarkSlateBlue, DarkSlateGray, DarkTurquoise, DarkViolet, DeepPink, DeepSkyBlue, DimGray, DodgerBlue, Firebrick, FloralWhite, ForestGreen, Fuchsia, Gainsboro, GhostWhite, Gold, GoldenRod, Gray, Green, GreenYellow, Honeydew, HotPink, IndianRed, Indigo, Ivory, Khaki, Lavender, LavenderBlush, LawnGreen, LemonChiffon, LightBlue, LightCoral, LightCyan, LightGoldenRodYellow, LightGreen, LightGrey, LightPink, LightSalmon, LightSeaGreen, LightSkyBlue, LightSlateGray, LightSteelBlue, LightYellow, Lime, LimeGreen, Linen, Magenta, Maroon, MediumAquaMarine, MediumBlue, MediumOrchid, MediumPurple, MediumSeaGreen, MediumSlateBlue, MediumSpringGreen, MediumTurquoise, MediumVioletRed, MidnightBlue, MintCream, MistyRose, Moccasin, NavajoWhite, Navy, NavyBlue, OldLace, Olive, OliveDrab, Orange, OrangeRed, Orchid, PaleGoldenRod, PaleGreen, PaleTurquoise, PaleVioletRed, PapayaWhip, PeachPuff, Peru, Pink, Plum, PowderBlue, Purple, Red, RosyBrown, RoyalBlue, SaddleBrown, Salmon, SandyBrown, SeaGreen, SeaShell, Sienna, Silver, SkyBlue, SlateBlue, SlateGray, Snow, SpringGreen, SteelBlue, Tan, Teal, Thistle, Tomato, Turquoise, Violet, Wheat, White, WhiteSmoke, Yellow, YellowGreen}

func Random() string {
        rand.Seed(time.Now().UnixNano())
        return Colors[rand.Intn(len(Colors))]
}
