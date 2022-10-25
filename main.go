package main

import (
	"assignment_3/controllers"
	"assignment_3/routers"
	"time"
)

func main() {

	for true {
		go controllers.GetStatus()
		time.Sleep(15 * time.Second)
	}
	routers.Start().Run("localhost:8001")

}
