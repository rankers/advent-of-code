# Golang Overview Notes

## Installation

Lovely brew

## First program

### Docs

* https://gobyexample.com/reading-files
* https://tour.golang.org/

### Basics:


1. Package and set up
      * Easy enough to follow
      * Package management - https://pkg.go.dev/rsc.io/quote
      * Go binary version management 
      * `go mod init advent`
      * `go run filename.go` installs for you then runs
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