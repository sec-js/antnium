package server

import (
	"testing"

	"github.com/dobin/antnium/pkg/client"
)

func TestConnectorHttp(t *testing.T) {
	port := "55044"
	computerId := "computerid-23"

	// Server in background, checking via client
	s := NewServer("127.0.0.1:" + port)

	s.Campaign.ClientUseWebsocket = true

	// Make a example packet the client should receive
	packetA := makeSimpleTestPacket(computerId, "001")
	s.Middleware.FrontendAddNewPacket(packetA)
	// make server go
	go s.Serve()

	// make client
	client := client.NewClient()
	client.Campaign.ServerUrl = "http://127.0.0.1:" + port
	client.Campaign.ClientUseWebsocket = true
	client.Config.ComputerId = computerId
	client.Start()

	// expect packet to be received upon connection (its already added)
	packetB := <-client.UpstreamManager.Channel
	if packetB.PacketId != "001" || packetB.ComputerId != computerId {
		t.Error("Err")
	}

	// Add a test packet via Frontend REST
	packetC := makeSimpleTestPacket(computerId, "002")
	s.Middleware.FrontendAddNewPacket(packetC)

	// Expect it
	packetD := <-client.UpstreamManager.Channel
	if packetD.PacketId != "002" || packetD.ComputerId != computerId {
		t.Error("Err")
	}

}
