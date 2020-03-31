package css

import (
	"math/rand"
	"time"
)

const (
	Reset = "\033[0m"

	AliceBlue            = "\033[38;2;240;248;255m"
	AntiqueWhite         = "\033[38;2;250;235;215m"
	Aqua                 = "\033[38;2;0;255;255m"
	AquaMarine           = "\033[38;2;127;255;212m"
	Azure                = "\033[38;2;1240;255;255m"
	Beige                = "\033[38;2;245;245;220m"
	Bisque               = "\033[38;2;255;228;196m"
	Black                = "\033[38;2;0;0;0m"
	BlanchedAlmond       = "\033[38;2;255;235;205m"
	Blue                 = "\033[38;2;0;0;255m"
	BlueViolet           = "\033[38;2;138;43;226m"
	Brown                = "\033[38;2;165;42;42m"
	BurlyWood            = "\033[38;2;222;184;135m"
	CadetBlue            = "\033[38;2;95;158;160m"
	Chartreuse           = "\033[38;2;95;158;160m"
	Chocolate            = "\033[38;2;210;105;30m"
	Coral                = "\033[38;2;255;127;80m"
	CornFlowerBlue       = "\033[38;2;100;149;237m"
	CornSilk             = "\033[38;2;255;248;220m"
	Crimson              = "\033[38;2;220;20;60m"
	Cyan                 = "\033[38;2;0;255;255m"
	DarkBlue             = "\033[38;2;0;0;139m"
	DarkCyan             = "\033[38;2;0;139;139m"
	DarkGoldenRod        = "\033[38;2;184;134;11m"
	DarkGray             = "\033[38;2;169;169;169m"
	DarkGreen            = "\033[38;2;0;100;0m"
	DarkKhaki            = "\033[38;2;189;183;107m"
	DarkMagenta          = "\033[38;2;139;0;139m"
	DarkOliveGreen       = "\033[38;2;85;107;47m"
	DarkOrange           = "\033[38;2;255;140;0m"
	DarkOrchid           = "\033[38;2;153;50;204m"
	DarkRed              = "\033[38;2;139;0;0m"
	DarkSalmon           = "\033[38;2;233;150;122m"
	DarkSeaGreen         = "\033[38;2;143;188;143m"
	DarkSlateBlue        = "\033[38;2;72;61;139m"
	DarkSlateGray        = "\033[38;2;47;79;79m"
	DarkTurquoise        = "\033[38;2;0;206;209m"
	DarkViolet           = "\033[38;2;148;0;211m"
	DeepPink             = "\033[38;2;255;20;147m"
	DeepSkyBlue          = "\033[38;2;0;191;255m"
	DimGray              = "\033[38;2;0;191;255m"
	DodgerBlue           = "\033[38;2;30;144;255m"
	Firebrick            = "\033[38;2;178;34;34m"
	FloralWhite          = "\033[38;2;255;250;240m"
	ForestGreen          = "\033[38;2;34;139;34m"
	Fuchsia              = "\033[38;2;255;0;255m"
	Gainsboro            = "\033[38;2;220;220;220m"
	GhostWhite           = "\033[38;2;248;248;255m"
	Gold                 = "\033[38;2;255;215;0m"
	GoldenRod            = "\033[38;2;218;165;32m"
	Gray                 = "\033[38;2;127;127;127m"
	Green                = "\033[38;2;0;128;0m"
	GreenYellow          = "\033[38;2;173;255;47m"
	Honeydew             = "\033[38;2;240;255;240m"
	HotPink              = "\033[38;2;255;105;180m"
	IndianRed            = "\033[38;2;205;92;92m"
	Indigo               = "\033[38;2;75;0;130m"
	Ivory                = "\033[38;2;255;255;240m"
	Khaki                = "\033[38;2;240;230;140m"
	Lavender             = "\033[38;2;230;230;250m"
	LavenderBlush        = "\033[38;2;255;240;245m"
	LawnGreen            = "\033[38;2;124;252;0m"
	LemonChiffon         = "\033[38;2;255;250;205m"
	LightBlue            = "\033[38;2;173;216;230m"
	LightCoral           = "\033[38;2;240;128;128m"
	LightCyan            = "\033[38;2;224;255;255m"
	LightGoldenRodYellow = "\033[38;2;250;250;210m"
	LightGreen           = "\033[38;2;144;238;144m"
	LightGrey            = "\033[38;2;211;211;211m"
	LightPink            = "\033[38;2;255;182;193m"
	LightSalmon          = "\033[38;2;255;160;122m"
	LightSeaGreen        = "\033[38;2;32;178;170m"
	LightSkyBlue         = "\033[38;2;135;206;250m"
	LightSlateGray       = "\033[38;2;119;136;153m"
	LightSteelBlue       = "\033[38;2;176;196;222m"
	LightYellow          = "\033[38;2;255;255;224m"
	Lime                 = "\033[38;2;0;255;0m"
	LimeGreen            = "\033[38;2;50;205;50m"
	Linen                = "\033[38;2;250;240;230m"
	Magenta              = "\033[38;2;255;0;255m"
	Maroon               = "\033[38;2;128;0;0m"
	MediumAquaMarine     = "\033[38;2;102;205;170m"
	MediumBlue           = "\033[38;2;0;0;205m"
	MediumOrchid         = "\033[38;2;186;85;211m"
	MediumPurple         = "\033[38;2;147;112;219m"
	MediumSeaGreen       = "\033[38;2;60;179;113m"
	MediumSlateBlue      = "\033[38;2;123;104;238m"
	MediumSpringGreen    = "\033[38;2;0;250;154m"
	MediumTurquoise      = "\033[38;2;72;209;204m"
	MediumVioletRed      = "\033[38;2;199;21;133m"
	MidnightBlue         = "\033[38;2;25;25;112m"
	MintCream            = "\033[38;2;245;255;250m"
	MistyRose            = "\033[38;2;255;228;225m"
	Moccasin             = "\033[38;2;255;228;181m"
	NavajoWhite          = "\033[38;2;255;222;173m"
	Navy                 = "\033[38;2;0;0;128m"
	NavyBlue             = "\033[38;2;159;175;223m"
	OldLace              = "\033[38;2;253;245;230m"
	Olive                = "\033[38;2;128;128;0m"
	OliveDrab            = "\033[38;2;107;142;35m"
	Orange               = "\033[38;2;255;165;0m"
	OrangeRed            = "\033[38;2;255;69;0m"
	Orchid               = "\033[38;2;218;112;214m"
	PaleGoldenRod        = "\033[38;2;238;232;170m"
	PaleGreen            = "\033[38;2;152;251;152m"
	PaleTurquoise        = "\033[38;2;175;238;238m"
	PaleVioletRed        = "\033[38;2;219;112;147m"
	PapayaWhip           = "\033[38;2;255;239;213m"
	PeachPuff            = "\033[38;2;255;218;185m"
	Peru                 = "\033[38;2;205;133;63m"
	Pink                 = "\033[38;2;255;192;203m"
	Plum                 = "\033[38;2;221;160;221m"
	PowderBlue           = "\033[38;2;176;224;230m"
	Purple               = "\033[38;2;128;0;128m"
	RebeccaPurple        = "\033[38;2;102;51;153m"
	Red                  = "\033[38;2;255;0;0m"
	RosyBrown            = "\033[38;2;188;143;143m"
	RoyalBlue            = "\033[38;2;65;105;225m"
	SaddleBrown          = "\033[38;2;139;69;19m"
	Salmon               = "\033[38;2;250;128;114m"
	SandyBrown           = "\033[38;2;244;164;96m"
	SeaGreen             = "\033[38;2;46;139;87m"
	SeaShell             = "\033[38;2;255;245;238m"
	Sienna               = "\033[38;2;160;82;45m"
	Silver               = "\033[38;2;192;192;192m"
	SkyBlue              = "\033[38;2;135;206;235m"
	SlateBlue            = "\033[38;2;106;90;205m"
	SlateGray            = "\033[38;2;112;128;144m"
	Snow                 = "\033[38;2;255;250;250m"
	SpringGreen          = "\033[38;2;0;255;127m"
	SteelBlue            = "\033[38;2;70;130;180m"
	Tan                  = "\033[38;2;210;180;140m"
	Teal                 = "\033[38;2;0;128;128m"
	Thistle              = "\033[38;2;216;191;216m"
	Tomato               = "\033[38;2;255;99;71m"
	Turquoise            = "\033[38;2;64;224;208m"
	Violet               = "\033[38;2;238;130;238m"
	Wheat                = "\033[38;2;245;222;179m"
	White                = "\033[38;2;255;255;255m"
	WhiteSmoke           = "\033[38;2;245;245;245m"
	Yellow               = "\033[38;2;255;255;0m"
	YellowGreen          = "\033[38;2;139;205;50m"

	AliceBlueBackground            = "\033[48;2;240;248;255m"
	AntiqueWhiteBackground         = "\033[48;2;250;235;215m"
	AquaBackground                 = "\033[48;2;0;255;255m"
	AquaMarineBackground           = "\033[48;2;127;255;212m"
	AzureBackground                = "\033[48;2;1240;255;255m"
	BeigeBackground                = "\033[48;2;245;245;220m"
	BisqueBackground               = "\033[48;2;255;228;196m"
	BlackBackground                = "\033[48;2;0;0;0m"
	BlanchedAlmondBackground       = "\033[48;2;255;235;205m"
	BlueBackground                 = "\033[48;2;0;0;255m"
	BlueVioletBackground           = "\033[48;2;138;43;226m"
	BrownBackground                = "\033[48;2;165;42;42m"
	BurlyWoodBackground            = "\033[48;2;222;184;135m"
	CadetBlueBackground            = "\033[48;2;95;158;160m"
	ChartreuseBackground           = "\033[48;2;95;158;160m"
	ChocolateBackground            = "\033[48;2;210;105;30m"
	CoralBackground                = "\033[48;2;255;127;80m"
	CornFlowerBlueBackground       = "\033[48;2;100;149;237m"
	CornSilkBackground             = "\033[48;2;255;248;220m"
	CrimsonBackground              = "\033[48;2;220;20;60m"
	CyanBackground                 = "\033[48;2;0;255;255m"
	DarkBlueBackground             = "\033[48;2;0;0;139m"
	DarkCyanBackground             = "\033[48;2;0;139;139m"
	DarkGoldenRodBackground        = "\033[48;2;184;134;11m"
	DarkGrayBackground             = "\033[48;2;169;169;169m"
	DarkGreenBackground            = "\033[48;2;0;100;0m"
	DarkKhakiBackground            = "\033[48;2;189;183;107m"
	DarkMagentaBackground          = "\033[48;2;139;0;139m"
	DarkOliveGreenBackground       = "\033[48;2;85;107;47m"
	DarkOrangeBackground           = "\033[48;2;255;140;0m"
	DarkOrchidBackground           = "\033[48;2;153;50;204m"
	DarkRedBackground              = "\033[48;2;139;0;0m"
	DarkSalmonBackground           = "\033[48;2;233;150;122m"
	DarkSeaGreenBackground         = "\033[48;2;143;188;143m"
	DarkSlateBlueBackground        = "\033[48;2;72;61;139m"
	DarkSlateGrayBackground        = "\033[48;2;47;79;79m"
	DarkTurquoiseBackground        = "\033[48;2;0;206;209m"
	DarkVioletBackground           = "\033[48;2;148;0;211m"
	DeepPinkBackground             = "\033[48;2;255;20;147m"
	DeepSkyBlueBackground          = "\033[48;2;0;191;255m"
	DimGrayBackground              = "\033[48;2;0;191;255m"
	DodgerBlueBackground           = "\033[48;2;30;144;255m"
	FirebrickBackground            = "\033[48;2;178;34;34m"
	FloralWhiteBackground          = "\033[48;2;255;250;240m"
	ForestGreenBackground          = "\033[48;2;34;139;34m"
	FuchsiaBackground              = "\033[48;2;255;0;255m"
	GainsboroBackground            = "\033[48;2;220;220;220m"
	GhostWhiteBackground           = "\033[48;2;248;248;255m"
	GoldBackground                 = "\033[48;2;255;215;0m"
	GoldenRodBackground            = "\033[48;2;218;165;32m"
	GrayBackground                 = "\033[48;2;127;127;127m"
	GreenBackground                = "\033[48;2;0;128;0m"
	GreenYellowBackground          = "\033[48;2;173;255;47m"
	HoneydewBackground             = "\033[48;2;240;255;240m"
	HotPinkBackground              = "\033[48;2;255;105;180m"
	IndianRedBackground            = "\033[48;2;205;92;92m"
	IndigoBackground               = "\033[48;2;75;0;130m"
	IvoryBackground                = "\033[48;2;255;255;240m"
	KhakiBackground                = "\033[48;2;240;230;140m"
	LavenderBackground             = "\033[48;2;230;230;250m"
	LavenderBlushBackground        = "\033[48;2;255;240;245m"
	LawnGreenBackground            = "\033[48;2;124;252;0m"
	LemonChiffonBackground         = "\033[48;2;255;250;205m"
	LightBlueBackground            = "\033[48;2;173;216;230m"
	LightCoralBackground           = "\033[48;2;240;128;128m"
	LightCyanBackground            = "\033[48;2;224;255;255m"
	LightGoldenRodYellowBackground = "\033[48;2;250;250;210m"
	LightGreenBackground           = "\033[48;2;144;238;144m"
	LightGreyBackground            = "\033[48;2;211;211;211m"
	LightPinkBackground            = "\033[48;2;255;182;193m"
	LightSalmonBackground          = "\033[48;2;255;160;122m"
	LightSeaGreenBackground        = "\033[48;2;32;178;170m"
	LightSkyBlueBackground         = "\033[48;2;135;206;250m"
	LightSlateGrayBackground       = "\033[48;2;119;136;153m"
	LightSteelBlueBackground       = "\033[48;2;176;196;222m"
	LightYellowBackground          = "\033[48;2;255;255;224m"
	LimeBackground                 = "\033[48;2;0;255;0m"
	LimeGreenBackground            = "\033[48;2;50;205;50m"
	LinenBackground                = "\033[48;2;250;240;230m"
	MagentaBackground              = "\033[48;2;255;0;255m"
	MaroonBackground               = "\033[48;2;128;0;0m"
	MediumAquaMarineBackground     = "\033[48;2;102;205;170m"
	MediumBlueBackground           = "\033[48;2;0;0;205m"
	MediumOrchidBackground         = "\033[48;2;186;85;211m"
	MediumPurpleBackground         = "\033[48;2;147;112;219m"
	MediumSeaGreenBackground       = "\033[48;2;60;179;113m"
	MediumSlateBlueBackground      = "\033[48;2;123;104;238m"
	MediumSpringGreenBackground    = "\033[48;2;0;250;154m"
	MediumTurquoiseBackground      = "\033[48;2;72;209;204m"
	MediumVioletRedBackground      = "\033[48;2;199;21;133m"
	MidnightBlueBackground         = "\033[48;2;25;25;112m"
	MintCreamBackground            = "\033[48;2;245;255;250m"
	MistyRoseBackground            = "\033[48;2;255;228;225m"
	MoccasinBackground             = "\033[48;2;255;228;181m"
	NavajoWhiteBackground          = "\033[48;2;255;222;173m"
	NavyBackground                 = "\033[48;2;0;0;128m"
	NavyBlueBackground             = "\033[48;2;159;175;223m"
	OldLaceBackground              = "\033[48;2;253;245;230m"
	OliveBackground                = "\033[48;2;128;128;0m"
	OliveDrabBackground            = "\033[48;2;107;142;35m"
	OrangeBackground               = "\033[48;2;255;165;0m"
	OrangeRedBackground            = "\033[48;2;255;69;0m"
	OrchidBackground               = "\033[48;2;218;112;214m"
	PaleGoldenRodBackground        = "\033[48;2;238;232;170m"
	PaleGreenBackground            = "\033[48;2;152;251;152m"
	PaleTurquoiseBackground        = "\033[48;2;175;238;238m"
	PaleVioletRedBackground        = "\033[48;2;219;112;147m"
	PapayaWhipBackground           = "\033[48;2;255;239;213m"
	PeachPuffBackground            = "\033[48;2;255;218;185m"
	PeruBackground                 = "\033[48;2;205;133;63m"
	PinkBackground                 = "\033[48;2;255;192;203m"
	PlumBackground                 = "\033[48;2;221;160;221m"
	PowderBlueBackground           = "\033[48;2;176;224;230m"
	PurpleBackground               = "\033[48;2;128;0;128m"
	RebeccaPurpleBackground        = "\033[48;2;102;51;153m"
	RedBackground                  = "\033[48;2;255;0;0m"
	RosyBrownBackground            = "\033[48;2;188;143;143m"
	RoyalBlueBackground            = "\033[48;2;65;105;225m"
	SaddleBrownBackground          = "\033[48;2;139;69;19m"
	SalmonBackground               = "\033[48;2;250;128;114m"
	SandyBrownBackground           = "\033[48;2;244;164;96m"
	SeaGreenBackground             = "\033[48;2;46;139;87m"
	SeaShellBackground             = "\033[48;2;255;245;238m"
	SiennaBackground               = "\033[48;2;160;82;45m"
	SilverBackground               = "\033[48;2;192;192;192m"
	SkyBlueBackground              = "\033[48;2;135;206;235m"
	SlateBlueBackground            = "\033[48;2;106;90;205m"
	SlateGrayBackground            = "\033[48;2;112;128;144m"
	SnowBackground                 = "\033[48;2;255;250;250m"
	SpringGreenBackground          = "\033[48;2;0;255;127m"
	SteelBlueBackground            = "\033[48;2;70;130;180m"
	TanBackground                  = "\033[48;2;210;180;140m"
	TealBackground                 = "\033[48;2;0;128;128m"
	ThistleBackground              = "\033[48;2;216;191;216m"
	TomatoBackground               = "\033[48;2;255;99;71m"
	TurquoiseBackground            = "\033[48;2;64;224;208m"
	VioletBackground               = "\033[48;2;238;130;238m"
	WheatBackground                = "\033[48;2;245;222;179m"
	WhiteBackground                = "\033[48;2;255;255;255m"
	WhiteSmokeBackground           = "\033[48;2;245;245;245m"
	YellowBackground               = "\033[48;2;255;255;0m"
	YellowGreenBackground          = "\033[48;2;139;205;50m"
)

