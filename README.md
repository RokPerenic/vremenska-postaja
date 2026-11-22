# Vremenska postaja

Preprosta aplikacija, napisana v programskem jeziku Go, za spremljanje vremenskih postaj in beleženje temperaturnih meritev.

## Funkcionalnosti

- Dodajanje vremenskih postaj.
- Beleženje temperaturnih meritev.
- Nastavitev dovoljenega temperaturnega območja.
- Izpis vseh meritev.
- Izračun minimalne, maksimalne in povprečne temperature.
- Izpis statističnega poročila.

## Zagon

Kloniranje repozitorija:

```bash
git clone https://github.com/RokPerenic/vremenska-postaja.git
cd vremenska-postaja
```

Zagon aplikacije:

```bash
go run .
```

## Nastavitve

Aplikacija omogoča uporabo naslednjih zastavic:

| Zastavica | Opis | Privzeta vrednost |
|---|---|---|
| `--minTemp` | Najnižja dovoljena temperatura | -50 |
| `--maxTemp` | Najvišja dovoljena temperatura | 50 |
| `--stMeritev` | Minimalno število meritev za statistiko | 3 |

Primer zagona z lastnimi nastavitvami:

```bash
go run . --minTemp=-20 --maxTemp=35 --stMeritev=5
```

## Struktura projekta

- `main.go` – glavni program in ukazna vrstica.
- `vreme/vreme.go` – funkcije za upravljanje postaj in temperaturnih meritev.

## Uporaba paketa v drugih Go projektih
Paket vreme lahko uporabite tudi v drugih Go projektih.

Namestitev paketa:
```go
go get github.com/RokPerenic/vremenska-postaja/vreme
```

Uvoz paketa v Go datoteko:
```go
import "github.com/RokPerenic/vremenska-postaja/vreme"
```