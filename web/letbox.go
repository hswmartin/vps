package web

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type Letbox struct {
}

func (l Letbox) Crawl() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://my.letbox.com/cart.php?a=add&pid=67", nil)
	checkError(err)
	req.Header.Add("cookie", `WHMCSAIWQpkmQ81cw=8b74436e9c26ba8b2ce8dcce946eeee5`)

	res, err := client.Do(req)

	checkError(err)
	defer res.Body.Close()
	dom, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)
	s := dom.Find("h1").Text()

	if s != "Out of Stock" {
		notice("letbox")
	}
	log.Println("letbox finshed!!!")
}
