// package utils
package utils
// import (
//     "src/users/internal/infrastructure/user"
//     "src/users/domain"
//     "context"
// )

// type GetUserCmd func(ctx context.Context, userID string) (*domain.User, error)

// func GetUser(storage user.Storage) GetUserCmd {
//     return func(ctx context.Context, userID string) (*domain.User, error) {
//         user, err := storage.Get(ctx, userID)

//         if err != nil {
//             return nil, err
//         }
//         return user, nil
//     }
// }