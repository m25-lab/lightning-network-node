package p2p

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type ipVersion uint8

const (
	ipv4 ipVersion = 0
	ipv6 ipVersion = 1
)

/*
Encode layout: |ip version|address|port|
*/
func EncodeTCPAddress(w io.Writer, addr *net.TCPAddr) error {
	var (
		ipVersion byte
		ip        []byte
	)

	if result := addr.IP.To4(); result != nil {
		ipVersion = byte(ipv4)
		ip = result
	} else {
		ipVersion = byte(ipv6)
		ip = addr.IP.To16()
	}

	if ip == nil {
		return fmt.Errorf("unable to encode IP %v", addr.IP)
	}

	if _, err := w.Write([]byte{ipVersion}); err != nil {
		return err
	}

	if _, err := w.Write(ip); err != nil {
		return err
	}

	var port [2]byte
	binary.BigEndian.PutUint16(port[:], uint16(addr.Port))
	if _, err := w.Write(port[:]); err != nil {
		return err
	}

	return nil
}

func decodePort(r io.Reader) (int, error) {
	var port [2]byte
	if _, err := r.Read(port[:]); err != nil {
		return 0, err
	}

	return int(binary.BigEndian.Uint16(port[:])), nil
}

func decodeIP(r io.Reader, ipVer ipVersion) (net.IP, error) {
	var len uint
	if ipVer == ipv4 {
		len = 4
	} else {
		len = 16
	}

	ip := make([]byte, len)
	if _, err := r.Read(ip[:]); err != nil {
		return nil, err
	}

	return net.IP(ip[:]), nil
}

func DecodeIPAddress(r io.Reader) (net.Addr, error) {
	var ipVer [1]byte
	if _, err := r.Read(ipVer[:]); err != nil {
		return nil, err
	}

	ip, err := decodeIP(r, ipVersion(ipVer[0]))
	if err != nil {
		return nil, err
	}

	port, err := decodePort(r)
	if err != nil {
		return nil, err
	}

	return &net.TCPAddr{
		IP:   ip,
		Port: port,
	}, nil
}
