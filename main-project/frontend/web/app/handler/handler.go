package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/dhanusaputra/k8s-exercises/web/app/util"
)

// View ...
func View(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusBadRequest)
		return
	}

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
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo := r.FormValue("todo")

	query := fmt.Sprintf("{\"query\":\"mutation createTodo ($todo:String!,$done:Boolean!) {\\n  createTodo(input:{text:$todo,done:$done}) {\\n    text\\n    done\\n  }\\n}\",\"variables\":{\"todo\":\"%s\",\"done\":false}}", todo)

	_, statusCode, err := util.ReqBackend(query)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// UpdateTodo ...
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "not yet implemented", http.StatusNotImplemented)
		return
	}
}
