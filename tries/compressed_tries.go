package tries

import "fmt"

type TrieNode struct {
	isEnd               bool
	data                string
	mp                  []*TrieNode
	occurrenceListIndex int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		mp:                  make([]*TrieNode, 36),
		isEnd:               false,
		occurrenceListIndex: -1,
	}
}

type Trie struct {
	root *TrieNode
	size int
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
		size: 0,
	}
}

func charIndex(char uint8) int {
	if char >= "a"[0] && char <= "z"[0] {
		return int(char - "a"[0])
	} else {
		return int(26 + char - "0"[0])
	}
}

func (t *Trie) GetRoot() *TrieNode {
	return t.root
}

func (t *Trie) Search(token string) (docId int, err error) {
	currNode := t.root
	i := 0
	j := 0

	for {
		for i < len(token) && j < len(currNode.data) && token[i] == currNode.data[j] {
			i++
			j++
		}

		if i == len(token) && j == len(currNode.data) {
			if currNode.isEnd {

				return currNode.occurrenceListIndex, nil
			} else {
				return -1, fmt.Errorf("Search string not found")
			}
		}

		if i < len(token) && j < len(currNode.data) {
			return -2, fmt.Errorf("Search string not found")
		}

		if i < len(token) && j == len(currNode.data) {
			if currNode.mp[charIndex(token[i])] != nil {
				currNode = currNode.mp[charIndex(token[i])]
				i += 1
				j = 0
			} else {
				return -1, fmt.Errorf("Search string not found")
			}
		}
		if i == len(token) && j < len(currNode.data) {
			return -1, fmt.Errorf("Search string not found")
		}
	}
}

func PrintTreeStructure(node *TrieNode) {

	p := fmt.Sprintf("node is %p, node ind is %v, node data is %s, isEnd is %t, ", node, node.occurrenceListIndex, node.data, node.isEnd)

	revInd := func(ind uint8) string {
		if ind < 26 {
			return string("a"[0] + ind)
		} else {
			return string("0"[0] + ind - 26)
		}
	}

	for i, v := range node.mp {
		if v != nil {
			p += fmt.Sprintf("%s-%p,", revInd(uint8(i)), v)
		}
	}

	fmt.Println(p)

	for _, v := range node.mp {
		if v != nil {
			PrintTreeStructure(v)
		}
	}
}

func (t *Trie) GetSize() int {
	return t.size
}

func (t *Trie) Insert(token string, docId int) error {
	currNode := t.root
	i := 0
	j := 0

	for {
		for i < len(token) && j < len(currNode.data) && token[i] == currNode.data[j] {
			i++
			j++
		}

		if i == len(token) {
			if j != len(currNode.data) {
				newNode := NewTrieNode()
				if j != len(currNode.data)-1 {
					newNode.data = currNode.data[j+1:]
				} else {
					newNode.data = ""
				}
				newNode.isEnd = currNode.isEnd
				newNode.occurrenceListIndex = currNode.occurrenceListIndex
				copy(newNode.mp, currNode.mp)
				currNode.isEnd = false
				currNode.mp = nil
				currNode.mp = make([]*TrieNode, 36)
				currNode.mp[charIndex(currNode.data[j])] = newNode
				currNode.data = currNode.data[:j]
				currNode.occurrenceListIndex = docId
				t.size++
			}
			currNode.isEnd = true
			currNode.occurrenceListIndex = docId

			return nil
		} else if j == len(currNode.data) {
			if currNode.mp[charIndex(token[i])] == nil {
				newNode := NewTrieNode()
				newNode.data = token[i+1:]
				newNode.isEnd = true
				newNode.occurrenceListIndex = docId
				currNode.mp[charIndex((token[i]))] = newNode
				t.size++
				return nil
			} else {
				currNode = currNode.mp[charIndex(token[i])]
				j = 0
				i += 1
			}
		} else {
			newNode := NewTrieNode()
			if j != len(currNode.data)-1 {
				newNode.data = currNode.data[j+1:]
			} else {
				newNode.data = ""
			}
			newNode.isEnd = currNode.isEnd
			newNode.occurrenceListIndex = currNode.occurrenceListIndex
			copy(newNode.mp, currNode.mp)
			currNode.isEnd = false
			currNode.mp = nil
			currNode.mp = make([]*TrieNode, 36)
			currNode.mp[charIndex(currNode.data[j])] = newNode
			currNode.data = currNode.data[:j]

			newNode = NewTrieNode()
			newNode.data = token[i+1:]
			newNode.isEnd = true
			currNode.mp[charIndex(token[i])] = newNode
			t.size += 2
		}
	}
}
