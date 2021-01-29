// Package css provides all the css colors.
package css

import (
	"math/rand"
	"time"
)

/* Color Hex Value Constants */

//
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
	LightGray            = 0x90EE90
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

	// Ansi escape codes that have no hex equivalent:
	AnsiReset = "\x1b[0m"
	AnsiX     = "\x1b[0m"
)

/* Lists of Colors */

// HexColors is an array of all of the css colors in hexadecimal format.
var HexColors = [...]int32{AliceBlue, AntiqueWhite, Aqua, AquaMarine, Azure, Beige, Bisque, Black, BlanchedAlmond, Blue, BlueViolet, Brown, BurlyWood, CadetBlue, Chartreuse, Chocolate, Coral, CornFlowerBlue, CornSilk, Crimson, Cyan, DarkBlue, DarkCyan, DarkGoldenRod, DarkGray, DarkGreen, DarkKhaki, DarkMagenta, DarkOliveGreen, DarkOrange, DarkOrchid, DarkRed, DarkSalmon, DarkSeaGreen, DarkSlateBlue, DarkSlateGray, DarkTurquoise, DarkViolet, DeepPink, DeepSkyBlue, DimGray, DodgerBlue, Firebrick, FloralWhite, ForestGreen, Fuchsia, Gainsboro, GhostWhite, Gold, GoldenRod, Gray, Green, GreenYellow, Honeydew, HotPink, IndianRed, Indigo, Ivory, Khaki, Lavender, LavenderBlush, LawnGreen, LemonChiffon, LightBlue, LightCoral, LightCyan, LightGoldenRodYellow, LightGreen, LightGray, LightPink, LightSalmon, LightSeaGreen, LightSkyBlue, LightSlateGray, LightSteelBlue, LightYellow, Lime, LimeGreen, Linen, Magenta, Maroon, MediumAquaMarine, MediumBlue, MediumOrchid, MediumPurple, MediumSeaGreen, MediumSlateBlue, MediumSpringGreen, MediumTurquoise, MediumVioletRed, MidnightBlue, MintCream, MistyRose, Moccasin, NavajoWhite, Navy, NavyBlue, OldLace, Olive, OliveDrab, Orange, OrangeRed, Orchid, PaleGoldenRod, PaleGreen, PaleTurquoise, PaleVioletRed, PapayaWhip, PeachPuff, Peru, Pink, Plum, PowderBlue, Purple, RebeccaPurple, Red, RosyBrown, RoyalBlue, SaddleBrown, Salmon, SandyBrown, SeaGreen, SeaShell, Sienna, Silver, SkyBlue, SlateBlue, SlateGray, Snow, SpringGreen, SteelBlue, Tan, Teal, Thistle, Tomato, Turquoise, Violet, Wheat, White, WhiteSmoke, Yellow, YellowGreen}

