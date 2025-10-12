# 🛍️ Shop Service — GameGear E-commerce

บริการสำหรับจัดการประสบการณ์การซื้อขายทั้งหมดของโปรเจกต์ **GameGear E-commerce** ตั้งแต่การแสดงสินค้า, การจัดการตะกร้า, ไปจนถึงการสร้างคำสั่งซื้อ เป็นหนึ่งใน microservices ที่ทำงานร่วมกับ **Kong API Gateway**

> 📖 **ดูเอกสารหลักของระบบ**: สำหรับ Kong Gateway setup, Architecture overview และการ integrate ทั้งระบบ → [Main README](../Mini-Project-Golang/README.md)

---

## 🏛️ Architectural Design & Responsibility

- เป็น **เจ้าของข้อมูล (Data Owner)** สำหรับตารางที่เกี่ยวข้องกับการซื้อขาย เช่น `products`, `carts`, `orders` เป็นต้น
- ข้อมูลของสินค้าและคำสั่งซื้อ **จะไม่ถูกเข้าถึงโดยตรงจาก service อื่น** — ต้องเรียกผ่าน API ของ Shop Service เท่านั้น
- บริการอื่น (Users, Admin) ที่ต้องการใช้ข้อมูลเหล่านี้ ต้องเรียกผ่าน Gateway เพื่อรักษาความถูกต้องและความปลอดภัยของข้อมูล

---

## ✅ สิ่งที่ต้องทำ (สำหรับผู้พัฒนา Service นี้)

> 👨‍💻 **Developers**:
>
> - **ณัฐพงษ์ ดีบุตร** - Product Catalog (branch: `feature/product-catalog`)
> - **วายุ กอคูณ** - Cart & Order (branch: `feature/cart-and-order`)

### 📝 ไฟล์ที่ต้องแก้ไข/เขียนโค้ด

คุณต้องเขียนโค้ดเฉพาะใน **2 จุดหลัก**:

#### 1. **`cmd/api/main.go`**

- เริ่มต้น Gin server
- Setup routes
- Connect database
- Middleware configuration

#### 2. **โฟลเดอร์ `internal/`**

