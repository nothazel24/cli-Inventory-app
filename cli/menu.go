package cli

import (
	"fmt"
	"inventory-app/service"
)

func Start(s service.ItemService) {
	for {
		fmt.Println("Inventory app (silahkan pilih opsi)")
		fmt.Println("1. Tambah barang")
		fmt.Println("2. Lihat barang")
		fmt.Println("3. Cari barang")
		fmt.Println("4. Update stok barang")
		fmt.Println("5. Hapus barang")
		fmt.Println("6. Keluar dari program")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var nama string
			var harga, stok int

			fmt.Print("Nama: ")
			fmt.Scanln(&nama)
			fmt.Print("Harga: ")
			fmt.Scanln(&harga)
			fmt.Print("Stok: ")
			fmt.Scanln(&stok)

			err := s.AddItem(nama, harga, stok)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			items, _ := s.GetAllItem()
			for _, item := range items {
				fmt.Println(item.Nama, item.Stok, item.Harga)
			}
		case 3:
			var nama string
			var qty int

			fmt.Print("Nama: ")
			fmt.Scanln(&nama)
			fmt.Print("Qty: ")
			fmt.Scanln(&qty)

			err := s.UpdateStock(nama, qty)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 4:
			var nama string
			fmt.Print("Nama: ")
			fmt.Scanln(&nama)

			err := s.DeleteItem(nama)
			if err != nil {
				fmt.Println(err)
			}
		case 5:
			var nama string
			fmt.Print("Nama: ")
			fmt.Scanln(&nama)

			item, err := s.FindItem(nama)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(item)
			}
		case 6:
			return
		default:
			fmt.Println("Perintah tidak dikenali")
		}
	}
}
