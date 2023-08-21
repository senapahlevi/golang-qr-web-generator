package m_test

// import (
// 	"fmt"
// 	"strings"
// )

// //The quick brown fox jumps over the lazy dog.

// func coba() {
// 	// text := "darwin city"
// 	// arr := []string{text}
// 	// arrs := arr[0]
// 	// urut := strings.Split(arrs, "")

// 	// sort.Sort(sort.StringSlice(urut))
// 	// fmt.Println(urut)
// 	// var result map[string]int = make(map[string]int)
// 	// for i := 0; i < len(arr); i++ {
// 	// 	// result = map[string]int{
// 	// 	// 	urut[i]: strings.Count(text, urut[i]),
// 	// 	// }
// 	// 	result = {"sena": strings.Count(text, urut[i])}

// 	// 	fmt.Println(result)
// 	// }
// 	// fmt.Println("Masukkan teks:")
// 	// var teks string
// 	// fmt.Scanln(&teks)

// 	// // Menghitung jumlah huruf yang muncul dalam teks
// 	// teks = strings.ToLower(teks)     // Mengubah semua huruf menjadi huruf kecil
// 	// hitungan := make(map[string]int) // Membuat map untuk menyimpan hasil perhitungan
// 	// for _, karakter := range teks {  // Mengiterasi setiap karakter dalam teks
// 	// 	// if karakter >= 'a' && karakter <= 'z' { // Mengecek apakah karakter adalah huruf
// 	// 	hitungan[string(karakter)]++ // Menambahkan jumlah huruf tersebut dengan satu
// 	// 	// }
// 	// 	fmt.Println("hellow", hitungan)
// 	// }

// 	// // Mengurutkan hasil perhitungan berdasarkan jumlah huruf terbanyak
// 	// huruf := make([]string, 0, len(hitungan)) // Membuat slice untuk menyimpan key dari map
// 	// for k := range hitungan {
// 	// 	huruf = append(huruf, k) // Mengisikan slice dengan key dari map
// 	// }
// 	// // sort.Slice(huruf, func(i, j int) bool { // Mengurutkan slice berdasarkan kriteria pengurutan
// 	// // 	return hitungan[huruf[i]] > hitungan[huruf[j]] // Kriteria pengurutan adalah jumlah kemunculan huruf yang lebih besar
// 	// // })

// 	// // Menampilkan hasil perhitungan dalam bentuk tabel
// 	// fmt.Println("Hasil perhitungan:")
// 	// fmt.Println("+-------+-------+")   // Mencetak garis horizontal
// 	// fmt.Printf("| Huruf | Jumlah |\n") // Mencetak header tabel
// 	// fmt.Println("+-------+-------+")   // Mencetak garis horizontal
// 	// for _, h := range huruf {          // Mengiterasi setiap elemen dalam slice yang sudah diurutkan
// 	// 	fmt.Printf("| %5s | %5d |\n", h, hitungan[h]) // Mencetak key (huruf) dan value (jumlah kemunculan) dari map dengan format tabel
// 	// }
// 	// fmt.Println("+-------+-------+")         // Mencetak garis horizontal
// 	// vocal := "hello"
// 	// strings.Split(vocal, "")
// 	// sound := 0
// 	// //ehlo
// 	// arr := []string{"a", "i", "u", "e", "o"} //aeiou
// 	// urut := strings.Split(vocal, "")
// 	// // fmt.Println(urut[0], "hello urut")
// 	// for i := 0; i < len(vocal); i++ {
// 	// 	for j := 0; j < len(vocal); j++ {
// 	// 		if urut[i] == arr[j] {
// 	// 			sound += 1
// 	// 		}
// 	// 	}
// 	// }

// 	// teks := "katak"
// 	// cobain := strings.Split(teks, "")
// 	// for i := 0; i < len(cobain); i++ {
// 	// 	if cobain[len(cobain)-i-1] == cobain[i] {
// 	// 		fmt.Println("yes")
// 	// 	} else {
// 	// 		fmt.Println("no")
// 	// 	}
// 	// }
// 	// fmt.Println("cobas", len(cobain))
// 	// FactorialPositive(5)
// 	// PrimaBilangan(13)
// 	// KebabCase("hello world sena")
// 	// WordCount("How are you today?")
// 	// HighLow([]int{2, 3, 1, 4, 5, 9})
// 	// duplicate([]int{1, 1, 2, 2, 3, 4, 5}) //1, 2, 3, 2, 4, 1, 5
// 	VocalSound("senapahlevi") //1, 2, 3, 2, 4, 1, 5

