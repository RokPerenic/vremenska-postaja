package main

import "fmt"

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func dodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64) {
	if temperatura > 50 || temperatura < -50 {
		fmt.Println("Temperatura mora biti med -50°C in 50°C")
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

func statistika(postaje map[string]Postaja, oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	if _, je := postaje[oznaka]; !je {
		fmt.Println("Ta postaja ne obstaja")
		ok = false
		return
	}
	najvisja = -51
	najnizja = 51

	postaja := postaje[oznaka]
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

	if j < 3 {
		ok = false
		return
	}
	ok = true
	povprecna /= float64(j)

	return
}

func izpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for i, v := range postaje {
		fmt.Print(i, " - ", v.ime, " (", v.kraj, ")", ": ", v.meritve, "\n")
	}
}

func izpisPorocila(postaje map[string]Postaja) {
	for i, v := range postaje {
		min, max, povp, ok := statistika(postaje, i)
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

func main() {
	temperature := map[string]Postaja{"LJ01": {"Ljubljana Center", "Ljubljana", []float64{21, 12}}, "PO01": {"Tržaška cesta", "Postojna", []float64{11, 8.3, 17.5, 21.5}}, "GO01": {"Rožna Dolina", "Nova Gorica", []float64{5, 13.3, 17.8, 21.5}}, "MB01": {"Trg Leona Štuklja", "Maribor", []float64{10.1, 11.1, 15.5, 22.5}}}
	dodajMeritev(temperature, "PO01", 22.4)
	dodajMeritev(temperature, "KR01", 12.5)
	dodajMeritev(temperature, "MB01", 63.3)
	izpisMeritev(temperature)
	fmt.Println("")
	izpisPorocila(temperature)

}
