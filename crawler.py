import os
from pathlib import Path
import sys
import queue
from bs4 import BeautifulSoup

import json


raw_pages = 'raw-pages'

working_dir = Path(__file__).resolve().parent / 'working-dir'


# text transformation functions
alphanum = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
num = '0123456789'
stop_words = ['a', 'an', 'the', 'aboard', 'elsewhere', 'such','about','above','across','after','against','along','amid','among','anti','around','as','at','before','behind','below','beneath','beside','besides','between','beyond','but','by','concerning','considering','despite','down','during','except','excepting','excluding','following','for','from','in','inside','into','like','minus','near','of','off','on','onto','opposite','outside','over','past','per','plus','regarding','round','save','since','than','through','to','toward','towards','under','underneath','unlike','until','up','upon','versus','via','with','within','without', 
              "you","i","he","she","it","we","they","my","your","his","her","its","our","their","this","that","these","those","who","whom","which","what","whose","all","any","each","every","none","some","anybody","anyone","anything","other","another","myself","yourself","himself","herself","itself","ourselves","themselves","is","also","can","be","or","and"]

def transform_alpha_numeric(c):
  if c not in alphanum:
    return " "
  return c

def transform_remove_stop_words(text):
  return [w for w in text if w not in stop_words]

def extract_info(doc_name, doc_id, res):
  files_dir = Path(__file__).resolve().parent / raw_pages
  doc_file_path = files_dir / doc_name
  if not os.path.exists(doc_file_path):
    print(f'doc file {doc_file_path} does not exist')
    exit(1)
  
  with open(doc_file_path) as f:
    soup = BeautifulSoup(f.read(), 'html.parser')
    soup = soup.find_all("div", {"class": "row"})[0]
    text = soup.get_text()
    text = transform_remove_stop_words(''.join([transform_alpha_numeric(c).lower() for c in text]).split())
    res['doc_ids'][doc_id] = doc_name
    
    # create frequency map
    fq_mp = {}
    for token in text:
      if token not in fq_mp:
        fq_mp[token] = 0
      fq_mp[token] += 1
    
    
    
    for token, freq in fq_mp.items():
      if token not in res['data']:
        res['data'][token] = []
      res['data'][token].append([doc_id, freq])


def main():
  
  if len(sys.argv) != 2:
    print('name of the file to start crawling from is missing')
    exit(1)
  
  start_from = sys.argv[1]
  
  files_dir = Path(__file__).resolve().parent / raw_pages
  
  start_file_path = files_dir / start_from
  
  if not os.path.exists(start_file_path):
    print('start file is invalid / doesn\'t exist')
    exit(1)
  
  
  # Crawl
  print(f'Crawling from {sys.argv[1]}')
  q = queue.Queue()
  visited = set()
  res = {
    'doc_ids': {},
    'data': {}
  }
  curr_doc_id = 0
  
  q.put(start_from)
  
  while not q.empty():
    doc = q.get()
    if doc not in visited:
      visited.add(doc)
      print(f'extracting content from {doc}')
      extract_info(doc, curr_doc_id, res)
      
      # extract links from current doc to crawl
      doc_file_path = files_dir / doc
      if not os.path.exists(doc_file_path):
        print(f'doc file {doc_file_path} does not exist')
        exit(1)
      with open(doc_file_path) as f:
        soup = BeautifulSoup(f.read(), 'html.parser')
        for link in soup.find_all('a', href=True):
          print(f"Found link to {link['href']}. Adding to queue")
          q.put(link['href'])
      
      curr_doc_id += 1
    
  
  with open(working_dir/'output.json', 'w') as f:
    f.write(json.dumps(res, indent=2))
  
  print('finished indexing all information')
  

if __name__ == "__main__":
  main()