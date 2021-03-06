package web

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/http/cookiejar"
	"sync"
)

type T667 struct {
}

func (t T667) Crawl() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if crawl("#product3") != "缺货" {
			notice("FinlandVps")
		}
	}()
	go func() {
		defer wg.Done()
		if crawl("#product4") != "缺货" {
			notice("GermanyVps")
		}
	}()

	wg.Wait()
	log.Println("T667 finshed!!!")

}
func crawl(s string) string {
	j, _ := cookiejar.New(nil)
	client := &http.Client{Jar: j}
	req, err := http.NewRequest("GET", "https://t667.com/cart.php?gid=1", nil)
	checkError(err)
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Mobile Safari/537.36")
	//req.Header.Add("cookie", `WHMCSHy6VuOzeyv1n=rbdj1huiuvmnrvs7r5v17dm4t2; __ancc_token=TiwcXNgL/QfU6E1ciz3DOw==`)
	res, err := client.Do(req)
	checkError(err)
	defer res.Body.Close()
	dom, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)
	var str string
	dom.Find(s).Find("div.price").Next().Each(func(i int, selection *goquery.Selection) {
		str = selection.Text()
	})
	return str

}
