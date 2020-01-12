package sms

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func InfoBip(sender string, number string, message string) {

	const APIURL = "https://vngmm.api.infobip.com/sms/2/text/single"

	usernamePassword := os.Getenv("INFO_BIP")
	encodedUsernamePassword := base64.StdEncoding.EncodeToString([]byte(usernamePassword))

	jsonData := map[string]string{
		"from": sender,
		"to":   NigeriaNumber(number),
		"text": message,
	}
	jsonStr, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("POST", APIURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Basic " + encodedUsernamePassword)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
