/* include libraries */
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <unistd.h>
#include <string.h>
#include <netdb.h>

int main (int argc, char *argv[])
{
   printf("HTTP Client program by Dena Schaeffer\n");
   if (argc!=3)
   {
   	/* code */
   	printf("Usage: %s <servername> <path>\n", argv[0] );
   	exit(1);
   }
   char *servername=argv[1];
   char *path= argv[2];
   if (strlen(servername) > 253 || strlen(path) > 253 )
   {
   	/* code */
   	printf("<servername> or <path> is too long\n");
   	exit(2);
   }
   printf("Got servername= %s, path=%s\n",servername,path);
   int socketfd = socket(AF_INET,SOCK_STREAM,0);
   if (socketfd<0)
   {
   	perror("ERROR opening socket");
   	exit(socketfd);
   }
   printf("A socket is opened!\n");
   struct addrinfo hints, *serveraddr;
   memset(&hints,0,sizeof hints);
   hints.ai_family = AF_INET;
   hints.ai_socktype = SOCK_STREAM;
   int addr_lookup = getaddrinfo(servername, "http", &hints, &serveraddr);
   if (addr_lookup != 0)
   {
   	fprintf(stderr, "getaddrinfo error: %s\n", gai_strerror(addr_lookup));
   	exit(3);
   }
   int connected = connect(socketfd, serveraddr->ai_addr,serveraddr->ai_addrlen);
   if (connected < 0)
   {
   	perror("Cannot connect to the server\n");
   	exit(connected);
   } else
   		printf("Connected to server: %s, path: %s\n", servername,path);
   freeaddrinfo(serveraddr);
   int BUFFERSIZE = 1024;
   char buffer[BUFFERSIZE];
   bzero(buffer,BUFFERSIZE);
   //printf("Enter your message to send: ");
   //fgets(buffer,BUFFERSIZE,stdin);
   sprintf(buffer,"GET %s HTTP/1.1\r\nHost: %s\r\n\r\n",path,servername);
   int byte_sent = send(socketfd,buffer,strlen(buffer), 0);
   bzero(buffer,BUFFERSIZE);
   int byte_received = recv(socketfd,buffer,BUFFERSIZE,0);
   if (byte_received < 0)
   {
   	perror("Error in receiving data.\n");
   	exit(byte_received);
   }
   printf("Received from server:%s\n",buffer);
   //main code
   close(socketfd);//end
}
