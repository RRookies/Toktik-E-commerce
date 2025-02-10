package model

import (
    "Tiktok/app/order/initialize"
    "Tiktok/app/order/idl/gen"
    "Tiktok/app/order/model"
    "testing"
)

func TestGetOrderList(t *testing.T) {
    initialize.InitializeTestDB()
    defer initialize.CleanupTestDB()

    // 插入测试数据
    testOrder := model.Order{
        UserID: 123,
        OrderNo: "1234567890",
        TotalPrice: 100,
        OrderStatus: 1,
        IsDeleted: false,
        OrderItems: []model.OrderItem{
            {GoodsID: 1, GoodsName: "Test Item 1", SellingPrice: 50, GoodsCount: 1},
            {GoodsID: 2, GoodsName: "Test Item 2", SellingPrice: 50, GoodsCount: 1},
        },
        Address: &model.OrderAddress{
            ProvinceName: "Test Province",
            CityName: "Test City",
            RegionName: "Test Region",
            DetailAddress: "Test Address",
        },
    }
    result := initialize.TestDB.Create(&testOrder)
    if result.Error != nil {
        t.Fatalf("failed to create test order, got error %v", result.Error)
    }

    // 调用 GetOrderList 方法
    orders, err := model.GetOrderList(initialize.TestDB,123, 1, 10)
    if err != nil {
        t.Fatalf("failed to get order list, got error %v", err)
    }

    if len(orders) != 1 {
        t.Fatalf("expected 1 order, got %d", len(orders))
    }
}

func TestCreateOrder(t *testing.T) {
    initialize.InitializeTestDB()
    defer initialize.CleanupTestDB()

    // 创建订单请求
    req := &gen.PlaceOrderReq{
        UserId: 123,
        TotalPrice: 200,
        Address: &gen.Address{
            ProvinceName: "Test Province",
            CityName: "Test City",
            RegionName: "Test Region",
            DetailAddress: "Test Address",
        },
        Items: []*gen.OrderItem{
            {ProductsId: 1, ProductsPrice: 100, Quantity: 1},
            {ProductsId: 2, ProductsPrice: 100, Quantity: 1},
        },
    }

    // 调用 CreateOrder 方法
    orderID, err := model.CreateOrder(initialize.TestDB,req)
    if err != nil {
        t.Fatalf("failed to create order, got error %v", err)
    }

    // 验证订单是否创建成功
    var order model.Order
    result := initialize.TestDB.Preload("OrderItems").Preload("Address").First(&order, orderID)
    if result.Error != nil {
        t.Fatalf("failed to find created order, got error %v", result.Error)
    }

    if order.UserID != 123 {
        t.Errorf("expected user_id to be 123, got %d", order.UserID)
    }

    if len(order.OrderItems) != 2 {
        t.Errorf("expected 2 order items, got %d", len(order.OrderItems))
    }

    if order.Address.ProvinceName != "Test Province" {
        t.Errorf("expected province_name to be 'Test Province', got '%s'", order.Address.ProvinceName)
    }
}

func TestCancelOrder(t *testing.T) {
    initialize.InitializeTestDB()
    defer initialize.CleanupTestDB()

    // 插入测试数据
    testOrder := model.Order{
        UserID: 123,
        OrderNo: "1234567890",
        TotalPrice: 100,
        OrderStatus: 1,
        IsDeleted: false,
    }
    result := initialize.TestDB.Create(&testOrder)
    if result.Error != nil {
        t.Fatalf("failed to create test order, got error %v", result.Error)
    }

    // 调用 CancelOrder 方法
    err := model.CancelOrder(initialize.TestDB,uint64(testOrder.OrderID))
    if err != nil {
        t.Fatalf("failed to cancel order, got error %v", err)
    }

    // 验证订单是否被标记为删除
    var order model.Order
    initialize.TestDB.First(&order, testOrder.OrderID)
    if !order.IsDeleted {
        t.Errorf("expected order to be marked as deleted, got is_deleted=%v", order.IsDeleted)
    }
}
