package storage

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"spider"
	"strings"
)

type StorageService struct {
	DB *gorm.DB
}

type StorageRedisService struct {
	Client *redis.Client
}

func NewStorageService(host string, port string, user string, password string) *StorageService {
	DB, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, "spider"))
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &StorageService{DB:DB}
}

func NewStorageRedisService(host string, port string, password string, db int) *StorageRedisService {
	Client := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%s", host, port), Password: password, DB: db})
	pong, err := Client.Ping().Result()
	if err != nil {
		log.Println(pong, err.Error())
	}
	return &StorageRedisService{Client:Client}
}

func (self *StorageService) CreateHouse(houses []*spider.House) {
	values := make([]string, 0, len(houses))
	for i:= 0; i < len(houses); i++ {
		value := fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', '%s', %d, %d, %d)", houses[i].Town, houses[i].Area,houses[i].Room,houses[i].TotalPrice, houses[i].Price, houses[i].Address, houses[i].Options, houses[i].CreateTime, houses[i].UpdateTime, houses[i].Page)
		values = append(values, value)
	}
	sql := "insert into houses(town, area, room, totalPrice, price, address, options, createTime, updateTime, page)values" +
		strings.Replace(strings.Trim(fmt.Sprint(values), "[]"), ") (", "),(", -1) +
		";"
	err := self.DB.Exec(sql).Error
	if err != nil {
		log.Println(err.Error())
	}
}