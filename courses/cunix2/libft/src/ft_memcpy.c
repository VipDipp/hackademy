#include <stddef.h>

void *ft_memcpy(void *restrict dest, const void *restrict src, size_t n)
{
    unsigned char *to = (unsigned char *) dest;
    unsigned char *from = (unsigned char *) src;

    while (n-- != 0)
    {
        *to = *from;
        to++;
        from++;
    }
    return dest;
}    
