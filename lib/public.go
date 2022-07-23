/***********************************************************************************
 *                         This file is part of showip
 *                    https://github.com/PucklaMotzer09/showip
 ***********************************************************************************
 * Copyright (c) 2022 PucklaMotzer09
 *
 * This software is provided 'as-is', without any express or implied warranty.
 * In no event will the authors be held liable for any damages arising from the
 * use of this software.
 *
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute it
 * freely, subject to the following restrictions:
 *
 * 1. The origin of this software must not be misrepresented; you must not claim
 * that you wrote the original software. If you use this software in a product,
 * an acknowledgment in the product documentation would be appreciated but is
 * not required.
 *
 * 2. Altered source versions must be plainly marked as such, and must not be
 * misrepresented as being the original software.
 *
 * 3. This notice may not be removed or altered from any source distribution.
 ************************************************************************************/

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

func GetPublicIPAddress() (net.IP, error) {
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
