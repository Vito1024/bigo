## Bigo: An efficient key-value storage server

### Overview
#### It supports four kinds of data structures:
    - string
    - list
    - set
    - hashtable
    
As the name applied, it's a key-value storage server, which means that bigo is a hashtable on the top level. Its key is
always a string type, value could be string, list, set and hashtable.

### How to implement
It's written in golang, which is an elegant and efficient programming language.
Golang supports lots of data structures. It has an implementation of hashtable called map which is built into the
language itself as a basic type, so here, we just reuse its map as the implementation of bigo's hashtable.
Meanwhile, golang as a common language also supports string type, it implements string as a read-only byte slice,
we'll not talk about it here. As a result, we need to implement list and set type with golang.


### How to use