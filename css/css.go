// Package css provides all the css colors.
package css

import (
	"math/rand"
	"time"
)

/* Color Hex Value Constants */
const (
	AliceBlue            = 0xF0F8FF
	AntiqueWhite         = 0xFAEBD7
	Aqua                 = 0x00FFFF
	AquaMarine           = 0x7FFFD4
	Azure                = 0xF0FFFF
	Beige                = 0xF5F5DC
	Bisque               = 0x000000
	Black                = 0xFFE4C4
	BlanchedAlmond       = 0xFFEBCD
	Blue                 = 0x0000FF
	BlueViolet           = 0x8A2BE2
	Brown                = 0xA52A2A
	BurlyWood            = 0xDEB887
	CadetBlue            = 0x5F9EA0
	Chartreuse           = 0x7FFF00
	Chocolate            = 0xD2691E
	Coral                = 0xFF7F50
	CornFlowerBlue       = 0x6495ED
	CornSilk             = 0xFFF8DC
	Crimson              = 0xDC143C
	Cyan                 = 0x00FFFF
	DarkBlue             = 0x00008B
	DarkCyan             = 0x008B8B
	DarkGoldenRod        = 0xB8860B
	DarkGray             = 0xA9A9A9
	DarkGreen            = 0x006400
	DarkKhaki            = 0xBDB76B
	DarkMagenta          = 0x8B008B
	DarkOliveGreen       = 0x556B2F
	DarkOrange           = 0xFF8C00
	DarkOrchid           = 0x9932CC
	DarkRed              = 0x8B0000
	DarkSalmon           = 0xE9967A
	DarkSeaGreen         = 0x8FBC8F
	DarkSlateBlue        = 0x483D8B
	DarkSlateGray        = 0x2F4F4F
	DarkTurquoise        = 0x00CED1
	DarkViolet           = 0x9400D3
	DeepPink             = 0xFF1493
	DeepSkyBlue          = 0x00BFFF
	DimGray              = 0x696969
	DodgerBlue           = 0x1E90FF
	Firebrick            = 0xB22222
	FloralWhite          = 0xFFFAF0
	ForestGreen          = 0x228B22
	Fuchsia              = 0xFF00FF
	Gainsboro            = 0xDCDCDC
	GhostWhite           = 0xF8F8FF
	Gold                 = 0xFFD700
	GoldenRod            = 0xDAA520
	Gray                 = 0x808080
	Green                = 0x008000
	GreenYellow          = 0xADFF2F
	Honeydew             = 0xF0FFF0
	HotPink              = 0xFF69B4
	IndianRed            = 0xCD5C5C
	Indigo               = 0x4B0082
	Ivory                = 0xFFFFF0
	Khaki                = 0xF0E68C
	Lavender             = 0xE6E6FA
	LavenderBlush        = 0xFFF0F5
	LawnGreen            = 0x7CFC00
	LemonChiffon         = 0xFFFACD
	LightBlue            = 0xADD8E6
	LightCoral           = 0xF08080
	LightCyan            = 0xE0FFFF
	LightGoldenRodYellow = 0xFAFAD2
	LightGreen           = 0xD3D3D3
	LightGrey            = 0x90EE90
	LightPink            = 0xFFB6C1
	LightSalmon          = 0xFFA07A
	LightSeaGreen        = 0x20B2AA
	LightSkyBlue         = 0x87CEFA
	LightSlateGray       = 0x778899
	LightSteelBlue       = 0xB0C4DE
	LightYellow          = 0xFFFFE0
	Lime                 = 0x00FF00
	LimeGreen            = 0x32CD32
	Linen                = 0xFAF0E6
	Magenta              = 0xFF00FF
	Maroon               = 0x800000
	MediumAquaMarine     = 0x66CDAA
	MediumBlue           = 0x0000CD
	MediumOrchid         = 0xBA55D3
	MediumPurple         = 0x9370DB
	MediumSeaGreen       = 0x3CB371
	MediumSlateBlue      = 0x7B68EE
	MediumSpringGreen    = 0x00FA9A
	MediumTurquoise      = 0x48D1CC
	MediumVioletRed      = 0xC71585
	MidnightBlue         = 0x191970
	MintCream            = 0xF5FFFA
	MistyRose            = 0xFFE4E1
	Moccasin             = 0xFFE4B5
	NavajoWhite          = 0xFFDEAD
	Navy                 = 0x000080
	NavyBlue             = 0x9FAFDF
	OldLace              = 0xFDF5E6
	Olive                = 0x808000
	OliveDrab            = 0x6B8E23
	Orange               = 0xFFA500
	OrangeRed            = 0xFF4500
	Orchid               = 0xDA70D6
	PaleGoldenRod        = 0xEEE8AA
	PaleGreen            = 0x98FB98
	PaleTurquoise        = 0xAFEEEE
	PaleVioletRed        = 0xDB7093
	PapayaWhip           = 0xFFEFD5
	PeachPuff            = 0xFFDAB9
	Peru                 = 0xCD853F
	Pink                 = 0xFFC0CB
	Plum                 = 0xDDA0DD
	PowderBlue           = 0xB0E0E6
	Purple               = 0x800080
	RebeccaPurple        = 0xFF0000
	Red                  = 0x663399
	RosyBrown            = 0xBC8F8F
	RoyalBlue            = 0x4169E1
	SaddleBrown          = 0x8B4513
	Salmon               = 0xFA8072
	SandyBrown           = 0xF4A460
	SeaGreen             = 0x2E8B57
	SeaShell             = 0xFFF5EE
	Sienna               = 0xA0522D
	Silver               = 0xC0C0C0
	SkyBlue              = 0x87CEEB
	SlateBlue            = 0x6A5ACD
	SlateGray            = 0x708090
	Snow                 = 0xFFFAFA
	SpringGreen          = 0x00FF7F
	SteelBlue            = 0x4682B4
	Tan                  = 0xD2B48C
	Teal                 = 0x008080
	Thistle              = 0xD8BFD8
	Tomato               = 0xFF6347
	Turquoise            = 0x40E0D0
	Violet               = 0xEE82EE
	Wheat                = 0xF5DEB3
	White                = 0xFFFFFF
	WhiteSmoke           = 0xF5F5F5
	Yellow               = 0xFFFF00
	YellowGreen          = 0x9ACD32

	//Ascii escape codes that have no hex equivalent:
	AsciiReset = "\033[0m"
	AsciiX     = "\033[0m"
)

