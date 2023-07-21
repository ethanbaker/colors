package color

import (
	"strconv"
	"strings"

	"github.com/ethanbaker/colors/css"
	"github.com/ethanbaker/colors/sol"
)

const (
	ClearScreen = "\033[2J\033[H"
	ClearLine   = "\033[2K\033[G"
	CursorOff   = "\033[?25l"
	CursorOn    = "\033[?25h"
	StrikeOut   = "\033[9m"
)

var colorReplacerAnsiSol *strings.Replacer
var colorReplacerAnsiCss *strings.Replacer

func init() {
	colorReplacerAnsiSol = strings.NewReplacer(
		sol.AnsiMap["base03"], "",
		sol.AnsiMap["base02"], "",
		sol.AnsiMap["base01"], "",
		sol.AnsiMap["base00"], "",
		sol.AnsiMap["base0"], "",
		sol.AnsiMap["base1"], "",
		sol.AnsiMap["base2"], "",
		sol.AnsiMap["base3"], "",
		sol.AnsiMap["yellow"], "",
		sol.AnsiMap["orange"], "",
		sol.AnsiMap["red"], "",
		sol.AnsiMap["magenta"], "",
		sol.AnsiMap["violet"], "",
		sol.AnsiMap["blue"], "",
		sol.AnsiMap["cyan"], "",
		sol.AnsiMap["green"], "",
		sol.AnsiMap["reset"], "",
	)

	colorReplacerAnsiCss = strings.NewReplacer(
		css.AnsiMap["aliceblue"], "",
		css.AnsiMap["antiquewhite"], "",
		css.AnsiMap["aqua"], "",
		css.AnsiMap["aquamarine"], "",
		css.AnsiMap["azure"], "",
		css.AnsiMap["beige"], "",
		css.AnsiMap["bisque"], "",
		css.AnsiMap["black"], "",
		css.AnsiMap["blanchedalmond"], "",
		css.AnsiMap["blue"], "",
		css.AnsiMap["blueviolet"], "",
		css.AnsiMap["brown"], "",
		css.AnsiMap["burlywood"], "",
		css.AnsiMap["cadetblue"], "",
		css.AnsiMap["chartreuse"], "",
		css.AnsiMap["chocolate"], "",
		css.AnsiMap["coral"], "",
		css.AnsiMap["cornflowerblue"], "",
		css.AnsiMap["cornsilk"], "",
		css.AnsiMap["crimson"], "",
		css.AnsiMap["cyan"], "",
		css.AnsiMap["darkblue"], "",
		css.AnsiMap["darkcyan"], "",
		css.AnsiMap["darkgoldenrod"], "",
		css.AnsiMap["darkgray"], "",
		css.AnsiMap["darkgreen"], "",
		css.AnsiMap["darkkhaki"], "",
		css.AnsiMap["darkmagenta"], "",
		css.AnsiMap["darkolivegreen"], "",
		css.AnsiMap["darkorange"], "",
		css.AnsiMap["darkorchid"], "",
		css.AnsiMap["darkred"], "",
		css.AnsiMap["darksalmon"], "",
		css.AnsiMap["darkseagreen"], "",
		css.AnsiMap["darkslateblue"], "",
		css.AnsiMap["darkslategray"], "",
		css.AnsiMap["darkturquoise"], "",
		css.AnsiMap["darkviolet"], "",
		css.AnsiMap["deeppink"], "",
		css.AnsiMap["deepskyblue"], "",
		css.AnsiMap["dimgray"], "",
		css.AnsiMap["dodgerblue"], "",
		css.AnsiMap["firebrick"], "",
		css.AnsiMap["floralwhite"], "",
		css.AnsiMap["forestgreen"], "",
		css.AnsiMap["fuchsia"], "",
		css.AnsiMap["gainsboro"], "",
		css.AnsiMap["ghostwhite"], "",
		css.AnsiMap["gold"], "",
		css.AnsiMap["goldenrod"], "",
		css.AnsiMap["gray"], "",
		css.AnsiMap["green"], "",
		css.AnsiMap["greenyellow"], "",
		css.AnsiMap["honeydew"], "",
		css.AnsiMap["hotpink"], "",
		css.AnsiMap["indianred"], "",
		css.AnsiMap["indigo"], "",
		css.AnsiMap["ivory"], "",
		css.AnsiMap["khaki"], "",
		css.AnsiMap["lavender"], "",
		css.AnsiMap["lavenderblush"], "",
		css.AnsiMap["lawngreen"], "",
		css.AnsiMap["lemonchiffon"], "",
		css.AnsiMap["lightblue"], "",
		css.AnsiMap["lightcoral"], "",
		css.AnsiMap["lightcyan"], "",
		css.AnsiMap["lightgoldenrodyellow"], "",
		css.AnsiMap["lightgreen"], "",
		css.AnsiMap["lightgray"], "",
		css.AnsiMap["lightpink"], "",
		css.AnsiMap["lightsalmon"], "",
		css.AnsiMap["lightseagreen"], "",
		css.AnsiMap["lightskyblue"], "",
		css.AnsiMap["lightslategray"], "",
		css.AnsiMap["lightsteelblue"], "",
		css.AnsiMap["lightyellow"], "",
		css.AnsiMap["lime"], "",
		css.AnsiMap["limegreen"], "",
		css.AnsiMap["linen"], "",
		css.AnsiMap["magenta"], "",
		css.AnsiMap["maroon"], "",
		css.AnsiMap["mediumaquamarine"], "",
		css.AnsiMap["mediumblue"], "",
		css.AnsiMap["mediumorchid"], "",
		css.AnsiMap["mediumpurple"], "",
		css.AnsiMap["mediumseagreen"], "",
		css.AnsiMap["mediumslateblue"], "",
		css.AnsiMap["mediumspringgreen"], "",
		css.AnsiMap["mediumturquoise"], "",
		css.AnsiMap["mediumvioletred"], "",
		css.AnsiMap["midnightblue"], "",
		css.AnsiMap["mintcream"], "",
		css.AnsiMap["mistyrose"], "",
		css.AnsiMap["moccasin"], "",
		css.AnsiMap["navajowhite"], "",
		css.AnsiMap["navy"], "",
		css.AnsiMap["navyblue"], "",
		css.AnsiMap["oldlace"], "",
		css.AnsiMap["olive"], "",
		css.AnsiMap["olivedrab"], "",
		css.AnsiMap["orange"], "",
		css.AnsiMap["orangered"], "",
		css.AnsiMap["orchid"], "",
		css.AnsiMap["palegoldenrod"], "",
		css.AnsiMap["palegreen"], "",
		css.AnsiMap["paleturquoise"], "",
		css.AnsiMap["palevioletred"], "",
		css.AnsiMap["papayawhip"], "",
		css.AnsiMap["peachpuff"], "",
		css.AnsiMap["peru"], "",
		css.AnsiMap["pink"], "",
		css.AnsiMap["plum"], "",
		css.AnsiMap["powderblue"], "",
		css.AnsiMap["purple"], "",
		css.AnsiMap["red"], "",
		css.AnsiMap["rosybrown"], "",
		css.AnsiMap["royalblue"], "",
		css.AnsiMap["saddlebrown"], "",
		css.AnsiMap["salmon"], "",
		css.AnsiMap["sandybrown"], "",
		css.AnsiMap["seagreen"], "",
		css.AnsiMap["seashell"], "",
		css.AnsiMap["sienna"], "",
		css.AnsiMap["silver"], "",
		css.AnsiMap["skyblue"], "",
		css.AnsiMap["slateblue"], "",
		css.AnsiMap["slategray"], "",
		css.AnsiMap["snow"], "",
		css.AnsiMap["springgreen"], "",
		css.AnsiMap["steelblue"], "",
		css.AnsiMap["tan"], "",
		css.AnsiMap["teal"], "",
		css.AnsiMap["thistle"], "",
		css.AnsiMap["tomato"], "",
		css.AnsiMap["turquoise"], "",
		css.AnsiMap["violet"], "",
		css.AnsiMap["wheat"], "",
		css.AnsiMap["white"], "",
		css.AnsiMap["whitesmoke"], "",
		css.AnsiMap["yellow"], "",
		css.AnsiMap["yellowgreen"], "",
	)
}

// DecolorAnsiSol removes any of the sol ansi color escapes known to this package.
func DecolorAnsiSol(s string) string {
	return colorReplacerAnsiSol.Replace(s)
}

// DecolorAnsiiCss removes any of the css ansi color escapes known to this package.
func DecolorAnsiCss(s string) string {
	return colorReplacerAnsiCss.Replace(s)
}

// AnsiRGB creates a new ansi escape code based on an RGB input
func AnsiRGB(rgb RGB) string {
	return "\x1b[38;2;" + strconv.FormatInt(int64(rgb.R), 10) + ";" + strconv.FormatInt(int64(rgb.G), 10) + ";" + strconv.FormatInt(int64(rgb.B), 10) + "m"
}

// AnsiRGBBackground creates a new background ansi escape code based on an RGB input
func AnsiRGBBackground(rgb RGB) Ansi {
	return Ansi("\x1b[48;2;" + strconv.FormatInt(int64(rgb.R), 10) + ";" + strconv.FormatInt(int64(rgb.G), 10) + ";" + strconv.FormatInt(int64(rgb.B), 10) + "m")
}
