# Update Implementasi (Soal Pertama)

Berikut adalah perubahan yang telah dilakukan untuk memenuhi kebutuhan soal tes:

## 1. Fitur Baru
### a. Tabel Mata Kuliah (Subject) & Relasi
- **Model**: Menambahkan `model/subject.go` dengan field `ID`, `Name`, dan `Credits`.
- **Relasi**: Menambahkan relasi **Many-to-Many** antara `Student` dan `Subject` melalui tabel pivot `student_subjects`.
- **Migrasi**: Update `database/migrate.go` untuk auto-migrate tabel subject.

### b. API Endpoints
Menambahkan endpoint baru di `route/api.go` dan handler di `handler/subject.go`:
- **CRUD Subject**:
  - `GET /api/v0/subject` - List semua mata kuliah
  - `POST /api/v0/subject` - Tambah mata kuliah baru
  - `GET /api/v0/subject/:uuid_subject` - Detail mata kuliah
  - `PUT /api/v0/subject/:uuid_subject` - Update mata kuliah
  - `DELETE /api/v0/subject/:uuid_subject` - Hapus mata kuliah
- **Integrasi**:
  - `POST /api/v0/student/subject` - Menambahkan mata kuliah ke mahasiswa (Attach).

## 2. Keamanan & Enkripsi
- **Kolom Email**: Menambahkan field `Email` pada struct `Student`.
- **Enkripsi AES-256**: Implementasi enkripsi otomatis pada kolom email menggunakan **GORM Hooks** (`BeforeCreate`, `BeforeUpdate`, `AfterFind`).
  - Email akan tersimpan dalam bentuk *ciphertext* di database.
  - Email akan otomatis didekripsi menjadi *plaintext* saat di-query melalui aplikasi.
- **Utils**: Helper enkripsi/dekripsi diletakkan di `utils/crypto.go`.

## 3. Optimasi & Validasi (Nilai Tambah)
- **Validasi Input**: Menambahkan tag `binding:"required"` pada struct request (`request/student.go`, `request/subject.go`) untuk mencegah input kosong.
- **Docker Alpine**: Mengupdate `Dockerfile` dan `docker-compose.yml` untuk menggunakan image berbasis **Alpine Linux** (`alpine:latest`, `postgres:alpine`). Ini mengurangi ukuran image secara signifikan.

## Cara Menjalankan Perubahan
1. Tambahkan `ENCRYPTION_KEY` (32 karakter) pada file `.env`.
2. Build ulang dan jalankan container:
   ```shell
   docker-compose up --build
   ```