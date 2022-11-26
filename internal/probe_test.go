package internal_test

import (
	"log"
	"socket-chatroom/internal"
	"testing"
)

func TestProbe(t *testing.T) {
	server := internal.Probe()
	log.Default().Println(server)
}
