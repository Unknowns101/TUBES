**Deskripsi Aplikasi**:

Aplikasi ini merupakan program berbasis **CLI (Command Line Interface)** yang dibuat menggunakan bahasa **Go (Golang)**. Aplikasi ini berfungsi untuk **mengelola data tim olahraga**, termasuk informasi pemain, statistik pertandingan, dan peringkat tim.

Setiap tim memiliki:

* Nama tim
* Jumlah menang dan kalah
* Total pertandingan
* Total poin
* Daftar 5 pemain, masing-masing memiliki nama dan poin

Aplikasi ini memungkinkan pengguna untuk menambahkan, memperbarui, menghapus, menampilkan, mencari, dan mengurutkan data tim secara efisien.

**Cara Mengoperasikan Aplikasi**:

1. **Menjalankan Program**

   * Simpan file `.go` ke dalam direktori kerja.
   * Jalankan dengan perintah:
     go run tubess.go

2. **Menggunakan Menu**
   Saat program berjalan, pengguna akan disuguhkan menu utama:

   ```
   ===== MENU =====
   1. Tambah Tim
   2. Update Tim
   3. Hapus Tim
   4. Tampilkan Semua Tim
   5. Cari Tim (berdasarkan Nama)
   6. Cari Tim (berdasarkan Poin)
   7. Urutkan Tim berdasarkan Nama
   8. Urutkan Tim berdasarkan Poin
   9. Keluar
   ```

   * Pilih menu dengan mengetikkan angka dan tekan **Enter**.
   * Ikuti petunjuk untuk memasukkan data sesuai kebutuhan.

3. **Penjelasan Operasi**

   * **Tambah Tim**: Input nama tim dan nama 5 pemain. Statistik diinisialisasi ke nol.
   * **Update Tim**: Update jumlah menang/kalah dan poin setiap pemain.
   * **Hapus Tim**: Masukkan nama tim yang ingin dihapus.
   * **Tampilkan Semua Tim**: Menampilkan seluruh data tim dan statistik pemain.
   * **Cari Tim**: Binary search berdasarkan nama atau poin (dengan pengurutan sebelumnya).
   * **Urutkan Tim**: Bisa berdasarkan nama (ascending) atau poin (descending).
