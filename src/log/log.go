package log

import (
	"log"
	"os"
)

func LogAction(action string) {
	f, err := os.OpenFile("actions.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		log.Println(err)
		log.Println("Couldn't append to actions.log")
	}

	logger := log.New(f, "action", log.LstdFlags)
	logger.Println(action)
}
