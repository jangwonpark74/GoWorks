# GoLangStudy

## Go Introduction
 - [go 언어 이렇게 배워 보세요!](http://www.bloter.net/archives/230851)

## Go Tutorial 
 - [An Introduction to Go(CERN)](https://speakerdeck.com/campoy/an-introduction-to-go-cern)
 - [Understanding Go Map](https://www.digitalocean.com/community/tutorials/understanding-maps-in-go)
 
## Learning Go - Miek Gieben
 - [Learning go by Miek Gieben](https://www.miek.nl/go/) : this site is good for quick speedup of go programming experience.


## Go for beginners 
 - [How do beginners Learn go](https://changelog.com/gotime/85)

## Uber Go Style Guide
 - [Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

## Go programming with errors
 - [pogramming with errors](https://peter.bourgon.org/blog/2019/09/11/programming-with-errors.html)
 - [Go errors in 1.13](https://blog.golang.org/go1.13-errors)
## I never use in go
 - [Never use guide in go](https://www.youtube.com/watch?v=5DVV36uqQ4E)

## Go gotchas
 - [50 common gotchas in go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)
 - [bad go with slice](https://philpearl.github.io/post/bad_go_not_sizing_arrays/)
 - [bad go slice of pointers](https://medium.com/@philpearl/bad-go-slices-of-pointers-ed3c06b8bb41)

## Go real world concurrency bug
 - [Go concurrency bug study](https://medium.com/a-journey-with-go/go-concurrency-bugs-in-go-7d3677a1f2a2)
 
## Go expert to follow in online
 - [Go experts](https://blog.newrelic.com/technology/golang-experts-follow-online/)

### Check go builtin first
 - $ godoc builtin | more
 - $ godoc net/http | more

## Go 1.12 upgrade (ubuntu)
 - [go 1.12 upgrade ](https://github.com/golang/go/wiki/Ubuntu)

## Go 1.13.2 Release Note
 - [go 1.13.2 Release and bug fix](https://groups.google.com/forum/#!topic/golang-nuts/uQlAyLkZ5kg)

### Go tool compile
```
 $ go tool compile unsafe.go  (object file 생성)
 $ go tool compile -pack unsafe.go (archive file 생성)
 $ go tool compile -race unsafe.go (race condition 감지)
 $ go tool compile -S unsafe.go (assembly)
 $ go tool compile -W hello.go (node tree 생성)

```
### Go race detector 
 - [Introducing the Go race detection tool](https://blog.golang.org/race-detector)
 
### Go Interpretor  
- [Go AST](http://notes.eatonphil.com/interpreting-go.html)


### Go dissembly
```
 $ go build -o add.a  
 $ go tool objdump add.a
```

### Go binary
 - [dissecting go binaries](https://www.grant.pizza/dissecting-go-binaries/)

#### Intel Assembly Tutorial
 - [x86 assembly](https://www.nayuki.io/page/a-fundamental-introduction-to-x86-assembly-programming#13-compared-to-other-architectures)

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
}
```

-executing go tool trace 
```
 $ go run runtime_trace.go 2>trace.out
 $ go tool trace trace.out

```
### Go App Tracing
- [DataDog Go App Tracing](https://www.datadoghq.com/dg/apm/go-tracing/?utm_source=Advertisement&utm_medium=CooperPress&utm_campaign=CooperPress-GoWeeklyS)


### Go profiling

#### CPU profiling

[CopherCon2019 : Dave Cheney - Two Go programs, Three Different Profiling Techniques](https://www.youtube.com/watch?v=nok0aYiGiYA&list=PL2ntRZ1ySWBdDyspRTNBIKES1Y-P__59_&index=9)

```

func main() {

    defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()  // Add this line to programe 
   
    f, err := os.Open(os.Args[1])
    if err != nil {
     ...
    }

}

:Goimports 

<< after execution >>

$ go tool pprof cpu.pprof

$ go tool pprof -http=:8080 cpu.pprof

 

```

#### Memory profiling

```

func main() {

    defer profile.Start(profile.MemProfile, profile.MemProfileRate(1),profile.ProfilePath(".")).Stop()  // Add this line to programe 
   
    f, err := os.Open(os.Args[1])
    if err != nil {
     ...
    }

}

:Goimports 

<< after execution >>

$ go tool pprof mem.pprof

$ go tool pprof -http=:8080 mem.pprof
 

```

### Memory Leak Profiling 

- [Memory leak in Microservice](https://blog.detectify.com/2019/09/05/how-we-tracked-down-a-memory-leak-in-one-of-our-go-microservices/)
- [Go 101 memory leak scenario](https://go101.org/article/memory-leaking.html)

Most of the leak in go program 
```
– Creating substrings and subslices.
– Wrong use of the defer statement.
– Unclosed HTTP response bodies (or unclosed resources in general).
– Orphaned hanging go routines.
– Global variables.

```
## Go adaptor
 - [ Go Mock Test](https://itnext.io/yet-another-tool-to-mock-interfaces-in-go-73de1b02c041)

## Go Test Coverage

 - [Instrumentation in Go ](https://medium.com/a-journey-with-go/go-instrumentation-in-go-e845cdae0c51)

## Go Parallel Programming
 - [Parallel processing in go with WaitGroup](https://medium.com/@arindamroy/easy-guide-to-parallel-processing-in-golang-60f455fb7af3)

### Dockerizing Go service
 - [Dockerizing Go service](https://xitonix.io/containerised-go-services/)


## Go Lang Tutorial Site
 - [learn go tutorials](https://stackify.com/learn-go-tutorials/)
 
## Go : the Good, the Bad and the Ugly
 - [the Good, the Bad and the Ugly](https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/)

## Go best practice
 - [Twelve Go Best Practices](https://talks.golang.org/2013/bestpractices.slide#1)
 - [Andrew Gerrand’s idiomatic Naming Convention](https://talks.golang.org/2014/names.slide#1)
 - [Peter Bourgon's Go best practices, six years in](https://peter.bourgon.org/go-best-practices-2016/)
 - [Phil Estes's 4 tips for learning Golang](https://opensource.com/article/18/11/learning-golang)

## Go style guide
 - [practical-go](https://dave.cheney.net/practical-go/presentations/qcon-china.html)

## Go Editor 
 -  Visual Studio Code + [vscode-go](https://github.com/Microsoft/vscode-go)

## Go CLI 
 - [build an awesome cli app in go](https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/)


## Build Web Application with Golang
 - [Git book for Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/en/)
### go web examples
 - [Go Web examples](https://gowebexamples.com/)

### Golang Programs
 - [Golang Programs](http:/www.golangprograms.com/)

### Golang string tips
 - [40+ practical string tips](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/)

### Go Utility
 - [Go-funk](https://github.com/thoas/go-funk)

### Go Coding Convention
 - [Unknow/go-code-convention](https://github.com/Unknwon/go-code-convention/blob/master/en-US/README.md) : For production level coding 

### State of Go 
 - [What's new since Go 1.10](https://speakerdeck.com/campoy/the-state-of-go-feb-2019)

### Robustness of Go
 - [Robustness of Go](https://www.youtube.com/watch?v=40d26ZGfhR8)

### Go performance pattern 
  - [Go performance pattern](https://stackimpact.com/docs/go-performance-tuning/#go-performance-patterns)
  
### Go train company
 - [ArdanLabs](https://www.ardanlabs.com/ultimate-go/)

## Go Example Program

### Streaming IO in Go
 - [Learning Streaming IO in go](https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185)

### Production Grade UDP streaming
 - [kcp go library for UDP](https://github.com/xtaci/kcp-go)

### Go TCP Communication
 - [Go TCP communication](https://blog.kesuskim.com/2018/08/go-tcp-implementation/)

### TLS1.3 for Go
 - [PKG for Gophers](https://about.sourcegraph.com/go/gophercon-2019-pki-for-gophers)

### Quic for Go
 - [QUIC-Go](https://github.com/lucas-clemente/quic-go)

### Go chat bot
 - [Go:code that grows with grace](https://talks.golang.org/2012/chat.slide#1)

### Go for Pythonistas
 - [Go vs Python](https://talks.golang.org/2013/go4python.slide#1)

### Go code generation
 - [Code generation using template](https://www.youtube.com/watch?v=3llI65DQB_w)

### Context 
 - [Context and Cancellation of goroutines](http://dahernan.github.io/2015/02/04/context-and-cancellation-of-goroutines/)


### Advanced Go Concurrency
 - [Advanced Go Concurrency](https://talks.golang.org/2013/advconc.slide#1)
 - [Visualizing Go Concurrency](https://divan.dev/posts/go_concurrency_visualize/)
 - [Visualizing Go Concurrency(YouTube)](https://www.youtube.com/watch?v=KyuFeiG3Y60)
 - [3D gotrace tool](https://github.com/divan/gotrace)
 - [Go 로 1M request 처리하기](http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/)


## Debugging Go Code with GDB
 - [gdb](https://golang.org/doc/gdb)


## Go refactoring

 - [How to fix tightly coupled go program](https://www.sage42.org/2019/01/30/how-to-fix-tightly-coupled-go-code/)

## Go Performance
 - [So You Wanna Go Fast? by Tyler Treat](https://www.youtube.com/watch?v=DJ4d_PZ6Gns&utm_source=golangweekly&utm_medium=email)

### Go scheduler
 - [Theory behind go-scheduler](https://povilasv.me/go-scheduler/)
 - [Go runtime scheduler](http://www.cs.columbia.edu/~aho/cs6998/reports/12-12-11_DeshpandeSponslerWeiss_GO.pdf)

### Go Memory Management
 - [A visual guide to Go Memory Allocator from scratch (Golang)](https://blog.learngoprogramming.com/a-visual-guide-to-golang-memory-allocator-from-ground-up-e132258453ed)
 - [Avoiding Memory Leak in Golang API](https://hackernoon.com/avoiding-memory-leak-in-golang-api-1843ef45fca8)


### Garbage collection & API gateway latency
 - [API gateway performance](https://blog.twitch.tv/go-memory-ballast-how-i-learnt-to-stop-worrying-and-love-the-heap-26c2462549a2)

### Garbage collection 
```
 $GODEBUG=gctrace=1 go run gColl.go
```

### Go Memory Leak Detection
 - [Investigating memory leaks in go](https://medium.freecodecamp.org/how-i-investigated-memory-leaks-in-go-using-pprof-on-a-large-codebase-4bec4325e192)

### Go test and coverage
 - [Go Testing](https://dave.cheney.net/paste/gopherchina-2019-testing-talk.pdf)
 - [Go cover](https://blog.golang.org/cover)

### Go Performance
 - [Go Performance in GCP](https://www.youtube.com/watch?v=b0o-xeEoug0)
 - [Go Performance tuning](https://eng.uber.com/optimizing-m3/)

## Go Compiler
 - [Rob Pike's Lexical Scanning in Go](https://talks.golang.org/2011/lex.slide#1)
 - [Go LLVM](https://aykevl.nl/2019/04/llvm-from-go)
 

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
 - [Go module blog](https://blog.golang.org/using-go-modules)
 - [Go 1.13 modules 소개](https://aidanbae.github.io/code/golang/modules/)
 - [Defining Go Moudles](https://research.swtch.com/vgo-module.pdf)
 
### Go module by example
 - [Go module example](https://github.com/go-modules-by-example/index)

### Go Unit Test by GoConvey
 - [Introduction to GoConvey](https://www.youtube.com/watch?v=wlUKRxWEELU)
 - [Go runtime scheduler](http://www.cs.columbia.edu/~aho/cs6998/reports/12-12-11_DeshpandeSponslerWeiss_GO.pdf)
 
## Go Library

### EVCache
 - [Netflix EVCache Architecture](https://docs.google.com/presentation/d/1hiwj0oV0nLMtFmARpoFLDoGNh-YdCdy9epTagmwUIDE/edit#slide=id.g1505988a49_0_0)

### Jingo : A Faster JSON Encoder for Go
 - [Jingo json encoder for go](https://bet365techblog.com/open-sourcing-jingo-a-faster-json-encoder-for-go)


## Go Rest API Tutorial

### Go Restful API tutorial
 - [Go Restful API from start](https://medium.com/@phuctm97/go-restful-series-a7addbfef5b1)
 - [Codementor Go Restful API handling](https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo)
 - [Go Restful API server 개발기](https://www.slideshare.net/Hyejong/golang-restful)
 - [7 Go Rest API framework Introduction](https://nordicapis.com/7-frameworks-to-build-a-rest-api-in-go/)

### Golang Restful API using GORM and Gorilla Mux
 - [RESTful Web API](http://www.golangprograms.com/advance-programs/golang-restful-api-using-grom-and-gorilla-mux.html)

### Go gRPC
 - [gopher academy on gRPC](https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/)
 - [go GRPC](https://www.youtube.com/watch?v=RoXT_Rkg8LA)
 - [golang gRPC](https://github.com/grpc/grpc-go)

### Protocol buffer
 - [Protocol buffer Jason Mapping](https://developers.google.com/protocol-buffers/docs/proto3#json)

### Go Encrypting Streaming
 - [Streaming Cyper for Large Data](https://medium.com/blend-engineering/encrypting-streams-in-go-6cff6062a107)

### Go graphql
 - [graphql-server in google app engine](https://outcrawl.com/graphql-server-go-google-app-engine?utm_source=golangweekly&utm_medium=email)

### Go TOML
- [Go TOML](https://github.com/pelletier/go-toml)

### GoCV (Computer Vision)
 - [gocv install](https://gocv.io/getting-started/linux/)
 
### Go XMPP Server
 - [Go XMPP server:jackal](https://github.com/ortuman/jackal)

### Go Microservice
 - [Go Microservice Blog Series](http://callistaenterprise.se/blogg/teknik/2017/02/17/go-blog-series-part1/)
 - [Go Microservice Toolkit from New York Times](https://github.com/nytimes/gizmo) 
  
### Open source go 배송조회 시스템
 - [Go 배송 조회](https://subicura.com/2016/06/13/start-go-shipment-tracking-opensource.html)

### Gorm for Postgresql 
 - [Managing Data in Golang Using Gorm](http://blog.tamizhvendan.in/blog/2017/07/23/managing-data-in-golang-using-gorm-part-1/)
 
### Go 2.0 
 - [Go 2.0 and Module](http://callistaenterprise.se/assets/presentationer/cadec-2019-go.pdf)

## Awesome Go
 - [Awesome go](https://awesome-go.com/)
