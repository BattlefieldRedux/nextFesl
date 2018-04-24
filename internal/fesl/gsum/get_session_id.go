package gsum

import (
	"github.com/Synaxis/nextFesl/network"
	"github.com/Synaxis/nextFesl/network/codec"
)

type ansGetSessionID struct {
	Txn string `fesl:"TXN"`
	// Games  []Game  `fesl:"games"`
	// Events []Event `fesl:"events"`
}

// GetSessionID handles gsum.GetSessionID command
func (gsum *GameSummary) GetSessionID(client *network.Client, event *codec.Command) {
	gsum.answer(client, 0, ansGetSessionID{
		Txn: gsumGetSessionID,
	})
}
