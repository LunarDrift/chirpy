# Chirpy API

A Twitter-like REST API built with Go. Made as part of the [Boot.dev](https://boot.dev) back-end development path.

## Features

- User registration and authentication with hashed passwords
- JWT-based access tokens and refresh tokens
- Create, retrieve, and delete chirps (posts limited to 140 characters)
- Automatic 'profanity' filtering
- Chirpy Red membership upgrades via webhook

## Tech Stack
- **Language:** Go
- **Database:** PostgreSQL
- **Migrations:** Goose
- **Query generation**: sqlc

## Prerequisites
- Go 1.22+
- PostgreSQL 15+
- Goose (`go install github.com/pressly/goose/v3/cmd/goose@latest`)

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/LunarDrift/chirpy.git
cd chirpy
```

2. Create a `.env` file and add your database connection:
```bash
DB_URL=postgres://username:password@localhost:5432/chirpy
JWT_SECRET=your_jwt_secret
POLKA_KEY=your_polka_api_key
```

3. Run database migrations:
```bash
goose -dir sql/schema postgres "$DB_URL" up
```

4. Start the server:
```bash
go build -o out && ./out
```

## API Endpoints

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | `/api/users` | None | Create a user |
| PUT | `/api/users` | Acesss token | Update email and password |
| POST | `/api/login` | None | Login, returns access and refresh tokens |
| POST | `/api/refresh` | Refresh token | Get a new access token |
| POST | `/api/revoke` | Refresh token | Revoke a refresh token |
| POST | `/api/chirps` | Acess token | Create a chirp |
| GET | `/api/chirps` | None | Get all chirps (filter by `?author_id=`, sort by `?sort=asc/desc`) |
| GET | `/api/chirps/{chirpID}` | None | Get a chirp by ID |
| DELETE | `/api/chirps/{chirpID}` | Access token | Delete a chirp |
| POST | `/api/polka/webhooks` | API key | Webhook for Chirpy Red upgrades |
