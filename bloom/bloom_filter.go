package bloom

import (
	"log"
	"strings"
	"time"
	"xc/myGit/go_filter_test/util"

	"github.com/bloom"
)

func PoolFilter(title string, filterWords string) bool {
	start_ts := time.Now().UnixNano() / 1000000
	//remove ,
	strFilterWords := strings.Split(filterWords, ",")
	log.Println("过滤词:", strFilterWords)

	//bloom filter
	n := uint(1000)
	filter := bloom.New(20*n, uint(len(strFilterWords))) // load of 20, 5 keys

	//add filter
	for i := 0; i < len(strFilterWords); i++ {
		//bloom filter add
		filter.Add([]byte(strFilterWords[i]))
	}

	titleTran := util.ConvertTitle2Slice(title)
	for _, v := range titleTran {
		//bloom
		if filter.Test([]byte(v)) {
			log.Println("bloom标题不合法-非法词:", v)
			end_ts := time.Now().UnixNano() / 1000000
			log.Println("耗时(ms):", end_ts-start_ts)
			return true
		}
	}
	return false
}
