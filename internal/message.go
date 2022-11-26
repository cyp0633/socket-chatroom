package internal

type Message struct {
	Type    string `json:"type"`    // from 代表自己发出，to 代表自己收到
	Content string `json:"content"` // 消息内容
	Name    string `json:"name"`    // 另一端的 IP 地址，不管发送还是接收
}

var Messages = []Message{}

var FakeMessages = []Message{
	{
		Type:    "from",
		Content: "Hello",
		Name:    "192.168.1.132",
	},
	{
		Type:    "to",
		Content: "Hi",
		Name:    "",
	},
	{
		Type:    "from",
		Content: "How are you?",
		Name:    "192.168.1.132",
	},
}
