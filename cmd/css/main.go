package main

import (
        "fmt"
        "os"

        "gitlab.com/ethanbaker.dev/color/css"
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
        case "Aquamarine":
                fmt.Println(css.Aquamarine)
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
                fmt.Println(css.Ghostwhite)
        case "Gold":
                fmt.Println(css.Gold)
        case "GoldenRod":
                fmt.Println(css.Goldenrod)
        case "Gray":
                fmt.Println(css.Gray)
        case "Green":
                fmt.Println(css.Green)
        case "GreenYellow":
                fmt.Println(css.GreenYellow)
        case "HoneyDew":
                fmt.Println(css.HoneyDew)
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
        case "Limegreen":
                fmt.Println(css.Limegreen)
        case "Linen":
                fmt.Println(css.Linen)
        case "Magenta":
                fmt.Println(css.Magenta)
        case "Maroon":
                fmt.Println(css.Maroon)
        case "MediumAquamarine":
                fmt.Println(css.MediumAquamarine)
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
                fmt.Println(css.Navyblue)
        case "OldLace":
                fmt.Println(css.Oldlace)
        case "Olive":
                fmt.Println(css.Olive)
        case "Olivedrab":
                fmt.Println(css.Olivedrab)
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
        case "Whitesmoke":
                fmt.Println(css.Whitesmoke)
        case "Yellow":
                fmt.Println(css.Yellow)
        case "Yellowgreen":
                fmt.Println(css.YellowGreen)
        case "Sample":
                fmt.Printf("%v%v", css.Aliceblue, "aliceblue")
                fmt.Printf("%v%v", css.Antiquewhite, "antiquewhite")
                fmt.Printf("%v%v", css.Aqua, "aqua")
                fmt.Printf("%v%v", css.Aquamarine, "aquamarine")
                fmt.Printf("%v%v", css.Azure, "azure")
                fmt.Printf("%v%v", css.Beige, "beige")
                fmt.Printf("%v%v", css.Bisque, "bisque")
                fmt.Printf("%v%v", css.Black, "black")
                fmt.Printf("%v%v", css.Blanchedalmond, "blanchedalmond")
                fmt.Printf("%v%v", css.Blue, "blue")
                fmt.Printf("%v%v", css.Blueviolet, "blueviolet")
                fmt.Printf("%v%v", css.Brown, "brown")
                fmt.Printf("%v%v", css.Burlywood, "burlywood")
                fmt.Printf("%v%v", css.Cadetblue, "cadetblue")
                fmt.Printf("%v%v", css.Chartreuse, "chartreuse")
                fmt.Printf("%v%v", css.Chocolate, "chocolate")
                fmt.Printf("%v%v", css.Coral, "coral")
                fmt.Printf("%v%v", css.Cornflowerblue, "cornflowerblue")
                fmt.Printf("%v%v", css.Cornsilk, "cornsilk")
                fmt.Printf("%v%v", css.Crimson, "crimson")
                fmt.Printf("%v%v", css.Cyan, "cyan")
                fmt.Printf("%v%v", css.Darkblue, "darkblue")
                fmt.Printf("%v%v", css.Darkcyan, "darkcyan")
                fmt.Printf("%v%v", css.Darkgoldenrod, "darkgoldenrod")
                fmt.Printf("%v%v", css.Darkgray, "darkgray")
                fmt.Printf("%v%v", css.Darkgreen, "darkgreen")
                fmt.Printf("%v%v", css.Darkkhaki, "darkkhaki")
                fmt.Printf("%v%v", css.Darkmagenta, "darkmagenta")
                fmt.Printf("%v%v", css.Darkolivegreen, "darkolivegreen")
                fmt.Printf("%v%v", css.Darkorange, "darkorange")
                fmt.Printf("%v%v", css.Darkorchid, "darkorchid")
                fmt.Printf("%v%v", css.Darkred, "darkred")
                fmt.Printf("%v%v", css.Darksalmon, "darksalmon")
                fmt.Printf("%v%v", css.Darkseagreen, "darkseagreen")
                fmt.Printf("%v%v", css.Darkslateblue, "darkslateblue")
                fmt.Printf("%v%v", css.Darkslategray, "darkslategray")
                fmt.Printf("%v%v", css.Darkturquoise, "darkturquoise")
                fmt.Printf("%v%v", css.Darkviolet, "darkviolet")
                fmt.Printf("%v%v", css.Deeppink, "deeppink")
                fmt.Printf("%v%v", css.Deepskyblue, "deepskyblue")
                fmt.Printf("%v%v", css.Dimgray, "dimgray")
                fmt.Printf("%v%v", css.dodgerblue, "dodgerblue") //TODO
                fmt.Printf("%v%v", css.firebrick, "firebrick")
                fmt.Printf("%v%v", css.floralwhite, "floralwhite")
                fmt.Printf("%v%v", css.forestgreen, "forestgreen")
                fmt.Printf("%v%v", css.fuchsia, "fuchsia")
                fmt.Printf("%v%v", css.gainsboro, "gainsboro")
                fmt.Printf("%v%v", css.ghostwhite, "ghostwhite")
                fmt.Printf("%v%v", css.gold, "gold")
                fmt.Printf("%v%v", css.goldenrod, "goldenrod")
                fmt.Printf("%v%v", css.gray, "gray")
                fmt.Printf("%v%v", css.green, "green")
                fmt.Printf("%v%v", css.greenyellow, "greenyellow")
                fmt.Printf("%v%v", css.honeydew, "honeydew")
                fmt.Printf("%v%v", css.hotpink, "hotpink")
                fmt.Printf("%v%v", css.indianred, "indianred")
                fmt.Printf("%v%v", css.indigo, "indigo")
                fmt.Printf("%v%v", css.ivory, "ivory")
                fmt.Printf("%v%v", css.khaki, "khaki")
                fmt.Printf("%v%v", css.lavender, "lavender")
                fmt.Printf("%v%v", css.lavenderblush, "lavenderblush")
                fmt.Printf("%v%v", css.lawngreen, "lawngreen")
                fmt.Printf("%v%v", css.lemonchiffon, "lemonchiffon")
                fmt.Printf("%v%v", css.lightblue, "lightblue")
                fmt.Printf("%v%v", css.lightcoral, "lightcoral")
                fmt.Printf("%v%v", css.lightcyan, "lightcyan")
                fmt.Printf("%v%v", css.lightgoldenrodyellow, "lightgoldenrodyellow")
                fmt.Printf("%v%v", css.lightgreen, "lightgreen")
                fmt.Printf("%v%v", css.lightgrey, "lightgrey")
                fmt.Printf("%v%v", css.lightpink, "lightpink")
                fmt.Printf("%v%v", css.lightsalmon, "lightsalmon")
                fmt.Printf("%v%v", css.lightseagreen, "lightseagreen")
                fmt.Printf("%v%v", css.lightskyblue, "lightskyblue")
                fmt.Printf("%v%v", css.lightslategray, "lightslategray")
                fmt.Printf("%v%v", css.lightsteelblue, "lightsteelblue")
                fmt.Printf("%v%v", css.lightyellow, "lightyellow")
                fmt.Printf("%v%v", css.lime, "lime")
                fmt.Printf("%v%v", css.limegreen, "limegreen")
                fmt.Printf("%v%v", css.linen, "linen")
                fmt.Printf("%v%v", css.magenta, "magenta")
                fmt.Printf("%v%v", css.maroon, "maroon")
                fmt.Printf("%v%v", css.mediumaquamarine, "mediumaquamarine")
                fmt.Printf("%v%v", css.mediumblue, "mediumblue")
                fmt.Printf("%v%v", css.mediumorchid, "mediumorchid")
                fmt.Printf("%v%v", css.mediumpurple, "mediumpurple")
                fmt.Printf("%v%v", css.mediumseagreen, "mediumseagreen")
                fmt.Printf("%v%v", css.mediumslateblue, "mediumslateblue")
                fmt.Printf("%v%v", css.mediumspringgreen, "mediumspringgreen")
                fmt.Printf("%v%v", css.mediumturquoise, "mediumturquoise")
                fmt.Printf("%v%v", css.mediumvioletred, "mediumvioletred")
                fmt.Printf("%v%v", css.midnightblue, "midnightblue")
                fmt.Printf("%v%v", css.mintcream, "mintcream")
                fmt.Printf("%v%v", css.mistyrose, "mistyrose")
                fmt.Printf("%v%v", css.moccasin, "moccasin")
                fmt.Printf("%v%v", css.navajowhite, "navajowhite")
                fmt.Printf("%v%v", css.navy, "navy")
                fmt.Printf("%v%v", css.navyblue, "navyblue")
                fmt.Printf("%v%v", css.oldlace, "oldlace")
                fmt.Printf("%v%v", css.olive, "olive")
                fmt.Printf("%v%v", css.olivedrab, "olivedrab")
                fmt.Printf("%v%v", css.orange, "orange")
                fmt.Printf("%v%v", css.orangered, "orangered")
                fmt.Printf("%v%v", css.orchid, "orchid")
                fmt.Printf("%v%v", css.palegoldenrod, "palegoldenrod")
                fmt.Printf("%v%v", css.palegreen, "palegreen")
                fmt.Printf("%v%v", css.paleturquoise, "paleturquoise")
                fmt.Printf("%v%v", css.palevioletred, "palevioletred")
                fmt.Printf("%v%v", css.papayawhip, "papayawhip")
                fmt.Printf("%v%v", css.peachpuff, "peachpuff")
                fmt.Printf("%v%v", css.peru, "peru")
                fmt.Printf("%v%v", css.pink, "pink")
                fmt.Printf("%v%v", css.plum, "plum")
                fmt.Printf("%v%v", css.powderblue, "powderblue")
                fmt.Printf("%v%v", css.purple, "purple")
                fmt.Printf("%v%v", css.red, "red")
                fmt.Printf("%v%v", css.rosybrown, "rosybrown")
                fmt.Printf("%v%v", css.royalblue, "royalblue")
                fmt.Printf("%v%v", css.saddlebrown, "saddlebrown")
                fmt.Printf("%v%v", css.salmon, "salmon")
                fmt.Printf("%v%v", css.sandybrown, "sandybrown")
                fmt.Printf("%v%v", css.seagreen, "seagreen")
                fmt.Printf("%v%v", css.seashell, "seashell")
                fmt.Printf("%v%v", css.sienna, "sienna")
                fmt.Printf("%v%v", css.silver, "silver")
                fmt.Printf("%v%v", css.skyblue, "skyblue")
                fmt.Printf("%v%v", css.slateblue, "slateblue")
                fmt.Printf("%v%v", css.slategray, "slategray")
                fmt.Printf("%v%v", css.snow, "snow")
                fmt.Printf("%v%v", css.springgreen, "springgreen")
                fmt.Printf("%v%v", css.steelblue, "steelblue")
                fmt.Printf("%v%v", css.tan, "tan")
                fmt.Printf("%v%v", css.teal, "teal")
                fmt.Printf("%v%v", css.thistle, "thistle")
                fmt.Printf("%v%v", css.tomato, "tomato")
                fmt.Printf("%v%v", css.turquoise, "turquoise")
                fmt.Printf("%v%v", css.violet, "violet")
                fmt.Printf("%v%v", css.wheat, "wheat")
                fmt.Printf("%v%v", css.white, "white")
                fmt.Printf("%v%v", css.whitesmoke, "whitesmoke")
                fmt.Printf("%v%v", css.yellow, "yellow")
                fmt.Printf("%v%v", css.yellowgreen, "yellowgreen")
        default:
                usage()
        }
}
