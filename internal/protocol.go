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
}
