package main

import (
	memecontrollers "meme/controllers/meme_controllers"

	"github.com/kataras/iris/v12"
	_ "github.com/lib/pq"

	"log"

	"github.com/jmoiron/sqlx"
)


func main(){

	//connect to the database
	db, err := sqlx.Connect("postgres", "user=piccasso dbname=notebook sslmode=disable")

	if err != nil{
		log.Fatalln(err)
	}
	println("Connected to Database", db)

	app := iris.New()
	app.Use(iris.Compression)

	memecontrollers.InitializeEndpoints(app,db)

	app.Listen(":3500")


}