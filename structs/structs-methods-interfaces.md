# Structs in Go
- structs are "passed by value" by default 

# Methods in Go
- structs can be "extended" to add methods
- works like type in TypeScript, BUT you define your own "methods" SEPARATELY:
```go
// in Go you do this:
type Circle struct {
  Radius float64
}

// already have the logic
func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}
```

```typescript
// in TypeScript you do this:
type Circle {
  Radius: number
  Area: () => void;
}

function something(...args): Circle {
  Area = k
}
```

# Interfaces in Go
