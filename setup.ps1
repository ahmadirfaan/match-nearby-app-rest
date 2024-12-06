Write-Output "Installing libraries..."
# Web framework
go get -u github.com/gin-gonic/gin

# Database ORM
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

# Authentication (JWT)
go get -u github.com/golang-jwt/jwt/v5

# Rate limiting (Redis)
go get -u github.com/go-redis/redis/v8
go get -u github.com/go-redis/redis_rate/v9

# Input validation
go get -u github.com/go-playground/validator/v10

# Logging
go get -u github.com/sirupsen/logrus
# OR
go get -u go.uber.org/zap

# Testing
go get -u github.com/stretchr/testify

# API documentation (Swagger)
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/swag


# Environment variables
go get -u github.com/joho/godotenv

Write-Output "Cleaning up dependencies..."
go mod tidy

Write-Output "Setup complete!"