// }

// func MultipleFormTextTravel(numPersons int, input1, input2, input3, input4 string) string {
// 	var DataMultipleFormTravel []string
// 	//using 4 params because in travel only max 4 families member
// 	if input1 != "" {
// 		DataMultipleFormTravel = append(DataMultipleFormTravel, input1)
// 	}
// 	if input2 != "" {
// 		DataMultipleFormTravel = append(DataMultipleFormTravel, input2)
// 	}
// 	if input3 != "" {
// 		DataMultipleFormTravel = append(DataMultipleFormTravel, input3)
// 	}
// 	if input4 != "" {
// 		DataMultipleFormTravel = append(DataMultipleFormTravel, input4)
// 	}
// 	data := strings.Join(DataMultipleFormTravel, ",")
// 	// fmt.Println("firstcoba", firstcoba)
// 	// fmt.Println("leng", len(children))

// 	if len(DataMultipleFormTravel) != numPersons {
// 		return ""
// 	}
// 	return data
// }

// func FactorialPositive(input int) {
// 	// input := 5
// 	// var result = 0
// 	// var pengali = 0
// 	simpan := 1
// 	for i := 1; i <= (input); i++ {
// 		simpan *= i
// 		fmt.Println(simpan, "simpan", i, "is")
// 	}
// }
// func VocalSound(input string) {
// 	inputArey := strings.Split(input, "")

// 	// sort.Strings(hasile)
// 	sum := 0
// 	vocale := "aiueo"
// 	vocaleArey := strings.Split(vocale, "")
// 	strings.Contains(vocale, "a")
// 	// fmt.Println("hasile", hasile)

// 	for i := 0; i < len(inputArey); i++ {
// 		for j := 0; j < len(vocaleArey); j++ {
// 			if inputArey[i] == vocaleArey[j] {
// 				sum += 1
// 			}
// 		}
// 	}
// 	fmt.Println(sum, "hasile")
// }

// func PrimaBilangan(input int) string {
// 	if input == 0 || input == 1 {
// 		fmt.Println("no")
// 	}
// 	for i := 2; i < input; i++ {
// 		if input%i == 0 {
// 			// fmt.Println("no")
// 			return "no"
// 		} else {
// 			fmt.Println("ye")
// 		}
// 	}
// 	// fmt.Println("ye")
// 	return "yeey"
// }
// func KebabCase(input string) {
// 	// pisah := strings.Split(input, "")
// 	coba := strings.Replace(input, " ", "-", 2)
// 	fmt.Println("coba", coba)
// }

// func WordCount(input string) {
// 	pisah := strings.Split(input, " ")

// 	fmt.Println(len(pisah), ":pisah")
// }
// func HighLow(arey []int) {
// 	// simpan := 1
// 	// temp := 0
// 	// arey := []int{2, 3, 1, 4, 5, 9}
// 	for i := 0; i < len(arey); i++ {
// 		for j := 0; j < len(arey)-1; j++ {
// 			if arey[j] > arey[j+1] {
// 				simpan := arey[j]
// 				// arey[j] = arey[j+1]
// 				// arey[j+1] = simpan
// 				fmt.Println("coba", simpan)
// 			}
// 		}
// 	}
// 	fmt.Println(arey, "hallo arrey")
// }

// func duplicate(arey []int) {
// 	// arey := []int{2, 3, 1, 1, 3, 2, 5}
// 	simpan := []int{}
// 	iseng := []int{}
// 	for i := 0; i < len(arey)-1; i++ {
// 		if arey[i] == arey[i+1] {
// 			simpan = append(simpan, arey[i])
// 			fmt.Println("coba", simpan)
// 		} else {
// 			iseng = append(iseng, arey[i])

// 		}

// 	}
// 	fmt.Println("simapn", simpan)
// 	fmt.Println("iseng", iseng)
// }
