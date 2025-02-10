package model

import (
	"Tiktok/app/order/global"
	"time"
)

type Order struct {
	BaseModel
	OrderID     int64       `gorm:"primaryKey;autoIncrement" json:"order_id"`
	OrderNo     string      `gorm:"type:varchar(20);not null;default:''" json:"order_no"`
	UserID      int64       `gorm:"not null;default:0" json:"user_id"`
	TotalPrice  int         `gorm:"not null;default:1" json:"total_price"`
	PayStatus   int8        `gorm:"not null;default:0" json:"pay_status"`
	PayType     int8        `gorm:"not null;default:0" json:"pay_type"`
	PayTime     time.Time   `json:"pay_time"`
	OrderStatus int8        `gorm:"not null;default:0" json:"order_status"`
	ExtraInfo   string      `gorm:"type:varchar(100);not null;default:''" json:"extra_info"`
	IsDeleted   bool        `gorm:"not null;default:0" json:"is_deleted"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order_items"`
	Address     *OrderAddress `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"address"`
}


type OrderAddress struct {
	BaseModel
	OrderID       int64  `gorm:"uniqueIndex;not null" json:"order_id"`
	ProvinceName  string `gorm:"type:varchar(32);not null;default:''" json:"province_name"`
	CityName      string `gorm:"type:varchar(32);not null;default:''" json:"city_name"`
	RegionName    string `gorm:"type:varchar(32);not null;default:''" json:"region_name"`
	DetailAddress string `gorm:"type:varchar(64);not null;default:''" json:"detail_address"`
}



type OrderItem struct {
	BaseModel
	OrderItemID   int64  `gorm:"primaryKey;autoIncrement" json:"order_item_id"`
	OrderID       int64  `gorm:"not null;default:0" json:"order_id"`
	GoodsID       int64  `gorm:"not null;default:0" json:"goods_id"`
	GoodsName     string `gorm:"type:varchar(200);not null;default:''" json:"goods_name"`
	GoodsCoverImg string `gorm:"type:varchar(200);not null;default:''" json:"goods_cover_img"`
	SellingPrice  int    `gorm:"not null;default:1" json:"selling_price"`
	GoodsCount    int    `gorm:"not null;default:1" json:"goods_count"`

	Order         Order  `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}


func GetOrderList(Userid int64,Pages, Pagesize int) ([]Order,error){
	var orderList []Order
	result := global.DB.Scopes(Paginate(int(Pages), int(Pagesize))).
	Where(Order{UserID: Userid}).Find(&orderList)
	if result.Error != nil {
		return nil, result.Error
	}

	return orderList,nil
}