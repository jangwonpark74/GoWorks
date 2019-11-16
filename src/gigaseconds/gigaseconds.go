package gigaseconds

import (
	"fmt"
	"time"
)

func AfterGigaSeconds() {
	t1 := time.Now()
	t2 := t1.Add(time.Second * 1000000000)
	fmt.Println("Now         :", t1)
	fmt.Println("AfterGigaSec:", t2)
}
