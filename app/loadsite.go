package main

import (
	"fmt"
	"./web"
	"flag"
	"./parse/xml"
)

func main() {
	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		page, _ := web.LoadPage(arg)
		//fmt.Println(page)
		data, err := xml.Parse(page)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data.ChannelTitle)
		fmt.Println(data.ChannelDescription)

		for _, item := range data.Items {
			fmt.Println(item.Title)
		}
	}


}