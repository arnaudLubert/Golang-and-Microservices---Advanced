// package handlers
package handlers
// import (
// 	"encoding/json"
// 	"net/http"
// 	"src/users/domain"
// 	"src/users/internal/utils"

// 	"github.com/gorilla/mux"
// )

// func GetUserHandler(cmd utils.GetUserCmd) http.HandlerFunc {
// 	return func(rw http.ResponseWriter, req *http.Request) {
// 		user, err := cmd(req.Context(), mux.Vars(req)["user_id"])

// 		if err != nil {
// 			switch err {
// 			case domain.ErrUserNotFound:
// 				http.Error(rw, err.Error(), http.StatusNotFound)
// 			default:
// 				http.Error(rw, err.Error(), http.StatusInternalServerError)
// 			}
// 		} else {
// 			rw.Header().Set("Content-type", "application/json")
// 			json.NewEncoder(rw).Encode(user)
// 		}
// 	}
// }
