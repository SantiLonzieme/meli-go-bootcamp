package models

type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"tipo"`
	Count        int     `json:"cantidad"`
	Price        float64 `json:"precio"`
	Id_Warehouse int     `json:"id_warehouse"`
}
