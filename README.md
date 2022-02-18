### Microservices API

## Installation

- ``` cd src```
- ``` docker-compose up --build ```

## Routes

# Authentication
- POST   /login
- GET    /logout
- GET    /session

# Users
- POST   /new-account
- GET    /auth/access/{user_id}
- POST   /auth/login
- GET    /users
- GET    /users/{user_id}
- PUT    /users (?user_id=1)
- PUT    /users/balance (?user_id=1)
- DELETE /users (?user_id=1)

# Ads

- GET /ad/
- POST /ad/
- GET /ad/{ad_id}
- GET /ad/{ad_id}
- GET /ad/{ad_id}
- GET /ad/{ad_id}

# Balance

- GET /balance
- GET /balance/user/{user_id}
- GET /balance/user/me
- PUT /balance
- PUT /balance/user/{user_id}
- PUT /balance/user/me
- POST /balance

# Transactions

- GET /transactions
- POST /transactions
- GET /transactions/{transaction_id}
- DEL /transactions/{transaction_id}
- PUT /transactions/{transaction_id}
- PUT /transactions/{transaction_id}/accept
- PUT /transactions/{transaction_id}/refuse
- PUT /transactions/{transaction_id}/cancel
