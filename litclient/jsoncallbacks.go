package litclient

import (
	"context"

	"github.com/lightninglabs/faraday/frdrpc"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"github.com/lightninglabs/loop/looprpc"
	"github.com/lightninglabs/pool/poolrpc"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/assetwalletrpc"
	"github.com/lightninglabs/taproot-assets/taprpc/mintrpc"
	"github.com/lightninglabs/taproot-assets/taprpc/priceoraclerpc"
	"github.com/lightninglabs/taproot-assets/taprpc/rfqrpc"
	"github.com/lightninglabs/taproot-assets/taprpc/tapchannelrpc"
	"github.com/lightninglabs/taproot-assets/taprpc/universerpc"
	"github.com/vanditshah99/lnd/lnrpc"
	"github.com/vanditshah99/lnd/lnrpc/autopilotrpc"
	"github.com/vanditshah99/lnd/lnrpc/chainrpc"
	"github.com/vanditshah99/lnd/lnrpc/invoicesrpc"
	"github.com/vanditshah99/lnd/lnrpc/routerrpc"
	"github.com/vanditshah99/lnd/lnrpc/signrpc"
	"github.com/vanditshah99/lnd/lnrpc/verrpc"
	"github.com/vanditshah99/lnd/lnrpc/walletrpc"
	"github.com/vanditshah99/lnd/lnrpc/watchtowerrpc"
	"github.com/vanditshah99/lnd/lnrpc/wtclientrpc"
	"google.golang.org/grpc"
)

// StubPackageRegistration defines the signature of a function that maps JSON
// method URIs to the function that should be called to handle the method.
type StubPackageRegistration func(map[string]func(context.Context,
	*grpc.ClientConn, string, func(string, error)))

// Registrations defines a list of StubPackageRegistrations that a lit client
// will have access to when using Lit.
var Registrations = []StubPackageRegistration{
	lnrpc.RegisterLightningJSONCallbacks,
	lnrpc.RegisterStateJSONCallbacks,
	autopilotrpc.RegisterAutopilotJSONCallbacks,
	chainrpc.RegisterChainNotifierJSONCallbacks,
	invoicesrpc.RegisterInvoicesJSONCallbacks,
	routerrpc.RegisterRouterJSONCallbacks,
	signrpc.RegisterSignerJSONCallbacks,
	verrpc.RegisterVersionerJSONCallbacks,
	walletrpc.RegisterWalletKitJSONCallbacks,
	watchtowerrpc.RegisterWatchtowerJSONCallbacks,
	wtclientrpc.RegisterWatchtowerClientJSONCallbacks,
	looprpc.RegisterSwapClientJSONCallbacks,
	poolrpc.RegisterTraderJSONCallbacks,
	frdrpc.RegisterFaradayServerJSONCallbacks,
	litrpc.RegisterSessionsJSONCallbacks,
	litrpc.RegisterAccountsJSONCallbacks,
	litrpc.RegisterAutopilotJSONCallbacks,
	litrpc.RegisterFirewallJSONCallbacks,
	litrpc.RegisterStatusJSONCallbacks,
	taprpc.RegisterTaprootAssetsJSONCallbacks,
	assetwalletrpc.RegisterAssetWalletJSONCallbacks,
	universerpc.RegisterUniverseJSONCallbacks,
	mintrpc.RegisterMintJSONCallbacks,
	priceoraclerpc.RegisterPriceOracleJSONCallbacks,
	rfqrpc.RegisterRfqJSONCallbacks,
	tapchannelrpc.RegisterTaprootAssetChannelsJSONCallbacks,
}
