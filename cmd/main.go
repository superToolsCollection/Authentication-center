package main

import (
    "fmt"
	"log"
	"authentication-center/internal/app/engine"
	"authentication-center/internal/app/model"
    "authentication-center/internal/app/service"
    "github.com/douyu/jupiter"
)

func main() {
	eng := engine.NewEngine()
	eng.RegisterHooks(jupiter.StageAfterStop, func() error {
        fmt.Println("exit jupiter app ...")
        return nil
      })

    model.Init()
    service.Init()
    if err := eng.Run(); err != nil {
    	log.Fatal(err)
    }
}

