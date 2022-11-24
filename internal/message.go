package internal

type Message struct {
	Type    string // from/to 从此发出/接收
	Content string // 消息内容
	Name    string // 另一端的 IP 地址，不管发送还是接收
}

var Messages = []Message{}
