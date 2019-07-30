package main

import (
    //"net/http"
    "log"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    )

func cr(url string){
    doc, err := goquery.NewDocument(url)
    if err != nil {
    	log.Fatal(err)
    }
    fmt.Println(doc)
}

func main (){
	cr("http://www.baidu.com")
}
