package tsp

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestFindShortestCycle(t *testing.T) {
	g, expected := getTestData("tests/test1.txt")
	actual := g.FindShortestCycle()
	assert(t, actual, expected)

	g, expected = getTestData("tests/test2.txt")
	actual = g.FindShortestCycle()
	assert(t, actual, expected)

	g, expected = getTestData("tests/test3.txt")
	actual = g.FindShortestCycle()
	assert(t, actual, expected)

	g, expected = getTestData("tests/test4.txt")
	actual = g.FindShortestCycle()
	assert(t, actual, expected)
}

func assert(t *testing.T, actual, expected any) {
	if expected != actual {
		t.Errorf("Result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func getTestData(fileName string) (Graph, int) {
	var g Graph
	var expected int

	inFile, _ := os.Open(fileName)
	in := bufio.NewReader(inFile)

	fmt.Fscan(in, &g.n)

	g.A = make([][]int, g.n)

	for i := 0; i < g.n; i++ {
		g.A[i] = make([]int, g.n)
		for j := 0; j < g.n; j++ {
			fmt.Fscan(in, &g.A[i][j])
		}
	}

	fmt.Fscan(in, &expected)

	return g, expected
}
