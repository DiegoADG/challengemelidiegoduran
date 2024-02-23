package dto

type ConsultarUserDto struct {
	Seller SellerDto `json:"seller"`
}

type SellerDto struct {
	Nickname string `json:"nickname"`
}
