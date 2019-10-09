package leetcode

type trieNode struct {
	val        rune
	isWord     bool
	childTable []*trieNode
}

type Trie struct {
	root *trieNode
}

/** Initialize your data structure here. */
func Constructor1() Trie {
	node := Trie{}
	node.root = new(trieNode)
	node.root.childTable = make([]*trieNode, 26, 26)
	node.root.val = ' '
	return node
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root

	for _, v := range word {
		if node.childTable[rune(v)-'a'] == nil {
			node.childTable[v-97] = &trieNode{val: v}
			node.childTable[v-97].childTable = make([]*trieNode, 26, 26)
		}
		node = node.childTable[v-97]
	}
	node.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {

	node := this.root

	for _, v := range word {
		if node.childTable[v-97] == nil || node.childTable[v-97].val != v {
			return false
		}
		node = node.childTable[v-97]
	}

	return node.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {

	node := this.root

	for _, v := range prefix {
		if node.childTable[v-97] == nil || node.childTable[v-97].val != v {
			return false
		}
		node = node.childTable[v-97]
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
