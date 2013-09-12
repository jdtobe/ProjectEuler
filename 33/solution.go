package main

import (
	fmt "fmt"
	log "log"
	strconv "strconv"
)

func main() {
	pn := 1.0
	pd := 1.0

	for i := 1; i < 10; i++ {
		for d := 1; d < 10; d++ {
			for n := 1; n < 10; n++ {
				if d == n {
					continue
				}

				bn, err := strconv.ParseFloat(fmt.Sprintf("%d%d", n, i), 64)
				if err != nil {
					log.Fatal(err)
				}

				bd, err := strconv.ParseFloat(fmt.Sprintf("%d%d", i, d), 64)
				if err != nil {
					log.Fatal(err)
				}

				sn := float64(n)
				sd := float64(d)

				if bn/bd == sn/sd {
					pn = pn * bn
					pd = pd * bd

					fmt.Printf("%0.0f\t/\t%0.0f = %f\n", bn, bd, bn/bd)
					fmt.Printf("%0.0f\t/\t%0.0f = %f\n", sn, sd, sn/sd)

					fmt.Println()
				}
			}
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Printf("%0.0f\t/\t%0.0f = %f\n", pn, pd, pn/pd)
}
