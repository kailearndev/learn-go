# Kai Shop BE (Go)


## Yêu cầu
- Go 1.21+
- Git
- (Tùy chọn) Docker, Make

## Cài đặt
- Cài dependency:
    - go mod tidy
- Tạo biến môi trường (ví dụ .env):
    - PORT=8080
    - DATABASE_URL=postgres://user:pass@localhost:5432/kai_shop?sslmode=disable
    - JWT_SECRET=change-me

## Chạy phát triển
- go run ./cmd/api

## Test
- go test ./...

## Build
- go build -o bin/kai-shop-be ./cmd/api

## Docker (tùy chọn)
- docker build -t kai-shop-be .
- docker run -p 8080:8080 --env-file .env kai-shop-be

## Cấu trúc (gợi ý)
- cmd/api: entrypoint (main.go)
- internal/...: business logic
- pkg/...: packages dùng chung
- configs, migrations, scripts, docs: tùy dự án

## Chất lượng mã
- go fmt ./...
- go vet ./...
- (Tùy chọn) golangci-lint run

## API Docs (tùy chọn)
- swag init -g cmd/api/main.go
- Truy cập /swagger/index.html (nếu cấu hình)