package main

import "fmt"
import (
    "io/ioutil"
    "path/filepath"
    "strings"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func palindroomiTest(sona []byte) {
	var pd bool
	var p int
	p = len(sona)
	// fmt.Println(string(sona), " ", p)
	pd = true
	for i := 0; i < p/2; i++ {
		// fmt.Println(sona[i], " ... ", sona[p - i - 1])
		if sona[i] != sona[p - i - 1] {
			pd = false
			break
		} 
	}
	if pd == true {
		fmt.Println("Palindroomsõna: ", string(sona))
	}	
}

// Tagastab pööratud sõna, iseseisva viiluna
func pooraSona(sona []byte) []byte {
	p := len(sona)
	pSona := make([]byte, p)
 	for i := 0; i < p; i = i + 1 {
        pSona[i] = sona[p - i - 1]
    }
    return pSona
}

// Reverse returns its argument string reversed rune-wise left to right.
// http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

// Selgitab välja, kas sõna on samasõna (tyyp 1), lõugsõna (tyyp 2)
// või tavasõna (tyyp 0)
func selgitaSonaTyyp(algSona string, pSona string, sonad []string) int {
	tyyp := 0 // eeldame tavasõna
	for _, sona := range sonad {
	 	if strings.Compare(pSona, sona) == 0 {
	 		if strings.Compare(algSona, pSona) == 0 {
	 			// samasõna
	 			tyyp = 1
				// fmt.Println(algSona)
				break		 			
	 		} else {
	 			// lõugsõnad
	 			tyyp = 2
				// fmt.Println(algSona, " -- ", pSona)
				break		 			
	 		}
		}
	}
	return tyyp
}

func poordsonatest(sonad []string) []string {
	var tulem []string
	for _, sona := range sonad {
  		pSona := Reverse(sona)
 		tyyp := selgitaSonaTyyp(sona, pSona, sonad)
 		switch {
 		case tyyp == 0:
 		case tyyp == 1:
 			tulem = append(tulem, sona)
 		case tyyp == 2:
 			tulem = append(tulem, sona + "  " + pSona)	
 		}
 	}
 	return tulem
}

func main() {

	sonafailinimi := filepath.Clean("C:/Users/Kasutaja/Desktop/KEEL/SONAD")
	tulemifailinimi := filepath.Clean("C:/Users/Kasutaja/Desktop/KEEL/TULEM.txt")

    loetudFail, err := ioutil.ReadFile(sonafailinimi)
    check(err)

    // Teisendamine stringimassiiviks
    sonad := strings.Split(string(loetudFail), "\r\n")

    // Ava fail tulemi salvestamiseks
    f, err := os.Create(tulemifailinimi)
    check(err)

    defer f.Close()

  	tulem := poordsonatest(sonad)

  	// Tulemi kirjutamine faili
  	for _, tulemiRida := range tulem {
  		f.WriteString(tulemiRida + "\r\n")
  	}

	f.Sync()

  	fmt.Println("Tulemi pikkus: ", len(tulem))

}