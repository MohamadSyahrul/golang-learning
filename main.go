package main

import (
	"fmt"
	"strconv"
)

func main() {
	// variabel i dideklarasikan dengan nilai 21
	i := 21
	// Format %v digunakan untuk menampilkan nilai variabel.
	fmt.Printf("nilai i : %v \n", i)
	// menampilkan nilai i dengan tipe data apa yang dimiliki oleh variabel i, yaitu int.
	fmt.Printf("tipe data variabel i : %T \n", i)
	// untuk menampilkan tanda persen. Karena tanda persen merupakan karakter escape,
	//  maka perlu menggunakan tanda persen ganda (%%) untuk menampilkan karakter tersebut.
	fmt.Printf("tanda persen : %% \n")

	// strconv.FormatInt untuk mengonversi nilai i menjadi bilangan biner
	binary := strconv.FormatInt(int64(i), 2)
	fmt.Printf("nilai binary nilai i : %v \n", binary)

	// variabel j dideklarasikan dengan nilai boolean true.
	j := true
	// untuk menampilkan nilai boolean j dengan format %v.
	fmt.Printf("nilai boolean j: %v\n", j)
	// untuk menampilkan nilai boolean j dengan format %t. Format %t digunakan untuk menampilkan nilai boolean.
	fmt.Printf("nilai boolean j : %t \n", j)
	//  untuk menampilkan karakter Я dalam bentuk unicode dengan format %c. Format %c digunakan untuk menampilkan karakter.
	fmt.Printf("nilai unicode russia : %c \n", 'Я')
	// untuk menampilkan nilai i dalam basis 10 dengan format %d. Format %d digunakan untuk menampilkan bilangan bulat dalam basis 10
	fmt.Printf("nilai base 10 : %d \n", i)
	// untuk menampilkan nilai i dalam basis 8 dengan format %o. Format %o digunakan untuk menampilkan bilangan bulat dalam basis 8.
	fmt.Printf("nilai base 8 : %o \n", i)

	// variabel hexLower dideklarasikan dengan nilai 0xf dalam notasi heksadesimal merepresentasikan angka 15 dalam sistem desimal.
	hexLower := 0xf
	// Format %x digunakan untuk menampilkan bilangan bulat dalam basis 16 dengan huruf kecil.
	fmt.Printf("nilai base 16 : %x \n", hexLower)

	// variabel hexUpper dideklarasikan dengan nilai 0xF13 setara dengan 3.539
	hexUpper := 0xF13
	// Format %X digunakan untuk menampilkan bilangan bulat dalam basis 16 dengan huruf besar.
	fmt.Printf("nilai base 16 : %X \n", hexUpper)

	// untuk menampilkan karakter Я dalam bentuk unicode dengan format %U. Format %U digunakan untuk menampilkan karakter dalam bentuk unicode.
	fmt.Printf("unicode karakter Я : %U \n", 'Я')

	// variabel k dideklarasikan dengan nilai float64 123.456.
	var k float64 = 123.456
	// untuk menampilkan nilai float k dengan format %f. Format %f digunakan untuk menampilkan nilai float dengan tanda desimal
	fmt.Printf("float : %f \n", k)
	// untuk menampilkan nilai float k dalam notasi scientific dengan format %e. Format %e digunakan untuk menampilkan nilai float dalam notasi scientific
	fmt.Printf("float scientific : %e \n", k)

}
