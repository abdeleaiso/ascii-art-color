package main

import (
	"strings"
)

func TransformToAscii(input string, filelines []string) string{
	if input == ""{
		return ""
	}
	input = strings.ReplaceAll(input, "\\n", "\n")
	inputLines := strings.Split(input, "\n")
	isEmpty := true
	for i:=0; i<len(inputLines); i++{
		if inputLines[i] != ""{
			isEmpty = false
			break
		}
	}

	if isEmpty{
		return strings.Repeat("\n", len(input))
	}
	
	
	var sb strings.Builder
	
	for _, line := range inputLines{
		if line == ""{
			sb.WriteByte('\n')
			continue
		}
		for i:=0; i<8; i++{
			for _,c := range line{
				if c < 32 || c > 126{
					continue
				}
				StartIndex := (int(c)-32)*9 + 1
				if StartIndex+i < len(filelines){
					sb.WriteString(filelines[StartIndex+i])
				}
			}
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}