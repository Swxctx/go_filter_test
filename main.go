package main

import (
	"log"
	"time"
	"xc/myGit/go_filter_test/bloom"
	"xc/myGit/go_filter_test/cuckoo"
	"xc/myGit/go_filter_test/usestring"
)

func main() {
	filterWords := "广告,裤袜,不会,大家,我们,在一起,最高级,最大,最小,全世界,全国,全部的,最小的,最好的,免费的,包好,全部免费,平台"
	// title := "全城疯抢,1元1GB还有腾讯VIP"
	// title := "屠龙一刀999"
	// title := "泰国3D眉笔眉粉染眉膏三合一防水防汗"
	title := "网易新闻 - 头条视频资讯阅读平台"
	// log.Println("blandcuk")
	// log.Println("flag:", blandcuk.PoolFilter(title))

	log.Println("bloom")
	log.Println("flag:", bloom.PoolFilter(title, filterWords))
	log.Println()

	log.Println("cuckoo")
	log.Println("flag:", cuckoo.PoolFilter(title, filterWords))
	log.Println()

	// strings
	log.Println("use_strings")
	// contains
	log.Println("contains")
	log.Println("flag:", usestring.PoolFilterContains(title, filterWords))
	log.Println()
	// containsAny
	log.Println("containsAny")
	log.Println("flag:", usestring.PoolFilterContainsAny(title, filterWords))
	log.Println()
	// indexAny
	log.Println("indexAny")
	log.Println("flag:", usestring.PoolFilterIndexAny(title, filterWords))
	log.Println()
}

func getTime() int64 {
	return time.Now().UnixNano() / 1000000
}
