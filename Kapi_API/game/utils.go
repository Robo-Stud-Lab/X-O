package game


func RemoveFromSliceBySelector(games []Game, GameID int) []Game {
	for i := len(games) - 1; i >= 0; i-- {
		if games[i].id == GameID {
		 return append(games[:i], games[i+1:]...)
		}
	}
	return games
}

func updateBoard(games []Game, game Game) []Game {
	for i := len(games) - 1; i >= 0; i-- {
		if games[i].id == game.id {
		 NewGames := append(games[:i], game)
		 NewGames = append(NewGames, games[i+1:]... )
		
		 return NewGames
		}
	}
	return games
}