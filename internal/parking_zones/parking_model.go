package parking_zones

type parking struct {
	Id             int    `json:"uuid"`
	Name           string `json:"name"`
	Total_Capacity uint   `json:"total_capacity"  gorm:"check:total_capacity > 0"`
	Location       int    `json:"location"`
}
