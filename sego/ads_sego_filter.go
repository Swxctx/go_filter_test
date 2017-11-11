package main

import (
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/bloom"
	cuckoo "github.com/goCuckoo"
	"github.com/huichen/sego"
)

var segmenter sego.Segmenter

func init() {
	// 载入词典
	segmenter.LoadDictionary("dictionary.txt")
}

func main() {
	// title := "全城疯抢,1元1GB还有腾讯VIP"
	// title := "屠龙一刀999"
	start_ts := time.Now().UnixNano() / 1000000
	title := "泰国3D眉笔眉粉染眉膏三合一防水防汗"
	log.Println("flag:", PoolFilter(title))
	end_ts := time.Now().UnixNano() / 1000000
	log.Printf("计算耗时:%v", end_ts-start_ts)
}

func PoolFilter(adstitle string) bool {
	filterWords := "广告"
	var adsTitleTwo []string

	title := ""
	titleTran := regexp.MustCompile(`[\p{Han}]+`).FindAllString(adstitle, -1)
	for _, str := range titleTran {
		title += str
	}

	log.Printf("标题:%s", title)

	// 分词
	text := []byte(title)
	segments := segmenter.Segment(text)
	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	adsTitleTwo = sego.SegmentsToSlice(segments, false)
	log.Printf("分词:", adsTitleTwo)

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

	//check filter
	//过滤
	for titleLen := 0; titleLen < len(adsTitleTwo); titleLen++ {

		//cockoo
		if cuckoofilter.Find([]byte(adsTitleTwo[titleLen])) {
			//fmt.Println("exist")
			log.Panicf("cuckoo-标题不合法1-非法词:%s", adsTitleTwo[titleLen])
			return true
		}

		//bloom
		if filter.Test([]byte(adsTitleTwo[titleLen])) {
			//fmt.Println("exist")
			log.Panicf("bloom标题不合法1-非法词:")
			return true
		}
	}

	for i := 0; i < len(strFilterWords); i++ {
		//bloom filter cant`t delete
		//cuckFilter insert
		cuckoofilter.Del([]byte(strFilterWords[i]))
	}
	return false
}
