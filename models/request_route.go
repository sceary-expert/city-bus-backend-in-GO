package models

type RequestRoute struct {
	StartingPointName string `json:"starting-point-name"`
	EndingPointName   string `json:"ending-point-name"`
}
