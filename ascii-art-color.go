package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type sliceOfFlags []string
type Line struct {
	lineText        string
	colorIndexSlice []string
}

// colorMap is a map that maps color names to the corresponding escape codes
// used to print colored text in the terminal
var colorMap = map[string]string{
	"black":  "\x1b[30m",
	"red":    "\x1b[31m",
	"green":  "\x1b[32m",
	"yellow": "\x1b[33m",
	"blue":   "\x1b[34m",
	"purple": "\x1b[35m",
	"cyan":   "\x1b[36m",
	"white":  "\x1b[37m",
	"reset":  "\x1b[0m",
	"orange": "\u001b[38;2;255;165;0m",
}

func main() {
	checkForAudit()
	colorSlice, _, _ := getFlags()
	characterMap, inputText := getArguments()
	listOfColors := getColorIndexes(inputText, colorSlice)
	lines := splitInput(inputText, listOfColors)
	printASCII(lines, characterMap)

}

func checkForAudit() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
	if strings.Contains(os.Args[1], "--") && !strings.Contains(os.Args[1], "=") {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
}

func getColorIndexes(inputText string, colorSlice sliceOfFlags) []string {
	listOfColors := make([]string, len(inputText))
	for i := range listOfColors {
		listOfColors[i] = "\x1b[0m"
	}
	for _, v := range colorSlice {
		if strings.Contains(v, "[") {
			color := v[:strings.IndexByte(v, '[')]
			color = translateColor(color)
			flag := v[strings.IndexByte(v, '[')+1 : strings.IndexByte(v, ']')]
			tag := v[strings.IndexByte(v, ']')+1:]
			switch flag {
			case "word":
				r, _ := regexp.Compile(tag)
				elina := r.FindAllStringIndex(inputText, -1)
				listOfColors = applyIndexToList(elina, color, listOfColors)
			case "letter":
				r, _ := regexp.Compile(`[` + tag + `]`)
				elina := r.FindAllStringIndex(inputText, -1)
				listOfColors = applyIndexToList(elina, color, listOfColors)
			case "index":
				elina := indexMagic(tag, inputText)
				listOfColors = applyIndexToList(elina, color, listOfColors)
			}
		} else {
			for i := range listOfColors {
				color := translateColor(v)
				listOfColors[i] = color
			}
		}
	}
	return listOfColors
}

func translateColor(color string) string {
	var tranlatedColor string
	if strings.Contains(color, "#") {
		tranlatedColor = HextoAnsi(color)
	} else if strings.Contains(color, "rgb") {
		tranlatedColor = RGBtoAnsi(color)
	} else if strings.Contains(color, "hsl") {
		tranlatedColor = HSLtoAnsi(color)
	} else {
		return colorMap[color]
	}
	return tranlatedColor
}

func indexMagic(tag string, inputText string) [][]int {
	si := strings.Split(tag, ",")
	elina := [][]int{}
	for _, v := range si {
		start, end := 0, 0
		temp, err := strconv.Atoi(v)
		start = temp
		end = temp + 1
		if err != nil {
			if strings.Contains(v, ":") {
				st := strings.Split(v, ":")
				start, err = strconv.Atoi(st[0])
				if err != nil {
					start = 0
				}
				end, err = strconv.Atoi(st[1])
				if err != nil {
					end = len(inputText)
				}
			} else {
				panic(err)
			}
		}
		elina = append(elina, []int{start, end})
	}
	return elina
}

func applyIndexToList(elina [][]int, color string, listOfColors []string) []string {
	for _, v := range elina {
		for i := v[0]; i < v[1]; i++ {
			listOfColors[i] = color
		}
	}
	return listOfColors
}

// accepting flags, using special variable to collect n number of options for coloring with use of single flag multiple times
func getFlags() (sliceOfFlags, string, string) {
	var colorSlice sliceOfFlags
	var output string
	var align string
	flag.Var(&colorSlice, "color", "please see README.md for usage")
	//flag.StringVar(&output, "outfile", "fileName.txt", "erroro message")
	//flag.StringVar(&align, "align", "left", "erroro message")
	flag.Parse()
	return colorSlice, output, align
}

// accepting arguments
func getArguments() ([855]string, string) {
	// returns non-flag command-text arguments
	args := flag.Args()
	var inputFile *os.File
	var err error
	if len(args) == 2 {
		if !strings.HasSuffix(args[1], ".txt") {
			args[1] += ".txt"
		}
		args[1] = "Font/" + args[1]
		inputFile, err = os.Open(args[1])
		if err != nil {
			log.Fatalf("\"%v\" is not valid name for character map", args[1])
		}
	} else if len(args) == 1 {
		inputFile, err = os.Open("Font/standard.txt")
		if err != nil {
			log.Fatal("Could not find Standard letter map!!")
		}
	} else {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
	}
	characterMap := scanFile(inputFile)
	inputText := args[0]
	return characterMap, inputText
}

// Printing out ascii text line by line
func printASCII(lines []Line, characterMap [855]string) {
	for i, v := range lines {
		if v.lineText == "" && i == 0 {
			continue
		}
		if len(v.lineText) == 0 {
			fmt.Println()
			continue
		}
		printLine(v.lineText, characterMap, v.colorIndexSlice)
	}
}

func printLine(inputText string, characterMap [855]string, colorIndexSlice []string) {
	byteText := []byte(inputText)
	var text [8]string
	for i, v := range byteText {
		indexOfFirstLineOfLetter := ((int(v) - 32) * 9) + 1
		for letter := range text {
			text[letter] = text[letter] + colorIndexSlice[i] + characterMap[indexOfFirstLineOfLetter] + colorMap["reset"]
			indexOfFirstLineOfLetter++
		}
	}
	for _, v := range text {
		fmt.Println(v)
	}
}

func scanFile(inputFile *os.File) [855]string {
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)
	var characterMap [855]string
	n := 0
	for scanner.Scan() {
		characterMap[n] = scanner.Text()
		n++
	}
	inputFile.Close()
	return characterMap
}

// Check if text starts with newline (special case) and split input text by the new text "\n" symbols
func splitInput(inputText string, listOfColors []string) []Line {
	l := Line{
		lineText:        "",
		colorIndexSlice: []string{},
	}
	lines := []Line{l}
	if len(inputText) == 0 {
		os.Exit(0)
	} else if string(inputText[0])+string(inputText[1]) == "\\n" {
		inputText = inputText[2:]
		fmt.Println()
		if len(inputText) == 0 {
			os.Exit(0)
		}
	}
	splittedInputText := strings.Split(inputText, "\\n")
	n, m := 0, 0
	for _, v := range splittedInputText {
		l.lineText = v
		m = n + len(v)
		splittedListOfColors := listOfColors[n:m]
		l.colorIndexSlice = splittedListOfColors
		lines = append(lines, l)
		n = m + 2 //we skip each time two indexes because we take out \n
	}
	return lines
}

func (i *sliceOfFlags) String() string {
	return "my string representation"
}

func (i *sliceOfFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}
