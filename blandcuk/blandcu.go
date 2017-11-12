package blandcuk

import (
	"log"
	"strings"
	"xc/myGit/go_filter_test/util"

	"github.com/bloom"
	cuckoo "github.com/goCuckoo"
)

func PoolFilter(title string, filterWords string) bool {

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

	titleTran := util.ConvertTitle2Slice(title)
	//check filter
	//过滤
	for _, v := range titleTran {
		//cockoo
		if cuckoofilter.Find([]byte(v)) {
			log.Println("cuckoo标题不合法-非法词:%s", v)
			return true
		}

		//bloom
		if filter.Test([]byte(v)) {
			log.Println("bloom标题不合法-非法词:%s", v)
			return true
		}
	}

	return false
}
