package main

import (
	"strings"
	"fmt"
)

func IndexSubString(s, sub string) [][2]int{
	var arr [][2]int
	offset := 0
	for {
		startIndex := strings.Index(s[offset:], sub)
		if startIndex == -1{
			break
		}
		startIndex += offset
		endIndex := startIndex + len(sub)-1
		//modstr := s[:startIndex] + s[startIndex+len(sub):]
		arr = append(arr, [2]int{startIndex, endIndex})
		offset = endIndex + 1
	}
	return arr
}

func ColorSubAscii(input string, filelines []string, sub string, colorTheme string) string{
	if input == ""{
		return ""
	}
	nbr := strings.Count(input,sub)
	slice := make([][2]int, nbr)
	slice = IndexSubString(input, sub)
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
			for pos,c := range line{
				if c < 32 || c > 126{
					continue
				}
				StartIndex := (int(c)-32)*9 + 1
				shouldbecolored := false
				for i:= 0; i < nbr; i++{
					if (pos >= slice[i][0] && pos <= slice[i][1]){
						shouldbecolored = true
						break
					}
				}

				if shouldbecolored && StartIndex < len(filelines){					
					colored := fmt.Sprintf("%s%s\033[0m",colorTheme,filelines[StartIndex+i])
					sb.WriteString(colored)
				}else if StartIndex+i < len(filelines){
					sb.WriteString(filelines[StartIndex+i])
				}
			}
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}