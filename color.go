package color

import (
	"strings"

	"gitlab.com/skilstak/go/term/ansi/color/css"
	"gitlab.com/skilstak/go/term/ansi/color/sol"
)

var colorReplacerSol *strings.Replacer
var colorReplacerCss *strings.Replacer

func init() {
	colorReplacerSol = strings.NewReplacer(
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

	colorReplacerCss = strings.NewReplacer(
		css.AliceBlue, "",
		css.AntiqueWhite, "",
		css.Aqua, "",
		css.AquaMarine, "",
		css.Azure, "",
		css.Beige, "",
		css.Bisque, "",
		css.Black, "",
		css.BlanchedAlmond, "",
		css.Blue, "",
		css.BlueViolet, "",
		css.Brown, "",
		css.BurlyWood, "",
		css.CadetBlue, "",
		css.Chartreuse, "",
		css.Chocolate, "",
		css.Coral, "",
		css.CornFlowerBlue, "",
		css.CornSilk, "",
		css.Crimson, "",
		css.Cyan, "",
		css.DarkBlue, "",
		css.DarkCyan, "",
		css.DarkGoldenRod, "",
		css.DarkGray, "",
		css.DarkGreen, "",
		css.DarkKhaki, "",
		css.DarkMagenta, "",
		css.DarkOliveGreen, "",
		css.DarkOrange, "",
		css.DarkOrchid, "",
		css.DarkRed, "",
		css.DarkSalmon, "",
		css.DarkSeaGreen, "",
		css.DarkSlateBlue, "",
		css.DarkSlateGray, "",
		css.DarkTurquoise, "",
		css.DarkViolet, "",
		css.DeepPink, "",
		css.DeepSkyBlue, "",
		css.DimGray, "",
		css.DodgerBlue, "",
		css.Firebrick, "",
		css.FloralWhite, "",
		css.ForestGreen, "",
		css.Fuchsia, "",
		css.Gainsboro, "",
		css.GhostWhite, "",
		css.Gold, "",
		css.GoldenRod, "",
		css.Gray, "",
		css.Green, "",
		css.GreenYellow, "",
		css.Honeydew, "",
		css.HotPink, "",
		css.IndianRed, "",
		css.Indigo, "",
		css.Ivory, "",
		css.Khaki, "",
		css.Lavender, "",
		css.LavenderBlush, "",
		css.LawnGreen, "",
		css.LemonChiffon, "",
		css.LightBlue, "",
		css.LightCoral, "",
		css.LightCyan, "",
		css.LightGoldenRodYellow, "",
		css.LightGreen, "",
		css.LightGrey, "",
		css.LightPink, "",
		css.LightSalmon, "",
		css.LightSeaGreen, "",
		css.LightSkyBlue, "",
		css.LightSlateGray, "",
		css.LightSteelBlue, "",
		css.LightYellow, "",
		css.Lime, "",
		css.LimeGreen, "",
		css.Linen, "",
		css.Magenta, "",
		css.Maroon, "",
		css.MediumAquaMarine, "",
		css.MediumBlue, "",
		css.MediumOrchid, "",
		css.MediumPurple, "",
		css.MediumSeaGreen, "",
		css.MediumSlateBlue, "",
		css.MediumSpringGreen, "",
		css.MediumTurquoise, "",
		css.MediumVioletRed, "",
		css.MidnightBlue, "",
		css.MintCream, "",
		css.MistyRose, "",
		css.Moccasin, "",
		css.NavajoWhite, "",
		css.Navy, "",
		css.NavyBlue, "",
		css.OldLace, "",
		css.Olive, "",
		css.OliveDrab, "",
		css.Orange, "",
		css.OrangeRed, "",
		css.Orchid, "",
		css.PaleGoldenRod, "",
		css.PaleGreen, "",
		css.PaleTurquoise, "",
		css.PaleVioletRed, "",
		css.PapayaWhip, "",
		css.PeachPuff, "",
		css.Peru, "",
		css.Pink, "",
		css.Plum, "",
		css.PowderBlue, "",
		css.Purple, "",
		css.Red, "",
		css.RosyBrown, "",
		css.RoyalBlue, "",
		css.SaddleBrown, "",
		css.Salmon, "",
		css.SandyBrown, "",
		css.SeaGreen, "",
		css.SeaShell, "",
		css.Sienna, "",
		css.Silver, "",
		css.SkyBlue, "",
		css.SlateBlue, "",
		css.SlateGray, "",
		css.Snow, "",
		css.SpringGreen, "",
		css.SteelBlue, "",
		css.Tan, "",
		css.Teal, "",
		css.Thistle, "",
		css.Tomato, "",
		css.Turquoise, "",
		css.Violet, "",
		css.Wheat, "",
		css.White, "",
		css.WhiteSmoke, "",
		css.Yellow, "",
		css.YellowGreen, "",
	)
}

// Decolor removes any of the ANSI color escapes known to this package.
func DecolorSol(s string) string {
	return colorReplacerSol.Replace(s)
}

func SearchGlobalIndex(s string) string {
	if val, ok := Index[strings.ToLower(s)]; ok {
		return val + strings.ToLower(s)
	} else {
		return s + " is not a valid color. To see what colors are valid, run the sample? command."
	}
}

func PrintGlobalIndex() string {
	s := ""
	for k, v := range Index {
		s = s + v + k + "\n"
	}
	return s
}

func Rgb(r string, g string, b string) string {
	s := "\033[38;2;" + r + ";" + g + ";" + b + "m"
	return s
}

func RgbBackground(r string, g string, b string) string {
	s := "\033[48;2;" + r + ";" + g + ";" + b + "m"
	return s
}

var Index = map[string]string{
	"sol.yellow":                         "\033[0;33m",
	"sol.orange":                         "\033[1;31m",
	"sol.red":                            "\033[0;31m",
	"sol.magenta":                        "\033[0;35m",
	"sol.violet":                         "\033[1;35m",
	"sol.blue":                           "\033[0;34m",
	"sol.cyan":                           "\033[0;36m",
	"sol.green":                          "\033[0;32m",
	"css.aliceblue":                      "\033[38;2;240;248;255m",
	"css.antiquewhite":                   "\033[38;2;250;235;215m",
	"css.aqua":                           "\033[38;2;0;255;255m",
	"css.aquamarine":                     "\033[38;2;127;255;212m",
	"css.azure":                          "\033[38;2;1240;255;255m",
	"css.beige":                          "\033[38;2;245;245;220m",
	"css.bisque":                         "\033[38;2;255;228;196m",
	"css.black":                          "\033[38;2;0;0;0m",
	"css.blanchedalmond":                 "\033[38;2;255;235;205m",
	"css.blue":                           "\033[38;2;0;0;255m",
	"css.blueviolet":                     "\033[38;2;138;43;226m",
	"css.brown":                          "\033[38;2;165;42;42m",
	"css.burlywood":                      "\033[38;2;222;184;135m",
	"css.cadetblue":                      "\033[38;2;95;158;160m",
	"css.chartreuse":                     "\033[38;2;95;158;160m",
	"css.chocolate":                      "\033[38;2;210;105;30m",
	"css.coral":                          "\033[38;2;255;127;80m",
	"css.cornflowerblue":                 "\033[38;2;100;149;237m",
	"css.cornsilk":                       "\033[38;2;255;248;220m",
	"css.crimson":                        "\033[38;2;220;20;60m",
	"css.cyan":                           "\033[38;2;0;255;255m",
	"css.darkblue":                       "\033[38;2;0;0;139m",
	"css.darkcyan":                       "\033[38;2;0;139;139m",
	"css.darkgoldenrod":                  "\033[38;2;184;134;11m",
	"css.darkgray":                       "\033[38;2;169;169;169m",
	"css.darkgreen":                      "\033[38;2;0;100;0m",
	"css.darkkhaki":                      "\033[38;2;189;183;107m",
	"css.darkmagenta":                    "\033[38;2;139;0;139m",
	"css.darkolivegreen":                 "\033[38;2;85;107;47m",
	"css.darkorange":                     "\033[38;2;255;140;0m",
	"css.darkorchid":                     "\033[38;2;153;50;204m",
	"css.darkred":                        "\033[38;2;139;0;0m",
	"css.darksalmon":                     "\033[38;2;233;150;122m",
	"css.darkseagreen":                   "\033[38;2;143;188;143m",
	"css.darkslateblue":                  "\033[38;2;72;61;139m",
	"css.darkslategray":                  "\033[38;2;47;79;79m",
	"css.darkturquoise":                  "\033[38;2;0;206;209m",
	"css.darkviolet":                     "\033[38;2;148;0;211m",
	"css.deeppink":                       "\033[38;2;255;20;147m",
	"css.deepskyblue":                    "\033[38;2;0;191;255m",
	"css.dimgray":                        "\033[38;2;0;191;255m",
	"css.dodgerblue":                     "\033[38;2;30;144;255m",
	"css.firebrick":                      "\033[38;2;178;34;34m",
	"css.floralwhite":                    "\033[38;2;255;250;240m",
	"css.forestgreen":                    "\033[38;2;34;139;34m",
	"css.fuchsia":                        "\033[38;2;255;0;255m",
	"css.gainsboro":                      "\033[38;2;220;220;220m",
	"css.ghostwhite":                     "\033[38;2;248;248;255m",
	"css.gold":                           "\033[38;2;255;215;0m",
	"css.goldenrod":                      "\033[38;2;218;165;32m",
	"css.gray":                           "\033[38;2;127;127;127m",
	"css.green":                          "\033[38;2;0;128;0m",
	"css.greenyellow":                    "\033[38;2;173;255;47m",
	"css.honeydew":                       "\033[38;2;240;255;240m",
	"css.hotpink":                        "\033[38;2;255;105;180m",
	"css.indianred":                      "\033[38;2;205;92;92m",
	"css.indigo":                         "\033[38;2;75;0;130m",
	"css.ivory":                          "\033[38;2;255;255;240m",
	"css.khaki":                          "\033[38;2;240;230;140m",
	"css.lavender":                       "\033[38;2;230;230;250m",
	"css.lavenderblush":                  "\033[38;2;255;240;245m",
	"css.lawngreen":                      "\033[38;2;124;252;0m",
	"css.lemonchiffon":                   "\033[38;2;255;250;205m",
	"css.lightblue":                      "\033[38;2;173;216;230m",
	"css.lightcoral":                     "\033[38;2;240;128;128m",
	"css.lightcyan":                      "\033[38;2;224;255;255m",
	"css.lightgoldenrodyellow":           "\033[38;2;250;250;210m",
	"css.lightgreen":                     "\033[38;2;144;238;144m",
	"css.lightgrey":                      "\033[38;2;211;211;211m",
	"css.lightpink":                      "\033[38;2;255;182;193m",
	"css.lightsalmon":                    "\033[38;2;255;160;122m",
	"css.lightseagreen":                  "\033[38;2;32;178;170m",
	"css.lightskyblue":                   "\033[38;2;135;206;250m",
	"css.lightslategray":                 "\033[38;2;119;136;153m",
	"css.lightsteelblue":                 "\033[38;2;176;196;222m",
	"css.lightyellow":                    "\033[38;2;255;255;224m",
	"css.lime":                           "\033[38;2;0;255;0m",
	"css.limegreen":                      "\033[38;2;50;205;50m",
	"css.linen":                          "\033[38;2;250;240;230m",
	"css.magenta":                        "\033[38;2;255;0;255m",
	"css.maroon":                         "\033[38;2;128;0;0m",
	"css.mediumaquamarine":               "\033[38;2;102;205;170m",
	"css.mediumblue":                     "\033[38;2;0;0;205m",
	"css.mediumorchid":                   "\033[38;2;186;85;211m",
	"css.mediumpurple":                   "\033[38;2;147;112;219m",
	"css.mediumseagreen":                 "\033[38;2;60;179;113m",
	"css.mediumslateblue":                "\033[38;2;123;104;238m",
	"css.mediumspringgreen":              "\033[38;2;0;250;154m",
	"css.mediumturquoise":                "\033[38;2;72;209;204m",
	"css.mediumvioletred":                "\033[38;2;199;21;133m",
	"css.midnightblue":                   "\033[38;2;25;25;112m",
	"css.mintcream":                      "\033[38;2;245;255;250m",
	"css.mistyrose":                      "\033[38;2;255;228;225m",
	"css.moccasin":                       "\033[38;2;255;228;181m",
	"css.navajowhite":                    "\033[38;2;255;222;173m",
	"css.navy":                           "\033[38;2;0;0;128m",
	"css.navyblue":                       "\033[38;2;159;175;223m",
	"css.oldlace":                        "\033[38;2;253;245;230m",
	"css.olive":                          "\033[38;2;128;128;0m",
	"css.olivedrab":                      "\033[38;2;107;142;35m",
	"css.orange":                         "\033[38;2;255;165;0m",
	"css.orangered":                      "\033[38;2;255;69;0m",
	"css.orchid":                         "\033[38;2;218;112;214m",
	"css.palegoldenrod":                  "\033[38;2;238;232;170m",
	"css.palegreen":                      "\033[38;2;152;251;152m",
	"css.paleturquoise":                  "\033[38;2;175;238;238m",
	"css.palevioletred":                  "\033[38;2;219;112;147m",
	"css.papayawhip":                     "\033[38;2;255;239;213m",
	"css.peachpuff":                      "\033[38;2;255;218;185m",
	"css.peru":                           "\033[38;2;205;133;63m",
	"css.pink":                           "\033[38;2;255;192;203m",
	"css.plum":                           "\033[38;2;221;160;221m",
	"css.powderblue":                     "\033[38;2;176;224;230m",
	"css.purple":                         "\033[38;2;128;0;128m",
	"css.red":                            "\033[38;2;255;0;0m",
	"css.rosybrown":                      "\033[38;2;188;143;143m",
	"css.royalblue":                      "\033[38;2;65;105;225m",
	"css.saddlebrown":                    "\033[38;2;139;69;19m",
	"css.salmon":                         "\033[38;2;250;128;114m",
	"css.sandybrown":                     "\033[38;2;244;164;96m",
	"css.seagreen":                       "\033[38;2;46;139;87m",
	"css.seashell":                       "\033[38;2;255;245;238m",
	"css.sienna":                         "\033[38;2;160;82;45m",
	"css.silver":                         "\033[38;2;192;192;192m",
	"css.skyblue":                        "\033[38;2;135;206;235m",
	"css.slateblue":                      "\033[38;2;106;90;205m",
	"css.slategray":                      "\033[38;2;112;128;144m",
	"css.snow":                           "\033[38;2;255;250;250m",
	"css.springgreen":                    "\033[38;2;0;255;127m",
	"css.steelblue":                      "\033[38;2;70;130;180m",
	"css.tan":                            "\033[38;2;210;180;140m",
	"css.teal":                           "\033[38;2;0;128;128m",
	"css.thistle":                        "\033[38;2;216;191;216m",
	"css.tomato":                         "\033[38;2;255;99;71m",
	"css.turquoise":                      "\033[38;2;64;224;208m",
	"css.violet":                         "\033[38;2;238;130;238m",
	"css.wheat":                          "\033[38;2;245;222;179m",
	"css.white":                          "\033[38;2;255;255;255m",
	"css.whitesmoke":                     "\033[38;2;245;245;245m",
	"css.yellow":                         "\033[38;2;255;255;0m",
	"css.yellowgreen":                    "\033[38;2;139;205;50m",
	"css.alicebluebackground":            "\033[48;2;240;248;255m",
	"css.antiquewhitebackground":         "\033[48;2;250;235;215m",
	"css.aquabackground":                 "\033[48;2;0;255;255m",
	"css.aquamarinebackground":           "\033[48;2;127;255;212m",
	"css.azurebackground":                "\033[48;2;1240;255;255m",
	"css.beigebackground":                "\033[48;2;245;245;220m",
	"css.bisquebackground":               "\033[48;2;255;228;196m",
	"css.blackbackground":                "\033[48;2;0;0;0m",
	"css.blanchedalmondbackground":       "\033[48;2;255;235;205m",
	"css.bluebackground":                 "\033[48;2;0;0;255m",
	"css.bluevioletbackground":           "\033[48;2;138;43;226m",
	"css.brownbackground":                "\033[48;2;165;42;42m",
	"css.burlywoodbackground":            "\033[48;2;222;184;135m",
	"css.cadetbluebackground":            "\033[48;2;95;158;160m",
	"css.chartreusebackground":           "\033[48;2;95;158;160m",
	"css.chocolatebackground":            "\033[48;2;210;105;30m",
	"css.coralbackground":                "\033[48;2;255;127;80m",
	"css.cornflowerbluebackground":       "\033[48;2;100;149;237m",
	"css.cornsilkbackground":             "\033[48;2;255;248;220m",
	"css.crimsonbackground":              "\033[48;2;220;20;60m",
	"css.cyanbackground":                 "\033[48;2;0;255;255m",
	"css.darkbluebackground":             "\033[48;2;0;0;139m",
	"css.darkcyanbackground":             "\033[48;2;0;139;139m",
	"css.darkgoldenrodbackground":        "\033[48;2;184;134;11m",
	"css.darkgraybackground":             "\033[48;2;169;169;169m",
	"css.darkgreenbackground":            "\033[48;2;0;100;0m",
	"css.darkkhakibackground":            "\033[48;2;189;183;107m",
	"css.darkmagentabackground":          "\033[48;2;139;0;139m",
	"css.darkolivegreenbackground":       "\033[48;2;85;107;47m",
	"css.darkorangebackground":           "\033[48;2;255;140;0m",
	"css.darkorchidbackground":           "\033[48;2;153;50;204m",
	"css.darkredbackground":              "\033[48;2;139;0;0m",
	"css.darksalmonbackground":           "\033[48;2;233;150;122m",
	"css.darkseagreenbackground":         "\033[48;2;143;188;143m",
	"css.darkslatebluebackground":        "\033[48;2;72;61;139m",
	"css.darkslategraybackground":        "\033[48;2;47;79;79m",
	"css.darkturquoisebackground":        "\033[48;2;0;206;209m",
	"css.darkvioletbackground":           "\033[48;2;148;0;211m",
	"css.deeppinkbackground":             "\033[48;2;255;20;147m",
	"css.deepskybluebackground":          "\033[48;2;0;191;255m",
	"css.dimgraybackground":              "\033[48;2;0;191;255m",
	"css.dodgerbluebackground":           "\033[48;2;30;144;255m",
	"css.firebrickbackground":            "\033[48;2;178;34;34m",
	"css.floralwhitebackground":          "\033[48;2;255;250;240m",
	"css.forestgreenbackground":          "\033[48;2;34;139;34m",
	"css.fuchsiabackground":              "\033[48;2;255;0;255m",
	"css.gainsborobackground":            "\033[48;2;220;220;220m",
	"css.ghostwhitebackground":           "\033[48;2;248;248;255m",
	"css.goldbackground":                 "\033[48;2;255;215;0m",
	"css.goldenrodbackground":            "\033[48;2;218;165;32m",
	"css.graybackground":                 "\033[48;2;127;127;127m",
	"css.greenbackground":                "\033[48;2;0;128;0m",
	"css.greenyellowbackground":          "\033[48;2;173;255;47m",
	"css.honeydewbackground":             "\033[48;2;240;255;240m",
	"css.hotpinkbackground":              "\033[48;2;255;105;180m",
	"css.indianredbackground":            "\033[48;2;205;92;92m",
	"css.indigobackground":               "\033[48;2;75;0;130m",
	"css.ivorybackground":                "\033[48;2;255;255;240m",
	"css.khakibackground":                "\033[48;2;240;230;140m",
	"css.lavenderbackground":             "\033[48;2;230;230;250m",
	"css.lavenderblushbackground":        "\033[48;2;255;240;245m",
	"css.lawngreenbackground":            "\033[48;2;124;252;0m",
	"css.lemonchiffonbackground":         "\033[48;2;255;250;205m",
	"css.lightbluebackground":            "\033[48;2;173;216;230m",
	"css.lightcoralbackground":           "\033[48;2;240;128;128m",
	"css.lightcyanbackground":            "\033[48;2;224;255;255m",
	"css.lightgoldenrodyellowbackground": "\033[48;2;250;250;210m",
	"css.lightgreenbackground":           "\033[48;2;144;238;144m",
	"css.lightgreybackground":            "\033[48;2;211;211;211m",
	"css.lightpinkbackground":            "\033[48;2;255;182;193m",
	"css.lightsalmonbackground":          "\033[48;2;255;160;122m",
	"css.lightseagreenbackground":        "\033[48;2;32;178;170m",
	"css.lightskybluebackground":         "\033[48;2;135;206;250m",
	"css.lightslategraybackground":       "\033[48;2;119;136;153m",
	"css.lightsteelbluebackground":       "\033[48;2;176;196;222m",
	"css.lightyellowbackground":          "\033[48;2;255;255;224m",
	"css.limebackground":                 "\033[48;2;0;255;0m",
	"css.limegreenbackground":            "\033[48;2;50;205;50m",
	"css.linenbackground":                "\033[48;2;250;240;230m",
	"css.magentabackground":              "\033[48;2;255;0;255m",
	"css.maroonbackground":               "\033[48;2;128;0;0m",
	"css.mediumaquamarinebackground":     "\033[48;2;102;205;170m",
	"css.mediumbluebackground":           "\033[48;2;0;0;205m",
	"css.mediumorchidbackground":         "\033[48;2;186;85;211m",
	"css.mediumpurplebackground":         "\033[48;2;147;112;219m",
	"css.mediumseagreenbackground":       "\033[48;2;60;179;113m",
	"css.mediumslatebluebackground":      "\033[48;2;123;104;238m",
	"css.mediumspringgreenbackground":    "\033[48;2;0;250;154m",
	"css.mediumturquoisebackground":      "\033[48;2;72;209;204m",
	"css.mediumvioletredbackground":      "\033[48;2;199;21;133m",
	"css.midnightbluebackground":         "\033[48;2;25;25;112m",
	"css.mintcreambackground":            "\033[48;2;245;255;250m",
	"css.mistyrosebackground":            "\033[48;2;255;228;225m",
	"css.moccasinbackground":             "\033[48;2;255;228;181m",
	"css.navajowhitebackground":          "\033[48;2;255;222;173m",
	"css.navybackground":                 "\033[48;2;0;0;128m",
	"css.navybluebackground":             "\033[48;2;159;175;223m",
	"css.oldlacebackground":              "\033[48;2;253;245;230m",
	"css.olivebackground":                "\033[48;2;128;128;0m",
	"css.olivedrabbackground":            "\033[48;2;107;142;35m",
	"css.orangebackground":               "\033[48;2;255;165;0m",
	"css.orangeredbackground":            "\033[48;2;255;69;0m",
	"css.orchidbackground":               "\033[48;2;218;112;214m",
	"css.palegoldenrodbackground":        "\033[48;2;238;232;170m",
	"css.palegreenbackground":            "\033[48;2;152;251;152m",
	"css.paleturquoisebackground":        "\033[48;2;175;238;238m",
	"css.palevioletredbackground":        "\033[48;2;219;112;147m",
	"css.papayawhipbackground":           "\033[48;2;255;239;213m",
	"css.peachpuffbackground":            "\033[48;2;255;218;185m",
	"css.perubackground":                 "\033[48;2;205;133;63m",
	"css.pinkbackground":                 "\033[48;2;255;192;203m",
	"css.plumbackground":                 "\033[48;2;221;160;221m",
	"css.powderbluebackground":           "\033[48;2;176;224;230m",
	"css.purplebackground":               "\033[48;2;128;0;128m",
	"css.redbackground":                  "\033[48;2;255;0;0m",
	"css.rosybrownbackground":            "\033[48;2;188;143;143m",
	"css.royalbluebackground":            "\033[48;2;65;105;225m",
	"css.saddlebrownbackground":          "\033[48;2;139;69;19m",
	"css.salmonbackground":               "\033[48;2;250;128;114m",
	"css.sandybrownbackground":           "\033[48;2;244;164;96m",
	"css.seagreenbackground":             "\033[48;2;46;139;87m",
	"css.seashellbackground":             "\033[48;2;255;245;238m",
	"css.siennabackground":               "\033[48;2;160;82;45m",
	"css.silverbackground":               "\033[48;2;192;192;192m",
	"css.skybluebackground":              "\033[48;2;135;206;235m",
	"css.slatebluebackground":            "\033[48;2;106;90;205m",
	"css.slategraybackground":            "\033[48;2;112;128;144m",
	"css.snowbackground":                 "\033[48;2;255;250;250m",
	"css.springgreenbackground":          "\033[48;2;0;255;127m",
	"css.steelbluebackground":            "\033[48;2;70;130;180m",
	"css.tanbackground":                  "\033[48;2;210;180;140m",
	"css.tealbackground":                 "\033[48;2;0;128;128m",
	"css.thistlebackground":              "\033[48;2;216;191;216m",
	"css.tomatobackground":               "\033[48;2;255;99;71m",
	"css.turquoisebackground":            "\033[48;2;64;224;208m",
	"css.violetbackground":               "\033[48;2;238;130;238m",
	"css.wheatbackground":                "\033[48;2;245;222;179m",
	"css.whitebackground":                "\033[48;2;255;255;255m",
	"css.whitesmokebackground":           "\033[48;2;245;245;245m",
	"css.yellowbackground":               "\033[48;2;255;255;0m",
	"css.yellowgreenbackground":          "\033[48;2;139;205;50m",
}
