package app
import (
	."./src/router"
)
func main()  {
	println("1")
	router := InitRouter()

	router.Run(":8000")
}