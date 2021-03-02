/*-
 * Copyright (c) 2021 F5 Networks, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package manager

import (
	"fmt"
	log "github.com/f5devcentral/f5-ipam-controller/pkg/vlogger"
)

// Manager defines the interface that the IPAM system should implement
type Manager interface {
	// Creates an A record
	CreateARecord(hostname, ipAddr string) bool
	// Deletes an A record and releases the IP address
	DeleteARecord(hostname, ipAddr string)
	// Gets IP Address associated with hostname
	GetIPAddress(hostname string) string
	// Gets and reserves the next available IP address
	GetNextIPAddress(cidr string) string
	// Allocates this particular ip from the CIDR
	AllocateIPAddress(cidr, ipAddr string) bool
	// Releases an IP address
	ReleaseIPAddress(ipAddr string)
}

const F5IPAMProvider = "f5-ip-provider"

type Params struct {
	Provider string
	IPAMManagerParams
}

func NewManager(params Params) (Manager, error) {
	switch params.Provider {
	case F5IPAMProvider:
		log.Debugf("[MGR] Creating Manager with Provider: %v", F5IPAMProvider)
		f5IPAMParams := IPAMManagerParams{Range: params.Range}
		return NewIPAMManager(f5IPAMParams)
	default:
		log.Errorf("[MGR] Unknown Provider: %v", params.Provider)
	}
	return nil, fmt.Errorf("manager cannot be initialized")
}
