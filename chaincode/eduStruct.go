package main

type Commodity struct {
	ObjectType	string	`json:"docType"`
	Name string `json:"Name"`
	EntityID string `json:"EntityID"`
	FileHash string `json:"FileHash"`
	FileType string `json:"FileType"`
	Desc string `json:"Desc"`
	Company string `json:"Company"`

	Historys []HistoryItem
}

type HistoryItem struct {
	TxId string
	Commodity Commodity
}