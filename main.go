package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"main/cc"
	"os"
	"regexp"
	"strings"
)

func lastOf(s []int) int {
	return s[len(s)-1]
}
type ByteArr []byte

func (input ByteArr)formatJson() ByteArr {
	isInJson:=false
	bracketColorStack := make([]int,5)
	indentLv := 0
	indentStr := "  "

	var sb bytes.Buffer

	for i := range input {
		if input[i] == '{' || (input[i] == '[' && isInJson == true) {
			indentLv++
			if isInJson == false {
				sb.WriteByte('\n')
				isInJson = true
			}

			bracketColorStack = append(bracketColorStack, cc.NextColor())
			sb.WriteString(cc.Colorize(lastOf(bracketColorStack), string(input[i])))
			sb.WriteString("\n"+strings.Repeat(indentStr, indentLv))

		} else if input[i] == '}' || (input[i] == ']' && isInJson == true) {
			indentLv--
			if indentLv >= 0 {
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
	return sb.Bytes()
}



func (input ByteArr)highlightKeywords() ByteArr {
	strIpt := string(input)
	keywordsMatch:=[]string{"(?i)error", "(?i)warn", "(?i)info", "(?i)debug"}
	colorList:=[]int{cc.Red, cc.Yellow, cc.Cyan, cc.Green}

	for i:=range keywordsMatch {
		r:=regexp.MustCompile(keywordsMatch[i])
		strIpt = r.ReplaceAllStringFunc(strIpt, func(s string) string{
			return cc.Colorize(colorList[i], s)
		})
	}
	return []byte(strIpt)
}

func main() {
	var input ByteArr
	input,_=ioutil.ReadAll(os.Stdin)
	res := input.highlightKeywords().formatJson()
	fmt.Print(string(res))
}
