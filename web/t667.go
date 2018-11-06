package web

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
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
			http.Get("https://sc.ftqq.com/SCU6677T8dbadb7d61c94c3acaf6df175948d2b858c8e9d9be2fb.send?text=FinlandVps")
		}
	}()
	go func() {
		defer wg.Done()
		if crawl("#product4") != "缺货" {
			http.Get("https://sc.ftqq.com/SCU6677T8dbadb7d61c94c3acaf6df175948d2b858c8e9d9be2fb.send?text=GermanyVps")
		}
	}()

	wg.Wait()
	log.Println("T667 finshed!!!")

}
func crawl(s string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://t667.com/cart.php?gid=1", nil)
	checkError(err)
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Mobile Safari/537.36")
	req.Header.Add("cookie", `WHMCSHy6VuOzeyv1n=rbdj1huiuvmnrvs7r5v17dm4t2; __ancc_token=1f36dGLuGy85cvuhr0Ed7A==`)
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
func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
