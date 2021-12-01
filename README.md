# Advent of Code

Attempting solutions in a mix of languages to see how I like them.

# Golang Overview Notes

## Installation

Lovely brew

`brew install go`

### Docs

* https://golang.org/doc/tutorial/
* https://gobyexample.com/reading-files
* https://tour.golang.org/

### Day 1 Puzzle:

1. Package and set up
      * Easy enough to follow
      * Package management - https://pkg.go.dev/rsc.io/quote
      * Create module via `go mod init advent`
      * Creates a `go.mod` file, declares Go version and any packages
      * `go run filename.go` installs any packages for you then runs the code
      * Need to delve deeper in Go version management to see how they solve this problem
      * checksum generated - assume commit this in same manner as other package manager lock files.
2. Var assignment:

    Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

    Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

    Little confusing between := and = at first.

3. Flow Control - use range for this use case

4. Map vs map confusion, coming from node world expect native functional expressions such as map, filter etc. but doesn't look like thats supported OOTB. From the docs:

    *We often need our programs to perform operations on collections of data, like selecting all items that satisfy a given predicate or mapping all items to a new collection with a custom function.*

    *In some languages it’s idiomatic to use generic data structures and algorithms. Go does not support generics; in Go it’s common to provide collection functions if and when they are specifically needed for your program and data types.*

5. **Decision** ditch map for now, can investigate later. Use range and convert as we go, not as clean as have to do at every interation

6. For loops with in each other for solutions, on part 2 takes a while to execute.

7. Operators multiplication * </br> Comparator ==

8. String interpolation **just works** very nice. Assume theres another form to do positional args but not looked into that yet.

9.  Up next:

    Tests, better code organisation, more complex control flow, maps,  methods, pointers, structs and more :raised_hands:


### Day 2 Puzzle

1. install go language supprot VS Code. Installs to `Tools environment: GOPATH=/Users/rankers/go
Installing 16 tools at /Users/rankers/go/bin in module mode.
  gocode`

1. Array destructuring. gGneral packing/unpacking as is done in Python is not supported. I think the way to go there is to define your own small ad-hoc function using multiple return values. 
2. Enums

