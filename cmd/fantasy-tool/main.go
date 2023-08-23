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
		"all": {"All", "Highest ranking players overall", players.CreateTableCallbackByPosition(nflplayers.ALL)},
		"qb":  {"QB", "Highest ranking Quarter Backs", players.CreateTableCallbackByPosition(nflplayers.QuarterBack)},
		"wr":  {"WR", "Highest ranking Wide Receivers", players.CreateTableCallbackByPosition(nflplayers.WideReceiver)},
		"rb":  {"RB", "Highest ranking Running Backs", players.CreateTableCallbackByPosition(nflplayers.RunningBack)},
		"te":  {"TE", "Highest ranking Tight Ends", players.CreateTableCallbackByPosition(nflplayers.TightEnd)},
		"k":   {"K", "Highest ranking Kickers", players.CreateTableCallbackByPosition(nflplayers.Kicker)},
		"def": {"DEF", "Highest ranking Defenses", players.CreateTableCallbackByPosition(nflplayers.Defense)},
	}

	// Set the main page to all players
	content := container.NewMax()
	title := widget.NewLabel(menu["all"].Title)
	description := widget.NewLabel(menu["all"].Description)
	content.Objects = []fyne.CanvasObject{menu["all"].View(myWindow)}
	content.Refresh()

	// Callback function to refresh the right side content
	sideMenuCB := func(s SideMenu) {
		title.SetText(s.Title)
		description.SetText(s.Description)

		content.Objects = []fyne.CanvasObject{s.View(myWindow)}
		content.Refresh()
	}

	border := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), description), nil, nil, nil, content)
	split := container.NewHSplit(makeNav(sideMenuCB, true, menu), border)
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
			return widget.NewLabel("")
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
