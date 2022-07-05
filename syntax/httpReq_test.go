package syntax

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
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

func ErrPrint(err error) {
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
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
func httpRequest(config *Config) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", config.URL, nil)
	request.Header.Set("Cookie", config.Cookie)
	ErrPrint(err)
	reqs, err := client.Do(request)
	ErrPrint(err)
	defer reqs.Body.Close()
	bytes, err := ioutil.ReadAll(reqs.Body)
	if err != nil {
		log.Panic(err)
	}
	return bytes
}
func TestHttpReq(t *testing.T) {
	file := "modellearn/config/http_req.conf"
	c := formatJson(httpRequest(getUrl(file)))
	fmt.Printf("%v", c.Data.Instances[0].HOSTNAME)
}
