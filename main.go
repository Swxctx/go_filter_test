package main

import (
	"go_filter_test/bloom"
	"go_filter_test/cuckoo"
	"go_filter_test/userune"
	"go_filter_test/usestring"
	"log"
	"time"
)

func main() {
	filterWords := "广告,电脑,不会,大家,我们,在一起,最高级,最大,最小,全世界,全国,全部的,最小的,最好的,免费的,包好,全部免费,微爱,"
	// filterWords := "微爱"
	// title := "屠龙一刀999"
	// title := "金利来袜子男士纯棉夏季商务男袜中筒袜黑短"
	// title := "网易新闻 - 头条视频资讯阅读平台"
	title := "网易新闻 - 头条视频资讯阅读平台免费的"
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

	// rune
	log.Println("rune range")
	log.Println("flag:", userune.PoolFilterRune(title, filterWords))
	log.Println()

	//bloom and contains
	log.Println("Bloom And Contains")
	log.Println("flag:", usestring.FilterBloomAndContains(title, filterWords))

}

func getTime() int64 {
	return time.Now().UnixNano()
}