var Index = map[string]string{
	"aliceblue":                      "\033[38;2;240;248;255m",
	"antiquewhite":                   "\033[38;2;250;235;215m",
	"aqua":                           "\033[38;2;0;255;255m",
	"aquamarine":                     "\033[38;2;127;255;212m",
	"azure":                          "\033[38;2;1240;255;255m",
	"beige":                          "\033[38;2;245;245;220m",
	"bisque":                         "\033[38;2;255;228;196m",
	"black":                          "\033[38;2;0;0;0m",
	"blanchedalmond":                 "\033[38;2;255;235;205m",
	"blue":                           "\033[38;2;0;0;255m",
	"blueviolet":                     "\033[38;2;138;43;226m",
	"brown":                          "\033[38;2;165;42;42m",
	"burlywood":                      "\033[38;2;222;184;135m",
	"cadetblue":                      "\033[38;2;95;158;160m",
	"chartreuse":                     "\033[38;2;95;158;160m",
	"chocolate":                      "\033[38;2;210;105;30m",
	"coral":                          "\033[38;2;255;127;80m",
	"cornflowerblue":                 "\033[38;2;100;149;237m",
	"cornsilk":                       "\033[38;2;255;248;220m",
	"crimson":                        "\033[38;2;220;20;60m",
	"cyan":                           "\033[38;2;0;255;255m",
	"darkblue":                       "\033[38;2;0;0;139m",
	"darkcyan":                       "\033[38;2;0;139;139m",
	"darkgoldenrod":                  "\033[38;2;184;134;11m",
	"darkgray":                       "\033[38;2;169;169;169m",
	"darkgreen":                      "\033[38;2;0;100;0m",
	"darkkhaki":                      "\033[38;2;189;183;107m",
	"darkmagenta":                    "\033[38;2;139;0;139m",
	"darkolivegreen":                 "\033[38;2;85;107;47m",
	"darkorange":                     "\033[38;2;255;140;0m",
	"darkorchid":                     "\033[38;2;153;50;204m",
	"darkred":                        "\033[38;2;139;0;0m",
	"darksalmon":                     "\033[38;2;233;150;122m",
	"darkseagreen":                   "\033[38;2;143;188;143m",
	"darkslateblue":                  "\033[38;2;72;61;139m",
	"darkslategray":                  "\033[38;2;47;79;79m",
	"darkturquoise":                  "\033[38;2;0;206;209m",
	"darkviolet":                     "\033[38;2;148;0;211m",
	"deeppink":                       "\033[38;2;255;20;147m",
	"deepskyblue":                    "\033[38;2;0;191;255m",
	"dimgray":                        "\033[38;2;0;191;255m",
	"dodgerblue":                     "\033[38;2;30;144;255m",
	"firebrick":                      "\033[38;2;178;34;34m",
	"floralwhite":                    "\033[38;2;255;250;240m",
	"forestgreen":                    "\033[38;2;34;139;34m",
	"fuchsia":                        "\033[38;2;255;0;255m",
	"gainsboro":                      "\033[38;2;220;220;220m",
	"ghostwhite":                     "\033[38;2;248;248;255m",
	"gold":                           "\033[38;2;255;215;0m",
	"goldenrod":                      "\033[38;2;218;165;32m",
	"gray":                           "\033[38;2;127;127;127m",
	"green":                          "\033[38;2;0;128;0m",
	"greenyellow":                    "\033[38;2;173;255;47m",
	"honeydew":                       "\033[38;2;240;255;240m",
	"hotpink":                        "\033[38;2;255;105;180m",
	"indianred":                      "\033[38;2;205;92;92m",
	"indigo":                         "\033[38;2;75;0;130m",
	"ivory":                          "\033[38;2;255;255;240m",
	"khaki":                          "\033[38;2;240;230;140m",
	"lavender":                       "\033[38;2;230;230;250m",
	"lavenderblush":                  "\033[38;2;255;240;245m",
	"lawngreen":                      "\033[38;2;124;252;0m",
	"lemonchiffon":                   "\033[38;2;255;250;205m",
	"lightblue":                      "\033[38;2;173;216;230m",
	"lightcoral":                     "\033[38;2;240;128;128m",
	"lightcyan":                      "\033[38;2;224;255;255m",
	"lightgoldenrodyellow":           "\033[38;2;250;250;210m",
	"lightgreen":                     "\033[38;2;144;238;144m",
	"lightgrey":                      "\033[38;2;211;211;211m",
	"lightpink":                      "\033[38;2;255;182;193m",
	"lightsalmon":                    "\033[38;2;255;160;122m",
	"lightseagreen":                  "\033[38;2;32;178;170m",
	"lightskyblue":                   "\033[38;2;135;206;250m",
	"lightslategray":                 "\033[38;2;119;136;153m",
	"lightsteelblue":                 "\033[38;2;176;196;222m",
	"lightyellow":                    "\033[38;2;255;255;224m",
	"lime":                           "\033[38;2;0;255;0m",
	"limegreen":                      "\033[38;2;50;205;50m",
	"linen":                          "\033[38;2;250;240;230m",
	"magenta":                        "\033[38;2;255;0;255m",
	"maroon":                         "\033[38;2;128;0;0m",
	"mediumaquamarine":               "\033[38;2;102;205;170m",
	"mediumblue":                     "\033[38;2;0;0;205m",
	"mediumorchid":                   "\033[38;2;186;85;211m",
	"mediumpurple":                   "\033[38;2;147;112;219m",
	"mediumseagreen":                 "\033[38;2;60;179;113m",
	"mediumslateblue":                "\033[38;2;123;104;238m",
	"mediumspringgreen":              "\033[38;2;0;250;154m",
	"mediumturquoise":                "\033[38;2;72;209;204m",
	"mediumvioletred":                "\033[38;2;199;21;133m",
	"midnightblue":                   "\033[38;2;25;25;112m",
	"mintcream":                      "\033[38;2;245;255;250m",
	"mistyrose":                      "\033[38;2;255;228;225m",
	"moccasin":                       "\033[38;2;255;228;181m",
	"navajowhite":                    "\033[38;2;255;222;173m",
	"navy":                           "\033[38;2;0;0;128m",
	"navyblue":                       "\033[38;2;159;175;223m",
	"oldlace":                        "\033[38;2;253;245;230m",
	"olive":                          "\033[38;2;128;128;0m",
	"olivedrab":                      "\033[38;2;107;142;35m",
	"orange":                         "\033[38;2;255;165;0m",
	"orangered":                      "\033[38;2;255;69;0m",
	"orchid":                         "\033[38;2;218;112;214m",
	"palegoldenrod":                  "\033[38;2;238;232;170m",
	"palegreen":                      "\033[38;2;152;251;152m",
	"paleturquoise":                  "\033[38;2;175;238;238m",
	"palevioletred":                  "\033[38;2;219;112;147m",
	"papayawhip":                     "\033[38;2;255;239;213m",
	"peachpuff":                      "\033[38;2;255;218;185m",
	"peru":                           "\033[38;2;205;133;63m",
	"pink":                           "\033[38;2;255;192;203m",
	"plum":                           "\033[38;2;221;160;221m",
	"powderblue":                     "\033[38;2;176;224;230m",
	"purple":                         "\033[38;2;128;0;128m",
	"red":                            "\033[38;2;255;0;0m",
	"rosybrown":                      "\033[38;2;188;143;143m",
	"royalblue":                      "\033[38;2;65;105;225m",
	"saddlebrown":                    "\033[38;2;139;69;19m",
	"salmon":                         "\033[38;2;250;128;114m",
	"sandybrown":                     "\033[38;2;244;164;96m",
	"seagreen":                       "\033[38;2;46;139;87m",
	"seashell":                       "\033[38;2;255;245;238m",
	"sienna":                         "\033[38;2;160;82;45m",
	"silver":                         "\033[38;2;192;192;192m",
	"skyblue":                        "\033[38;2;135;206;235m",
	"slateblue":                      "\033[38;2;106;90;205m",
	"slategray":                      "\033[38;2;112;128;144m",
	"snow":                           "\033[38;2;255;250;250m",
	"springgreen":                    "\033[38;2;0;255;127m",
	"steelblue":                      "\033[38;2;70;130;180m",
	"tan":                            "\033[38;2;210;180;140m",
	"teal":                           "\033[38;2;0;128;128m",
	"thistle":                        "\033[38;2;216;191;216m",
	"tomato":                         "\033[38;2;255;99;71m",
	"turquoise":                      "\033[38;2;64;224;208m",
	"violet":                         "\033[38;2;238;130;238m",
	"wheat":                          "\033[38;2;245;222;179m",
	"white":                          "\033[38;2;255;255;255m",
	"whitesmoke":                     "\033[38;2;245;245;245m",
	"yellow":                         "\033[38;2;255;255;0m",
	"yellowgreen":                    "\033[38;2;139;205;50m",
	"alicebluebackground":            "\033[48;2;240;248;255m",
	"antiquewhitebackground":         "\033[48;2;250;235;215m",
	"aquabackground":                 "\033[48;2;0;255;255m",
	"aquamarinebackground":           "\033[48;2;127;255;212m",
	"azurebackground":                "\033[48;2;1240;255;255m",
	"beigebackground":                "\033[48;2;245;245;220m",
	"bisquebackground":               "\033[48;2;255;228;196m",
	"blackbackground":                "\033[48;2;0;0;0m",
	"blanchedalmondbackground":       "\033[48;2;255;235;205m",
	"bluebackground":                 "\033[48;2;0;0;255m",
	"bluevioletbackground":           "\033[48;2;138;43;226m",
	"brownbackground":                "\033[48;2;165;42;42m",
	"burlywoodbackground":            "\033[48;2;222;184;135m",
	"cadetbluebackground":            "\033[48;2;95;158;160m",
	"chartreusebackground":           "\033[48;2;95;158;160m",
	"chocolatebackground":            "\033[48;2;210;105;30m",
	"coralbackground":                "\033[48;2;255;127;80m",
	"cornflowerbluebackground":       "\033[48;2;100;149;237m",
	"cornsilkbackground":             "\033[48;2;255;248;220m",
	"crimsonbackground":              "\033[48;2;220;20;60m",
	"cyanbackground":                 "\033[48;2;0;255;255m",
	"darkbluebackground":             "\033[48;2;0;0;139m",
	"darkcyanbackground":             "\033[48;2;0;139;139m",
	"darkgoldenrodbackground":        "\033[48;2;184;134;11m",
	"darkgraybackground":             "\033[48;2;169;169;169m",
	"darkgreenbackground":            "\033[48;2;0;100;0m",
	"darkkhakibackground":            "\033[48;2;189;183;107m",
	"darkmagentabackground":          "\033[48;2;139;0;139m",
	"darkolivegreenbackground":       "\033[48;2;85;107;47m",
	"darkorangebackground":           "\033[48;2;255;140;0m",
	"darkorchidbackground":           "\033[48;2;153;50;204m",
	"darkredbackground":              "\033[48;2;139;0;0m",
	"darksalmonbackground":           "\033[48;2;233;150;122m",
	"darkseagreenbackground":         "\033[48;2;143;188;143m",
	"darkslatebluebackground":        "\033[48;2;72;61;139m",
	"darkslategraybackground":        "\033[48;2;47;79;79m",
	"darkturquoisebackground":        "\033[48;2;0;206;209m",
	"darkvioletbackground":           "\033[48;2;148;0;211m",
	"deeppinkbackground":             "\033[48;2;255;20;147m",
	"deepskybluebackground":          "\033[48;2;0;191;255m",
	"dimgraybackground":              "\033[48;2;0;191;255m",
	"dodgerbluebackground":           "\033[48;2;30;144;255m",
	"firebrickbackground":            "\033[48;2;178;34;34m",
	"floralwhitebackground":          "\033[48;2;255;250;240m",
	"forestgreenbackground":          "\033[48;2;34;139;34m",
	"fuchsiabackground":              "\033[48;2;255;0;255m",
	"gainsborobackground":            "\033[48;2;220;220;220m",
	"ghostwhitebackground":           "\033[48;2;248;248;255m",
	"goldbackground":                 "\033[48;2;255;215;0m",
	"goldenrodbackground":            "\033[48;2;218;165;32m",
	"graybackground":                 "\033[48;2;127;127;127m",
	"greenbackground":                "\033[48;2;0;128;0m",
	"greenyellowbackground":          "\033[48;2;173;255;47m",
	"honeydewbackground":             "\033[48;2;240;255;240m",
	"hotpinkbackground":              "\033[48;2;255;105;180m",
	"indianredbackground":            "\033[48;2;205;92;92m",
	"indigobackground":               "\033[48;2;75;0;130m",
	"ivorybackground":                "\033[48;2;255;255;240m",
	"khakibackground":                "\033[48;2;240;230;140m",
	"lavenderbackground":             "\033[48;2;230;230;250m",
	"lavenderblushbackground":        "\033[48;2;255;240;245m",
	"lawngreenbackground":            "\033[48;2;124;252;0m",
	"lemonchiffonbackground":         "\033[48;2;255;250;205m",
	"lightbluebackground":            "\033[48;2;173;216;230m",
	"lightcoralbackground":           "\033[48;2;240;128;128m",
	"lightcyanbackground":            "\033[48;2;224;255;255m",
	"lightgoldenrodyellowbackground": "\033[48;2;250;250;210m",
	"lightgreenbackground":           "\033[48;2;144;238;144m",
	"lightgreybackground":            "\033[48;2;211;211;211m",
	"lightpinkbackground":            "\033[48;2;255;182;193m",
	"lightsalmonbackground":          "\033[48;2;255;160;122m",
	"lightseagreenbackground":        "\033[48;2;32;178;170m",
	"lightskybluebackground":         "\033[48;2;135;206;250m",
	"lightslategraybackground":       "\033[48;2;119;136;153m",
	"lightsteelbluebackground":       "\033[48;2;176;196;222m",
	"lightyellowbackground":          "\033[48;2;255;255;224m",
	"limebackground":                 "\033[48;2;0;255;0m",
	"limegreenbackground":            "\033[48;2;50;205;50m",
	"linenbackground":                "\033[48;2;250;240;230m",
	"magentabackground":              "\033[48;2;255;0;255m",
	"maroonbackground":               "\033[48;2;128;0;0m",
	"mediumaquamarinebackground":     "\033[48;2;102;205;170m",
	"mediumbluebackground":           "\033[48;2;0;0;205m",
	"mediumorchidbackground":         "\033[48;2;186;85;211m",
	"mediumpurplebackground":         "\033[48;2;147;112;219m",
	"mediumseagreenbackground":       "\033[48;2;60;179;113m",
	"mediumslatebluebackground":      "\033[48;2;123;104;238m",
	"mediumspringgreenbackground":    "\033[48;2;0;250;154m",
	"mediumturquoisebackground":      "\033[48;2;72;209;204m",
	"mediumvioletredbackground":      "\033[48;2;199;21;133m",
	"midnightbluebackground":         "\033[48;2;25;25;112m",
	"mintcreambackground":            "\033[48;2;245;255;250m",
	"mistyrosebackground":            "\033[48;2;255;228;225m",
	"moccasinbackground":             "\033[48;2;255;228;181m",
	"navajowhitebackground":          "\033[48;2;255;222;173m",
	"navybackground":                 "\033[48;2;0;0;128m",
	"navybluebackground":             "\033[48;2;159;175;223m",
	"oldlacebackground":              "\033[48;2;253;245;230m",
	"olivebackground":                "\033[48;2;128;128;0m",
	"olivedrabbackground":            "\033[48;2;107;142;35m",
	"orangebackground":               "\033[48;2;255;165;0m",
	"orangeredbackground":            "\033[48;2;255;69;0m",
	"orchidbackground":               "\033[48;2;218;112;214m",
	"palegoldenrodbackground":        "\033[48;2;238;232;170m",
	"palegreenbackground":            "\033[48;2;152;251;152m",
	"paleturquoisebackground":        "\033[48;2;175;238;238m",
	"palevioletredbackground":        "\033[48;2;219;112;147m",
	"papayawhipbackground":           "\033[48;2;255;239;213m",
	"peachpuffbackground":            "\033[48;2;255;218;185m",
	"perubackground":                 "\033[48;2;205;133;63m",
	"pinkbackground":                 "\033[48;2;255;192;203m",
	"plumbackground":                 "\033[48;2;221;160;221m",
	"powderbluebackground":           "\033[48;2;176;224;230m",
	"purplebackground":               "\033[48;2;128;0;128m",
	"redbackground":                  "\033[48;2;255;0;0m",
	"rosybrownbackground":            "\033[48;2;188;143;143m",
	"royalbluebackground":            "\033[48;2;65;105;225m",
	"saddlebrownbackground":          "\033[48;2;139;69;19m",
	"salmonbackground":               "\033[48;2;250;128;114m",
	"sandybrownbackground":           "\033[48;2;244;164;96m",
	"seagreenbackground":             "\033[48;2;46;139;87m",
	"seashellbackground":             "\033[48;2;255;245;238m",
	"siennabackground":               "\033[48;2;160;82;45m",
	"silverbackground":               "\033[48;2;192;192;192m",
	"skybluebackground":              "\033[48;2;135;206;235m",
	"slatebluebackground":            "\033[48;2;106;90;205m",
	"slategraybackground":            "\033[48;2;112;128;144m",
	"snowbackground":                 "\033[48;2;255;250;250m",
	"springgreenbackground":          "\033[48;2;0;255;127m",
	"steelbluebackground":            "\033[48;2;70;130;180m",
	"tanbackground":                  "\033[48;2;210;180;140m",
	"tealbackground":                 "\033[48;2;0;128;128m",
	"thistlebackground":              "\033[48;2;216;191;216m",
	"tomatobackground":               "\033[48;2;255;99;71m",
	"turquoisebackground":            "\033[48;2;64;224;208m",
	"violetbackground":               "\033[48;2;238;130;238m",
	"wheatbackground":                "\033[48;2;245;222;179m",
	"whitebackground":                "\033[48;2;255;255;255m",
	"whitesmokebackground":           "\033[48;2;245;245;245m",
	"yellowbackground":               "\033[48;2;255;255;0m",
	"yellowgreenbackground":          "\033[48;2;139;205;50m",
}

