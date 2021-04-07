var net = require('net');

if(process.argv.length != 4) {
	console.log("Usage: node %s <host> <port>", process.argv[1]);
	process.exit(0);
}
var host=process.argv[2];
var port=process.argv[3];
if(host.length >253 || port.length >5) {
	console.log("Invalid host or port. Try again!\nUsage: node %s <port>", process.argv[1]);
	process.exit(1);
}
var client = new net.Socket();
console.log("Connecting to: %s:%s", host, port);
client. connect(port, host, connected);

var authenticated = false;

const rl = require('readline');

function connected() {
	console.log("Connected to: %s:%s", client.remoteAddress, client.remotePort);
	console.log("You need to login before sending and receiving messages\n");
	loginsync();
}

var readlineSync = require("readline-sync");


var username;
var password;
function loginsync() {
	username = readlineSync.question('Username: ');
	if (!inputValidated(username)) {
		console.log("Username must have at least five characters. Please try again.");
		loginsync();
		return;
	}

	//cover password text
	password = readlineSync.question('Password: ', {
		hideEchoBack: true // hide with *
	});

	var login = '{"Username":"' + username + '","Password":"'+password+'"}';

	client.write(login);
}

client.on("data", data => {
	console.log("Received data: " + data);
	if (!authenticated){
		if(username && data.toString().includes(username)) {
			console.log("You have logged in successfully witth username " + username);
			authenticated = true;
			chat();
		}else{
			console.log("Authenticated failed. Please try again.");
			loginsync();
		}
	}
});

client.on("error", function(err){
	console.log("Error");
	process.exit(2);
});

client.on("close", function(data){
	console.log("Connection has been disconnected");
	process.exit(3);
});

function chat(){
	var keyboard = rl.createInterface({
		input: process.stdin,
		output: process.stdout
	});

	console.log("Welcome to the Chat System. Type anything to send to public chat.\n");
	console.log("Type .help to view all options");

	keyboard.on('line', (input) => {
		if(input === ".exit") {
			client.write('{"Command":"exit"}');
			setTimeout(() =>{
				client.destroy();
				console.log("Disconnected!");
				process.exit();}, 1);
		}else if(input === ".userlist"){
			client.write('{"Command":"userlist"}');
		}else if(input === ".help"){
			console.log("Welcome to the Chat System. Below are all options available to authenticated users.");
			console.log("Type anything to send to public chat.\n");
			console.log("Type '[To:Receiver] Message' to send to a specific user.");
			console.log("Type .userlist to request latest online users.\nType .exit to logout and close the connection");
		}else if(input.includes("[To:")){
			endname = input.search("]");
			receiver = input.substring(4,endname);
			if ((endname<0) || (receiver == "" || undefined)){
				console.log("Unknown receiver. Try again!\n");
				return;
			}else
				client.write('{"ChatType":"private","Receiver":"'+receiver+
					'","Message":"'+input.substring(endname+1,input.length)+'"}');
		}else
			client.write('{"ChatType":"public","Message":"'+input+'"}');
	});
}

function inputValidated(input){
	//checks to see if the username is longer than five characters
	if (input && input.length>5 && input.length<1000){ //prevents buffer overflow for input
		return true;
	}
}