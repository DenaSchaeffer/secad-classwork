/* Simple EchoServer in GoLang by Phu Phung, customized by Dena Schaeffer for SecAD*/
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"encoding/json"
)

const BUFFERSIZE int = 1024
// const TESTUSER1= "username" // {"username":"TESTUSER1","password":"TESTUSERPASS1"}
// const TESTUSERPASS1 = "password" 

var AuthenticatedClient_conns = make(map[net.Conn]string)
var lostclient = make(chan net.Conn) //step 1
var newclient = make(chan net.Conn)
var userslist = make(map[string]bool)


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

	go func () {
		for {
			client_conn, _ := server.Accept()
			//newclient <- client_conn
			go login(client_conn)
		}
	}()
	for{
        select{
            case client_conn := <- newclient:
            	go authenticating(client_conn)
			case client_conn := <- lostclient: //step 3
                delete (AuthenticatedClient_conns, client_conn)
                byemessage := fmt.Sprintf("A client '%s' disconnected!\n#of connected clients: %d\n",
								client_conn.RemoteAddr().String(), len(AuthenticatedClient_conns))
                fmt.Println(byemessage)
				go sendtoAll([]byte (byemessage))
		}
	}
}

// func privateMessage(client_conn net.Conn, userMessage userMessage){
// 	//add stuff here=
// }

func authenticating(client_conn net.Conn) {
    AuthenticatedClient_conns[client_conn]=client_conn.RemoteAddr().String()
	sendto(client_conn, []byte("You are authenticated! Welcome to the chat system!\n"))
	welcomemessage := fmt.Sprintf("A new client '%s' connected!\n#of connected clients: %d\n",
				 client_conn.RemoteAddr().String(), len(AuthenticatedClient_conns))
	fmt.Println(welcomemessage)
	go sendtoAll([]byte (welcomemessage))
	go client_goroutine(client_conn)
}

func client_goroutine(client_conn net.Conn) {
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
			//compare the data
			
			if len(clientdata) >= 5 && strings.Compare(string(clientdata[0:5]), "login") == 0 {
				fmt.Println("DEBUG> strings.Compare: login data")
				sendto(client_conn, []byte("login data\n"))
			} else {
				fmt.Println("Debug>strings.Compare: non-login data\n")
				sendto(client_conn, []byte("non-login data\n"))
			}
			go sendtoAll(buffer[0:byte_received])
			//go sendtoAll([]byte("Message from " + allClient_conns[client_conn] + ":" + string(buffer[0:byte_received])))
			
			}		
	}() //execute go routine
}

func sendtoAll(data []byte){
	for client_conn, _:= range AuthenticatedClient_conns{
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

func login(client_conn net.Conn) {
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
			//compare the data
			checklogin, username, message := checklogin(clientdata)
			// if len(clientdata) >= 5 && strings.Compare(string(clientdata[0:5]), "login") == 0 {
			if checklogin{
				fmt.Println("DEBUG> Valid JSON login format and login information. Username = " + username + ". Message: " + message)
				newclient <- client_conn
			} else {
				fmt.Println("Debug>Invalid JSON login format\n")
				sendto(client_conn, []byte(message))
				login(client_conn) //call the function again to let them try again
			}
		}		
	}() 
}

func checklogin(data []byte) (bool, string, string){
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
	
	if username == "backd1" && password == "test" {
		fmt.Println("Login access granted!")
		return true
	}
	return false
}
