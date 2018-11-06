package main

import "github.com/hswmartin/vps/web"

func main() {
	var t667, letbox web.Web
	t667 = new(web.T667)
	ch := make(chan int)
	go func() {
		t667.Crawl()
		ch <- 1
	}()
	letbox = new(web.Letbox)
	letbox.Crawl()
	<-ch
}
