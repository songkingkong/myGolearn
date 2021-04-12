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
	URL    string `json:"url"`
	Cookie string `json:"cookie"`
}
type HostInfo struct {
	HOSTNAME    string `json:"hostname"`
	TYPE        string `json:"type"`
	IDC         string `json:"idc"`
	VENDOR      string `json:"vendor"`
	HostID      string `json:"host_id"`
	InternalIp  string `json:"internal_ip"`
	RackName    string `json:"rack_name"`
	AssetNumber string `json:"asset_number"`
	BelongTo    string `json:"belong_to"`
}
type Instances struct {
	Instances []*HostInfo `json:"instances"`
}
type JsonData struct {
	Data *Instances `json:"data"`
}

func getUrl(path string) *Config {
	content, err := ioutil.ReadFile(path)
	ErrPrint(err)
	config := Config{}
	err = json.Unmarshal(content, &config)
	ErrPrint(err)
	return &config
}
func formatJson(jsonapi []byte) *JsonData {
	data := JsonData{}
	err := json.Unmarshal(jsonapi, &data)
	ErrPrint(err)
	return &data

}
func ErrPrint(err error) {
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
func httpRequest(config *Config) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", config.URL, nil)
	request.Header.Set("Cookie", config.Cookie)
	ErrPrint(err)
	reqs, err := client.Do(request)
	ErrPrint(err)
	defer reqs.Body.Close()
	bytes, err := ioutil.ReadAll(reqs.Body)
	return bytes
}
func main() {
	file := "modellearn/config/http_req.conf"
	c := formatJson(httpRequest(getUrl(file)))
	fmt.Printf("%#v", c.Data.Instances[0].HOSTNAME)
}
