# 🛍️ Shop Service

Service หลักสำหรับจัดการประสบการณ์การซื้อขายทั้งหมดของโปรเจกต์ **GameGear E-commerce** ตั้งแต่การแสดงสินค้า, จัดการตะกร้า, ไปจนถึงการสร้างคำสั่งซื้อ

---

## 🏛️ Architectural Design

Service นี้ถูกออกแบบให้เป็น **"เจ้าของข้อมูล" (Data Owner)** โดยมีหน้าที่รับผิดชอบในการจัดการข้อมูลและตารางในฐานข้อมูลที่เกี่ยวกับการซื้อขาย (`products`, `carts`, `orders`, etc.) โดยตรง
Service อื่นๆ ที่ต้องการเข้าถึงข้อมูลเหล่านี้ จะต้องเรียกใช้งานผ่าน API ของ Service นี้เท่านั้น

---

## ✨ Features & Responsibilities

Service นี้รับผิดชอบฟีเจอร์หลัก 3 ส่วน ซึ่งทีมพัฒนาแบ่งงานกันชัดเจน:

* **Product Catalog** (`/api/products`, `/api/products/:id`):

  * แสดงรายการสินค้าทั้งหมดพร้อมระบบแบ่งหน้า (Pagination)
  * แสดงรายละเอียดของสินค้าแต่ละชิ้น
  * จัดการข้อมูลสินค้าและหมวดหมู่

* **Shopping Cart** (`/api/cart/*`):

  * จัดการตะกร้าสินค้าของผู้ใช้ (ดู, เพิ่ม, แก้ไข, ลบ)
  * คำนวณราคารวมของสินค้าในตะกร้า

* **Order Processing** (`/api/orders`, `/api/orders/history`):

  * สร้างคำสั่งซื้อจากข้อมูลในตะกร้า
  * จัดการ Flow การชำระเงินและที่อยู่จัดส่ง
  * แสดงประวัติการสั่งซื้อของผู้ใช้

---

## 📂 Project Structure

โครงสร้างโปรเจกต์ถูกออกแบบให้รองรับการทำงานร่วมกันของทีม โดยแยกโค้ดตามความรับผิดชอบ

<table>
<tr>
<td width="50%">
<pre>
.
├── cmd/api/
│   └── main.go
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
├── .env.example
├── .gitignore
├── go.mod
└── README.md
</pre>
</td>
<td>
<ul>
<li><b>cmd/api</b>: จุดเริ่มต้นการรันโปรแกรม</li>
<li><b>internal</b>: โค้ดหลักของ Service</li>
<ul>
<li><b>handlers</b>: รับ Request และส่ง Response</li>
<li><b>services</b>: จัดการ Business Logic</li>
<li><b>repositories</b>: ติดต่อกับฐานข้อมูล</li>
<li><b>models</b>: กำหนดโครงสร้างข้อมูล (Structs)</li>
</ul>
<li><b>.env.example</b>: ไฟล์ตัวอย่าง Configuration</li>
<li><b>.gitignore</b>: รายการไฟล์ที่ไม่ต้องนำขึ้น Git</li>
</ul>
</td>
</tr>
</table>

---

## 🚀 Getting Started (Step-by-step)

### Step 1 — Clone the Repository (Standard)

```bash
git clone https://github.com/Wattanaroj2567/shop-service.git
cd shop-service
```

### Step 1 (Alt) — Direct Branch Clone (Feature Branch)

ใช้เมื่อทำงานตาม Feature ที่รับผิดชอบ

**สำหรับ ณัฐพงษ์ ดีบุตร (Product Catalog)**

```bash
git clone -b feature/product-catalog https://github.com/Wattanaroj2567/shop-service.git
cd shop-service
```

**สำหรับ วายุ กอคูณ (Cart & Order)**

```bash
git clone -b feature/cart-and-order https://github.com/Wattanaroj2567/shop-service.git
cd shop-service
```

### Step 2 — Install Dependencies

```bash
go mod tidy
```

### Step 3 — Setup Database

ตรวจสอบให้ PostgreSQL Server ทำงานอยู่ และสร้างฐานข้อมูล (ถ้ายังไม่มี)

**(A) SQL โดยตรง**

```sql
CREATE DATABASE gamegear_shop_db;
```

**(B) ใช้ psql ผ่าน bash**

```bash
psql -U your_user -h localhost -p 5432 -c "CREATE DATABASE gamegear_shop_db;"
```

### Step 4 — Configure Environment Variables

สร้างไฟล์ `.env` และกำหนดค่าตามตัวอย่าง

```env
# Core Configuration
APPLICATION_PORT=8081

# PostgreSQL Database
DATABASE_URL="host=localhost user=your_user password=your_password dbname=gamegear_shop_db port=5432 sslmode=disable"
```

### Step 5 — Run the Service

```bash
go run cmd/api/main.go
```

> ระบบจะทำการ migrate ตารางที่จำเป็น และเซิร์ฟเวอร์จะเริ่มที่ `http://localhost:8081`

---

## 📝 API Documentation: Swagger (OpenAPI)

### Step 1 — ติดตั้ง `swag`

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Step 2 — สร้างไฟล์ Swagger Docs

```bash
swag init
```

> จะสร้างโฟลเดอร์ `docs` และไฟล์ที่จำเป็นอัตโนมัติ

### Step 3 — เปิดดู API Docs

```
http://localhost:8081/swagger/index.html
```

---

## 🤝 Remote Development (ngrok)

### Step 1 — ติดตั้ง ngrok

[ดาวน์โหลดได้ที่ ngrok.com](https://ngrok.com)

### Step 2 — รัน Service

```bash
go run cmd/api/main.go
```

### Step 3 — เปิดอุโมงค์

```bash
ngrok http 8081
```

### Step 4 — แชร์ URL

คัดลอก URL ที่ขึ้นต้นด้วย `https://...` ส่งให้ทีม

### Step 5 — ตั้งค่าใน .env ของ admin-service

```env
SHOP_SERVICE_URL="<THE_NGROK_URL_YOU_SENT>"
```

---

## 🌱 Branching Strategy

เพื่อความเป็นระเบียบในการทำงานร่วมกัน จะใช้ Git Flow แบบง่าย

* `main` → ห้าม push โดยตรง
* `develop` → branch หลักสำหรับการพัฒนา
* ใช้ **Feature Branch** สำหรับงานของแต่ละคน

  * ณัฐพงษ์ ดีบุตร → `feature/product-catalog`
  * วายุ กอคูณ → `feature/cart-and-order`
* เมื่อเสร็จงาน → สร้าง Pull Request (PR) ไปที่ `develop`
* ต้องมี Code Review โดยอีกคนก่อน Merge
