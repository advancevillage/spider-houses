package main

import (
	"fmt"
	"spider"
	"storage"
	"sync"
)

type Test struct {
	x int
	y int
}

func main() {
	host := "192.168.1.101"
	port := "3306"
	user := "root"
	password := "password"
	repo := storage.NewStorageService(host, port, user, password)
	jujiaku := spider.NewJujiakeService()
	group := sync.WaitGroup{}
	//城市-->多城区-->多格局
	src := []*spider.Town{
		&spider.Town{
			Key:  "wuhan",
			Name: "武汉",
			Func: jujiaku,
			Areas: []*spider.Area{
				&spider.Area{
					Key: "hongshana",
					Name: "洪山区",
					Rooms:[]*spider.RoomType{
						{Key: "b142", Name:"2室"},
					},
				},
				&spider.Area{
					Key: "dongxihu",
					Name: "东西湖",
					Rooms:[]*spider.RoomType{
						{Key: "b142", Name:"2室"},
					},
				},
			},
		},
	}
	for c := 0; c < len(src); c++ {
		cityName := src[c].Name
		cityKey := src[c].Key
		for a := 0; a < len(src[c].Areas); a++ {
			areaName := src[c].Areas[a].Name
			areaKey := src[c].Areas[a].Key
			for r := 0; r <len(src[c].Areas[a].Rooms); r++ {
				roomName := src[c].Areas[a].Rooms[r].Name
				roomKey := src[c].Areas[a].Rooms[r].Key
				for p := 1; p <= 10; p++ {
					group.Add(1)
					//go func(ic, ia, ir, ip int) {
					//	time.Sleep(1 * time.Second)
						//log.Println(ic, ia, ir, ip)
						houses, err := src[c].Func.QueryFangJia(cityName, areaName, roomName, p, src[c].Func.UrlFormat(cityKey, areaKey, roomKey, p))
						if err != nil {
							group.Done()
						} else if len(houses) > 0 {
							repo.CreateHouse(houses)
							group.Done()
						} else {
							group.Done()
						}
					//}(c, a, r, p)
				}
			}

		}
	}
	group.Wait()
	fmt.Println("finish done...")
}