// AnsiColors is an array of all the css colors as ANSI escape codes.
var AnsiColors = [...]string{AnsiIndex["sol.yellow"], AnsiIndex["sol.orange"], AnsiIndex["sol.red"], AnsiIndex["sol.magenta"], AnsiIndex["sol.violet"], AnsiIndex["sol.blue"], AnsiIndex["sol.cyan"], AnsiIndex["sol.green"], AnsiIndex["aliceblue"], AnsiIndex["antiquewhite"], AnsiIndex["aqua"], AnsiIndex["aquamarine"], AnsiIndex["azure"], AnsiIndex["beige"], AnsiIndex["bisque"], AnsiIndex["black"], AnsiIndex["blanchedalmond"], AnsiIndex["blue"], AnsiIndex["blueviolet"], AnsiIndex["brown"], AnsiIndex["burlywood"], AnsiIndex["cadetblue"], AnsiIndex["chartreuse"], AnsiIndex["chocolate"], AnsiIndex["coral"], AnsiIndex["cornflowerblue"], AnsiIndex["cornsilk"], AnsiIndex["crimson"], AnsiIndex["cyan"], AnsiIndex["darkblue"], AnsiIndex["darkcyan"], AnsiIndex["darkgoldenrod"], AnsiIndex["darkgray"], AnsiIndex["darkgreen"], AnsiIndex["darkkhaki"], AnsiIndex["darkmagenta"], AnsiIndex["darkolivegreen"], AnsiIndex["darkorange"], AnsiIndex["darkorchid"], AnsiIndex["darkred"], AnsiIndex["darksalmon"], AnsiIndex["darkseagreen"], AnsiIndex["darkslateblue"], AnsiIndex["darkslategray"], AnsiIndex["darkturquoise"], AnsiIndex["darkviolet"], AnsiIndex["deeppink"], AnsiIndex["deepskyblue"], AnsiIndex["dimgray"], AnsiIndex["dodgerblue"], AnsiIndex["firebrick"], AnsiIndex["floralwhite"], AnsiIndex["forestgreen"], AnsiIndex["fuchsia"], AnsiIndex["gainsboro"], AnsiIndex["ghostwhite"], AnsiIndex["gold"], AnsiIndex["goldenrod"], AnsiIndex["gray"], AnsiIndex["green"], AnsiIndex["greenyellow"], AnsiIndex["honeydew"], AnsiIndex["hotpink"], AnsiIndex["indianred"], AnsiIndex["indigo"], AnsiIndex["ivory"], AnsiIndex["khaki"], AnsiIndex["lavender"], AnsiIndex["lavenderblush"], AnsiIndex["lawngreen"], AnsiIndex["lemonchiffon"], AnsiIndex["lightblue"], AnsiIndex["lightcoral"], AnsiIndex["lightcyan"], AnsiIndex["lightgoldenrodyellow"], AnsiIndex["lightgreen"], AnsiIndex["lightgray"], AnsiIndex["lightpink"], AnsiIndex["lightsalmon"], AnsiIndex["lightseagreen"], AnsiIndex["lightskyblue"], AnsiIndex["lightslategray"], AnsiIndex["lightsteelblue"], AnsiIndex["lightyellow"], AnsiIndex["lime"], AnsiIndex["limegreen"], AnsiIndex["linen"], AnsiIndex["magenta"], AnsiIndex["maroon"], AnsiIndex["mediumaquamarine"], AnsiIndex["mediumblue"], AnsiIndex["mediumorchid"], AnsiIndex["mediumpurple"], AnsiIndex["mediumseagreen"], AnsiIndex["mediumslateblue"], AnsiIndex["mediumspringgreen"], AnsiIndex["mediumturquoise"], AnsiIndex["mediumvioletred"], AnsiIndex["midnightblue"], AnsiIndex["mintcream"], AnsiIndex["mistyrose"], AnsiIndex["moccasin"], AnsiIndex["navajowhite"], AnsiIndex["navy"], AnsiIndex["navyblue"], AnsiIndex["oldlace"], AnsiIndex["olive"], AnsiIndex["olivedrab"], AnsiIndex["orange"], AnsiIndex["orangered"], AnsiIndex["orchid"], AnsiIndex["palegoldenrod"], AnsiIndex["palegreen"], AnsiIndex["paleturquoise"], AnsiIndex["palevioletred"], AnsiIndex["papayawhip"], AnsiIndex["peachpuff"], AnsiIndex["peru"], AnsiIndex["pink"], AnsiIndex["plum"], AnsiIndex["powderblue"], AnsiIndex["purple"], AnsiIndex["red"], AnsiIndex["rosybrown"], AnsiIndex["royalblue"], AnsiIndex["saddlebrown"], AnsiIndex["salmon"], AnsiIndex["sandybrown"], AnsiIndex["seagreen"], AnsiIndex["seashell"], AnsiIndex["sienna"], AnsiIndex["silver"], AnsiIndex["skyblue"], AnsiIndex["slateblue"], AnsiIndex["slategray"], AnsiIndex["snow"], AnsiIndex["springgreen"], AnsiIndex["steelblue"], AnsiIndex["tan"], AnsiIndex["teal"], AnsiIndex["thistle"], AnsiIndex["tomato"], AnsiIndex["turquoise"], AnsiIndex["violet"], AnsiIndex["wheat"], AnsiIndex["white"], AnsiIndex["whitesmoke"], AnsiIndex["yellow"], AnsiIndex["yellowgreen"]}

