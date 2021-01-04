package main

import "strings"

var ez = [][]string{
	{
		"random",
		"regular",
		"fakie",
		"switch",
		"nollie",
		"x",
	},
	{
		"random",
		"frontside",
		"frontside",
		"backside",
		"backside",
		"x",
	},
	{
		"random",
		"360",
		"180",
		"360",
		"180",
		"x",
	},
	{
		"random",
		"flip",
		"flip",
		"heel",
		"heel",
		"x",
	},
	{
		"random",
		"no comply",
		"no comply",
		"double",
		"double",
		"x",
	},
	{
		"random",
		"grab",
		"grab",
		"rewind",
		"rewind",
		"x",
	},
}

var reg = [][]string{
	{
		"random",
		"regular",
		"fakie",
		"switch",
		"nollie",
		"x",
	},
	{
		"random",
		"frontside",
		"frontside",
		"backside",
		"backside",
		"x",
	},
	{
		"random",
		"360",
		"180",
		"Big",
		"Shuv",
		"x",
	},
	{
		"random",
		"flip",
		"flip",
		"heel",
		"heel",
		"x",
	},
	{
		"random",
		"no comply",
		"no comply",
		"double",
		"double",
		"x",
	},
	{
		"random",
		"grab",
		"grab",
		"rewind",
		"rewind",
		"x",
	},
}

// var vert = [][]string{
// 	{
// 		"random",
// 		"regular",
// 		"fakie",
// 		"switch",
// 		"nollie",
// 		"x",
// 	},
// 	{
// 		"random",
// 		"frontside",
// 		"frontside",
// 		"backside",
// 		"backside",
// 		"x",
// 	},
// 	{
// 		"random",
// 		"air",
// 		"grind",
// 		"stall",
// 		"grab",
// 		"x",
// 	},
// 	{
// 		"random",
// 		"flip",
// 		"flip",
// 		"heel",
// 		"heel",
// 		"x",
// 	},
// 	{
// 		"random",
// 		"no comply",
// 		"no comply",
// 		"double",
// 		"double",
// 		"x",
// 	},
// 	{
// 		"random",
// 		"grab",
// 		"grab",
// 		"rewind",
// 		"rewind",
// 		"x",
// 	},
// }

var kook = [][]string{
	{
		"random",
		"regular",
		"fakie",
		"switch",
		"nollie",
		"x",
	},
	{
		"random",
		"frontside",
		"frontside",
		"backside",
		"backside",
		"x",
	},
	{
		"random",
		"360",
		"180",
		"Big",
		"Shuv",
		"x",
	},
	{
		"random",
		"boneless",
		"bonlesss",
		"benihanna",
		"benihanna",
		"x",
	},
	{
		"random",
		"no comply",
		"no comply",
		"double",
		"double",
		"x",
	},
	{
		"random",
		"grab",
		"grab",
		"rewind",
		"rewind",
		"x",
	},
}

var qs = map[string][][]string{
	"Grom":  ez,
	"Flow":  ez,
	"Am":    reg,
	"Pro":   reg,
	"Yummy": reg,
}

func getTitle(t string) string {
	strings.Replace(t, "fakie backside 360", "caballerial", -1)
	strings.Replace(t, "fakie backside 180", "half cab", -1)
	strings.Replace(t, "360 flip", "tre", -1)
	//strings.Replace(t, "360 heel", "laser", -1)
	return t
}
