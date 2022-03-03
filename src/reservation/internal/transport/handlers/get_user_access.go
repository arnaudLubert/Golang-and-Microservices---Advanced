package handlers

// import (
//     "src/users/internal/utils"
//     "src/users/domain"
//     "github.com/gorilla/mux"
//     "encoding/json"
//     "net/http"
// )

// type GetUserAccessResponse struct {
//     Access      int8          `json:"access"`
// }

// func GetUserAccessHandler(cmd utils.GetUserAccessCmd) http.HandlerFunc {
//     return func(rw http.ResponseWriter, req *http.Request) {
//         access, err := cmd(req.Context(), mux.Vars(req)["user_id"])

//         if err != nil {
//             switch err {
//             case domain.ErrUserNotFound: http.Error(rw, err.Error(), http.StatusNotFound)
//             default: http.Error(rw, err.Error(), http.StatusInternalServerError)
//             }
//         } else {
//             rw.Header().Set("Content-type", "application/json")
//             json.NewEncoder(rw).Encode(GetUserAccessResponse{ Access: access })
//         }
//     }
// }