<h1 align="center">Selamat datang di xray 👋</h1>
<p>
  <img src="https://img.shields.io/github/release/chaitin/xray.svg" />
  <img src="https://img.shields.io/github/release-date/chaitin/xray.svg?color=blue&label=update" />
  <img src="https://img.shields.io/badge/go report-A+-brightgreen.svg" />
  <a href="https://docs.xray.cool/#/">
    <img alt="Dokumentasi" src="https://img.shields.io/badge/dokumentasi-ya-brightgreen.svg" target="_blank" />
  </a>
</p>

[**Versi Bahasa Inggris**](./README_EN.md)

> Sebuah alat evaluasi keamanan yang kuat 

## ✨ Demo

![](https://docs.xray.cool/assets/term.svg)

🏠[Dokumentasi Pengguna](https://docs.xray.cool)  
⬇️[Alamat Unduh untuk Pengguna di Indonesia](https://stack.chaitin.com/tool/detail?id=1)  
⬇️[Alamat Unduh di GitHub](https://github.com/chaitin/xray/releases)

> Catatan: xray tidak bersifat open source, cukup unduh file biner yang telah dibangun. Repositori ini terutama berisi poc yang dikontribusikan oleh komunitas, dan setiap rilis xray akan di-packing secara otomatis.

## xray 2.0

Untuk mengatasi masalah kompleksitas dan keterlaluan xray 1.0 dalam proses peningkatan fungsionalitas, kami memperkenalkan xray 2.0.

Versi baru ini bertujuan untuk meningkatkan kelancaran penggunaan fitur, mengurangi ambang batas penggunaan, dan membantu lebih banyak praktisi keamanan untuk mendapatkan pengalaman yang lebih baik dengan mode yang lebih efisien. xray 2.0 akan mengintegrasikan serangkaian alat keamanan baru, membentuk seperangkat alat keamanan yang komprehensif.

**Alat pertama dari seri xray 2.0, yaitu xpoc, sudah diluncurkan. Selamat mencoba!**

- [**xpoc**](https://github.com/chaitin/xpoc)

## 🚀 Penggunaan Cepat

**Sebelum menggunakan, harap baca dan setujui [Lisensi](https://github.com/chaitin/xray/blob/master/LICENSE.md) yang ada dalam file ini. Jika tidak, jangan instal dan gunakan alat ini.**

1. Gunakan *crawler* dasar untuk mengambil tautan dan lakukan pemindaian kerentanannya.

    ```bash
    xray webscan --basic-crawler http://contoh.com --html-output vuln.html
    ```

2. Gunakan *HTTP proxy* untuk pemindaian pasif.

    ```bash
    xray webscan --listen 127.0.0.1:7777 --html-output proxy.html
    ```
   Atur *HTTP proxy* browser ke `http://127.0.0.1:7777` untuk otomatis menganalisis lalu lintas proxy dan melakukan pemindaian.

   > Untuk pemindaian lalu lintas HTTPS, baca bagian "Menangkap lalu lintas HTTPS" di bawah.

3. Lakukan pemindaian hanya untuk satu URL tanpa menggunakan *crawler*.

    ```bash
    xray webscan --url http://contoh.com/?a=b --html-output single-url.html
    ```

4. Tentukan plugin yang akan dijalankan secara manual.

    Secara default, semua plugin bawaan akan diaktifkan. Anda dapat menggunakan perintah berikut untuk menentukan plugin yang akan diaktifkan.

    ```bash
    xray webscan --plugins cmd-injection,sqldet --url http://contoh.com
    xray webscan --plugins cmd-injection,sqldet --listen 127.0.0.1:7777
    ```
      
5. Tentukan output plugin.

    Anda dapat menentukan tempat keluaran informasi kerentanan pemindaian ke file tertentu.

    ```bash
    xray webscan --url http://contoh.com/?a=b \
    --text-output result.txt --json-output result.json --html-output report.html
    ```

    [Contoh laporan](https://docs.xray.cool/assets/report_example.html)

Untuk penggunaan lainnya, harap baca dokumen: https://docs.xray.cool

## 🪟 Modul Deteksi

Modul deteksi baru akan terus ditambahkan.

| Nama               | Kunci            | Versi | Penjelasan |
|-------------------|------------------|------|--------------------------------------------------------------------|
| Deteksi Kerentanan XSS | `xss`            | Komunitas | Mendeteksi kerentanan XSS menggunakan analisis semantik |
| Deteksi Kerentanan SQL | `sqldet`         | Komunitas | Mendukung jenis kerentanan injeksi kesalahan, injeksi boolean, dan injeksi waktu buta |
| Deteksi Injeksi Perintah/Kode | `cmd-injection`  | Komunitas | Mendukung injeksi perintah shell, eksekusi kode PHP, dan injeksi template lainnya |
| Pemindaian Direktori | `dirscan`        | Komunitas | Mendeteksi file cadangan, file sementara, halaman debug, file konfigurasi, dll. |
| Deteksi Traversal Path | `path-traversal` | Komunitas | Mendukung platform dan encoding umum |
| Deteksi Injeksi Entitas XML | `xxe`            | Komunitas | Mendukung deteksi entitas XML dengan respon langsung atau entitas terbalik |
| Manajemen poc | `phantasm`       | Komunitas | Mengandung beberapa poc umum dan dapat disesuaikan oleh pengguna. Dokumen: [POC](https://docs.xray.cool/#/guide/poc) |
| Deteksi Unggah File | `upload`         | Komunitas | Mendukung bahasa backend umum |
| Deteksi Kata Sandi Lemah | `brute-force`    | Komunitas | Mendukung deteksi kata sandi HTTP dasar dan kata sandi formulir sederhana, kamus nama pengguna dan kata sandi umum terintegrasi |
| Deteksi JSONP | `jsonp`          | Komunitas | Mendeteksi antarmuka JSONP yang mengandung informasi sensitif yang dapat dibaca melalui lintas domain |
| Deteksi SSRF | `ssrf`           | Komunitas | Modul deteksi SSRF, mendukung teknik bypass umum dan deteksi platform terbalik |
| Pemeriksaan Dasar | `baseline`       | Komunitas | Mendeteksi versi SSL rendah, header HTTP yang hilang atau salah, dll. |
| Deteksi Redireksi Bebas | `redirect`       | Komunitas | Mendukung meta HTML redirect, redirect 30x, dll. |
| CRLF Injection | `crlf-injection` | Komunitas | Mendeteksi penyisipan header HTTP, mendukung parameter di posisi query, body, dll. |
| Deteksi Kerentanan XStream | `xstream`        | Komunitas | Mendeteksi kerentanan seri XStream |
| Deteksi Kerentanan Seri Struts2 | `struts`         | Advanced | Mendeteksi apakah situs target memiliki kerentanan seri Struts2, termasuk s2-016, s2-032, s2-045, s2-059, s2-061, dll. |
| Deteksi Kerentanan Seri Thinkphp | `thinkphp`       | Advanced | Mendeteksi kerentanan pada situs web yang dikembangkan dengan ThinkPHP |
| Deteksi Kerentanan Seri Shiro | `shiro`          | Advanced | Mendeteksi kerentanan seri Shiro deserialization |
| Deteksi Kerentanan Seri Fastjson | `fastjson`       | Advanced | Mendeteksi kerentanan seri Fastjson |

## ⚡️ Penggunaan Lanjutan

Untuk penggunaan lanjutan, silakan lihat https://docs.xray.cool/ menggunakan.

 - Mengubah file konfigurasi
 - Menangkap lalu lintas HTTPS
 - Mengubah konfigurasi pengiriman HTTP
 - Penggunaan platform terbalik
 - ...

## 😘 Kontribusi POC

Kemajuan xray tidak terlepas dari dukungan para mastah. Dengan semangat saling membantu, agar kita bisa berkembang bersama, xray membuka saluran "Penerimaan POC"! Di sini Anda akan mendapatkan:

### Prosedur Pengajuan

1. Kontributor mengajukan dengan cara membuat PR ke repositori komunitas xray di github, lokasi pengajuan POC ada di sini: https://github.com/chaitin/xray/tree/master/pocs, lokasi script pengenalan sidik jari ada di sini: https://github.com/chaitin/xray/tree/master/fingerprints
2. Di dalam PR, isi informasi POC sesuai dengan template Pull Request.
3. PR akan diverifikasi secara internal untuk menentukan apakah akan dimasukkan ke repositori.
4. Harap dicatat, jika ingin mendapatkan hadiah untuk POC yang diajukan, Anda perlu mengirimkan POC Anda ke CT stack untuk mendapatkan hadiah.

### Hadiah Melimpah

- Kontributor POC akan mendapatkan **hadiah berupa koin emas yang melimpah**, memberikan rasa pencapaian.
- Area penukaran hadiah yang **banyak dan beragam**, dengan lebih dari 50 pilihan hadiah merchandise yang bisa Anda pilih.
- Penawaran teratur dari kartu JD (Jingdong) untuk pertukaran, mendekatkan diri pada **kebebasan finansial**.
- Kesempatan masuk ke dalam komunitas inti, mengambil tugas khusus, dan mendapatkan **reward bounty yang tinggi**.

### Tutorial Lengkap

- **Panduan lengkap dalam membuat dan menguji POC**, membantu Anda untuk cepat memahami, mengurangi kesalahan.
  
### Belajar dan Berbagi

- Kesempatan **belajar langsung dan berbagi dengan kontributor, pengembang** secara langsung, meningkatkan keterampilan secara keseluruhan.
- Kesempatan untuk mendapatkan pekerjaan **tanpa tes tulis**, membawa ke arah pekerjaan yang baik.

Jika Anda sudah berhasil menyumbangkan POC tetapi belum bergabung dengan grup, tambahkan akun servis pelanggan kami di WeChat:

<img src="https://docs.xray.cool/assets/customer_service.png?cache=_none" height="200px">

Sediakan ID pendaftaran platform untuk diverifikasi dan setelah itu Anda dapat bergabung!

Lihat juga: https://docs.xray.cool/#/guide/contribute

## 🔧 Ekosistem Sekitar

### Alat Bantu Penulisan POC

Alat ini dapat membantu dalam membuat POC dan versi online mendukung **pemeriksaan duplikat POC**, sementara versi lokal mendukung verifikasi pengiriman langsung.

#### Versi Online
- [**Laboratorium Aturan**](https://poc.xray.cool)
- Versi online mendukung **pemeriksaan duplikat POC**

#### Versi Lokal
- [**gamma-gui**](https://github.com/zeoxisca/gamma-gui)

### Alat Bantu GUI xray

Alat ini hanya bungkusan baris perintah sederhana dan bukan pemanggilan langsung ke metode. Dalam perencanaan xray, di masa depan akan ada alat GUI XrayPro yang sebenarnya dan lengkap. Nantikan itu.

- [**super-xray**](https://github.com/4ra1n/super-xray)

## 📝 Forum Diskusi

Untuk laporan palsu atau laporan kebutuhan lainnya, harap baca https://docs.xray.cool/#/guide/feedback terlebih dahulu.

Jika ada masalah, silakan buat isu di GitHub, atau bergabung di grup diskusi di bawah ini:

1. Isu GitHub: https://github.com/chaitin/xray/issues
2. Akun WeChat: Pindai kode QR di bawah, dan ikuti kami.

<img src="https://docs.xray.cool/assets/wechat.jpg?cache=_none" height="200px">

3. Grup WeChat: Tambahkan akun WeChat dan klik "Hubungi Kami" -> "Bergabung dengan Grup", lalu pindai kode QR untuk bergabung.
4. Grup QQ: 717365081
