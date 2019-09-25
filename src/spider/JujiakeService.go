package spider

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
	"time"
)

type JujiakeService struct {

}

func NewJujiakeService() *JujiakeService {
	return &JujiakeService{}
}

func (self *JujiakeService) QueryFangJia(city string, url string) ([]*House, error) {
	response, err := http.Get(url)
	time.Sleep(time.Second * 5)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer func() { err = response.Body.Close() }()
	if response.StatusCode != 200 {
		log.Println(url, err.Error())
		return nil, errors.New("not 200")
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	houses := make([]*House,0)
	doc.Find("li.list-item").Each(func(i int, s *goquery.Selection) {
		house := &House{}
		houseListTitle := self.TrimString(s.Find("a.houseListTitle").Text())
		pattern := self.TrimString(s.Find("div.details-item").Text())
		address := self.TrimString(s.Find("div.details-item").Text())
		options := self.TrimString(s.Find("div.tags-bottom").Text())
		totalPrice := self.TrimString(s.Find("span.price-det").Text())
		price := self.TrimString(s.Find("span.unit-price").Text())
		house.City = city
		house.Pattern = pattern
		house.Address = address
		house.Options = options
		house.TotalPrice = totalPrice
		house.Price = price
		house.Title = houseListTitle
		house.CreateTime = time.Now().Unix()
		house.UpdateTime = time.Now().Unix()
		houses = append(houses, house)
	})
	return houses, nil
}

func (self *JujiakeService) TrimString(s string) string {
	s = strings.Replace(s, "\n"," ", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.TrimSpace(s)
	return s
}

