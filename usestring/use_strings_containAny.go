package usestring

import (
	"log"
	"strings"
	"time"
	"xc/myGit/go_filter_test/util"
)

func PoolFilterContainsAny(title string, filterWords string) bool {
	start_ts := time.Now().UnixNano() / 1000000
	//remove ,
	// strFilterWords := strings.Split(filterWords, ",")
	// log.Printf("过滤词：%v", strFilterWords)

	//简单分词
	titleTran := util.ConvertTitle2Slice(title)
	for _, v := range titleTran {
		//containsAny
		if strings.ContainsAny(filterWords, v) {
			//fmt.Println("exist")
			log.Println("标题不合法-非法词:%s", v)
			return true
		}
	}

	end_ts := time.Now().UnixNano() / 1000000
	log.Println("耗时(ms):", end_ts-start_ts)

	return false
}
