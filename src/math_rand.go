package main

import( 
	"fmt"
	"math/rand"
	"time"
)

func randInt(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return min + rand.Intn(max-min)
}

func main() {
	fmt.Printf("%v\n", randInt(1000, 1000000))
}
