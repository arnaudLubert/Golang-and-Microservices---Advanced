// package utils
package utils
// import (
//     "src/users/internal/infrastructure/user"
//     "context"
// )

// type GetUserAccessCmd func(ctx context.Context, userID string) (int8, error)

// func GetUserAccess(storage user.Storage) GetUserAccessCmd {
//     return func(ctx context.Context, userID string) (int8, error) {
//         access, err := storage.GetAccess(ctx, userID)

//         if err != nil {
//             return -1, err
//         }
//         return access, nil
//     }
// }