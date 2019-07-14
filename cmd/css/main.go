package main

import (
        "fmt"
        "os"

        "gitlab.com/ethandbaker/color/css"
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
        case "aliceblue":
                fmt.Println(css.aliceblue)
        case "antiquewhite":
                fmt.Println(css.antiquewhite)
        case "aqua":
                fmt.Println(css.aqua)
        case "aquamarine":
                fmt.Println(css.aquamarine)
        case "azure":
                fmt.Println(css.azure)
        case "beige":
                fmt.Println(css.beige)
        case "bisque":
                fmt.Println(css.bisque)
        case "black":
                fmt.Println(css.black)
        case "blanchedalmond":
                fmt.Println(css.blanchedalmond)
        case "blue":
                fmt.Println(css.blue)
        case "blueviolet":
                fmt.Println(css.blueviolet)
        case "brown":
                fmt.Println(css.brown)
        case "burlywood":
                fmt.Println(css.burlywood)
        case "cadetblue":
                fmt.Println(css.cadetblue)
        case "chartreuse":
                fmt.Println(css.chartreuse)
        case "chocolate":
                fmt.Println(css.chocolate)
        case "coral":
                fmt.Println(css.coral)
        case "cornflowerblue":
                fmt.Println(css.cornflowerblue)
        case "cornsilk":
                fmt.Println(css.cornsilk)
        case "crimson":
                fmt.Println(css.crimson)
        case "cyan":
                fmt.Println(css.cyan)
        case "darkblue":
                fmt.Println(css.darkblue)
        case "darkcyan":
                fmt.Println(css.darkcyan)
        case "darkgoldenrod":
                fmt.Println(css.darkgoldenrod)
        case "darkgray":
                fmt.Println(css.darkgray)
        case "darkgreen":
                fmt.Println(css.darkgreen)
        case "darkkhaki":
                fmt.Println(css.darkkhaki)
        case "darkmagenta":
                fmt.Println(css.darkmagenta)
        case "darkolivegreen":
                fmt.Println(css.darkolivegreen)
        case "darkorange":
                fmt.Println(css.darkorange)
        case "darkorchid":
                fmt.Println(css.darkorchid)
        case "darkred":
                fmt.Println(css.darkred)
        case "darksalmon":
                fmt.Println(css.darksalmon)
        case "darkseagreen":
                fmt.Println(css.darkseagreen)
        case "darkslateblue":
                fmt.Println(css.darkslateblue)
        case "darkslategray":
                fmt.Println(css.darkslategray)
        case "darkturquoise":
                fmt.Println(css.darkturquoise)
        case "darkviolet":
                fmt.Println(css.darkviolet)
        case "deeppink":
                fmt.Println(css.deeppink)
        case "deepskyblue":
                fmt.Println(css.deepskyblue)
        case "dimgray":
                fmt.Println(css.dimgray)
        case "dodgerblue":
                fmt.Println(css.dodgerblue)
        case "firebrick":
                fmt.Println(css.firebrick)
        case "floralwhite":
                fmt.Println(css.floralwhite)
        case "forestgreen":
                fmt.Println(css.forestgreen)
        case "fuchsia":
                fmt.Println(css.fuchsia)
        case "gainsboro":
                fmt.Println(css.gainsboro)
        case "ghostwhite":
                fmt.Println(css.ghostwhite)
        case "gold":
                fmt.Println(css.gold)
        case "goldenrod":
                fmt.Println(css.goldenrod)
        case "gray":
                fmt.Println(css.gray)
        case "green":
                fmt.Println(css.green)
        case "greenyellow":
                fmt.Println(css.greenyellow)
        case "honeydew":
                fmt.Println(css.honeydew)
        case "hotpink":
                fmt.Println(css.hotpink)
        case "indianred":
                fmt.Println(css.indianred)
        case "indigo":
                fmt.Println(css.indigo)
        case "ivory":
                fmt.Println(css.ivory)
        case "khaki":
                fmt.Println(css.khaki)
        case "lavender":
                fmt.Println(css.lavender)
        case "lavenderblush":
                fmt.Println(css.lavenderblush)
        case "lawngreen":
                fmt.Println(css.lawngreen)
        case "lemonchiffon":
                fmt.Println(css.lemonchiffon)
        case "lightblue":
                fmt.Println(css.lightblue)
        case "lightcoral":
                fmt.Println(css.lightcoral)
        case "lightcyan":
                fmt.Println(css.lightcyan)
        case "lightgoldenrodyellow":
                fmt.Println(css.lightgoldenrodyellow)
        case "lightgreen":
                fmt.Println(css.lightgreen)
        case "lightgrey":
                fmt.Println(css.lightgrey)
        case "lightpink":
                fmt.Println(css.lightpink)
        case "lightsalmon":
                fmt.Println(css.lightsalmon)
        case "lightseagreen":
                fmt.Println(css.lightseagreen)
        case "lightskyblue":
                fmt.Println(css.lightskyblue)
        case "lightslategray":
                fmt.Println(css.lightslategray)
        case "lightsteelblue":
                fmt.Println(css.lightsteelblue)
        case "lightyellow":
                fmt.Println(css.lightyellow)
        case "lime":
                fmt.Println(css.lime)
        case "limegreen":
                fmt.Println(css.limegreen)
        case "linen":
                fmt.Println(css.linen)
        case "magenta":
                fmt.Println(css.magenta)
        case "maroon":
                fmt.Println(css.maroon)
        case "mediumaquamarine":
                fmt.Println(css.mediumaquamarine)
        case "mediumblue":
                fmt.Println(css.mediumblue)
        case "mediumorchid":
                fmt.Println(css.mediumorchid)
        case "mediumpurple":
                fmt.Println(css.mediumpurple)
        case "mediumseagreen":
                fmt.Println(css.mediumseagreen)
        case "mediumslateblue":
                fmt.Println(css.mediumslateblue)
        case "mediumspringgreen":
                fmt.Println(css.mediumspringgreen)
        case "mediumturquoise":
                fmt.Println(css.mediumturquoise)
        case "mediumvioletred":
                fmt.Println(css.mediumvioletred)
        case "midnightblue":
                fmt.Println(css.midnightblue)
        case "mintcream":
                fmt.Println(css.mintcream)
        case "mistyrose":
                fmt.Println(css.mistyrose)
        case "moccasin":
                fmt.Println(css.moccasin)
        case "navajowhite":
                fmt.Println(css.navajowhite)
        case "navy":
                fmt.Println(css.navy)
        case "navyblue":
                fmt.Println(css.navyblue)
        case "oldlace":
                fmt.Println(css.oldlace)
        case "olive":
                fmt.Println(css.olive)
        case "olivedrab":
                fmt.Println(css.olivedrab)
        case "orange":
                fmt.Println(css.orange)
        case "orangered":
                fmt.Println(css.orangered)
        case "orchid":
                fmt.Println(css.orchid)
        case "palegoldenrod":
                fmt.Println(css.palegoldenrod)
        case "palegreen":
                fmt.Println(css.palegreen)
        case "paleturquoise":
                fmt.Println(css.paleturquoise)
        case "palevioletred":
                fmt.Println(css.palevioletred)
        case "papayawhip":
                fmt.Println(css.papayawhip)
        case "peachpuff":
                fmt.Println(css.peachpuff)
        case "peru":
                fmt.Println(css.peru)
        case "pink":
                fmt.Println(css.pink)
        case "plum":
                fmt.Println(css.plum)
        case "powderblue":
                fmt.Println(css.powderblue)
        case "purple":
                fmt.Println(css.purple)
        case "red":
                fmt.Println(css.red)
        case "rosybrown":
                fmt.Println(css.rosybrown)
        case "royalblue":
                fmt.Println(css.royalblue)
        case "saddlebrown":
                fmt.Println(css.saddlebrown)
        case "salmon":
                fmt.Println(css.salmon)
        case "sandybrown":
                fmt.Println(css.sandybrown)
        case "seagreen":
                fmt.Println(css.seagreen)
        case "seashell":
                fmt.Println(css.seashell)
        case "sienna":
                fmt.Println(css.sienna)
        case "silver":
                fmt.Println(css.silver)
        case "skyblue":
                fmt.Println(css.skyblue)
        case "slateblue":
                fmt.Println(css.slateblue)
        case "slategray":
                fmt.Println(css.slategray)
        case "snow":
                fmt.Println(css.snow)
        case "springgreen":
                fmt.Println(css.springgreen)
        case "steelblue":
                fmt.Println(css.steelblue)
        case "tan":
                fmt.Println(css.tan)
        case "teal":
                fmt.Println(css.teal)
        case "thistle":
                fmt.Println(css.thistle)
        case "tomato":
                fmt.Println(css.tomato)
        case "turquoise":
                fmt.Println(css.turquoise)
        case "violet":
                fmt.Println(css.violet)
        case "wheat":
                fmt.Println(css.wheat)
        case "white":
                fmt.Println(css.white)
        case "whitesmoke":
                fmt.Println(css.whitesmoke)
        case "yellow":
                fmt.Println(css.yellow)
        case "yellowgreen":
                fmt.Println(css.yellowgreen)
        case "sample":
                fmt.Printf("%v%v", css.aliceblue, "aliceblue")
                fmt.Printf("%v%v", css.antiquewhite, "antiquewhite")
                fmt.Printf("%v%v", css.aqua, "aqua")
                fmt.Printf("%v%v", css.aquamarine, "aquamarine")
                fmt.Printf("%v%v", css.azure, "azure")
                fmt.Printf("%v%v", css.beige, "beige")
                fmt.Printf("%v%v", css.bisque, "bisque")
                fmt.Printf("%v%v", css.black, "black")
                fmt.Printf("%v%v", css.blanchedalmond, "blanchedalmond")
                fmt.Printf("%v%v", css.blue, "blue")
                fmt.Printf("%v%v", css.blueviolet, "blueviolet")
                fmt.Printf("%v%v", css.brown, "brown")
                fmt.Printf("%v%v", css.burlywood, "burlywood")
                fmt.Printf("%v%v", css.cadetblue, "cadetblue")
                fmt.Printf("%v%v", css.chartreuse, "chartreuse")
                fmt.Printf("%v%v", css.chocolate, "chocolate")
                fmt.Printf("%v%v", css.coral, "coral")
                fmt.Printf("%v%v", css.cornflowerblue, "cornflowerblue")
                fmt.Printf("%v%v", css.cornsilk, "cornsilk")
                fmt.Printf("%v%v", css.crimson, "crimson")
                fmt.Printf("%v%v", css.cyan, "cyan")
                fmt.Printf("%v%v", css.darkblue, "darkblue")
                fmt.Printf("%v%v", css.darkcyan, "darkcyan")
                fmt.Printf("%v%v", css.darkgoldenrod, "darkgoldenrod")
                fmt.Printf("%v%v", css.darkgray, "darkgray")
                fmt.Printf("%v%v", css.darkgreen, "darkgreen")
                fmt.Printf("%v%v", css.darkkhaki, "darkkhaki")
                fmt.Printf("%v%v", css.darkmagenta, "darkmagenta")
                fmt.Printf("%v%v", css.darkolivegreen, "darkolivegreen")
                fmt.Printf("%v%v", css.darkorange, "darkorange")
                fmt.Printf("%v%v", css.darkorchid, "darkorchid")
                fmt.Printf("%v%v", css.darkred, "darkred")
                fmt.Printf("%v%v", css.darksalmon, "darksalmon")
                fmt.Printf("%v%v", css.darkseagreen, "darkseagreen")
                fmt.Printf("%v%v", css.darkslateblue, "darkslateblue")
                fmt.Printf("%v%v", css.darkslategray, "darkslategray")
                fmt.Printf("%v%v", css.darkturquoise, "darkturquoise")
                fmt.Printf("%v%v", css.darkviolet, "darkviolet")
                fmt.Printf("%v%v", css.deeppink, "deeppink")
                fmt.Printf("%v%v", css.deepskyblue, "deepskyblue")
                fmt.Printf("%v%v", css.dimgray, "dimgray")
                fmt.Printf("%v%v", css.dodgerblue, "dodgerblue")
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
