## **Tugas-6-Network-Programming-Kelompok**

### **Anggota:**
        
        Fadillah Rizky R                1301164493
        Mazaya Z D                      1301154508
        Renaning Karutami Susilo        1301154466
        
### **HTTPS**

HTTPS, Hyper Text Transfer Protocol Secure, adalah ekstensi dari HTTP yang berguna untuk pengamanan komunikasi lewat internet. Data atau informasi yang dikomunikasikan di-enkripsi menggunakan TLS.

### **TLS**

TLS, Transport Layer Security, adalah versi yang lebih update dari SSL.

### **SSL**

SSL, Secure Sockets Layer, adalah standar untuk pengamanan komunikasi lewat internet. Data atau informasi yang sedang dikomunikasikan dari sebuah system ke system lain akan di-proteksi, dengan cara adalah mengacak informasi tersebut menggunakan algoritma enkripsi.

![https](https://user-images.githubusercontent.com/33456025/56866882-e83b6500-6a08-11e9-8546-29201a229589.PNG)
        
### **HTML**

HTML (HyperText Markup Language) adalah bahasa pemrograman standar yang digunakan untuk membuat sebuah halaman web, yang kemudian dapat diakses untuk menampilkan berbagai informasi di dalam sebuah penjelajah web Internet (Browser).
Fungsi html yang pertama adalah untuk membuat suatu halaman website yang dapat dibaca dan dipahami oleh pengguna dengan lebih mudah. Lalu, menandai teks pada suatu laman.

### **HTTP**

HTTP (Hypertext Transfer Protocol) merupakan istilah yang diberikan pada sebuah protokol dan dipergunakan untuk mengirimkan dokumen dari WWW (World Wide Web). HTTP dapat pula diartikan sebagai protokol jaringan untuk pendistribusian sistem informasi hypermedia secara kolaboratif. Fungsi utama dari protokol HTTP sebenarnya cukup sederhana, yaitu untuk mengkomunikasikan satu komputer dengan lainnya.Selain itu, HTTP juga berfungsi untuk menentukan bagaimana sebuah data atau pesan dapat ditransmisikan maupun diformat menjadi bentuk yang dapat merespon browser untuk menampilkan data-data tersebut.

Cara Kerja HTTP:
1. Pertama-tama, komputer klien (HTTP klien) membuat sambungan, lalu mengirimkan permintaan dokumen ke web server.
2. HTTP server kemudian memproses permintaan klien, sementara itu, HTTP klien menunggu respon dari server (berupa HTML) tersebut.
3. Web server merespon permintaan dengan kode status data, lalu barulah menutup sambungan ketika telah selesai memproses    permintaan. 

**Soal**

Buatlah perancangan aplikasi Web Server yang dapat melakukan serve koneksi HTTPS menggunakan diagram FSM serta jelaskan cara kerjanya!

**Jawaban:**

![secure web design](https://user-images.githubusercontent.com/33456025/56866888-f5f0ea80-6a08-11e9-83c3-29e9c25b2694.jpg)


1. Browser mengunjungi suatu halaman website yang menggunakan SSL.
2. Secara bersamaan browser komputer client meminta komputer server memberikan identitas SSL website untuk memvalidasi status keamanan website tersebut.
3. Komputer server mengirim duplikat dari sertifikat SSL yang dimilikinya.
4. Browser komputer client memeriksa kebenaran sertifikat SSL yang dimiliki website tersebut. Jika benar, browser akan mengirim permintaan untuk melakukan koneksi terenskripsi kepada server.
5. Komputer server mengirim anak kunci berupa tanda tangan digital untuk melakukan melakukan transfer data dengan komputer client dengan identitas khusus (Private Key) yang hanya bisa digunakan komputer client tersebut. Anak kunci digital ini digunakan oleh browser membuka informasi yang diminta dari website tersebut serta mengamankan data yang dikirim ke website tersebut, karena setiap data yang dialirkan akan dibungkus oleh sistem SSL dengan identitas khusus.
6. Koneksi terenskripsi dapat dilakukan dan browser komputer client menampilkan status “Connection is secure”. Untuk memberi tahu pengguna bahwa website tersebut aman untuk melindungi datanya.

**Soal**

Implementasikan aplikasi secure web server dari design yang sudah anda buat, aplikasi harus mempunyai config file untuk melakukan konfigurasi aplikasi!

**Jawaban:**

gagal :')

![secure](https://user-images.githubusercontent.com/33456025/56866896-0e610500-6a09-11e9-8a0e-9fc0e8fc8c01.PNG)

![https 2](https://user-images.githubusercontent.com/33456025/56866911-29cc1000-6a09-11e9-8ef1-dbc3e50621d8.PNG)

![https 3](https://user-images.githubusercontent.com/33456025/56866900-1a4cc700-6a09-11e9-91c7-875d1f660834.PNG)

maap yaa pak :((




