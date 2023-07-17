package model

type ShoppingCart struct {
	Items []*ShoppingCartItem `json:"items"`
}

type ShoppingCartItem struct {
	Id     int     `json:"id"`
	Recipe *Recipe `json:"recipe"`
}
