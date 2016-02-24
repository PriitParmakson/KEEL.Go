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

// Reverse returns its argument string reversed rune-wise left to right.
// http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

// Sisendiks sõnade viil (sõnad esitatud stringidena)
// Tagastab tulemiridade viilu
func selgitaSonadeTyybid(sonad []string) []string {
	var tulem []string
	for i, sona1 := range sonad {
  		pSona := Reverse(sona1)

  		if i == 30000 {
  			break
  		}

		for j, sona2 := range sonad {

			if j == 30000 {
				break
			}

			if strings.Contains(sona2, pSona) {
			 	if strings.Compare(pSona, sona2) == 0 {
			 		if strings.Compare(sona1, pSona) == 0 {
			 			// samasõna
			 			tulem = append(tulem, sona1)
			 		} else {
			 			// lõugsõnad
			 			tulem = append(tulem, sona1 + " <-> " + pSona)
			 		}
				} else { // osaline sisalduvus
					// registreeri ainult vähemalt 4-tähelised sisalduvused
					if len(sona1) > 3 {
						tulem = append(tulem, sona1 + " -> " + sona2)
					}
				}
			}
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

  	tulem := selgitaSonadeTyybid(sonad)

  	// Tulemi kirjutamine faili
  	for _, tulemiRida := range tulem {
  		f.WriteString(tulemiRida + "\r\n")
  	}

	f.Sync()

  	fmt.Println("Tulemi pikkus: ", len(tulem))

}