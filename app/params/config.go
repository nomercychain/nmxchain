package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// Bech32MainPrefix defines the main network Bech32 prefix for accounts
	Bech32MainPrefix = "nmx"

	// CoinType is the SLIP-44 coin type for NoMercy Chain
	CoinType = 118 // Same as Cosmos Hub

	// Bech32PrefixAccAddr defines the Bech32 prefix for accounts
	Bech32PrefixAccAddr = Bech32MainPrefix
	// Bech32PrefixAccPub defines the Bech32 prefix for account public keys
	Bech32PrefixAccPub = Bech32MainPrefix + "pub"
	// Bech32PrefixValAddr defines the Bech32 prefix for validator addresses
	Bech32PrefixValAddr = Bech32MainPrefix + "valoper"
	// Bech32PrefixValPub defines the Bech32 prefix for validator public keys
	Bech32PrefixValPub = Bech32MainPrefix + "valoperpub"
	// Bech32PrefixConsAddr defines the Bech32 prefix for consensus node addresses
	Bech32PrefixConsAddr = Bech32MainPrefix + "valcons"
	// Bech32PrefixConsPub defines the Bech32 prefix for consensus node public keys
	Bech32PrefixConsPub = Bech32MainPrefix + "valconspub"
)

// SetupSDKConfig sets up the SDK configuration
func SetupSDKConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	config.SetAddressVerifier(address.Verifier)
	config.SetCoinType(CoinType)
	config.Seal()
}
