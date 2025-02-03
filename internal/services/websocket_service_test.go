package services

import (
	"testing"
)

func TestNewWebsocketService(t *testing.T) {
	server := NewWebsocketService()

	if server == nil {
		t.Errorf("NewWebsocketService() = nil; want WebsocketService")
	}

}
