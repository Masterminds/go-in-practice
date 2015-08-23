package safely

import (
	"log"
)

type GoDoer func()

func Go(todo GoDoer) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic in safely.Go: %s", err)
			}
		}()
		todo()
	}()
}
