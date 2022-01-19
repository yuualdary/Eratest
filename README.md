# Eratest

1. Untuk structure code, saya pisahkan tiap-tiap kegunaan, misal model saya pisahkan jadi pada folder model hanya terdapat table dan coloumn untuk database, 
begitu juga pada folder product dimana proses crud terdapat pada folder tersebut , repository, service, dan format inputtan. Saya membuat struktur seperti ini agar terlihat rapi
dan mudah di maintenance


//
1. Untuk membuat product baru dapat menggunakan endpoint
localhost:8000/api/v1/product/ dengan pilihan POST pada postman

//berikut format jsonnya
{
   "productname":"Asus ROG",
   "productprice":522355,
   "productdescription":"Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to ",
   "productquantity":1
}

2. Untuk mensort sesuai kriteria saya menggunakan Query Param dimana menggunakan url untuk melakukan sorting

Key               Value
cheapest_product true // set true untuk melakukan sorting barang termurah
atoz             true // set true untuk melakukan sorting A-Z
ztoa             true // set true untuk melakukan sorting Z-A  note: saya membuat agar tidak dapat memilih A-Z dan Z-A secara bersamaan, jadi hanya saya set salah satu A-Z atau Z-A
lowest_product   110000 // set jumlah harga terendah yang diinginkan
newest_product   true // set true untuk melakukan sorting barang terbaru sesuai dengan kapan product dibuat
page             1//mau di page berapa, untuk page saya set limit = 5


localhost:8000/api/v1/product/?page=1 //endpoint untuk GET semua product pada page 1, untuk melakukan custom dapat melakukan query param pada postman sesuai dengan key dan value yang sudah saya buat
localhost:8000/api/v1/product/?cheapest_product=true&atoz=true&lowest_product=110000&page=1 // contoh endpoint untuk mencari product termurah, A-Z, harga terendah di 110000 pada page 1

Maaf jika format tulisan berantakan, untuk melihat tidak berantakan dapat membuka read.md dengan mode plain https://github.com/yuualdary/Eratest/blob/master/README.md?plain=1

