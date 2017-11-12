package cuckoo

import (
	"log"
	"strings"
	"time"
	"xc/myGit/go_filter_test/util"

	cuckoo "github.com/goCuckoo"
)

func PoolFilter(title string, filterWords string) bool {
	start_ts := time.Now().UnixNano() / 1000000
	//remove ,
	strFilterWords := strings.Split(filterWords, ",")
	log.Printf("过滤词：%v", strFilterWords)

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
			log.Println("cuckoo标题不合法-非法词:%s", v)
			return true
		}
	}

	for i := 0; i < len(strFilterWords); i++ {
		//cuckFilter insert
		cuckoofilter.Del([]byte(strFilterWords[i]))
	}
	end_ts := time.Now().UnixNano() / 1000000
	log.Println("耗时(ms):", end_ts-start_ts)

	return false
}
