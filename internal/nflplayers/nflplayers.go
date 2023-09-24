package nflplayers

import (
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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
type NFLPlayers struct {
	allPlayers []NFLPlayer
	qbs        []NFLPlayer
	rbs        []NFLPlayer
	wrs        []NFLPlayer
	tes        []NFLPlayer
	def        []NFLPlayer
	kickers    []NFLPlayer
	mux        sync.Mutex
}

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
func GetNFLPlayersFromCSV(playersCSV [][]string) *NFLPlayers {
	players := &NFLPlayers{}

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
			players.allPlayers = append(players.allPlayers, player)
		}
	}

	players.qbs = players.sortPlayersByPosition(QuarterBack)
	players.rbs = players.sortPlayersByPosition(RunningBack)
	players.wrs = players.sortPlayersByPosition(WideReceiver)
	players.tes = players.sortPlayersByPosition(TightEnd)
	players.def = players.sortPlayersByPosition(Defense)
	players.kickers = players.sortPlayersByPosition(Kicker)

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

func (n *NFLPlayers) sortPlayersByPosition(position NFL_POSITION) []NFLPlayer {
	posPlayers := []NFLPlayer{}

	for _, player := range n.allPlayers {
		if player.Position == string(position) {
			posPlayers = append(posPlayers, player)
		}
	}

	return posPlayers
}

func (n *NFLPlayers) getPlayers(position NFL_POSITION) []NFLPlayer {
	switch position {
	case QuarterBack:
		return n.qbs
	case WideReceiver:
		return n.wrs
	case RunningBack:
		return n.rbs
	case TightEnd:
		return n.tes
	case Kicker:
		return n.kickers
	case Defense:
		return n.def
	}

	return n.allPlayers
}

// CreateTableCallbackByPosition creates a callback function for creating tables of players
func (n *NFLPlayers) CreateTableCallbackByPosition(position NFL_POSITION) func(w fyne.Window) fyne.CanvasObject {
	players := n.getPlayers(position)

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
		table.CreateHeader = func() fyne.CanvasObject {
			return widget.NewLabel("Header")
		}
		table.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
			label, ok := template.(*widget.Label)
			if !ok {
				return
			}
			if id.Row == -1 {
				switch id.Col {
				case 0:
					label.SetText("Rank")
				case 1:
					label.SetText("Player Name")
				case 2:
					label.SetText("Position")
				case 3:
					label.SetText("Team")
				case 4:
					label.SetText("Bye Week")
				}

				label.TextStyle.Bold = true
			}
		}
		table.ShowHeaderRow = true

		table.SetColumnWidth(0, 50)
		table.SetColumnWidth(1, 250)
		table.SetColumnWidth(2, 100)
		table.SetColumnWidth(3, 100)

		table.OnSelected = func(id widget.TableCellID) {
			table.Unselect(id)
			if id.Col != 1 {
				// user should click the name to get the option to remove the player
				// unless every column would prompt a remove player screen
				return
			}

			n.createPopUp(fyne.CurrentApp().Driver().AllWindows()[0], players[id.Row].Name, id.Row, table)
		}

		return table
	}
}

func (n *NFLPlayers) createPopUp(w fyne.Window, playerName string, index int, table *widget.Table) {
	var modal *widget.PopUp

	// Split box to ask if we want to remove a player
	split := container.NewVBox(widget.NewLabel("Remove "+playerName+"?"), container.NewHBox(
		widget.NewButton("Yes", func() {
			n.removePlayerFromAll(playerName)
			modal.Hide()
			table.Refresh()
		}),
		widget.NewSeparator(),
		widget.NewButton("No", func() {
			modal.Hide()
		}),
	))

	modal = widget.NewModalPopUp(
		split,
		w.Canvas(),
	)

	modal.Show()
}

func (n *NFLPlayers) removePlayerFromAll(name string) {
	// in order to get the tables and sub tables to remove correctly
	// we must remove the player from allPlayers then remove that player
	// from his sub group
	n.mux.Lock()
	defer n.mux.Unlock()

	index, position := n.findPlayer(n.allPlayers, name)
	if index != -1 {
		n.allPlayers = append(n.allPlayers[:index], n.allPlayers[index+1:]...)
	}

	switch position {
	case string(QuarterBack):
		n.qbs = n.removePlayerByPosition(n.qbs, name)
	case string(WideReceiver):
		n.wrs = n.removePlayerByPosition(n.wrs, name)
	case string(RunningBack):
		n.rbs = n.removePlayerByPosition(n.rbs, name)
	case string(TightEnd):
		n.tes = n.removePlayerByPosition(n.tes, name)
	case string(Kicker):
		n.kickers = n.removePlayerByPosition(n.kickers, name)
	case string(Defense):
		n.def = n.removePlayerByPosition(n.def, name)
	}
}

func (n *NFLPlayers) findPlayer(players []NFLPlayer, name string) (int, string) {
	for i, player := range players {
		if player.Name == name {
			return i, player.Position
		}
	}

	return -1, ""
}

func (n *NFLPlayers) removePlayerByPosition(posPlayers []NFLPlayer, name string) []NFLPlayer {
	newPlayers := []NFLPlayer{}
	index, _ := n.findPlayer(posPlayers, name)

	if index != -1 {
		newPlayers = append(posPlayers[:index], posPlayers[index+1:]...)
	}

	return newPlayers
}
