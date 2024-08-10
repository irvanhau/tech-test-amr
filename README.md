## Technical Test PT AMANTRA INVESTAMA INDODANA

1. Jalankan perintah `git clone https://github.com/irvanhau/tech-test-amr` di command line atau terminal
2. Jalankan perintah `go mod tidy`
3. Jalankan perintah `cp .env.example .env`
4. Buatlah database di postgresql lalu isi `.env`
5. Jalankan perintah `go run main.go`
6. Selamat, Aplikasi berhasil dijalankan

### Tech Stack
- Framework Echo
- Database PostgreSQL dan Redis
- ORM menggunakan GORM

### Unit Test
Unit Test sudah dilakukan di service dan 100% coverage

### API Documentation
[Link API Documentation](https://documenter.getpostman.com/view/33387055/2sA3s3HWrb)

### What if there are thousands of products in the database?
- Saya menggunakan pagination untuk menghandle data yang banyak, dengan memakai offset dinamis dan limit yang statis

### What if many users are accessing your API at same time?
- Saya menggunakan redis agar saat orang banyak mengakses tidak terjadi loading yang lambat

### What if users perform stored xss and how to prevent it?
- Saya menggunakan middleware CORS dari framework echo dan saya custom sesuai yang dibutuhkan
