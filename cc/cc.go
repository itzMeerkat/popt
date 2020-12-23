package cc

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

const (
	Red    = iota
	Green
	Yellow
	Blue
	Purple
	Cyan
	Gray
)

const colorNum = 6

var colorIdx = -1
var colorCycle = [7]string{red, green, yellow, blue, purple, cyan, gray}

func Colorize(color int, f string) string {
	return colorCycle[color]+f+reset
}

func NextColor() int {
	colorIdx++
	colorIdx %= colorNum
	return colorIdx
}