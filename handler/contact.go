package handler

import (
	"context"
	"net/http"
	"portfolio/view/contact"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	component := contact.ShowContact()
	err := component.Render(ctx, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
