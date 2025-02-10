package model

import (
	"Tiktok/app/order/idl/gen"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
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



func GetOrderList(db *gorm.DB,Userid int64,Pages, Pagesize int) ([]Order,error){
	var orderList []Order
	result := db.Scopes(Paginate(int(Pages), int(Pagesize))).
	Where(Order{UserID: Userid}).Where("is_deleted = ?",0).Find(&orderList)
	if result.Error != nil {
		return nil, result.Error
	}

	return orderList,nil
}

func GenerateOrderSn(userId int32) string {
	now := time.Now()
	rand.Seed(uint64(time.Now().UnixNano()))
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(),
		userId, rand.Intn(90)+10,
	)
	return orderSn
}

func CreateOrder(db *gorm.DB, order *gen.PlaceOrderReq) (uint64, error) {
    orderNO := GenerateOrderSn(int32(order.UserId))

    newAddress := OrderAddress{
        ProvinceName: order.Address.ProvinceName,
        CityName:     order.Address.CityName,
        RegionName:   order.Address.RegionName,
        DetailAddress: order.Address.DetailAddress,
    }

    newOrderItems := make([]OrderItem, len(order.Items))
    for i, item := range order.Items {
        newOrderItems[i] = OrderItem{
            GoodsID:     int64(item.ProductsId),
            SellingPrice: int(item.ProductsPrice),
            GoodsCount:   int(item.Quantity),
        }
    }

    newOrder := Order{
        OrderNo:    orderNO,
        UserID:     int64(order.UserId),
        TotalPrice: int(order.TotalPrice),
        Address:    &newAddress,
        OrderItems: newOrderItems,
    }

    var orderID uint64

    // 使用事务保证原子性
    err := db.Transaction(func(tx *gorm.DB) error {
        // 创建订单
        if err := tx.Create(&newOrder).Error; err != nil {
            return err
        }

        // 创建订单项
        for i := range newOrderItems {
            newOrderItems[i].OrderID = newOrder.OrderID
            if err := tx.Create(&newOrderItems[i]).Error; err != nil {
                return err
            }
        }

        // 设置生成的 OrderID
        orderID = uint64(newOrder.OrderID)

        // 创建地址
        newAddress.OrderID = newOrder.OrderID
        if err := tx.Create(&newAddress).Error; err != nil {
            return err
        }

        return nil
    })

    // 返回 orderID 和可能出现的错误
    if err != nil {
        return 0, err
    }

    return orderID, nil
}




func CancelOrder(db *gorm.DB,orderId uint64) error {
	result := db.Model(&Order{}).Where("order_id = ?", orderId).Update("is_deleted", 1)
	if result.Error!= nil {
		return result.Error
	}
	return nil
}
