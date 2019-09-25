package storage

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"spider"
)

type StorageService struct {
	DB *gorm.DB
}

func NewStorageService(host string, port string, user string, password string) *StorageService {
	DB, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, "spider"))
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &StorageService{DB:DB}
}

func (self *StorageService) CreateHouse(houses []*spider.House) {
	for i := 0; i < len(houses); i++ {
		err := self.DB.Table("houses").Create(houses[i]).Error
		if err != nil {
			log.Println(err.Error())
		}
	}
}