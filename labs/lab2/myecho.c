/* include libraries */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main (int argc, char *argv[])
{
	char buffer[126];
	if(strlen(argv[1]) > 126){
		printf("Input is too long");
		exit(1);
	}
	strcpy(buffer,argv[1]);
    printf("%s\n",buffer);
}