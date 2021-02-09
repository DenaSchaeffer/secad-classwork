/* include libraries */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main (int argc, char *argv[])
{
	char buffer[126];
	strcpy(buffer,argv[1]);
    printf("%s\n", buffer);
}