// AnsiBackgrounds is an array of all the css colors as ANSI background escape codes
var AnsiBackgrounds = [...]string{AnsiIndex["antiquewhitebackground"], AnsiIndex["aquabackground"], AnsiIndex["aquamarinebackground"], AnsiIndex["azurebackground"], AnsiIndex["beigebackground"], AnsiIndex["bisquebackground"], AnsiIndex["blackbackground"], AnsiIndex["blanchedalmondbackground"], AnsiIndex["bluebackground"], AnsiIndex["bluevioletbackground"], AnsiIndex["brownbackground"], AnsiIndex["burlywoodbackground"], AnsiIndex["cadetbluebackground"], AnsiIndex["chartreusebackground"], AnsiIndex["chocolatebackground"], AnsiIndex["coralbackground"], AnsiIndex["cornflowerbluebackground"], AnsiIndex["cornsilkbackground"], AnsiIndex["crimsonbackground"], AnsiIndex["cyanbackground"], AnsiIndex["darkbluebackground"], AnsiIndex["darkcyanbackground"], AnsiIndex["darkgoldenrodbackground"], AnsiIndex["darkgraybackground"], AnsiIndex["darkgreenbackground"], AnsiIndex["darkkhakibackground"], AnsiIndex["darkmagentabackground"], AnsiIndex["darkolivegreenbackground"], AnsiIndex["darkorangebackground"], AnsiIndex["darkorchidbackground"], AnsiIndex["darkredbackground"], AnsiIndex["darksalmonbackground"], AnsiIndex["darkseagreenbackground"], AnsiIndex["darkslatebluebackground"], AnsiIndex["darkslategraybackground"], AnsiIndex["darkturquoisebackground"], AnsiIndex["darkvioletbackground"], AnsiIndex["deeppinkbackground"], AnsiIndex["deepskybluebackground"], AnsiIndex["dimgraybackground"], AnsiIndex["dodgerbluebackground"], AnsiIndex["firebrickbackground"], AnsiIndex["floralwhitebackground"], AnsiIndex["forestgreenbackground"], AnsiIndex["fuchsiabackground"], AnsiIndex["gainsborobackground"], AnsiIndex["ghostwhitebackground"], AnsiIndex["goldbackground"], AnsiIndex["goldenrodbackground"], AnsiIndex["graybackground"], AnsiIndex["greenbackground"], AnsiIndex["greenyellowbackground"], AnsiIndex["honeydewbackground"], AnsiIndex["hotpinkbackground"], AnsiIndex["indianredbackground"], AnsiIndex["indigobackground"], AnsiIndex["ivorybackground"], AnsiIndex["khakibackground"], AnsiIndex["lavenderbackground"], AnsiIndex["lavenderblushbackground"], AnsiIndex["lawngreenbackground"], AnsiIndex["lemonchiffonbackground"], AnsiIndex["lightbluebackground"], AnsiIndex["lightcoralbackground"], AnsiIndex["lightcyanbackground"], AnsiIndex["lightgoldenrodyellowbackground"], AnsiIndex["lightgreenbackground"], AnsiIndex["lightgraybackground"], AnsiIndex["lightpinkbackground"], AnsiIndex["lightsalmonbackground"], AnsiIndex["lightseagreenbackground"], AnsiIndex["lightskybluebackground"], AnsiIndex["lightslategraybackground"], AnsiIndex["lightsteelbluebackground"], AnsiIndex["lightyellowbackground"], AnsiIndex["limebackground"], AnsiIndex["limegreenbackground"], AnsiIndex["linenbackground"], AnsiIndex["magentabackground"], AnsiIndex["maroonbackground"], AnsiIndex["mediumaquamarinebackground"], AnsiIndex["mediumbluebackground"], AnsiIndex["mediumorchidbackground"], AnsiIndex["mediumpurplebackground"], AnsiIndex["mediumseagreenbackground"], AnsiIndex["mediumslatebluebackground"], AnsiIndex["mediumspringgreenbackground"], AnsiIndex["mediumturquoisebackground"], AnsiIndex["mediumvioletredbackground"], AnsiIndex["midnightbluebackground"], AnsiIndex["mintcreambackground"], AnsiIndex["mistyrosebackground"], AnsiIndex["moccasinbackground"], AnsiIndex["navajowhitebackground"], AnsiIndex["navybackground"], AnsiIndex["navybluebackground"], AnsiIndex["oldlacebackground"], AnsiIndex["olivebackground"], AnsiIndex["olivedrabbackground"], AnsiIndex["orangebackground"], AnsiIndex["orangeredbackground"], AnsiIndex["orchidbackground"], AnsiIndex["palegoldenrodbackground"], AnsiIndex["palegreenbackground"], AnsiIndex["paleturquoisebackground"], AnsiIndex["palevioletredbackground"], AnsiIndex["papayawhipbackground"], AnsiIndex["peachpuffbackground"], AnsiIndex["perubackground"], AnsiIndex["pinkbackground"], AnsiIndex["plumbackground"], AnsiIndex["powderbluebackground"], AnsiIndex["purplebackground"], AnsiIndex["redbackground"], AnsiIndex["rosybrownbackground"], AnsiIndex["royalbluebackground"], AnsiIndex["saddlebrownbackground"], AnsiIndex["salmonbackground"], AnsiIndex["sandybrownbackground"], AnsiIndex["seagreenbackground"], AnsiIndex["seashellbackground"], AnsiIndex["siennabackground"], AnsiIndex["silverbackground"], AnsiIndex["skybluebackground"], AnsiIndex["slatebluebackground"], AnsiIndex["slategraybackground"], AnsiIndex["snowbackground"], AnsiIndex["springgreenbackground"], AnsiIndex["steelbluebackground"], AnsiIndex["tanbackground"], AnsiIndex["tealbackground"], AnsiIndex["thistlebackground"], AnsiIndex["tomatobackground"], AnsiIndex["turquoisebackground"], AnsiIndex["violetbackground"], AnsiIndex["wheatbackground"], AnsiIndex["whitebackground"], AnsiIndex["whitesmokebackground"], AnsiIndex["yellowbackground"], AnsiIndex["yellowgreenbackground"], AnsiIndex["alicebluebackgroundbackground"], AnsiIndex["antiquewhitebackgroundbackground"], AnsiIndex["aquabackgroundbackground"], AnsiIndex["aquamarinebackgroundbackground"], AnsiIndex["azurebackgroundbackground"], AnsiIndex["beigebackgroundbackground"], AnsiIndex["bisquebackgroundbackground"], AnsiIndex["blackbackgroundbackground"], AnsiIndex["blanchedalmondbackgroundbackground"], AnsiIndex["bluebackgroundbackground"], AnsiIndex["bluevioletbackgroundbackground"], AnsiIndex["brownbackgroundbackground"], AnsiIndex["burlywoodbackgroundbackground"], AnsiIndex["cadetbluebackgroundbackground"], AnsiIndex["chartreusebackgroundbackground"], AnsiIndex["chocolatebackgroundbackground"], AnsiIndex["coralbackgroundbackground"], AnsiIndex["cornflowerbluebackgroundbackground"], AnsiIndex["cornsilkbackgroundbackground"], AnsiIndex["crimsonbackgroundbackground"], AnsiIndex["cyanbackgroundbackground"], AnsiIndex["darkbluebackgroundbackground"], AnsiIndex["darkcyanbackgroundbackground"], AnsiIndex["darkgoldenrodbackgroundbackground"], AnsiIndex["darkgraybackgroundbackground"], AnsiIndex["darkgreenbackgroundbackground"], AnsiIndex["darkkhakibackgroundbackground"], AnsiIndex["darkolivegreenbackgroundbackground"], AnsiIndex["darkorangebackgroundbackground"], AnsiIndex["darkorchidbackgroundbackground"], AnsiIndex["darkredbackgroundbackground"], AnsiIndex["darksalmonbackgroundbackground"], AnsiIndex["darkseagreenbackgroundbackground"], AnsiIndex["darkslatebluebackgroundbackground"], AnsiIndex["darkslategraybackgroundbackground"], AnsiIndex["darkturquoisebackgroundbackground"], AnsiIndex["darkvioletbackgroundbackground"], AnsiIndex["deeppinkbackgroundbackground"], AnsiIndex["deepskybluebackgroundbackground"], AnsiIndex["dimgraybackgroundbackground"], AnsiIndex["dodgerbluebackgroundbackground"], AnsiIndex["firebrickbackgroundbackground"], AnsiIndex["floralwhitebackgroundbackground"], AnsiIndex["forestgreenbackgroundbackground"], AnsiIndex["fuchsiabackgroundbackground"], AnsiIndex["gainsborobackgroundbackground"], AnsiIndex["ghostwhitebackgroundbackground"], AnsiIndex["goldbackgroundbackground"], AnsiIndex["goldenrodbackgroundbackground"], AnsiIndex["graybackgroundbackground"], AnsiIndex["greenbackgroundbackground"], AnsiIndex["greenyellowbackgroundbackground"], AnsiIndex["honeydewbackgroundbackground"], AnsiIndex["hotpinkbackgroundbackground"], AnsiIndex["indianredbackgroundbackground"], AnsiIndex["indigobackgroundbackground"], AnsiIndex["ivorybackgroundbackground"], AnsiIndex["khakibackgroundbackground"], AnsiIndex["lavenderbackgroundbackground"], AnsiIndex["lavenderblushbackgroundbackground"], AnsiIndex["lawngreenbackgroundbackground"], AnsiIndex["lemonchiffonbackgroundbackground"], AnsiIndex["lightbluebackgroundbackground"], AnsiIndex["lightcoralbackgroundbackground"], AnsiIndex["lightcyanbackgroundbackground"], AnsiIndex["lightgoldenrodyellowbackgroundbackground"], AnsiIndex["lightgreenbackgroundbackground"], AnsiIndex["lightgraybackgroundbackground"], AnsiIndex["lightpinkbackgroundbackground"], AnsiIndex["lightsalmonbackgroundbackground"], AnsiIndex["lightseagreenbackgroundbackground"], AnsiIndex["lightskybluebackgroundbackground"], AnsiIndex["lightslategraybackgroundbackground"], AnsiIndex["lightsteelbluebackgroundbackground"], AnsiIndex["lightyellowbackgroundbackground"], AnsiIndex["limebackgroundbackground"], AnsiIndex["limegreenbackgroundbackground"], AnsiIndex["linenbackgroundbackground"], AnsiIndex["magentabackgroundbackground"], AnsiIndex["maroonbackgroundbackground"], AnsiIndex["mediumaquamarinebackgroundbackground"], AnsiIndex["mediumbluebackgroundbackground"], AnsiIndex["mediumorchidbackgroundbackground"], AnsiIndex["mediumpurplebackgroundbackground"], AnsiIndex["mediumseagreenbackgroundbackground"], AnsiIndex["mediumslatebluebackgroundbackground"], AnsiIndex["mediumspringgreenbackgroundbackground"], AnsiIndex["mediumturquoisebackgroundbackground"], AnsiIndex["mediumvioletredbackgroundbackground"], AnsiIndex["midnightbluebackgroundbackground"], AnsiIndex["mintcreambackgroundbackground"], AnsiIndex["mistyrosebackgroundbackground"], AnsiIndex["moccasinbackgroundbackground"], AnsiIndex["navajowhitebackgroundbackground"], AnsiIndex["navybackgroundbackground"], AnsiIndex["navybluebackgroundbackground"], AnsiIndex["oldlacebackgroundbackground"], AnsiIndex["olivebackgroundbackground"], AnsiIndex["olivedrabbackgroundbackground"], AnsiIndex["orangebackgroundbackground"], AnsiIndex["orangeredbackgroundbackground"], AnsiIndex["orchidbackgroundbackground"], AnsiIndex["palegoldenrodbackgroundbackground"], AnsiIndex["palegreenbackgroundbackground"], AnsiIndex["paleturquoisebackgroundbackground"], AnsiIndex["palevioletredbackgroundbackground"], AnsiIndex["papayawhipbackgroundbackground"], AnsiIndex["peachpuffbackgroundbackground"], AnsiIndex["perubackgroundbackground"], AnsiIndex["pinkbackgroundbackground"], AnsiIndex["plumbackgroundbackground"], AnsiIndex["powderbluebackgroundbackground"], AnsiIndex["purplebackgroundbackground"], AnsiIndex["redbackgroundbackground"], AnsiIndex["rosybrownbackgroundbackground"], AnsiIndex["royalbluebackgroundbackground"], AnsiIndex["saddlebrownbackgroundbackground"], AnsiIndex["salmonbackgroundbackground"], AnsiIndex["sandybrownbackgroundbackground"], AnsiIndex["seagreenbackgroundbackground"], AnsiIndex["seashellbackgroundbackground"], AnsiIndex["siennabackgroundbackground"], AnsiIndex["silverbackgroundbackground"], AnsiIndex["skybluebackgroundbackground"], AnsiIndex["slatebluebackgroundbackground"], AnsiIndex["slategraybackgroundbackground"], AnsiIndex["snowbackgroundbackground"], AnsiIndex["springgreenbackgroundbackground"], AnsiIndex["steelbluebackgroundbackground"], AnsiIndex["tanbackgroundbackground"], AnsiIndex["tealbackgroundbackground"], AnsiIndex["thistlebackgroundbackground"], AnsiIndex["tomatobackgroundbackground"], AnsiIndex["turquoisebackgroundbackground"], AnsiIndex["violetbackgroundbackground"], AnsiIndex["wheatbackgroundbackground"], AnsiIndex["whitebackgroundbackground"], AnsiIndex["whitesmokebackgroundbackground"], AnsiIndex["yellowbackgroundbackground"], AnsiIndex["yellowgreenbackgroundbackground"]}

