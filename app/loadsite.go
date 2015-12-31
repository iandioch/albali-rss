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
		fmt.Println(data)
		fmt.Println(data.ChannelTitle)
	}


}