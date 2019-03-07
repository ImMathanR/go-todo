package model

import (
	"encoding/json"
	"time"
)

// Crash -- Model to take all the values
type Crash struct {
	AI5            string    `json:"ai5"` // header
	AppN           string    `json:"appn"`
	AppV           string    `json:"appv"`
	AndroidVersion string    `json:"android_v"`
	CampaignID     string    `json:"campaign_id"`
	GameID         string    `json:"game_id"` // header
	Model          string    `json:"model"`
	Stacktrace     string    `json:"stacktrace"`
	SDKV           string    `json:"sdv"`   //header
	SDKN           string    `json:"sdkn"`  //header
	AdvID          string    `json:"advid"` //header
	Platform       string    `json:"platform"`
	CrashTs        string    `json:"crash_ts"`
	Tag            string    `json:"tag"`
	NonFatal       bool      `json:"non_fatal"`
	CountryID      string    `json:"country_id"`
	CreatedAt      time.Time `json:"created_at"`
	IP             string    `json:"ip"`
}

// ToString -- To serialize the Tracker model string.
func (c *Crash) ToString() string {
	output, _ := json.Marshal(c)
	return string(output)
}
