// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
)

// Proxy Contains information about a proxy server
type Proxy struct {
	tdCommon
	ID           int32     `json:"id"`             // Unique identifier of the proxy
	Server       string    `json:"server"`         // Proxy server IP address
	Port         int32     `json:"port"`           // Proxy server port
	LastUsedDate int32     `json:"last_used_date"` // Point in time (Unix timestamp) when the proxy was last used; 0 if never
	IsEnabled    bool      `json:"is_enabled"`     // True, if the proxy is enabled now
	Type         ProxyType `json:"type"`           // Type of the proxy
}

// MessageType return the string telegram-type of Proxy
func (proxy *Proxy) MessageType() string {
	return "proxy"
}

// NewProxy creates a new Proxy
//
// @param iD Unique identifier of the proxy
// @param server Proxy server IP address
// @param port Proxy server port
// @param lastUsedDate Point in time (Unix timestamp) when the proxy was last used; 0 if never
// @param isEnabled True, if the proxy is enabled now
// @param typeParam Type of the proxy
func NewProxy(iD int32, server string, port int32, lastUsedDate int32, isEnabled bool, typeParam ProxyType) *Proxy {
	proxyTemp := Proxy{
		tdCommon:     tdCommon{Type: "proxy"},
		ID:           iD,
		Server:       server,
		Port:         port,
		LastUsedDate: lastUsedDate,
		IsEnabled:    isEnabled,
		Type:         typeParam,
	}

	return &proxyTemp
}

// UnmarshalJSON unmarshal to json
func (proxy *Proxy) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID           int32  `json:"id"`             // Unique identifier of the proxy
		Server       string `json:"server"`         // Proxy server IP address
		Port         int32  `json:"port"`           // Proxy server port
		LastUsedDate int32  `json:"last_used_date"` // Point in time (Unix timestamp) when the proxy was last used; 0 if never
		IsEnabled    bool   `json:"is_enabled"`     // True, if the proxy is enabled now

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	proxy.tdCommon = tempObj.tdCommon
	proxy.ID = tempObj.ID
	proxy.Server = tempObj.Server
	proxy.Port = tempObj.Port
	proxy.LastUsedDate = tempObj.LastUsedDate
	proxy.IsEnabled = tempObj.IsEnabled

	fieldType, _ := unmarshalProxyType(objMap["type"])
	proxy.Type = fieldType

	return nil
}