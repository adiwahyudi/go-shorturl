# go-shorturl
Projek ini bertujuan untuk mempersingkat suatu *URL* panjang menjadi lebih ringkas agar lebih mudah untuk disebarkan.

## Masalah
Ketika membagikan suatu *URL*, sering kali *URL* tersebut terlihat panjang, kompleks dan sulit diingat.

## Solusi   
Projek ini akan mengambil *URL* asli dan mengubahnya menjadi lebih ringkas sepanjang 7 karakter acak. Adapun pada projek ini tahapan untuk mempersingkat *URL* sebagai berikut:
1. Server menerima permintaan *URL* dari client.
2. Permintaan *URL* tersebut akan di *hash* dengan cara MD5.
3. Dari hasil hash tersebut akan di *encode* dengan cara base62.
4. Hasil dari encode base62 tersebut akan diambil 7 karakter dari depan sebagai hasil ringkasan URL.
5. Jika 7 karakter tersebut sudah ada di dalam database, maka dari hasil encode base62 akan diambil 7 karakter lagi maju selangkah dari karakter sebelumnya.
6. Dan jika masih sama lagi, akan berulang hingga ditemukan 7 karakter yang tidak ada di database atau berakhir pada 7 karakter terakhir pada hasil encode.

**Alur**:
- Server menerima request dari client berupa *URL* original yaitu `https://google.com`.
- URL akan di *hash* dengan  MD5 dan menghasilkan hasil berupa tipe data array 16 *byte* `[81 53 242 98 178 25 94 34 92 30 223 6 146 77 66 4]`.
- Selanjutnya akan di *encode* ke base62 dan hasilnya berikut `f0SA6fJn06dw34t9zrnZmJ`.
- Hasil tersebut akan diambil sepanjang 7 karakter yaitu `f0SA6fJ`.
- Sebelum pada database, akan dilakukan pengecekan 7 karakter tersebut belum ada dalam database.
- Jika ada, akan mengambil 7 karakter berikutnya sebagai contoh `0SA6fJn`, dan seterusnya.
- Jika tidak ada, akan disimpan di dalam database.
## Penggunaan aplikasi
**Menjalankan Aplikasi**
1. *Clone* repository dan masuk ke directori lokal:
```bash
    git clone https://github.com/adiwahyudi/go-shorturl.git
    cd go-shorturl
```
2. Instal dependensi 
```go
    go get
```
3. Jalankan aplikasi
```go
    go run ./
```

**Contoh Shorten URL**

- **Request:** /short-url body: {url}
```curl
    curl -d '{"url": "https://facebook.com"}' http://localhost:8080/short-url
```

- **Response:**
```json
    {"url":"https://facebook.com","short_url":"http://localhost:8080/0END2iy"}
```

**Contoh Get Long URL**
- **Request:** /long-url body: {short_url}
```curl
    curl -d '{"short_url": "http://localhost:8080/0END2iy"}' http://localhost:8080/long-url
```

- **Response:**
```json
    {"url":"https://facebook.com","short_url":"http://localhost:8080/0END2iy"}
```

- Validation request url: `format: https or http`
```
    Valid:
    "url" : "https://facebook.com",
    "url" : "http://facebook.com",

    Invalid:
    "url" : "facebook.com",
    "url" : "https://facebook",
```