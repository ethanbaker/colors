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
	cadetblue            = "\x1b[38;2;95;158;160m"
	chartreuse           = "\x1b[38;2;95;158;160m"
	chocolate            = "\x1b[38;2;210;105;30m"
	coral                = "\x1b[38;2;255;127;80m"
	cornflowerblue       = "\x1b[38;2;100;149;237m"
	cornsilk             = "\x1b[38;2;255;248;220m"
	crimson              = "\x1b[38;2;220;20;60m"
	cyan                 = "\x1b[38;2;0;255;255m"
	darkblue             = "\x1b[38;2;0;0;139m"
	darkcyan             = "\x1b[38;2;0;139;139m"
	darkgoldenrod        = "\x1b[38;2;184;134;11m"
	darkgray             = "\x1b[38;2;169;169;169m"
	darkgreen            = "\x1b[38;2;0;100;0m"
	darkkhaki            = "\x1b[38;2;189;183;107m"
	darkmagenta          = "\x1b[38;2;139;0;139m"
	darkolivegreen       = "\x1b[38;2;85;107;47m"
	darkorange           = "\x1b[38;2;255;140;0m"
	darkorchid           = "\x1b[38;2;153;50;204m"
	darkred              = "\x1b[38;2;139;0;0m"
	darksalmon           = "\x1b[38;2;233;150;122m"
	darkseagreen         = "\x1b[38;2;143;188;143m"
	darkslateblue        = "\x1b[38;2;72;61;139m"
	darkslategray        = "\x1b[38;2;47;79;79m"
	darkturquoise        = "\x1b[38;2;0;206;209m"
	darkviolet           = "\x1b[38;2;148;0;211m"
	deeppink             = "\x1b[38;2;255;20;147m"
	deepskyblue          = "\x1b[38;2;0;191;255m"
	dimgray              = "\x1b[38;2;0;191;255m"
	dodgerblue           = "\x1b[38;2;30;144;255m"
	firebrick            = "\x1b[38;2;178;34;34m"
	floralwhite          = "\x1b[38;2;255;250;240m"
	forestgreen          = "\x1b[38;2;34;139;34m"
	fuchsia              = "\x1b[38;2;255;0;255m"
	gainsboro            = "\x1b[38;2;220;220;220m"
	ghostwhite           = "\x1b[38;2;248;248;255m"
	gold                 = "\x1b[38;2;255;215;0m"
	goldenrod            = "\x1b[38;2;218;165;32m"
	gray                 = "\x1b[38;2;127;127;127m"
	green                = "\x1b[38;2;0;128;0m"
	greenyellow          = "\x1b[38;2;173;255;47m"
	honeydew             = "\x1b[38;2;240;255;240m"
	hotpink              = "\x1b[38;2;255;105;180m"
	indianred            = "\x1b[38;2;205;92;92m"
	indigo               = "\x1b[38;2;75;0;130m"
	ivory                = "\x1b[38;2;255;255;240m"
	khaki                = "\x1b[38;2;240;230;140m"
	lavender             = "\x1b[38;2;230;230;250m"
	lavenderblush        = "\x1b[38;2;255;240;245m"
	lawngreen            = "\x1b[38;2;124;252;0m"
	lemonchiffon         = "\x1b[38;2;255;250;205m"
	lightblue            = "\x1b[38;2;173;216;230m"
	lightcoral           = "\x1b[38;2;240;128;128m"
	lightcyan            = "\x1b[38;2;224;255;255m"
	lightgoldenrodyellow = "\x1b[38;2;250;250;210m"
	lightgreen           = "\x1b[38;2;144;238;144m"
	lightgrey            = "\x1b[38;2;211;211;211m"
	lightpink            = "\x1b[38;2;255;182;193m"
	lightsalmon          = "\x1b[38;2;255;160;122m"
	lightseagreen        = "\x1b[38;2;32;178;170m"
	lightskyblue         = "\x1b[38;2;135;206;250m"
	lightslategray       = "\x1b[38;2;119;136;153m"
	lightsteelblue       = "\x1b[38;2;176;196;222m"
	lightyellow          = "\x1b[38;2;255;255;224m"
	lime                 = "\x1b[38;2;0;255;0m"
	limegreen            = "\x1b[38;2;50;205;50m"
	linen                = "\x1b[38;2;250;240;230m"
	magenta              = "\x1b[38;2;255;0;255m"
	maroon               = "\x1b[38;2;128;0;0m"
	mediumaquamarine     = "\x1b[38;2;102;205;170m"
	mediumblue           = "\x1b[38;2;0;0;205m"
	mediumorchid         = "\x1b[38;2;186;85;211m"
	mediumpurple         = "\x1b[38;2;147;112;219m"
	mediumseagreen       = "\x1b[38;2;60;179;113m"
	mediumslateblue      = "\x1b[38;2;123;104;238m"
	mediumspringgreen    = "\x1b[38;2;0;250;154m"
	mediumturquoise      = "\x1b[38;2;72;209;204m"
	mediumvioletred      = "\x1b[38;2;199;21;133m"
	midnightblue         = "\x1b[38;2;25;25;112m"
	mintcream            = "\x1b[38;2;245;255;250m"
	mistyrose            = "\x1b[38;2;255;228;225m"
	moccasin             = "\x1b[38;2;255;228;181m"
	navajowhite          = "\x1b[38;2;255;222;173m"
	navy                 = "\x1b[38;2;0;0;128m"
	navyblue             = "\x1b[38;2;159;175;223m"
	oldlace              = "\x1b[38;2;253;245;230m"
	olive                = "\x1b[38;2;128;128;0m"
	olivedrab            = "\x1b[38;2;107;142;35m"
	orange               = "\x1b[38;2;255;165;0m"
	orangered            = "\x1b[38;2;255;69;0m"
	orchid               = "\x1b[38;2;218;112;214m"
	palegoldenrod        = "\x1b[38;2;238;232;170m"
	palegreen            = "\x1b[38;2;152;251;152m"
	paleturquoise        = "\x1b[38;2;175;238;238m"
	palevioletred        = "\x1b[38;2;219;112;147m"
	papayawhip           = "\x1b[38;2;255;239;213m"
	peachpuff            = "\x1b[38;2;255;218;185m"
	peru                 = "\x1b[38;2;205;133;63m"
	pink                 = "\x1b[38;2;255;192;203m"
	plum                 = "\x1b[38;2;221;160;221m"
	powderblue           = "\x1b[38;2;176;224;230m"
	purple               = "\x1b[38;2;128;0;128m"
	red                  = "\x1b[38;2;255;0;0m"
	rosybrown            = "\x1b[38;2;188;143;143m"
	royalblue            = "\x1b[38;2;65;105;225m"
	saddlebrown          = "\x1b[38;2;139;69;19m"
	salmon               = "\x1b[38;2;250;128;114m"
	sandybrown           = "\x1b[38;2;244;164;96m"
	seagreen             = "\x1b[38;2;46;139;87m"
	seashell             = "\x1b[38;2;255;245;238m"
	sienna               = "\x1b[38;2;160;82;45m"
	silver               = "\x1b[38;2;192;192;192m"
	skyblue              = "\x1b[38;2;135;206;235m"
	slateblue            = "\x1b[38;2;106;90;205m"
	slategray            = "\x1b[38;2;112;128;144m"
	snow                 = "\x1b[38;2;255;250;250m"
	springgreen          = "\x1b[38;2;0;255;127m"
	steelblue            = "\x1b[38;2;70;130;180m"
	tan                  = "\x1b[38;2;210;180;140m"
	teal                 = "\x1b[38;2;0;128;128m"
	thistle              = "\x1b[38;2;216;191;216m"
	tomato               = "\x1b[38;2;255;99;71m"
	turquoise            = "\x1b[38;2;64;224;208m"
	violet               = "\x1b[38;2;238;130;238m"
	wheat                = "\x1b[38;2;245;222;179m"
	white                = "\x1b[38;2;255;255;255m"
	whitesmoke           = "\x1b[38;2;245;245;245m"
	yellow               = "\x1b[38;2;255;255;0m"
	yellowgreen          = "\x1b[38;2;139;205;50m"
)

