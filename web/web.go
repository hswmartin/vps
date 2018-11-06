package web

import (
	"log"
	"net/http"
)

type Web interface {
	Crawl()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func notice(s string) {
	http.Get("https://sc.ftqq.com/SCU6677T8dbadb7d61c94c3acaf6df175948d2b858c8e9d9be2fb.send?text=" + s)
}
