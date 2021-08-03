package main

import (
	"encoding/json"
	"fmt"
	"time"
)


func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
func test() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	fmt.Printf("该函数执行完成耗时：%+v", elapsed)
}
func DayBeginStamp(now int64) int64 {

	_, offset := time.Now().Zone()
	//fmt.Println(zone, offset)
	return now - (now+int64(offset)) % int64(3600 * 24)
	//return (now + int64(offset))/int64(3600 * 24) * int64(3600 * 24) - int64(offset)

}

type OriginalReq struct {
	Msg struct {
		Trade struct {
			PromotionDetails struct {
				PromotionDetail []PromotionDetail `json:"promotion_detail"`
			} `json:"promotion_details"`
			Modified string `json:"modified"`
		} `json:"trade"`
	} `json:"msg"`
}

type PromotionDetail struct {
	DiscountFee   string `json:"discount_fee"`
	ID            string `json:"id"`
	PromotionDesc string `json:"promotion_desc"`
	PromotionID   string `json:"promotion_id"`
	PromotionName string `json:"promotion_name"`
}

func main() {
	//daystr := os.Args[1]
	//tm, _ := time.ParseInLocation("2006-01-02 15:04:05", daystr, time.Local)
	//ts := CalValidBindts(tm.Unix())
	//year, month, _ := time.Now().Date()
	//thisMonth := time.Date(year,month,1,0,0,0,0,time.Local)
	//date := thisMonth.AddDate(0, -6, 0)
	//start := date.Format("2006-01-02 15:04:05")
	//fmt.Println(start)
	//fmt.Println(date.Unix())
	//fmt.Println(time.Date(year,month,1,0,0,0,0,time.Local).Format("2006-01"))
	//test()
	//fmt.Println()
	////stamp := DayBeginStamp(time.Now().Unix())
	//var count int64 = 8
	// k := float64(count) * 0.2
	// j := int32(k)
	//fmt.Printf("===：%+v , %+v, %+v", k,int32(math.Round(k)),j)
	//fmt.Println()
	originalReq := "{\"trade\":{\"modified\":\"2021-07-14 11:00:34\",\"created\":\"2021-07-14 11:00:28\",\"new_presell\":false,\"num\":1,\"credit_card_fee\":\"1.00\",\"orders\":{\"order\":[{\"adjust_fee\":\"0.00\",\"buyer_rate\":false,\"cid\":50023728,\"discount_fee\":\"0.00\",\"is_daixiao\":false,\"is_oversold\":false,\"nr_outer_iid\":\"testkid_328975537580124\",\"num\":1,\"num_iid\":647641059623,\"oid\":\"1948423070178064731\",\"oid_str\":\"1948423070178064731\",\"order_from\":\"WAP,WAP\",\"outer_iid\":\"testkid_328975537580124\",\"part_mjz_discount\":\"200.00\",\"payment\":\"1.00\",\"price\":\"201.00\",\"refund_status\":\"NO_REFUND\",\"seller_rate\":false,\"seller_type\":\"B\",\"snapshot_url\":\"t:1948423070178064731_1\",\"status\":\"WAIT_SELLER_SEND_GOODS\",\"title\":\"测试商品1：测试天猫下单正常首购套餐；购买无效\",\"total_fee\":\"201.00\",\"fqg_num\":0,\"is_fqg_s_fee\":false,\"divide_order_fee\":\"\"}]},\"payment\":\"1.00\",\"promotion_details\":{\"promotion_detail\":null},\"receiver_mobile\":\"18501232436\",\"status\":\"WAIT_SELLER_SEND_GOODS\",\"tid\":\"1948423070178064731\",\"tid_str\":\"1948423070178064731\",\"title\":\"伴鱼旗舰店\",\"you_xiang\":false}}"
	fun := "ParsingOriginalReq -->"
	var info OriginalReq
	if err := json.Unmarshal([]byte(originalReq), &info); err != nil {
		fmt.Printf( "%s json unmarshal failed err:%s", fun, err.Error())
	}
	fmt.Printf("%s originalReq :%+v", fun, info)

}
