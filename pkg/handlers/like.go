package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"jilt.com/m/pkg/models"
)

func LikeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LikeHandler") // /like?post_id=1&user_id=1 get from url
	postUUID := r.FormValue("post_uuid")
	userUUID := r.FormValue("user_uuid")

	like, _ := models.GetLikeByPostAndUserUUID(postUUID, userUUID)
	
	if like.UUID != "" {
		fmt.Println("Delete like")
		like.Delete()
	} else {
		fmt.Println("Create like")
		like = &models.Like{
			UserUUID: userUUID,
			PostUUID: postUUID,
		}
		like.Create()
	}

	likes, err := models.GetLikesByPostUUID(postUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(likes)
	fmt.Println("LikeHandler end")
}

func DislikeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DislikeHandler")
	postUUID := r.FormValue("post_uuid")
	userUUID := r.FormValue("user_uuid")

	dislike, _ := models.GetDislikeByPostAndUserUUID(postUUID, userUUID)

	if dislike.UUID != "" {
		fmt.Println("Delete dislike")
		dislike.Delete()
	} else {
		fmt.Println("Create dislike")
		dislike := &models.Dislike{
			UserUUID: userUUID,
			PostUUID: postUUID,
		}
		dislike.Create()
	}

	json.NewEncoder(w).Encode(dislike)
	fmt.Println("DislikeHandler end")
}
