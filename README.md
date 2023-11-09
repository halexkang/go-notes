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

## Interfaces & Generics in Go
### Type Embedding & Embedded Interfaces
### Embedded Structs

## Concurrency in Go
### Goroutines
### Channels
### Synchronization
