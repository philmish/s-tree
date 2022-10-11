# s-tree - a simple tree DS in go

![Tests](https://github.com/philmish/s-tree/actions/workflows/go.yml/badge.svg)

## Introduction
To learn about tree data structures I created a simple tree and
a radix-tree in go. As an example, I implemented them in a very basic
key-value based in-memeory storage application, which provides a server and
a client, which communicate over unix domain sockets.

## Roadmap

- [ ] refactor tree pkg for better access to stored values
- [ ] typed nodes for storing other data types than strings in kvdb
    - [ x ] create TypedNode struct
    - [ x ] create StrNode struct
    - [ x ] create IntNode struct
- [ ] data persistance to file and data loading from file 
- [ ] benchmarking
- [ ] docs
