package color

import (
	"strings"

	"gitlab.com/skilstak/code/go/color/sol"
)

var colorReplacer *strings.Replacer
var ansiReplacer *strings.Replacer

func init() {
	colorReplacer = strings.NewReplacer(
		sol.Base03, "",
		sol.Base02, "",
		sol.Base01, "",
		sol.Base00, "",
		sol.Base0, "",
		sol.Base1, "",
		sol.Base2, "",
		sol.Base3, "",
		sol.Yellow, "",
		sol.Orange, "",
		sol.Red, "",
		sol.Magenta, "",
		sol.Violet, "",
		sol.Blue, "",
		sol.Cyan, "",
		sol.Green, "",
		sol.Reset, "",
	)

	ansiReplacer = strings.NewReplacer(
		sol.ClearScreen, "",
		sol.ClearLine, "",
		sol.CursorOff, "",
		sol.CursorOn, "",
		sol.StrikeOut, "",
	)
}

// Decolor removes any of the ANSI color escapes known to this package.
func Decolor(s string) string {
	return colorReplacer.Replace(s)
}

// Strip strips any ANSI escapes known to this package.
// TODO strip all ANSI escapes
func Strip(s string) string {
	return colorReplacer.Replace(ansiReplacer.Replace(s))
}

func GlobalIndex(s string) string {
        if val, ok := Index[strings.ToLower(s)]; ok {
                return val + strings.ToLower(s)
        } else {
                return s + " is not a valid color. To see what colors are valid, run the sample? command."
        }
}

func RGB(r int g int b int) string {
        return "\x1b[38;2;" + r + ";" + g + ";" + b ";m"
}

