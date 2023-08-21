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
// 	FixingCapital("terimakasih")
// 	// capitalizeWords("terimakasih")
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

// func FixingCapital(input string) {
// 	coba := strings.Split(input, " ")
// 	// fmt.Println("hello", coba)
// 	// hasile := strings.ToUpper(coba[0])
// 	for _, word := range coba {
// 		first := strings.ToUpper(string(word[0])) + strings.ToLower(word[1:6])
// 		last := strings.ToUpper(string(word[1]))
// 		fmt.Println("hasile", result)
// 	}

// }

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
