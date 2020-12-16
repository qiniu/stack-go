package bss

import (
	"time"

	"github.com/qiniu/stack-go/params"
)

// OrderInfo 订单信息
type OrderInfo struct {
	Status       string    `json:"status"`
	CreateTime   time.Time `json:"create_time"`
	CompleteTime time.Time `json:"complete_time"`
	CancelTime   time.Time `json:"cancel_time"`
	PayTime      time.Time `json:"pay_time"`
	Order        struct {
		OrderDetail struct {
			ID          int       `json:"id"`
			Order       string    `json:"order"`
			Fee         float64   `json:"fee"`
			CFee        float64   `json:"c_fee"`
			UpdateTime  time.Time `json:"update_time"`
			CreateTime  time.Time `json:"create_time"`
			PayTime     time.Time `json:"pay_time"`
			ExpiredTime time.Time `json:"expired_time"`
			Status      int       `json:"status"`
		} `json:"bo_order"`
		OrderTag string `json:"order_tag"`
	} `json:"order"`
	Resources []struct {
		OrderType      string              `json:"order_type"`
		ResourceType   params.ResourceType `json:"resource_type"`
		ResourceID     string              `json:"resource_id"`
		RegionID       string              `json:"region"`
		Price          float64             `json:"price"`
		Status         string              `json:"status"`
		LifeStartTime  time.Time           `json:"life_start_time"`
		LifeExpireTime time.Time           `json:"life_expire_time"`
	} `json:"resource_order_infos"`
}
