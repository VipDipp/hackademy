#include <stdlib.h>

char *my_itoa(int nmb)
{
    if (nmb == 0) 
    {
        char *str = (char *) malloc(2);
        *str = '0';
        str++;
        *str = '\0';
        return --str;
    }
    int len = 0;
    int i = 10;
    int pos_nmb = nmb > 0 ? nmb : -nmb;
    int num = nmb > 0 ? 1 : 2;
    while (pos_nmb - i > 0) 
    {
        i *= 10;
        len++;
    }
    char *str = (char *) malloc(len + num);
    if (nmb < 0)
    {
        str[0] = '-';
    }    
    str[len + num] = '\0';
    while (len >= 0)
    {
        str[len + num - 1] = pos_nmb % 10 + '0';
        len--;
        pos_nmb /= 10;
    }
    return str;
}  
