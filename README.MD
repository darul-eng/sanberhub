# Sanberhub Back-End Assessment


## Cara menggunakan

### Prasyarat

Sebelum menggunakan aplikasi ini pastikan anda telah memenuhi prasyarat sebagai berikut:

- Go 1.21 atau yang lebih baru
- Docker

### Konfigurasi

Untuk Menggunakan Aplikasi ini, dapat melalui dua cara oleh karena itu anda perlu melakukan konfigurasi terlebih dahulu. Berikut ini adalah cara melakukan konfigurasi:

#### Setup PostgreSQL
1. Download PostgreSQL image dari docker hub dengan perintah `docker pull postgres`
2. Jalankan container dari image sebelumnya dengan perintah `docker run --name postgres-container -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres`. Anda juga dapat menyesuaikan kredensialnya sesuai kebutuhan anda. `postgres-container` adalah nama container, `mysecretpassword` adalah database password , `5432` adalah database port, dan `postgres` yang di akhir string adalah database user
3. Pastikan container sudah running dengan perintah `docker ps` 
4. Masuk ke container PostgreSQL yang sudah berjalan dengan perintah `docker exec -it postgres-container bash`
5. Kemudian masuk ke PostgreSQL dengan perintah `psql -U postgres`.
6. Lalu Buat database dengan perintah `CREATE DATABASE db_sanberhub;`. `db_sanberhub` adalah nama database.

#### Setup Aplikasi
1. Clone project dari github `https://github.com/darul-eng/sanberhub.git`
2. Buat file .env dari duplikasi file .env-example. Sesuaikan semua nilai dengan yang dimasukkan sebelumnya pada saat setup PostgreSQL, seperti database password, database port, database user, dan nama database yang sudah anda buat.
3. Jalankan perintah `migrate -database "postgres://postgres:mysecretpassword@localhost:5432/db_sanberhub?sslmode=disable"
   -path migrations up` untuk melakukan migrasi database atau ganti perintah `up` menjadi `down` untuk melakukan rollback. Sesuikan krendensial yang ada di `.env`
4. Jalankan aplikasi dengan perintah `go run main.go`
5. Aplikasi Siap digunakan

### Endpoints


