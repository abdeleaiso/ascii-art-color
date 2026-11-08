package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
Some control escape sequences, like \e for ESC,
are not guaranteed to work in all languages and compilers.
It is recommended to use the decimal, octal or hex representation as escape code.
*/

func ColorMap(color string) string{
	m := map[string]string{
		"red":       "\033[38;5;196m",
		"green":     "\033[38;5;46m",
		"yellow":    "\033[38;5;226m",
		"blue":      "\033[38;5;21m",
		"purple":    "\033[38;5;129m",
		"cyan":      "\033[38;5;51m",
		"white":     "\033[38;5;15m",
		"black":     "\033[38;5;0m",

		"orange":    "\033[38;5;208m",
		"pink":      "\033[38;5;205m",
		"gold":      "\033[38;5;220m",
		"brown":     "\033[38;5;94m",
		"violet":    "\033[38;5;177m",
		"magenta":   "\033[38;5;201m",
		"turquoise": "\033[38;5;80m",
		"gray":      "\033[38;5;245m",
	}
	for k,v := range m{
		if color == k{
			return v
		} 
	}
	return ""
}

func main() {
	color := flag.String("color", "", "for coloring the txt")
	flag.Parse()
	if len(os.Args) == 2 {
		input := os.Args[1]
		FileData, err := os.ReadFile("standard.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		FileTxt := strings.ReplaceAll(string(FileData), "\r\n", "\n")
		Filelines := strings.Split(FileTxt, "\n")
		res := TransformToAscii(input, Filelines)
		fmt.Print(res)
	} else if len(os.Args) == 3 {
		*color = strings.ToLower(*color)
		colorTheme := ColorMap(*color)
		if colorTheme == ""{
			fmt.Println("Error: color not available")
			return
		}
		input := os.Args[2]
		FileData, err := os.ReadFile("standard.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		FileTxt := strings.ReplaceAll(string(FileData), "\r\n", "\n")
		Filelines := strings.Split(FileTxt, "\n")
		res := TransformToAscii(input, Filelines)
		fmt.Printf("%s%s\033[0m",colorTheme,res)
	} else if len(os.Args) == 4 {
		*color = strings.ToLower(*color)
		colorTheme := ColorMap(*color)
		if colorTheme == ""{
			fmt.Println("Error: color not available")
			return
		}
		subinput := os.Args[2]
		input := os.Args[3]
		FileData, err := os.ReadFile("standard.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		FileTxt := strings.ReplaceAll(string(FileData), "\r\n", "\n")
		Filelines := strings.Split(FileTxt, "\n")
		if input == "" || subinput == "" {
			fmt.Println("Substring doesn't exist")
			return
		}
		if strings.Contains(input, subinput) {
			res := ColorSubAscii(input, Filelines, subinput, colorTheme)
			fmt.Print(res)
		}else{
			fmt.Println("Substring doesn't exist")
			return
		}
	} else {
		fmt.Println("Usage: go run . --color=<color> <substring to be colored> 'something'")
		return
	}
}