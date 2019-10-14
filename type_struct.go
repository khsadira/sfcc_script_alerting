package main

type Sites struct {
	Data []struct {
		Type string `json:"_type"`
		ID   string `json:"id"`
	} `json:"data"`
	Total int `json:"total"`
	Count int `json:"count"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}



type Data struct {
	Hits []struct {
		Data struct {
			LastModified string `json:"last_modified"`
		} `json:"data"`
	} `json:"hits"`
}