var Index = map[string]string{
	"sol-yellow":               "\033[0;33m",
	"sol-orange":               "\033[1;31m",
	"sol-red":                  "\033[0;31m",
	"sol-magenta":              "\033[0;35m",
	"sol-violet":               "\033[1;35m",
	"sol-blue":                 "\033[0;34m",
	"sol-cyan":                 "\033[0;36m",
	"sol-green":                "\033[0;32m",
	"css-aliceblue":            "\x1b[38;2;240;248;255m",
	"css-antiquewhite":         "\x1b[38;2;250;235;215m",
	"css-aqua":                 "\x1b[38;2;0;255;255m",
	"css-aquamarine":           "\x1b[38;2;127;255;212m",
	"css-azure":                "\x1b[38;2;1240;255;255m",
	"css-beige":                "\x1b[38;2;245;245;220m",
	"css-bisque":               "\x1b[38;2;255;228;196m",
	"css-black":                "\x1b[38;2;0;0;0m",
	"css-blanchedalmond":       "\x1b[38;2;255;235;205m",
	"css-blue":                 "\x1b[38;2;0;0;255m",
	"css-blueviolet":           "\x1b[38;2;138;43;226m",
	"css-brown":                "\x1b[38;2;165;42;42m",
	"css-burlywood":            "\x1b[38;2;222;184;135m",
	"css-cadetblue":            "\x1b[38;2;95;158;160m",
	"css-chartreuse":           "\x1b[38;2;95;158;160m",
	"css-chocolate":            "\x1b[38;2;210;105;30m",
	"css-coral":                "\x1b[38;2;255;127;80m",
	"css-cornflowerblue":       "\x1b[38;2;100;149;237m",
	"css-cornsilk":             "\x1b[38;2;255;248;220m",
	"css-crimson":              "\x1b[38;2;220;20;60m",
	"css-cyan":                 "\x1b[38;2;0;255;255m",
	"css-darkblue":             "\x1b[38;2;0;0;139m",
	"css-darkcyan":             "\x1b[38;2;0;139;139m",
	"css-darkgoldenrod":        "\x1b[38;2;184;134;11m",
	"css-darkgray":             "\x1b[38;2;169;169;169m",
	"css-darkgreen":            "\x1b[38;2;0;100;0m",
	"css-darkkhaki":            "\x1b[38;2;189;183;107m",
	"css-darkmagenta":          "\x1b[38;2;139;0;139m",
	"css-darkolivegreen":       "\x1b[38;2;85;107;47m",
	"css-darkorange":           "\x1b[38;2;255;140;0m",
	"css-darkorchid":           "\x1b[38;2;153;50;204m",
	"css-darkred":              "\x1b[38;2;139;0;0m",
	"css-darksalmon":           "\x1b[38;2;233;150;122m",
	"css-darkseagreen":         "\x1b[38;2;143;188;143m",
	"css-darkslateblue":        "\x1b[38;2;72;61;139m",
	"css-darkslategray":        "\x1b[38;2;47;79;79m",
	"css-darkturquoise":        "\x1b[38;2;0;206;209m",
	"css-darkviolet":           "\x1b[38;2;148;0;211m",
	"css-deeppink":             "\x1b[38;2;255;20;147m",
	"css-deepskyblue":          "\x1b[38;2;0;191;255m",
	"css-dimgray":              "\x1b[38;2;0;191;255m",
	"css-dodgerblue":           "\x1b[38;2;30;144;255m",
	"css-firebrick":            "\x1b[38;2;178;34;34m",
	"css-floralwhite":          "\x1b[38;2;255;250;240m",
	"css-forestgreen":          "\x1b[38;2;34;139;34m",
	"css-fuchsia":              "\x1b[38;2;255;0;255m",
	"css-gainsboro":            "\x1b[38;2;220;220;220m",
	"css-ghostwhite":           "\x1b[38;2;248;248;255m",
	"css-gold":                 "\x1b[38;2;255;215;0m",
	"css-goldenrod":            "\x1b[38;2;218;165;32m",
	"css-gray":                 "\x1b[38;2;127;127;127m",
	"css-green":                "\x1b[38;2;0;128;0m",
	"css-greenyellow":          "\x1b[38;2;173;255;47m",
	"css-honeydew":             "\x1b[38;2;240;255;240m",
	"css-hotpink":              "\x1b[38;2;255;105;180m",
	"css-indianred":            "\x1b[38;2;205;92;92m",
	"css-indigo":               "\x1b[38;2;75;0;130m",
	"css-ivory":                "\x1b[38;2;255;255;240m",
	"css-khaki":                "\x1b[38;2;240;230;140m",
	"css-lavender":             "\x1b[38;2;230;230;250m",
	"css-lavenderblush":        "\x1b[38;2;255;240;245m",
	"css-lawngreen":            "\x1b[38;2;124;252;0m",
	"css-lemonchiffon":         "\x1b[38;2;255;250;205m",
	"css-lightblue":            "\x1b[38;2;173;216;230m",
	"css-lightcoral":           "\x1b[38;2;240;128;128m",
	"css-lightcyan":            "\x1b[38;2;224;255;255m",
	"css-lightgoldenrodyellow": "\x1b[38;2;250;250;210m",
	"css-lightgreen":           "\x1b[38;2;144;238;144m",
	"css-lightgrey":            "\x1b[38;2;211;211;211m",
	"css-lightpink":            "\x1b[38;2;255;182;193m",
	"css-lightsalmon":          "\x1b[38;2;255;160;122m",
	"css-lightseagreen":        "\x1b[38;2;32;178;170m",
	"css-lightskyblue":         "\x1b[38;2;135;206;250m",
	"css-lightslategray":       "\x1b[38;2;119;136;153m",
	"css-lightsteelblue":       "\x1b[38;2;176;196;222m",
	"css-lightyellow":          "\x1b[38;2;255;255;224m",
	"css-lime":                 "\x1b[38;2;0;255;0m",
	"css-limegreen":            "\x1b[38;2;50;205;50m",
	"css-linen":                "\x1b[38;2;250;240;230m",
	"css-magenta":              "\x1b[38;2;255;0;255m",
	"css-maroon":               "\x1b[38;2;128;0;0m",
	"css-mediumaquamarine":     "\x1b[38;2;102;205;170m",
	"css-mediumblue":           "\x1b[38;2;0;0;205m",
	"css-mediumorchid":         "\x1b[38;2;186;85;211m",
	"css-mediumpurple":         "\x1b[38;2;147;112;219m",
	"css-mediumseagreen":       "\x1b[38;2;60;179;113m",
	"css-mediumslateblue":      "\x1b[38;2;123;104;238m",
	"css-mediumspringgreen":    "\x1b[38;2;0;250;154m",
	"css-mediumturquoise":      "\x1b[38;2;72;209;204m",
	"css-mediumvioletred":      "\x1b[38;2;199;21;133m",
	"css-midnightblue":         "\x1b[38;2;25;25;112m",
	"css-mintcream":            "\x1b[38;2;245;255;250m",
	"css-mistyrose":            "\x1b[38;2;255;228;225m",
	"css-moccasin":             "\x1b[38;2;255;228;181m",
	"css-navajowhite":          "\x1b[38;2;255;222;173m",
	"css-navy":                 "\x1b[38;2;0;0;128m",
	"css-navyblue":             "\x1b[38;2;159;175;223m",
	"css-oldlace":              "\x1b[38;2;253;245;230m",
	"css-olive":                "\x1b[38;2;128;128;0m",
	"css-olivedrab":            "\x1b[38;2;107;142;35m",
	"css-orange":               "\x1b[38;2;255;165;0m",
	"css-orangered":            "\x1b[38;2;255;69;0m",
	"css-orchid":               "\x1b[38;2;218;112;214m",
	"css-palegoldenrod":        "\x1b[38;2;238;232;170m",
	"css-palegreen":            "\x1b[38;2;152;251;152m",
	"css-paleturquoise":        "\x1b[38;2;175;238;238m",
	"css-palevioletred":        "\x1b[38;2;219;112;147m",
	"css-papayawhip":           "\x1b[38;2;255;239;213m",
	"css-peachpuff":            "\x1b[38;2;255;218;185m",
	"css-peru":                 "\x1b[38;2;205;133;63m",
	"css-pink":                 "\x1b[38;2;255;192;203m",
	"css-plum":                 "\x1b[38;2;221;160;221m",
	"css-powderblue":           "\x1b[38;2;176;224;230m",
	"css-purple":               "\x1b[38;2;128;0;128m",
	"css-red":                  "\x1b[38;2;255;0;0m",
	"css-rosybrown":            "\x1b[38;2;188;143;143m",
	"css-royalblue":            "\x1b[38;2;65;105;225m",
	"css-saddlebrown":          "\x1b[38;2;139;69;19m",
	"css-salmon":               "\x1b[38;2;250;128;114m",
	"css-sandybrown":           "\x1b[38;2;244;164;96m",
	"css-seagreen":             "\x1b[38;2;46;139;87m",
	"css-seashell":             "\x1b[38;2;255;245;238m",
	"css-sienna":               "\x1b[38;2;160;82;45m",
	"css-silver":               "\x1b[38;2;192;192;192m",
	"css-skyblue":              "\x1b[38;2;135;206;235m",
	"css-slateblue":            "\x1b[38;2;106;90;205m",
	"css-slategray":            "\x1b[38;2;112;128;144m",
	"css-snow":                 "\x1b[38;2;255;250;250m",
	"css-springgreen":          "\x1b[38;2;0;255;127m",
	"css-steelblue":            "\x1b[38;2;70;130;180m",
	"css-tan":                  "\x1b[38;2;210;180;140m",
	"css-teal":                 "\x1b[38;2;0;128;128m",
	"css-thistle":              "\x1b[38;2;216;191;216m",
	"css-tomato":               "\x1b[38;2;255;99;71m",
	"css-turquoise":            "\x1b[38;2;64;224;208m",
	"css-violet":               "\x1b[38;2;238;130;238m",
	"css-wheat":                "\x1b[38;2;245;222;179m",
	"css-white":                "\x1b[38;2;255;255;255m",
	"css-whitesmoke":           "\x1b[38;2;245;245;245m",
	"css-yellow":               "\x1b[38;2;255;255;0m",
	"css-yellowgreen":          "\x1b[38;2;139;205;50m",
}
