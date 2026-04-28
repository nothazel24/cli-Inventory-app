package cli

import "fmt"

func ShowMenu() {
	fmt.Println("Inventory app (silahkan pilih opsi)")
	fmt.Println("1. Tambah barang")
	fmt.Println("2. Lihat barang")
	fmt.Println("3. Cari barang")
	fmt.Println("4. Update stok barang")
	fmt.Println("5. Hapus barang")
}
