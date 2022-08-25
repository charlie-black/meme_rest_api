package memecontrollers

import (
	"fmt"
	"meme/models"

	"github.com/kataras/iris/v12"

	_ "log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitializeEndpoints(app *iris.Application, db *sqlx.DB) {

	//get all memes

	app.Get("/memes", func(ctx iris.Context) {
		memes := []models.Meme{}
		err := db.Select(&memes, "SELECT * FROM funny_memes")

		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.JSON(memes)
	})

	// create meme

	app.Post("/create_memes", func(ctx iris.Context) {
		var user_params models.CreateMeme
		err := ctx.ReadJSON(&user_params)

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		_, err = db.NamedExec(`INSERT INTO funny_memes (creator, url)
        VALUES (:creator, :url)`, user_params)

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return

		}

		ctx.JSON(iris.Map{"message": "Meme Created Successfully"})

	})

	//delete a meme
	app.Delete("/delete_meme", func(ctx iris.Context) {
		var user_params models.DeleteMeme
		err := ctx.ReadJSON(&user_params)
		var count int
		err2 := db.QueryRow(fmt.Sprint("SELECT COUNT(*) from funny_memes where id =", user_params.ID)).Scan(&count)

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		if count == 0 {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": "no meme with given id"})
			return
		}
		if err2 != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		_, err = db.Exec("DELETE FROM funny_memes WHERE id =$1", user_params.ID)

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		ctx.JSON(iris.Map{"message": "Meme deleted"})

	})

	//update a meme

	app.Post("/update_meme", func(ctx iris.Context) {
		var user_params models.UpdateMeme
		err := ctx.ReadJSON(&user_params)
		var count int
		err2 := db.QueryRow(fmt.Sprint("SELECT COUNT(*) from funny_memes where id =", user_params.ID)).Scan(&count)

		if count == 0 {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": "no meme with given id"})
			return
		}

		if err2 != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		_, err = db.NamedExec("UPDATE funny_memes SET creator=:creator, url=:url WHERE id=:id", user_params)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}

		ctx.JSON(iris.Map{"message": "Meme updated"})

	})

	app.Get("/single_meme", func(ctx iris.Context) {
		var user_params models.GetMemeByID
		memes := []models.GetMemeByID{}
		err := db.Select(&memes, "SELECT id FROM funny_memes WHERE id=$1", user_params.ID)

		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.JSON(memes)

	})
}
