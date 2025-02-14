package main

import (
	"fmt"
	"net"

	"github.com/algo7/wgcfg/internal/config"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main() {

	privateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PublicKey()

	config := config.Config{
		Interface: config.Interface{
			Address: net.IPNet{
				IP:   net.ParseIP("12.12.12.12"),
				Mask: net.CIDRMask(24, 32),
			},
			PrivateKey: privateKey,
			DNS: []net.IP{
				net.ParseIP("12.12.123.12"),
			},
			MTU: 1420,
		},
		Peers: []config.Peer{
			{
				Endpoint: &net.UDPAddr{
					IP:   net.ParseIP("12.12.12.12"),
					Port: 51820,
				},
				AllowedIPs: []net.IPNet{
					{
						IP:   net.ParseIP("12.12.32.12"),
						Mask: net.CIDRMask(24, 32),
					},
				},
				PublicKey: publicKey,
			},
		},
	}

	cfg, err := config.Generate()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)

}
