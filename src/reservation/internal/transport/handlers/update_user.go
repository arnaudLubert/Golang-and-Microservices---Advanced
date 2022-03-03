// package handlers
package handlers
// import (
//     "src/users/internal/utils"
//     "src/users/domain"
//     "encoding/json"
//     "net/http"
// )

// type UpdateUserRequest struct {
//     Email       string          `json:"email"`
//     Firstname   string          `json:"firstname"`
//     Lastname    string          `json:"lastname"`
//     Phone       string          `json:"phone"`
//     Address     domain.Address  `json:"address"`
// }

// func UpdateUserHandler(cmd utils.UpdateUserCmd) http.HandlerFunc {
//     return func(rw http.ResponseWriter, req *http.Request) {
//         var updateUserReq UpdateUserRequest

//         decoder := json.NewDecoder(req.Body)
//         err := decoder.Decode(&updateUserReq)

//         if err != nil {
//             http.Error(rw, err.Error(), http.StatusBadRequest)
//             return
//         }

//         user := domain.User{
//             Email: updateUserReq.Email,
//             Firstname: updateUserReq.Firstname,
//             Lastname: updateUserReq.Lastname,
//             Phone: updateUserReq.Phone,
//             Address: updateUserReq.Address,
//         }

//         if err = cmd(req.Context(), req.URL.Query().Get("user_id"), user); err != nil {
//             switch err {
//             case domain.ErrUserNotFound: http.Error(rw, err.Error(), http.StatusNotFound)
//             case domain.ErrInvalidIBAN: http.Error(rw, err.Error(), http.StatusBadRequest)
//             case domain.ErrUserAlreadyExists: http.Error(rw, err.Error(), http.StatusConflict)
//             case domain.ErrAccessForbidden, domain.ErrOperationNotPermitted:
//                 http.Error(rw, err.Error(), http.StatusForbidden)
//             default: http.Error(rw, err.Error(), http.StatusInternalServerError)
//             }
//         } else {
//             rw.WriteHeader(http.StatusNoContent)
//         }
//     }
// }