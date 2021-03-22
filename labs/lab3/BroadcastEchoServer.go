/* Simple EchoServer in GoLang by Phu Phung, customized by Dena Schaeffer for SecAD*/
package main

import (
	"fmt"
	"net"
	"os"
)

const BUFFERSIZE int = 1024
var allClients_conns = make(map[net.Conn]string)
var lostclient = make(chan net.Conn)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <port>\n", os.Args[0])
		os.Exit(0)
	}
	port := os.Args[1]
	if len(port) > 5 {
		fmt.Println("Invalid port value. Try again!")
		os.Exit(1)
	}
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Cannot listen on port '" + port + "'!\n")
		os.Exit(2)
	}
	fmt.Println("EchoServer in GoLang developed by Phu Phung, SecAD, revised by Dena Schaeffer")
	fmt.Printf("EchoServer is listening on port '%s' ...\n", port)
	newclient := make(chan net.Conn)

	go func(){
		for {
			client_conn, _ := server.Accept()
			newclient <- client_conn
		}
	}()
	for{
		select{
			case client_conn := <- newclient:
				allClients_conns[client_conn]=client_conn.RemoteAddr().String()
				go client_goroutine(client_conn)
			case client_conn := <- lostclient:
				delete (allClients_conns, client_conn)
				byemessage := fmt.Sprintf("A client '%s' disconnected!\n#of connected clients: %d\n",
					 client_conn.RemoteAddr().String(), len(allClients_conns))
				fmt.Println(byemessage)
				go sendtoAll([]byte (byemessage))
		}
	}
}
func client_goroutine(client_conn net.Conn) {
	welcomemessage := fmt.Sprintf("A new client '%s' connected!\n#of connected clients: %d\n",
					 client_conn.RemoteAddr().String(), len(allClients_conns))
	fmt.Println(welcomemessage)
	go sendtoAll([]byte (welcomemessage))
	var buffer [BUFFERSIZE]byte
	go func(){	
		for {
			byte_received, read_err := client_conn.Read(buffer[0:])
			if read_err != nil {
				fmt.Println("Error in receiving...")
				lostclient <- client_conn
				return
			}
			go sendtoAll(buffer[0:byte_received])
		}
	}() //execute go routine
}
func sendtoAll(data []byte){
	for client_conn, _:= range allClients_conns{
		_, write_err := client_conn.Write(data)
		if write_err != nil {
			fmt.Println("Error in sending...")
			continue
		}
	}
		fmt.Printf("Received data: %sSent to all connected clients!\n", data)
}
func sendto(client_conn net.Conn, data []byte){
		_, write_err := client_conn.Write(data)
		if write_err != nil {
			fmt.Println("Error in sending...")
			return
		}
		fmt.Printf("Received data: %sSent to connected client!\n", data)
}