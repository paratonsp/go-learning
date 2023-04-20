# learning-go
Go Rest API:
- Simple CRUD to MySQL
- JWT using Cookie
- Middleware using MUX

Installing:
- docker build -t learning-go .
- docker run -d -p 8080:8080 --name=learning-go learning-go:latest

InstallinV2:
- docker compose -f nginx-compose.yaml up -d
- docker compose -f app-compose.yaml up -d