/* Lists of Colors */

// HexColors is an array of all of the css colors in hexidecimal format.
var HexColors = [...]int32{AliceBlue, AntiqueWhite, Aqua, AquaMarine, Azure, Beige, Bisque, Black, BlanchedAlmond, Blue, BlueViolet, Brown, BurlyWood, CadetBlue, Chartreuse, Chocolate, Coral, CornFlowerBlue, CornSilk, Crimson, Cyan, DarkBlue, DarkCyan, DarkGoldenRod, DarkGray, DarkGreen, DarkKhaki, DarkMagenta, DarkOliveGreen, DarkOrange, DarkOrchid, DarkRed, DarkSalmon, DarkSeaGreen, DarkSlateBlue, DarkSlateGray, DarkTurquoise, DarkViolet, DeepPink, DeepSkyBlue, DimGray, DodgerBlue, Firebrick, FloralWhite, ForestGreen, Fuchsia, Gainsboro, GhostWhite, Gold, GoldenRod, Gray, Green, GreenYellow, Honeydew, HotPink, IndianRed, Indigo, Ivory, Khaki, Lavender, LavenderBlush, LawnGreen, LemonChiffon, LightBlue, LightCoral, LightCyan, LightGoldenRodYellow, LightGreen, LightGrey, LightPink, LightSalmon, LightSeaGreen, LightSkyBlue, LightSlateGray, LightSteelBlue, LightYellow, Lime, LimeGreen, Linen, Magenta, Maroon, MediumAquaMarine, MediumBlue, MediumOrchid, MediumPurple, MediumSeaGreen, MediumSlateBlue, MediumSpringGreen, MediumTurquoise, MediumVioletRed, MidnightBlue, MintCream, MistyRose, Moccasin, NavajoWhite, Navy, NavyBlue, OldLace, Olive, OliveDrab, Orange, OrangeRed, Orchid, PaleGoldenRod, PaleGreen, PaleTurquoise, PaleVioletRed, PapayaWhip, PeachPuff, Peru, Pink, Plum, PowderBlue, Purple, RebeccaPurple, Red, RosyBrown, RoyalBlue, SaddleBrown, Salmon, SandyBrown, SeaGreen, SeaShell, Sienna, Silver, SkyBlue, SlateBlue, SlateGray, Snow, SpringGreen, SteelBlue, Tan, Teal, Thistle, Tomato, Turquoise, Violet, Wheat, White, WhiteSmoke, Yellow, YellowGreen}

