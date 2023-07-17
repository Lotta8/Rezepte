package model

type ShoppingCart struct {
	Items []*ShoppingCartItem `json:"items"`
}

type ShoppingCartItem struct {
	Id     int     `json:"id"`
	Count  int     `json:"count"`
	Recipe *Recipe `json:"recipe"`
}
