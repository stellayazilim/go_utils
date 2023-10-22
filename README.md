# go_utils
some utility functions for go codebases


## slices
play with slices like javascripts find, filter , map

### Find
Finds item on a slice at first matches on prediction and return it's pointer
```go
func Find[T <itemType>](*[]T, ComporablePrediction[T]) *T
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

filter items for given prediction and returns new slice

```go
func Filter[T <itemType>](*[]T, ComporablePrediction[T]) []T
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

filtered := Filter(&s, func(item struct{ name string }, index item) bool {
    return strings.Contains(item.name, "t")
})

fmt.Println(*filtered)

// {{earth} {jupiter}}
```


### Map

converts given slice to another type of slice

```go
func Map[T any, X any](slice *[]T, prediction DeterminablePrediction[T, X]) []X 
```

```go

	type name struct {
		name string
	}

	type planet struct {
		planetName string
	}

	s := []name{{ name: "mars" }, { name: "venus" }}

    var planets []planet 

	planets = slices.Map(&s, func(item name, index int) planet {
		return planet{
			planetName: item.name,
		}
	})

	fmt.Println(planets)

    // [{mars} {venus}]
```

### prediction types

```go
type CompareblePrediction[T comparable] func(el T, index int) bool
```

```go
type DeterminablePrediction[T any, X any] func(el T, index int) X
```