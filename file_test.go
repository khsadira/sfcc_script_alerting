package main

import "testing"

func TestGetHosts(t *testing.T) {
	hosts := getHosts()
	for _, host := range hosts {
		if host == "" {
			t.Errorf("getTarget was incorrect, got nothing in one string of the slice")
		}
	}
}

func TestGetSitesSfcc(t *testing.T) {
	var i int

	token, _ := getToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	sites, len := getSitesSfcc("https://store-dev.ubi.com", token)

	for _, site := range sites {
		if site == "" {
			t.Errorf("getTarget was incorrect, got nothing in one string of the slice")
		}
		i += 1
	}

	if i != len {
		t.Errorf("getTarget was incorrect, len was %d, suppose to be %d", len, i)
	}
}