package main

import (
	"testing"
)

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

func TestCompareInfo(t *testing.T) {
	compare := []byte(`{"hits":[{"data":{"last_modified":"2019-10-11T13:24:41.345Z"}}]}`)
	compare2 := []byte(`{"hits":[{"data":{"last_modified":"2019-10-11T13:24:41.34512Z"}}]}`)
	compare3 := []byte(`{"hits":[{"type": "salut","data":{"last_modified":"2019-10-11T13:24:41.34512Z"}}]}`)

	_, retT1 := compareInfo(compare, compare)
	_, retT2 := compareInfo(compare2, compare2)
	_, retT3 := compareInfo(compare2, compare3)
	_, retF1 := compareInfo(compare, compare2)
	_, retF2 := compareInfo(compare, nil)
	if retT1 == false {
		t.Errorf("getSfccInfo not working well, retT1")
	}
	if retF1 == true {
		t.Errorf("getSfccInfo not working well, retF1")
	}
	if retF2 == true {
		t.Errorf("getSfccInfo not working well, retF2")
	}
	if retT2 == false {
		t.Errorf("getSfccInfo not working well, retT2")
	}
	if retT3 == false {
		t.Errorf("getSfccInfo not working well, retT3")
	}
}