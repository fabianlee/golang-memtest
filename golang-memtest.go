package main

import (
 "fmt"
 "os"
 "strconv"
 "runtime"
 "math/rand"
 "time"
 "os/signal"
 "syscall"
)

func main() {

    // alert to any OS signals sent while running
    CatchOSSignals()

    // get number of Mb to allocate from param (default=1)
    nmb := ReadEnvOrArgs(1,"nmb","1")
    // get number of milliseconds to wait between 1Mb allocations (default=100)
    nms := ReadEnvOrArgs(2,"nms","100")


    // show initial usage 
    PrintMemUsage()
    fmt.Printf("Asked to allocate %dMb\n\n",nmb)

    // allocate memory 1Mb at a time
    rand.Seed(time.Now().UTC().UnixNano())
    var resarr = make([][]byte,nmb)
    for i:=0; i<nmb; i++ {
      resarr[i] = make([]byte, 1024*1024)
      // populate array so it takes up memory
      // if this is not done, it will not fill up memory space
      rand.Read(resarr[i])
      PrintMemUsage()
      time.Sleep( time.Duration(nms) * time.Millisecond)
      //fmt.Printf("Total allocated: %dMb\n",i+1)
    }
    fmt.Printf("\n")

    // show final usage
    PrintMemUsage()
    fmt.Printf("SUCCESS allocating %dMb\n",len(resarr))
}


// Credit: https://golangcode.com/print-the-current-memory-usage/
func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

// prove that docker container kill for OOM is abrupt, no OS signals to catch
func CatchOSSignals() {
    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc,
        syscall.SIGHUP,
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGQUIT)
    go func() {
        s := <-sigc
        fmt.Printf("SIGNAL from OS!!!!!!!!!!!!")
        fmt.Println(s)
    }()
}

// finds the 'nmb' parameter in either command line param or OS env
// defaults to 1 if no values found
func ReadEnvOrArgs(posIndex int,pname string,defaultString string) int {

    nmbstr := defaultString
    //fmt.Printf("str arg default. going to look at index %d, else using %s\n",posIndex,defaultString)
    if len(os.Args)>posIndex {
      nmbstr = os.Args[posIndex]
    }else if len(os.Getenv(pname))>0 {
      nmbstr = os.Getenv(pname)
    }
    //fmt.Printf("final str arg: %s\n",nmbstr)

    nmb,err := strconv.Atoi(nmbstr)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    return nmb
}
