package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// func main() {

// 	// Vocale("senapahle")
// 	// JsonInput("budi tono 21")
// 	// FixingCapital("terimakasih")
// 	// capitalizeWords("terimakasih")
// 	SunlightTree([]int{1, 2, 8, 6, 7, 9})
// 	sunlitTrees([]int{1, 2, 8, 6, 7, 20, 9})
// }

// type Manager struct {
// 	FirstName string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// 	Age       string `json:"age"`
// }

// func JsonInput(input string) {
// 	hasil := strings.Split(input, " ")
// 	fmt.Println(hasil)

// 	coba, err := json.Marshal(Manager{
// 		FirstName: hasil[0],
// 		Lastname:  hasil[1],
// 		Age:       hasil[2],
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(coba))
// }

// func FixingTeks(input string) {
// 	hasile := ""
// 	for i := 0; i < len(input); i++ {
// 		_, err := strconv.Atoi(string(input[i]))
// 		if err != nil {
// 			hasile += string(input[i])
// 		}
// 	} // membaikkan hasile
// 	balik := ""
// 	for i := len(hasile) - 1; i >= 0; i-- {
// 		balik += string(hasile[i])
// 	}
// 	fmt.Println(balik)
// }

// // func FixingTeks(input string) {
// // 	coba := strings.Split(input, "")
// // 	arey := []string{}
// // 	hasile := ""
// // 	for i := 0; i < len(coba); i++ {
// // 		_, err := strconv.Atoi(coba[i])
// // 		if err != nil {
// // 			arey = append(arey, coba[i])
// // 		}

// // 	}
// // 	cobasin := sort.Reverse(sort.StringSlice(arey))

// // 	fmt.Println(cobasin, "hello")
// // 	fmt.Println(hasile, "hasile")
// // }

// func cleanAndReverse(s string) string { // loop over all digits from 0 to 9
// 	for i := 0; i < 10; i++ { // convert the digit to a string
// 		digit := strconv.Itoa(i) // replace the digit with an empty string
// 		s = strings.ReplaceAll(s, digit, "")
// 	} // split the string into a slice of characters
// 	sSlice := strings.Split(s, "")
// 	// sort.Strings(sSlice)                   // reverse the order of the slice
// 	sort.Reverse(sort.StringSlice(sSlice)) // join the slice into a string
// 	s = strings.Join(sSlice, "")
// 	fmt.Println(s, "finale")

// 	return s
// }

// // func FixingCapital(input string) {
// // 	coba := strings.Split(input, " ")

// // 	for _, word := range coba {
// // 		first := strings.ToUpper(string(word[0])) + strings.ToLower(word[1:6])
// // 		last := strings.ToUpper(string(word[1]))
// // 		fmt.Println("hasile", result)
// // 	}

// // }

// func capitalizeWords(s string) string {
// 	// Variabel untuk menyimpan hasil output
// 	var result string

// 	// Membuat slice of string yang berisi kata-kata dalam string input
// 	words := strings.Split(s, " ")

// 	// Loop for sepanjang panjang slice
// 	for _, word := range words {
// 		// Mengubah huruf pertama menjadi huruf kapital dan menggabungkan dengan sisa kata
// 		word = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])

// 		// Menambahkan kata yang sudah diubah ke variabel result
// 		result += word + " "
// 	}

// 	// Mengembalikan variabel result sebagai nilai balik fungsi
// 	fmt.Println("heyy result", result)
// 	return result
// }

// func SunlightTree(input []int) {
// 	sum := 0
// 	hasil := []int{}
// 	max := 0

// 	for i := 0; i < len(input)-1; i++ {
// 		// if i == 0 {
// 		// 	sum += 1
// 		// }
// 		if input[i] < input[i+1] {
// 			// max = input[i+1]
// 			hasil = append(hasil, input[i+1])
// 		}

// 	}
// 	// 1, 2, 8, 6, 7, 9
// 	fmt.Println(sum, "sum")
// 	fmt.Println(hasil, "hasil")
// 	fmt.Println(max, "max")
// }

// func sunlitTrees(meter []int) []int {
// 	// Variabel untuk menyimpan tinggi maksimum pohon yang terkena sinar matahari
// 	var max int

// 	// Variabel untuk menyimpan indeks-indeks pohon yang terkena sinar matahari
// 	var result []int

// 	// Loop for sepanjang panjang array meter
// 	for i, m := range meter {
// 		// Memeriksa apakah tinggi pohon lebih besar dari variabel max atau tidak
// 		fmt.Println(i, "jello i", m, " hello m")
// 		fmt.Println("hoy max", max)
// 		if m > max {
// 			// Menambahkan indeks pohon ke variabel result
// 			result = append(result, m)
// 			// Mengubah nilai variabel max menjadi tinggi pohon
// 			max = m
// 		}
// 	}

// 	// Mengembalikan variabel result sebagai nilai balik fungsi
// 	fmt.Println("finale", result)
// 	return result

// }
