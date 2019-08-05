package models

// ImageModel ...
type ImageModel struct {
	Index            int    `json:"index"`
	LastModified     string `json:"lastModified,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	Name             string `json:"name,omitempty"`
	Size             string `json:"size,omitempty"`
	Title            string `json:"title,omitempty"`
	Type             string `json:"type,omitempty"`
	URL              string `json:"url,omitempty"`
	Username         string `json:"username,omitempty"`
}
