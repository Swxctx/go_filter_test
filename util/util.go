package util

import (
	"log"
	"regexp"
)

// ConvertTitle2Slice 简单分词
func ConvertTitle2Slice(adstitle string) []string {
	var (
		titleSlice []string
	)
	title := ""
	titleTran := regexp.MustCompile(`[\p{Han}]+`).FindAllString(adstitle, -1)
	for _, str := range titleTran {
		title += str
	}
	log.Printf("标题:%s", title)

	//分词 两字
	titleRune := []rune(title)
	//fmt.Println(len(titleRune))
	if len(titleRune) >= 2 {
		for m := 0; m < len(titleRune)-1; m++ {
			titleSlice = append(titleSlice, string(titleRune[m:m+2]))
		}
	}

	if len(titleRune) >= 3 {
		for m := 0; m < len(titleRune)-2; m++ {
			titleSlice = append(titleSlice, string(titleRune[m:m+3]))
		}

	}

	if len(titleRune) >= 4 {
		for m := 0; m < len(titleRune)-3; m++ {
			titleSlice = append(titleSlice, string(titleRune[m:m+3]))
		}
	}
	return titleSlice
}