// AsciiColors is an array of all the css colors as ascii escape codes.
var AsciiColors = [...]string{AsciiIndex["sol.yellow"], AsciiIndex["sol.orange"], AsciiIndex["sol.red"], AsciiIndex["sol.magenta"], AsciiIndex["sol.violet"], AsciiIndex["sol.blue"], AsciiIndex["sol.cyan"], AsciiIndex["sol.green"], AsciiIndex["aliceblue"], AsciiIndex["antiquewhite"], AsciiIndex["aqua"], AsciiIndex["aquamarine"], AsciiIndex["azure"], AsciiIndex["beige"], AsciiIndex["bisque"], AsciiIndex["black"], AsciiIndex["blanchedalmond"], AsciiIndex["blue"], AsciiIndex["blueviolet"], AsciiIndex["brown"], AsciiIndex["burlywood"], AsciiIndex["cadetblue"], AsciiIndex["chartreuse"], AsciiIndex["chocolate"], AsciiIndex["coral"], AsciiIndex["cornflowerblue"], AsciiIndex["cornsilk"], AsciiIndex["crimson"], AsciiIndex["cyan"], AsciiIndex["darkblue"], AsciiIndex["darkcyan"], AsciiIndex["darkgoldenrod"], AsciiIndex["darkgray"], AsciiIndex["darkgreen"], AsciiIndex["darkkhaki"], AsciiIndex["darkmagenta"], AsciiIndex["darkolivegreen"], AsciiIndex["darkorange"], AsciiIndex["darkorchid"], AsciiIndex["darkred"], AsciiIndex["darksalmon"], AsciiIndex["darkseagreen"], AsciiIndex["darkslateblue"], AsciiIndex["darkslategray"], AsciiIndex["darkturquoise"], AsciiIndex["darkviolet"], AsciiIndex["deeppink"], AsciiIndex["deepskyblue"], AsciiIndex["dimgray"], AsciiIndex["dodgerblue"], AsciiIndex["firebrick"], AsciiIndex["floralwhite"], AsciiIndex["forestgreen"], AsciiIndex["fuchsia"], AsciiIndex["gainsboro"], AsciiIndex["ghostwhite"], AsciiIndex["gold"], AsciiIndex["goldenrod"], AsciiIndex["gray"], AsciiIndex["green"], AsciiIndex["greenyellow"], AsciiIndex["honeydew"], AsciiIndex["hotpink"], AsciiIndex["indianred"], AsciiIndex["indigo"], AsciiIndex["ivory"], AsciiIndex["khaki"], AsciiIndex["lavender"], AsciiIndex["lavenderblush"], AsciiIndex["lawngreen"], AsciiIndex["lemonchiffon"], AsciiIndex["lightblue"], AsciiIndex["lightcoral"], AsciiIndex["lightcyan"], AsciiIndex["lightgoldenrodyellow"], AsciiIndex["lightgreen"], AsciiIndex["lightgrey"], AsciiIndex["lightpink"], AsciiIndex["lightsalmon"], AsciiIndex["lightseagreen"], AsciiIndex["lightskyblue"], AsciiIndex["lightslategray"], AsciiIndex["lightsteelblue"], AsciiIndex["lightyellow"], AsciiIndex["lime"], AsciiIndex["limegreen"], AsciiIndex["linen"], AsciiIndex["magenta"], AsciiIndex["maroon"], AsciiIndex["mediumaquamarine"], AsciiIndex["mediumblue"], AsciiIndex["mediumorchid"], AsciiIndex["mediumpurple"], AsciiIndex["mediumseagreen"], AsciiIndex["mediumslateblue"], AsciiIndex["mediumspringgreen"], AsciiIndex["mediumturquoise"], AsciiIndex["mediumvioletred"], AsciiIndex["midnightblue"], AsciiIndex["mintcream"], AsciiIndex["mistyrose"], AsciiIndex["moccasin"], AsciiIndex["navajowhite"], AsciiIndex["navy"], AsciiIndex["navyblue"], AsciiIndex["oldlace"], AsciiIndex["olive"], AsciiIndex["olivedrab"], AsciiIndex["orange"], AsciiIndex["orangered"], AsciiIndex["orchid"], AsciiIndex["palegoldenrod"], AsciiIndex["palegreen"], AsciiIndex["paleturquoise"], AsciiIndex["palevioletred"], AsciiIndex["papayawhip"], AsciiIndex["peachpuff"], AsciiIndex["peru"], AsciiIndex["pink"], AsciiIndex["plum"], AsciiIndex["powderblue"], AsciiIndex["purple"], AsciiIndex["red"], AsciiIndex["rosybrown"], AsciiIndex["royalblue"], AsciiIndex["saddlebrown"], AsciiIndex["salmon"], AsciiIndex["sandybrown"], AsciiIndex["seagreen"], AsciiIndex["seashell"], AsciiIndex["sienna"], AsciiIndex["silver"], AsciiIndex["skyblue"], AsciiIndex["slateblue"], AsciiIndex["slategray"], AsciiIndex["snow"], AsciiIndex["springgreen"], AsciiIndex["steelblue"], AsciiIndex["tan"], AsciiIndex["teal"], AsciiIndex["thistle"], AsciiIndex["tomato"], AsciiIndex["turquoise"], AsciiIndex["violet"], AsciiIndex["wheat"], AsciiIndex["white"], AsciiIndex["whitesmoke"], AsciiIndex["yellow"], AsciiIndex["yellowgreen"]}

