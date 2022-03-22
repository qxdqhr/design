package service

import "github.com/jinzhu/gorm"

type Juice struct {
	gorm.Model
	JuiceId     string `json:"juice_id"`
	JuiceName string `json:"juice_name"`
	JuiceType string `json:"juice_type"`
	LastOrderingTime string `json:"last_ordering_time"`
	Price string `json:"price"`
	Profit string `json:"profit"`
	Cost string `json:"cost"`
	CurEvaluate string `json:"cur_evaluate"` //最近评价
	JuiceSoldNumber string `json:"juice_sold_number"`
	SellingTotalPrice string `json:"selling_total_price"`//历史总销售额
	GoodEvaluateNum string `json:"good_evaluate_num"`//好评数
}