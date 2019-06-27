package game

import (
	"fmt" 
	"github.com/brainfucker/zero"
	"github.com/Betra/Kapi_API/ml"
	"math/rand"
	"time"
	"strconv"
	"encoding/json"
)

var games []Game

// Game constructor
func Init() {
	fmt.Println("Game server started ")
}

// Game struct
type Game struct {
	id     int
	board  [9]int // 0-2
	winner string
}

// Starts new Game, returns game Id
func Start(srv *zero.Server) {
	rand.Seed(time.Now().UTC().UnixNano())

	Gameid := rand.Int()

	board := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	game := Game{Gameid, board, ""}
	games = append(games, game)

	srv.Resp(zero.H{"game_id": game.id})

	fmt.Println("\n---------------------\nNew game started: gameId:", game.id, "\n board: ", game.board,"\n-------------------------------")
	fmt.Println("Amount of running games: ", len(games))
}


func Finish(srv *zero.Server) {
	GameIDStr := srv.GetPathParam("game_id")
	if GameIDStr == "" {
		srv.Err("param_int", "game_id wasn't passed")
		return
	}

	GameID, ErrorOnStrToInt := strconv.Atoi(GameIDStr)

	if (ErrorOnStrToInt != nil) {
		srv.Err("param_int", ErrorOnStrToInt)
		return
	}

	_, ErrorOnFindGame := getGameByID(srv)

	if (ErrorOnFindGame) {
		return
	}

	games = RemoveFromSliceBySelector(games, GameID)

	fmt.Println("Game ", &GameID, " was deleted. Active games: ", len(games))

	srv.RespOk()
}


func GetActive(srv *zero.Server) {

	resp := []zero.H{}
	for _, game := range games {
		resp = append(resp, zero.H{"game_id": game.id, "board": game.board})
	}

	srv.Resp(zero.H{"items": resp})
	fmt.Println("Active games: ", resp)

}

type BoardInput struct {
	Board [9]int `json:"board"`
}

func EditBoard(srv *zero.Server) {
	EditedGame, err := getGameByID(srv)
	if err {
		return
	}
	input := BoardInput{}

	erre := json.Unmarshal(srv.GetBody(), &input)

	if erre != nil {
		srv.Err("err", "can't parse body")
		return
	}

	EditedGame.board = input.Board

	games = updateBoard(games, *EditedGame)

	fmt.Println("Game updated: ", EditedGame.id, EditedGame.board)

	fmt.Println(EditedGame)
	
	AiBoard, winner := ml.FindSolution(EditedGame.board)

	if (winner == "") {
		fmt.Println(AiBoard)
	} else {
		srv.Resp(zero.H{"result": winner})
		Finish(srv)
	}

}

func getGameByID(srv *zero.Server) (*Game, bool) {
	GameIDStr := srv.GetPathParam("game_id")

	if GameIDStr == "" {
		srv.Err("param_int", "game_id wasn't passed")
		return nil, true
	}

	GameID, ErrorOnStrToInt := strconv.Atoi(GameIDStr)

	if (ErrorOnStrToInt != nil) {
		srv.Err("param_int", ErrorOnStrToInt)
			return nil, true
	}

	var game *Game

	for i := len(games) - 1; i >= 0; i-- {
		if games[i].id == GameID {
			game = &games[i]
		}
	}


	if (game == nil) {
		srv.Err("Couldn't find game","")
		return nil, true
	}

	return game, false
}