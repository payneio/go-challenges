package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Event struct {
	TS          string
	TSParsed    time.Time
	Utilization int
}

func main() {

	sc := make(chan Event, 100)
	ec := make(chan Event, 100)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			var s, e string
			var u int
			fmt.Sscanf(line, "%s %s %d", &s, &e, &u)
			start, err := time.Parse(time.RFC3339, s)
			if err != nil {
				fmt.Println("Bad format.", err, s)
				os.Exit(1)
			}
			end, err := time.Parse(time.RFC3339, e)
			if err != nil {
				fmt.Println("Bad format.", err, e)
				os.Exit(1)
			}

			sc <- Event{s, start, u}
			ec <- Event{e, end, -u}
		}
		sc <- Event{"end", time.Now(), 0}
		ec <- Event{"end", time.Now(), 0}

	}()

	sTop := <-sc
	eTop := <-ec

	for {
		select {
		case sTop = <-sc:
		case eTop = <-ec:
		}
		if eTop.TS == "end" {
			os.Exit(0)
		}
		if sTop.TS != "end" &&
			sTop.TSParsed.Before(eTop.TSParsed) {
			fmt.Println(sTop.TS, sTop.Utilization)
		} else {
			fmt.Println(eTop.TS, eTop.Utilization)
		}

	}

}
