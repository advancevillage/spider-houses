package main

import (
	"fmt"
	"log"
	"spider"
	"storage"
	"time"
)

type Test struct {
	x int
	y int
}

func main() {
	host := "100.100.20.36"
	port := "3306"
	user := "root"
	password := "richardsun"
	repo := storage.NewStorageService(host, port, user, password)
	jujiaku := spider.NewJujiakeService()
	for i := 1; i <= 7; i++ {
		//go func() {
			houses, err := jujiaku.QueryFangJia("武汉", fmt.Sprintf("https://wuhan.anjuke.com/sale/b142-p%d/", i))
			if err != nil {
				log.Println(i, err.Error())
			}
			if len(houses) > 0 {
				repo.CreateHouse(houses)
			}
		//}()
	}
	log.Println("wait ...")
	time.Sleep(time.Minute * 20)
}