// ColorNames is an array of all the css color names
var ColorNames = [...]string{"AliceBlue", "AntiqueWhite", "Aqua", "AquaMarine", "Azure", "Beige", "Bisque", "Black", "BlanchedAlmond", "Blue", "BlueViolet", "Brown", "BurlyWood", "CadetBlue", "Chartreuse", "Chocolate", "Coral", "CornFlowerBlue", "CornSilk", "Crimson", "DarkBlue", "DarkCyan", "DarkGoldenRod", "DarkGray", "DarkGreen", "DarkKhaki", "DarkMagenta", "DarkOliveGreen", "DarkOrange", "DarkOrchid", "DarkRed", "DarkSalmon", "DarkSeaGreen", "DarkSlateBlue", "DarkSlateGray", "DarkTurquoise", "DarkViolet", "DeepPink", "DeepSkyBlue", "DimGray", "DodgerBlue", "Firebrick", "FloralWhite", "ForestGreen", "Fuchsia", "Gainsboro", "GhostWhite", "Gold", "GoldenRod", "Gray", "Green", "GreenYellow", "Honeydew", "HotPink", "IndianRed", "Indigo", "Ivory", "Khaki", "Lavender", "LavenderBlush", "LawnGreen", "LemonChiffon", "LightBlue", "LightCoral", "LightCyan", "LightGoldenRodYellow", "LightGreen", "LightGray", "LightPink", "LightSalmon", "LightSeaGreen", "LightSkyBlue", "LightSlateGray", "LightSteelBlue", "LightYellow", "Lime", "LimeGreen", "Linen", "Magenta", "Maroon", "MediumAquaMarine", "MediumBlue", "MediumOrchid", "MediumPurple", "MediumSeaGreen", "MediumSlateBlue", "MediumSpringGreen", "MediumTurquoise", "MediumVioletRed", "MidnightBlue", "MintCream", "MistyRose", "Moccasin", "NavajoWhite", "Navy", "NavyBlue", "OldLace", "Olive", "OliveDrab", "Orange", "OrangeRed", "Orchid", "PaleGoldenRod", "PaleGreen", "PaleTurquoise", "PaleVioletRed", "PapayaWhip", "PeachPuff", "Peru", "Pink", "Plum", "PowderBlue", "Purple", "RebeccaPurple", "Red", "RosyBrown", "RoyalBlue", "SaddleBrown", "Salmon", "SandyBrown", "SeaGreen", "SeaShell", "Sienna", "Silver", "SkyBlue", "SlateBlue", "SlateGray", "Snow", "SpringGreen", "SteelBlue", "Tan", "Teal", "Thistle", "Tomato", "Turquoise", "Violet", "Wheat", "White", "WhiteSmoke", "Yellow", "YellowGreen"}

