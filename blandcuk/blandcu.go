package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/bloom"
	cuckoo "github.com/goCuckoo"
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

	log.Printf("标题:%s", title)
	//remove ,
	strFilterWords := strings.Split(filterWords, ",")
	log.Printf("过滤词：%v", strFilterWords)

	//cuckooFilter
	cuckoofilter := cuckoo.NewFilter(10000)

	//bloom filter
	n := uint(1000)
	filter := bloom.New(20*n, uint(len(strFilterWords))) // load of 20, 5 keys

	//add filter
	for i := 0; i < len(strFilterWords); i++ {
		//bloom filter add
		filter.Add([]byte(strFilterWords[i]))
		//cuckFilter insert
		cuckoofilter.Insert([]byte(strFilterWords[i]))
	}

	//分词 两字
	titleRune := []rune(title)
	//fmt.Println(len(titleRune))
	if len(titleRune) >= 2 {
		for m := 0; m < len(titleRune)-1; m++ {
			adsTitleTwo = append(adsTitleTwo, string(titleRune[m:m+2]))
		}
		//check filter
		//过滤
		for titleLen := 0; titleLen < len(adsTitleTwo); titleLen++ {

			//cockoo
			if cuckoofilter.Find([]byte(adsTitleTwo[titleLen])) {
				log.Panicf("cuckoo标题不合法1-非法词:%s", adsTitleTwo[titleLen])
				return true
			}

			//bloom
			if filter.Test([]byte(adsTitleTwo[titleLen])) {
				log.Panicf("bloom标题不合法1-非法词:%s", adsTitleTwo[titleLen])
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
			//cockoo
			if cuckoofilter.Find([]byte(adsTitleThree[titleLen])) {
				log.Panicf("cuckoo标题不合法2-非法词:%s", adsTitleThree[titleLen])
				return true
			}
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
			//cockoo
			if cuckoofilter.Find([]byte(adsTitleTwo[titleLen])) {
				log.Panicf("cuckoo标题不合法3-非法词:%s", adsTitleThree[titleLen])
				return true
			}
			//bloom
			if filter.Test([]byte(adsTitleFour[titleLen])) {
				log.Panicf("bloom标题不合法3-非法词")
				return true
			}
		}
	}
	//bloom filter cant`t delete
	return false
}
