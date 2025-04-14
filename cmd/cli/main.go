package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/jaychillin2607/go-http-server"
)

const dbFileName = "game.db.json"

func main() {
	fileSystemPlayerStore, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println(`Type "{Name} wins" to record a win`)

	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), fileSystemPlayerStore)
	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
