package models

type TestStruct struct {
	X1 string `json:"x1"`
	X2 string `json:"x2"`
}

type Account struct {
	AccountID int                `bson:"account_id"`
	Limit     int                `bson:"limit"`
	Products  []string           `bson:"products"`
}