package main

import (
	"urils/application/initialize"
)

func main() {
	Router := initialize.InitRouter()
	Router.Run(":8000")

}
