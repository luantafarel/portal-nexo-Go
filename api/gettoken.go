package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func get() (string, int64) {

	url := "https://homologacaoapi.nexoweb.com/nexoapirest/token"
	method := "POST"

	payload := strings.NewReader("password=123456&username=USU1&grant_type=password&client_id=FUNDACAO")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	sec := now.Unix()
	token := fmt.Sprintf("%v", data["access_token"])
	return token, 14400 + sec
}
