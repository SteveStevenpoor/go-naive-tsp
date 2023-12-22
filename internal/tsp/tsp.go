package tsp

type Graph struct {
	n int
	A [][]int
}

func GivePermutations(pref, post []int) [][]int {
	var perms [][]int
	l := len(post)
	if l == 0 {
		perms = append(perms, pref)
		return perms
	}

	for r := 0; r < l; r++ {
		tmp := []int{}
		tmp = append(tmp, post[:r]...)
		tmp = append(tmp, post[r+1:]...)
		perms = append(perms, GivePermutations(append(pref, post[r]), tmp)...)
	}

	return perms
}

func (g *Graph) FindShortestCycle() int {
	if g.n <= 1 {
		return 0
	}

	startPerm := giveStartPerm(g.n)
	perms := GivePermutations([]int{}, startPerm)

	var min int = 1e9

	for i := range perms {
		l := len(perms[i])
		pathLen := 0
		isCyclic := true

		if g.A[0][perms[i][0]] == 0 || g.A[perms[i][l-1]][0] == 0 {
			continue
		} else {
			pathLen += g.A[0][perms[i][0]] + g.A[perms[i][l-1]][0]
		}

		for j := 0; j < l-1; j++ {
			from := perms[i][j]
			to := perms[i][j+1]
			if g.A[from][to] == 0 {
				isCyclic = false
				break
			} else {
				pathLen += g.A[from][to]
			}
		}

		if isCyclic && pathLen < min {
			min = pathLen
		}
	}

	return min
}

func giveStartPerm(n int) []int {
	startPerm := []int{}

	for i := 1; i < n; i++ {
		startPerm = append(startPerm, i)
	}

	return startPerm
}
