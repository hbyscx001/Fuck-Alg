package leetcode

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	children map[rune]*Trie
	isEnd bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	instance := Trie{
		isEnd: true,
		children: make(map[rune]*Trie),
	}

	return instance
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	cur := this

	for _, r := range word {
		if _, ok := cur.children[r]; !ok {
			cur.children[r] = &Trie{
				isEnd: false,
				children: make(map[rune]*Trie),
			}
		}
		cur = cur.children[r]
	}

	cur.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this

	for _, r := range word {
		if n, ok := cur.children[r]; ok {
			cur = n
		} else {
			return false
		}
	}

	return cur.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this

	for _, r := range prefix {
		if n, ok := cur.children[r]; ok {
			cur = n
		} else {
			return false
		}
	}

	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

