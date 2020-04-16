# -*- coding: utf-8 -*-
import os
import pandas as pd
from pycorenlp import StanfordCoreNLP
import json
nlp = StanfordCoreNLP('http://localhost:9000')

###testing the annotater is working/ can resolve quotes###
texty = '"Is it working again now" asked Laura. "Yeah I reckon so", said Tom'

testy = nlp.annotate(texty, properties = {'annotators': 'tokenize,ssplit,pos,lemma,ner,parse,depparse,coref,quote',
                                          'outputFormat': 'text'})

###importing data###
testoneday = pd.read_csv('Jan14_articles_clean.csv', sep = ',')

###looking at data###
testoneday.info(verbose = False)
print(testoneday.head(3))
testoneday.columns
print(testoneday["content"].head(3))
articles = testoneday
articlebody=testoneday["content"].head(1)

###successful singular article test
##this is called greedy test because i was playing around with the coref greedyness
##greedyness is a measure of how many coref connections the model wants to make
#https://stanfordnlp.github.io/CoreNLP/coref.html


greedyarticletest = nlp.annotate(articlebody[0], properties = {'timeout':60000,'annotators': 'tokenize,ssplit,pos,lemma,ner,parse,coref',
                                                               'coref.maxdist': -1,
                                                               'coref.algorithm': 'neural',
                                                               'coref.neural.greedyness': .7,
                                                               'outputFormat': 'json'})

#####iterating over a number of articles
results = []
for i in range(len(articles)):
    results.append(nlp.annotate(articles["content"][i], properties = {'timeout':60000,'annotators': 'tokenize,ssplit,pos,lemma,ner,depparse,coref',
                                                                      'outputFormat': 'json'}))



    ###saving

    with open ('cutler_maybenn6.json', 'w') as json_file:
        json.dump(greedyarticletest4,json_file)


    results.to_csv("quoteytod.csv")



