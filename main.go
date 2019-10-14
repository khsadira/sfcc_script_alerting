package main

import (
	"log"
	"os"
)

func main() {
	path := "data_x.json"
	hosts := getHosts()
	token, _ := getToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	for _, host := range hosts {
		lst, id := getNoneProds(host)

		/*		sites, _ := getSitesSfcc(host, token)
				for _, site := range sites {
					/*if site == "fr_ubisoft" {
						log.Println("Get Datas...")
						data := getSfccInfo(host, site, token)
						dataBfr := getFileInfo(path)

						log.Println("Comparing Datas...")
						report, ret := compareInfo(dataBfr, data)

						if !ret {
							log.Println("Alerting...")
							alerting(report)
							log.Println("Replacing...")
							replaceFileInfo(path, data)
						} else {
							log.Println("No change detected")
						}
					}
				}*/
	}
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
