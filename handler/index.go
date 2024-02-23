package handler

import (
	"context"
	"fmt"
	"net/http"
	"portfolio/view/index"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	component := index.ShowIndex()
	err := component.Render(ctx, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	//ctx := context.Background()
	fmt.Println("Post handler")
	w.Write([]byte("Post handler"))
}
