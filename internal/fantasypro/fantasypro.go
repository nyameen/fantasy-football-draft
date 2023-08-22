package fantasypro

import (
	"encoding/csv"
	"net/http"
)

func GetFantasyProCSV() ([][]string, error) {
	resp, err := http.Get("https://partners.fantasypros.com/api/v1/consensus-rankings.php?sport=NFL&year=2023&week=0&id=1663&position=ALL&type=PPR&scoring=HALF&export=csv")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}
