package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/juaniferro/fake-twitter/internal/models"
	"github.com/juaniferro/fake-twitter/internal/usecases"
)

type PostTweetUsecaseInterface interface{
	TweetPoster(user int, tweet string) error
}

type PostTweetHandler struct {
	postTweetUsecase PostTweetUsecaseInterface
}

func NewPostTweetHandler(postTweetUsecase usecases.PostTweetUsecase) *PostTweetHandler{
	return &PostTweetHandler{postTweetUsecase: postTweetUsecase}
}

func (pth PostTweetHandler)HandlePostTweet(w http.ResponseWriter, r *http.Request) {
	user_id_str := r.Header.Get("user_id")
	user_id, err := strconv.Atoi(user_id_str) 
    if err != nil {
        http.Error(w, "Invalid user_id", http.StatusBadRequest)
        return
    }

	var tweet models.Tweet
	err = json.NewDecoder(r.Body).Decode(&tweet)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	err = pth.postTweetUsecase.TweetPoster(user_id, tweet.Content)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	w.WriteHeader(http.StatusCreated)
}