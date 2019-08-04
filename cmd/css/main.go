package main

import (
        "fmt"
        "os"

        "gitlab.com/skilstak/code/go/color/css"
)

const usagetxt = `
usage: css COLOR

COLOR: 
aliceblue, antiquewhite, aqua, aquamarine, azure, beige, bisque, black, blanchedalmond, blue, blueviolet, brown, burlywood, cadetblue, chartreuse, chocolate, coral, cornflowerblue, cornsilk, crimson, cyan, darkblue, darkcyan, darkgoldenrod, darkgray, darkgreen, darkkhaki, darkmagenta, darkolivegreen, darkorange, darkorchid, darkred, darksalmon, darkseagreen, darkslateblue, darkslategray, darkturquoise, darkviolet, deeppink, deepskyblue, dimgray, dodgerblue, firebrick, floralwhite, forestgreen, fuchsia, gainsboro, ghostwhite, gold, goldenrod, gray, green, greenyellow, honeydew, hotpink, indianred, indigo, ivory, khaki, lavender, lavenderblush, lawngreen, lemonchiffon, lightblue, lightcoral, lightcyan, lightgoldenrodyellow, lightgreen, lightgrey, lightpink, lightsalmon, lightseagreen, lightskyblue, lightslategray, lightsteelblue, lightyellow, lime, limegreen, linen, magenta, maroon, mediumaquamarine, mediumblue, mediumorchid, mediumpurple, mediumseagreen, mediumslateblue, mediumspringgreen, mediumturquoise, mediumvioletred, midnightblue, mintcream, mistyrose, moccasin, navajowhite, navy, navyblue, oldlace, olive, olivedrab, orange, orangered, orchid, palegoldenrod, palegreen, paleturquoise, palevioletred, papayawhip, peachpuff, peru, pink, plum, powderblue, purple, red, rosybrown, royalblue, saddlebrown, salmon, sandybrown, seagreen, seashell, sienna, silver, skyblue, slateblue, slategray, snow, springgreen, steelblue, tan, teal, thistle, tomato, turquoise, violet, wheat, white, whitesmoke, yellow, yellowgreen
`

func usage() {
        fmt.Println(usagetxt)
        os.Exit(1)
}