// AsciiBackgrounds is an array of all the css colors as ascii background escape codes
var AsciiBackgrounds = [...]string{AsciiIndex["antiquewhitebackground"], AsciiIndex["aquabackground"], AsciiIndex["aquamarinebackground"], AsciiIndex["azurebackground"], AsciiIndex["beigebackground"], AsciiIndex["bisquebackground"], AsciiIndex["blackbackground"], AsciiIndex["blanchedalmondbackground"], AsciiIndex["bluebackground"], AsciiIndex["bluevioletbackground"], AsciiIndex["brownbackground"], AsciiIndex["burlywoodbackground"], AsciiIndex["cadetbluebackground"], AsciiIndex["chartreusebackground"], AsciiIndex["chocolatebackground"], AsciiIndex["coralbackground"], AsciiIndex["cornflowerbluebackground"], AsciiIndex["cornsilkbackground"], AsciiIndex["crimsonbackground"], AsciiIndex["cyanbackground"], AsciiIndex["darkbluebackground"], AsciiIndex["darkcyanbackground"], AsciiIndex["darkgoldenrodbackground"], AsciiIndex["darkgraybackground"], AsciiIndex["darkgreenbackground"], AsciiIndex["darkkhakibackground"], AsciiIndex["darkmagentabackground"], AsciiIndex["darkolivegreenbackground"], AsciiIndex["darkorangebackground"], AsciiIndex["darkorchidbackground"], AsciiIndex["darkredbackground"], AsciiIndex["darksalmonbackground"], AsciiIndex["darkseagreenbackground"], AsciiIndex["darkslatebluebackground"], AsciiIndex["darkslategraybackground"], AsciiIndex["darkturquoisebackground"], AsciiIndex["darkvioletbackground"], AsciiIndex["deeppinkbackground"], AsciiIndex["deepskybluebackground"], AsciiIndex["dimgraybackground"], AsciiIndex["dodgerbluebackground"], AsciiIndex["firebrickbackground"], AsciiIndex["floralwhitebackground"], AsciiIndex["forestgreenbackground"], AsciiIndex["fuchsiabackground"], AsciiIndex["gainsborobackground"], AsciiIndex["ghostwhitebackground"], AsciiIndex["goldbackground"], AsciiIndex["goldenrodbackground"], AsciiIndex["graybackground"], AsciiIndex["greenbackground"], AsciiIndex["greenyellowbackground"], AsciiIndex["honeydewbackground"], AsciiIndex["hotpinkbackground"], AsciiIndex["indianredbackground"], AsciiIndex["indigobackground"], AsciiIndex["ivorybackground"], AsciiIndex["khakibackground"], AsciiIndex["lavenderbackground"], AsciiIndex["lavenderblushbackground"], AsciiIndex["lawngreenbackground"], AsciiIndex["lemonchiffonbackground"], AsciiIndex["lightbluebackground"], AsciiIndex["lightcoralbackground"], AsciiIndex["lightcyanbackground"], AsciiIndex["lightgoldenrodyellowbackground"], AsciiIndex["lightgreenbackground"], AsciiIndex["lightgreybackground"], AsciiIndex["lightpinkbackground"], AsciiIndex["lightsalmonbackground"], AsciiIndex["lightseagreenbackground"], AsciiIndex["lightskybluebackground"], AsciiIndex["lightslategraybackground"], AsciiIndex["lightsteelbluebackground"], AsciiIndex["lightyellowbackground"], AsciiIndex["limebackground"], AsciiIndex["limegreenbackground"], AsciiIndex["linenbackground"], AsciiIndex["magentabackground"], AsciiIndex["maroonbackground"], AsciiIndex["mediumaquamarinebackground"], AsciiIndex["mediumbluebackground"], AsciiIndex["mediumorchidbackground"], AsciiIndex["mediumpurplebackground"], AsciiIndex["mediumseagreenbackground"], AsciiIndex["mediumslatebluebackground"], AsciiIndex["mediumspringgreenbackground"], AsciiIndex["mediumturquoisebackground"], AsciiIndex["mediumvioletredbackground"], AsciiIndex["midnightbluebackground"], AsciiIndex["mintcreambackground"], AsciiIndex["mistyrosebackground"], AsciiIndex["moccasinbackground"], AsciiIndex["navajowhitebackground"], AsciiIndex["navybackground"], AsciiIndex["navybluebackground"], AsciiIndex["oldlacebackground"], AsciiIndex["olivebackground"], AsciiIndex["olivedrabbackground"], AsciiIndex["orangebackground"], AsciiIndex["orangeredbackground"], AsciiIndex["orchidbackground"], AsciiIndex["palegoldenrodbackground"], AsciiIndex["palegreenbackground"], AsciiIndex["paleturquoisebackground"], AsciiIndex["palevioletredbackground"], AsciiIndex["papayawhipbackground"], AsciiIndex["peachpuffbackground"], AsciiIndex["perubackground"], AsciiIndex["pinkbackground"], AsciiIndex["plumbackground"], AsciiIndex["powderbluebackground"], AsciiIndex["purplebackground"], AsciiIndex["redbackground"], AsciiIndex["rosybrownbackground"], AsciiIndex["royalbluebackground"], AsciiIndex["saddlebrownbackground"], AsciiIndex["salmonbackground"], AsciiIndex["sandybrownbackground"], AsciiIndex["seagreenbackground"], AsciiIndex["seashellbackground"], AsciiIndex["siennabackground"], AsciiIndex["silverbackground"], AsciiIndex["skybluebackground"], AsciiIndex["slatebluebackground"], AsciiIndex["slategraybackground"], AsciiIndex["snowbackground"], AsciiIndex["springgreenbackground"], AsciiIndex["steelbluebackground"], AsciiIndex["tanbackground"], AsciiIndex["tealbackground"], AsciiIndex["thistlebackground"], AsciiIndex["tomatobackground"], AsciiIndex["turquoisebackground"], AsciiIndex["violetbackground"], AsciiIndex["wheatbackground"], AsciiIndex["whitebackground"], AsciiIndex["whitesmokebackground"], AsciiIndex["yellowbackground"], AsciiIndex["yellowgreenbackground"], AsciiIndex["alicebluebackgroundbackground"], AsciiIndex["antiquewhitebackgroundbackground"], AsciiIndex["aquabackgroundbackground"], AsciiIndex["aquamarinebackgroundbackground"], AsciiIndex["azurebackgroundbackground"], AsciiIndex["beigebackgroundbackground"], AsciiIndex["bisquebackgroundbackground"], AsciiIndex["blackbackgroundbackground"], AsciiIndex["blanchedalmondbackgroundbackground"], AsciiIndex["bluebackgroundbackground"], AsciiIndex["bluevioletbackgroundbackground"], AsciiIndex["brownbackgroundbackground"], AsciiIndex["burlywoodbackgroundbackground"], AsciiIndex["cadetbluebackgroundbackground"], AsciiIndex["chartreusebackgroundbackground"], AsciiIndex["chocolatebackgroundbackground"], AsciiIndex["coralbackgroundbackground"], AsciiIndex["cornflowerbluebackgroundbackground"], AsciiIndex["cornsilkbackgroundbackground"], AsciiIndex["crimsonbackgroundbackground"], AsciiIndex["cyanbackgroundbackground"], AsciiIndex["darkbluebackgroundbackground"], AsciiIndex["darkcyanbackgroundbackground"], AsciiIndex["darkgoldenrodbackgroundbackground"], AsciiIndex["darkgraybackgroundbackground"], AsciiIndex["darkgreenbackgroundbackground"], AsciiIndex["darkkhakibackgroundbackground"], AsciiIndex["darkolivegreenbackgroundbackground"], AsciiIndex["darkorangebackgroundbackground"], AsciiIndex["darkorchidbackgroundbackground"], AsciiIndex["darkredbackgroundbackground"], AsciiIndex["darksalmonbackgroundbackground"], AsciiIndex["darkseagreenbackgroundbackground"], AsciiIndex["darkslatebluebackgroundbackground"], AsciiIndex["darkslategraybackgroundbackground"], AsciiIndex["darkturquoisebackgroundbackground"], AsciiIndex["darkvioletbackgroundbackground"], AsciiIndex["deeppinkbackgroundbackground"], AsciiIndex["deepskybluebackgroundbackground"], AsciiIndex["dimgraybackgroundbackground"], AsciiIndex["dodgerbluebackgroundbackground"], AsciiIndex["firebrickbackgroundbackground"], AsciiIndex["floralwhitebackgroundbackground"], AsciiIndex["forestgreenbackgroundbackground"], AsciiIndex["fuchsiabackgroundbackground"], AsciiIndex["gainsborobackgroundbackground"], AsciiIndex["ghostwhitebackgroundbackground"], AsciiIndex["goldbackgroundbackground"], AsciiIndex["goldenrodbackgroundbackground"], AsciiIndex["graybackgroundbackground"], AsciiIndex["greenbackgroundbackground"], AsciiIndex["greenyellowbackgroundbackground"], AsciiIndex["honeydewbackgroundbackground"], AsciiIndex["hotpinkbackgroundbackground"], AsciiIndex["indianredbackgroundbackground"], AsciiIndex["indigobackgroundbackground"], AsciiIndex["ivorybackgroundbackground"], AsciiIndex["khakibackgroundbackground"], AsciiIndex["lavenderbackgroundbackground"], AsciiIndex["lavenderblushbackgroundbackground"], AsciiIndex["lawngreenbackgroundbackground"], AsciiIndex["lemonchiffonbackgroundbackground"], AsciiIndex["lightbluebackgroundbackground"], AsciiIndex["lightcoralbackgroundbackground"], AsciiIndex["lightcyanbackgroundbackground"], AsciiIndex["lightgoldenrodyellowbackgroundbackground"], AsciiIndex["lightgreenbackgroundbackground"], AsciiIndex["lightgreybackgroundbackground"], AsciiIndex["lightpinkbackgroundbackground"], AsciiIndex["lightsalmonbackgroundbackground"], AsciiIndex["lightseagreenbackgroundbackground"], AsciiIndex["lightskybluebackgroundbackground"], AsciiIndex["lightslategraybackgroundbackground"], AsciiIndex["lightsteelbluebackgroundbackground"], AsciiIndex["lightyellowbackgroundbackground"], AsciiIndex["limebackgroundbackground"], AsciiIndex["limegreenbackgroundbackground"], AsciiIndex["linenbackgroundbackground"], AsciiIndex["magentabackgroundbackground"], AsciiIndex["maroonbackgroundbackground"], AsciiIndex["mediumaquamarinebackgroundbackground"], AsciiIndex["mediumbluebackgroundbackground"], AsciiIndex["mediumorchidbackgroundbackground"], AsciiIndex["mediumpurplebackgroundbackground"], AsciiIndex["mediumseagreenbackgroundbackground"], AsciiIndex["mediumslatebluebackgroundbackground"], AsciiIndex["mediumspringgreenbackgroundbackground"], AsciiIndex["mediumturquoisebackgroundbackground"], AsciiIndex["mediumvioletredbackgroundbackground"], AsciiIndex["midnightbluebackgroundbackground"], AsciiIndex["mintcreambackgroundbackground"], AsciiIndex["mistyrosebackgroundbackground"], AsciiIndex["moccasinbackgroundbackground"], AsciiIndex["navajowhitebackgroundbackground"], AsciiIndex["navybackgroundbackground"], AsciiIndex["navybluebackgroundbackground"], AsciiIndex["oldlacebackgroundbackground"], AsciiIndex["olivebackgroundbackground"], AsciiIndex["olivedrabbackgroundbackground"], AsciiIndex["orangebackgroundbackground"], AsciiIndex["orangeredbackgroundbackground"], AsciiIndex["orchidbackgroundbackground"], AsciiIndex["palegoldenrodbackgroundbackground"], AsciiIndex["palegreenbackgroundbackground"], AsciiIndex["paleturquoisebackgroundbackground"], AsciiIndex["palevioletredbackgroundbackground"], AsciiIndex["papayawhipbackgroundbackground"], AsciiIndex["peachpuffbackgroundbackground"], AsciiIndex["perubackgroundbackground"], AsciiIndex["pinkbackgroundbackground"], AsciiIndex["plumbackgroundbackground"], AsciiIndex["powderbluebackgroundbackground"], AsciiIndex["purplebackgroundbackground"], AsciiIndex["redbackgroundbackground"], AsciiIndex["rosybrownbackgroundbackground"], AsciiIndex["royalbluebackgroundbackground"], AsciiIndex["saddlebrownbackgroundbackground"], AsciiIndex["salmonbackgroundbackground"], AsciiIndex["sandybrownbackgroundbackground"], AsciiIndex["seagreenbackgroundbackground"], AsciiIndex["seashellbackgroundbackground"], AsciiIndex["siennabackgroundbackground"], AsciiIndex["silverbackgroundbackground"], AsciiIndex["skybluebackgroundbackground"], AsciiIndex["slatebluebackgroundbackground"], AsciiIndex["slategraybackgroundbackground"], AsciiIndex["snowbackgroundbackground"], AsciiIndex["springgreenbackgroundbackground"], AsciiIndex["steelbluebackgroundbackground"], AsciiIndex["tanbackgroundbackground"], AsciiIndex["tealbackgroundbackground"], AsciiIndex["thistlebackgroundbackground"], AsciiIndex["tomatobackgroundbackground"], AsciiIndex["turquoisebackgroundbackground"], AsciiIndex["violetbackgroundbackground"], AsciiIndex["wheatbackgroundbackground"], AsciiIndex["whitebackgroundbackground"], AsciiIndex["whitesmokebackgroundbackground"], AsciiIndex["yellowbackgroundbackground"], AsciiIndex["yellowgreenbackgroundbackground"]}

