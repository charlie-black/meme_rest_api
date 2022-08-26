package main

import (
	_ "fmt"
	authcontroller "meme/controllers"
	memecontrollers "meme/controllers/meme_controllers"
	"meme/models"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	_ "github.com/lib/pq"

	"log"

	"github.com/jmoiron/sqlx"
)

func main() {

	var SECRET = []byte("@$%^&@*()^^---^%*")
	signer := jwt.NewSigner(jwt.HS256, SECRET, 1*time.Minute)
	verifier := jwt.NewVerifier(jwt.HS256, SECRET)

	verifyMiddleWare := verifier.Verify(func() interface{} {
		return new(models.UserClaims)
	})

	//connect to the database
	db, err := sqlx.Connect("postgres", "user=piccasso dbname=memeDB sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}
	println("Connected to Database", db)

	app := iris.New()
	app.Use(iris.Compression)

	authcontroller.InitializeAuthEndpoints(signer, app, db)
	memecontrollers.InitializeEndpoints(app, db, verifyMiddleWare)

	app.Listen(":3500")

}
