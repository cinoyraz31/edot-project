### 📦 Order Service
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

![Order Service Flow](https://i.imgur.com/xdTttRN.png)