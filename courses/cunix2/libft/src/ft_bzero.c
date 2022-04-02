#include <stddef.h>

void ft_bzero(void *s, size_t n)
{
    char *str = (char *) s;
    while (n-- != 0) 
    {
        *str = '\0';
        str++;
    }
}
