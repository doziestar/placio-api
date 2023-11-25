package order

import "placio-app/ent"

type OrderWithItemsDTO struct {
    Order ent.Order `json:"order"`
    Items map[string]int `json:"items"`
}


