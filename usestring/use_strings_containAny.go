package usestring

import (
	"go_filter_test/util"
	"log"
	"strings"
	"time"
)

// PoolFilterContainsAny 是否包含多个
func PoolFilterContainsAny(title string, filterWords string) bool {
	start_ts := time.Now().UnixNano()
	//remove ,
	// strFilterWords := strings.Split(filterWords, ",")
	// log.Printf("过滤词：%v", strFilterWords)

	//简单分词
	titleTran := util.ConvertTitle2Slice(title)
	for _, v := range titleTran {
		//containsAny
		if strings.ContainsAny(filterWords, v) {
			log.Println("标题不合法-非法词:", v)
			end_ts := time.Now().UnixNano()
			log.Println("耗时(ns):", end_ts-start_ts)
			return true
		}
	}
	return false
}
