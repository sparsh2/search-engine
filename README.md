# Search Engine

There are 2 main components to this project: 
- The Crawler
- The Search Engine

## The Crawler

The Crawler is written in Python - `crawler.py` in the root directory. It is responsible for crawling and preprocessing the local html pages content present under `raw-pages` directory. The processed data only contains **lowercase english alphabets** and **numbers**. It uses [Beautiful Soup 4](https://pypi.org/project/beautifulsoup4/) to parse the html and extract the text. It also filters out the stop-words. This preprocessed data is then saved in the disk at `working-dir/output.json` file.

## The Search Engine
The Search Engine is written in golang. It uses the preprocessed data from the crawler to construct compressed tries. The code for the compressed tries is under `tries/`. The value stored in the nodes is the corresponding index of the occurrence list as described in book. The occurrence list for all tokens is stored on the disk at `working-dir/processed-data.json` which will be queried at runtime to fetch list for each search token.

Each occurrence list for a token has the following form:
```
[[0, 2], [3, 7]]
```
This means that the token occurs in 2 different documents - document id `0` and `3`, with the corresponding frequency i.e it occurs 2 times in document `0` and `7` times in document `3`.

The `processed-data.json` has the following form:
```json
{
	'doc_ids': {
		0: 'home.html',
		1: 'books.html',
		2: 'cars.html'
	},
	'data' : [[[0, 1], [1, 4], [2, 2]], [[0, 2]]]
}
```
The `data` field is just a list of occurrence list for each token. The index of each occurrence list is stored in the trie. The `doc_ids` field maps document ids to the document names.

### Document ranking
Once the common documents are found, they are ranked according to the total occurrence frequency of all the tokens in that document.

For example, consider the search tokens: `chess`, `games`. `chess` occurs 5 times in `magnus-carlsen.html` and 3 times in `stockfish.html`. And `games` occurs 10 times in `magnus-carlsen.html` and 2 times in `stockfish.html`. Total frequency for `magnus-carlsen-html` = 5 + 10 = 15. Total frequency for `stockfish.html` = 3 + 2 = 5. Therefore, `magnus-carlsen.html` will be ranked higher than `stockfish.html`

## How to run the project?
> **Note**: 
> 
> Ensure you have installed the following:
> * Install [Beautiful Soup 4](https://pypi.org/project/beautifulsoup4/) pypi package
> * Install [Golang](https://go.dev/doc/install)

### Run the crawler
From the root directory, run `python crawler.py psychology.html` to start crawling from `psychology.html` page. You should see a similar output:
```
(cs600) sparsh@Zitkel:~/Acad/Adv-Algo/Project$ python crawler.py psychology.html
Crawling from psychology.html
extracting content from psychology.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
.
.
.
Found link to world-war-2.html. Adding to queue
extracting content from world-war-2.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
Found link to world-war-2.html. Adding to queue
extracting content from stockfish.html
Found link to magnus-carlsen.html. Adding to queue
finished indexing all information
```
This will output a json file - `working-dir/output.json` which will be used by the Search Engine.

### Run the Search Engine
From the root directory, run `go run search-engine.go`. Enter comma-separated search keywords to search through the documents. Example output:

```
(cs600) sparsh@Zitkel:~/Acad/Adv-Algo/Project$ go run search-engine.go 
Constructing the trie from the crawled data
Trie has been successfully constructed. Number of nodes in the trie: 6588

Enter the keywords to search through the pages: games, chess

keyword "games" found in: [magnus-carlsen.html,rafael-nadal.html,stockfish.html,]
keyword "chess" found in: [magnus-carlsen.html,stockfish.html,]

Relevant docs: [magnus-carlsen.html stockfish.html]
```
Another example output:

```
(cs600) sparsh@Zitkel:~/Acad/Adv-Algo/Project$ go run search-engine.go 
Constructing the trie from the crawled data
Trie has been successfully constructed. Number of nodes in the trie: 6588

Enter the keywords to search through the pages: Ice, Global,WARMING,history

keyword "ice" found in: [paleolithic.html,quarternary-glaciation.html,]
keyword "global" found in: [paleolithic.html,quarternary-glaciation.html,world-war-2.html,]
keyword "warming" found in: [paleolithic.html,quarternary-glaciation.html,]
keyword "history" found in: [psychology.html,magnus-carlsen.html,quarternary-glaciation.html,rafael-nadal.html,world-war-2.html,stockfish.html,]

Relevant docs: [quarternary-glaciation.html]
```


