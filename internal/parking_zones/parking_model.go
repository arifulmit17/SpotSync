package parking_zones

type parking struct {
	Id             int    `json:"uuid"`
	Name           string `json:"name"`
	Total_Capacity int    `json:"total_capacity"`
	Location       int    `json:"location"`
}
