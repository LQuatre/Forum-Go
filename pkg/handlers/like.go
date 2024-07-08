package handlers

import (
	"encoding/json"
	"net/http"

	"jilt.com/m/pkg/models"
)

func LikeHandler(w http.ResponseWriter, r *http.Request) {
	// si le post n'est pas  liké alors on like et si il est liké alors on unlike
	// on recupere le post_id et le user_id
	// on verifie si le post est liké
	// si il est liké alors on unlike
	// si il n'est pas liké alors on like
	// on retourne le nombre de like

	// on recupere le post_id et le user_id
	postID := r.FormValue("post_id")
	userID := r.FormValue("user_id")

	// on verifie si le post est liké
	like, err := models.GetLikeByPostID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if like.UUID != "" {
		like.Delete()
	} else {
		like = &models.Like{
			UserID: userID,
			PostID: postID,
		}
		like.Create()
	}

	// on retourne le nombre de like
	likes, err := models.GetLikeByPostID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(likes)
}

func DislikeHandler(w http.ResponseWriter, r *http.Request) {
	// pareille mais avec un dislike
	postID := r.FormValue("post_id")
	userID := r.FormValue("user_id")

	dislike, err := models.GetDislikeByPostID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if dislike.UUID != "" {
		dislike.Delete()
	} else {
		dislike = &models.Dislike{
			UserID: userID,
			PostID: postID,
		}
		dislike.Create()
	}

	dislikes, err := models.GetDislikesByPostID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dislikes)
}
