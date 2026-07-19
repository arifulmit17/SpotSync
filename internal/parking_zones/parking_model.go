package parking_zones

type parking struct {
	Id             int     `json:"uuid"`
	Name           string  `json:"name"`
	Total_Capacity uint    `json:"total_capacity"  gorm:"check:total_capacity > 0"`
	Location       int     `json:"location"`
	Price_per_hour float64 `json:"price_per_hour" gorm:"check:price_per_hour > 0"`
	Created_at     string  `json:"created_at"`
	Updated_at     string  `json:"updated_at"`
}
