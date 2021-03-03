var net = require('net');
 
//console.log("process.argv.length=" + process.argv.length +"\n");
if(process.argv.length != 3){
	console.log("Usage: node %s <port>", process.argv[1]);
	process.exit(0);	
}
var port=process.argv[2];
if(port.length >5 ){
	console.log("Invalid port! Try again.\nUsage: node %s <port>", process.argv[1]);
	process.exit(1);	
}
var server = net.createServer();
server.listen(port);
console.log("Simple EchoServer.js developed by Phu Phung, SecAD");
console.log("EchoServer is listenning on port " + port + " ...");

server.on("error", error => {console.log(`Error: ${error}. Exit!`); process.exit(2);});
server.on("connection", handleclient);

function handleclient(client_conn){
	console.log("A new client is connected: %s:%s", 
					client_conn.remoteAddress, 
					client_conn.remotePort);
	client_conn.on("data", data => {
		console.log("Received data:" + data + "Echoed back!");
		client_conn.write(data);
	});
	client_conn.on("close", function(data){
		console.log("Client %s:%s has been disconnected", 
					  client_conn.remoteAddress, client_conn.remotePort);
	});
} 
