package main

import "fmt"
import "net"
import "bufio"

func main() {
    
    fmt.Printf("Hello, world\n")

    // open AT port

    // enable net

    // try connection
    conn, err := net.Dial("tcp", "baidu.com:80")
    if err != nil {
        fmt.Printf("Connection error!\n")
    } else {
	    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	    status, err := bufio.NewReader(conn).ReadString('\n')
	    if err != nil {
	    	fmt.Printf("Read error!\n")
	    } else {
	    	fmt.Printf("Connection OK \n%s", status)
	    }
	}
}
