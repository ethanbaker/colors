package main

import (
        "fmt"
        "os"

        "gitlab.com/skilstak/code/go/color/css"
)

const usagetxt = `
usage: css COLOR message
-> Prints message in cooresponding color 

COLOR: 
AliceBlue, AntiqueWhite, Aqua, AquaMarine, Azure, Beige, Bisque, Black, BlanchedAlmond, Blue, BlueViolet, Brown, BurlyWood, CadetBlue, ChartReuse, Chocolate, Coral, CornFlowerBlue, CornSilk, Crimson, Cyan, DarkBlue, DarkCyan, DarkGoldenRod, DarkGray, DarkGreen, DarkKhaki, DarkMagenta, DarkOliveGreen, DarkOrange, DarkOrchid, DarkRed, DarkSalmon, DarkSeaGreen, DarkSlateBlue, DarkSlateGray, DarkTurquoise, DarkViolet, DeepPink, DeepSkyBlue, DimGray, DodgerBlue, FireBrick, FloralWhite, ForestGreen, Fuchsia, Gainsboro, GhostWhite, Gold, GoldenRod, Gray, Green, GreenYellow, HoneyDew, HotPink, IndianRed, Indigo, Ivory, Khaki, Lavender, LavenderBlush, LawnGreen, LemonChiffon, LightBlue, LightCoral, LightCyan, LightGoldenRodYellow, LightGreen, LightGrey, LightPink, LightSalmon, LightSeaGreen, LightSkyBlue, LightSlateGray, LightSteelBlue, LightYellow, Lime, LimeGreen, Linen, Magenta, Maroon, MediumAquaMarine, MediumBlue, MediumOrchid, MediumPurple, MediumSeaGreen, MediumSlateBlue, MediumSpringGreen, MediumTurquoise, MediumVioletRed, MidnightBlue, MintCream, MistyRose, Moccasin, NavajoWhite, Navy, NavyBlue, OldLace, Olive, OliveDrab, Orange, OrangeRed, Orchid, PaleGoldenRod, PaleGreen, PaleTurquoise, PaleVioletRed, PapayaWhip, PeachPuff, Peru, Pink, Plum, PowderBlue, Purple, Red, RosyBrown, RoyalBlue, SaddleBrown, Salmon, SandyBrown, SeaGreen, SeaShell, Sienna, Silver, SkyBlue, SlateBlue, SlateGray, Snow, SpringGreen, SteelBlue, Tan, Teal, Thistle, Tomato, Turquoise, Violet, Wheat, White, Whitesmoke, Yellow, YellowGreen

OR usage: css Sample
-> Shows all color names in the cooresponding color
`

func usage() {
        fmt.Println(usagetxt)
        os.Exit(1)
}

