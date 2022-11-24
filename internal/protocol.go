package internal

import (
	"log"
	"net"
	"regexp"
	"strings"
	"time"
)

var conn *net.Conn

// 尝试连接服务器
func Connect() {
	c, err := net.DialTimeout("tcp", "localhost:65432", 2*time.Second)
	if err != nil {
		log.Panicln(err)
	}
	conn = &c
}

var heloReplyRegex = regexp.MustCompile(`^CLIENTS .+`)

// 发送 HELO 命令
func DoHelo() {
	c := *conn
	_, err := c.Write([]byte("HELO"))
	if err != nil {
		log.Println(err)
	}

	buf := make([]byte, 1024)
	n, err := c.Read(buf) // 返回其他客户端列表
	str := string(buf[:n])
	if err != nil || !heloReplyRegex.MatchString(str) {
		log.Println(err)
	}

	clients := strings.Split(string(buf[8:n]), " ")
	for _, client := range clients {
		AddClient(client)
	}
}

var sendReplyRegex = regexp.MustCompile(`^OK.+`)

func DoSend(msg string, to string) {
	c := *conn
	str := "SEND " + to + " MSG " + msg
	_, err := c.Write([]byte(str))
	if err != nil {
		log.Println(err)
	}

	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	str = string(buf[:n])
	if err != nil || !sendReplyRegex.MatchString(str) {
		log.Println(err)
	}

	Messages = append(Messages, Message{
		Type:    "from",
		Content: msg,
		Name:    to,
	})
}

var pullReplyRegex = regexp.MustCompile(`^(FROM .+ CONTENT .+\n)+END`)

var fromRegex = regexp.MustCompile(`^FROM .+ CONTENT`)

var contentRegex = regexp.MustCompile(`CONTENT .+\n`)

func DoPull() {
	c := *conn
	str := "PULL"
	_, err := c.Write([]byte(str))
	if err != nil {
		log.Println(err)
	}

	buf := make([]byte, 4096) // 为消息留大点 buffer
	n, err := c.Read(buf)
	str = string(buf[:n])
	if err != nil || !pullReplyRegex.MatchString(str) {
		log.Println(err)
	}

	// 解析消息
	msg := strings.Split(str, "\n")
	for _, m := range msg {
		if m == "END" {
			return
		}
		from := fromRegex.FindString(m)[5 : len(m)-8]
		content := contentRegex.FindString(m)[8:]
		Messages = append(Messages, Message{
			Type:    "to",
			Content: content,
			Name:    from,
		})
	}
}
