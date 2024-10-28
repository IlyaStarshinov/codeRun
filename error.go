package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	nStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	probabilites := make([]float64, n)
	totalPropability := 0.0

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		var ai, bi int
		fmt.Sscanf(line, "%d %d", &ai, &bi)

		p := float64(ai) / 100.0 * float64(bi) / 100.0
		probabilites[i] += p
		totalPropability += p
	}

	for _, p := range probabilites {
		probabilityOfError := p / totalPropability
		fmt.Fprintf(writer, "%.12f\n", probabilityOfError)
	}

}
