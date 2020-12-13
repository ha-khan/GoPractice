package utils


// Trie is the wrapper/top-level object which repr the Trie as an entire entity itself
type Trie struct {
	Root *TrieNode
}

// TrieNode is the recursive
type TrieNode struct {
	OutEdge []*TrieNode
	Val     string
}
