package controllers

import (
	"os"
	"fmt"
	"strings"
	"path/filepath"

	"github.com/cutebex/sportvenue-server/database"
	"github.com/cutebex/sportvenue-server/models"
	"github.com/gin-gonic/gin"
)

// VenueController for VenueType
type VenueController struct{}

// Search Venue Type
func (V *VenueController) Get(c *gin.Context) {
	var (
		venues []models.Venue
	)

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the current working directory:", err)
		return
	}

	// Create a relative file path
	relativePath := "database/mockdata.json"
	fullPath := filepath.Join(currentDir, relativePath)

	fmt.Println("Full path:", fullPath)

	mockData, err := database.ReadJSONFromFile(fullPath)

	if err != nil {
		c.JSON(400, "An error is occured reading the file.")
	}

	for _, venueMap := range mockData {
		// Assuming models.Venue has a constructor function NewVenueFromMap
		venue := models.NewVenueFromMap(venueMap)
		venues = append(venues, venue)
	}

	keyword := c.Param("keyword")

	// Assuming models.Venue has a field named Type
	filteredVenues := V.filterVenuesByType(venues, keyword)

	// Do something with the filteredVenues, such as returning it as JSON
	c.JSON(200, filteredVenues)
}

// filterVenuesByType filters venues based on the keyword in the Venue.Type field
func (V *VenueController) filterVenuesByType(venues []models.Venue, keyword string) []models.Venue {
	var filteredVenues []models.Venue

	for _, venue := range venues {
		if strings.Contains(strings.ToLower(venue.Type), strings.ToLower(keyword)) {
			// Add the venue to the filteredVenues if the keyword is found in the Type field
			filteredVenues = append(filteredVenues, venue)
		}
	}

	return filteredVenues
}