func main() {
        if len(os.Args) != 2 {
                usage()
        }
        switch os.Args[1] {
        case "AliceBlue":
                fmt.Println(css.AliceBlue)
        case "AntiqueWhite":
                fmt.Println(css.AntiqueWhite)
        case "Aqua":
                fmt.Println(css.Aqua)
        case "AquaMarine":
                fmt.Println(css.AquaMarine)
        case "Azure":
                fmt.Println(css.Azure)
        case "Beige":
                fmt.Println(css.Beige)
        case "Bisque":
                fmt.Println(css.Bisque)
        case "Black":
                fmt.Println(css.Black)
        case "BlanchedAlmond":
                fmt.Println(css.BlanchedAlmond)
        case "Blue":
                fmt.Println(css.Blue)
        case "BlueViolet":
                fmt.Println(css.BlueViolet)
        case "Brown":
                fmt.Println(css.Brown)
        case "BurlyWood":
                fmt.Println(css.BurlyWood)
        case "CadetBlue":
                fmt.Println(css.CadetBlue)
        case "Chartreuse":
                fmt.Println(css.Chartreuse)
        case "Chocolate":
                fmt.Println(css.Chocolate)
        case "Coral":
                fmt.Println(css.Coral)
        case "CornflowerBlue":
                fmt.Println(css.CornFlowerBlue)
        case "CornSilk":
                fmt.Println(css.CornSilk)
        case "Crimson":
                fmt.Println(css.Crimson)
        case "Cyan":
                fmt.Println(css.Cyan)
        case "DarkBlue":
                fmt.Println(css.DarkBlue)
        case "DarkCyan":
                fmt.Println(css.DarkCyan)
        case "DarkGoldenRod":
                fmt.Println(css.DarkGoldenRod)
        case "DarkGray":
                fmt.Println(css.DarkGray)
        case "DarkGreen":
                fmt.Println(css.DarkGreen)
        case "DarkKhaki":
                fmt.Println(css.DarkKhaki)
        case "DarkMagenta":
                fmt.Println(css.DarkMagenta)
        case "DarkOliveGreen":
                fmt.Println(css.DarkOliveGreen)
        case "DarkOrange":
                fmt.Println(css.DarkOrange)
        case "DarkOrchid":
                fmt.Println(css.DarkOrchid)
        case "DarkRed":
                fmt.Println(css.DarkRed)
        case "DarkSalmon":
                fmt.Println(css.DarkSalmon)
        case "DarkSeaGreen":
                fmt.Println(css.DarkSeaGreen)
        case "DarkSlateBlue":
                fmt.Println(css.DarkSlateBlue)
        case "DarkSlateGray":
                fmt.Println(css.DarkSlateGray)
        case "DarkTurquoise":
                fmt.Println(css.DarkTurquoise)
        case "DarkViolet":
                fmt.Println(css.DarkViolet)
        case "DeepPink":
                fmt.Println(css.DeepPink)
        case "DeepSkyBlue":
                fmt.Println(css.DeepSkyBlue)
        case "DimGray":
                fmt.Println(css.DimGray)
        case "DodgerBlue":
                fmt.Println(css.DodgerBlue)
        case "Firebrick":
                fmt.Println(css.Firebrick)
        case "FloralWhite":
                fmt.Println(css.FloralWhite)
        case "ForestGreen":
                fmt.Println(css.ForestGreen)
        case "Fuchsia":
                fmt.Println(css.Fuchsia)
        case "Gainsboro":
                fmt.Println(css.Gainsboro)
        case "GhostWhite":
                fmt.Println(css.GhostWhite)
        case "Gold":
                fmt.Println(css.Gold)
        case "GoldenRod":
                fmt.Println(css.GoldenRod)
        case "Gray":
                fmt.Println(css.Gray)
        case "Green":
                fmt.Println(css.Green)
        case "GreenYellow":
                fmt.Println(css.GreenYellow)
        case "Honeydew":
                fmt.Println(css.Honeydew)
        case "HotPink":
                fmt.Println(css.HotPink)
        case "IndianRed":
                fmt.Println(css.IndianRed)
        case "Indigo":
                fmt.Println(css.Indigo)
        case "Ivory":
                fmt.Println(css.Ivory)
        case "Khaki":
                fmt.Println(css.Khaki)
        case "Lavender":
                fmt.Println(css.Lavender)
        case "LavenderBlush":
                fmt.Println(css.LavenderBlush)
        case "LawnGreen":
                fmt.Println(css.LawnGreen)
        case "LemonChiffon":
                fmt.Println(css.LemonChiffon)
        case "LightBlue":
                fmt.Println(css.LightBlue)
        case "LightCoral":
                fmt.Println(css.LightCoral)
        case "LightCyan":
                fmt.Println(css.LightCyan)
        case "LightGoldenRodYellow":
                fmt.Println(css.LightGoldenRodYellow)
        case "LightGreen":
                fmt.Println(css.LightGreen)
        case "LightGrey":
                fmt.Println(css.LightGrey)
        case "LightPink":
                fmt.Println(css.LightPink)
        case "LightSalmon":
                fmt.Println(css.LightSalmon)
        case "LightSeaGreen":
                fmt.Println(css.LightSeaGreen)
        case "LightSkyBlue":
                fmt.Println(css.LightSkyBlue)
        case "LightSlateGray":
                fmt.Println(css.LightSlateGray)
        case "LightSteelBlue":
                fmt.Println(css.LightSteelBlue)
        case "LightYellow":
                fmt.Println(css.LightYellow)
        case "Lime":
                fmt.Println(css.Lime)
        case "LimeGreen":
                fmt.Println(css.LimeGreen)
        case "Linen":
                fmt.Println(css.Linen)
        case "Magenta":
                fmt.Println(css.Magenta)
        case "Maroon":
                fmt.Println(css.Maroon)
        case "MediumAquamarine":
                fmt.Println(css.MediumAquaMarine)
        case "MediumBlue":
                fmt.Println(css.MediumBlue)
        case "MediumOrchid":
                fmt.Println(css.MediumOrchid)
        case "MediumPurple":
                fmt.Println(css.MediumPurple)
        case "MediumSeaGreen":
                fmt.Println(css.MediumSeaGreen)
        case "MediumSlateBlue":
                fmt.Println(css.MediumSlateBlue)
        case "MediumSpringGreen":
                fmt.Println(css.MediumSpringGreen)
        case "MediumTurquoise":
                fmt.Println(css.MediumTurquoise)
        case "MediumVioletRed":
                fmt.Println(css.MediumVioletRed)
        case "MidnightBlue":
                fmt.Println(css.MidnightBlue)
        case "MintCream":
                fmt.Println(css.MintCream)
        case "MistyRose":
                fmt.Println(css.MistyRose)
        case "Moccasin":
                fmt.Println(css.Moccasin)
        case "NavajoWhite":
                fmt.Println(css.NavajoWhite)
        case "Navy":
                fmt.Println(css.Navy)
        case "NavyBlue":
                fmt.Println(css.NavyBlue)
        case "OldLace":
                fmt.Println(css.OldLace)
        case "Olive":
                fmt.Println(css.Olive)
        case "OliveDrab":
                fmt.Println(css.OliveDrab)
        case "Orange":
                fmt.Println(css.Orange)
        case "OrangeRed":
                fmt.Println(css.OrangeRed)
        case "Orchid":
                fmt.Println(css.Orchid)
        case "PaleGoldenRod":
                fmt.Println(css.PaleGoldenRod)
        case "PaleGreen":
                fmt.Println(css.PaleGreen)
        case "PaleTurquoise":
                fmt.Println(css.PaleTurquoise)
        case "PaleVioletRed":
                fmt.Println(css.PaleVioletRed)
        case "PapayaWhip":
                fmt.Println(css.PapayaWhip)
        case "PeachPuff":
                fmt.Println(css.PeachPuff)
        case "Peru":
                fmt.Println(css.Peru)
        case "Pink":
                fmt.Println(css.Pink)
        case "Plum":
                fmt.Println(css.Plum)
        case "PowderBlue":
                fmt.Println(css.PowderBlue)
        case "Purple":
                fmt.Println(css.Purple)
        case "Red":
                fmt.Println(css.Red)
        case "RosyBrown":
                fmt.Println(css.RosyBrown)
        case "RoyalBlue":
                fmt.Println(css.RoyalBlue)
        case "SaddleBrown":
                fmt.Println(css.SaddleBrown)
        case "Salmon":
                fmt.Println(css.Salmon)
        case "SandyBrown":
                fmt.Println(css.SandyBrown)
        case "SeaGreen":
                fmt.Println(css.SeaGreen)
        case "SeaShell":
                fmt.Println(css.SeaShell)
        case "Sienna":
                fmt.Println(css.Sienna)
        case "Silver":
                fmt.Println(css.Silver)
        case "SkyBlue":
                fmt.Println(css.SkyBlue)
        case "SlateBlue":
                fmt.Println(css.SlateBlue)
        case "SlateGray":
                fmt.Println(css.SlateGray)
        case "Snow":
                fmt.Println(css.Snow)
        case "SpringGreen":
                fmt.Println(css.SpringGreen)
        case "SteelBlue":
                fmt.Println(css.SteelBlue)
        case "Tan":
                fmt.Println(css.Tan)
        case "Teal":
                fmt.Println(css.Teal)
        case "Thistle":
                fmt.Println(css.Thistle)
        case "Tomato":
                fmt.Println(css.Tomato)
        case "Turquoise":
                fmt.Println(css.Turquoise)
        case "Violet":
                fmt.Println(css.Violet)
        case "Wheat":
                fmt.Println(css.Wheat)
        case "White":
                fmt.Println(css.White)
        case "WhiteSmoke":
                fmt.Println(css.WhiteSmoke)
        case "Yellow":
                fmt.Println(css.Yellow)
        case "Yellowgreen":
                fmt.Println(css.YellowGreen)
        case "Sample":
                fmt.Printf("%v%v", css.AliceBlue, "aliceblue")
                fmt.Printf("%v%v", css.AntiqueWhite, "antiquewhite")
                fmt.Printf("%v%v", css.Aqua, "aqua")
                fmt.Printf("%v%v", css.AquaMarine, "aquamarine")
                fmt.Printf("%v%v", css.Azure, "azure")
                fmt.Printf("%v%v", css.Beige, "beige")
                fmt.Printf("%v%v", css.Bisque, "bisque")
                fmt.Printf("%v%v", css.Black, "black")
                fmt.Printf("%v%v", css.BlanchedAlmond, "blanchedalmond")
                fmt.Printf("%v%v", css.Blue, "blue")
                fmt.Printf("%v%v", css.BlueViolet, "blueviolet")
                fmt.Printf("%v%v", css.Brown, "brown")
                fmt.Printf("%v%v", css.BurlyWood, "burlywood")
                fmt.Printf("%v%v", css.CadetBlue, "cadetblue")
                fmt.Printf("%v%v", css.Chartreuse, "chartreuse")
                fmt.Printf("%v%v", css.Chocolate, "chocolate")
                fmt.Printf("%v%v", css.Coral, "coral")
                fmt.Printf("%v%v", css.CornFlowerBlue, "cornflowerblue")
                fmt.Printf("%v%v", css.CornSilk, "cornsilk")
                fmt.Printf("%v%v", css.Crimson, "crimson")
                fmt.Printf("%v%v", css.Cyan, "cyan")
                fmt.Printf("%v%v", css.DarkBlue, "darkblue")
                fmt.Printf("%v%v", css.DarkCyan, "darkcyan")
                fmt.Printf("%v%v", css.DarkGoldenRod, "darkgoldenrod")
                fmt.Printf("%v%v", css.DarkGray, "darkgray")
                fmt.Printf("%v%v", css.DarkGreen, "darkgreen")
                fmt.Printf("%v%v", css.DarkKhaki, "darkkhaki")
                fmt.Printf("%v%v", css.DarkMagenta, "darkmagenta")
                fmt.Printf("%v%v", css.DarkOliveGreen, "darkolivegreen")
                fmt.Printf("%v%v", css.DarkOrange, "darkorange")
                fmt.Printf("%v%v", css.DarkOrchid, "darkorchid")
                fmt.Printf("%v%v", css.DarkRed, "darkred")
                fmt.Printf("%v%v", css.DarkSalmon, "darksalmon")
                fmt.Printf("%v%v", css.DarkSeaGreen, "darkseagreen")
                fmt.Printf("%v%v", css.DarkSlateBlue, "darkslateblue")
                fmt.Printf("%v%v", css.DarkSlateGray, "darkslategray")
                fmt.Printf("%v%v", css.DarkTurquoise, "darkturquoise")
                fmt.Printf("%v%v", css.DarkViolet, "darkviolet")
                fmt.Printf("%v%v", css.DeepPink, "deeppink")
                fmt.Printf("%v%v", css.DeepSkyBlue, "deepskyblue")
                fmt.Printf("%v%v", css.DimGray, "dimgray")
                fmt.Printf("%v%v", css.DodgerBlue, "dodgerblue") //TODO
                fmt.Printf("%v%v", css.Firebrick, "firebrick")
                fmt.Printf("%v%v", css.FloralWhite, "floralwhite")
                fmt.Printf("%v%v", css.ForestGreen, "forestgreen")
                fmt.Printf("%v%v", css.Fuchsia, "fuchsia")
                fmt.Printf("%v%v", css.Gainsboro, "gainsboro")
                fmt.Printf("%v%v", css.GhostWhite, "ghostwhite")
                fmt.Printf("%v%v", css.Gold, "gold")
                fmt.Printf("%v%v", css.GoldenRod, "goldenrod")
                fmt.Printf("%v%v", css.Gray, "gray")
                fmt.Printf("%v%v", css.Green, "green")
                fmt.Printf("%v%v", css.GreenYellow, "greenyellow")
                fmt.Printf("%v%v", css.Honeydew, "honeydew")
                fmt.Printf("%v%v", css.HotPink, "hotpink")
                fmt.Printf("%v%v", css.IndianRed, "indianred")
                fmt.Printf("%v%v", css.Indigo, "indigo")
                fmt.Printf("%v%v", css.Ivory, "ivory")
                fmt.Printf("%v%v", css.Khaki, "khaki")
                fmt.Printf("%v%v", css.Lavender, "lavender")
                fmt.Printf("%v%v", css.LavenderBlush, "lavenderblush")
                fmt.Printf("%v%v", css.LawnGreen, "lawngreen")
                fmt.Printf("%v%v", css.LemonChiffon, "lemonchiffon")
                fmt.Printf("%v%v", css.LightBlue, "lightblue")
                fmt.Printf("%v%v", css.LightCoral, "lightcoral")
                fmt.Printf("%v%v", css.LightCyan, "lightcyan")
                fmt.Printf("%v%v", css.LightGoldenRodYellow, "lightgoldenrodyellow")
                fmt.Printf("%v%v", css.LightGreen, "lightgreen")
                fmt.Printf("%v%v", css.LightGrey, "lightgrey")
                fmt.Printf("%v%v", css.LightPink, "lightpink")
                fmt.Printf("%v%v", css.LightSalmon, "lightsalmon")
                fmt.Printf("%v%v", css.LightSeaGreen, "lightseagreen")
                fmt.Printf("%v%v", css.LightSkyBlue, "lightskyblue")
                fmt.Printf("%v%v", css.LightSlateGray, "lightslategray")
                fmt.Printf("%v%v", css.LightSteelBlue, "lightsteelblue")
                fmt.Printf("%v%v", css.LightYellow, "lightyellow")
                fmt.Printf("%v%v", css.Lime, "lime")
                fmt.Printf("%v%v", css.LimeGreen, "limegreen")
                fmt.Printf("%v%v", css.Linen, "linen")
                fmt.Printf("%v%v", css.Magenta, "magenta")
                fmt.Printf("%v%v", css.Maroon, "maroon")
                fmt.Printf("%v%v", css.MediumAquaMarine, "mediumaquamarine")
                fmt.Printf("%v%v", css.MediumBlue, "mediumblue")
                fmt.Printf("%v%v", css.MediumOrchid, "mediumorchid")
                fmt.Printf("%v%v", css.MediumPurple, "mediumpurple")
                fmt.Printf("%v%v", css.MediumSeaGreen, "mediumseagreen")
                fmt.Printf("%v%v", css.MediumSlateBlue, "mediumslateblue")
                fmt.Printf("%v%v", css.MediumSpringGreen, "mediumspringgreen")
                fmt.Printf("%v%v", css.MediumTurquoise, "mediumturquoise")
                fmt.Printf("%v%v", css.MediumVioletRed, "mediumvioletred")
                fmt.Printf("%v%v", css.MidnightBlue, "midnightblue")
                fmt.Printf("%v%v", css.MintCream, "mintcream")
                fmt.Printf("%v%v", css.MistyRose, "mistyrose")
                fmt.Printf("%v%v", css.Moccasin, "moccasin")
                fmt.Printf("%v%v", css.NavajoWhite, "navajowhite")
                fmt.Printf("%v%v", css.Navy, "navy")
                fmt.Printf("%v%v", css.NavyBlue, "navyblue")
                fmt.Printf("%v%v", css.OldLace, "oldlace")
                fmt.Printf("%v%v", css.Olive, "olive")
                fmt.Printf("%v%v", css.OliveDrab, "olivedrab")
                fmt.Printf("%v%v", css.Orange, "orange")
                fmt.Printf("%v%v", css.OrangeRed, "orangered")
                fmt.Printf("%v%v", css.Orchid, "orchid")
                fmt.Printf("%v%v", css.PaleGoldenRod, "palegoldenrod")
                fmt.Printf("%v%v", css.PaleGreen, "palegreen")
                fmt.Printf("%v%v", css.PaleTurquoise, "paleturquoise")
                fmt.Printf("%v%v", css.PaleVioletRed, "palevioletred")
                fmt.Printf("%v%v", css.PapayaWhip, "papayawhip")
                fmt.Printf("%v%v", css.PeachPuff, "peachpuff")
                fmt.Printf("%v%v", css.Peru, "peru")
                fmt.Printf("%v%v", css.Pink, "pink")
                fmt.Printf("%v%v", css.Plum, "plum")
                fmt.Printf("%v%v", css.PowderBlue, "powderblue")
                fmt.Printf("%v%v", css.Purple, "purple")
                fmt.Printf("%v%v", css.Red, "red")
                fmt.Printf("%v%v", css.RosyBrown, "rosybrown")
                fmt.Printf("%v%v", css.RoyalBlue, "royalblue")
                fmt.Printf("%v%v", css.SaddleBrown, "saddlebrown")
                fmt.Printf("%v%v", css.Salmon, "salmon")
                fmt.Printf("%v%v", css.SandyBrown, "sandybrown")
                fmt.Printf("%v%v", css.SeaGreen, "seagreen")
                fmt.Printf("%v%v", css.SeaShell, "seashell")
                fmt.Printf("%v%v", css.Sienna, "sienna")
                fmt.Printf("%v%v", css.Silver, "silver")
                fmt.Printf("%v%v", css.SkyBlue, "skyblue")
                fmt.Printf("%v%v", css.SlateBlue, "slateblue")
                fmt.Printf("%v%v", css.SlateGray, "slategray")
                fmt.Printf("%v%v", css.Snow, "snow")
                fmt.Printf("%v%v", css.SpringGreen, "springgreen")
                fmt.Printf("%v%v", css.SteelBlue, "steelblue")
                fmt.Printf("%v%v", css.Tan, "tan")
                fmt.Printf("%v%v", css.Teal, "teal")
                fmt.Printf("%v%v", css.Thistle, "thistle")
                fmt.Printf("%v%v", css.Tomato, "tomato")
                fmt.Printf("%v%v", css.Turquoise, "turquoise")
                fmt.Printf("%v%v", css.Violet, "violet")
                fmt.Printf("%v%v", css.Wheat, "wheat")
                fmt.Printf("%v%v", css.White, "white")
                fmt.Printf("%v%v", css.WhiteSmoke, "whitesmoke")
                fmt.Printf("%v%v", css.Yellow, "yellow")
                fmt.Printf("%v%v", css.YellowGreen, "yellowgreen")
        default:
                usage()
        }
}
