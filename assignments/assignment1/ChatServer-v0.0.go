/* Simple EchoServer in GoLang by Phu Phung, customized by Dena Schaeffer for SecAD*/
package main

import (
	"fmt"
	"net"
	"os"
	"encoding/json"
)

const BUFFERSIZE int = 1024
const TESTUSER1= "username" // {"username":"TESTUSER1","password":"TESTUSERPASS1"}
const TESTUSERPASS1 = "password" 

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
	// ncusername := make(chan string)
	go func () {
		for {
			client_conn, _ := server.Accept()
			newclient <- client_conn
		}
	}()

	for{
        select{
            case client_conn := <- newclient:
                allClients_conns[client_conn]=client_conn.RemoteAddr().String()
                authenticated, _ := login(client_conn)
			
				if (authenticated){
	                go client_goroutine(client_conn)
				}
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
			clientdata := buffer[0:byte_received]
			fmt.Printf("Received data: %s, len=%d\n", clientdata, len(clientdata))
			go sendtoAll([]byte("Message from " + allClients_conns[client_conn] + ":" + string(buffer[0:byte_received])))
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

func login(client_conn net.Conn) (bool, string) {
	var buffer [BUFFERSIZE]byte
	for {
		byte_received, read_err := client_conn.Read(buffer[0:])
			if read_err != nil {
				fmt.Println("Error in receiving...")
				lostclient <- client_conn
				return false, "Error"
			}
		clientdata := buffer[0:byte_received] // rn the clientdata shoudl be the json string
		authenticated, _, message  := checkLogin(clientdata)

		loginmessage := fmt.Sprintf("%s\n", message)
		go sendto(client_conn, []byte (loginmessage))

		if authenticated{
			return authenticated, message
		}
	}
}

func checkLogin(data []byte) (bool, string, string){
	//expecting format of ("username":"...","password":"...")
	type Account struct{
		Username string
		Password string
	}
	var account Account
	err := json.Unmarshal(data, &account)
	if err!=nil ||account.Username =="" || account.Password == "" {
		fmt.Printf("JSON parsing error: %s\n", err)
		return false, "", "[BAD LOGIN] Expected: {'username':'...','password':'...'}"
	}
	fmt.Println("DEBUG>Got account=%s\n", account)
	fmt.Println("DEBUG>Got username=%s\n, password=%s\n", account.Username, account.Password)

	if checkAccount(account.Username, account.Password) {
		return true, account.Username, "logged"
	}
	
	return false, "", "Invalid username or password\n"
}

func checkAccount(username string, password string) (bool){
	
	if username == "TESTUSER1" && password == "TESTUSERPASS1" {
		fmt.Println("Login access granted!")
		return true
	}
	return false
}
