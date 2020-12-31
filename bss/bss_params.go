package bss

import (
	"time"

	"github.com/qiniu/stack-go/params"
)

// OrderInfo 订单信息
type OrderInfo struct {

	// 订单状态
	// - `new`:      新建状态
	// - `paid`:     已经支付状态
	// - `canceled`: 被取消状态
	// - `complete`: 订单下所有子订单都已经完成
	// - `refund`:   用户订单执行了退款
	Status string `json:"status"`

	CreateTime   time.Time `json:"create_time"`   // 订单创建时间
	CompleteTime time.Time `json:"complete_time"` // 订单完成时间
	CancelTime   time.Time `json:"cancel_time"`   // 订单取消时间
	PayTime      time.Time `json:"pay_time"`      // 订单支付时间

	Order struct {
		OrderDetail struct {
			ID          int       `json:"id"`           // ID
			Order       string    `json:"order"`        // 订单号
			Fee         float64   `json:"fee"`          // 订单原价
			CFee        float64   `json:"c_fee"`        // 订单实际价格
			CreateTime  time.Time `json:"create_time"`  // 创建时间
			UpdateTime  time.Time `json:"update_time"`  // 更新时间
			PayTime     time.Time `json:"pay_time"`     // 支付时间
			ExpiredTime time.Time `json:"expired_time"` // 订单过期时间
			Status      int       `json:"status"`       // 订单状态 1: 未支付 2: 已支付 3:作废, 4: 垫付，5: 后付费
		} `json:"bo_order"`

		OrderTag string `json:"order_tag"` // 订单标签
	} `json:"order"`

	Resources []struct {

		// 订单类型
		//
		// - `create`:              新建
		// - `renew`:               续费
		// - `auto_renew`:          自动续费
		// - `to_prepaid`:          按需转包月
		// - `upgrade`:             升级
		// - `replace_system_disk`: 更换系统盘
		// - `to_prepaid_disk`:     按需磁盘转包月
		// - `renew_with_spec`:     续费变配
		OrderType string `json:"order_type"`

		// 资源类型
		ResourceType params.ResourceType `json:"resource_type"`

		// 资源 ID
		ResourceID string `json:"resource_id"`

		// 地域
		RegionID string `json:"region"`

		// 金额
		Price float64 `json:"price"`

		// 状态
		//
		// - `non_executable`: 不可执行状态
		// - `executable`:     可执行状态
		// - `processing`:     正在执行中
		// - `success`:        执行成功
		// - `failed`:         执行失败
		// - `fail_over`:      失败次数过多
		// - `cancelled`:      被取消
		// - `refunded`:       已退款
		Status string `json:"status"`

		LifeStartTime  time.Time `json:"life_start_time"`  // 计费周期
		LifeExpireTime time.Time `json:"life_expire_time"` // 计费周期
	} `json:"resource_order_infos"`
}
