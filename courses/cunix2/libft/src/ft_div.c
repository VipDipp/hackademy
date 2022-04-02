#include <stdlib.h>

div_t ft_div(int num, int denom)
{
    div_t res =
    {
        quot: num / denom,
        rem: num % denom
    };
    return res;
}
