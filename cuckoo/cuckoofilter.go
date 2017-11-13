package cuckoo

import (
	"log"
	"strings"
	"time"

	"github.com/Swxctx/go_filter_test/util"
	cuckoo "github.com/goCuckoo"
)

func PoolFilter(title string, filterWords string) bool {
	start_ts := time.Now().UnixNano()
	//remove ,
	strFilterWords := strings.Split(filterWords, ",")

	//cuckooFilter
	cuckoofilter := cuckoo.NewFilter(10000)

	//add filter
	for i := 0; i < len(strFilterWords); i++ {
		//cuckFilter insert
		cuckoofilter.Insert([]byte(strFilterWords[i]))
	}

	//标题分词
	titleTran := util.ConvertTitle2Slice(title)
	//check filter
	//过滤
	for _, v := range titleTran {

		//cockoo
		if cuckoofilter.Find([]byte(v)) {
			//fmt.Println("exist")
			log.Println("cuckoo标题不合法-非法词:", v)
			end_ts := time.Now().UnixNano()
			log.Println("耗时(ns):", end_ts-start_ts)

			return true
		}
	}

	for i := 0; i < len(strFilterWords); i++ {
		//cuckFilter insert
		cuckoofilter.Del([]byte(strFilterWords[i]))
	}
	return false
}
