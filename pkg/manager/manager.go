package manager

import log "github.com/f5devcentral/f5-ipam-controller/pkg/vlogger"

// Manager defines the interface that the IPAM system should implement
type Manager interface {
	// Creates an A record
	CreateARecord(params ...interface{}) bool
	// Deletes an A record and releases the IP address
	DeleteARecord(params ...interface{})
	// Gets IP Address associated with hostname
	GetIPAddress(hostname string) string
	// Gets and reserves the next available IP address
	GetNextIPAddress(cidr string) string
	// Allocates this particular ip from the CIDR
	AllocateIPAddress(cidr, ipAddr string) bool
	// Releases an IP address
	ReleaseIPAddress(ipAddr string)
}

// Search String for Default Static IP Approach
const F5IPAMProvider = "f5"

// search string for Infoblox Provider
const INFOBLOXProvider = "infoblox"
const INFOBLOXcmpType = "f5"

type Params struct {
	Provider string
	IPAMManagerParams
	InfobloxParams
}

func NewManager(params Params) Manager {
	switch params.Provider {
	case F5IPAMProvider:
		log.Debugf("[MGR] Creating Manager with Provider: %v", F5IPAMProvider)
		f5IPAMParams := IPAMManagerParams{Range: params.Range}
		return NewIPAMManager(f5IPAMParams)
	case INFOBLOXProvider:
		log.Debugf("[MGR] Creating Manager with Provider: %v", INFOBLOXProvider)
		ibclient, err := NewInfobloxManager(&params.InfobloxParams, INFOBLOXcmpType)
		if err == nil {
			return ibclient
		}
	default:
		log.Errorf("[MGR] Unknown Provider: %v", params.Provider)
	}
	return nil
}
