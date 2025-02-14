// Package config provides a WireGuard configuration parser and generator.
package config

import (
	"bytes"
	"net"
	"strconv"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"gopkg.in/ini.v1"
)

const (
	DefaultMTU string = "1420"
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

func (c *Config) Generate() (*ini.File, error) {

	wgcfg := ini.Empty()

	// Add the interface section.
	ifaceSec, err := wgcfg.NewSection("Interface")
	if err != nil {
		return nil, err
	}

	// Add the key-value pairs to the interface section.
	ifaceSec.NewKey("Address", c.Interface.Address.String())
	ifaceSec.NewKey("PrivateKey", c.Interface.PrivateKey.String())
	if len(c.Interface.DNS) > 0 {
		ifaceSec.NewKey("DNS", c.Interface.DNS[0].String())
		for _, dns := range c.Interface.DNS[1:] {
			ifaceSec.NewKey("", dns.String())
		}
	}

	if c.Interface.MTU > 0 {
		ifaceSec.NewKey("MTU", strconv.Itoa(c.Interface.MTU))
	} else {
		ifaceSec.NewKey("MTU", DefaultMTU)
	}

	// Add the peer sections.
	for _, peer := range c.Peers {
		peerSec, err := wgcfg.NewSection("Peer")
		if err != nil {
			return nil, err
		}

		// Add the key-value pairs to the peer section.
		peerSec.NewKey("PublicKey", peer.PublicKey.String())
		peerSec.NewKey("AllowedIPs", peer.AllowedIPs[0].String())
		for _, allowedIP := range peer.AllowedIPs[1:] {
			peerSec.NewKey("", allowedIP.String())
		}

		if peer.Endpoint != nil {
			peerSec.NewKey("Endpoint", peer.Endpoint.String())
		}

	}

	return wgcfg, nil
}

func (c *Config) String() (string, error) {
	cfg, err := c.Generate()
	if err != nil {
		return "", err
	}

	// Convert ini file to string
	var outputBuffer bytes.Buffer
	_, err = cfg.WriteTo(&outputBuffer)
	if err != nil {
		return "", err
	}

	return outputBuffer.String(), nil
}

func (c *Config) Save(fileName string) error {
	cfg, err := c.Generate()
	if err != nil {
		return err
	}

	return cfg.SaveTo(fileName)

}
