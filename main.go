package main

import (
	"awesomeProject1"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var ch chan int
	//定时任务
	ticker := time.NewTicker(time.Second * 30)
	go func() {
		for range ticker.C {
			if Status() {
				awesomeProject1.Alert()
				fmt.Println("error")
			} else {
				fmt.Println("ok")
			}
			// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
		ch <- 1
	}()
	<-ch
}

func Status() bool {
	result := Get("https://www.dxpool.io/index.php?id_product=52&rewrite=hs5-miner&controller=product")
	res := strings.Contains(result, "type=\"submit\"\n                          disabled")
	fmt.Println(result[20047: 20052])
	price, _ := strconv.Atoi(result[20047: 20052])
	if price < 29000 && !res {
		return true
	} else {
		return false
	}
}
