#include <stdlib.h>

char *ft_strjoin(const char *s1, const char *s2)
{
    char *str = (char *)malloc(sizeof(char));
    int i = 0;
    while (*s1 != '\0')
    {
        *str = *s1;
        str++;
        s1++;
        i++;
    }
    while (*s2 != '\0')
    {
        *str = *s2;
        str++;
        s2++;
        i++;
    }
    *str = '\0';
    str -= i;
    char *data = (char *)malloc(sizeof(char) * (i+1));
    while (*str != '\0')
    {
        *data = *str;
        str++;
        data++;
    }
    free(str);
    return data - i;
}
