# Project Execution Output

## Sample 1
### Running `crawler.py`
```
(cs600) sparsh@Zitkel:~/Acad/Adv-Algo/Project$ python crawler.py paleolithic.html
Crawling from paleolithic.html
extracting content from paleolithic.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
Found link to world-war-2.html. Adding to queue
extracting content from magnus-carlsen.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
Found link to world-war-2.html. Adding to queue
Found link to stockfish.html. Adding to queue
extracting content from psychology.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
Found link to world-war-2.html. Adding to queue
extracting content from quarternary-glaciation.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
Found link to world-war-2.html. Adding to queue
extracting content from rafael-nadal.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
Found link to world-war-2.html. Adding to queue
extracting content from texas-revolution.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
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
(cs600) sparsh@Zitkel:~/Acad/Adv-Algo/Project$ 

```
### Running Search Engine
```
(cs600) sparsh@Zitkel:~/Acad/Adv-Algo/Project$ go run search-engine.go 
Constructing the trie from the crawled data
Trie has been successfully constructed. Number of nodes in the trie: 6587

Enter the keywords to search through the pages: games, Chess,

keyword "games" found in: [magnus-carlsen.html,rafael-nadal.html,stockfish.html,]
keyword "chess" found in: [magnus-carlsen.html,stockfish.html,]

Relevant docs: [magnus-carlsen.html stockfish.html]

Enter the keywords to search through the pages: 

Enter the keywords to search through the pages: GAMES,CHESS

keyword "games" found in: [magnus-carlsen.html,rafael-nadal.html,stockfish.html,]
keyword "chess" found in: [magnus-carlsen.html,stockfish.html,]

Relevant docs: [magnus-carlsen.html stockfish.html]

```

## Sample 2

### Running `crawler.py`

```
(cs600) sparsh@Zitkel:~/Acad/Adv-Algo/Project$ python crawler.py rafael-nadal.html
Crawling from rafael-nadal.html
extracting content from rafael-nadal.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
Found link to world-war-2.html. Adding to queue
extracting content from magnus-carlsen.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
Found link to world-war-2.html. Adding to queue
Found link to stockfish.html. Adding to queue
extracting content from paleolithic.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
Found link to world-war-2.html. Adding to queue
extracting content from psychology.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
Found link to world-war-2.html. Adding to queue
extracting content from quarternary-glaciation.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
Found link to world-war-2.html. Adding to queue
extracting content from texas-revolution.html
Found link to magnus-carlsen.html. Adding to queue
Found link to paleolithic.html. Adding to queue
Found link to psychology.html. Adding to queue
Found link to quarternary-glaciation.html. Adding to queue
Found link to rafael-nadal.html. Adding to queue
Found link to texas-revolution.html. Adding to queue
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
### Running Search Engine
```
(cs600) sparsh@Zitkel:~/Acad/Adv-Algo/Project$ go run search-engine.go 
Constructing the trie from the crawled data
Trie has been successfully constructed. Number of nodes in the trie: 6587

Enter the keywords to search through the pages: Ice, Global, warming,HISTORY;

keyword "ice" found in: [paleolithic.html,quarternary-glaciation.html,]
keyword "global" found in: [paleolithic.html,quarternary-glaciation.html,world-war-2.html,]
keyword "warming" found in: [paleolithic.html,quarternary-glaciation.html,]
keyword "history" found in: [rafael-nadal.html,magnus-carlsen.html,psychology.html,quarternary-glaciation.html,world-war-2.html,stockfish.html,]

Relevant docs: [quarternary-glaciation.html]

```