// ColorNames is an array of all the css color names
var ColorNames = [...]string{"AliceBlue", "AntiqueWhite", "Aqua", "AquaMarine", "Azure", "Beige", "Bisque", "Black", "BlanchedAlmond", "Blue", "BlueViolet", "Brown", "BurlyWood", "CadetBlue", "Chartreuse", "Chocolate", "Coral", "CornFlowerBlue", "CornSilk", "Crimson", "DarkBlue", "DarkCyan", "DarkGoldenRod", "DarkGray", "DarkGreen", "DarkKhaki", "DarkMagenta", "DarkOliveGreen", "DarkOrange", "DarkOrchid", "DarkRed", "DarkSalmon", "DarkSeaGreen", "DarkSlateBlue", "DarkSlateGray", "DarkTurquoise", "DarkViolet", "DeepPink", "DeepSkyBlue", "DimGray", "DodgerBlue", "Firebrick", "FloralWhite", "ForestGreen", "Fuchsia", "Gainsboro", "GhostWhite", "Gold", "GoldenRod", "Gray", "Green", "GreenYellow", "Honeydew", "HotPink", "IndianRed", "Indigo", "Ivory", "Khaki", "Lavender", "LavenderBlush", "LawnGreen", "LemonChiffon", "LightBlue", "LightCoral", "LightCyan", "LightGoldenRodYellow", "LightGreen", "LightGrey", "LightPink", "LightSalmon", "LightSeaGreen", "LightSkyBlue", "LightSlateGray", "LightSteelBlue", "LightYellow", "Lime", "LimeGreen", "Linen", "Magenta", "Maroon", "MediumAquaMarine", "MediumBlue", "MediumOrchid", "MediumPurple", "MediumSeaGreen", "MediumSlateBlue", "MediumSpringGreen", "MediumTurquoise", "MediumVioletRed", "MidnightBlue", "MintCream", "MistyRose", "Moccasin", "NavajoWhite", "Navy", "NavyBlue", "OldLace", "Olive", "OliveDrab", "Orange", "OrangeRed", "Orchid", "PaleGoldenRod", "PaleGreen", "PaleTurquoise", "PaleVioletRed", "PapayaWhip", "PeachPuff", "Peru", "Pink", "Plum", "PowderBlue", "Purple", "RebeccaPurple", "Red", "RosyBrown", "RoyalBlue", "SaddleBrown", "Salmon", "SandyBrown", "SeaGreen", "SeaShell", "Sienna", "Silver", "SkyBlue", "SlateBlue", "SlateGray", "Snow", "SpringGreen", "SteelBlue", "Tan", "Teal", "Thistle", "Tomato", "Turquoise", "Violet", "Wheat", "White", "WhiteSmoke", "Yellow", "YellowGreen"}

