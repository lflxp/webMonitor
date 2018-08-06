package models

type Bucket struct {
	Tablename string `json:"tablename"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}
