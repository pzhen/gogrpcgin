package utils

import (
	"strings"
	"net/http"
	"io/ioutil"
)

func HttPost(url string, jsonStr string) (r []byte, err error){

	payload := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("authorization", "")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if err != nil{
		return nil ,err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil{
		return nil ,err
	}

	return body,nil
}