// BackgroundNames is an array of all the css background color names
var BackgroundNames = [...]string{"AliceBlueBackground", "AntiqueWhiteBackground", "AquaBackground", "AquaMarineBackground", "AzureBackground", "BeigeBackground", "BisqueBackground", "BlackBackground", "BlanchedAlmondBackground", "BlueBackground", "BlueVioletBackground", "BrownBackground", "BurlyWoodBackground", "CadetBlueBackground", "ChartreuseBackground", "ChocolateBackground", "CoralBackground", "CornFlowerBlueBackground", "CornSilkBackground", "CrimsonBackground", "DarkBlueBackground", "DarkCyanBackground", "DarkGoldenRodBackground", "DarkGrayBackground", "DarkGreenBackground", "DarkKhakiBackground", "DarkMagentaBackground", "DarkOliveGreenBackground", "DarkOrangeBackground", "DarkOrchidBackground", "DarkRedBackground", "DarkSalmonBackground", "DarkSeaGreenBackground", "DarkSlateBlueBackground", "DarkSlateGrayBackground", "DarkTurquoiseBackground", "DarkVioletBackground", "DeepPinkBackground", "DeepSkyBlueBackground", "DimGrayBackground", "DodgerBlueBackground", "FirebrickBackground", "FloralWhiteBackground", "ForestGreenBackground", "FuchsiaBackground", "GainsboroBackground", "GhostWhiteBackground", "GoldBackground", "GoldenRodBackground", "GrayBackground", "GreenBackground", "GreenYellowBackground", "HoneydewBackground", "HotPinkBackground", "IndianRedBackground", "IndigoBackground", "IvoryBackground", "KhakiBackground", "LavenderBackground", "LavenderBlushBackground", "LawnGreenBackground", "LemonChiffonBackground", "LightBlueBackground", "LightCoralBackground", "LightCyanBackground", "LightGoldenRodYellowBackground", "LightGreenBackground", "LightGreyBackground", "LightPinkBackground", "LightSalmonBackground", "LightSeaGreenBackground", "LightSkyBlueBackground", "LightSlateGrayBackground", "LightSteelBlueBackground", "LightYellowBackground", "LimeBackground", "LimeGreenBackground", "LinenBackground", "MagentaBackground", "MaroonBackground", "MediumAquaMarineBackground", "MediumBlueBackground", "MediumOrchidBackground", "MediumPurpleBackground", "MediumSeaGreenBackground", "MediumSlateBlueBackground", "MediumSpringGreenBackground", "MediumTurquoiseBackground", "MediumVioletRedBackground", "MidnightBlueBackground", "MintCreamBackground", "MistyRoseBackground", "MoccasinBackground", "NavajoWhiteBackground", "NavyBackground", "NavyBlueBackground", "OldLaceBackground", "OliveBackground", "OliveDrabBackground", "OrangeBackground", "OrangeRedBackground", "OrchidBackground", "PaleGoldenRodBackground", "PaleGreenBackground", "PaleTurquoiseBackground", "PaleVioletRedBackground", "PapayaWhipBackground", "PeachPuffBackground", "PeruBackground", "PinkBackground", "PlumBackground", "PowderBlueBackground", "PurpleBackground", "RebeccaPurpleBackground", "RedBackground", "RosyBrownBackground", "RoyalBlueBackground", "SaddleBrownBackground", "SalmonBackground", "SandyBrownBackground", "SeaGreenBackground", "SeaShellBackground", "SiennaBackground", "SilverBackground", "SkyBlueBackground", "SlateBlueBackground", "SlateGrayBackground", "SnowBackground", "SpringGreenBackground", "SteelBlueBackground", "TanBackground", "TealBackground", "ThistleBackground", "TomatoBackground", "TurquoiseBackground", "VioletBackground", "WheatBackground", "WhiteBackground", "WhiteSmokeBackground", "YellowBackground", "YellowGreen"}

