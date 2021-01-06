package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
)

type Element struct {
	value int
	nb []int //neighborhood
	visited bool
}

var count int = 0

func  main() {
	graf  := [13]Element{}

	for i := 0; i < 13; i++ {
		graf[i].value = i + 1
		graf[i].visited = false
	}
	//yep, it is hardcode
	graf[0].nb = []int{1, 2} //number of element in array
	graf[1].nb = []int{3}
	graf[2].nb = []int{4}
	graf[3].nb = []int{5, 6}
	graf[4].nb = []int{6, 7}
	graf[5].nb = []int{8}
	graf[7].nb = []int{8, 10}
	graf[8].nb = []int{9}
	graf[10].nb = []int{11, 12}

	for i := 0; i < 13; i++ {
		fmt.Println(graf[i])
	}

	search, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatalf("incorect value second element")
	}
	fmt.Printf("\n")
	fmt.Printf("DFS:\n")
	result1 := dfs(graf, 0, search)
	fmt.Printf("\n")

	fmt.Println(count)
	fmt.Println(result1)

	for i := 0; i < 13; i++ {
		graf[i].visited = false
	}
	fmt.Printf("\n")
	fmt.Printf("BFS:\n")
	result2 := bfs(graf, 0, search)
	fmt.Println(result2)
}

func dfs(graf [13]Element, verIndex int, srchValue int) bool {
	if graf[verIndex].value == srchValue {
		return true
	}
	if graf[verIndex].visited {
		return false
	}

	graf[verIndex].visited = true

	for i := 0; i < len(graf[verIndex].nb); i++ {
		indexNB := graf[verIndex].nb[i]

		if (!graf[indexNB].visited) {
			reached := dfs(graf, indexNB, srchValue)
			if (reached) {
				fmt.Printf("%v ", indexNB + 1)
				count++
				return true
			}
		}
	}

	return false
}

func bfs(graf [13]Element, verIndex int, srchValue int) bool {
	queue := []int{}

	queue = append(queue, verIndex)
	graf[verIndex].visited = true
	

	for len(queue) > 0 {
		v, a := queue[0], queue[1:]
		queue = a

		for i := 0; i < len(graf[v].nb); i++ {
			indexNB := graf[v].nb[i]

			if (!graf[indexNB].visited) {
				graf[indexNB].visited = true
				queue = append(queue, indexNB)

				if (graf[indexNB].value == srchValue) {
					return true
				}
			}
		}
		
	}

	return false
}
	
