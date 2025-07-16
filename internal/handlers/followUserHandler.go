package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/juaniferro/fake-twitter/internal/usecases"
)

type FollowUserUsecaseInterface interface{
	UserFollower(user, followedUser int) error
}

type FollowUserHandler struct {
	followUserUsecase FollowUserUsecaseInterface
}

func NewFollowUserHandler(followUserUsecase usecases.FollowUserUsecase) *FollowUserHandler{
	return &FollowUserHandler{followUserUsecase: followUserUsecase}
}

func (fuh FollowUserHandler)HandleFollowUser(w http.ResponseWriter, r *http.Request) {
	user_id_str := r.Header.Get("user_id")
	user_id, err := strconv.Atoi(user_id_str) 
	if err != nil {
        http.Error(w, "Invalid user_id", http.StatusBadRequest)
        return
    }
	
	vars := mux.Vars(r)
	followed_user_str := vars["followed_user_id"]
    followedUserID, err := strconv.Atoi(followed_user_str)
    if err != nil {
        http.Error(w, "Invalid followed user id", http.StatusBadRequest)
        return
    }

	err = fuh.followUserUsecase.UserFollower(user_id, followedUserID)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	w.WriteHeader(http.StatusCreated)
}