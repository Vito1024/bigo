package datastructure

/*
In golang, a string is represented as a StringHeader in runtime, which is defined as below:
	type StringHeader struct {
		Data  uintptr
		Len   int64
	}
From the definition, we know that we can get the length of any string in O(1) time, which is unlike C.
On the other hand, Go provides lots of protections to avoid problems that may occur in C
So, the conclusion is: we use go-string type as our bigo-string type directly.
*/
