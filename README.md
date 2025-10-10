# 🛍️ Shop Service — GameGear E-commerce

บริการสำหรับจัดการประสบการณ์การซื้อขายทั้งหมดของโปรเจกต์ **GameGear E-commerce** ตั้งแต่การแสดงสินค้า, การจัดการตะกร้า, ไปจนถึงการสร้างคำสั่งซื้อ เป็นหนึ่งใน microservices ที่ทำงานร่วมกับ **Kong API Gateway**

---

## 🏛️ Architectural Design & Responsibility

* เป็น **เจ้าของข้อมูล (Data Owner)** สำหรับตารางที่เกี่ยวข้องกับการซื้อขาย เช่น `products`, `carts`, `orders` เป็นต้น
* ข้อมูลของสินค้าและคำสั่งซื้อ **จะไม่ถูกเข้าถึงโดยตรงจาก service อื่น** — ต้องเรียกผ่าน API ของ Shop Service เท่านั้น
* บริการอื่น (Users, Admin) ที่ต้องการใช้ข้อมูลเหล่านี้ ต้องเรียกผ่าน Gateway เพื่อรักษาความถูกต้องและความปลอดภัยของข้อมูล

---

## ✨ Features & Endpoints

| Feature          | HTTP Method               | Path                                 | Description                              |
| ---------------- | ------------------------- | ------------------------------------ | ---------------------------------------- |
| Product Catalog  | GET                       | `/api/products`, `/api/products/:id` | แสดงสินค้าทั้งหมด / รายละเอียดสินค้า     |
| Shopping Cart    | GET / POST / PUT / DELETE | `/api/cart/*`                        | จัดการตะกร้าสินค้าของผู้ใช้              |
| Order Processing | GET / POST                | `/api/orders`, `/api/orders/history` | สร้างคำสั่งซื้อ / แสดงประวัติการสั่งซื้อ |

นอกจากนี้ ควรมี endpoint สำหรับ **Healthcheck**:

```
GET /healthz → 200 OK
```

---

## 📂 Project Structure (Standard Go Layout)

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   ├── product_handler.go
│   │   └── cart_order_handler.go
│   ├── services/
│   │   ├── product_service.go
│   │   └── cart_order_service.go
│   ├── repositories/
│   │   ├── product_repository.go
│   │   └── cart_order_repository.go
│   └── models/
│       ├── product_model.go
│       ├── cart_model.go
│       └── order_model.go
├── docs/
│   └── swagger (ไฟล์ที่สร้างโดย swag)
├── .env.example
├── go.mod
├── go.sum
├── Dockerfile.dev
├── docker-compose.override.yml
├── docker-compose.kong.yml
├── .dockerignore
└── README.md
```

คำอธิบาย:

* **cmd/api** — จุดเริ่มต้นของโปรแกรมหลัก (main entry point)
* **internal/handlers** — ชั้นรับคำขอ (Request) และส่งผลลัพธ์กลับ (Response)
* **internal/services** — ชั้นของ Business Logic ที่จัดการการทำงานหลักของระบบ
* **internal/repositories** — ชั้นเชื่อมต่อกับฐานข้อมูล เช่น Query, Insert, Update, Delete
* **internal/models** — กำหนดโครงสร้างข้อมูล (Struct) สำหรับ ORM (GORM)
* **docs/swagger** — เอกสาร API ที่สร้างจาก `swag init`
* **.env.example** — ตัวอย่างไฟล์ตั้งค่า environment variables
* **Dockerfile.dev** — Dockerfile สำหรับโหมดพัฒนา (Dev Mode) พร้อม hot-reload ด้วย `air`
* **docker-compose.override.yml** — ใช้รัน shop-service พร้อม PostgreSQL ในโหมดพัฒนา
* **docker-compose.kong.yml** — ใช้เป็นแนวทางการเชื่อมต่อกับ Kong Gateway ตามสถาปัตยกรรมของโปรเจกต์หลัก (Mini-Project-Golang)
* **.dockerignore** — รายการไฟล์/โฟลเดอร์ที่ไม่ต้องการนำเข้าเวลาสร้าง image
* **README.md** — เอกสารอธิบายวิธีติดตั้งและใช้งาน service

---

## 🚀 Getting Started (Local Development)

### 1. Clone Repository (ตามหน้าที่ของแต่ละคน)

**สำหรับ ณัฐพงษ์ ดีบุตร — Product Catalog**

```bash
git clone -b feature/product-catalog https://github.com/Wattanaroj2567/shop-service.git
cd shop-service
```

**สำหรับ วายุ กอคูณ — Cart & Order**

```bash
git clone -b feature/cart-and-order https://github.com/Wattanaroj2567/shop-service.git
cd shop-service
```

> แต่ละคนให้ clone branch ของตนเองโดยตรง เพื่อแยกพื้นที่การพัฒนาและลดโอกาส conflict

---

### 2. ติดตั้ง dependencies

```bash
go mod tidy
```

### 3. ตั้งค่า Database

ตรวจสอบว่า PostgreSQL ทำงานอยู่ และสร้างฐานข้อมูล:

```sql
CREATE DATABASE gamegear_shop_db;
```

### 4. ตั้งค่า Environment Variables

สร้าง `.env` โดยใช้ค่าจากตัวอย่าง:

```env
APPLICATION_PORT=8081
DATABASE_URL="host=localhost user=dev password=dev dbname=gamegear_shop_db port=5432 sslmode=disable"
JWT_SECRET_KEY="supersecretkey"
```

### 5. รัน Service

```bash
go run cmd/api/main.go
```

---

## 🐋 Run with Docker (Dev)

### Dockerfile.dev

```dockerfile
FROM golang:1.22-alpine
RUN apk add --no-cache git bash build-base tzdata ca-certificates \
    && update-ca-certificates \
    && go install github.com/cosmtrek/air@latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8081
