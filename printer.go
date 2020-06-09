package main

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func printGames(nowPlaying []nowPlaying) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "My Turn", "Opponent", "Last Move", "Board"})
	for _, game := range nowPlaying {
		for _, row := range getRowStrings(game.Fen, game.Color == "black") {
			t.AppendRow([]interface{}{game.GameID, game.IsMyTurn, game.Opponent.Username, game.LastMove, row}, rowConfigAutoMerge)
		}
		t.AppendSeparator()
	}
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true},
		{Number: 2, AutoMerge: true},
		{Number: 3, AutoMerge: true},
		{Number: 4, AutoMerge: true},
		{Number: 5, AutoMerge: true},
	})
	t.SetStyle(table.StyleBold)
	t.Render()
}

func printMoveMessage(move string, message string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Attempted Move", "Message"})
	t.AppendRow([]interface{}{move, message})
	t.SetStyle(table.StyleBold)
	t.Render()
}