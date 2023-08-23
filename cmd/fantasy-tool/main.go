package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/nyameen/fantasy-football-draft/internal/fantasypro"
	"github.com/nyameen/fantasy-football-draft/internal/nflplayers"
)

var (
	// Side menu tabs
	Index = map[string][]string{
		"": {"all", "qb", "rb", "wr", "te", "k", "def"},
	}
)

type SideMenu struct {
	Title       string
	Description string
	View        func(w fyne.Window) fyne.CanvasObject
}

func main() {

	data, err := fantasypro.GetFantasyProCSV()
	if err != nil {
		fmt.Println("Error, could not get CSV: ", err)
		return
	}

	players := nflplayers.GetNFLPlayersFromCSV(data[3:])

	myApp := app.New()
	myWindow := myApp.NewWindow("Fantasy Football Draft Tool")

	// The menu with menu Title, Description, and Table of players
	menu := map[string]SideMenu{
		"all": {"All", "Highest ranking players overall", func(w fyne.Window) fyne.CanvasObject {
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
		}},
		"qb": {"QB", "Highest Ranking Quarter Backs", func(w fyne.Window) fyne.CanvasObject {
			table := widget.NewTable(
				func() (int, int) {
					return len(nflplayers.GetPlayersByPosition("QB", players)), 5
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("label")
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					posPlayers := nflplayers.GetPlayersByPosition("QB", players)
					switch i.Col {
					case 0:
						o.(*widget.Label).SetText(posPlayers[i.Row].Rank)
					case 1:
						o.(*widget.Label).SetText(posPlayers[i.Row].Name)
					case 2:
						o.(*widget.Label).SetText(posPlayers[i.Row].Position)
					case 3:
						o.(*widget.Label).SetText(posPlayers[i.Row].Team)
					case 4:
						o.(*widget.Label).SetText(posPlayers[i.Row].ByeWeek)
					}
				},
			)

			table.SetColumnWidth(0, 50)
			table.SetColumnWidth(1, 250)
			table.SetColumnWidth(2, 100)
			table.SetColumnWidth(3, 100)
			return table
		}},
		"wr": {"WR", "Highest Ranking Wide Receivers", func(w fyne.Window) fyne.CanvasObject {
			table := widget.NewTable(
				func() (int, int) {
					return len(nflplayers.GetPlayersByPosition("WR", players)), 5
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("label")
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					posPlayers := nflplayers.GetPlayersByPosition("WR", players)
					switch i.Col {
					case 0:
						o.(*widget.Label).SetText(posPlayers[i.Row].Rank)
					case 1:
						o.(*widget.Label).SetText(posPlayers[i.Row].Name)
					case 2:
						o.(*widget.Label).SetText(posPlayers[i.Row].Position)
					case 3:
						o.(*widget.Label).SetText(posPlayers[i.Row].Team)
					case 4:
						o.(*widget.Label).SetText(posPlayers[i.Row].ByeWeek)
					}
				},
			)

			table.SetColumnWidth(0, 50)
			table.SetColumnWidth(1, 250)
			table.SetColumnWidth(2, 100)
			table.SetColumnWidth(3, 100)
			return table
		}},
		"rb": {"RB", "Highest Ranking Running Backs", func(w fyne.Window) fyne.CanvasObject {
			table := widget.NewTable(
				func() (int, int) {
					return len(nflplayers.GetPlayersByPosition("RB", players)), 5
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("label")
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					posPlayers := nflplayers.GetPlayersByPosition("RB", players)
					switch i.Col {
					case 0:
						o.(*widget.Label).SetText(posPlayers[i.Row].Rank)
					case 1:
						o.(*widget.Label).SetText(posPlayers[i.Row].Name)
					case 2:
						o.(*widget.Label).SetText(posPlayers[i.Row].Position)
					case 3:
						o.(*widget.Label).SetText(posPlayers[i.Row].Team)
					case 4:
						o.(*widget.Label).SetText(posPlayers[i.Row].ByeWeek)
					}
				},
			)

			table.SetColumnWidth(0, 50)
			table.SetColumnWidth(1, 250)
			table.SetColumnWidth(2, 100)
			table.SetColumnWidth(3, 100)
			return table
		}},
		"te": {"TE", "Highest Ranking Tight Ends", func(w fyne.Window) fyne.CanvasObject {
			table := widget.NewTable(
				func() (int, int) {
					return len(nflplayers.GetPlayersByPosition("TE", players)), 5
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("label")
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					posPlayers := nflplayers.GetPlayersByPosition("TE", players)
					switch i.Col {
					case 0:
						o.(*widget.Label).SetText(posPlayers[i.Row].Rank)
					case 1:
						o.(*widget.Label).SetText(posPlayers[i.Row].Name)
					case 2:
						o.(*widget.Label).SetText(posPlayers[i.Row].Position)
					case 3:
						o.(*widget.Label).SetText(posPlayers[i.Row].Team)
					case 4:
						o.(*widget.Label).SetText(posPlayers[i.Row].ByeWeek)
					}
				},
			)

			table.SetColumnWidth(0, 50)
			table.SetColumnWidth(1, 250)
			table.SetColumnWidth(2, 100)
			table.SetColumnWidth(3, 100)
			return table
		}},
		"k": {"K", "Highest Ranking Kickers", func(w fyne.Window) fyne.CanvasObject {
			table := widget.NewTable(
				func() (int, int) {
					return len(nflplayers.GetPlayersByPosition("K", players)), 5
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("label")
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					posPlayers := nflplayers.GetPlayersByPosition("K", players)
					switch i.Col {
					case 0:
						o.(*widget.Label).SetText(posPlayers[i.Row].Rank)
					case 1:
						o.(*widget.Label).SetText(posPlayers[i.Row].Name)
					case 2:
						o.(*widget.Label).SetText(posPlayers[i.Row].Position)
					case 3:
						o.(*widget.Label).SetText(posPlayers[i.Row].Team)
					case 4:
						o.(*widget.Label).SetText(posPlayers[i.Row].ByeWeek)
					}
				},
			)

			table.SetColumnWidth(0, 50)
			table.SetColumnWidth(1, 250)
			table.SetColumnWidth(2, 100)
			table.SetColumnWidth(3, 100)
			return table
		}},
		"def": {"DEF", "Highest Ranking Defenses", func(w fyne.Window) fyne.CanvasObject {
			table := widget.NewTable(
				func() (int, int) {
					return len(nflplayers.GetPlayersByPosition("DST", players)), 5
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("label")
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					posPlayers := nflplayers.GetPlayersByPosition("DST", players)
					switch i.Col {
					case 0:
						o.(*widget.Label).SetText(posPlayers[i.Row].Rank)
					case 1:
						o.(*widget.Label).SetText(posPlayers[i.Row].Name)
					case 2:
						o.(*widget.Label).SetText(posPlayers[i.Row].Position)
					case 3:
						o.(*widget.Label).SetText(posPlayers[i.Row].Team)
					case 4:
						o.(*widget.Label).SetText(posPlayers[i.Row].ByeWeek)
					}
				},
			)

			table.SetColumnWidth(0, 50)
			table.SetColumnWidth(1, 250)
			table.SetColumnWidth(2, 100)
			table.SetColumnWidth(3, 100)
			return table
		}},
	}

	content := container.NewMax()
	title := widget.NewLabel("Nick's Player Rankings")
	intro := widget.NewLabel("")
	sideMenuCB := func(s SideMenu) {
		title.SetText(s.Title)
		intro.SetText(s.Description)

		content.Objects = []fyne.CanvasObject{s.View(myWindow)}
		content.Refresh()
	}

	tutorial := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)
	split := container.NewHSplit(makeNav(sideMenuCB, true, menu), tutorial)
	split.Offset = 0.2
	myWindow.SetContent(split)

	myWindow.Resize(fyne.NewSize(840, 460))
	myWindow.ShowAndRun()
}

func makeNav(sideMenuCB func(menu SideMenu), loadPrevious bool, menu map[string]SideMenu) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return Index[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := Index[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := menu[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
			obj.(*widget.Label).TextStyle = fyne.TextStyle{}

		},
		OnSelected: func(uid string) {
			if t, ok := menu[uid]; ok {
				a.Preferences().SetString("currentTutorial", uid)
				sideMenuCB(t)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback("currentTutorial", "welcome")
		tree.Select(currentPref)
	}

	return container.NewBorder(nil, nil, nil, nil, tree)
}
