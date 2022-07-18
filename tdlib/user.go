// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
)

// User Represents a user
type User struct {
	tdCommon
	ID                int64         `json:"id"`                 // User identifier
	FirstName         string        `json:"first_name"`         // First name of the user
	LastName          string        `json:"last_name"`          // Last name of the user
	Username          string        `json:"username"`           // Username of the user
	PhoneNumber       string        `json:"phone_number"`       // Phone number of the user
	Status            UserStatus    `json:"status"`             // Current online status of the user
	ProfilePhoto      *ProfilePhoto `json:"profile_photo"`      // Profile photo of the user; may be null
	IsContact         bool          `json:"is_contact"`         // The user is a contact of the current user
	IsMutualContact   bool          `json:"is_mutual_contact"`  // The user is a contact of the current user and the current user is a contact of the user
	IsVerified        bool          `json:"is_verified"`        // True, if the user is verified
	IsSupport         bool          `json:"is_support"`         // True, if the user is Telegram support account
	RestrictionReason string        `json:"restriction_reason"` // If non-empty, it contains a human-readable description of the reason why access to this user must be restricted
	IsScam            bool          `json:"is_scam"`            // True, if many users reported this user as a scam
	IsFake            bool          `json:"is_fake"`            // True, if many users reported this user as a fake account
	HaveAccess        bool          `json:"have_access"`        // If false, the user is inaccessible, and the only information known about the user is inside this class. It can't be passed to any method except GetUser
	Type              UserType      `json:"type"`               // Type of the user
	LanguageCode      string        `json:"language_code"`      // IETF language tag of the user's language; only available to bots
}

// MessageType return the string telegram-type of User
func (user *User) MessageType() string {
	return "user"
}

// NewUser creates a new User
//
// @param iD User identifier
// @param firstName First name of the user
// @param lastName Last name of the user
// @param username Username of the user
// @param phoneNumber Phone number of the user
// @param status Current online status of the user
// @param profilePhoto Profile photo of the user; may be null
// @param isContact The user is a contact of the current user
// @param isMutualContact The user is a contact of the current user and the current user is a contact of the user
// @param isVerified True, if the user is verified
// @param isSupport True, if the user is Telegram support account
// @param restrictionReason If non-empty, it contains a human-readable description of the reason why access to this user must be restricted
// @param isScam True, if many users reported this user as a scam
// @param isFake True, if many users reported this user as a fake account
// @param haveAccess If false, the user is inaccessible, and the only information known about the user is inside this class. It can't be passed to any method except GetUser
// @param typeParam Type of the user
// @param languageCode IETF language tag of the user's language; only available to bots
func NewUser(iD int64, firstName string, lastName string, username string, phoneNumber string, status UserStatus, profilePhoto *ProfilePhoto, isContact bool, isMutualContact bool, isVerified bool, isSupport bool, restrictionReason string, isScam bool, isFake bool, haveAccess bool, typeParam UserType, languageCode string) *User {
	userTemp := User{
		tdCommon:          tdCommon{Type: "user"},
		ID:                iD,
		FirstName:         firstName,
		LastName:          lastName,
		Username:          username,
		PhoneNumber:       phoneNumber,
		Status:            status,
		ProfilePhoto:      profilePhoto,
		IsContact:         isContact,
		IsMutualContact:   isMutualContact,
		IsVerified:        isVerified,
		IsSupport:         isSupport,
		RestrictionReason: restrictionReason,
		IsScam:            isScam,
		IsFake:            isFake,
		HaveAccess:        haveAccess,
		Type:              typeParam,
		LanguageCode:      languageCode,
	}

	return &userTemp
}

// UnmarshalJSON unmarshal to json
func (user *User) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID                int64         `json:"id"`                 // User identifier
		FirstName         string        `json:"first_name"`         // First name of the user
		LastName          string        `json:"last_name"`          // Last name of the user
		Username          string        `json:"username"`           // Username of the user
		PhoneNumber       string        `json:"phone_number"`       // Phone number of the user
		ProfilePhoto      *ProfilePhoto `json:"profile_photo"`      // Profile photo of the user; may be null
		IsContact         bool          `json:"is_contact"`         // The user is a contact of the current user
		IsMutualContact   bool          `json:"is_mutual_contact"`  // The user is a contact of the current user and the current user is a contact of the user
		IsVerified        bool          `json:"is_verified"`        // True, if the user is verified
		IsSupport         bool          `json:"is_support"`         // True, if the user is Telegram support account
		RestrictionReason string        `json:"restriction_reason"` // If non-empty, it contains a human-readable description of the reason why access to this user must be restricted
		IsScam            bool          `json:"is_scam"`            // True, if many users reported this user as a scam
		IsFake            bool          `json:"is_fake"`            // True, if many users reported this user as a fake account
		HaveAccess        bool          `json:"have_access"`        // If false, the user is inaccessible, and the only information known about the user is inside this class. It can't be passed to any method except GetUser
		LanguageCode      string        `json:"language_code"`      // IETF language tag of the user's language; only available to bots
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	user.tdCommon = tempObj.tdCommon
	user.ID = tempObj.ID
	user.FirstName = tempObj.FirstName
	user.LastName = tempObj.LastName
	user.Username = tempObj.Username
	user.PhoneNumber = tempObj.PhoneNumber
	user.ProfilePhoto = tempObj.ProfilePhoto
	user.IsContact = tempObj.IsContact
	user.IsMutualContact = tempObj.IsMutualContact
	user.IsVerified = tempObj.IsVerified
	user.IsSupport = tempObj.IsSupport
	user.RestrictionReason = tempObj.RestrictionReason
	user.IsScam = tempObj.IsScam
	user.IsFake = tempObj.IsFake
	user.HaveAccess = tempObj.HaveAccess
	user.LanguageCode = tempObj.LanguageCode

	fieldStatus, _ := unmarshalUserStatus(objMap["status"])
	user.Status = fieldStatus

	fieldType, _ := unmarshalUserType(objMap["type"])
	user.Type = fieldType

	return nil
}
