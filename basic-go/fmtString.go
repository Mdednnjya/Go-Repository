package main 
import "fmt"

func main() {
    // mendefiniskan variabel
    var name string = "Danan"
    var age = 18
    var alamat = "Klojen, Malang"
    
    // print
    fmt.Println("Result:")
    fmt.Printf("Nama: %v \n", name)
    fmt.Printf("Umur: %d \n", age)
    fmt.Printf("Alamat: %v \n", alamat)


    // \"%s\" agar tanda petik terbaca sebagai String
    fmt.Println(" ")
    fmt.Printf("kata \"%s\" meiliki %d huruf", name, len(name))

}