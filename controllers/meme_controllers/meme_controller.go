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

	// create meme

	app.Post("/create_memes", func(ctx iris.Context){
		var user_params models.CreateMeme
		err := ctx.ReadJSON(&user_params)

		if err != nil{
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		_, err = db.NamedExec(`INSERT INTO notes (creator, url)
        VALUES (:creator, :url)`, user_params)

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return

		}

		ctx.JSON(iris.Map{"message": "Meme Added Successfully"})

	
		
	})
}