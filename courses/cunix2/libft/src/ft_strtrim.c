#include <stdlib.h>

char *ft_strtrim(char const *s)
{
    int i = 0;
    int check = 0;
    char *str = (char *)malloc(sizeof(char));
    while (*s != '\0')
    {
        if (*s == ' ' || *s == '\n' || *s == '\t')
        {
            if (check == 1)
            {
                *str = *s;
                str++;
                i++;
            }
        }
        else
        {
            check = 1;
            *str = *s;
            str++;
            i++;
        }
        s++;
    }
    *str = '\0';
    str--;
    while (*str == ' ' || *str == '\n' || *str == '\t')
    {
        *str = '\0';
        str--;
        i--;
    }
    str++;
    char *data = (char *)malloc(sizeof(char) * (i + 1));
    str -= i;
    while (*str != '\0')
    {
        *data = *str;
        data++;
        str++;
    }
    free(str);
    *data = '\0';
    return data - i;
}
