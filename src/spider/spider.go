package spider

type Spider interface {
	//@brief: 查询房价接口
	//@Param: town 城市
	//@Param: area 城区
	//@Param: room 格局
	//@Param: URL
	QueryFangJia(town string, area string, room string, page int, url string) ([]*House, error)
	UrlFormat(city string, area string, room string, page int) string
}


type House struct {
	ID      int 		`gorm:"column:id;type:int;auto_increment"`
	Town 	string   	`gorm:"column:town;type:varchar(255);not null"`	//城市
	Area    string 		`gorm:"column:area;type:varchar(255);not null"`	//卖点
	Room    string 		`gorm:"column:room;type:varchar(255);not null"`	//房型
	Address string		`gorm:"column:address;type:varchar(255);not null"`	//小区
	Options string 		`gorm:"column:options;type:varchar(255);not null"`	//卖点
	TotalPrice string 	`gorm:"column:totalPrice;type:varchar(255);not null"`//总价
	Price 	string 		`gorm:"column:price;type:varchar(255);not null"`	//单价
	CreateTime int64 	`gorm:"column:createTime;type:int(64);not null"`	//创建时间
	UpdateTime int64 	`gorm:"column:updateTime;type:int(64);not null"`	//更新时间
	Page 	int 		`gorm:"column:page;type:int;not null"`
}

type Town struct {
	Key   string  //ID
	Name  string  //名称
	Func  Spider //接口
	Areas []*Area
}

type Area struct {
	Key   string //ID
	Name  string //名称
	Rooms []*RoomType
}

type RoomType struct {
	Key  string  //ID
	Name string  //名称
}