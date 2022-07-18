// AUTOGENERATED - DO NOT EDIT

package tdlib

// DatabaseStatistics Contains database statistics
type DatabaseStatistics struct {
	tdCommon
	Statistics string `json:"statistics"` // Database statistics in an unspecified human-readable format
}

// MessageType return the string telegram-type of DatabaseStatistics
func (databaseStatistics *DatabaseStatistics) MessageType() string {
	return "databaseStatistics"
}

// NewDatabaseStatistics creates a new DatabaseStatistics
//
// @param statistics Database statistics in an unspecified human-readable format
func NewDatabaseStatistics(statistics string) *DatabaseStatistics {
	databaseStatisticsTemp := DatabaseStatistics{
		tdCommon:   tdCommon{Type: "databaseStatistics"},
		Statistics: statistics,
	}

	return &databaseStatisticsTemp
}
