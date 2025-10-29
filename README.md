# 🧱 README.md – Kai Shop Backend

> 🧰 Backend API viết bằng Go, kiến trúc DDD (Domain-Driven Design)  
> Framework: Gin • ORM: GORM • DB: PostgreSQL
>
> Mục tiêu: Hệ thống quản lý sản phẩm (PC, thiết bị, phụ kiện…) với phân quyền, upload ảnh, và JWT auth.

---

## 📂 Cấu trúc thư mục chuẩn

```
kai-shop-be/
├── cmd/
│   └── main.go                # entrypoint: gọi app.Run()
├── internal/
│   ├── app/                   # setup toàn bộ hệ thống (db, router, handler)
│   │   └── app.go
│   ├── domain/                # nơi chứa các "module" (Domain) tách biệt
│   │   ├── user/
│   │   ├── product/
│   │   ├── category/
│   │   └── upload/
│   ├── server/
│   │   └── router.go
│   └── pkg/                   # các gói tiện ích chung
│       ├── db/
│       ├── jwt/
│       ├── response/
│       └── validator/
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Cách khởi chạy dự án

```bash
git clone https://github.com/yourusername/kai-shop-be.git
cd kai-shop-be

cp .env.example .env

air
# hoặc
go run cmd/main.go
```

---

## 🧩 Kiến trúc tổng thể (DDD)

| Layer                | Vai trò                                                      |
| -------------------- | ------------------------------------------------------------ |
| Model (Entity)       | Đại diện dữ liệu (mapping DB với struct Go)                  |
| Repository           | Làm việc với database (CRUD, query logic)                    |
| Service (Use case)   | Xử lý nghiệp vụ, gọi repo, validate logic                    |
| Handler (Controller) | Giao tiếp với HTTP (Gin), parse JSON, trả response           |
| Router               | Đăng ký route, gom nhóm domain                               |
| App                  | Khởi tạo toàn hệ thống (DB, router, migration, server start) |

---

## ⚒️ Quy trình tạo module mới

1. **Tạo Model (Entity)**  
   → `internal/domain/<module>/model.go`

2. **Repository**  
   → CRUD, làm việc trực tiếp với DB

3. **Service (Business Logic)**  
   → Validate, xử lý logic nghiệp vụ

4. **Handler (HTTP)**  
   → Giao tiếp REST API

5. **Router**  
   → Gắn các route của module

6. **App.go**  
   → Khởi tạo repo, service, handler và truyền vào router

7. **Migration**  
   → Thêm vào AutoMigrate()

---

## 🧰 Chuẩn hóa JSON Response

```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "name": "Laptop"
  }
}
```

---

## 🔒 Middleware / JWT Auth

Middleware `JWTAuth()` dùng để parse và kiểm tra token.  
Khi token hợp lệ, nó set vào Gin Context:

```go
c.Set("userID", claims.UserID)
c.Set("email", claims.Email)
```

---

## 🧱 Best Practices

✅ Mỗi module có 4 file: `model.go`, `repository.go`, `service.go`, `handler.go`  
✅ Không gọi DB trực tiếp trong handler  
✅ Handler chỉ giao tiếp qua Service  
✅ Mỗi Handler có hàm `RegisterRoutes()`  
✅ Không return model nhạy cảm (vd: password)  
✅ Migrate theo thứ tự: User → Category → Product  
✅ Preload dữ liệu quan hệ chỉ khi cần thiết

---

## 🧩 Tạo module mới (ví dụ: Order)

1️⃣ Tạo thư mục `internal/domain/order/`  
2️⃣ Copy template từ `category`  
3️⃣ Sửa tên struct / repo / service / handler  
4️⃣ Đăng ký trong `router.go` và `app.go`  
5️⃣ Thêm vào AutoMigrate()

→ Mất chưa đến **5 phút** để có CRUD hoàn chỉnh ✨

---

## 🚀 Run example

```
POST /categories
{
  "name": "Laptop",
  "description": "Máy tính xách tay"
}

GET /products
GET /categories
```

---

## 📄 License

MIT © 2025 Kai Shop Project


## 🛠️ Công nghệ sử dụng

- **PostgreSQL**: Hệ quản trị cơ sở dữ liệu quan hệ mạnh mẽ, hỗ trợ nhiều tính năng nâng cao.
- **GORM**: ORM (Object-Relational Mapping) cho Go, giúp tương tác với cơ sở dữ liệu một cách dễ dàng và hiệu quả.

--- 