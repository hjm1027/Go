package main

import ("log"
	"github.com/PuerkitoBio/goquery"
)

func main() {
    doc, err := goquery.NewDocument("https://studygolang.com/topics")
	if err != nil {
	    log.Fatal(err)
	    }
	doc.Find(".topics .topic").Each(func(i int, contentSelection *goquery.Selection) {
	title := contentSelection.Find(".title a").Text()
	log.Println("第", i+1, "个帖子的标题：", title)
	 })
}
