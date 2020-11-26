package leetcode

/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start

type Tree struct {
	isRoot bool
	isLeaf bool
	value string
	children map[byte]*Tree
}

func findWords(board [][]byte, words []string) []string {
	wordTree := &Tree{
		isRoot: true,
	}
	
	for _, word := range words {
		curNode := wordTree
		for i := 0; i < len(word); i++ {
			if curNode.children == nil {
				curNode.children = make(map[byte]*Tree)
			}

			child, _ := curNode.children[word[i]]
			if child == nil {
				child = &Tree{}
			}
			curNode.children[word[i]] = child

			curNode = child
		}
		curNode.isLeaf = true
		curNode.value = word
	}

	result := []string{}
	xSize := len(board)
	ySize := len(board[0])

	// x 和 y 表示当前位置，parent则是遍历树
	// 如果board[x][y] in parent.children则进行递归探索，测试是否是叶子
	// 此时尝试x-1, x+1, y-1, y+1四个方向，并且将当前node标记为visited
	var backtracing func(x, y int, parent *Tree) 

	backtracing = func(x, y int, parent *Tree) {
		if parent == nil {
			return
		}

		curLetter := board[x][y]

		if curLetter == '#' {
			return
		}

		child, _ := parent.children[curLetter]
		if child == nil {
			return
		}
		
		if child.isLeaf {
			result = append(result, child.value)
		}

		offsets := [][]int{
			[]int{-1, 0},
			[]int{0, -1},
			[]int{1, 0},
			[]int{0, 1},
		}
		
		board[x][y] = '#'
		for _, offset := range offsets {
			newX, newY := x+offset[0], y+offset[1]

			if newX >=0 && newX < xSize && newY >=0 && newY < ySize {
				backtracing(newX, newY, child)
			}
		}
		board[x][y] = curLetter
	}

	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			backtracing(x, y, wordTree)
		}
	}

	return unique(result)
}

func unique(a []string) []string {
	result := []string{}
	amap := map[string]bool{}

	for _, s := range a {
		if _, has := amap[s]; !has {
			result = append(result, s)
			amap[s] = true
		}
	}

	return result
}
// @lc code=end

