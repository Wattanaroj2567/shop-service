# 🛍️ Shop Service

![Go](https://img.shields.io/badge/Go-1.24.6-00ADD8?style=for-the-badge\&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge\&logo=postgresql)
![GORM](https://img.shields.io/badge/GORM-B93527?style=for-the-badge)

Service หลักสำหรับจัดการประสบการณ์การซื้อขายทั้งหมดของโปรเจกต์ **GameGear E-commerce** ตั้งแต่การแสดงสินค้า, จัดการตะกร้า, ไปจนถึงการสร้างคำสั่งซื้อ

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

ทำตามขั้นตอนเหล่านี้เพื่อตั้งค่าโปรเจกต์และรัน Service ในเครื่องของคุณ

### 1. Clone the Repository

```bash
git clone https://github.com/Wattanaroj2567/shop-service.git
```
```bash
cd shop-service
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Setup Database

* ตรวจสอบให้แน่ใจว่า PostgreSQL Server ของคุณทำงานอยู่
* สร้างฐานข้อมูลสำหรับโปรเจกต์ (หากยังไม่มี):

```sql
CREATE DATABASE gamegear_db;
```

### 4. Configure Environment Variables

สร้างไฟล์ `.env` และใส่ค่าการเชื่อมต่อฐานข้อมูล (อย่าลืมแก้ `your_user` และ `your_password`)

```env
# PostgreSQL Database Connection URL
DATABASE_URL="host=localhost user=your_user password=your_password dbname=gamegear_db port=5432 sslmode=disable"
```

### 5. Run the Service

```bash
go run cmd/api/main.go
```

* เมื่อรันคำสั่งนี้ โปรแกรมจะทำการ **Migrate** สร้างตารางที่จำเป็นทั้งหมด (`products`, `categories`, `carts`, `cart_items`, `orders`, `order_items`)
* เซิร์ฟเวอร์จะเริ่มต้นทำงานที่ `http://localhost:8081` 
