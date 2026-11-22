// Package vreme omogoča dodajanje in statistično obdelavo
// temperaturnih meritev vremenskih postaj.
package vreme

import "fmt"

// Postaja struktura, ki predstavlja vremensko postajo z njenimi meritvami.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo doda novo vremensko postajo v slovar postaj.
func DodajPostajo(postaje map[string]Postaja, oznaka string, ime string, kraj string) {
	postaje[oznaka] = Postaja{
		ime:  ime,
		kraj: kraj,
	}
}

// DodajMeritev doda meritev postaje v map s ključem oznaka
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp float64, maxTemp float64) {
	if temperatura > maxTemp || temperatura < minTemp {
		fmt.Printf("Temperatura mora biti med %.2f in %.2f°C\n", minTemp, maxTemp)
		return
	}

	if _, ok := postaje[oznaka]; !ok {
		fmt.Println("Ta postaja ne obstaja")
		return
	}

	p := postaje[oznaka]
	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p
}

// statistika pomožna funkcija za poročilo
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	if _, je := postaje[oznaka]; !je {
		fmt.Println("Ta postaja ne obstaja")
		ok = false
		return
	}
	postaja := postaje[oznaka]

	if len(postaja.meritve) == 0 {
		ok = false
		return
	}

	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]

	j := 0
	for _, v := range postaja.meritve {
		if v > najvisja {
			najvisja = v
		}
		if v < najnizja {
			najnizja = v
		}
		povprecna += v
		j++
	}
	if j < stMeritev {
		ok = false
		return
	}
	ok = true
	povprecna /= float64(j)

	return
}

// IzpisMeritev izpiše meritve vseh postaj
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for i, v := range postaje {
		fmt.Print(i, " - ", v.ime, " (", v.kraj, ")", ": ", v.meritve, "\n")
	}
}

// IzpisPorocila izpiše v cmd poročilo vseh postaj, min, max in povprečje temperatur
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for i, v := range postaje {
		min, max, povp, ok := statistika(postaje, i, stMeritev)
		izpis := ""
		switch {
		case povp >= 25.0:
			izpis = "Vroče"
		case povp >= 15.0:
			izpis = "Toplo"
		case povp >= 5.0:
			izpis = "Hladno"
		case povp < 5.0:
			izpis = "Mrzlo"
		}
		if ok == false {
			fmt.Printf("%s: premalo meritev\n", v.ime)
			continue
		}
		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", v.ime, min, max, povp, izpis)
	}
}
