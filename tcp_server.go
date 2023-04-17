package main

import (
	"io"
	"log"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup

<<<<<<< HEAD
	server, err := net.Listen("tcp", "172.17.0.3:8500")
=======
	server, err := net.Listen("tcp", "172.0.0.1:")
>>>>>>> 7108ca7031cf8da3e55d465b9b996370359ee84e
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bundet til %s", server.Addr().String())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			log.Println("før server.Accept() kallet")
			conn, err := server.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer c.Close()
				for {
					buf := make([]byte, 1024)
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return // fra for løkke
					}
					switch msg := string(buf[:n]); msg {
  				        case "ping":
						_, err = c.Write([]byte("pong"))
					default:
						_, err = c.Write(buf[:n])
					}
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return // fra for løkke
					}
				}
			}(conn)
		}
	}()
	wg.Wait()
}
