package main

import (
	"fmt"
	"golang.org/x/net/context"
	"spider"
	"storage"
	"time"
)

type Test struct {
	x int
	y int
}

func main() {
	host := "100.100.25.66"
	port := "6379"
	password := ""
	repo := storage.NewStorageRedisService(host, port, password, 0)
	jujiaku := spider.NewJujiakeService()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute * 30)
	for i := 1; i <= 7; i++ {
		//go func(ctx context.Context) {
			houses, err := jujiaku.QueryFangJia("武汉", fmt.Sprintf("https://wuhan.anjuke.com/sale/hongshana/b142-p%d/", i))
			if err != nil {
				<- ctx.Done()
			}
			if len(houses) > 0 {
				repo.CreateHouse(houses)
			} else {
				<- ctx.Done()
			}
		//}(ctx)
	}
	fmt.Println("finish done...")
	cancel()
}