# Go Cyclic Handler

A minimalist go cyclic handler

## Examples

Import package

```go
    import (
        "github.com/Lucas32-dev/go-cyclic-handler"
    )
```

Create new cycle

```go
    handler := func (t time.Time) {
        log.Println("called at ", t)
        // do something
    }

    c := cycle.New(
        handler,        // handler called every cycle
        time.Second,    // interval between each cycle
        time.Second*10, // stop cycling after 10 seconds
    )
    c.Run()
```

Change options to a existing cycle

```go
    // ... same code above

    c.ResetInterval(time.Second*4)
    c.ResetStopTime(time.Second*40)
```
