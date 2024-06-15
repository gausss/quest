package main

import "os"

func main() {
	a := App{}
	a.Init(os.Getenv("APP_DB_NAME"), os.Getenv("APP_DB_PASSWORD"))
	a.Run(":8080")
}
