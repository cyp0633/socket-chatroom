package internal

// 其他客户端的数组
var Clients = []string{}

func AddClient(client string) {
	Clients = append(Clients, client)
}

func RemoveClient(client string) {
	for i, c := range Clients {
		if c == client {
			Clients = append(Clients[:i], Clients[i+1:]...)
		}
	}
}