package main

import (
	"database/sql"
	"net/http"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/juaniferro/fake-twitter/internal/handlers"
	"github.com/juaniferro/fake-twitter/internal/repositories"
	"github.com/juaniferro/fake-twitter/internal/services"
	"github.com/juaniferro/fake-twitter/internal/usecases"
	_ "github.com/juaniferro/fake-twitter/pkg"
)

func main()  {

	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/fake_tw_database?parseTime=true")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
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

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}