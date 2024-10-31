package net

import "testing"

func TestGetOutboundIP(t *testing.T) {
	ip, err := GetOutboundIP()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip)
}

func TestGetLocalIP(t *testing.T) {
	ip, err := GetLocalIP()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip)
}
