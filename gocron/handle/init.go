// Package handle init
// Created by chenguolin 2018-11-17
package handle

var (
	crawlAddressTxsHan *CrawlAddressTxsHandle
)

// built-in init function
func init() {
	crawlAddressTxsHan = newCrawlAddressTxsHandle()
}

// GetCrawlAddressTxsHandle get CrawlAddressTxsHandle
func GetCrawlAddressTxsHandle() *CrawlAddressTxsHandle {
	return crawlAddressTxsHan
}
