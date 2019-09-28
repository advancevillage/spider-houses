package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	url2 "net/url"
	"strings"
	"time"
)

type JujiakeService struct {

}

func NewJujiakeService() *JujiakeService {
	return &JujiakeService{}
}

func (self *JujiakeService) QueryFangJia(town string, area string, room string, page int, url string) ([]*House, error) {
	//从IP池获取IP
	xUrl := "http://www.xicidaili.com/wt/"
	xRequest, err := http.NewRequest("GET", xUrl, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	xRequest.Header.Add("User-Agent", self.Agent())
	xClient := &http.Client{}
	xResponse, err := xClient.Do(xRequest)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	proxyPool := make([]string, 0)
	dom, err := goquery.NewDocumentFromReader(xResponse.Body)
	dom.Find(".odd").Each(func(i int, context *goquery.Selection) {
		ip := context.Find("td").Eq(1).Text()
		port := context.Find("td").Eq(2).Text()
		httpType := context.Find("td").Eq(5).Text()
		proxyIp := strings.ToLower(httpType) + "://" + ip + ":" + port
		//nim := context.Find("td").Eq(4).Text() //是否是高匿,高匿的可以隐藏你的原始IP
		proxyPool = append(proxyPool, proxyIp)
	})
	r := rand.New(rand.NewSource(time.Now().Unix()))
	proxy, err := url2.Parse(proxyPool[r.Intn(len(proxyPool))])
	//
	client := http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	request, err :=  http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	request.Header.Add("User-Agent", self.Agent())
	//request.Header.Add("origin", "https://wuhan.anjuke.com")
	//request.Header.Add("referer", url)
	request.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	request.Header.Add("accept-encoding", "gzip, deflate, br")
	request.Header.Add("cache-control", "no-cache")
	request.Header.Add("sec-fetch-mode", "navigate")
	request.Header.Add("sec-fetch-site", "none")
	request.Header.Add("sec-fetch-user", "?1")
	request.Header.Add("upgrade-insecure-requests", "1")
	request.Header.Add("cookie", "sessid=FB2E0FE3-2017-4D3A-22D3-923708D31810; aQQ_ajkguid=DE11B51E-1FCF-3BF7-F7B0-D9376100AF16; lps=http%3A%2F%2Fwww.anjuke.com%2F%7Chttps%3A%2F%2Fwww.google.com%2F; twe=2; _ga=GA1.2.977696046.1569423652; 58tj_uuid=1506c123-5f99-4b77-a8e4-7e464e7010c8; als=0; ctid=22; wmda_uuid=a6317853f130ba7b4e73b250661ea212; wmda_new_uuid=1; wmda_visited_projects=%3B6289197098934; ajk_member_captcha=9c69d59894aef187364ce977c3e59bd0; _gid=GA1.2.1616239134.1569569740; isp=true; Hm_lvt_c5899c8768ebee272710c9c5f365a6d8=1569607130; Hm_lpvt_c5899c8768ebee272710c9c5f365a6d8=1569630627; browse_comm_ids=639328%7C847342; propertys=uaytnc-pyin6r_uk6qgl-pye7ts_; init_refer=; new_uv=7; wmda_session_id_6289197098934=1569654432932-a74f4e73-6613-9328; new_session=0; _gat=1")
	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer func() { err = response.Body.Close() }()
	if response.StatusCode != 200 {
		body, err := ioutil.ReadAll(response.Body)
		log.Println(err, string(body))
		return nil, errors.New("not 200")
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	houses := make([]*House,0)
	doc.Find("li.list-item").Each(func(i int, s *goquery.Selection) {
		house := &House{}
		address := self.TrimString(s.Find("div.details-item").Text())
		options := self.TrimString(s.Find("div.tags-bottom").Text())
		totalPrice := self.TrimString(s.Find("span.price-det").Text())
		price := self.TrimString(s.Find("span.unit-price").Text())
		house.Town = town
		house.Area = area
		house.Room = room
		house.Address = address
		house.Options = options
		house.TotalPrice = totalPrice
		house.Price = price
		house.Page = page
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

func (self *JujiakeService) UrlFormat(city string, area string, room string, page int) string {
	return fmt.Sprintf("https://%s.anjuke.com/sale/%s/%s-p%d/", city, area, room, page)
}


func (self *JujiakeService) Agent() string {
	agent  := [...]string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := len(agent)
	return agent[r.Intn(len)]
}