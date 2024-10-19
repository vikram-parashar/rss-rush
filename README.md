# RSS Rush

RSS Rush is a fast and simple RSS aggregator built with Golang. It enables users to manage and follow their favorite channels, fetch articles, and handle authentication using API keys. The project leverages the Gin-Gonic framework for REST APIs, Goose for database migrations, SQLC for query generation, and PostgreSQL as the database.

## Features
- User authentication via API keys.
- Channel management (add, delete, list).
- Follow and unfollow channels.
- Fetch articles from followed channels.

## Tech Stack
- **Golang**
- **Gin-Gonic**: REST API framework
- **Goose**: Database migrations
- **SQLC**: Go code generation for SQL queries
- **PostgreSQL**: Database

## Setup

### Prerequisites
- Golang installed (Go 1.20 or higher)
- PostgreSQL installed
- Goose installed for database migrations
- SQLC installed for query generation

### 1. Clone the repository
```bash
git clone https://github.com/yourusername/rss-rush.git
cd rss-rush
```
### 2. Update the .env file with your postgres db url
```
PORT=3001
DB_URL=postgres://user:password@localhost:5432/
```
### 3. Run migrations and seed.sql in /sql directiory
```bash
goose postgresql {dbURL} up
```
### 4. Generate queries
```bash
sqlc generate
```
### 5. start the app
```bash
make run
```
or 
```bash
go run .
```

## Usage

### Endpoints

#### User Management
- **Create User**:  
  `POST /user?name=...&email=...`
  
- **Get User (authenticated)**:  
  `GET /user`  
  *Header*: `Authorization: api_key`

#### Channel Management
- **Create Channel**:  
  `POST /channel?name=...&htmlUrl=...&xmlUrl=...`
  
- **Get Channels (paginated)**:  
  `GET /channels?limit=...&offset=...`
  
- **Delete Channel**:  
  `DELETE /channel/:channelId`

#### Follow Management
- **Follow a Channel**:  
  `POST /follow/:channelId`
  
- **Get Followed Channels**:  
  `GET /follows`
  
- **Unfollow a Channel**:  
  `DELETE /follow/:channelId`

#### Article Fetching
- **Get Articles (paginated)**:  
  `GET /articles?limit=...&offset=...`

### Authentication
All endpoints (except creating a user) require API key authentication. Provide the API key in the `Authorization` header:
```
Authorization: api_key
```