CMD ["air"]
```

### docker-compose.override.yml

```yaml
version: "3.9"
services:
  shop-db:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: gamegear_shop_db
      POSTGRES_USER: dev
      POSTGRES_PASSWORD: dev
    ports:
      - "5433:5432"

  shop-service:
    build:
      context: .
      dockerfile: Dockerfile.dev
    environment:
      DATABASE_URL: postgres://dev:dev@shop-db:5432/gamegear_shop_db?sslmode=disable
      APPLICATION_PORT: 8081
      JWT_SECRET_KEY: "supersecretkey"
    ports:
      - "8081:8081"
    depends_on:
      - shop-db
    volumes:
      - .:/app
```

### docker-compose.kong.yml

```yaml
version: "3.9"
services:
  shop-service:
    build:
      context: .
      dockerfile: Dockerfile.dev
    container_name: shop-service
    environment:
      APPLICATION_PORT: 8081
      DATABASE_URL: postgres://dev:dev@shop-db:5432/gamegear_shop_db?sslmode=disable
    ports:
      - "8081:8081"
    networks:
      - gamegear-network
    depends_on:
      - shop-db

  shop-db:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: gamegear_shop_db
      POSTGRES_USER: dev
      POSTGRES_PASSWORD: dev
    networks:
      - gamegear-network

networks:
  gamegear-network:
    external: true  # ใช้ network เดียวกับ Kong จากโปรเจกต์หลัก
```

> ไฟล์นี้ใช้เป็นแนวทางการเชื่อมต่อกับ Kong Gateway ตามสถาปัตยกรรมของโปรเจกต์หลัก (Mini-Project-Golang)

---

## 📝 API Documentation (Swagger / OpenAPI)

### ติดตั้ง Swag

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### สร้างเอกสาร Swagger

```bash
swag init
```

จะสร้างโฟลเดอร์ `docs` และไฟล์ Swagger JSON/YAML ให้อัตโนมัติ

### เปิดใช้งาน Swagger UI

```
http://localhost:8080/swagger/index.html
```

---

## 🔁 Remote Development (ngrok)

ใช้ในกรณีที่ทีมพัฒนาอยู่คนละเครื่อง และต้องเชื่อม service ผ่านอินเทอร์เน็ต:

```bash
ngrok http 8081
```

จากนั้นนำ URL ไปตั้งใน `.env` ของ `admin-service` หรือ Gateway เช่น:

```env
SHOP_SERVICE_URL=https://abc1234.ngrok.io
```

---

## 🧭 ทำไมต้องใช้ ngrok ร่วมกับ Kong Gateway

| เหตุผล                         | รายละเอียด                                                                                                                                       |
| ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| 🌍 การทำงานคนละเครื่อง         | เมื่อสมาชิกทีมรัน service ของตนเอง (เช่น คุณรัน shop-service แต่เพื่อนรัน users-service) ต้องใช้ ngrok แชร์ URL เพื่อให้เชื่อมต่อได้ผ่าน Gateway |
| ⚙️ Integration กับ Kong        | Kong สามารถใช้ URL จาก ngrok เป็นปลายทาง service ได้ เช่น `SHOP_SERVICE_URL=https://abc1234.ngrok.io`                                            |
| 🚀 สะดวกต่อการทดสอบ            | ใช้ทดสอบระบบรวม (Integration Test) โดยไม่ต้อง deploy ขึ้นเซิร์ฟเวอร์กลาง                                                                         |
| 🎓 ใช้สาธิต/ส่งอาจารย์ได้ทันที | เปิดระบบในเครื่องแล้วแชร์ให้ผู้สอนหรือผู้ทดสอบเข้าถึงผ่านอินเทอร์เน็ตได้ทันที                                                                    |

> 🔸 ถ้าอยู่ใน Docker Compose เดียวกัน → ไม่ต้องใช้ ngrok
> 🔹 ถ้าอยู่คนละเครื่อง → ใช้ ngrok เพื่อให้ Gateway เชื่อมต่อได้ครบ

---

## 🌱 Branching Strategy

| Branch                    | Description                                 |
| ------------------------- | ------------------------------------------- |
| `main`                    | ห้าม push โดยตรง ใช้สำหรับ release เท่านั้น |
| `develop`                 | branch หลักสำหรับการพัฒนา                   |
| `feature/product-catalog` | งานของ **ณัฐพงษ์ ดีบุตร** (จัดการสินค้า)    |
| `feature/cart-and-order`  | งานของ **วายุ กอคูณ** (ตะกร้าและคำสั่งซื้อ) |

> เมื่อเสร็จงานให้สร้าง Pull Request (PR) ไปที่ `develop` และต้องผ่านการ review อย่างน้อย 1 คนก่อน merge

---

## ✅ Summary

* README นี้อัปเดตให้สอดคล้องกับ **แนวทางหลักของโปรเจกต์ Mini-Project-Golang**
* รองรับทั้งการพัฒนา, ทดสอบ และเชื่อมต่อกับ Kong Gateway
* มีวิธีรันแบบ local, Docker, Swagger, Kong Integration และ Remote Dev พร้อมใช้งานจริง
