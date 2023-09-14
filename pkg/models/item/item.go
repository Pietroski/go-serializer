package item_models

type SubItem struct {
	Date     int64  `json:"date,omitempty"`
	Amount   int64  `json:"amount,omitempty"`
	ItemCode string `json:"itemCode,omitempty"`
}

type Item struct {
	Id      string   `json:"id,omitempty"`
	ItemId  uint64   `json:"itemId,omitempty"`
	Number  int64    `json:"number,omitempty"`
	SubItem *SubItem `json:"subItem,omitempty"`
}
