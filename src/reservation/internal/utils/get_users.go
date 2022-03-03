// package utils
package utils
// import (
//     "src/users/internal/infrastructure/user"
//     "src/users/domain"
//     "context"
// )

// type GetUsersCmd func(ctx context.Context) (*[]domain.User, error)

// func GetUsers(storage user.Storage) GetUsersCmd {
//     return func(ctx context.Context) (*[]domain.User, error) {
//         users, err := storage.GetAll(ctx)

//         if err != nil {
//             return nil, err
//         }
//         return users, nil
//     }
// }