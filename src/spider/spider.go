package spider

type Spider interface {
	QueryFangJia(city string, url string) ([]*House, error)
}


type House struct {
	ID      int 		`gorm:"column:id;type:int;auto_increment"`
	City 	string   	`gorm:"column:city;type:varchar(255);not null"`	//城市
	Title   string 		`gorm:"column:title;type:varchar(255);not null"`	//卖点
	Pattern string 		`gorm:"column:pattern;type:varchar(255);not null"`	//房型
	Address string		`gorm:"column:address;type:varchar(255);not null"`	//小区
	Options string 		`gorm:"column:options;type:varchar(255);not null"`	//卖点
	TotalPrice string 	`gorm:"column:totalPrice;type:varchar(255);not null"`//总价
	Price 	string 		`gorm:"column:price;type:varchar(255);not null"`	//单价
	CreateTime int64 	`gorm:"column:createTime;type:int(64);not null"`	//创建时间
	UpdateTime int64 	`gorm:"column:updateTime;type:int(64);not null"`	//更新时间
}