# Go Notes

## What Go feels like:
- easy to use like python
- C performance & control
- garbage collection like  Java but faster since no VM
- extra builtin concurrency tools
- cool error handling
- mix of functional & OOP paradigms?

## Cool Things in Go
### Slices
### Pointer Recievers
- methods with pointer receivers can modify the value to which the receiver points
- can act like object member functions
```
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
```
### iota
- instead of manually assigning sequence of numbers, automatically assign values with keyword `iota`
### Variadics
- allows functions to be called with any number of arguments
- within the function, treat variadic as slice
```
// definition
func myFunc(nums ...int)
// Ex.1 pass multiple params
myFunc(1, 2, 3)
// Ex.2 pass slice as param
myFunc(nums...)
```
### Defer
- defer statement defers the execution of a function until the surrounding function returns
- can be put near cleanup / error handling code
- can make functions with multiple returns look clean
### Interfaces
- interfaces are implicit in Go, and can be used to decouple things
```
type Abser interface {
	Abs() float64
}
```
### Generics
- generics can take different types without sacrificing type safety
- can be used to reduce duplicate code
```
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

## Concurrency in Go
### Goroutines
### Channels
### Synchronization
