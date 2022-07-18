// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/satyshef/go-tdlib/tdlib"
)

// GetStickerSet Returns information about a sticker set by its identifier
// @param setID Identifier of the sticker set
func (client *Client) GetStickerSet(setID *tdlib.JSONInt64) (*tdlib.StickerSet, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":  "getStickerSet",
		"set_id": setID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var stickerSet tdlib.StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SearchStickerSet Searches for a sticker set by its name
// @param name Name of the sticker set
func (client *Client) SearchStickerSet(name string) (*tdlib.StickerSet, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "searchStickerSet",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var stickerSet tdlib.StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// CreateNewStickerSet Creates a new sticker set. Returns the newly created sticker set
// @param userID Sticker set owner; ignored for regular users
// @param title Sticker set title; 1-64 characters
// @param name Sticker set name. Can contain only English letters, digits and underscores. Must end with *"_by_<bot username>"* (*<bot_username>* is case insensitive) for bots; 1-64 characters
// @param isMasks True, if stickers are masks. Animated stickers can't be masks
// @param stickers List of stickers to be added to the set; must be non-empty. All stickers must be of the same type. For animated stickers, uploadStickerFile must be used before the sticker is shown
// @param source Source of the sticker set; may be empty if unknown
func (client *Client) CreateNewStickerSet(userID int64, title string, name string, isMasks bool, stickers []tdlib.InputSticker, source string) (*tdlib.StickerSet, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "createNewStickerSet",
		"user_id":  userID,
		"title":    title,
		"name":     name,
		"is_masks": isMasks,
		"stickers": stickers,
		"source":   source,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var stickerSet tdlib.StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// AddStickerToSet Adds a new sticker to a set; for bots only. Returns the sticker set
// @param userID Sticker set owner
// @param name Sticker set name
// @param sticker Sticker to add to the set
func (client *Client) AddStickerToSet(userID int64, name string, sticker tdlib.InputSticker) (*tdlib.StickerSet, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "addStickerToSet",
		"user_id": userID,
		"name":    name,
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var stickerSet tdlib.StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SetStickerSetThumbnail Sets a sticker set thumbnail; for bots only. Returns the sticker set
// @param userID Sticker set owner
// @param name Sticker set name
// @param thumbnail Thumbnail to set in PNG or TGS format. Animated thumbnail must be set for animated sticker sets and only for them. Pass a zero InputFileId to delete the thumbnail
func (client *Client) SetStickerSetThumbnail(userID int64, name string, thumbnail tdlib.InputFile) (*tdlib.StickerSet, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":     "setStickerSetThumbnail",
		"user_id":   userID,
		"name":      name,
		"thumbnail": thumbnail,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var stickerSet tdlib.StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}
