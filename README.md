## Bigo: An Efficient Key-Value Storage Server


### Overview

#### It supports four kinds of data structures:
    - string
    - list
    - set
    - hashtable
    
As the name implies, it’s a key-value storage server, which means that Bigo is a hashtable at the top level. Its key is always of string type, and the value could be a string, list, set, or hashtable.


### How to implement
It’s written in Golang, which is an elegant and efficient programming language. Golang supports many data structures. It has an implementation of a hashtable called a map, which is built into the language itself as a basic type, so here, we just reuse its map as the implementation of Bigo’s hashtable. Meanwhile, Golang also supports string types. It implements strings as read-only byte slices, so we won’t discuss it here. As a result, we need to implement the list and set types with Golang.

### How to use
Run the build.sh script, and you’ll get two runnable files: client and server, at $HOME/bin directory. Start the server first, and then the client. The client will connect to the server.
Enjoy!
