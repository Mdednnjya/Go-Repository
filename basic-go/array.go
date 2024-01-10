package main 
import "fmt"

func main() {

	// --array--
	var id = [3]string{"121", "131", "214"}

	fmt.Println(id, len(id))

	// --slice--
	fmt.Printf("")
	var num = []int{10, 20, 30, 40, 50, 60 ,70 ,80, 90, 100}
	rangeOne := num[0:3]
	rangeTwo := num[2:5]
	rangeThree := num[4:7]
	rangeFour := num[7:10]
	fmt.Printf("Sebuah array dengan nilai: %d \n", num)
	fmt.Printf("Sebuah array dengan total: %d buah \n \n", len(num))

	
	fmt.Println("Yang sudah dibagi menjadi beebrapa bagian:")
	fmt.Printf("bagian 1: %v \n", rangeOne)
	fmt.Printf("bagian 2: %d \n", rangeTwo)
	fmt.Printf("bagian 3: %d \n", rangeThree)
	fmt.Printf("bagian 4: %d \n \n", rangeFour)

	num = append(num, 110)
	fmt.Println("nilai pada array telah ditambahkan")

	fmt.Printf("nilai array sekarang adalah: %v\n", num)
	fmt.Printf("dengan panjang %d \n", len(num))



}