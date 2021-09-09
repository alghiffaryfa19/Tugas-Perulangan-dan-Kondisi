package main

import "fmt"

func main() {
	var a, batas int

    fmt.Printf("Pilih:\n")
	fmt.Printf("1. Segitiga Tidak Terbalik:\n")
	fmt.Printf("2. Segitiga Terbalik :\n")
	fmt.Scanf("%d",&a)
	fmt.Print("Batas : ")
	fmt.Scan(&batas)

	switch a{
		case 1:
			i := 1
			for i <= batas {
				j := 1
				for j <= batas-i {
					fmt.Print(" ")
					j+=1
				}
				k := 0
				for k != (2*i - 1) {
					fmt.Print("*")
					k+=1
				}
				fmt.Println()
				i+=1
			}

		case 2:
			i := batas
			for i > 0 {
				j := 0
				for j <= batas-i {
					fmt.Print(" ")
					j+=1
				}

				k := 0
				for k < (2*i - 1) {
					fmt.Print("*")
					k+=1
				}

				fmt.Println()
				i-=1
			}
		default:fmt.Println("Tidak ada")
	}
}