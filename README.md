# Introduction

A convenience package to measure wall clock execution time for of go code.

# Usage and API

```go
func MysteriouslySlowFunc() {
     timeit.G.Trace() // Records and prints the entry and exit time for the MysteriouslySlowFunc

     timeit.G.Record("before an action")
     AnAction()
     timeit.G.Record("after an action") // prints the time taken by AnAction()
}
```

# Sample output