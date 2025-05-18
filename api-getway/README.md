# API Gateway Overview

## Apa Itu API Gateway?

**API Gateway** adalah sebuah komponen yang berfungsi sebagai *single entry point* untuk seluruh permintaan (request) dari klien ke berbagai layanan (services) dalam sistem. Dengan kata lain, API Gateway menjadi pintu gerbang utama sebelum request diteruskan ke masing-masing layanan backend yang relevan.

## Fungsi Utama API Gateway

- 🔐 **Security**: Menyediakan autentikasi dan otorisasi secara terpusat.
- 📊 **Rate Limiting & Throttling**: Mengontrol jumlah request dalam waktu tertentu.
- 🔁 **Proxying**: Meneruskan request ke service yang dituju tanpa mengekspos detail internal.
- 🧾 **Logging & Monitoring**: Menyediakan log dan metrik dari seluruh permintaan.
- ⚡ **Caching**: Menyimpan response sementara untuk meningkatkan performa.
- 🔀 **Routing**: Mengarahkan request berdasarkan prefix URL, header, atau metode HTTP.

## Manfaat Penggunaan API Gateway

### 1. Penyederhanaan Pengembangan
Setiap service tidak perlu tahu lokasi service lain. Cukup gunakan API Gateway untuk mengakses semua service dengan rute seperti:
- `api/users/...` → User Service
- `api/orders/...` → Order Service
- `api/warehouses/...` → Warehouse Service
- `api/products/...` → Product Service
- `api/shop/...` → Shop Service

Ini mengurangi kebutuhan mengelola banyak environment URL dan membuat kode lebih bersih dan maintainable.

### 2. Keamanan dan Isolasi
Layanan internal tidak perlu diekspos ke luar. Hanya API Gateway yang terbuka untuk publik, sehingga permukaan serangan lebih kecil dan sistem lebih aman.

### 3. Efisiensi Komunikasi Internal
Dengan API Gateway yang berjalan dalam satu cluster atau jaringan privat, komunikasi antarlayanan bisa dilakukan secara lokal (misalnya via `localhost`). Ke
