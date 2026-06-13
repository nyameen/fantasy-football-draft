package fantasypro

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"github.com/gocolly/colly/v2"
)

var teamMap = map[string]string{
	"Arizona":       "ARI",
	"Atlanta":       "ATL",
	"Baltimore":     "BAL",
	"Buffalo":       "BUF",
	"Carolina":      "CAR",
	"Chicago":       "CHI",
	"Cincinnati":    "CIN",
	"Cleveland":     "CLE",
	"Dallas":        "DAL",
	"Denver":        "DEN",
	"Detroit":       "DET",
	"Green Bay":     "GB",
	"Houston":       "HOU",
	"Indianapolis":  "IND",
	"Jacksonville":  "JAC",
	"Kansas City":   "KC",
	"Las Vegas":     "LV",
	"LA Chargers":   "LAC",
	"LA Rams":       "LAR",
	"Miami":         "MIA",
	"Minnesota":     "MIN",
	"New England":   "NE",
	"New Orleans":   "NO",
	"NY Giants":     "NYG",
	"NY Jets":       "NYJ",
	"Philadelphia":  "PHI",
	"Pittsburgh":    "PIT",
	"San Francisco": "SF",
	"Seattle":       "SEA",
	"Tampa Bay":     "TB",
	"Tennessee":     "TEN",
	"Washington":    "WAS",
}

type Bye struct {
	Week  string
	Teams []string
}

// GetFantasyProCSV get a list of players from Fantasy Pros in CSV format
func GetFantasyProCSV() ([][]string, error) {
	year, _, _ := time.Now().Date()
	url := fmt.Sprintf("https://partners.fantasypros.com/api/v1/consensus-rankings.php?sport=NFL&year=%d&week=0&position=ALL&type=PPR&scoring=FULL&export=csv", year)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'

	return reader.ReadAll()
}

// GetByeWeeks get the teams bye week based on year
func GetByeWeeks() []Bye {
	var byes []Bye

	c := colly.NewCollector()

	// Target the tbody within the specific table (e.g., matching by ID)
	c.OnHTML("table#nfl-bye-weeks tbody", func(e *colly.HTMLElement) {

		if e.Index > 0 {
			return
		}
		// Loop through each table row (tr)
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			b := Bye{}

			// Loop through each cell (td) in the row
			row.ForEach("td", func(index int, cell *colly.HTMLElement) {
				switch index {
				case 0:
					b.Week = cell.Text
				case 1:
					b.Teams = strings.Split(cell.Text, ",")
					for i, team := range b.Teams {
						b.Teams[i] = strings.TrimSpace(team)
					}
				default:
				}
			})

			byes = append(byes, b)
		})
	})

	year, _, _ := time.Now().Date()
	url := fmt.Sprintf("https://gridirongames.com/football-schedules/nfl-bye-weeks-schedule/?Year=%d", year)
	if err := c.Visit(url); err != nil {
		fyne.LogError("could not get Bye weeks: ", err)
	}

	for _, b := range byes {
		convertNameToAbbreviation(b.Teams)
	}

	return byes
}

func convertNameToAbbreviation(teams []string) {
	for i, team := range teams {
		teams[i] = teamMap[team]
	}
}