| Folder               | ต้องทำ       | หมายเหตุ                                        |
| -------------------- | ------------ | ----------------------------------------------- |
| ✅ **handlers/**     | ✅ ต้องทำ    | เขียน HTTP handlers สำหรับ products/cart/orders |
| ❌ **models/**       | ❌ ไม่ต้องทำ | **PM ทำให้แล้ว** (วรรธนโรจน์)                   |
| ✅ **repositories/** | ✅ ต้องทำ    | เขียน database operations (CRUD)                |
| ✅ **services/**     | ✅ ต้องทำ    | เขียน business logic                            |

#### 3. **ไฟล์อื่นๆ**

- ✅ **`.env`** - ตั้งค่า environment variables
- ✅ **`go.mod`** - เพิ่ม dependencies ตามต้องการ

### 🚫 ไฟล์ที่ไม่ต้องแก้

- ❌ `internal/models/` - PM ทำให้แล้ว
- ❌ `README.md` - มีให้แล้ว (แต่สามารถเพิ่มเติมได้)

### 📋 Checklist

**สำหรับ ณัฐพงษ์ (Products):**

- [ ] เขียน `product_handler.go`
- [ ] เขียน `product_repository.go`
- [ ] เขียน `product_service.go`
- [ ] API: GET /products, GET /products/:id, POST /products, PUT /products/:id, DELETE /products/:id

**สำหรับ วายุ (Cart & Orders):**

- [ ] เขียน `cart_order_handler.go`
- [ ] เขียน `cart_order_repository.go`
- [ ] เขียน `cart_order_service.go`
- [ ] API: Cart (GET, POST, PUT, DELETE) + Orders (GET, POST)

**ทั้งคู่:**

- [ ] รวม routes ใน `cmd/api/main.go`
- [ ] ตั้งค่า `.env`
- [ ] ทดสอบ API ด้วย Swagger
- [ ] ทดสอบผ่าน Kong Gateway

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
└── README.md
```

คำอธิบาย:

- **cmd/api** — จุดเริ่มต้นของโปรแกรมหลัก (main entry point)
- **internal/handlers** — ชั้นรับคำขอ (Request) และส่งผลลัพธ์กลับ (Response)
- **internal/services** — ชั้นของ Business Logic ที่จัดการการทำงานหลักของระบบ
- **internal/repositories** — ชั้นเชื่อมต่อกับฐานข้อมูล เช่น Query, Insert, Update, Delete
- **internal/models** — กำหนดโครงสร้างข้อมูล (Struct) สำหรับ ORM (GORM)
- **docs/swagger** — เอกสาร API ที่สร้างจาก `swag init`
- **.env.example** — ตัวอย่างไฟล์ตั้งค่า environment variables
- **Kong Gateway** — จัดการโดย PM (วรรธนโรจน์) ใน admin-service
- **README.md** — เอกสารอธิบายวิธีติดตั้งและใช้งาน service

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


## 🦍 Kong API Gateway Integration

Service นี้เป็นส่วนหนึ่งของระบบ Microservices ที่ใช้ **Kong Gateway** เป็นจุดเข้าถึงหลัก (API Gateway)

> 👤 **ผู้ดูแล Kong Gateway**: วรรธนโรจน์ บุตรดี (Project Manager)  
> 💡 **คำแนะนำ**: หากมีปัญหาเกี่ยวกับ Kong setup, Routes หรือ Plugins โปรดติดต่อผู้ดูแล

### 🚀 Quick Start Options

#### Option 1: รันพร้อม Kong + Konga (แนะนำสำหรับ Integration Testing)

```bash
# จาก root directory (GameGear-Ecommerce/)
# Kong Gateway จัดการโดย PM ใน admin-service
# สำหรับการทดสอบ service เดี่ยว ให้ใช้:
go run cmd/api/main.go
```

**Services ที่จะรัน:**

- 🦍 Kong Gateway (port 8000, 8001)
- 🖥️ Konga Admin UI (port 1337)
- 🛍️ Shop Service (port 8081)
- 🗄️ PostgreSQL Databases (Kong, Konga, Shop)

#### Option 2: รัน Service เดี่ยว (สำหรับ Local Development)

```bash
# จาก service directory
go run cmd/api/main.go
```

**Service ที่จะรัน:**

- 🛍️ Shop Service (port 8081)

**หมายเหตุ:** ต้องมี PostgreSQL database รันอยู่แล้ว หรือใช้ cloud database

### 🌐 การใช้ ngrok สำหรับ External Access

```bash
# รัน service
go run cmd/api/main.go

# เปิด terminal ใหม่ และรัน ngrok
ngrok http 8081
```

**ผลลัพธ์:**

- Service รันที่: `http://localhost:8081`
- External URL: `https://abc123.ngrok.io` (สำหรับ PM เรียกใช้)

---

### 📍 API Endpoints

| Access Method             | URL                                             | Use Case       | Note                      |
| ------------------------- | ----------------------------------------------- | -------------- | ------------------------- |
| **Via Kong Gateway**      | `http://localhost:8000/shop/*`                  | ✅ Production  | เรียกผ่าน Gateway (แนะนำ) |
| **Direct Access**         | `http://localhost:8081/*`                       | 🔧 Development | เรียกตรง (Dev/Test only)  |
| **Swagger UI (via Kong)** | `http://localhost:8000/shop/swagger/index.html` | 📖 API Docs    | ผ่าน Gateway              |
| **Swagger UI (direct)**   | `http://localhost:8081/swagger/index.html`      | 📖 API Docs    | เรียกตรง                  |

### 🔧 Kong Configuration

หากต้องการตั้งค่า Service นี้ใน Kong Gateway:

1. **เปิด Konga UI**: http://localhost:1337
2. **เพิ่ม Service**:
   ```
   Name: shop-service
   Protocol: http
   Host: host.docker.internal
   Port: 8081
   Path: /
   ```
3. **เพิ่ม Route**:
   ```
   Name: shop-route
   Paths: /shop
   Strip Path: false
   ```

### 🧪 ทดสอบการเชื่อมต่อ

```bash
# ทดสอบผ่าน Kong Gateway
curl http://localhost:8000/shop/healthz

# ทดสอบเรียกตรง (Dev only)
curl http://localhost:8081/healthz

# ทดสอบดูสินค้าทั้งหมด via Kong
curl http://localhost:8000/shop/products

# ทดสอบดูรายละเอียดสินค้า
curl http://localhost:8000/shop/products/1
```

### 📖 เอกสารเพิ่มเติม

สำหรับข้อมูลเพิ่มเติมเกี่ยวกับ Kong Gateway และการตั้งค่าทั้งระบบ:

- **Kong + Konga Setup Guide**: [Main README - Kong Setup](../Mini-Project-Golang/README.md#-quick-start-ติดตั้งและรัน-kong--konga)
- **System Architecture**: [Main README - Architecture](../Mini-Project-Golang/README.md#%EF%B8%8F-system-architecture-overview)
- **Ports Summary**: [Main README - Ports](../Mini-Project-Golang/README.md#-ports-summary)
- **Troubleshooting**: [Main README - Troubleshooting](../Mini-Project-Golang/README.md#-troubleshooting)

> 💡 **หมายเหตุ**: สำหรับการ setup Kong, Konga และ Plugins (CORS, JWT, Rate Limiting) โปรดดูเอกสารหลักที่ [Main README](../Mini-Project-Golang/README.md)

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

- README นี้อัปเดตให้สอดคล้องกับ **แนวทางหลักของโปรเจกต์ Mini-Project-Golang**
- รองรับทั้งการพัฒนา, ทดสอบ และเชื่อมต่อกับ Kong Gateway
- มีวิธีรันแบบ local, Docker, Swagger, Kong Integration และ Remote Dev พร้อมใช้งานจริง
