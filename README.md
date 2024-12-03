# URL Crawl

[![License](https://img.shields.io/github/license/rytsh/urlcrawl?color=blue&style=flat-square)](https://raw.githubusercontent.com/rytsh/urlcrawl/main/LICENSE)
[![Coverage](https://img.shields.io/sonar/coverage/rytsh_urlcrawl?logo=sonarcloud&server=https%3A%2F%2Fsonarcloud.io&style=flat-square)](https://sonarcloud.io/summary/overall?id=rytsh_urlcrawl)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/rytsh/urlcrawl/test.yml?branch=main&logo=github&style=flat-square&label=ci)](https://github.com/rytsh/urlcrawl/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/rytsh/urlcrawl?style=flat-square)](https://goreportcard.com/report/github.com/rytsh/urlcrawl)

Get a list of URLs from a given URL.

> Download data to destination directory.  
> If file already exists, skip download again.
> URI if different from the given URL will be skipped.

## Usage

```sh
urlcrawl -d "./testdata" <url>
```