// BackgroundNames is an array of all the css background color names
var BackgroundNames = [...]string{"AliceBlueBackground", "AntiqueWhiteBackground", "AquaBackground", "AquaMarineBackground", "AzureBackground", "BeigeBackground", "BisqueBackground", "BlackBackground", "BlanchedAlmondBackground", "BlueBackground", "BlueVioletBackground", "BrownBackground", "BurlyWoodBackground", "CadetBlueBackground", "ChartreuseBackground", "ChocolateBackground", "CoralBackground", "CornFlowerBlueBackground", "CornSilkBackground", "CrimsonBackground", "DarkBlueBackground", "DarkCyanBackground", "DarkGoldenRodBackground", "DarkGrayBackground", "DarkGreenBackground", "DarkKhakiBackground", "DarkMagentaBackground", "DarkOliveGreenBackground", "DarkOrangeBackground", "DarkOrchidBackground", "DarkRedBackground", "DarkSalmonBackground", "DarkSeaGreenBackground", "DarkSlateBlueBackground", "DarkSlateGrayBackground", "DarkTurquoiseBackground", "DarkVioletBackground", "DeepPinkBackground", "DeepSkyBlueBackground", "DimGrayBackground", "DodgerBlueBackground", "FirebrickBackground", "FloralWhiteBackground", "ForestGreenBackground", "FuchsiaBackground", "GainsboroBackground", "GhostWhiteBackground", "GoldBackground", "GoldenRodBackground", "GrayBackground", "GreenBackground", "GreenYellowBackground", "HoneydewBackground", "HotPinkBackground", "IndianRedBackground", "IndigoBackground", "IvoryBackground", "KhakiBackground", "LavenderBackground", "LavenderBlushBackground", "LawnGreenBackground", "LemonChiffonBackground", "LightBlueBackground", "LightCoralBackground", "LightCyanBackground", "LightGoldenRodYellowBackground", "LightGreenBackground", "LightGrayBackground", "LightPinkBackground", "LightSalmonBackground", "LightSeaGreenBackground", "LightSkyBlueBackground", "LightSlateGrayBackground", "LightSteelBlueBackground", "LightYellowBackground", "LimeBackground", "LimeGreenBackground", "LinenBackground", "MagentaBackground", "MaroonBackground", "MediumAquaMarineBackground", "MediumBlueBackground", "MediumOrchidBackground", "MediumPurpleBackground", "MediumSeaGreenBackground", "MediumSlateBlueBackground", "MediumSpringGreenBackground", "MediumTurquoiseBackground", "MediumVioletRedBackground", "MidnightBlueBackground", "MintCreamBackground", "MistyRoseBackground", "MoccasinBackground", "NavajoWhiteBackground", "NavyBackground", "NavyBlueBackground", "OldLaceBackground", "OliveBackground", "OliveDrabBackground", "OrangeBackground", "OrangeRedBackground", "OrchidBackground", "PaleGoldenRodBackground", "PaleGreenBackground", "PaleTurquoiseBackground", "PaleVioletRedBackground", "PapayaWhipBackground", "PeachPuffBackground", "PeruBackground", "PinkBackground", "PlumBackground", "PowderBlueBackground", "PurpleBackground", "RebeccaPurpleBackground", "RedBackground", "RosyBrownBackground", "RoyalBlueBackground", "SaddleBrownBackground", "SalmonBackground", "SandyBrownBackground", "SeaGreenBackground", "SeaShellBackground", "SiennaBackground", "SilverBackground", "SkyBlueBackground", "SlateBlueBackground", "SlateGrayBackground", "SnowBackground", "SpringGreenBackground", "SteelBlueBackground", "TanBackground", "TealBackground", "ThistleBackground", "TomatoBackground", "TurquoiseBackground", "VioletBackground", "WheatBackground", "WhiteBackground", "WhiteSmokeBackground", "YellowBackground", "YellowGreen"}

