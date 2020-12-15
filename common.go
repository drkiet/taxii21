package main

import (
	"encoding/json"
	"github.com/drkiet/stix21"
	"time"
)

type ApiRoot struct{
	Title string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
	Versions []string `json:"versions" binding:"required"`
	MaxContentLength int `json:"max_content_length" binding:"required"`
}

type Collection struct {
	Id stix21.Identifier `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Alias string `json:"alias,omitempty"`
	CanRead bool `json:"can_read" binding:"required"`
	CanWrite bool `json:"can_write" binding:"required"`
	MediaType []string `json:"media_types,omitempty"`
}

type Collections struct {
	Collections []Collection `json:"collections,omitempty"`
}

type Discovery struct {
	Title string `json:"discovery" binding:"required"`
	Description string `json:"description,omitempty"`
	Contact string `json:"contact,omitempty"`
	Default string `json:"default,omitempty"`
	ApiRoots []string `json:"api_roots,omitempty"`
}

type Envelop struct {
	More bool `json:"more,omitempty"`
	Next string `json:"next,omitempty"`
	Objects []stix21.STIXObject `json:"objects,omitempty"`
}

type Error struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
	ErrorId string `json:"error_id,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
	HttpStatus string `json:"http_status,omitempty"`
	ExternalDetails string `json:"external_details,omitempty"`
	Details string `json:"details,omitempty"`
}

type Manifest struct {
	More bool `json:"more,omitempty"`
	Objects []ManifestRecord `json:"objects,omitempty"`
}

type ManifestRecord struct {
	Id string `json:"id" binding:"required"`
	DateAdded time.Time `json:"date_added" binding:"required"`
	Version string `json:"version" binding:"required"`
	MediaType string `json:"media_type,omitempty"`
}

type Status struct {
	Id stix21.Identifier `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
	RequestTimestamp time.Time `json:"request_timestamp,omitempty"`
	TotalCount int `json:"total_count" binding:"required"`
	SuccessCount int `json:"success_count" binding:"required"`
	Successes []StatusDetail `json:"successes,omitempty"`
	FailureCount int `json:"failure_count" binding:"required"`
	Failures []StatusDetail `json:"failures,omitempty"`
	PendingCount int `json:"pending_count" binding:"required"`
	Pendings []StatusDetail `json:"pending,omitempty"`
}

type StatusDetail struct {
	Id string `json:"id" binding:"required"`
	Version string `json:"version" binding:"required"`
	Message string `json:"message,omitempty"`
}

type Versions struct {
	More bool `json:"more,omitempty"`
	Versions []string `json:"versions,omitempty"`
}

func MarshalError(error Error) (jsonData string){
	data, _ := json.MarshalIndent(error, "", "  ")
	return string(data)
}