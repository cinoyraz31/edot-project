# 🛒 E-Commerce Microservices System

## 📦 Overview

Proyek ini adalah sistem e-commerce berbasis **microservices** yang terdiri dari beberapa layanan terpisah namun saling terintegrasi. Sistem ini menyerupai model **retail multi-outlet** seperti Erafone, di mana satu produk dapat dijual oleh banyak toko dengan harga yang berbeda-beda.

## 📁 Services

### 1. 👤 User Service
Manajemen pengguna sistem:

- Registrasi dan login menggunakan nomor telepon dengan verifikasi OTP
- Mendukung dua tipe pengguna:
    - **Client**: Pengguna akhir (customer)
    - **Shop User**: Pengelola toko

### 2. 🧾 Product Service
Mengelola informasi produk:

- Data produk seperti nama, varian, ukuran, dan kategori
- Produk bersifat **tersentralisasi**
- Harga ditentukan oleh masing-masing toko (bukan di product service)

### 3. 🏪 Shop Service
Mengelola data toko:

- Toko memiliki relasi ke produk (`ShopProduct`)
- Toko dapat menentukan **harga berbeda** untuk produk yang sama
- Setiap toko dapat memiliki satu atau lebih gudang (`Warehouse`)

### 4. 🏬 Warehouse Service
Mengelola stok produk:

- Stok disimpan per gudang
- Fitur:
    - Penambahan stok oleh toko
    - Transfer stok antar gudang
    - Aktivasi dan non-aktivasi gudang
    - Lock stok saat order dan release jika pembayaran gagal

### 5. 📦 Order Service
Mengelola proses pemesanan:

- Checkout produk
- Pengecekan dan penguncian stok di warehouse
- Support order dari beberapa toko dalam satu transaksi
- Struktur order:
    - `Order` memiliki banyak `ShopOrder`
    - `ShopOrder` memiliki banyak `OrderItem`
- Release stok otomatis jika tidak dibayar dalam waktu tertentu

---

## 🔄 Order Flow

1. User membuat order berisi produk dari satu atau lebih toko.
2. Sistem mengecek dan mengunci stok dari warehouse toko.
3. Jika pembayaran dilakukan:
    - Stok dikurangi permanen
    - Order diproses ke pengiriman
4. Jika tidak dibayar dalam waktu tertentu:
    - Stok otomatis dirilis kembali

---

## 🚀 Business Highlights

- Produk tersentralisasi, harga fleksibel per toko
- Order bisa berasal dari banyak toko dalam satu transaksi
- Mekanisme lock/unlock stok untuk mencegah over-selling
- Arsitektur microservice: scalable dan modular

---

## ⚙️ Tech Stack

- **Language**: Go (Golang)
- **Framework**: Fiber
- **Database**: Mysql
- **ORM**: GORM
- **Auth**: JWT
- **Containerization**: Docker & Docker Compose

---

## 📦 Installation

### Step 1: You Will Prepare Create Database
- edot-user-service
- edot-product-service
- edot-shop-service
- edot-warehouse-service
- edot-order-service

### Step 2: Clone Every Repository
```bash
git clone https://github.com/example.git
cd project-name

go mod download
go mod tidy
go run main.go

```

Saya membahas secara flow ilustrasi disetiap service README.md