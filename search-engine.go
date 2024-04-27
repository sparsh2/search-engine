package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"search-engine/tries"
	"sort"
	"strings"
)

var workingDir = "working-dir"
var processedDataFilename = "processed-data.json"
var processedFilepath = path.Join(".", workingDir, processedDataFilename)
var crawledDataFilepath = path.Join(".", workingDir, "output.json")

// run crawler.py
func runCrawler() {
	cmd := exec.Command("python", "crawler.py", "magnus-carlsen.html")
	// get command output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	// fmt.Println("Running the crawler")
	// runCrawler()
	// fmt.Println()
	// fmt.Printf("Crawler has successfully finished crawling the pages\n\n")
	fmt.Println("Constructing the trie from the crawled data")
	trie := readAndConstructTrie()
	fmt.Printf("Trie has been successfully constructed. Number of nodes in the trie: %v\n\n", trie.GetSize())
	for {
		searchString := ""
		fmt.Printf("Enter the keywords to search through the pages: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			searchString = scanner.Text()
		}
		fmt.Println()
		lowerSearchString := strings.ToLower(searchString)
		alphaNumeric := "abcdefghijklmnopqrstuvwxyz0123456789,"
		searchString = ""
		for _, v := range lowerSearchString {
			if strings.Contains(alphaNumeric, string(v)) {
				searchString += string(v)
			}
		}

		searchTokens := []string{}
		for _, v := range strings.Split(searchString, ",") {
			if v != "" {
				searchTokens = append(searchTokens, strings.Trim(v, " "))
			}
		}

		docsLists := [][]DocAndFreq{}
		ok := true

		for _, v := range searchTokens {
			docId, err := trie.Search(v)
			if err != nil {
				fmt.Printf("Error getting doc list for the key \"%v\": %v\n\n", v, err)
				ok = false
				break
			}
			docsList := getDocNameAndFrequency(docId)
			docsLists = append(docsLists, docsList)
			s := ""
			for _, doc := range docsList {
				s += fmt.Sprintf("%s,", doc.DocName)
			}
			fmt.Printf("keyword \"%v\" found in: [%v]\n", v, s)
		}

		if !ok {
			fmt.Printf("No matching docs found containing all keywords\n\n")
			continue
		}

		// find intersection of the docs across all keywords
		docIdxs := []int{}
		commonDocsLists := []DocAndFreq{}
		for i := 0; i < len(docsLists); i++ {
			docIdxs = append(docIdxs, 0)
		}

		for {
			// check if all the pointers are less than the length of the lists
			lessThanLen := true
			for i, v := range docIdxs {
				if v >= len(docsLists[i]) {
					lessThanLen = false
					break
				}
			}
			if !lessThanLen {
				break
			}

			// get the max doc index from all the pointers
			maxDocIdx := 0
			for i, v := range docIdxs {
				if docsLists[i][v].DocId > maxDocIdx {
					maxDocIdx = docsLists[i][v].DocId
				}
			}

			// for each pointer, increment the pointer as long as it is less than the max doc index
			done := false
			for i := 0; i < len(docIdxs); i++ {
				for docIdxs[i] < len(docsLists[i]) && docsLists[i][docIdxs[i]].DocId < maxDocIdx {
					docIdxs[i]++
				}
				if docIdxs[i] == len(docsLists[i]) {
					done = true
					break
				}
			}
			if done {
				break
			}

			// if all the pointers point to the same doc, add the doc to the commonDocsList
			allEqual := true
			for i, v := range docIdxs {
				if docsLists[i][v].DocId != maxDocIdx {
					allEqual = false
					break
				}
			}

			if allEqual {
				commonDocsLists = append(commonDocsLists, docsLists[0][docIdxs[0]])
				// increment one of the pointers to move to the next doc
				docIdxs[0]++
			}
		}

		if len(commonDocsLists) == 0 {
			fmt.Printf("No matching docs found containing all keywords\n\n")
		} else {
			sort.Slice(commonDocsLists, func(i, j int) bool {
				return commonDocsLists[i].Freq > commonDocsLists[j].Freq
			})
			relevantDocNames := []string{}
			for _, doc := range commonDocsLists {
				relevantDocNames = append(relevantDocNames, doc.DocName)
			}
			fmt.Printf("\nRelevant docs: %v\n\n", relevantDocNames)
		}

	}
}

func getDocNameAndFrequency(docListId int) []DocAndFreq {
	bytes, err := os.ReadFile(processedFilepath)
	if err != nil {
		panic(err)
	}

	processedData := &ProcessedData{}
	err = json.Unmarshal(bytes, processedData)
	if err != nil {
		panic(err)
	}

	list := processedData.Data[docListId]

	docList := []DocAndFreq{}
	for _, v := range list {
		docList = append(docList, DocAndFreq{
			DocId:   v[0],
			DocName: processedData.DocIds[v[0]],
			Freq:    v[1],
		})
	}

	return docList
}

func readAndConstructTrie() *tries.Trie {
	data, err := os.ReadFile(crawledDataFilepath)
	if err != nil {
		panic(err)
	}
	jsonData := &CrawlerData{}
	err = json.Unmarshal(data, jsonData)
	if err != nil {
		panic(err)
	}

	tokenTrie := tries.NewTrie()
	docIdList := [][][]int{}
	i := 0
	for k, v := range jsonData.Data {
		tokenTrie.Insert(k, i)
		i++
		docIdList = append(docIdList, v)
	}

	pData := &ProcessedData{
		DocIds: jsonData.DocIds,
		Data:   docIdList,
	}

	bytes, err := json.MarshalIndent(pData, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(processedFilepath, bytes, 0777)
	if err != nil {
		panic(err)
	}

	return tokenTrie
}

type DocAndFreq struct {
	DocId   int
	DocName string
	Freq    int
}

type CrawlerData struct {
	DocIds map[int]string     `json:"doc_ids"`
	Data   map[string][][]int `json:"data"`
}

type ProcessedData struct {
	DocIds map[int]string `json:"doc_ids"`
	Data   [][][]int      `json:"data"`
}
