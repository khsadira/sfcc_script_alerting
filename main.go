package main

import (
	"log"
	"os"
)

func main() {
	hosts := getHosts()
	token, _ := getToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	for _, host := range hosts {
		sites, _ := getSitesSfcc(host, token)
		for _, site := range sites {
			dataBfr := getLocalInfo("./bfr.json") //UNITEST
			println(dataBfr, site)
			/*		data := getSfccInfo(host, site, token) //UNITEST
					ret, report := compareInfo(dataBfr, data) //UNITEST
					if ret {
						alerting(report) //UNITEST
					}*/
		} //GLOBAL TEST
	} //GLOBAL TEST
}

func getHosts() []string {
	args := os.Args[1:]
	if len(args) == 0 {
		return []string{"https://store-dev.ubi.com"}
		//LATER
		//return []string{"https://store.ubi.com", "https://staging-ubisoftstore-ubisoft.demandware.net"} LATER LATER LATER
	}
	for _, arg := range args {
		if arg == "" {
			return []string{"https://store-dev.ubi.com"}
			//LATER
			//return []string{"https://store.ubi.com", "https://staging-ubisoftstore-ubisoft.demandware.net"} LATER LATER LATER
		}
	}
	return args
}

func getLocalInfo(filepath string) string {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	return "salut"
}
