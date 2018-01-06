package parser58

import (
	//"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hopkings/crawler/parser"
)

type Parser58 struct {
}

func (p5 *Parser58) Parse(doc *goquery.Document) (*parser.WarehouseInfo, error) {
	//get the location
	location_content := ""
	hasLoc := 0
	whi := parser.CreateWarehouseInfo()
	doc.Find("head").Each(func(i int, sel *goquery.Selection) {
		sel.Find("meta").Each(func(ii int, metas *goquery.Selection) {
			for _, n := range metas.Nodes {
				for _, attr := range n.Attr {
					if attr.Key == "name" && attr.Val == "location" {
						hasLoc = 1
						continue
					}
					if attr.Key == "content" {
						location_content = attr.Val
						break
					}
				}
				if 1 == hasLoc {
					break
				}
			}
		})
	})

	if 1 == hasLoc {
		whi.Location = location_content
		whi.IsValid = 1
	}

	doc.Find(".info").Each(func(i int, contentSelection *goquery.Selection) {
		contentSelection.Find("li").Each(func(i int, elems *goquery.Selection) {
			switch i {
			case 0:
				{
					loc := elems.Find("a").Text()
					loc = strings.Replace(loc, " ", "", -1)
					loc = strings.Replace(loc, "\n", "", -1)
					whi.Region = loc
				}
			case 1:
				{
					if len(elems.Nodes) > 0 && elems.Nodes[0].LastChild != nil {
						diduan := elems.Nodes[0].LastChild.Data
						diduan = strings.Replace(diduan, " ", "", -1)
						diduan = strings.Replace(diduan, "\n", "", -1)
						whi.District = diduan
					}
				}
			case 2:
				{
					if len(elems.Nodes) > 0 && elems.Nodes[0].LastChild != nil {
						class := elems.Nodes[0].LastChild.Data
						class = strings.Replace(class, " ", "", -1)
						class = strings.Replace(class, "\n", "", -1)
						whi.Class = class
					}
				}
			case 3:
				{
					if len(elems.Nodes) > 0 && elems.Nodes[0].LastChild != nil {
						squre := elems.Nodes[0].LastChild.Data
						squre = strings.Replace(squre, " ", "", -1)
						squre = strings.Replace(squre, "\n", "", -1)
						whi.Square = squre
					}
				}
			case 4:
				{
					money := elems.Find("em").Text()
					whi.Money = money
				}
			}
		})
	})

	return whi, nil
}
