package main

import (
 "fmt"
 "os"
 "strconv"
 "runtime"
 "math/rand"
 "time"
)

func main() {

    // get number of Mb to allocate from param
    nmb := ReadEnvOrArgs("nmb")

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


// finds the 'nmb' parameter in either command line param or OS env
// defaults to 1 if no values found
func ReadEnvOrArgs(pname string) int {

    nmbstr := "1"
    if len(os.Args)>1 {
	    nmbstr = os.Args[1]
    }else if len(os.Getenv("nmb"))>0 {
	    nmbstr = os.Getenv("nmb")
    }

    nmb,err := strconv.Atoi(nmbstr)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    return nmb
}
