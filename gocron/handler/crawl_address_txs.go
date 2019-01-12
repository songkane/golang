// Package handler 定时抓取某个地址的交易记录
// Created by chenguolin 2018-09-27
package handler

import (
	"fmt"

	"github.com/antchfx/htmlquery"
	"gitlab.local.com/golang/golog"
)

type tx struct {
	TxHash string `json:"tx_hash"`
	From   string `json:"from"`
}

// CrawlAddressTxsHandler 抓取某个地址的交易记录
func CrawlAddressTxsHandler() {
	maxPage := 10
	page := 1
	for page <= maxPage {
		// 1. 数据抓取
		txs, err := crawlTxs(page)
		if err != nil {
			page++
			continue
		}
		if len(txs) <= 0 {
			return
		}
		fmt.Println(txs)
		page++
	}
}

// crawlTxs 从etherscan抓取交易
// 0xc77aa121eb48ac3fda90cca5c65317a82ccd0f4a 以太坊地址
// https://etherscan.io/address/0xc77aa121eb48ac3fda90cca5c65317a82ccd0f4a
func crawlTxs(page int) ([]*tx, error) {
	txs := make([]*tx, 0)
	url := fmt.Sprintf("https://etherscan.io/txs?a=0xc77aa121eb48ac3fda90cca5c65317a82ccd0f4a&p=%d", page)
	pageDoc, err := htmlquery.LoadURL(url)
	if err != nil {
		golog.Error("[crawl_address_txs - crawlTxs] htmlquery.LoadURL, error:",
			golog.String("URL", url),
			golog.Object("Error", err))
		return txs, err
	}
	if pageDoc == nil {
		return txs, nil
	}

	// 解析出所有的入账交易
	txItems := htmlquery.Find(pageDoc, `//div[@class="table-responsive"]/table/tbody/tr`)
	if len(txItems) == 0 {
		golog.Warn("[crawl_address_txs - crawlTxs] pageDoc not found any tx",
			golog.String("URL", url))
		return txs, nil
	}
	for _, item := range txItems {
		// extract txhash
		text := htmlquery.Find(item, `/td[1]/span/a/text()`)
		if len(text) <= 0 {
			continue
		}
		txHash := text[0].Data

		// extract from
		text = htmlquery.Find(item, `/td[4]/span/a/text()`)
		if len(text) <= 0 {
			continue
		}
		from := text[0].Data

		// new tx
		newTx := &tx{
			TxHash: txHash,
			From:   from,
		}
		txs = append(txs, newTx)
	}
	return txs, nil
}
