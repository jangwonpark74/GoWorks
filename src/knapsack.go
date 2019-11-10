package main

import(
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sort"
	"errors"
)

type item struct {
	name string
	worth, weight int
}

type bag struct {
	bagWeight, currItemsWeight, maxItemsWeight, totalWeight int
	items []item
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readItem(path string) {
	dat, err := ioutil.ReadFile(path)
	check(err)
	fmt.Print(string(dat))
}



func readItems( path string) []item {
	
	dat, err := ioutil.ReadFile(path)
	check(err)

	fmt.Println("finish ReadFile")
	lines := strings.Split(string(dat), "\n")

	var itemList []item

	for i, v := range (lines) {
		if i == 0 {
			continue
		}
	 	s :=strings.Split(v, ",")
		newItemWorth, _ := strconv.Atoi(s[1])
		newItemWeigth, _ :=strconv.Atoi(s[2])
		newItem := item{name:s[0], worth:newItemWorth, weight: newItemWeigth}
		itemList = append(itemList, newItem)
	}
	fmt.Println("finish reading items")
	return itemList
}

// sort item by worth
// put an item  in the bag
// check to see if the back is full

func (b *bag) addItem(i item) error {
	if b.currItemsWeight + i.weight <= b.maxItemsWeight {
		b.currItemsWeight +=i.weight
		b.items = append(b.items, i)
		return nil
	}
	return errors.New("could not fit item")
}

func greedy( is []item, b bag) {

	sort.Slice(is, func(i, j int) bool {
		return is[i].worth > is[j].worth
	})

	for i := range is {
		b.addItem(is[i])
	}

	b.totalWeight = b.bagWeight + b.currItemsWeight

	for _, v := range b.items {
		b.totalWeight += v.worth
	}
}

func main() {

	minaal := bag{bagWeight: 1415, currItemsWeight: 0, maxItemsWeight: 5585}
	itemList := readItems("Objects.csv")

	greedy(itemList, minaal)
}



