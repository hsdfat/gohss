package models

type PublicId struct {
	ID        string `json:"publicId,omitempty"`
	State     int    `json:"state"`
	PrivateId string `json:"privateId,omitempty"`
}

type PrivateId struct {
	ID         string `json:"privateId,omitempty"`
	UeId       string `json:"ueId,omitempty"`
	ServerName string `json:"serverName"`
}
