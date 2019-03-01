# GoLangStudy


## Go Tutorial 
 - [An Introduction to Go(CERN)](https://speakerdeck.com/campoy/an-introduction-to-go-cern)

## Learning Go - Miek Gieben
 - [Learning go by Miek Gieben](https://www.miek.nl/go/) : this site is good for quick speedup of go programming experience.
 
### Check go builtin first
 - $ godoc builtin | more
 - $ godoc net/http | more
 
### Go tool compile
```
 $ go tool compile unsafe.go  (object file 생성)
 $ go tool compile -pack unsafe.go (archive file 생성)
 $ go tool compile -race unsafe.go (race condition 감지)
 $ go tool compile -S unsafe.go (assembly)
 $ go tool compile -W hello.go (node tree 생성)
 
```

### Got tool trace
- example runtime_trace.go file for trace
```
package main

import (
	"os"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan int)

	go func() {
		ch <-42
	}()
	
	<-ch

```

-executing go tool trace 
```
 $ go run runtime_trace.go 2>trace.out
 $ go tool trace trace.out

```


## Go Lang Tutorial Site
 - [learn go tutorials](https://stackify.com/learn-go-tutorials/)
 
## Go : the Good, the Bad and the Ugly
 - [the Good, the Bad and the Ugly](https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/)

## Go best practice
 - [Twelve Go Best Practices](https://talks.golang.org/2013/bestpractices.slide#1)
 - [Andrew Gerrand’s idiomatic Naming Convention](https://talks.golang.org/2014/names.slide#1)
 - [Peter Bourgon's Go best practices, six years in](https://peter.bourgon.org/go-best-practices-2016/)
 - [Phil Estes's 4 tips for learning Golang](https://opensource.com/article/18/11/learning-golang)

## Go Editor 
 -  Visual Studio Code + [vscode-go](https://github.com/Microsoft/vscode-go)

## Build Web Application with Golang
 - [Git book for Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/en/)

### Golang Programs
 - [Golan Programs](http:/www.golangprograms.com/)

### Go Coding Convention
 - [Unknow/go-code-convention](https://github.com/Unknwon/go-code-convention/blob/master/en-US/README.md) : For production level coding 


## Go Example Program

### Go TCP Communication
 - [Go TCP communication](https://blog.kesuskim.com/2018/08/go-tcp-implementation/)

### Go chat bot
 - [Go:code that grows with grace](https://talks.golang.org/2012/chat.slide#1)

### Go for Pythonistas
 - [Go vs Python](https://talks.golang.org/2013/go4python.slide#1)

### Advanced Go Concurrency
 - [Advanced Go Concurrency](https://talks.golang.org/2013/advconc.slide#1)

## Debugging Go Code with GDB
 - [gdb](https://golang.org/doc/gdb)

### Garbage collection 
```
 $GODEBUG=gctrace=1 go run gColl.go
```

### Go test and coverage
 - [Go cover](https://blog.golang.org/cover)

## Go Compiler
 - [Rob Pike's Lexical Scanning in Go](https://talks.golang.org/2011/lex.slide#1)

### Go Parser (goyacc)
 - [PlanetScale's CTO on go parser](https://www.youtube.com/watch?v=NG0s3-s3whY)
 - [GopherCon 2018 - How to Write a Parser in Go](https://about.sourcegraph.com/go/gophercon-2018-how-to-write-a-parser-in-go)
 - [GopherAcademy : Parsing with ANTLR 4 and Go](https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/)
 - [GopherAcademy : Handwritten Parsers & Lexers in Go](https://blog.gopheracademy.com/advent-2014/parsers-lexers/)


### Go Assembly 
 - [Quick guide to assembly](https://golang.org/doc/asm)
 - [Rob Pike's talk on go assembly](https://talks.golang.org/2016/asm.slide#1)
 - [go asm complete reference](https://quasilyte.github.io/blog/post/go-asm-complementary-reference/)
 - [SIMD optimization in go](https://goroutines.com/asm)
 - [go assembly related blog](https://quasilyte.github.io/blog/)
 - [go Compiler Explorer](https://go.godbolt.org/)
 - assembly method :  $go tool compile -S <go_src.go>

### Go Logging with uber zap
 - [golang log library zap](https://github.com/golangkorea/gophercon-talks/blob/master/slides/201901/golang-log-library-zap.pdf)

## Go runtime 
### Go Garbage Collection Algorithm
 - tricolor mark-and-sweep algorithm (GO GC algorithm)
 - On-the-fly Garbage Collection: An Exercise in Cooperation 논문에서 처음 발표됨 
### Go Garbage Collection

check runtime Garbage Stats implementation
```go
 printStats(mem runtime.MemStats) {
   runtime.ReadMemStats(&mem)
   ...
 }
```

## Go module
 - [Go module roadmap](https://blog.golang.org/modules2019) 
 
### Go module by example
 - [Go module example](https://github.com/go-modules-by-example/index)

### Go Unit Test by GoConvey
 - [Introduction to GoConvey](https://www.youtube.com/watch?v=wlUKRxWEELU)

## Go Library

### EVCache
 - [Netflix EVCache Architecture](https://docs.google.com/presentation/d/1hiwj0oV0nLMtFmARpoFLDoGNh-YdCdy9epTagmwUIDE/edit#slide=id.g1505988a49_0_0)

## Go Rest API Tutorial
 
### Go Restful API tutorial
 - [Codementor Go Restful API handling](https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo)
 - [Go Restful API server 개발기](https://www.slideshare.net/Hyejong/golang-restful)
 - [7 Go Rest API framework Introduction](https://nordicapis.com/7-frameworks-to-build-a-rest-api-in-go/)

### Golang Restful API using GORM and Gorilla Mux
 - [RESTful Web API](http://www.golangprograms.com/advance-programs/golang-restful-api-using-grom-and-gorilla-mux.html)

### Go gRPC
 - [gopher academy on gRPC](https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/)

### Go Microservice
 - [Go Microservice Blog Series](http://callistaenterprise.se/blogg/teknik/2017/02/17/go-blog-series-part1/)

### Gorm for Postgresql 
 - [Managing Data in Golang Using Gorm](http://blog.tamizhvendan.in/blog/2017/07/23/managing-data-in-golang-using-gorm-part-1/)
 
### Go 2.0 
 - [Go 2.0 and Module](http://callistaenterprise.se/assets/presentationer/cadec-2019-go.pdf)

## Awesome Go
 - [Awesome go](https://awesome-go.com/)
