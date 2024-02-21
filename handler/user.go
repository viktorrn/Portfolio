package handler

import (
	"context"
	"net/http"
	"portfolio/view/user"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	component := user.Show(r.URL.Path)
	err := component.Render(ctx, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
