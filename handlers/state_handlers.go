package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type StateData struct {
	States map[string][]string `json:"states"`
}

var data StateData

// LoadData loads the JSON data into memory
func LoadData() error {
	//cwd, _ := os.Getwd()
	//log.Println("Current working directory:", cwd)

	file, err := os.Open("data/states.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}
	return nil
}

// GetAllStates returns the list of all states
func GetAllStates(c *gin.Context) {
	states := make([]string, 0, len(data.States))
	for state := range data.States {
		states = append(states, state)
	}
	c.JSON(http.StatusOK, gin.H{"states": states})
}

// GetStateLGAs returns the LGAs of a specified state
func GetStateLGAs(c *gin.Context) {
	state := c.Param("state")
	lgas, exists := data.States[state]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "State not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"state": state, "lgas": lgas})
}


// SearchStatesAndLGAs handles searching for states or local governments
func SearchStatesAndLGAs(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	query = strings.ToLower(query) // Convert query to lowercase for case-insensitive search
	results := make(map[string][]string)

	// Search through the states and their LGAs
	for state, lgas := range data.States {
		if strings.Contains(strings.ToLower(state), query) {
			results[state] = lgas // Match state
			continue
		}
		// Check for matches in LGAs
		for _, lga := range lgas {
			if strings.Contains(strings.ToLower(lga), query) {
				results[state] = append(results[state], lga)
			}
		}
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No results found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}
