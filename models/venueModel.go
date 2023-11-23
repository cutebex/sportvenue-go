package models

// venue model
type Venue struct {
	ID       int
	Name     string
	Type     string
	Date     string
	Location string
}

func NewVenueFromMap(venueMap map[string]interface{}) Venue {
	// Assuming Id, Name and others are the fields in the map
	id, _ := venueMap["Id"].(int)
	name, _ := venueMap["Name"].(string)
	venuetype, _ := venueMap["Type"].(string)
	date, _ := venueMap["Date"].(string)
	venuelocation, _ := venueMap["Location"].(string)

	return Venue{
		ID:       id,
		Name:     name,
		Type:     venuetype,
		Date:     date,
		Location: venuelocation,
	}
}