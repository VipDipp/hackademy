#include <stddef.h>

int ft_memcmp(const void *s1, const void *s2, size_t n)
{
    unsigned char *data1 = (unsigned char *) s1;
    unsigned char *data2 = (unsigned char *) s2;
    int diff = 0;
    while (n-- != 0)
    {
        if (*data1 == *data2)
        {
        }
        else
        {
            (*data1 - *data2) > 0 ? diff++ : diff--;
        }
        data1++;
        data2++;
    }
    return diff;
}
