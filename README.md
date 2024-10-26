# MSDS 431: Assignment 5 - Go Web Scraper
Migus Wong\
10/27/2024

## Project Description
### Assignment Scenario
A technology firm has decided to create its own online library, a knowledge base focused on its current research and development efforts into intelligent systems and robotics. Some of the information for the knowledge base will be collected from the World Wide Web. There will be an initial web crawl, followed by web page scraping and text file parsing. 

The World Wide Web has become our primary information resource. But with the web's immense size and lack of organization, finding our way to the information we need can be a frustrating, time-consuming process, even with the best general-purpose search engines at our disposal.

Fortunately, the firm believes it can collect much of the general information it needs by web scraping Wikipedia pages. Drawing on code from Ryan (2018), they have written a simple crawler/scraper in Python using Scrapy. Unfortunately, the program runs very slowly, requiring sequential searches through ordered lists of web pages.

Hearing how fast Go is compared with Python, due largely to Go's ability to take advantage of today's multicore processors, the managers have decided to convert from Python Scrapy to Go.  Their thinking is that web crawling, scraping, and parsing can be carried out with concurrent processes across many target websites or web pages. If there are hundreds or thousands of targets to crawl, the data scientists can initiate hundreds or thousands of goroutines. Crawling, scraping, and parsing could be executed concurrently across many websites and web pages. 

### Assignment Deliverables
* Include the web address text (URL) for the GitHub repository in the comments form of the assignment posting.  The web address you provide should be the URL for the cloneable GitHub repository. It should end with the .git extension.   
* The README.md Markdown text file of the repository should provide complete documentation for the assignment.
* The repository should include a JSON lines file with the text extracted from the web pages of the crawl (the results from crawling and scraping of the Web) [use .jl extension for the JSON lines file] A JSON lines file has separate JSON objects (one for each page scraped) separated by line feeds (on macOS and Linux) or by carriage returns and line feeds (on Windows). We ask for a JSON lines file because this is a common file format for input to database or knowledge base systems.

### Discussion of Development
This web scraper was mainly built on functionality from
[Go colly](https://github.com/gocolly/colly) was the main library used to scape the wiki pages. The program works by placing extracted wiki data in a struct Type named Site. 4 fields are captured:
* name - The url of the wiki
* title - The article name of the wiki
* bodytext - Main content of the wiki.
* tags - potential key words of the article based on URL parsing.

Colly extracts the text for each field based off html ids and classes and returns the respective text for each field in the format of a json lines file. Outputs were validated through the use of https://jsonlines.org/validator/.



## Testing
Unit testing was employed via Go's standard testing library to mainy test scraping functionality, ability to remove stopwords from any tags identified. 

### Discussion of Concurrency
Timing was also conducted on how quickly the entire scraping process could occur. Total time to execute the program took around 600 - 700 ms to fully execute. This seems to be significantly quicker than the scrappy python program ran which took anywhere from a whole second to longer times. It appears that that the main attributes that could explain the difference between the two results is Go's ability for concurrency. 

In the program I wrote, each website triggers an individual go routine meaning that multiple sites are simultaneously scraped. This is evidently seen in the output file and the fact that websites are not listed in the same order as they are in the provided *urls* list.
