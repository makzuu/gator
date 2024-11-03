package main

import (
	"fmt"
	"github.com/makzuu/gator/internal/config"
	"log"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}
	conf.SetUser("Makz")
	conf, err = config.Read()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(conf)
}
