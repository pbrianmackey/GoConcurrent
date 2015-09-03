#Concurrent programming in Go

An introduction and sandbox for programming in [GoLang](https://golang.org/).

Go is a C like language.  

- Has garbage collection.  
- ~25 keywords.  
- Compiled language
- Statically typed
- Supported type inference
- functions are first class citizens
- Go is case sensitive
- code is mostly written in lower case, see standards: https://golang.org/doc/code.html
- ***Exported names need a capital letter***

Go source is available on github and locally when you install go.

###Folder structure

Go requires a "workspace". This is a root namespace for apps that you will develop. This workspace contains 3 folders

- package (shared libraries)
- source (your source code)
- bin (compiled stuff)

###Commands

- go env (display environment variables)

###Go Routine
Light weight threads managed by the go runtime.  To run a function in a go routine just put go in front of the function call.  

###Channels

Channels are mediums for synchronization and communication between go subroutines.  Kind of like a pipe.

##References
https://golang.org/  
https://vimeo.com/69237265  
http://www.pluralsight.com/courses/go-fundamentals  

##Cool Stuff
https://www.docker.com/  
http://kubernetes.io/ (Orchestrator for docker!)
https://coreos.com/
https://www.openstack.org/