var HexIndex = map[string]int{
	"AliceBlue":            0xF0F8FF,
	"AntiqueWhite":         0xFAEBD7,
	"Aqua":                 0x00FFFF,
	"AquaMarine":           0x7FFFD4,
	"Azure":                0xF0FFFF,
	"Beige":                0xF5F5DC,
	"Black":                0x000000,
	"Bisque":               0xFFE4C4,
	"BlanchedAlmond":       0xFFEBCD,
	"Blue":                 0x0000FF,
	"BlueViolet":           0x8A2BE2,
	"Brown":                0xA52A2A,
	"BurlyWood":            0xDEB887,
	"CadetBlue":            0x5F9EA0,
	"Chartreuse":           0x7FFF00,
	"Chocolate":            0xD2691E,
	"Coral":                0xFF7F50,
	"CornFlowerBlue":       0x6495ED,
	"CornSilk":             0xFFF8DC,
	"Crimson":              0xDC143C,
	"Cyan":                 0x00FFFF,
	"DarkBlue":             0x00008B,
	"DarkCyan":             0x008B8B,
	"DarkGoldenRod":        0xB8860B,
	"DarkGray":             0xA9A9A9,
	"DarkGreen":            0x006400,
	"DarkKhaki":            0xBDB76B,
	"DarkMagenta":          0x8B008B,
	"DarkOliveGreen":       0x556B2F,
	"DarkOrange":           0xFF8C00,
	"DarkOrchid":           0x9932CC,
	"DarkRed":              0x8B0000,
	"DarkSalmon":           0xE9967A,
	"DarkSeaGreen":         0x8FBC8F,
	"DarkSlateBlue":        0x483D8B,
	"DarkslateGray":        0x2F4F4F,
	"DarkTurquoise":        0x00CED1,
	"DarkViolet":           0x9400D3,
	"DeepPink":             0xFF1493,
	"DeepSkyBlue":          0x00BFFF,
	"DimGray":              0x696969,
	"DodgerBlue":           0x1E90FF,
	"FireBrick":            0xB22222,
	"FloralWhite":          0xFFFAF0,
	"ForestGreen":          0x228B22,
	"Fuchsia":              0xFF00FF,
	"Gainsboro":            0xDCDCDC,
	"GhostWhite":           0xF8F8FF,
	"Gold":                 0xFFD700,
	"GoldenRod":            0xDAA520,
	"Gray":                 0x808080,
	"Green":                0x008000,
	"GreenYellow":          0xADFF2F,
	"HoneyDew":             0xF0FFF0,
	"HotPink":              0xFF69B4,
	"IndianRed":            0xCD5C5C,
	"Indigo":               0x4B0082,
	"Ivory":                0xFFFFF0,
	"Khaki":                0xF0E68C,
	"Lavender":             0xE6E6FA,
	"LavenderBlush":        0xFFF0F5,
	"LawnGreen":            0x7CFC00,
	"LemonChiffon":         0xFFFACD,
	"LightBlue":            0xADD8E6,
	"LightCoral":           0xF08080,
	"LightCyan":            0xE0FFFF,
	"LightGoldenRodYellow": 0xFAFAD2,
	"LightGray":            0xD3D3D3,
	"LightGreen":           0x90EE90,
	"LightPink":            0xFFB6C1,
	"LightSalmon":          0xFFA07A,
	"LightSeaGreen":        0x20B2AA,
	"LightSkyBlue":         0x87CEFA,
	"LightSlateGray":       0x778899,
	"LightSteelBlue":       0xB0C4DE,
	"LightYellow":          0xFFFFE0,
	"Lime":                 0x00FF00,
	"LimeGreen":            0x32CD32,
	"Linen":                0xFAF0E6,
	"Magenta":              0xFF00FF,
	"Maroon":               0x800000,
	"MediumAquaMarine":     0x66CDAA,
	"MediumBlue":           0x0000CD,
	"MediumOrchid":         0xBA55D3,
	"MediumPurple":         0x9370DB,
	"MediumSeaGreen":       0x3CB371,
	"MediumSlateBlue":      0x7B68EE,
	"MediumSpringGreen":    0x00FA9A,
	"MediumTurquoise":      0x48D1CC,
	"MediumVioletRed":      0xC71585,
	"MidnightBlue":         0x191970,
	"MintCream":            0xF5FFFA,
	"MistyRose":            0xFFE4E1,
	"Moccasin":             0xFFE4B5,
	"NavajoWhite":          0xFFDEAD,
	"Navy":                 0x000080,
	"NavyBlue":             0x9FAFDF,
	"OldLace":              0xFDF5E6,
	"Olive":                0x808000,
	"OliveDrab":            0x6B8E23,
	"Orange":               0xFFA500,
	"OrangeRed":            0xFF4500,
	"Orchid":               0xDA70D6,
	"PaleGoldenRod":        0xEEE8AA,
	"PaleGreen":            0x98FB98,
	"PaleTurquoise":        0xAFEEEE,
	"PaleVioletRed":        0xDB7093,
	"PapayaWhip":           0xFFEFD5,
	"PeachPuff":            0xFFDAB9,
	"Peru":                 0xCD853F,
	"Pink":                 0xFFC0CB,
	"Plum":                 0xDDA0DD,
	"PowderBlue":           0xB0E0E6,
	"Purple":               0x800080,
	"Red":                  0xFF0000,
	"RebeccaPurple":        0x663399,
	"RosyBrown":            0xBC8F8F,
	"RoyalBlue":            0x4169E1,
	"SaddleBrown":          0x8B4513,
	"Salmon":               0xFA8072,
	"SandyBrown":           0xF4A460,
	"SeaGreen":             0x2E8B57,
	"SeaShell":             0xFFF5EE,
	"Sienna":               0xA0522D,
	"Silver":               0xC0C0C0,
	"SkyBlue":              0x87CEEB,
	"SlateBlue":            0x6A5ACD,
	"SlateGray":            0x708090,
	"Snow":                 0xFFFAFA,
	"SpringGreen":          0x00FF7F,
	"SteelBlue":            0x4682B4,
	"Tan":                  0xD2B48C,
	"Teal":                 0x008080,
	"Thistle":              0xD8BFD8,
	"Tomato":               0xFF6347,
	"Turquoise":            0x40E0D0,
	"Violet":               0xEE82EE,
	"Wheat":                0xF5DEB3,
	"White":                0xFFFFFF,
	"WhiteSmoke":           0xF5F5F5,
	"Yellow":               0xFFFF00,
	"YellowGreen":          0x9ACD32,
}

