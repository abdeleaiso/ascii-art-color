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

func ColorMap(color string) int{
	m := make(map[string]int)
	m["black"]=30 
	m["red"]=31
	m["green"]=32
	m["yellow"]=33
	m["blue"]=34
	m["purple"]=35
	m["cyan"]=36
	m["white"]=37
	m["default"]=39
	m["gray"]=90
	for k,v := range m{
		if color == k{
			return v
		} 
	}
	return -1
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
		colorN := ColorMap(*color)
		if colorN == -1{
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
		fmt.Printf("\033[1;%dm", colorN)
		fmt.Print(res)
	} else if len(os.Args) == 4 {
		*color = strings.ToLower(*color)
		colorN := ColorMap(*color)
		if colorN == -1{
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
		if strings.Contains(input, subinput) {
			res := ColorSubAscii(input, Filelines, subinput, colorN)
			fmt.Print(res)
		}else{
			fmt.Println("Error: subinput not found")
			return
		}
	} else {
		fmt.Println("Usage: go run . --color=<color> <substring to be colored> 'something'")
		return
	}
}