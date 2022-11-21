package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var host string
var port string

func init() {
	flag.Parse()
	host = flag.Arg(0)
	port = flag.Arg(1)

}

func writer(conn net.Conn, wg *sync.WaitGroup, cancel context.CancelFunc, ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)
	select {
	case <-ctx.Done():
		wg.Done()
	default:
		for {
			fmt.Print(">> ")
			if !scanner.Scan() {
				break
			}

			str := scanner.Text()

			if str == "exit" {
				log.Print("Exiting...")
				cancel()
				break
			}

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}
	}
}

func reader(conn net.Conn, wg *sync.WaitGroup, ctx context.Context) {
	sc := bufio.NewScanner(conn)
	log.Print("Starting reading from server")
	select {
	case <-ctx.Done():
		wg.Done()
	default:
		for {
			if !sc.Scan() {
				log.Print("error reading")
				break
			}
			str := sc.Text()
			log.Printf("<< %s", str)
		}
	}
}

func main() {
	fmt.Println(host, port)

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("connection error: %s", err)
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(2)
	go reader(conn, &wg, ctx)
	go writer(conn, &wg, cancel, ctx)

	wg.Wait()
	conn.Close()
}