var Colors = [...]string{AliceBlue, AntiqueWhite, Aqua, AquaMarine, Azure, Beige, Bisque, Black, BlanchedAlmond, Blue, BlueViolet, Brown, BurlyWood, CadetBlue, Chartreuse, Chocolate, Coral, CornFlowerBlue, CornSilk, Crimson, Cyan, DarkBlue, DarkCyan, DarkGoldenRod, DarkGray, DarkGreen, DarkKhaki, DarkMagenta, DarkOliveGreen, DarkOrange, DarkOrchid, DarkRed, DarkSalmon, DarkSeaGreen, DarkSlateBlue, DarkSlateGray, DarkTurquoise, DarkViolet, DeepPink, DeepSkyBlue, DimGray, DodgerBlue, Firebrick, FloralWhite, ForestGreen, Fuchsia, Gainsboro, GhostWhite, Gold, GoldenRod, Gray, Green, GreenYellow, Honeydew, HotPink, IndianRed, Indigo, Ivory, Khaki, Lavender, LavenderBlush, LawnGreen, LemonChiffon, LightBlue, LightCoral, LightCyan, LightGoldenRodYellow, LightGreen, LightGrey, LightPink, LightSalmon, LightSeaGreen, LightSkyBlue, LightSlateGray, LightSteelBlue, LightYellow, Lime, LimeGreen, Linen, Magenta, Maroon, MediumAquaMarine, MediumBlue, MediumOrchid, MediumPurple, MediumSeaGreen, MediumSlateBlue, MediumSpringGreen, MediumTurquoise, MediumVioletRed, MidnightBlue, MintCream, MistyRose, Moccasin, NavajoWhite, Navy, NavyBlue, OldLace, Olive, OliveDrab, Orange, OrangeRed, Orchid, PaleGoldenRod, PaleGreen, PaleTurquoise, PaleVioletRed, PapayaWhip, PeachPuff, Peru, Pink, Plum, PowderBlue, Purple, Red, RosyBrown, RoyalBlue, SaddleBrown, Salmon, SandyBrown, SeaGreen, SeaShell, Sienna, Silver, SkyBlue, SlateBlue, SlateGray, Snow, SpringGreen, SteelBlue, Tan, Teal, Thistle, Tomato, Turquoise, Violet, Wheat, White, WhiteSmoke, Yellow, YellowGreen}

