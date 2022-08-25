package memecontrollers

import (
	"context"
	"fmt"
	"meme/models"

	"github.com/kataras/iris/v12"

	_ "log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitializeEndpoints(app *iris.Application, db *sqlx.DB){


	//get all memes

	app.Get("/memes", func(ctx iris.Context) {
		memes := []models.Meme{}
		err := db.Select(&memes, "SELECT * FROM notes")

		if err != nil{
			fmt.Println(err)
			return
		}
		ctx.JSON(memes)
	} )
}