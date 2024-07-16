package api

// Texts in backticks represent tags
// so that it can serialize into json with those tag names as key names in JSON

// ReceiptItem represents an item in the receipt
type ReceiptItem struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Receipt struct represents the data structure of a receipt
type Receipt struct {
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Items        []ReceiptItem `json:"items"`
	Total        string        `json:"total"`
}

// ReceiptPoints struct represents the points response
type ReceiptPoints struct {
	Points int `json:"points"`
}
