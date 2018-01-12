package ycparser

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hopkings/crawler/parser"
)

type YcParser struct {
}

func (yp *YcParser) Parse(doc *goquery.Document) (*parser.WarehouseInfo, error) {
	//get the location
	whi := parser.CreateWarehouseInfo()
	whi.IsValid = 1
	doc.Find(".address").Each(func(i int, sel *goquery.Selection) {
		sel.Find(".see-map").Each(func(ii int, addr *goquery.Selection) {
			for _, p := range addr.Nodes {
				n := p.FirstChild
				fmt.Printf("infowarehouse: 仓库地址: %s", n.Data)
			}
		})
	})

	doc.Find(".warehouse-msg-box").Each(func(i int, sel *goquery.Selection) {
		sel.Find(".msg-name").Each(func(ii int, detail *goquery.Selection) {
			fmt.Printf(" %s", detail.Text())
		})
	})

	doc.Find(".warehouse-list-bold").Each(func(i int, ar *goquery.Selection) {
		ar.Find(".dl-left dd").Each(func(i int, area *goquery.Selection) {
			if 0 == i {
				fmt.Printf(" 仓库面积: %s", area.Text())
			}
		})
	})

	doc.Find(".warehouse-mb-bg").Each(func(i int, bs *goquery.Selection) {
		bs.Find("dl dd").Each(func(ii int, bss *goquery.Selection) {
			v := bss.Text()
			v = strings.Replace(v, " ", "", -1)
			v = strings.Replace(v, "\n", "", -1)
			switch ii {
			case 0:
				fmt.Printf(" 起租面积: %s", v)
			case 1:
				fmt.Printf(" 仓库服务: %s", v)
			case 2:
				fmt.Printf(" 仓库模式: %s", v)
			case 3:
				fmt.Printf(" 仓库类型: %s", v)
			case 4:
				fmt.Printf(" 仓库地坪: %s", v)
			case 5:
				fmt.Printf(" 子仓库数: %s", v)
			}
		})
	})

	//doc.Find(".warehouse-mb-list .warehouse-mb-bg").Each(func(i int,

	doc.Find(".warehouse-detail-message").Each(func(i int, contentSelection *goquery.Selection) {
		contentSelection.Find(".moiety").Each(func(i int, elems *goquery.Selection) {
			switch i {
			case 0:
				{
					loc := elems.Find("p").Text()
					loc = strings.Replace(loc, " ", "", -1)
					loc = strings.Replace(loc, "\n", "", -1)
					fmt.Printf(" 底层承重: %s", loc)
				}
			case 1:
				{
					diduan := elems.Find("p").Text()
					diduan = strings.Replace(diduan, " ", "", -1)
					diduan = strings.Replace(diduan, "\n", "", -1)
					fmt.Printf(" 第二层及以上承重: %s", diduan)
				}
			case 2:
				{
					class := elems.Find("p").Text()
					class = strings.Replace(class, " ", "", -1)
					class = strings.Replace(class, "\n", "", -1)
					fmt.Printf(" 公摊面积: %s", class)
				}
			case 3:
				{
					squre := elems.Find("p").Text()
					squre = strings.Replace(squre, " ", "", -1)
					squre = strings.Replace(squre, "\n", "", -1)
					fmt.Printf(" 起租期限: %s", squre)
				}
			case 4:
				{
					money := elems.Find("p").Text()
					money = strings.Replace(money, " ", "", -1)
					money = strings.Replace(money, "\n", "", -1)
					fmt.Printf(" 免租期: %s", money)
				}
			}
		})

		contentSelection.Find(".message-left-content").Each(func(i int, content *goquery.Selection) {
			if content.Find(".message-left-title").Text() == "消防安保" {
				content.Find(".ml-content-list").Each(func(iii int, child *goquery.Selection) {
					child.Find("li").Each(func(iii int, subchild *goquery.Selection) {
						anfang := strings.Replace(subchild.Text(), " ", "", -1)
						anfang = strings.Replace(anfang, "\n", "", -1)
						fmt.Printf(" %s", anfang)
					})
				})
			}
		})
	})

	fmt.Printf("\n")

	return whi, nil
}
