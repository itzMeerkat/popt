package main

import (
	"fmt"
	"io/ioutil"
	"main/cc"
	"os"
	"strings"
)

func lastOf(s []int) int {
	return s[len(s)-1]
}

func main() {
	isInJson:=false
	bracketColorStack := make([]int,5)
	indentLv := 0
	indentStr := "  "
	input,_:=ioutil.ReadAll(os.Stdin)
	var sb strings.Builder

	for i := range input {
		if input[i] == '{' || input[i] == '['{
			indentLv++
			if isInJson == false {
				sb.WriteByte('\n')
				isInJson = true
			}

			bracketColorStack = append(bracketColorStack, cc.NextColor())
			sb.WriteString(cc.Colorize(lastOf(bracketColorStack), string(input[i])))
			sb.WriteString("\n"+strings.Repeat(indentStr, indentLv))

		} else if input[i] == '}' || input[i] == ']'{
			indentLv--
			if indentLv >= 0{
				sb.WriteByte('\n')
				sb.WriteString(strings.Repeat(indentStr, indentLv))
				sb.WriteString(cc.Colorize(lastOf(bracketColorStack), string(input[i])))
				bracketColorStack = bracketColorStack[:len(bracketColorStack)-1]
			}
			if indentLv <= 0 {
				isInJson = false
				if indentLv < 0 {
					sb.WriteString(cc.Colorize(cc.Gray, "}"))
				}
				indentLv=0
			}
		} else {
			if isInJson == false {
				sb.WriteByte(input[i])
			} else {
				if input[i] == ',' {
					sb.WriteString(",\n" + strings.Repeat(indentStr, indentLv))
				}else if input[i] != ' ' && input[i]!='\n'{
					sb.WriteByte(input[i])
				}
			}
		}
	}
	sb.WriteByte('\n')
	fmt.Print(sb.String())
}