/* Color Functions */

// RandomHex returns a random css color as a hexadecimal value
func RandomHex() int32 {
	rand.Seed(time.Now().UnixNano())
	return HexColors[rand.Intn(len(HexColors))]
}

// RandomAnsi returns a random css color as an ANSI escape code
func RandomAnsi() string {
	rand.Seed(time.Now().UnixNano())
	return AnsiColors[rand.Intn(len(AnsiColors))]
}

// RandomAnsiBackground returns a random css color as an ANSI background escape code
func RandomAnsiBackground() string {
	rand.Seed(time.Now().UnixNano())
	return AnsiBackgrounds[rand.Intn(len(AnsiBackgrounds))]
}

// MultipleColor colorizes every rune in a string randomly from the AnsiColors
func MultipleColor(s string) string {
	var m string
	for _, r := range s {
		m += RandomAnsi() + string(r)
	}
	m += AnsiReset
	return m
}

// MultipleBackground colorizes every rune's background in a string randomly from the AnsiColors
func MultipleBackground(s string) string {
	var m string
	for _, r := range s {
		m += RandomAnsiBackground() + string(r)
	}
	m += AnsiReset
	return m
}

/* Indexes that match names to colors in different types */

// HexIndex is a map of all css colors matched with their corresponding
// decimal values
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

// AnsiIndex is a map that contains all of the css color names and their
// corresponding ansi escape codes for text color
var AnsiIndex = map[string]string{
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

// AnsiBackgroundIndex is a map of all css color names matched with their
// corresponding ansi escape codes for backgrounds
var AnsiBackgroundIndex = map[string]string{
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
