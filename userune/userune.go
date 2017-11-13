package userune

import (
	"go_filter_test/util"
	"log"
	"strings"
	"time"
)

// PoolFilterRune 遍历标题以及过滤词数组过滤
func PoolFilterRune(title string, filterWords string) bool {
	start_ts := time.Now().UnixNano()

	//remove ,
	strFilterWords := strings.Split(filterWords, ",")

	//简单分词
	titleTran := util.ConvertTitle2Slice(title)
	for _, v := range titleTran {
		for _, s := range strFilterWords {
			if v == s {
				log.Println("标题不合法-非法词:", v)
				end_ts := time.Now().UnixNano()
				log.Println("耗时(us):", end_ts-start_ts)
				return true
			}
		}
	}
	return false
}
