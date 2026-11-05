package main

import (
	"strings"
	"fmt"
)

func IndexSubString(s, sub string) (int, int){
	startIndex := strings.Index(s, sub)
	if startIndex == -1{
		return -1 , -1
	}
	endIndex := startIndex + len(sub)-1
	//modstr := s[:startIndex] + s[startIndex+len(sub):]
	return startIndex, endIndex
}

func TransformToAscii(input string, filelines []string, sub string, colorN int) string{
	if input == ""{
		return ""
	}
	start, end := IndexSubString(input, sub)
	//fmt.Println("------------------- ", start ," -------------------------")
	startInAscii := (int(input[start])-32)*9+1
	endInAscii := (int(input[end])-32)*9+1
	fmt.Println("------------------- ", endInAscii ," -------------------------")
	input = strings.ReplaceAll(input, "\\n", "\n")
	/*fmt.Println("-----------",input,"-----------")
	fmt.Println(len(input))
	for _, c := range input {
		fmt.Println(string(c))
	}*/
	inputLines := strings.Split(input, "\n")
	/*fmt.Println("-----------",inputLines,"-----------")
	fmt.Println(len(inputLines))
	for _, c := range inputLines {
		fmt.Println(string(c))
	}*/
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
		/*fmt.Print("index=",index," | ")
		fmt.Println("inputlines=",inputLines)*/
		if line == ""{
			sb.WriteByte('\n')
			continue
		}
		for i:=0; i<8; i++{
			for _,c := range line{
				//fmt.Print("c=",string(c), " | ")
				if c < 32 || c > 126{
					continue
				}
				StartIndex := (int(c)-32)*9 + 1
				//fmt.Print("startIndex=",StartIndex, " | ")
				if startInAscii == StartIndex && StartIndex+i < len(filelines){
					colored := fmt.Sprintf("\033[1;%dm%s\033[0m",colorN,filelines[StartIndex+i])
					sb.WriteString(colored)
				}else if  endInAscii == StartIndex && StartIndex+i < len(filelines){
					colored := fmt.Sprintf("\033[1;%dm%s\033[0m",colorN,filelines[StartIndex+i])
					sb.WriteString(colored)
				}else if StartIndex+i < len(filelines){
					//fmt.Printf("StartIndex+%d=%d\n",i, StartIndex+i)
					sb.WriteString(filelines[StartIndex+i])
				}
			}
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}