func main() {
        if len(os.Args) == 2 {
                if os.Args[1] != "Sample" {
                        usage()
                }
        } else if len(os.Args) == 1 {
                usage()
        }
        switch os.Args[1] {
        case "AliceBlue":
                fmt.Print(css.AliceBlue)
        case "AntiqueWhite":
                fmt.Print(css.AntiqueWhite)
        case "Aqua":
                fmt.Print(css.Aqua)
        case "AquaMarine":
                fmt.Print(css.AquaMarine)
        case "Azure":
                fmt.Print(css.Azure)
        case "Beige":
                fmt.Print(css.Beige)
        case "Bisque":
                fmt.Print(css.Bisque)
        case "Black":
                fmt.Print(css.Black)
        case "BlanchedAlmond":
                fmt.Print(css.BlanchedAlmond)
        case "Blue":
                fmt.Print(css.Blue)
        case "BlueViolet":
                fmt.Print(css.BlueViolet)
        case "Brown":
                fmt.Print(css.Brown)
        case "BurlyWood":
                fmt.Print(css.BurlyWood)
        case "CadetBlue":
                fmt.Print(css.CadetBlue)
        case "Chartreuse":
                fmt.Print(css.Chartreuse)
        case "Chocolate":
                fmt.Print(css.Chocolate)
        case "Coral":
                fmt.Print(css.Coral)
        case "CornflowerBlue":
                fmt.Print(css.CornFlowerBlue)
        case "CornSilk":
                fmt.Print(css.CornSilk)
        case "Crimson":
                fmt.Print(css.Crimson)
        case "Cyan":
                fmt.Print(css.Cyan)
        case "DarkBlue":
                fmt.Print(css.DarkBlue)
        case "DarkCyan":
                fmt.Print(css.DarkCyan)
        case "DarkGoldenRod":
                fmt.Print(css.DarkGoldenRod)
        case "DarkGray":
                fmt.Print(css.DarkGray)
        case "DarkGreen":
                fmt.Print(css.DarkGreen)
        case "DarkKhaki":
                fmt.Print(css.DarkKhaki)
        case "DarkMagenta":
                fmt.Print(css.DarkMagenta)
        case "DarkOliveGreen":
                fmt.Print(css.DarkOliveGreen)
        case "DarkOrange":
                fmt.Print(css.DarkOrange)
        case "DarkOrchid":
                fmt.Print(css.DarkOrchid)
        case "DarkRed":
                fmt.Print(css.DarkRed)
        case "DarkSalmon":
                fmt.Print(css.DarkSalmon)
        case "DarkSeaGreen":
                fmt.Print(css.DarkSeaGreen)
        case "DarkSlateBlue":
                fmt.Print(css.DarkSlateBlue)
        case "DarkSlateGray":
                fmt.Print(css.DarkSlateGray)
        case "DarkTurquoise":
                fmt.Print(css.DarkTurquoise)
        case "DarkViolet":
                fmt.Print(css.DarkViolet)
        case "DeepPink":
                fmt.Print(css.DeepPink)
        case "DeepSkyBlue":
                fmt.Print(css.DeepSkyBlue)
        case "DimGray":
                fmt.Print(css.DimGray)
        case "DodgerBlue":
                fmt.Print(css.DodgerBlue)
        case "Firebrick":
                fmt.Print(css.Firebrick)
        case "FloralWhite":
                fmt.Print(css.FloralWhite)
        case "ForestGreen":
                fmt.Print(css.ForestGreen)
        case "Fuchsia":
                fmt.Print(css.Fuchsia)
        case "Gainsboro":
                fmt.Print(css.Gainsboro)
        case "GhostWhite":
                fmt.Print(css.GhostWhite)
        case "Gold":
                fmt.Print(css.Gold)
        case "GoldenRod":
                fmt.Print(css.GoldenRod)
        case "Gray":
                fmt.Print(css.Gray)
        case "Green":
                fmt.Print(css.Green)
        case "GreenYellow":
                fmt.Print(css.GreenYellow)
        case "Honeydew":
                fmt.Print(css.Honeydew)
        case "HotPink":
                fmt.Print(css.HotPink)
        case "IndianRed":
                fmt.Print(css.IndianRed)
        case "Indigo":
                fmt.Print(css.Indigo)
        case "Ivory":
                fmt.Print(css.Ivory)
        case "Khaki":
                fmt.Print(css.Khaki)
        case "Lavender":
                fmt.Print(css.Lavender)
        case "LavenderBlush":
                fmt.Print(css.LavenderBlush)
        case "LawnGreen":
                fmt.Print(css.LawnGreen)
        case "LemonChiffon":
                fmt.Print(css.LemonChiffon)
        case "LightBlue":
                fmt.Print(css.LightBlue)
        case "LightCoral":
                fmt.Print(css.LightCoral)
        case "LightCyan":
                fmt.Print(css.LightCyan)
        case "LightGoldenRodYellow":
                fmt.Print(css.LightGoldenRodYellow)
        case "LightGreen":
                fmt.Print(css.LightGreen)
        case "LightGrey":
                fmt.Print(css.LightGrey)
        case "LightPink":
                fmt.Print(css.LightPink)
        case "LightSalmon":
                fmt.Print(css.LightSalmon)
        case "LightSeaGreen":
                fmt.Print(css.LightSeaGreen)
        case "LightSkyBlue":
                fmt.Print(css.LightSkyBlue)
        case "LightSlateGray":
                fmt.Print(css.LightSlateGray)
        case "LightSteelBlue":
                fmt.Print(css.LightSteelBlue)
        case "LightYellow":
                fmt.Print(css.LightYellow)
        case "Lime":
                fmt.Print(css.Lime)
        case "LimeGreen":
                fmt.Print(css.LimeGreen)
        case "Linen":
                fmt.Print(css.Linen)
        case "Magenta":
                fmt.Print(css.Magenta)
        case "Maroon":
                fmt.Print(css.Maroon)
        case "MediumAquaMarine":
                fmt.Print(css.MediumAquaMarine)
        case "MediumBlue":
                fmt.Print(css.MediumBlue)
        case "MediumOrchid":
                fmt.Print(css.MediumOrchid)
        case "MediumPurple":
                fmt.Print(css.MediumPurple)
        case "MediumSeaGreen":
                fmt.Print(css.MediumSeaGreen)
        case "MediumSlateBlue":
                fmt.Print(css.MediumSlateBlue)
        case "MediumSpringGreen":
                fmt.Print(css.MediumSpringGreen)
        case "MediumTurquoise":
                fmt.Print(css.MediumTurquoise)
        case "MediumVioletRed":
                fmt.Print(css.MediumVioletRed)
        case "MidnightBlue":
                fmt.Print(css.MidnightBlue)
        case "MintCream":
                fmt.Print(css.MintCream)
        case "MistyRose":
                fmt.Print(css.MistyRose)
        case "Moccasin":
                fmt.Print(css.Moccasin)
        case "NavajoWhite":
                fmt.Print(css.NavajoWhite)
        case "Navy":
                fmt.Print(css.Navy)
        case "NavyBlue":
                fmt.Print(css.NavyBlue)
        case "OldLace":
                fmt.Print(css.OldLace)
        case "Olive":
                fmt.Print(css.Olive)
        case "OliveDrab":
                fmt.Print(css.OliveDrab)
        case "Orange":
                fmt.Print(css.Orange)
        case "OrangeRed":
                fmt.Print(css.OrangeRed)
        case "Orchid":
                fmt.Print(css.Orchid)
        case "PaleGoldenRod":
                fmt.Print(css.PaleGoldenRod)
        case "PaleGreen":
                fmt.Print(css.PaleGreen)
        case "PaleTurquoise":
                fmt.Print(css.PaleTurquoise)
        case "PaleVioletRed":
                fmt.Print(css.PaleVioletRed)
        case "PapayaWhip":
                fmt.Print(css.PapayaWhip)
        case "PeachPuff":
                fmt.Print(css.PeachPuff)
        case "Peru":
                fmt.Print(css.Peru)
        case "Pink":
                fmt.Print(css.Pink)
        case "Plum":
                fmt.Print(css.Plum)
        case "PowderBlue":
                fmt.Print(css.PowderBlue)
        case "Purple":
                fmt.Print(css.Purple)
        case "Red":
                fmt.Print(css.Red)
        case "RosyBrown":
                fmt.Print(css.RosyBrown)
        case "RoyalBlue":
                fmt.Print(css.RoyalBlue)
        case "SaddleBrown":
                fmt.Print(css.SaddleBrown)
        case "Salmon":
                fmt.Print(css.Salmon)
        case "SandyBrown":
                fmt.Print(css.SandyBrown)
        case "SeaGreen":
                fmt.Print(css.SeaGreen)
        case "SeaShell":
                fmt.Print(css.SeaShell)
        case "Sienna":
                fmt.Print(css.Sienna)
        case "Silver":
                fmt.Print(css.Silver)
        case "SkyBlue":
                fmt.Print(css.SkyBlue)
        case "SlateBlue":
                fmt.Print(css.SlateBlue)
        case "SlateGray":
                fmt.Print(css.SlateGray)
        case "Snow":
                fmt.Print(css.Snow)
        case "SpringGreen":
                fmt.Print(css.SpringGreen)
        case "SteelBlue":
                fmt.Print(css.SteelBlue)
        case "Tan":
                fmt.Print(css.Tan)
        case "Teal":
                fmt.Print(css.Teal)
        case "Thistle":
                fmt.Print(css.Thistle)
        case "Tomato":
                fmt.Print(css.Tomato)
        case "Turquoise":
                fmt.Print(css.Turquoise)
        case "Violet":
                fmt.Print(css.Violet)
        case "Wheat":
                fmt.Print(css.Wheat)
        case "White":
                fmt.Print(css.White)
        case "WhiteSmoke":
                fmt.Print(css.WhiteSmoke)
        case "Yellow":
                fmt.Print(css.Yellow)
        case "YellowGreen":
                fmt.Print(css.YellowGreen)
        case "Sample":
                fmt.Printf("%v%v\n", css.AliceBlue, "aliceblue")
                fmt.Printf("%v%v\n", css.AntiqueWhite, "antiquewhite")
                fmt.Printf("%v%v\n", css.Aqua, "aqua")
                fmt.Printf("%v%v\n", css.AquaMarine, "aquamarine")
                fmt.Printf("%v%v\n", css.Azure, "azure")
                fmt.Printf("%v%v\n", css.Beige, "beige")
                fmt.Printf("%v%v\n", css.Bisque, "bisque")
                fmt.Printf("%v%v\n", css.Black, "black")
                fmt.Printf("%v%v\n", css.BlanchedAlmond, "blanchedalmond")
                fmt.Printf("%v%v\n", css.Blue, "blue")
                fmt.Printf("%v%v\n", css.BlueViolet, "blueviolet")
                fmt.Printf("%v%v\n", css.Brown, "brown")
                fmt.Printf("%v%v\n", css.BurlyWood, "burlywood")
                fmt.Printf("%v%v\n", css.CadetBlue, "cadetblue")
                fmt.Printf("%v%v\n", css.Chartreuse, "chartreuse")
                fmt.Printf("%v%v\n", css.Chocolate, "chocolate")
                fmt.Printf("%v%v\n", css.Coral, "coral")
                fmt.Printf("%v%v\n", css.CornFlowerBlue, "cornflowerblue")
                fmt.Printf("%v%v\n", css.CornSilk, "cornsilk")
                fmt.Printf("%v%v\n", css.Crimson, "crimson")
                fmt.Printf("%v%v\n", css.Cyan, "cyan")
                fmt.Printf("%v%v\n", css.DarkBlue, "darkblue")
                fmt.Printf("%v%v\n", css.DarkCyan, "darkcyan")
                fmt.Printf("%v%v\n", css.DarkGoldenRod, "darkgoldenrod")
                fmt.Printf("%v%v\n", css.DarkGray, "darkgray")
                fmt.Printf("%v%v\n", css.DarkGreen, "darkgreen")
                fmt.Printf("%v%v\n", css.DarkKhaki, "darkkhaki")
                fmt.Printf("%v%v\n", css.DarkMagenta, "darkmagenta")
                fmt.Printf("%v%v\n", css.DarkOliveGreen, "darkolivegreen")
                fmt.Printf("%v%v\n", css.DarkOrange, "darkorange")
                fmt.Printf("%v%v\n", css.DarkOrchid, "darkorchid")
                fmt.Printf("%v%v\n", css.DarkRed, "darkred")
                fmt.Printf("%v%v\n", css.DarkSalmon, "darksalmon")
                fmt.Printf("%v%v\n", css.DarkSeaGreen, "darkseagreen")
                fmt.Printf("%v%v\n", css.DarkSlateBlue, "darkslateblue")
                fmt.Printf("%v%v\n", css.DarkSlateGray, "darkslategray")
                fmt.Printf("%v%v\n", css.DarkTurquoise, "darkturquoise")
                fmt.Printf("%v%v\n", css.DarkViolet, "darkviolet")
                fmt.Printf("%v%v\n", css.DeepPink, "deeppink")
                fmt.Printf("%v%v\n", css.DeepSkyBlue, "deepskyblue")
                fmt.Printf("%v%v\n", css.DimGray, "dimgray")
                fmt.Printf("%v%v\n", css.DodgerBlue, "dodgerblue") //TODO
                fmt.Printf("%v%v\n", css.Firebrick, "firebrick")
                fmt.Printf("%v%v\n", css.FloralWhite, "floralwhite")
                fmt.Printf("%v%v\n", css.ForestGreen, "forestgreen")
                fmt.Printf("%v%v\n", css.Fuchsia, "fuchsia")
                fmt.Printf("%v%v\n", css.Gainsboro, "gainsboro")
                fmt.Printf("%v%v\n", css.GhostWhite, "ghostwhite")
                fmt.Printf("%v%v\n", css.Gold, "gold")
                fmt.Printf("%v%v\n", css.GoldenRod, "goldenrod")
                fmt.Printf("%v%v\n", css.Gray, "gray")
                fmt.Printf("%v%v\n", css.Green, "green")
                fmt.Printf("%v%v\n", css.GreenYellow, "greenyellow")
                fmt.Printf("%v%v\n", css.Honeydew, "honeydew")
                fmt.Printf("%v%v\n", css.HotPink, "hotpink")
                fmt.Printf("%v%v\n", css.IndianRed, "indianred")
                fmt.Printf("%v%v\n", css.Indigo, "indigo")
                fmt.Printf("%v%v\n", css.Ivory, "ivory")
                fmt.Printf("%v%v\n", css.Khaki, "khaki")
                fmt.Printf("%v%v\n", css.Lavender, "lavender")
                fmt.Printf("%v%v\n", css.LavenderBlush, "lavenderblush")
                fmt.Printf("%v%v\n", css.LawnGreen, "lawngreen")
                fmt.Printf("%v%v\n", css.LemonChiffon, "lemonchiffon")
                fmt.Printf("%v%v\n", css.LightBlue, "lightblue")
                fmt.Printf("%v%v\n", css.LightCoral, "lightcoral")
                fmt.Printf("%v%v\n", css.LightCyan, "lightcyan")
                fmt.Printf("%v%v\n", css.LightGoldenRodYellow, "lightgoldenrodyellow")
                fmt.Printf("%v%v\n", css.LightGreen, "lightgreen")
                fmt.Printf("%v%v\n", css.LightGrey, "lightgrey")
                fmt.Printf("%v%v\n", css.LightPink, "lightpink")
                fmt.Printf("%v%v\n", css.LightSalmon, "lightsalmon")
                fmt.Printf("%v%v\n", css.LightSeaGreen, "lightseagreen")
                fmt.Printf("%v%v\n", css.LightSkyBlue, "lightskyblue")
                fmt.Printf("%v%v\n", css.LightSlateGray, "lightslategray")
                fmt.Printf("%v%v\n", css.LightSteelBlue, "lightsteelblue")
                fmt.Printf("%v%v\n", css.LightYellow, "lightyellow")
                fmt.Printf("%v%v\n", css.Lime, "lime")
                fmt.Printf("%v%v\n", css.LimeGreen, "limegreen")
                fmt.Printf("%v%v\n", css.Linen, "linen")
                fmt.Printf("%v%v\n", css.Magenta, "magenta")
                fmt.Printf("%v%v\n", css.Maroon, "maroon")
                fmt.Printf("%v%v\n", css.MediumAquaMarine, "mediumaquamarine")
                fmt.Printf("%v%v\n", css.MediumBlue, "mediumblue")
                fmt.Printf("%v%v\n", css.MediumOrchid, "mediumorchid")
                fmt.Printf("%v%v\n", css.MediumPurple, "mediumpurple")
                fmt.Printf("%v%v\n", css.MediumSeaGreen, "mediumseagreen")
                fmt.Printf("%v%v\n", css.MediumSlateBlue, "mediumslateblue")
                fmt.Printf("%v%v\n", css.MediumSpringGreen, "mediumspringgreen")
                fmt.Printf("%v%v\n", css.MediumTurquoise, "mediumturquoise")
                fmt.Printf("%v%v\n", css.MediumVioletRed, "mediumvioletred")
                fmt.Printf("%v%v\n", css.MidnightBlue, "midnightblue")
                fmt.Printf("%v%v\n", css.MintCream, "mintcream")
                fmt.Printf("%v%v\n", css.MistyRose, "mistyrose")
                fmt.Printf("%v%v\n", css.Moccasin, "moccasin")
                fmt.Printf("%v%v\n", css.NavajoWhite, "navajowhite")
                fmt.Printf("%v%v\n", css.Navy, "navy")
                fmt.Printf("%v%v\n", css.NavyBlue, "navyblue")
                fmt.Printf("%v%v\n", css.OldLace, "oldlace")
                fmt.Printf("%v%v\n", css.Olive, "olive")
                fmt.Printf("%v%v\n", css.OliveDrab, "olivedrab")
                fmt.Printf("%v%v\n", css.Orange, "orange")
                fmt.Printf("%v%v\n", css.OrangeRed, "orangered")
                fmt.Printf("%v%v\n", css.Orchid, "orchid")
                fmt.Printf("%v%v\n", css.PaleGoldenRod, "palegoldenrod")
                fmt.Printf("%v%v\n", css.PaleGreen, "palegreen")
                fmt.Printf("%v%v\n", css.PaleTurquoise, "paleturquoise")
                fmt.Printf("%v%v\n", css.PaleVioletRed, "palevioletred")
                fmt.Printf("%v%v\n", css.PapayaWhip, "papayawhip")
                fmt.Printf("%v%v\n", css.PeachPuff, "peachpuff")
                fmt.Printf("%v%v\n", css.Peru, "peru")
                fmt.Printf("%v%v\n", css.Pink, "pink")
                fmt.Printf("%v%v\n", css.Plum, "plum")
                fmt.Printf("%v%v\n", css.PowderBlue, "powderblue")
                fmt.Printf("%v%v\n", css.Purple, "purple")
                fmt.Printf("%v%v\n", css.Red, "red")
                fmt.Printf("%v%v\n", css.RosyBrown, "rosybrown")
                fmt.Printf("%v%v\n", css.RoyalBlue, "royalblue")
                fmt.Printf("%v%v\n", css.SaddleBrown, "saddlebrown")
                fmt.Printf("%v%v\n", css.Salmon, "salmon")
                fmt.Printf("%v%v\n", css.SandyBrown, "sandybrown")
                fmt.Printf("%v%v\n", css.SeaGreen, "seagreen")
                fmt.Printf("%v%v\n", css.SeaShell, "seashell")
                fmt.Printf("%v%v\n", css.Sienna, "sienna")
                fmt.Printf("%v%v\n", css.Silver, "silver")
                fmt.Printf("%v%v\n", css.SkyBlue, "skyblue")
                fmt.Printf("%v%v\n", css.SlateBlue, "slateblue")
                fmt.Printf("%v%v\n", css.SlateGray, "slategray")
                fmt.Printf("%v%v\n", css.Snow, "snow")
                fmt.Printf("%v%v\n", css.SpringGreen, "springgreen")
                fmt.Printf("%v%v\n", css.SteelBlue, "steelblue")
                fmt.Printf("%v%v\n", css.Tan, "tan")
                fmt.Printf("%v%v\n", css.Teal, "teal")
                fmt.Printf("%v%v\n", css.Thistle, "thistle")
                fmt.Printf("%v%v\n", css.Tomato, "tomato")
                fmt.Printf("%v%v\n", css.Turquoise, "turquoise")
                fmt.Printf("%v%v\n", css.Violet, "violet")
                fmt.Printf("%v%v\n", css.Wheat, "wheat")
                fmt.Printf("%v%v\n", css.White, "white")
                fmt.Printf("%v%v\n", css.WhiteSmoke, "whitesmoke")
                fmt.Printf("%v%v\n", css.Yellow, "yellow")
                fmt.Printf("%v%v\n", css.YellowGreen, "yellowgreen")
        default:
                usage()
        }
        for i := 2; i < len(os.Args); i++ {
                fmt.Print(os.Args[i] + " ")
        }
        fmt.Println()
}
