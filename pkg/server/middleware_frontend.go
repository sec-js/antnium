package server

import (
	"github.com/dobin/antnium/pkg/model"
)

func (s *Middleware) FrontendAddNewPacket(packet *model.Packet) error {
	// Add to packet DB and get packetInfo
	packetInfo, err := s.packetDb.addFromFrontend(packet)
	if err != nil {
		return err
	}

	// Notify UI immediately (for initial STATE_RECORDED)
	s.frontendManager.FrontendWs.broadcastPacket(*packetInfo)

	// Send to client, if they are connected via Websocket
	ok := s.connectorManager.ConnectorWs.TryViaWebsocket(&packetInfo.Packet)
	if ok {
		s.packetDb.sentToClient(packet.PacketId, "")

		// only notify UI if we really sent a packet
		s.frontendManager.FrontendWs.broadcastPacket(*packetInfo)
	}

	return nil
}

func (s *Middleware) FrontendGetAllPacket() []*PacketInfo {
	return s.packetDb.getAll()
}

func (s *Middleware) FrontendGetPacketById(computerId string) []PacketInfo {
	var filteredPackets []PacketInfo = make([]PacketInfo, 0)
	packetInfos := s.packetDb.getAll()
	for _, packetInfo := range packetInfos {
		if packetInfo.Packet.ComputerId == computerId {
			filteredPackets = append(filteredPackets, *packetInfo)
		}
	}
	return filteredPackets
}

func (s *Middleware) FrontendGetAllClients() []ClientInfo {
	return s.clientInfoDb.getAsList()
}

/*
func (s *Middleware) FrontendGetCampaign() campaign.Campaign {
	return *s.Campaign
}

func (s *Middleware) FrontendDirUpload() []model.DirEntry {
}

func (s *Middleware) FrontendDirStatic() {
}
*/