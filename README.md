# s-tree - a simple tree DS in go

![Tests](https://github.com/philmish/s-tree/actions/workflows/go.yml/badge.svg)

## Introduction
To learn about tree data structures I created a simple tree and
a radix-tree in go. As an example, I implemented them in a very basic
key-value based in-memeory storage application, which provides a server and
a client, which communicate over unix domain sockets.

## Roadmap

- [x] typed nodes for storing other data types than strings in kvdb
    - [x] create TypedNode type
    - [x] create StrNode type
    - [x] create IntNode type
    - [x] create StrSliceNode type
    - [x] create IntSliceNode type
    - [x] create BoolSliceNode type
    - [x] create StrStrMapNode type
    - [x] create StrIntMapNode type
    - [x] create StrBoolMapNode type
    - [x] create a TypedTree type
- [ ] refactor kvdb
    - [x] implement typed radix tree
    - [ ] refactor for use of typed tree
    - [ ] re-implement querying the kvdb
- [ ] data persistance for kvdb to file and data loading from file 
    - [ ] implement encoding db data to JSON
    - [ ] implement loading db data from JSON
- [ ] benchmarking
- [ ] docs
