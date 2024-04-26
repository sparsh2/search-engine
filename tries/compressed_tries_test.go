package tries

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompressedTries(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello", 0)
	trie.Insert("where1", 1)
	require.Equal(t, 2, trie.size)
	fmt.Println(len(trie.root.data))

	docId, err := trie.Search("hello")
	require.Equal(t, 0, docId)
	require.NoError(t, err)

	docId, err = trie.Search("where1")
	require.Equal(t, 1, docId)
	require.NoError(t, err)
}

func Test_AnotherCompressedTries(t *testing.T) {
	mp := map[string]int{
		"information": 0,
		"inform":      1,
		"hunter":      2,
		"example":     3,
		"sda7":        4,
		"sda77":       5,
		"558da":       6,
		"558damp":     7,
		"total":       8,
		"bunch":       9,
		"hunt":        10,
		"hunker":      11,
	}

	trie := NewTrie()

	for key, value := range mp {
		trie.Insert(key, value)
	}

	PrintTreeStructure(trie.root)

	for key, value := range mp {
		id, err := trie.Search(key)
		require.Equal(t, value, id, fmt.Sprintf("error for %s and %v", key, value))
		require.NoError(t, err, fmt.Sprintf("error for %s and %v", key, value))
	}
}

func Test_AnotherCompressedTries1(t *testing.T) {
	trie := NewTrie()

	trie.Insert("hunt", 0)
	PrintTreeStructure(trie.root)
	fmt.Println()
	trie.Insert("hunter", 1)

	PrintTreeStructure(trie.root)
	trie.Insert("hunker", 2)
	fmt.Println()
	PrintTreeStructure(trie.root)

	docId, err := trie.Search("hunt")
	require.Equal(t, 0, docId)
	require.NoError(t, err)

	docId, err = trie.Search("hunter")
	require.Equal(t, 1, docId)
	require.NoError(t, err)

	docId, err = trie.Search("hunker")
	require.Equal(t, 2, docId)
	require.NoError(t, err)

	_, err = trie.Search("98sdfso")
	require.Error(t, err)
}

func Test_AnotherCompressedTriesPermutation(t *testing.T) {
	trie := NewTrie()

	trie.Insert("tent", 0)
	trie.Insert("tester", 1)
	trie.Insert("test", 2)

	docInd, err := trie.Search("tent")
	require.Equal(t, 0, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("tester")
	require.Equal(t, 1, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("test")
	require.Equal(t, 2, docInd)
	require.NoError(t, err)

	trie = NewTrie()

	trie.Insert("tent", 0)
	trie.Insert("test", 2)
	trie.Insert("tester", 1)

	docInd, err = trie.Search("tent")
	require.Equal(t, 0, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("tester")
	require.Equal(t, 1, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("test")
	require.Equal(t, 2, docInd)
	require.NoError(t, err)

	trie = NewTrie()

	trie.Insert("tester", 1)
	trie.Insert("tent", 0)
	trie.Insert("test", 2)

	docInd, err = trie.Search("tent")
	require.Equal(t, 0, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("tester")
	require.Equal(t, 1, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("test")
	require.Equal(t, 2, docInd)
	require.NoError(t, err)

	trie = NewTrie()

	trie.Insert("tester", 1)
	trie.Insert("test", 2)
	trie.Insert("tent", 0)

	docInd, err = trie.Search("tent")
	require.Equal(t, 0, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("tester")
	require.Equal(t, 1, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("test")
	require.Equal(t, 2, docInd)
	require.NoError(t, err)

	trie = NewTrie()

	trie.Insert("test", 2)
	trie.Insert("tester", 1)
	trie.Insert("tent", 0)

	docInd, err = trie.Search("tent")
	require.Equal(t, 0, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("tester")
	require.Equal(t, 1, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("test")
	require.Equal(t, 2, docInd)
	require.NoError(t, err)

	trie = NewTrie()

	trie.Insert("test", 2)
	trie.Insert("tent", 0)
	trie.Insert("tester", 1)

	docInd, err = trie.Search("tent")
	require.Equal(t, 0, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("tester")
	require.Equal(t, 1, docInd)
	require.NoError(t, err)

	docInd, err = trie.Search("test")
	require.Equal(t, 2, docInd)
	require.NoError(t, err)

	_, err = trie.Search("hello4")
	require.Error(t, err)
	_, err = trie.Search("testit")
	require.Error(t, err)
	_, err = trie.Search("tes")
	require.Error(t, err)
	_, err = trie.Search("testers")
	require.Error(t, err)
}
