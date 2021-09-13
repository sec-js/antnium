package client

import (
	"bytes"
	"net/http"
)

func (d UpstreamRest) PacketGetUrl() string {
	return d.campaign.ServerUrl + d.campaign.PacketGetPath + d.config.ComputerId
}

func (d UpstreamRest) PacketSendUrl() string {
	return d.campaign.ServerUrl + d.campaign.PacketSendPath
}

func (d UpstreamRest) HttpGet(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Session-Token", d.campaign.ApiKey)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d UpstreamRest) HttpPost(url string, data *bytes.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Session-Token", d.campaign.ApiKey)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}