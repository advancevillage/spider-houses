package storage

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"spider"
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
	for i := 0; i < len(houses); i++ {
		err := self.DB.Table("houses").Create(houses[i]).Error
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (self *StorageRedisService) CreateHouse(houses []*spider.House) {
	for i := 0; i < len(houses); i++ {
		err := self.Client.LPush(houses[i].City, houses[i]).Err()
		if err != nil {
			log.Println(err.Error())
		}
	}
}