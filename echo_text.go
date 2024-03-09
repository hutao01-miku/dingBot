package main

import (
	"context"
	"dingding/config"
	"fmt"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/client"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/logger"
)

func OnChatBotMessageReceived(ctx context.Context, data *chatbot.BotCallbackDataModel) ([]byte, error) {
	// 将接收到的消息内容存储到userMessage中
	userMessage := data.Text.Content
	fmt.Println(userMessage)
	// 获取选择消息
	choiceMessage, err := GetChoiceMessage(userMessage)
	fmt.Println(choiceMessage)
	if err != nil {
		return nil, fmt.Errorf("error getting choice message: %v", err)
	}

	// 使用 ChatbotReplier 发送简单文本消息
	chatbotReplier := chatbot.NewChatbotReplier()
	if err := chatbotReplier.SimpleReplyText(ctx, data.SessionWebhook, []byte(choiceMessage)); err != nil {
		return nil, fmt.Errorf("error replying with text message: %v", err)
	}

	// 返回响应数据，通知钉钉服务器消息已经处理完毕
	return []byte(""), nil
}

func main() {
	var clientId, clientSecret string
	//flag.StringVar(&clientId, "client_id", "", "your-client-id")
	//flag.StringVar(&clientSecret, "client_secret", "", "your-client-secret")
	//flag.Parse()
	//使用： go run echo_text.go --client_id dingoovx5nyx0wcrrilz --client_secret 6Kez_CEttRULR-yyIZDbDkt_NRE-U7
	//TzPH4d5Ok4VBv0jWVR08z7w4VRoDyAU8Qw
	clientId = config.ClientID
	//clientId = viper.GetString("client_id")
	clientSecret = config.ClientSecret

	if len(clientId) == 0 || len(clientSecret) == 0 {
		panic("command line options --client_id and --client_secret required")
	}

	logger.SetLogger(logger.NewStdTestLogger())

	cli := client.NewStreamClient(client.WithAppCredential(client.NewAppCredentialConfig(clientId, clientSecret)))

	cli.RegisterChatBotCallbackRouter(OnChatBotMessageReceived)

	err := cli.Start(context.Background())
	if err != nil {
		panic(err)
	}

	select {}
}
