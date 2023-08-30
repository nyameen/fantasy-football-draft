package fantasypro

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"time"
)

// GetFantasyProCSV get a list of players from Fantasy Pros in CSV format
func GetFantasyProCSV() ([][]string, error) {
	year, _, _ := time.Now().Date()
	url := fmt.Sprintf("https://partners.fantasypros.com/api/v1/consensus-rankings.php?sport=NFL&year=%d&week=0&position=ALL&type=PPR&scoring=HALF&export=csv", year)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'

	return reader.ReadAll()
}