//all of the css named colors are assigned a rbg value using TrueColor

var Colors = []string{aliceblue, antiquewhite, aqua, aquamarine, azure, beige, bisque, black, blanchedalmond, blue, blueviolet, brown, burlywood, cadetblue, chartreuse, chocolate, coral, cornflowerblue, cornsilk, crimson, cyan, darkblue, darkcyan, darkgoldenrod, darkgray, darkgreen, darkkhaki, darkmagenta, darkolivegreen, darkorange, darkorchid, darkred, darksalmon, darkseagreen, darkslateblue, darkslategray, darkturquoise, darkviolet, deeppink, deepskyblue, dimgray, dodgerblue, firebrick, floralwhite, forestgreen, fuchsia, gainsboro, ghostwhite, gold, goldenrod, gray, green, greenyellow, honeydew, hotpink, indianred, indigo, ivory, khaki, lavender, lavenderblush, lawngreen, lemonchiffon, lightblue, lightcoral, lightcyan, lightgoldenrodyellow, lightgreen, lightgrey, lightpink, lightsalmon, lightseagreen, lightskyblue, lightslategray, lightsteelblue, lightyellow, lime, limegreen, linen, magenta, maroon, mediumaquamarine, mediumblue, mediumorchid, mediumpurple, mediumseagreen, mediumslateblue, mediumspringgreen, mediumturquoise, mediumvioletred, midnightblue, mintcream, mistyrose, moccasin, navajowhite, navy, navyblue, oldlace, olive, olivedrab, orange, orangered, orchid, palegoldenrod, palegreen, paleturquoise, palevioletred, papayawhip, peachpuff, peru, pink, plum, powderblue, purple, red, rosybrown, royalblue, saddlebrown, salmon, sandybrown, seagreen, seashell, sienna, silver, skyblue, slateblue, slategray, snow, springgreen, steelblue, tan, teal, thistle, tomato, turquoise, violet, wheat, white, whitesmoke, yellow, yellowgreen}

func Random() string {
	rand.Seed(time.Now().UnixNano())
	return Colors[rand.Intn(len(Colors))]
}