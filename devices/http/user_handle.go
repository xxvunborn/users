package http

import (
	"net/http"

	"github.com/unclebob/user"
	"github.com/unclebob/user/entities"
)

func handlers() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", UserGet)
}

// UserGet return the get of the controller
func UserGet(w http.ResponseWriter, r *http.Request) {

	var u *entities.UserQuery
	if err := json.Unmarshal(body, &u); err != nil {
		http.Error(w, "error", 422)
		return
	}
	//c, err := user.Controller.Get(q)
}
