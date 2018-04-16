/*
 * palindrometer.c
 *
 *  Created on: Sep 13, 2012
 *      Author: Omar
 */

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define TRUE -1
#define FALSE 0

#define MAX_DIGITS 16

int isPalindrome(char* str) 
{
	int match = TRUE;

	int start, finish;
	for (start = 0, finish = strlen(str) - 1; start < finish; start++, finish--) 
        {
		if (str[start] != str[finish])
                {
			match = FALSE;
			break;
		}
	}
	return match;
}

int nextPalindrome(char* str, char* next)
{
	int start, finish;
	int miles_to_go = 0;
	int first, last;
	int place;
	int carry_check;

	strcpy(next, str);
	for (start = 0, finish = strlen(str) - 1, place = 1; start < finish; start++, finish--, place *= 10)
	{
		/* If digits are different, change finish to match */
		if (next[start] != next[finish])
		{
			first = next[start] - '0';
			last  = next[finish] - '0';
			/* If first is bigger, this is trivial */
			if (first > last)
			{
				miles_to_go += (first - last) * place;
				next[finish] = next[start];
			}
			/* Otherwise, we have to carry */
			else
			{
				miles_to_go += (first - last + 10) * place;
				next[finish] = next[start];
				/* Check for carry */
				for (carry_check = finish - 1; carry_check >= start; carry_check--)
				{
					if (next[carry_check] == '9')
					{
						next[carry_check] = 0;
					}
					else
					{
						next[carry_check]++;
						break;
					}
				}
				/* If carried to start, recheck start by resetting loop parameters */
				if (carry_check == start)
				{
					start--;
					finish++;
					place /= 10;
				}
			}
		}
	}
	return miles_to_go;
}

int main()
{
	char buffer[MAX_DIGITS];
	char next[MAX_DIGITS];
	char* start;

	FILE* input;

	input = (FILE*)fopen("test.txt", "r");
	if (input == NULL )
	{
		printf("Unable to open input file.\n");
	}
	else
	{
		while (fgets(buffer, MAX_DIGITS, input) != NULL)
		{
			start = buffer;
			/* Remove newlines or carriage returns at end */
			while ( (buffer[strlen(buffer)-1] == '\r') ||
                              (buffer[strlen(buffer)-1] == '\n') )
			{
				buffer[strlen(buffer)-1] = '\0';
			}
			/* Remove leading zeros */
			while ( (start[0] == '0') && (start < (buffer + strlen(buffer) - 1)) )
			{
				start++;
			}      
			printf("%s\tSize: %lu\t", start, strlen(start));
			if (isPalindrome(start))
			{
				printf("Yes\t");
			}
			else
			{
				printf("No\t");
			}
			printf("%d\t%s\n", nextPalindrome(start, next), next);
		}
	}
	return 0;
}

