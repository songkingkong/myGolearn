package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Config struct {
	URL string `json:"url"`
}

func getUrl(path string) *Config {
	content, err := ioutil.ReadFile(path)
	ErrPrint(err)
	config := Config{}
	err = json.Unmarshal(content, &config)
	ErrPrint(err)
	return &config
}
func ErrPrint(err error) {
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
func httpRequest(url string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Cookie", "_AJSESSIONID=cb5b76a9782f10ae222a73147ea67331; username=songjingang; _gitlab_session=8f8eb29a39b38fbe1d344602a440765e; experimentation_subject_id=eyJfcmFpbHMiOnsibWVzc2FnZSI6IkltWXlOVGhoT0dNNUxURTVOV0V0TkdZeU1pMWhZakl4TFRBM1lUWXlNbVJqTUdVeU5pST0iLCJleHAiOm51bGwsInB1ciI6ImNvb2tpZS5leHBlcmltZW50YXRpb25fc3ViamVjdF9pZCJ9fQ; known_sign_in=Mzk3ODI3N0RiUnhNQm5KRXkvdTIvWmZSUyszUU5Ea29vcVJ6clJjeGM5d1gvQ2p3L2d4eTZqb0Z6RXFDNlRmNytHZXBwcm9NbWVFcmJUWFA3eDBkdXZuVVo0Uk9QaVhnTENrOUZ3eklDODlHdlpLb3YvTjRjZEs4M0hFR0hmK3otLTRSSGlNM1RFTHJWbU1WRFdiTXF1SHc9PQ")
	ErrPrint(err)
	reqs, err := client.Do(request)
	ErrPrint(err)
	defer reqs.Body.Close()
	bytes, err := ioutil.ReadAll(reqs.Body)
	return bytes
}
func main() {
	file := "modellearn/config/http_req.conf"
	a := getUrl(file)
	fmt.Println(a.URL)
	b := httpRequest(a.URL)
	fmt.Printf(string(b))

}
