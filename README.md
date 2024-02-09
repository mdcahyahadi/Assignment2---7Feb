Order adalah sebuah objek yang terdiri atas Item <br><br>
Endpoint: <br>
- Create: Setiap pembuatan order tidak wajib menyertakan item, tetapi item dapat ditambahan ke dalam Order.
- Delete: Setiap kali order dihapus, maka item-item didalamnya akan ikut terhapus, tetapi jika item dihapus, maka order tetap ada.
- Update: Update dilakukan pada order dan item tidak mempengaruhi satu sama lain, update dilakukan menggunakan ID rownya.
- Get All: Setiap order akan diurutkan berdasarkan ID dan didalamnya akan disertakan itemnya. Sedangkan untuk item wajib mengisi order ID nya karena get all item dianggap sebagai detail order sehingga ditampilkan per order.
- Get By ID: Pengambilan data order dan item berdasarkan id rownya.
