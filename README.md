# 🛍️ Shop Service

Service หลักสำหรับจัดการประสบการณ์การซื้อขายทั้งหมดของโปรเจกต์ **GameGear E-commerce** ตั้งแต่การแสดงสินค้า, จัดการตะกร้า, ไปจนถึงการสร้างคำสั่งซื้อ

---

## 🏛️ Architectural Design

Service นี้ถูกออกแบบให้เป็น **"เจ้าของข้อมูล" (Data Owner)** โดยมีหน้าที่รับผิดชอบในการจัดการข้อมูลและตารางในฐานข้อมูลที่เกี่ยวกับการซื้อขายทั้งหมด (`products`, `carts`, `orders`, etc.) โดยตรง
Service อื่นๆ ที่ต้องการเข้าถึงข้อมูลเหล่านี้ จะต้องเรียกใช้งานผ่าน API ที่ Service นี้มีให้เท่านั้น

---

## ✨ Features & Responsibilities

Service นี้รับผิดชอบฟีเจอร์หลัก 3 ส่วน ซึ่งจะแบ่งการพัฒนากันในทีม:

* **Product Catalog** (`/api/products`, `/api/products/:id`)

  * แสดงรายการสินค้าทั้งหมดพร้อมระบบแบ่งหน้า (Pagination)
  * แสดงรายละเอียดของสินค้าแต่ละชิ้น
  * จัดการข้อมูลสินค้าและหมวดหมู่

* **Shopping Cart** (`/api/cart/*`)

  * จัดการตะกร้าสินค้าของผู้ใช้ (ดู, เพิ่ม, แก้ไข, ลบ)
  * คำนวณราคารวมของสินค้าในตะกร้า

* **Order Processing** (`/api/orders`, `/api/orders/history`)

  * สร้างคำสั่งซื้อจากข้อมูลในตะกร้าสินค้า
  * จัดการ Flow การชำระเงินและที่อยู่จัดส่ง
  * แสดงประวัติการสั่งซื้อของผู้ใช้

---

## 📂 Project Structure

โครงสร้างของโปรเจกต์ถูกออกแบบมาเพื่อรองรับการทำงานร่วมกันของทีมพัฒนา 2 คน โดยแยกไฟล์ตามความรับผิดชอบ

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
├── .env
├── go.mod
└── README.md
</pre>
</td>
<td>
  <ul>
    <li><b>cmd/api</b>: จุดเริ่มต้นในการรันโปรแกรม</li>
    <li><b>internal</b>: โค้ดหลักทั้งหมดของ Service</li>
    <ul>
      <li><b>handlers</b>: ส่วนรับ Request และส่ง Response</li>
      <li><b>services</b>: ส่วนจัดการ Logic หลัก</li>
      <li><b>repositories</b>: ส่วนสื่อสารกับฐานข้อมูล</li>
      <li><b>models</b>: ส่วนกำหนดโครงสร้างข้อมูล (Structs)</li>
    </ul>
    <li><b>.env</b>: ไฟล์เก็บ Configuration</li>
  </ul>
</td>
</tr>
</table>

---

## 🚀 Getting Started

ทำตามขั้นตอนทีละสเต็ป เพื่อเตรียมและรัน Service ในเครื่องของคุณ

### Step 1 — Clone the Repository (Standard) สำหรับใช้งานได้จริง

```bash
git clone https://github.com/Wattanaroj2567/shop-service.git
cd shop-service
```

### Step 1 (Alt) — Direct Branch Clone สำหรับผู้พัฒนาให้เริ่มต้นที่นี่เท่านั้น
คำสั่งนี้จะ Clone โปรเจกต์ทั้งหมดมาที่เครื่องของคุณ โดยจะได้ Branch ของชื่อผู้รับผิดชอบตามที่ได้รับหน้าที่
> เลือกใช้เมื่อทราบชื่อ branch ที่ต้องการทำงานแล้ว

**สำหรับณัฐพงษ์ ดีบุตร (Product Catalog)**

```bash
git clone -b feature/product-catalog https://github.com/Wattanaroj2567/shop-service.git
cd shop-service
```

**สำหรับวายุ กอคูณ (Cart & Order)**

```bash
git clone -b feature/cart-and-order https://github.com/Wattanaroj2567/shop-service.git
cd shop-service
```

### Step 2 — Install Dependencies

```bash
go mod tidy
```

### Step 3 — Setup Database

เลือกวิธีใดวิธีหนึ่ง

**(A) ใช้ SQL โดยตรง**

```sql
CREATE DATABASE gamegear_db;
```

**(B) ใช้ psql ผ่าน bash (one‑liner)**

```bash
psql -U your_user -h localhost -p 5432 -c "CREATE DATABASE gamegear_db;"
```

### Step 4 — Configure Environment Variables

สร้างไฟล์ `.env` แล้วใส่ค่าเชื่อมต่อฐานข้อมูล

```env
# PostgreSQL Database Connection URL
DATABASE_URL="host=localhost user=your_user password=your_password dbname=gamegear_db port=5432 sslmode=disable"
```

### Step 5 — Run the Service

```bash
go run cmd/api/main.go
```

> เมื่อรันคำสั่งนี้ ระบบจะทำการ **migrate** ตารางที่จำเป็นทั้งหมด และเซิร์ฟเวอร์จะเริ่มที่ `http://localhost:8081`

---

## 🤝 Remote Development (Working from Different Locations)

ต้องการให้เพื่อนร่วมทีมเข้าถึง `shop-service` ของคุณจากภายนอก? ใช้ **ngrok** ตามขั้นตอนนี้

### Step 1 — ติดตั้ง/ดาวน์โหลด ngrok

ไปที่ [ngrok.com](https://ngrok.com) และติดตั้งตามระบบปฏิบัติการของคุณ

### Step 2 — รัน User Service ในเครื่อง

```bash
go run cmd/api/main.go
```

### Step 3 — เปิดอุโมงค์ไปยังพอร์ต 8081

เปิด Terminal ใหม่แล้วรัน

```bash
ngrok http 8081
```

### Step 4 — แชร์ URL ให้ทีม

คัดลอก URL ที่ขึ้นต้นด้วย `https://...` ส่งให้เพื่อนร่วมทีม

### Step 5 — เพื่อนนำ URL ไปตั้งใน .env (ฝั่ง admin-service)

```env
# .env file on admin-service
SHOP_SERVICE_URL="<THE_NGROK_URL_YOU_SENT>"
```

---

## 🌱 Branching Strategy

เพื่อให้การทำงานร่วมกันของทีมเป็นไปอย่างราบรื่น เราจะใช้ Git Flow แบบง่าย โดยมี Branch หลักคือ `main` และ `develop`

### Workflow

* **ห้าม** Push Code ขึ้น `main` โดยตรง
* `develop` = Branch หลักสำหรับการพัฒนา → ทุกงานต้อง Pull ล่าสุดมาก่อน
* ทำงานบน **Feature Branch** ของตัวเองเสมอ:

  * **ณัฐพงษ์ ดีบุตร (Product Catalog):** `feature/product-catalog`
  * **วายุ กอคูณ (Cart & Order):** `feature/cart-and-order`
* เมื่อทำฟีเจอร์เสร็จ → สร้าง Pull Request (PR) เข้าสู่ `develop`
* เพื่อนอีกคนต้องช่วยตรวจสอบ (Code Review) ก่อน Merge
