package config

import (
	"net"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// Interface represents a WireGuard interface section configuration.
type Interface struct {
	// Address is the IP address and subnet of the WireGuard interface.
	Address net.IPNet `ini:"Address"`
	// PrivateKey is the private key of the WireGuard interface.
	PrivateKey wgtypes.Key `ini:"PrivateKey"`
	// DNS is a list of DNS servers to use for the WireGuard interface.
	DNS []net.IP `ini:"DNS"`
	// MTU is the MTU of the WireGuard interface.
	MTU int `ini:"MTU"`
}

// Peer represents a WireGuard peer section configuration.
type Peer struct {
	// Endpoint is the IP address and port of the peer.
	Endpoint *net.UDPAddr `ini:"Endpoint"`
	// AllowedIPs is a list of IP subnets that are allowed to be routed.
	AllowedIPs []net.IPNet `ini:"AllowedIPs"`
	// PublicKey is the public key of the peer.
	PublicKey wgtypes.Key `ini:"PublicKey"`
}

// Config represents a WireGuard configuration.
type Config struct {
	Interface Interface `ini:"Interface"`
	Peers     []Peer    `ini:"Peer"`
}
