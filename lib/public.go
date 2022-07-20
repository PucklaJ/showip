package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type IPAPIResponse struct {
	Query  string
	Status string
}

func GetPublicIPv4Address() (net.IP, error) {
	resp, err := http.Get("http://ip-api.com/json/?fields=query,status")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var query IPAPIResponse
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&query)
	if err != nil {
		return nil, err
	}

	if query.Status != "success" {
		return nil, errors.New("Failed to request IP Address")
	}

	ip := net.ParseIP(query.Query)
	if ip == nil {
		return nil, fmt.Errorf("Failed to parse IP: \"%s\"", query.Query)
	}

	return ip, nil
}
