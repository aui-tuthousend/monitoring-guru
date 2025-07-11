# Monitoring Guru SMKN 02 SBY

Requirements :
- [PostgreSQL] CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
- [Go] 1.20 up

RESTful API :
- [Go](https://golang.org/)
- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://gorm.io/)
- JWT Authentication
- Swagger Docs (via swaggo)
- Dependency Injection
- WebSocket

---

## 🔧 How to Clone

### 1. Clone
```bash
git clone https://gitlab.com/aui-tuthousend/monitoring-guru.git
cd monitoring-guru
```

### 2. Create .env file
```bash
DATABASE_URL='host=yourhost user=userName password=pass dbname=dbName port=5432 sslmode=require'
JWT_SECRET='your_jwt_secret'
ENV='development'
```

### 3. Install Dependency
```bash
go mod tidy
go install github.com/air-verse/air@latest
go install github.com/swaggo/swag/cmd/swag@latest
```

### 4. Generate Swagger
```bash
swag init
air init
```

### 5. Run
```bash
air
```