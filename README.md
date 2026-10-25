# Heran-Music

Web-based music catalog, playlist management, and listening queue.

## Stack

- Frontend: React, TypeScript, Vite
- Backend: Go
- Database: MariaDB
- Deployment: Docker and Docker Swarm

## Development

Copy `.env.example` to `.env` and set strong local database passwords.

Install frontend dependencies:

```bash
cd frontend
npm install
```

Run the frontend:

```bash
npm run dev
```

Run the API from the backend directory:

```bash
cd backend
go run ./cmd/api
```

Health check: `http://localhost:8080/health`

## Notes

The initial API currently exposes a health endpoint only. CRUD endpoints,
database connection, and frontend-to-backend integration still need to be
implemented before the application is feature-complete.

The Swarm stack requires Docker Hub image names to be replaced and the
external database secrets to be created before deployment.
