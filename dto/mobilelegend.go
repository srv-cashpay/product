package dto

type MobileLegendRequest struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	BuyerSkuCode string `json:"buyer_sku_code"`
	CustomerNo   string `json:"customer_no"`
	RefID        string `json:"ref_id"`
	Sign         string `json:"sign"`
}

type MobileLegendResponse struct {
	Data MobileLegendData `json:"data"`
}

type MobileLegendData struct {
	RefID        string `json:"ref_id"`
	CustomerNo   string `json:"customer_no"`
	BuyerSkuCode string `json:"buyer_sku_code"`
	Message      string `json:"message"`
	Status       string `json:"status"`
	Rc           string `json:"rc"`
	Sn           string `json:"sn"`
}
