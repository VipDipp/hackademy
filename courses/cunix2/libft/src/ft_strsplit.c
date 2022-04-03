#include <stdlib.h>

char **ft_strsplit(const char *s, char c)
{
    char **arr = (char **)malloc(sizeof(char *));
    char *str = NULL;
    char *str1 = NULL;
    int i = 0;
    int j = 0;
    while (*s != '\0')
    {
        j = 0;
        if (*s == c)
        {
            s++;
        }
        else
        {
            str = (char *)malloc(sizeof(char));
            while (*s != c && *s != '\0')
            {
                *str = *s;
                j++;
                s++;
                str++;
            }

            *str = '\0';
            str -= j;
            str1 = (char *)malloc(sizeof(char) * (j + 1));

            while (*str != '\0')
            {
                *str1++ = *str++;
            }

            *str1 = '\0';
                
            *arr++ = str1 - j;
            free(str1);
            free(str);
            i++;
        }    
    }
    *arr = NULL;
    arr -= i;
    char **data = (char **)malloc(sizeof(char *) * (i + 1));
    while (*arr != NULL)
    {
        *data = *arr;
        data++;
        arr++;
    }
    *data = NULL;
    free(arr);
    return data - i;

}
