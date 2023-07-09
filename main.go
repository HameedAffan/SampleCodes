package main


import(
	"fmt"
	"time"
)

func main(){
	fmt.Println("This is a sample code")
	go warn()
	time.Sleep(1*time.Second)
}

func warn(){
	fmt.Println("This func is running under Goroutines--")
}