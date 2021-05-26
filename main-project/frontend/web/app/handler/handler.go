package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/dhanusaputra/k8s-exercises/web/app/util"
	"github.com/go-chi/chi/v5"
)

// View ...
func View(w http.ResponseWriter, r *http.Request) {
	if err := util.DownloadImageToVolume("https://picsum.photos/1200", "./static/test"); err != nil {
		http.Error(w, fmt.Sprintf("failed when download image, err: %v", err), http.StatusInternalServerError)
		return
	}

	_, err := util.ReadImageFromVolume("./static/test")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed when read image, err: %v", err), http.StatusInternalServerError)
		return
	}

	query := fmt.Sprintf("{\"query\":\"query listTodo {\\n    todos {\\n      id\\n      text\\n     done\\n     }\\n}\",\"variables\":{}}")

	resp, statusCode, err := util.ReqBackend(query)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	t := template.Must(template.New("view.html").ParseFiles("view.html"))
	if err := t.Execute(w, resp.Data); err != nil {
		http.Error(w, fmt.Sprintf("failed when execute template, err: %v", err), http.StatusInternalServerError)
		return
	}
}

// CreateTodo ...
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	text := r.FormValue("text")

	query := fmt.Sprintf("{\"query\":\"mutation createTodo ($text:String!,$done:Boolean!) {\\n  createTodo(input:{text:$text,done:$done}) {\\n    text\\n    done\\n  }\\n}\",\"variables\":{\"text\":\"%s\",\"done\":false}}", text)

	_, statusCode, err := util.ReqBackend(query)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ToggleTodo ...
func ToggleTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	v := r.FormValue(id)

	done := "true"
	if len(v) > 0 {
		done = "false"
	}

	query := fmt.Sprintf("{\"query\":\"mutation updateTodo ($id:ID!,$done:Boolean!) {\\n  updateTodo(id:$id, modifications:{done:$done}) {\\n    id\\n    text\\n    done\\n  }\\n}\",\"variables\":{\"id\":%s,\"done\":%s}}", id, done)

	_, statusCode, err := util.ReqBackend(query)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
