package main

import (
	"math"
	"regexp"
	"strconv"
)

// HSLtoRGB converts HSL values to RGB values
func HSLtoRGB(h int, s int, l int) (int, int, int) {
	s2 := float64(s) / float64(100)
	l2 := float64(l) / float64(100)
	c := (1 - math.Abs(2*l2-float64(1))) * s2
	x := c * (float64(1) - math.Abs(math.Mod(float64(h)/60, 2)-float64(1)))
	m := l2 - c/2
	var r, g, b float64
	switch {
	case h < 60:
		r = c
		g = x
		b = 0
	case h < 120 && h >= 60:
		r = x
		g = c
		b = 0
	case h < 180 && h >= 120:
		r = 0
		g = c
		b = x
	case h < 240 && h >= 180:
		r = 0
		g = x
		b = c
	case h < 300 && h >= 240:
		r = x
		g = 0
		b = c
	case h < 360 && h >= 300:
		r = c
		g = 0
		b = x
	}
	R := int(math.Round((r + m) * 255))
	G := int(math.Round((g + m) * 255))
	B := int(math.Round((b + m) * 255))
	return R, G, B
}

// HextoRGB converts a hexadecimal string to RGB values
func HextoRGB(hex string) (int, int, int) {
	if hex[0:1] == "#" {
		hex = hex[1:]
	}
	r := string(hex)[0:2]
	g := string(hex)[2:4]
	b := string(hex)[4:6]
	R, _ := strconv.ParseInt(r, 16, 0)
	G, _ := strconv.ParseInt(g, 16, 0)
	B, _ := strconv.ParseInt(b, 16, 0)

	return int(R), int(G), int(B)
}

// RGBtoAnsi converts a color code string to an Ansi escape code
func RGBtoAnsi(color string) string {
	re := regexp.MustCompile("[0-9]+")
	colors := re.FindAllString(color, -1)
	r, _ := strconv.Atoi(colors[0])
	g, _ := strconv.Atoi(colors[1])
	b, _ := strconv.Atoi(colors[2])
	str := "\x1b[38;2;" + strconv.FormatInt(int64(r), 10) + ";" + strconv.FormatInt(int64(g), 10) + ";" + strconv.FormatInt(int64(b), 10) + "m"
	return str
}

// HSLtoAnsi converts a color code string to an Ansi escape code
func HSLtoAnsi(color string) string {
	re := regexp.MustCompile("[0-9]+")
	colors := re.FindAllString(color, -1)
	h, _ := strconv.Atoi(colors[0])
	s, _ := strconv.Atoi(colors[1])
	l, _ := strconv.Atoi(colors[2])
	r, g, b := HSLtoRGB(h, s, l)
	str := "\x1b[38;2;" + strconv.FormatInt(int64(r), 10) + ";" + strconv.FormatInt(int64(g), 10) + ";" + strconv.FormatInt(int64(b), 10) + "m"
	return str
}

// HextoAnsi converts a hexadecimal string to an Ansi escape code
func HextoAnsi(hex string) string {
	R, G, B := HextoRGB(hex)
	str := "\x1b[38;2;" + strconv.FormatInt(int64(R), 10) + ";" + strconv.FormatInt(int64(G), 10) + ";" + strconv.FormatInt(int64(B), 10) + "m"
	return str
}