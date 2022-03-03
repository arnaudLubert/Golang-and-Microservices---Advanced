package handlers

// import (
//     "src/users/internal/utils"
//     "src/users/domain"
//     "net/http"
// )

// func DeleteUserHandler(cmd utils.DeleteUserCmd) http.HandlerFunc {
//     return func(rw http.ResponseWriter, req *http.Request) {
//         err := cmd(req.Context(), req.URL.Query().Get("user_id"))

//         if err != nil {
//             switch err {
//             case domain.ErrUserNotFound: http.Error(rw, err.Error(), http.StatusNotFound)
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