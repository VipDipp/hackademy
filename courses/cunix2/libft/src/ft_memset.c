#include <stddef.h>

void *ft_memset(void *s, int c, size_t n)
{
    unsigned char *data = (unsigned char *) s;
    while (n-- != 0)
    {
        *data = (unsigned char) c;
        data++;
    }
    return s;
}