/* Color Functions */

// RandomHex returns a random css color as a hexidecimal value
func RandomHex() int32 {
	rand.Seed(time.Now().UnixNano())
	return HexColors[rand.Intn(len(HexColors))]
}

// RandomAscii returns a random css color as an ascii escape code
func RandomAscii() string {
	rand.Seed(time.Now().UnixNano())
	return AsciiColors[rand.Intn(len(AsciiColors))]
}

// RandomAsciiBackground returns a random css color as an ascii background escape code
func RandomAsciiBackground() string {
	rand.Seed(time.Now().UnixNano())
	return AsciiBackgrounds[rand.Intn(len(AsciiBackgrounds))]
}

// MultipleColor colorizes every rune in a string randomly from the AsciiColors
func MultipleColor(s string) string {
	var m string
	for _, r := range s {
		m += RandomAscii() + string(r)
	}
	m += AsciiReset
	return m
}

// MultipleBackground colorizes every rune's background in a string randomly from the AsciiColors
func MultipleBackground(s string) string {
	var m string
	for _, r := range s {
		m += RandomAsciiBackground() + string(r)
	}
	m += AsciiReset
	return m
}

/* Indexes that match names to colors in different types */

var HexIndex = map[string]int32{
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
	"limeGreen":            0x32CD32,
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

var AsciiIndex = map[string]string{
	"aliceblue":            "\033[38;2;240;248;255m",
	"antiquewhite":         "\033[38;2;250;235;215m",
	"aqua":                 "\033[38;2;0;255;255m",
	"aquamarine":           "\033[38;2;127;255;212m",
	"azure":                "\033[38;2;1240;255;255m",
	"beige":                "\033[38;2;245;245;220m",
	"bisque":               "\033[38;2;255;228;196m",
	"black":                "\033[38;2;0;0;0m",
	"blanchedalmond":       "\033[38;2;255;235;205m",
	"blue":                 "\033[38;2;0;0;255m",
	"blueviolet":           "\033[38;2;138;43;226m",
	"brown":                "\033[38;2;165;42;42m",
	"burlywood":            "\033[38;2;222;184;135m",
	"cadetblue":            "\033[38;2;95;158;160m",
	"chartreuse":           "\033[38;2;95;158;160m",
	"chocolate":            "\033[38;2;210;105;30m",
	"coral":                "\033[38;2;255;127;80m",
	"cornflowerblue":       "\033[38;2;100;149;237m",
	"cornsilk":             "\033[38;2;255;248;220m",
	"crimson":              "\033[38;2;220;20;60m",
	"cyan":                 "\033[38;2;0;255;255m",
	"darkblue":             "\033[38;2;0;0;139m",
	"darkcyan":             "\033[38;2;0;139;139m",
	"darkgoldenrod":        "\033[38;2;184;134;11m",
	"darkgray":             "\033[38;2;169;169;169m",
	"darkgreen":            "\033[38;2;0;100;0m",
	"darkkhaki":            "\033[38;2;189;183;107m",
	"darkmagenta":          "\033[38;2;139;0;139m",
	"darkolivegreen":       "\033[38;2;85;107;47m",
	"darkorange":           "\033[38;2;255;140;0m",
	"darkorchid":           "\033[38;2;153;50;204m",
	"darkred":              "\033[38;2;139;0;0m",
	"darksalmon":           "\033[38;2;233;150;122m",
	"darkseagreen":         "\033[38;2;143;188;143m",
	"darkslateblue":        "\033[38;2;72;61;139m",
	"darkslategray":        "\033[38;2;47;79;79m",
	"darkturquoise":        "\033[38;2;0;206;209m",
	"darkviolet":           "\033[38;2;148;0;211m",
	"deeppink":             "\033[38;2;255;20;147m",
	"deepskyblue":          "\033[38;2;0;191;255m",
	"dimgray":              "\033[38;2;0;191;255m",
	"dodgerblue":           "\033[38;2;30;144;255m",
	"firebrick":            "\033[38;2;178;34;34m",
	"floralwhite":          "\033[38;2;255;250;240m",
	"forestgreen":          "\033[38;2;34;139;34m",
	"fuchsia":              "\033[38;2;255;0;255m",
	"gainsboro":            "\033[38;2;220;220;220m",
	"ghostwhite":           "\033[38;2;248;248;255m",
	"gold":                 "\033[38;2;255;215;0m",
	"goldenrod":            "\033[38;2;218;165;32m",
	"gray":                 "\033[38;2;127;127;127m",
	"green":                "\033[38;2;0;128;0m",
	"greenyellow":          "\033[38;2;173;255;47m",
	"honeydew":             "\033[38;2;240;255;240m",
	"hotpink":              "\033[38;2;255;105;180m",
	"indianred":            "\033[38;2;205;92;92m",
	"indigo":               "\033[38;2;75;0;130m",
	"ivory":                "\033[38;2;255;255;240m",
	"khaki":                "\033[38;2;240;230;140m",
	"lavender":             "\033[38;2;230;230;250m",
	"lavenderblush":        "\033[38;2;255;240;245m",
	"lawngreen":            "\033[38;2;124;252;0m",
	"lemonchiffon":         "\033[38;2;255;250;205m",
	"lightblue":            "\033[38;2;173;216;230m",
	"lightcoral":           "\033[38;2;240;128;128m",
	"lightcyan":            "\033[38;2;224;255;255m",
	"lightgoldenrodyellow": "\033[38;2;250;250;210m",
	"lightgreen":           "\033[38;2;144;238;144m",
	"lightgrey":            "\033[38;2;211;211;211m",
	"lightpink":            "\033[38;2;255;182;193m",
	"lightsalmon":          "\033[38;2;255;160;122m",
	"lightseagreen":        "\033[38;2;32;178;170m",
	"lightskyblue":         "\033[38;2;135;206;250m",
	"lightslategray":       "\033[38;2;119;136;153m",
	"lightsteelblue":       "\033[38;2;176;196;222m",
	"lightyellow":          "\033[38;2;255;255;224m",
	"lime":                 "\033[38;2;0;255;0m",
	"limegreen":            "\033[38;2;50;205;50m",
	"linen":                "\033[38;2;250;240;230m",
	"magenta":              "\033[38;2;255;0;255m",
	"maroon":               "\033[38;2;128;0;0m",
	"mediumaquamarine":     "\033[38;2;102;205;170m",
	"mediumblue":           "\033[38;2;0;0;205m",
	"mediumorchid":         "\033[38;2;186;85;211m",
	"mediumpurple":         "\033[38;2;147;112;219m",
	"mediumseagreen":       "\033[38;2;60;179;113m",
	"mediumslateblue":      "\033[38;2;123;104;238m",
	"mediumspringgreen":    "\033[38;2;0;250;154m",
	"mediumturquoise":      "\033[38;2;72;209;204m",
	"mediumvioletred":      "\033[38;2;199;21;133m",
	"midnightblue":         "\033[38;2;25;25;112m",
	"mintcream":            "\033[38;2;245;255;250m",
	"mistyrose":            "\033[38;2;255;228;225m",
	"moccasin":             "\033[38;2;255;228;181m",
	"navajowhite":          "\033[38;2;255;222;173m",
	"navy":                 "\033[38;2;0;0;128m",
	"navyblue":             "\033[38;2;159;175;223m",
	"oldlace":              "\033[38;2;253;245;230m",
	"olive":                "\033[38;2;128;128;0m",
	"olivedrab":            "\033[38;2;107;142;35m",
	"orange":               "\033[38;2;255;165;0m",
	"orangered":            "\033[38;2;255;69;0m",
	"orchid":               "\033[38;2;218;112;214m",
	"palegoldenrod":        "\033[38;2;238;232;170m",
	"palegreen":            "\033[38;2;152;251;152m",
	"paleturquoise":        "\033[38;2;175;238;238m",
	"palevioletred":        "\033[38;2;219;112;147m",
	"papayawhip":           "\033[38;2;255;239;213m",
	"peachpuff":            "\033[38;2;255;218;185m",
	"peru":                 "\033[38;2;205;133;63m",
	"pink":                 "\033[38;2;255;192;203m",
	"plum":                 "\033[38;2;221;160;221m",
	"powderblue":           "\033[38;2;176;224;230m",
	"purple":               "\033[38;2;128;0;128m",
	"red":                  "\033[38;2;255;0;0m",
	"rosybrown":            "\033[38;2;188;143;143m",
	"royalblue":            "\033[38;2;65;105;225m",
	"saddlebrown":          "\033[38;2;139;69;19m",
	"salmon":               "\033[38;2;250;128;114m",
	"sandybrown":           "\033[38;2;244;164;96m",
	"seagreen":             "\033[38;2;46;139;87m",
	"seashell":             "\033[38;2;255;245;238m",
	"sienna":               "\033[38;2;160;82;45m",
	"silver":               "\033[38;2;192;192;192m",
	"skyblue":              "\033[38;2;135;206;235m",
	"slateblue":            "\033[38;2;106;90;205m",
	"slategray":            "\033[38;2;112;128;144m",
	"snow":                 "\033[38;2;255;250;250m",
	"springgreen":          "\033[38;2;0;255;127m",
	"steelblue":            "\033[38;2;70;130;180m",
	"tan":                  "\033[38;2;210;180;140m",
	"teal":                 "\033[38;2;0;128;128m",
	"thistle":              "\033[38;2;216;191;216m",
	"tomato":               "\033[38;2;255;99;71m",
	"turquoise":            "\033[38;2;64;224;208m",
	"violet":               "\033[38;2;238;130;238m",
	"wheat":                "\033[38;2;245;222;179m",
	"white":                "\033[38;2;255;255;255m",
	"whitesmoke":           "\033[38;2;245;245;245m",
	"yellow":               "\033[38;2;255;255;0m",
	"yellowgreen":          "\033[38;2;139;205;50m",
}

var AsciiBackgroundIndex = map[string]string{
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
