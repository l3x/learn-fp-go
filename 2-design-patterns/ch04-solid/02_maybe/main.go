package main

import (
    "fmt"
    "maybe"
)

type Work struct { Succeeded bool }

func main() {
    fmt.Println("Has value:", maybe.SomeString("Hi"))
    fmt.Println("Is empty :", maybe.EmptyString())
    fmt.Println()
    printResults(Work{Succeeded:true})
    fmt.Println()
    printResults(Work{Succeeded:false})
}

func runTask(success bool) maybe.Either {
    if success {
        return maybe.EitherSuccess("I succeeded")
    } else {
        return maybe.EitherFailure(fmt.Errorf("ERROR: %s", "I failed"))
    }
}

func printResults(work Work) {
    maybeWorked := runTask(work.Succeeded)
    fmt.Printf("Either: %+v\n", maybeWorked)
    fmt.Println("maybeWorked.Succeeded:",  maybeWorked.Succeeded())
    fmt.Println("maybeWorked.Failed:", maybeWorked.Failed())

    fmt.Println("maybeWorked.Succeeded().Val():", maybeWorked.Succeeded().Val())
    fmt.Printf("maybeWorked.Failed().Err(): %v\n", maybeWorked.Failed().Err())
}
