#include <stdio.h>
#include <string.h>
#include <ctype.h>

void reverse(char *str)
{
    char *start = str;
    char *end = start + strlen(str) - 1; 
    char temp;

    while (end > start)
    {
        temp = *start;
        *start = *end;
        *end = temp;

        ++start;
        --end;
    }
}

int main()
{
    FILE* fp;
    char  line[26];
    fp = fopen("/usr/share/dict/words", "r");
    while (fgets(line, sizeof(line), fp) != NULL)
    {
        int linelen = strlen(line) + 1;
        char word[linelen];
        strcpy(word, line);
        line[0] = tolower(line[0]);
        char backwards[linelen];
        strcpy(backwards, line);
        reverse(backwards);
        if ( strcmp(line,backwards) == 0 )
            printf("%s is a palindrone\n", word);
    }
}
