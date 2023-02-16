# Ascii-Art-Color

This program receives color(s) and a string passed in as arguments and prints out a string as a colorful ASCII art. The project is written in Go.

## How to add color to string?

To add colors to the string you have to use --color flag, specify the desired color and parts of text you would like to color. This program provides following coloring options:

### Color whole text

```*
go run . --color=green "I've got 99 problems, but code ain't 1"
```

### Color specific word

Add [word] followed by word you would like to color

```*
go run . --color=green[word]problems "I've got 99 problems, but code ain't 1"
go run . --color=green[word]99 "I've got 99 problems, but code ain't 1"
```

### Color specific letter(s)

Add [letter] followed by one or multiple letters you would like to color

```*

go run . --color=green[letter]o "I've got 99 problems, but code ain't 1"
go run . --color=green[letter]oas "I've got 99 problems, but code ain't 1"
go run . --color=green[letter]1 "I've got 99 problems, but code ain't 1"

```

### Color letter(s) by index

Add [index] followed by one or multiple indexes seperated by comma you would like to color

Use ":" to color group of indexes, for example, 4:6 will color letters on indexes 4 and 5

```*

go run . --color=green[index]0 "I've got 99 problems, but code ain't 1"
go run . --color=green[index]1,9,12 "I've got 99 problems, but code ain't 1"
go run . --color=green[index]:3,5:7,12: "I've got 99 problems, but code ain't 1"

```

### Combine different flags

Mix and match multiple color flags to create a colorful text. In case of conflict, flag that is later in the line will be executed

```*

go run . --color=red[index]0 --color=orange[letter]A --color=yellow[index]2 --color=green[index]3 --color=blue[index]4 --color=cyan[index]5 --color=purple[index]6 "RAINBOW"
```

### Wrong flag

The flag must have exactly the same format as mentioned above, any other formats will return the following usage message

```*

go run . --color red "banana"

Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <letters to be colored> "something"
```

## Supported color codes

This program supports ANSI, RGB, HSL and HEX color code systems. Please use them as following:

### ANSI

```*
go run . --color=red "Eat your tomatoes"
```

> Supported colors: black, red, green, yellow, blue, purple, cyan, white, orange and reset

### HEX

```*
go run . --color=#32ad64 "Eat your broccoli"
```

### RGB

```*
go run . --color="rgb(252, 140, 3)" "Eat your carrots"
```

### HSL

```*
go run . --color="hsl(281, 79%, 62%)" "Eat your eggplant"
```

</li>
</ul>
**All color codes can be combined with all flag types: [word], [letter], [index]**

## Choose ASCII font

In addition to colors, you can choose a font in which ASCII art should be dislayed. Font name should be mentioned as a last argument, for example:

```*
go run . --color=red "Eat your tomatoes" shadow
```

> Supported fonts: standard, shadow, thinkertoy

## Test the program

Give permissions to access and read audit.sh file

```*
cmhod +x audit.sh
```

Run the audit.sh file

```*
./audit.sh
```
