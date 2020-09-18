package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
	dht "github.com/libp2p/go-libp2p-kad-dht"

	mrand "math/rand"

	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	ctx := context.Background()

	// libp2p.New constructs a new libp2p Host.
	// Other options can be added here.
	sourceMultiAddr, _ := ma.NewMultiaddr("/ip4/0.0.0.0/tcp/3000")

	r := mrand.New(mrand.NewSource(int64(10)))

	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)

	if err != nil {
		panic(err)
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrs(sourceMultiAddr),
		//libp2p.NATPortMap(),
		//libp2p.DefaultSecurity,
		//libp2p.EnableRelay(circuit.OptActive, circuit.OptHop, circuit.OptDiscovery),
		libp2p.Identity(prvKey),
	}

	host, err := libp2p.New(ctx, opts...)
	if err != nil {
		panic(err)
	}

	fmt.Println("This node: ", host.ID().Pretty(), " ", host.Addrs())

	_, err = dht.New(ctx, host)
	if err != nil {
		panic(err)
	}

	select {}
}
