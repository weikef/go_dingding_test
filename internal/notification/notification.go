package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Notifier 定义通知接口
type Notifier interface {
	SendNotification(message string) error
}

// Product 定义商品数据结构
type Product struct {
	Name     string
	Link     string
	ImageURL string
	Price    float64
}

// DingTalkNotifier 钉钉机器人通知器
type DingTalkNotifier struct {
	WebhookURL string
}

// SendNotification 实现 Notifier 接口，发送 Markdown 格式的消息到钉钉
func (d *DingTalkNotifier) SendNotification(message string) error {
	payload := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": "商品更新推荐",
			"text":  message,
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON 序列化失败: %v", err)
	}

	resp, err := http.Post(d.WebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("发送消息失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("钉钉返回错误状态: %d", resp.StatusCode)
	}

	return nil
}

// FormatProductsToMarkdown 格式化商品数据为 Markdown 文本
func FormatProductsToMarkdown(products []Product) string {
	var dd strings.Builder
	dd.WriteString("### 📢 今日推荐商品\n\n")
	for _, product := range products {
		dd.WriteString(fmt.Sprintf("**[%s](%s)**  \n", product.Name, product.Link))
		dd.WriteString(fmt.Sprintf("![图片](%s)  \n", product.ImageURL))
		dd.WriteString(fmt.Sprintf("💰 价格: **￥%.2f**  \n", product.Price))
		dd.WriteString("---\n")
	}
	return dd.String()
}
