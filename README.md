# go_utils
some utility functions for go codebases


## slices

### Find
Finds item on a slice at first matches on prediction and return it's pointer
```go
func Find[T <itemType>](*[]T, ComporablePrediction[T])
```
```go
s := []struct{
    name string
}{
    {
        name: "earth",
    },
    {
        name: "jupiter",
    },
    {
        name: "venus",
    },
}

item := Find(&s, func(item struct{ name string }, index item) bool {
    return item.name == "jupiter"
})


fmt.Println(*item)
// { jupiter }

```

### Filter

```go

s := []struct{
    name string
}{
    {
        name: "earth",
    },
    {
        name: "jupiter",
    },
    {
        name: "venus",
    },
}

filtered := Filter(&s, func(item struct{ name string }, index item) bool {
    return strings.Contains(item.name, "t")
})
fmt.Println(*filtered)

// {{earth} {jupiter}}
```