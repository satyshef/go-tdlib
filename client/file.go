// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/satyshef/go-tdlib/tdlib"
)

// GetFile Returns information about a file; this is an offline request
// @param fileID Identifier of the file to get
func (client *Client) GetFile(fileID int32) (*tdlib.File, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getFile",
		"file_id": fileID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var fileDummy tdlib.File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// GetRemoteFile Returns information about a file by its remote ID; this is an offline request. Can be used to register a URL as a file for further uploading, or sending as a message. Even the request succeeds, the file can be used only if it is still accessible to the user. For example, if the file is from a message, then the message must be not deleted and accessible to the user. If the file database is disabled, then the corresponding object with the file must be preloaded by the application
// @param remoteFileID Remote identifier of the file to get
// @param fileType File type, if known
func (client *Client) GetRemoteFile(remoteFileID string, fileType tdlib.FileType) (*tdlib.File, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "getRemoteFile",
		"remote_file_id": remoteFileID,
		"file_type":      fileType,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var fileDummy tdlib.File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// DownloadFile Downloads a file from the cloud. Download progress and completion of the download will be notified through updateFile updates
// @param fileID Identifier of the file to download
// @param priority Priority of the download (1-32). The higher the priority, the earlier the file will be downloaded. If the priorities of two files are equal, then the last one for which downloadFile was called will be downloaded first
// @param offset The starting position from which the file should be downloaded
// @param limit Number of bytes which should be downloaded starting from the "offset" position before the download will be automatically canceled; use 0 to download without a limit
// @param synchronous If false, this request returns file state just after the download has been started. If true, this request returns file state only after the download has succeeded, has failed, has been canceled or a new downloadFile request with different offset/limit parameters was sent
func (client *Client) DownloadFile(fileID int32, priority int32, offset int32, limit int32, synchronous bool) (*tdlib.File, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":       "downloadFile",
		"file_id":     fileID,
		"priority":    priority,
		"offset":      offset,
		"limit":       limit,
		"synchronous": synchronous,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var fileDummy tdlib.File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// UploadFile Asynchronously uploads a file to the cloud without sending it in a message. updateFile will be used to notify about upload progress and successful completion of the upload. The file will not have a persistent remote identifier until it will be sent in a message
// @param file File to upload
// @param fileType File type
// @param priority Priority of the upload (1-32). The higher the priority, the earlier the file will be uploaded. If the priorities of two files are equal, then the first one for which uploadFile was called will be uploaded first
func (client *Client) UploadFile(file tdlib.InputFile, fileType tdlib.FileType, priority int32) (*tdlib.File, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":     "uploadFile",
		"file":      file,
		"file_type": fileType,
		"priority":  priority,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var fileDummy tdlib.File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// UploadStickerFile Uploads a PNG image with a sticker; returns the uploaded file
// @param userID Sticker file owner; ignored for regular users
// @param sticker Sticker file to upload
func (client *Client) UploadStickerFile(userID int64, sticker tdlib.InputSticker) (*tdlib.File, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "uploadStickerFile",
		"user_id": userID,
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var file tdlib.File
	err = json.Unmarshal(result.Raw, &file)
	return &file, err

}

// GetMapThumbnailFile Returns information about a file with a map thumbnail in PNG format. Only map thumbnail files with size less than 1MB can be downloaded
// @param location Location of the map center
// @param zoom Map zoom level; 13-20
// @param width Map width in pixels before applying scale; 16-1024
// @param height Map height in pixels before applying scale; 16-1024
// @param scale Map scale; 1-3
// @param chatID Identifier of a chat, in which the thumbnail will be shown. Use 0 if unknown
func (client *Client) GetMapThumbnailFile(location *tdlib.Location, zoom int32, width int32, height int32, scale int32, chatID int64) (*tdlib.File, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "getMapThumbnailFile",
		"location": location,
		"zoom":     zoom,
		"width":    width,
		"height":   height,
		"scale":    scale,
		"chat_id":  chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var file tdlib.File
	err = json.Unmarshal(result.Raw, &file)
	return &file, err

}
