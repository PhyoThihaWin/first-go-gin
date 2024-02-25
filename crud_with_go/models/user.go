package models

type User struct {
	CustomModel
	USERNAME  string `json:"username"`
	DEVICE_ID string `json:"device_id"`
}
