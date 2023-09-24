// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "https://www.baidu.com/", "https://www.sina.com.cn/", "https://www.google.com/")
	}

	for _, url := range os.Args[1:] {
		// start a goroutine
		// 创建一个新的 goroutine
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		// 注意：此处的语法
		// receive from channel ch
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	// don't leak resources
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	// 给 channel 发送数据
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
