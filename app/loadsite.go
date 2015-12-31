package main

import (
	"fmt"
	"./web"
	"flag"
)

func main() {
	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		page, _ := web.LoadPage(arg)
		fmt.Println(page)
	}

	/*page, _ := web.LoadPage("http://nwerc.eu")
	//fmt.Println(page)

	page, _ = web.LoadPage("http://mycode.doesnot.run")

	fmt.Println(page)*/
}