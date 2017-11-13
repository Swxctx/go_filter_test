package usestring

import (
	"go_filter_test/util"
	"log"
	"strings"
	"time"
)

// PoolFilterIndexAny 第一次出现的位置
func PoolFilterIndexAny(title string, filterWords string) bool {
	start_ts := time.Now().UnixNano()
	//remove ,
	// strFilterWords := strings.Split(filterWords, ",")
	// log.Printf("过滤词：%v", strFilterWords)

	//简单分词
	titleTran := util.ConvertTitle2Slice(title)

	for _, v := range titleTran {
		//index
		if strings.IndexAny(filterWords, v) != -1 {
			//fmt.Println("exist")
			log.Println("标题不合法-非法词:", v)
			end_ts := time.Now().UnixNano()
			log.Println("耗时(ns):", end_ts-start_ts)

			return true
		}
	}
	return false
}
