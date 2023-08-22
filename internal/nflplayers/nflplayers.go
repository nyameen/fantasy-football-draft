package nflplayers

import (
	"strings"
)

type NFLPlayers struct {
	Rank     string
	Name     string
	Team     string
	Position string
}

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
			}
			players = append(players, player)
		}
	}

	return players
}
