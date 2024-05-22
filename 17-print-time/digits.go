package main

type digit [5]string

var zero = digit{
	" ███ ",
	"█   █",
	"█   █",
	"█   █",
	" ███ ",
}

var one = digit{
	"  █  ",
	" ██  ",
	"  █  ",
	"  █  ",
	" ███ ",
}

var two = digit{
	" ███ ",
	"█   █",
	"   █ ",
	"  █  ",
	"█████",
}

var three = digit{
	" ███ ",
	"█   █",
	"  ██ ",
	"█   █",
	" ███ ",
}

var four = digit{
	"█   █",
	"█   █",
	"█████",
	"    █",
	"    █",
}

var five = digit{
	"█████",
	"█    ",
	"████ ",
	"    █",
	"████ ",
}

var six = digit{
	" ███ ",
	"█    ",
	"████ ",
	"█   █",
	" ███ ",
}

var seven = digit{
	"█████",
	"   █ ",
	"  █  ",
	" █   ",
	"█    ",
}

var eight = digit{
	" ███ ",
	"█   █",
	" ███ ",
	"█   █",
	" ███ ",
}

var nine = digit{
	" ███ ",
	"█   █",
	" ████",
	"    █",
	" ███ ",
}

var colon = digit{
	"     ",
	"  ▒  ",
	"     ",
	"  ▒  ",
	"     ",
}

var blank = digit{
	"     ",
	"     ",
	"     ",
	"     ",
	"     ",
}

var digits = [...]digit{zero, one, two, three, four, five, six, seven, eight, nine, colon, blank}
