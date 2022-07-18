// AUTOGENERATED - DO NOT EDIT

package tdlib

// NetworkStatistics A full list of available network statistic entries
type NetworkStatistics struct {
	tdCommon
	SinceDate int32                    `json:"since_date"` // Point in time (Unix timestamp) from which the statistics are collected
	Entries   []NetworkStatisticsEntry `json:"entries"`    // Network statistics entries
}

// MessageType return the string telegram-type of NetworkStatistics
func (networkStatistics *NetworkStatistics) MessageType() string {
	return "networkStatistics"
}

// NewNetworkStatistics creates a new NetworkStatistics
//
// @param sinceDate Point in time (Unix timestamp) from which the statistics are collected
// @param entries Network statistics entries
func NewNetworkStatistics(sinceDate int32, entries []NetworkStatisticsEntry) *NetworkStatistics {
	networkStatisticsTemp := NetworkStatistics{
		tdCommon:  tdCommon{Type: "networkStatistics"},
		SinceDate: sinceDate,
		Entries:   entries,
	}

	return &networkStatisticsTemp
}