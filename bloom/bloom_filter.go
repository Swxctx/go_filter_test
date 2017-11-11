package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/bloom"
)

func main() {
	// title := "全城疯抢,1元1GB还有腾讯VIP"
	// title := "屠龙一刀999"
	// title := "泰国3D眉笔眉粉染眉膏三合一防水防汗"
	title := "网易新闻 - 头条视频资讯阅读平台"
	log.Println("flag:", PoolFilter(title))
}

func PoolFilter(adstitle string) bool {
	filterWords := "广告"
	var adsTitleTwo []string
	var adsTitleThree []string
	var adsTitleFour []string

	title := ""
	titleTran := regexp.MustCompile(`[\p{Han}]+`).FindAllString(adstitle, -1)
	for _, str := range titleTran {
		title += str
	}

	log.Println("标题:%s", title)
	//remove ,
	strFilterWords := strings.Split(filterWords, ",")
	log.Println("过滤词：%v", strFilterWords)

	//bloom filter
	n := uint(1000)
	filter := bloom.New(20*n, uint(len(strFilterWords))) // load of 20, 5 keys

	//add filter
	for i := 0; i < len(strFilterWords); i++ {
		//bloom filter add
		filter.Add([]byte(strFilterWords[i]))
	}

	//分词 两字
	titleRune := []rune(title)
	if len(titleRune) >= 2 {
		for m := 0; m < len(titleRune)-1; m++ {
			adsTitleTwo = append(adsTitleTwo, string(titleRune[m:m+2]))
		}
		//check filter
		//过滤
		for titleLen := 0; titleLen < len(adsTitleTwo); titleLen++ {
			//bloom
			if filter.Test([]byte(adsTitleTwo[titleLen])) {
				log.Panic("bloom标题不合法1-非法词:")
				return true
			}
		}
	}

	if len(titleRune) >= 3 {
		for m := 0; m < len(titleRune)-2; m++ {
			adsTitleThree = append(adsTitleThree, string(titleRune[m:m+3]))
		}

		//check filter
		for titleLen := 0; titleLen < len(adsTitleThree); titleLen++ {
			//bloom
			if filter.Test([]byte(adsTitleThree[titleLen])) {
				log.Panicf("标题不合法2-非法词:%s", adsTitleThree[titleLen])
				return true
			}
		}

	}

	if len(titleRune) >= 4 {
		for m := 0; m < len(titleRune)-3; m++ {
			adsTitleFour = append(adsTitleFour, string(titleRune[m:m+3]))
		}
		//check filter
		for titleLen := 0; titleLen < len(adsTitleFour); titleLen++ {
			//bloom
			if filter.Test([]byte(adsTitleFour[titleLen])) {
				log.Panic("bloom标题不合法3-非法词")
				return true
			}
		}
	}
	return false
}
