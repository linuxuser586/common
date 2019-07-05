package net

import (
	"errors"
	"net"
	"testing"

	"github.com/linuxuser586/common/pkg/os"
)

const testIP = "10.0.2.2"

var emptyEvn = func(key string) string {
	return ""
}

var emptyIP = func() ([]net.Addr, error) {
	return []net.Addr{}, nil
}

var loopbackIP = func() ([]net.Addr, error) {
	ip := []byte{127, 0, 0, 1}
	a := &net.IPNet{IP: ip}
	return []net.Addr{a}, nil
}

var loopbackWithValidIP = func() ([]net.Addr, error) {
	a := &net.IPNet{IP: []byte{127, 0, 0, 1}}
	b := &net.IPNet{IP: []byte{10, 1, 1, 1}}
	return []net.Addr{a, b}, nil
}

func TestPodEnvDefined(t *testing.T) {
	os.Getenv = func(key string) string {
		return testIP
	}
	ip, err := PodIP()
	if err != nil {
		t.Fatal(err)
	}
	if ip != testIP {
		t.Errorf("got: %v, want: %v\n", ip, testIP)
	}
}

func TestAddressError(t *testing.T) {
	os.Getenv = emptyEvn
	InterfaceAddrs = func() ([]net.Addr, error) {
		return nil, errors.New("fake error")
	}
	_, err := PodIP()
	if err == nil {
		t.Error("got: nil error, want: fake error")
	}
}

func TestAddressNoIPFound(t *testing.T) {
	os.Getenv = emptyEvn
	InterfaceAddrs = emptyIP
	_, err := PodIP()
	if err == nil {
		t.Errorf("got: %v, want: nil error", errNoIP)
	}
}

func TestLoopbackIPOnly(t *testing.T) {
	os.Getenv = emptyEvn
	InterfaceAddrs = loopbackIP
	_, err := PodIP()
	if err == nil {
		t.Errorf("got: %v, want: nil error", errNoIP)
	}
}

func TestLoopbackWithValidIP(t *testing.T) {
	os.Getenv = emptyEvn
	InterfaceAddrs = loopbackWithValidIP
	ip, err := PodIP()
	exp := "10.1.1.1"
	if err != nil {
		t.Fatalf("got: %v, want: %v", err, exp)
	}
	if ip != exp {
		t.Errorf("got: %v, want: %v", ip, exp)
	}
}
