package ajl

import (
	"fmt"
	"testing"
)

func TestAJL(t *testing.T) {
	ajl := New[int]()
	//       /--------_\
	//  1 -> 2 -> 3 -> 4
	//   \______-/
	ajl.Add(1, 2, 1)
	ajl.Add(1, 2, 1)
	ajl.Add(2, 3, 1)
	ajl.Add(1, 3, 1)
	ajl.Add(3, 4, 1)
	//ajl.Add(4, 3, 1)
	ajl.Add(2, 4, 1)
	fmt.Println(ajl)
	fmt.Println("----------")
	fmt.Println("BFS")
	ajl.BfsRange(1, func(a, b int, w int) bool {
		fmt.Println(a, "->", b)
		return true
	})

	fmt.Println("----------")
	fmt.Println("DFS")
	ajl.DfsRange(1, func(a, b int, w int) bool {
		fmt.Println(a, "->", b)
		return true
	})

	fmt.Println("----------")
	fmt.Println(ajl)
	fmt.Println("TopSort")
	topList, ok := ajl.TopSort()
	if ok {
		fmt.Println(topList)
		fmt.Println(ajl.cacheTopSort)
	}

	//       /---------\
	//  1   2 -> 3 <-> 4
	//   \_______/
	fmt.Println("----------")
	fmt.Println("remove 1->2", ajl.Remove(1, 2))
	ajl.Remove(5, 4)
	ajl.Remove(4, 5)
	ajl.Remove(5, 6)
	fmt.Println("DFS")
	ajl.DfsRange(1, func(a, b int, w int) bool {
		fmt.Println(a, "->", b)
		return true
	})

}