var BackgroundColors = [...]string{AliceBlueBackground, AntiqueWhiteBackground, AquaBackground, AquaMarineBackground, AzureBackground, BeigeBackground, BisqueBackground, BlackBackground, BlanchedAlmondBackground, BlueBackground, BlueVioletBackground, BrownBackground, BurlyWoodBackground, CadetBlueBackground, ChartreuseBackground, ChocolateBackground, CoralBackground, CornFlowerBlueBackground, CornSilkBackground, CrimsonBackground, CyanBackground, DarkBlueBackground, DarkCyanBackground, DarkGoldenRodBackground, DarkGrayBackground, DarkGreenBackground, DarkKhakiBackground, DarkMagentaBackground, DarkOliveGreenBackground, DarkOrangeBackground, DarkOrchidBackground, DarkRedBackground, DarkSalmonBackground, DarkSeaGreenBackground, DarkSlateBlueBackground, DarkSlateGrayBackground, DarkTurquoiseBackground, DarkVioletBackground, DeepPinkBackground, DeepSkyBlueBackground, DimGrayBackground, DodgerBlueBackground, FirebrickBackground, FloralWhiteBackground, ForestGreenBackground, FuchsiaBackground, GainsboroBackground, GhostWhiteBackground, GoldBackground, GoldenRodBackground, GrayBackground, GreenBackground, GreenYellowBackground, HoneydewBackground, HotPinkBackground, IndianRedBackground, IndigoBackground, IvoryBackground, KhakiBackground, LavenderBackground, LavenderBlushBackground, LawnGreenBackground, LemonChiffonBackground, LightBlueBackground, LightCoralBackground, LightCyanBackground, LightGoldenRodYellowBackground, LightGreenBackground, LightGreyBackground, LightPinkBackground, LightSalmonBackground, LightSeaGreenBackground, LightSkyBlueBackground, LightSlateGrayBackground, LightSteelBlueBackground, LightYellowBackground, LimeBackground, LimeGreenBackground, LinenBackground, MagentaBackground, MaroonBackground, MediumAquaMarineBackground, MediumBlueBackground, MediumOrchidBackground, MediumPurpleBackground, MediumSeaGreenBackground, MediumSlateBlueBackground, MediumSpringGreenBackground, MediumTurquoiseBackground, MediumVioletRedBackground, MidnightBlueBackground, MintCreamBackground, MistyRoseBackground, MoccasinBackground, NavajoWhiteBackground, NavyBackground, NavyBlueBackground, OldLaceBackground, OliveBackground, OliveDrabBackground, OrangeBackground, OrangeRedBackground, OrchidBackground, PaleGoldenRodBackground, PaleGreenBackground, PaleTurquoiseBackground, PaleVioletRedBackground, PapayaWhipBackground, PeachPuffBackground, PeruBackground, PinkBackground, PlumBackground, PowderBlueBackground, PurpleBackground, RedBackground, RosyBrownBackground, RoyalBlueBackground, SaddleBrownBackground, SalmonBackground, SandyBrownBackground, SeaGreenBackground, SeaShellBackground, SiennaBackground, SilverBackground, SkyBlueBackground, SlateBlueBackground, SlateGrayBackground, SnowBackground, SpringGreenBackground, SteelBlueBackground, TanBackground, TealBackground, ThistleBackground, TomatoBackground, TurquoiseBackground, VioletBackground, WheatBackground, WhiteBackground, WhiteSmokeBackground, YellowBackground, YellowGreenBackground}

func RandomColor() string {
	rand.Seed(time.Now().UnixNano())
	return Colors[rand.Intn(len(Colors))]
}

func RandomBackground() string {
	rand.Seed(time.Now().UnixNano())
	return BackgroundColors[rand.Intn(len(BackgroundColors))]
}
