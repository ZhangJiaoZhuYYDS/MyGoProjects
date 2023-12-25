package main

import (
 "sync"
 "time"
)

func main()  {
  i := 1
 mutex := sync.Mutex{}
 go func() {
  for x := 0 ; x < 100 ; x++ {
   mutex.Lock()
   i--
   mutex.Unlock()
   println(i)
  }
 }()
  go func() {
   for x := 0 ; x < 100 ; x++ {
    mutex.Lock()
    i++
    mutex.Unlock()
    println(i)
   }
  }()

 time.Sleep(5*time.Second)
 print(i)
}