package internal

import (
	"net"
	"regexp"
	"strings"
	"time"

	"go.uber.org/zap"
)

var conn *net.Conn

// 尝试连接服务器
func Connect(ip string) {
	if ip == "" {
		ip = "127.0.0.1"
	}
	c, err := net.DialTimeout("tcp", ip+":65432", 2*time.Second)
	if err != nil {
		Logger.Panic("Failed to connect to server", zap.Error(err))
	}
	conn = &c
}

var heloReplyRegex = regexp.MustCompile(`^CLIENTS .+`)

// 发送 HELO 命令
func DoHelo() {
	c := *conn
	_, err := c.Write([]byte("HELO"))
	if err != nil {
		Logger.Error("Failed to send HELO", zap.Error(err))
	}

	buf := make([]byte, 1024)
	n, err := c.Read(buf) // 返回其他客户端列表

	if err != nil || !heloReplyRegex.Match(buf) {
		Logger.Error("Failed to receive HELO reply", zap.Error(err))
	}

	clients := strings.Split(string(buf[8:n]), " ")
	for _, client := range clients {
		if client == "" {
			continue
		}
		AddClient(client)
	}
	Logger.Info("Received client list", zap.Strings("clients", clients))
}

var sendReplyRegex = regexp.MustCompile(`^OK`)

func DoSend(msg string, to string) {
	c := *conn
	str := "SEND " + to + " MSG " + msg
	_, err := c.Write([]byte(str))
	if err != nil {
		Logger.Error("Failed to send SEND", zap.Error(err))
	}

	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil || !sendReplyRegex.Match(buf) {
		Logger.Error("Failed to receive SEND reply", zap.String("msg", string(buf[:n])), zap.Error(err))
	}

	Messages = append(Messages, Message{
		Type:    "from",
		Content: msg,
		Name:    to,
	})
	Logger.Info("Sent message successfully", zap.String("to", to), zap.String("message", msg))
}

var pullReplyRegex = regexp.MustCompile(`^[0-9]+ MESSAGES\n(FROM .+ CONTENT .+\n)*END PULL\n`)

var pullUserRegex = regexp.MustCompile(`[0-9]+ USERS\n(.+\n)*END USER`)

var fromRegex = regexp.MustCompile(`^FROM .+ CONTENT`)

var contentRegex = regexp.MustCompile(`CONTENT .+`)

func DoPull() {
	c := *conn
	str := "PULL"
	_, err := c.Write([]byte(str))
	if err != nil {
		Logger.Error("Failed to send PULL", zap.Error(err))
	}

	buf := make([]byte, 4096) // 为消息留大点 buffer
	n, err := c.Read(buf)
	str = string(buf[:n])
	if err != nil {
		Logger.Error("Failed to receive PULL reply", zap.Error(err))
		return
	}
	if !pullReplyRegex.MatchString(str) {
		Logger.Error("Not reply format", zap.String("msg", str))
	}
	if !pullUserRegex.MatchString(str) {
		Logger.Error("Not user format", zap.String("msg", str))
	}

	// 消息部分
	replyPart := pullReplyRegex.FindString(str)
	msg := strings.Split(replyPart, "\n")
	for i, m := range msg {
		if m == "END PULL" || i == 0 || m == "" {
			continue
		}
		temp := ""
		temp = fromRegex.FindString(m)
		from := temp[5 : len(temp)-8]
		temp = contentRegex.FindString(m)
		content := temp[8:]
		Messages = append(Messages, Message{
			Type:    "to",
			Content: content,
			Name:    from,
		})
		Logger.Info("Received message", zap.String("from", from), zap.String("message", content))
	}

	// 用户部分
	userPart := pullUserRegex.FindString(str)
	users := strings.Split(userPart, "\n")
	tempClients := []string{}
	for i, user := range users {
		if user == "END USER" || user == "" || i == 0 {
			continue
		}
		tempClients = append(tempClients, user)
	}
	Clients = tempClients
}

var exitRegex = regexp.MustCompile(`^OK`)

func DoExit() {
	if conn == nil {
		Logger.Info("Trying to connect a closed connection")
		return
	}
	c := *conn
	str := "EXIT"
	_, err := c.Write([]byte(str))
	if err != nil {
		Logger.Error("Failed to send EXIT", zap.Error(err))
	}

	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	str = string(buf[:n])
	if err != nil || !exitRegex.MatchString(str) {
		Logger.Error("Failed to receive EXIT reply", zap.Error(err))
	}

	c.Close()
	Logger.Info("Closed connection")
	conn = nil
}

var userRegex = regexp.MustCompile(`[0-9]+ USERS(.|\n)*END`)

func DoUser() {
	c := *conn
	str := "USER"
	_, err := c.Write([]byte(str))
	if err != nil {
		Logger.Error("Failed to send USER", zap.Error(err))
	}

	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil || !userRegex.MatchString(string(buf[:n])) {
		Logger.Error("Failed to receive USER reply", zap.String("msg", string(buf[:n])), zap.Error(err))
	}

	users := strings.Split(string(buf[:n]), "\n")
	tempClients := []string{}
	for i, user := range users {
		if user == "END" || user == "" || i == 0 {
			continue
		}
		tempClients = append(tempClients, user)
	}
	Clients = tempClients
}
