package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"search-engine/tries"
	"sort"
	"strings"
)

var workingDir = "working-dir"
var processedDataFilename = "processed-data.json"
var processedFilepath = path.Join(".", workingDir, processedDataFilename)
var crawledDataFilepath = path.Join(".", workingDir, "output.json")

func main() {
	trie := readAndConstructTrie()
	fmt.Println(trie.GetSize())
	for {
		searchString := ""
		fmt.Printf("Enter the keywords to search through the pages: ")
		fmt.Scanf("%s", &searchString)
		fmt.Println()
		searchString = strings.ToLower(searchString)

		searchTokens := []string{}
		for _, v := range strings.Split(searchString, ",") {
			if v != "" {
				searchTokens = append(searchTokens, v)
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
			continue
		}

		// find intersection of the docs across all keywords
		docIdx := 0
		commonDocsLists := []DocAndFreq{}
		for {
			done := false
			// check if the docIdx is less than len of all docsLists
			for i := 0; i < len(docsLists); i++ {
				if docIdx == len(docsLists[i]) {
					// there will not be any other common docs, so break the loop
					done = true
					break
				}
			}

			if done {
				break
			}

			commonDocName := docsLists[0][docIdx].DocName
			sameDocName := true
			commonDocFreq := 0

			// check if all the indices point to the same doc
			for i := 0; i < len(docsLists); i++ {
				if docsLists[i][docIdx].DocName != commonDocName {
					sameDocName = false
				}
				commonDocFreq += docsLists[i][docIdx].Freq
			}

			// if all indices point to the same doc, append the doc to commonDocsList with sum of the frequency
			if sameDocName {
				commonDocsLists = append(commonDocsLists, DocAndFreq{
					DocName: commonDocName,
					Freq:    commonDocFreq,
				})
			}

			docIdx++
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
			fmt.Printf("Relevant docs: %v\n\n", relevantDocNames)
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
