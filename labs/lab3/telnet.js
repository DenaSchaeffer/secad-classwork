var net = require('net');
 
if(process.argv.length != 4){
	console.log("Usage: node %s <host> <port>", process.argv[1]);
	process.exit(1);	
}

var host=process.argv[2];
var port=process.argv[3];

if(host.length >253 || port.length >5 ){
	console.log("Invalid host or port. Try again!\nUsage: node %s <port>", process.argv[1]);
	process.exit(1);	
}

var client = new net.Socket();
console.log("Simple telnet.js developed by Phu Phung, SecAD");
console.log("Connecting to: %s:%s", host, port);

client.connect(port,host, connected);

function connected(){
	console.log("Connected to: %s:%s", client.remoteAddress, client.remotePort);
}

client.on("data", data => {
	console.log("Received data: " + data);
});
client.on("error", function(err) {
	console.log("Error");
	process.exit(2);
});
client.on("close", function(data) {
	console.log("Connection has been disconnected");
	process.exit(3);
});
const keyboard = require('readline').createInterface({
	input: process.stdin,
	output: process.stdout
});
keyboard.on('line', (input) => {
	console.log(`You typed: ${input}`);
	if(input === ".exit"){
		client.destroy();
		console.log("disconnected!");
		process.exit();
	} else
		client.write(input);
});