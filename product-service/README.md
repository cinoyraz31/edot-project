📦 Product Service

Sebelum memulai running service ini sebaiknya mnigration dulu terkait product di migration/product_store.go
Karena saya tidak ada aksi/ API untuk CRUD product

```azure
anda bisa jalankan
go run migration/product_store.go
```

Product Service digunakan untuk mengelola seluruh informasi produk:

- Detail produk (nama, deskripsi, kategori)
- SKU, model, dan varian (ukuran, warna, dll)
- Produk hanya dibuat oleh pusat, bukan per toko
- Harga dapat disesuaikan oleh setiap toko (relasi dengan Shop)

#### 🖼️ Ilustrasi Product Service

![Product Service Flow](https://i.imgur.com/U8Lybjd.png)