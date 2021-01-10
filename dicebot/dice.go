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
		"kickflip",
		"kickflip",
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

var jb = [][]string{
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
		"manual",
		"manual",
		"nose manual",
		"nose manual",
		"x",
	},
	{
		"random",
		"frontside 180 in",
		"backside 180 in",
		"frontside 180 in",
		"backside 180 in",
		"x",
	},
	{
		"random",
		"1 try",
		"2 tries",
		"3 tries",
		"4 tries",
		"x",
	},
	{
		"random",
		"shuv out",
		"front shove out",
		"shuv out",
		"front shove out",
		"x",
	},
	{
		"random",
		"with a flip",
		"with a flip",
		"with a heel",
		"with a heel",
		"x",
	},
}

var grind = [][]string{
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
		"nose grind",
		"50 50",
		"5-0 grind",
		"board slide",
		"tail silde",
		"x",
	},
	{
		"random",
		"1 try",
		"2 tries",
		"3 tries",
		"4 tries",
		"x",
	},
	{
		"random",
		"tweaked up",
		"tweaked down",
		"reverse",
		"reverse",
		"x",
	},
	{
		"random",
		"blunt",
		"blunt",
		"over",
		"over",
		"x",
	},
}

var gr = [][]string{
	{
		"random",
		"regular",
		"fakie",
		"regular",
		"fakie",
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
		"",
		"180",
		"",
		"",
		"x",
	},
	{
		"random",
		"kickflip",
		"pop shuv",
		"shuv-it",
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
		"big",
		"shuv",
		"x",
	},
	{
		"random",
		"kickflip",
		"kickflip",
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

var vert = [][]string{
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
		"air",
		"grind",
		"stall",
		"grab",
		"x",
	},
	{
		"random",
		"to fakie",
		"to tail",
		"to disaster",
		"to fakie",
		"x",
	},
	{
		"random",
		"180",
		"180",
		"360",
		"540",
		"x",
	},
	{
		"random",
		"kickflip",
		"kickflip",
		"heelflip",
		"heelflip",
		"x",
	},
}

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
	"Grom":  gr,
	"Flow":  ez,
	"Am":    reg,
	"Pro":   reg,
	"Yummy": reg,
}

func getTitle(t string) string {
	t = strings.Replace(t, "fakie backside 360", "caballerial", -1)
	t = strings.Replace(t, "fakie backside 180", "half cab", -1)
	t = strings.Replace(t, "360 kickflip", "tre", -1)
	t = strings.Replace(t, "regular", "", -1)
	t = strings.Replace(t, "180 pop shuv", "pop shuv", -1)
	t = strings.Replace(t, "180 shuv-it", "shuv-it", -1)
	return t
}
