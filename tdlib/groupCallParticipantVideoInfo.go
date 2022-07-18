// AUTOGENERATED - DO NOT EDIT

package tdlib

// GroupCallParticipantVideoInfo Contains information about a group call participant's video channel
type GroupCallParticipantVideoInfo struct {
	tdCommon
	SourceGroups []GroupCallVideoSourceGroup `json:"source_groups"` // List of synchronization source groups of the video
	EndpointID   string                      `json:"endpoint_id"`   // Video channel endpoint identifier
	IsPaused     bool                        `json:"is_paused"`     // True if the video is paused. This flag needs to be ignored, if new video frames are received
}

// MessageType return the string telegram-type of GroupCallParticipantVideoInfo
func (groupCallParticipantVideoInfo *GroupCallParticipantVideoInfo) MessageType() string {
	return "groupCallParticipantVideoInfo"
}

// NewGroupCallParticipantVideoInfo creates a new GroupCallParticipantVideoInfo
//
// @param sourceGroups List of synchronization source groups of the video
// @param endpointID Video channel endpoint identifier
// @param isPaused True if the video is paused. This flag needs to be ignored, if new video frames are received
func NewGroupCallParticipantVideoInfo(sourceGroups []GroupCallVideoSourceGroup, endpointID string, isPaused bool) *GroupCallParticipantVideoInfo {
	groupCallParticipantVideoInfoTemp := GroupCallParticipantVideoInfo{
		tdCommon:     tdCommon{Type: "groupCallParticipantVideoInfo"},
		SourceGroups: sourceGroups,
		EndpointID:   endpointID,
		IsPaused:     isPaused,
	}

	return &groupCallParticipantVideoInfoTemp
}