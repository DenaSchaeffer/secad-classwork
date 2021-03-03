import java.net.*;
import java.io.*;
import java.util.ArrayList;

public class EchoServer {
    static ThreadList threadlist = new ThreadList();
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
                // new EchoServerThread(clientSocket).start(); 
                new EchoServerThread(threadlist, clientSocket).start();    
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
    private ThreadList threadlist = null;
        public EchoServerThread(Socket socket){
        clientSocket = socket;
    }
    public EchoServerThread(ThreadList threadlist, Socket socket){
        clientSocket = socket;
        this.threadlist = threadlist;
    }
    public void run(){
        System.out.println("A new thread for client is running...");
        if(threadlist != null) {
            threadlist.addThread(this);
            System.out.println("Inside thread: total connected clients = " + threadlist.getNumberofThreads());
        }
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

class ThreadList{
    private ArrayList<EchoServerThread> threadlist = new ArrayList<EchoServerThread>();
    public ThreadList(){
    }
    public synchronized int getNumberofThreads(){
        //return number of current threads
        return threadlist.size();
    }
    public synchronized void addThread(EchoServerThread newthread){
        //add thread to the threadlist
        threadlist.add(newthread);
    }
    public synchronized void removeThread(EchoServerThread thread){
        //remove thread from threadlist
        threadlist.remove(thread);
    }
    public synchronized void sendToAll(String message) {
        //ask each thread in the threadlist to send the given message to its client
    }
}
