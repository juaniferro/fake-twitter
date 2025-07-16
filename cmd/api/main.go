package main

import (
	"database/sql"
	"net/http"
	"log"
	"os"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/juaniferro/fake-twitter/internal/handlers"
	"github.com/juaniferro/fake-twitter/internal/repositories"
	"github.com/juaniferro/fake-twitter/internal/services"
	"github.com/juaniferro/fake-twitter/internal/usecases"
	_ "github.com/juaniferro/fake-twitter/pkg"
)

func main()  {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
	shouldParseTime := os.Getenv("SHOULD_PARSE_TIME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%s", dbUser, dbPass, dbHost, dbPort, dbName, shouldParseTime)
    db, err := sql.Open("mysql", dsn)
	 if err != nil {
        panic(err.Error())  
    }
	defer db.Close()

	//Repository
	fakeTwitterRepo := repositories.NewFakeTwitterRepo(db)

	//Services
	postTweetService := services.NewPostTweetService(*fakeTwitterRepo)
	followUserService := services.NewFollowUserService(*fakeTwitterRepo)
	getTimelineService := services.NewGetTimelineService(*fakeTwitterRepo)

	//UseCases
	postTweetUsecase := usecases.NewPostTweetUsecase(*postTweetService)
	followUserUsecase := usecases.NewFollowUsertUsecase(*followUserService)
	getTimelineUsecase := usecases.NewGetTimelineUsecase(*getTimelineService)

	//Handlers
	postTweetHandler := handlers.NewPostTweetHandler(*postTweetUsecase)
	followUserHandler := handlers.NewFollowUserHandler(*followUserUsecase)
	getTimelineHandler := handlers.NewGetTimelineHandler(*getTimelineUsecase)

	router := mux.NewRouter()

	router.HandleFunc("/tweet", postTweetHandler.HandlePostTweet).Methods("POST")
	router.HandleFunc("/follow/{followed_user_id}", followUserHandler.HandleFollowUser).Methods("POST")
	router.HandleFunc("/timeline", getTimelineHandler.HandleGetTimeline).Methods("GET")

	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}