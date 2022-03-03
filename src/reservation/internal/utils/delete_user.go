// package utils
package utils
// import (
//     "src/users/internal/infrastructure/user"
//     "src/users/domain"
//     "context"
// )

// type DeleteUserCmd func(ctx context.Context, userID string) error

// func DeleteUser(storage user.Storage) DeleteUserCmd {
//     return func(ctx context.Context, userID string) error {
//         sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)

//         if !ok {
//             return domain.ErrCannotRetreiveSession
//         }

//         if userID == "" {
//             userID = sessionInfo.UserID
//         }

//         if sessionInfo.UserID != userID && sessionInfo.Access < 1 {
//             return domain.ErrOperationNotPermitted
//         }
//         return storage.Delete(ctx, userID)
//     }
// }