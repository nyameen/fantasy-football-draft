package nflplayers

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// NFLPlayer description of a player
type NFLPlayer struct {
	Rank     string
	Name     string
	Team     string
	Position string
	ByeWeek  string
}

// NFLPlayers collection of individual players
type NFLPlayers []NFLPlayer

// NFL_POSITION type of position
type NFL_POSITION string

var (
	ALL          NFL_POSITION = "ALL"
	QuarterBack  NFL_POSITION = "QB"
	WideReceiver NFL_POSITION = "WR"
	RunningBack  NFL_POSITION = "RB"
	TightEnd     NFL_POSITION = "TE"
	Defense      NFL_POSITION = "DST"
	Kicker       NFL_POSITION = "K"
)

// GetNFLPlayersFromCSV extracts players from a CSV format
func GetNFLPlayersFromCSV(playersCSV [][]string) NFLPlayers {
	players := []NFLPlayer{}

	for _, playerCSV := range playersCSV {
		for _, playerInfo := range playerCSV {
			playerStr := strings.Split(playerInfo, ",")

			player := NFLPlayer{
				Rank:     playerStr[0],
				Name:     playerStr[1],
				Team:     playerStr[2],
				Position: playerStr[3],
				ByeWeek:  getByeWeek((playerStr[2])),
			}
			players = append(players, player)
		}
	}

	return players
}

func getByeWeek(team string) string {
	switch team {
	case "CLE", "LAC", "SEA", "TB":
		return "5"
	case "GB", "PIT":
		return "6"
	case "CAR", "CIN", "DAL", "HOU", "NYJ", "TEN":
		return "7"
	case "DEN", "DET", "JAX", "SF", "JAC":
		return "9"
	case "KC", "LAR", "MIA", "PHI":
		return "10"
	case "ATL", "IND", "NE", "NO":
		return "11"
	case "BAL", "BUF", "CHI", "LV", "MIN", "NYG":
		return "13"
	case "ARI", "WAS":
		return "14"
	}

	return ""
}

func (n NFLPlayers) getPlayersByPosition(position NFL_POSITION) NFLPlayers {
	posPLayers := []NFLPlayer{}

	for _, player := range n {
		if player.Position == string(position) {
			posPLayers = append(posPLayers, player)
		}
	}

	return posPLayers
}

// CreateTableCallbackByPosition creates a callback function for creating tables of players
func (n NFLPlayers) CreateTableCallbackByPosition(position NFL_POSITION) func(w fyne.Window) fyne.CanvasObject {
	var players NFLPlayers
	if position == ALL {
		players = n
	} else {
		players = n.getPlayersByPosition(position)
	}

	return func(w fyne.Window) fyne.CanvasObject {
		table := widget.NewTable(
			func() (int, int) {
				return len(players), 5
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("label")
			},
			func(i widget.TableCellID, o fyne.CanvasObject) {
				switch i.Col {
				case 0:
					o.(*widget.Label).SetText(players[i.Row].Rank)
				case 1:
					o.(*widget.Label).SetText(players[i.Row].Name)
				case 2:
					o.(*widget.Label).SetText(players[i.Row].Position)
				case 3:
					o.(*widget.Label).SetText(players[i.Row].Team)
				case 4:
					o.(*widget.Label).SetText(players[i.Row].ByeWeek)
				}
			},
		)

		table.SetColumnWidth(0, 50)
		table.SetColumnWidth(1, 250)
		table.SetColumnWidth(2, 100)
		table.SetColumnWidth(3, 100)
		return table
	}
}
