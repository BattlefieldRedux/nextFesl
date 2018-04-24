package acct

import (
	"github.com/Synaxis/nextFesl/network"
	"github.com/Synaxis/nextFesl/network/codec"
	"github.com/Synaxis/nextFesl/storage/database"
)

const (
	acctGetTelemetryToken = "GetTelemetryToken"
	acctNuGetAccount      = "NuGetAccount"
	acctNuGetPersonas     = "NuGetPersonas"
	acctNuLogin           = "NuLogin"
	acctNuLoginPersona    = "NuLoginPersona"
	acctNuLookupUserInfo  = "NuLookupUserInfo"
)

const (
	clientTypeServer = "server"
)

// Account probably stands for "Account"
type Account struct {
	DB database.Adapter
}

func (acct *Account) answer(client *network.Client, pnum uint32, payload interface{}) {
	client.WriteEncode(&codec.Answer{
		Type:         codec.FeslAccount,
		PacketNumber: pnum,
		Payload:      payload,
	})
}
