package matchmaking

import (
	"github.com/Synaxis/nextFesl/network"
)

type Game struct {
	ID         int
	LobbyID    int
	GameServer *network.Client

	PlayersJoining int
	PlayersPlaying int
	PlayerSlots    int
	// VipSlots       int
}
