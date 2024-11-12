//go:build !dev
// +build !dev

package perms

import (
	"net"

	"github.com/vanditshah99/lnd/autopilot"
	"github.com/vanditshah99/lndditshah99/chainreg"
	"github.com/vanditshah99/lndditshah99/lnrpc"
	"github.com/vanditshah99/lndditshah99/lnrpc/autopilotrpc"
	"github.com/vanditshah99/lndditshah99/lnrpc/chainrpc"
	"github.com/vanditshah99/lndditshah99/lnrpc/devrpc"
	"github.com/vanditshah99/lndditshah99/lnrpc/invoicesrpc"
	"github.com/vanditshah99/lndditshah99/lnrpc/neutrinorpc"
	"github.com/vanditshah99/lndditshah99/lnrpc/peersrpc"
	"github.com/vanditshah99/lndditshah99/lnrpc/routerrpc"
	"github.com/vanditshah99/lndditshah99/lnrpc/signrpc"
	"github.com/vanditshah99/lndditshah99/lnrpc/walletrpc"
	"github.com/vanditshah99/lndditshah99/lnrpc/watchtowerrpc"
	"github.com/vanditshah99/lndditshah99/lnrpc/wtclientrpc"
	"github.com/vanditshah99/lndditshah99/lntest/mock"
	"github.com/vanditshah99/lndditshah99/routing"
	"github.com/vanditshah99/lndditshah99/sweep"
)

// mockConfig implements lnrpc.SubServerConfigDispatcher. It provides the
// functionality required so that the lnrpc.GrpcHandler.CreateSubServer
// function can be called without panicking.
type mockConfig struct{}

var _ lnrpc.SubServerConfigDispatcher = (*mockConfig)(nil)

// FetchConfig is a mock implementation of lnrpc.SubServerConfigDispatcher. It
// is used as a parameter to lnrpc.GrpcHandler.CreateSubServer and allows the
// function to be called without panicking. This is useful because
// CreateSubServer can be used to extract the permissions required by each
// registered subserver.
//
// TODO(elle): remove this once the sub-server permission lists in LND have been
// exported
func (t *mockConfig) FetchConfig(subServerName string) (interface{}, bool) {
	switch subServerName {
	case "InvoicesRPC":
		return &invoicesrpc.Config{}, true
	case "WatchtowerClientRPC":
		return &wtclientrpc.Config{
			Resolver: func(_, _ string) (*net.TCPAddr, error) {
				return nil, nil
			},
		}, true
	case "AutopilotRPC":
		return &autopilotrpc.Config{
			Manager: &autopilot.Manager{},
		}, true
	case "ChainRPC":
		return &chainrpc.Config{
			ChainNotifier: &chainreg.NoChainBackend{},
			Chain:         &mock.ChainIO{},
		}, true
	case "DevRPC":
		return &devrpc.Config{}, true
	case "NeutrinoKitRPC":
		return &neutrinorpc.Config{}, true
	case "PeersRPC":
		return &peersrpc.Config{}, true
	case "RouterRPC":
		return &routerrpc.Config{
			Router: &routing.ChannelRouter{},
		}, true
	case "SignRPC":
		return &signrpc.Config{
			Signer: &mock.DummySigner{},
		}, true
	case "WalletKitRPC":
		return &walletrpc.Config{
			FeeEstimator: &chainreg.NoChainBackend{},
			Wallet:       &mock.WalletController{},
			KeyRing:      &mock.SecretKeyRing{},
			Sweeper:      &sweep.UtxoSweeper{},
			Chain:        &mock.ChainIO{},
		}, true
	case "WatchtowerRPC":
		return &watchtowerrpc.Config{}, true
	default:
		return nil, false
	}
}
