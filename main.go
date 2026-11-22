package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/RokPerenic/vremenska-postaja/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vremenska-postaja",
		Usage: "Aplikacija za spremljanje vremenskih postaj",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Value: -50,
				Usage: "Najnižja dovoljena temperatura",
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "Najvišja dovoljena temperatura",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "Število meritev",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			fmt.Println("Minimalna temperatura:", minTemp)
			temperature := make(map[string]vreme.Postaja)

			vreme.DodajPostajo(temperature, "LJ01", "Ljubljana Center", "Ljubljana")
			vreme.DodajPostajo(temperature, "PO01", "Tržaška cesta", "Postojna")
			vreme.DodajPostajo(temperature, "GO01", "Rožna Dolina", "Nova Gorica")
			vreme.DodajPostajo(temperature, "MB01", "Trg Leona Štuklja", "Maribor")

			vreme.DodajMeritev(temperature, "PO01", 22.4, minTemp, maxTemp)
			vreme.DodajMeritev(temperature, "LJ01", 12.5, minTemp, maxTemp)
			vreme.DodajMeritev(temperature, "GO01", 13.3, minTemp, maxTemp)
			vreme.IzpisMeritev(temperature)
			fmt.Println("")
			vreme.IzpisPorocila(temperature, stMeritev)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
