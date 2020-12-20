[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)
[![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](https://GitHub.com/Naereen/ama)
[![made-for-VSCode](https://img.shields.io/badge/Made%20for-VSCode-1f425f.svg)](https://code.visualstudio.com/)
[![GitHub forks](https://img.shields.io/github/forks/saswatamcode/grpc_using_go?)](https://GitHub.com/saswatamcode/search_engine/network/)
[![GitHub stars](https://img.shields.io/github/stars/saswatamcode/grpc_using_go?)](https://GitHub.com/saswatamcode/search_engine/stargazers/)
[![GitHub issues](https://img.shields.io/github/issues/saswatamcode/grpc_using_go.svg)](https://GitHub.com/saswatamcode/search_engine/issues/)
[![Open Source Love svg1](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

# Quote Search Engine
This is a search engine built using [Go](https://golang.org/) and [Elasticsearch](https://www.elastic.co/elastic-stack).

<div align="center">
	<img width="80%" src="screenshots/Screenshot-1.png">	
</div>

## How it works
It functions by scraping quotes from [Goodreads](https://www.goodreads.com/quotes/) and indexing the data into elasticsearch. With elasticsearch's robust text analysis capabilities, the data can then be queried. Here, we simply query based on the content of a quote or the name of the author. This application has both CLI and web UI for search and indexing operations.

## Dependencies
- github.com/gocolly/colly v1.2.0: For web scraping
- github.com/olivere/elastic/v7 v7.0.22: For interacting with elasticsearch
- github.com/spf13/cobra v1.1.1: For CLI
- github.com/sirupsen/logrus v1.2.0: For server logs
- Next.js v10 and Tailwind v2: For frontend web app

## Get Started
- Clone into repo
- Install elasticsearch and start service. Using [homebrew](https://brew.sh/) simply run,
```
brew install elasticsearch
elasticsearch
```
- Using CLI, index data by, (for indexname and number of quotes to be indexed you can pass in flags) 
```
go run main.go index
```
- Search with CLI,
```
go run main.go search Mark Twain
```
OR
- Search with UI, by first starting server,
```
go run main.go server
```
- Then start frontend,
``` 
cd web-ui
yarn
yarn start
```
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-typescript.svg)](https://forthebadge.com)
