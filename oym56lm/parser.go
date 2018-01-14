package oym56lm

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	//"github.com/axgle/mahonia"
	"github.com/hopkings/crawler/parser"
)

type OymParser struct {
}

func (yp *OymParser) Parse(doc *goquery.Document) (*parser.WarehouseInfo, error) {
	//get the location
	whi := parser.CreateWarehouseInfo()
	whi.IsValid = 1

	doc.Find(".dt_boxsec .first").Each(func(i int, sel *goquery.Selection) {
		fmt.Printf("infowarehouse: %s ", sel.Text())
		/*for _, p := range addr.Nodes {
			n := p.FirstChild
			fmt.Printf("warehouse: 仓库地址: %s", n.Data)
		}*/
	})

	doc.Find(".midbox dl").Each(func(i int, sel *goquery.Selection) {
		data := strings.Replace(sel.Text(), "\n", " ", -1)
		data = strings.Replace(data, "\t", " ", -1)
		fmt.Printf("%s ", data)
	})

	doc.Find(".profile .floortwo p").Each(func(i int, sel *goquery.Selection) {
		data := strings.Replace(sel.Text(), "\n", " ", -1)
		data = strings.Replace(data, "\t", " ", -1)
		fmt.Printf("%s ", data)
	})

	doc.Find(".basic_information .basic_info_table").Each(func(i int, sel *goquery.Selection) {
		data := strings.Replace(sel.Text(), "\n", " ", -1)
		data = strings.Replace(data, "\t", " ", -1)
		fmt.Printf("%s ", data)
	})

	fmt.Printf("\n")

	return whi, nil
}
