package main

import (
	"fmt"
	"testDingding/internal/notification"
)

func main() {
	// 构造商品数据
	products := []notification.Product{
		{
			Name:     "product1",
			Link:     "https://www.ralphlauren.com/boys-clothing-sweaters/cable-knit-cotton-sweater/578341.html?dwvar578341_colorname=Oasis%20Yellow",
			ImageURL: "https://dtcralphlauren.scene7.com/is/image/PoloGSI/s7-AI322702674073_lifestyle?$plpDeskRF$",
			Price:    91.00,
		},
		{
			Name:     "product2",
			Link:     "https://example.com/macbookpro",
			ImageURL: "https://example.com/images/macbookpro.jpg",
			Price:    99.00,
		},
	}

	// 格式化商品数据为 Markdown 格式
	markdownMessage := notification.FormatProductsToMarkdown(products)

	// 实例化钉钉机器人通知器（请替换为你的 Webhook 地址）
	dingTalkNotifier := &notification.DingTalkNotifier{
		WebhookURL: "https://oapi.dingtalk.com/robot/send?access_token=??????",
	}

	// 调用一次发送通知
	if err := dingTalkNotifier.SendNotification(markdownMessage); err != nil {
		fmt.Println("发送通知失败：", err)
	} else {
		fmt.Println("通知发送成功！")
	}
}
