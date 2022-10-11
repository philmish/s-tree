# s-tree - a simple tree DS in go

![Tests](https://github.com/philmish/s-tree/actions/workflows/go.yml/badge.svg)

## Introduction
To learn about tree data structures I created a simple tree and
a radix-tree in go. As an example, I implemented them in a very basic
key-value based in-memeory storage application, which provides a server and
a client, which communicate over unix domain sockets.

## Roadmap

- [ ] typed nodes for storing other data types than strings in kvdb
    - [x] create TypedNode type
    - [x] create StrNode type
    - [x] create IntNode type
    - [x] create StrSliceNode type
    - [ ] create IntSliceNode type
    - [ ] create BoolSliceNode type
    - [ ] create StrStrMapNode type
    - [ ] create StrIntMapNode type
    - [ ] create StrBoolMapNode type
    - [ ] create a TypedTree type
- [ ] refactor kvdb
    - [ ] refactor for use of typed tree
    - [ ] re-implement querying the kvdb
- [ ] data persistance to file and data loading from file 
    - [ ] figure out how to which file type fits
    - [ ] implement transforming in-memory data to fitting format
- [ ] benchmarking
- [ ] docs
