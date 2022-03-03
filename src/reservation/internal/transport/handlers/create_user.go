package handlers

// import (
//     "src/users/internal/utils"
//     "src/users/domain"
//     "encoding/json"
//     "net/http"
// )

// type CreateUserRequest struct {
//     Pseudo       string         `json:"pseudo"`
//     Email       string          `json:"email"`
//     Password    string          `json:"password"`
//     Firstname   string          `json:"firstname"`
//     Lastname    string          `json:"lastname"`
//     Phone       string          `json:"phone"`
//     Access      int8            `json:"access"`
//     Address     domain.Address  `json:"address"`
// }

// type CreateUserResponse struct {
//     UserID      string          `json:"user_id"`
// }

// func CreateUserHandler(cmd utils.CreateUserCmd) http.HandlerFunc {
//     return func(rw http.ResponseWriter, req *http.Request) {
//         var createUserReq CreateUserRequest

//         decoder := json.NewDecoder(req.Body)
//         err := decoder.Decode(&createUserReq)

//         if err != nil {
//             http.Error(rw, err.Error(), http.StatusBadRequest)
//             return
//         }

//         if createUserReq.Pseudo == "" || createUserReq.Email == "" || createUserReq.Password == "" {
//             http.Error(rw, "missing pseudo, email nor password", http.StatusBadRequest)
//             return
//         }

//         user := domain.User{
//             Pseudo: createUserReq.Pseudo,
//             Email: createUserReq.Email,
//             Password: createUserReq.Password,
//             Firstname: createUserReq.Firstname,
//             Lastname: createUserReq.Lastname,
//             Phone: createUserReq.Phone,
//             Address: createUserReq.Address,
//             Access: createUserReq.Access,
//         }
//         userID, err := cmd(req.Context(), user)

//         if err != nil {
//             switch err {
//             case domain.ErrUserAlreadyExists:
//                 http.Error(rw, err.Error(), http.StatusConflict)
//             case domain.ErrUnsecuredPassword:
//                 http.Error(rw, err.Error(), http.StatusBadRequest)
//             case domain.ErrAccessForbidden, domain.ErrOperationNotPermitted:
//                 http.Error(rw, err.Error(), http.StatusForbidden)
//             default: http.Error(rw, err.Error(), http.StatusInternalServerError)
//             }
//         } else {
//             rw.WriteHeader(http.StatusCreated)
//             rw.Header().Set("Content-type", "application/json")
//             json.NewEncoder(rw).Encode(CreateUserResponse{UserID: userID})
//         }
//     }
// }
