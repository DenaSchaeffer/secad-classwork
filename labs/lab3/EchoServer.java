import java.net.*;
import java.io.*;

public class EchoServer {
    public static void main(String[] args) throws IOException {
        
        if (args.length != 1) {
            System.err.println("Usage: java EchoServer <port number>");
            System.exit(1);
        }
        
        int portNumber = Integer.parseInt(args[0]);
        
        try {
            ServerSocket serverSocket =
                new ServerSocket(Integer.parseInt(args[0]));
            System.out.println("EchoServer is running at port " + Integer.parseInt(args[0]));
            while (true) {
                Socket clientSocket = serverSocket.accept(); 
                System.out.println("A client is connected "); 
                System.out.println("Hello from Dena Schaeffer");
                //old code
                new EchoServerThread(clientSocket).start();     
            }
        } catch (IOException e) {
            System.out.println("Exception caught when trying to listen on port "
                + portNumber + " or listening for a connection");
            System.out.println(e.getMessage());
        }
    }
}

class EchoServerThread extends Thread {
    private Socket clientSocket = null;
    public EchoServerThread(Socket socket){
        clientSocket = socket;
    }
    public void run(){
        System.out.println("A new thread for client is running...");
        try{
        PrintWriter out =
            new PrintWriter(clientSocket.getOutputStream(), true);
        BufferedReader in = new BufferedReader(
            new InputStreamReader(clientSocket.getInputStream()));
        String inputLine;
        while ((inputLine = in.readLine()) != null) {
            System.out.println("received from client: " + inputLine);
            System.out.println("Echo back");
            out.println(inputLine);
        }
    } catch(IOException ioe){
        System.out.println("Exception in thread:"
                                + ioe.getMessage());
    }
    }
 }

