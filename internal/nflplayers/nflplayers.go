package nflplayers

import (
	"strings"
)

type NFLPlayers struct {
	Rank     string
	Name     string
	Team     string
	Position string
	ByeWeek  string
}

type NFL_POSITION string

var (
	QuarterBack  NFL_POSITION = "QB"
	WideReceiver NFL_POSITION = "WR"
	RunningBack  NFL_POSITION = "RB"
	TightEnd     NFL_POSITION = "TE"
	Defense      NFL_POSITION = "DST"
	Kicker       NFL_POSITION = "K"
)

func GetNFLPlayersFromCSV(playersCSV [][]string) []NFLPlayers {
	players := []NFLPlayers{}

	for _, playerCSV := range playersCSV {
		for _, playerInfo := range playerCSV {
			playerStr := strings.Split(playerInfo, ",")

			player := NFLPlayers{
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

func GetPlayersByPosition(position string, players []NFLPlayers) []NFLPlayers {
	posPLayers := []NFLPlayers{}

	for _, player := range players {
		if player.Position == position {
			posPLayers = append(posPLayers, player)
		}
	}

	return posPLayers
}
