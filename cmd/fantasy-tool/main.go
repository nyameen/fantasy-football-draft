package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	"github.com/nyameen/fantasy-football-draft/internal/fantasypro"
	"github.com/nyameen/fantasy-football-draft/internal/nflplayers"
)

func main() {

	data, err := fantasypro.GetFantasyProCSV()
	if err != nil {
		fmt.Println("Error, could not get CSV: ", err)
		return
	}

	players := nflplayers.GetNFLPlayersFromCSV(data[3:])

	myApp := app.New()
	myWindow := myApp.NewWindow("Fantasy Football Draft Tool")
	myWindow.Resize(fyne.NewSize(600, 400))

	table := widget.NewTable(
		func() (int, int) {
			return len(players), 4
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
			}
		},
	)

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 100)

	myWindow.SetContent(table)
	myWindow.ShowAndRun()
}
