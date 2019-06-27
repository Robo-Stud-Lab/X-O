package main

import (
	"fmt"
	"runtime"
	"time"
	"github.com/Betra/Kapi_API/test"
	"github.com/Betra/Kapi_API/game"
	"github.com/brainfucker/zero"
)

func main() {
	httpServer := zero.HTTP{}
	rest := httpServer.Rest("/api")

	rest.POST("/ping", test.Pong)

	rest.GET("/game/start", game.Start)
	rest.GET("/game/all", game.GetActive)
	rest.DELETE("/game/:game_id", game.Finish)
	rest.PATCH("/game/:game_id", game.EditBoard)

	go func() {
		for {
			fmt.Println("------------\ngoroutines num:", runtime.NumGoroutine())

			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
			fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
			fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
			fmt.Printf("\tNumGC = %v\n", m.NumGC)

			time.Sleep(time.Minute)
		}
	}()

	httpServer.Serve("9000")

}
