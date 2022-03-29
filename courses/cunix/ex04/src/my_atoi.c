int my_atoi(const char *nptr)
{
    int sign = 1;
    if (nptr[0] == '-') 
    {
        sign = -1;
        nptr++;
    }    
    int num = 0;
    if (!(*nptr > '0' && *nptr <= '9'))
    {    
        return 0;
    }
    while (*nptr >= '0' && *nptr <= '9')
    {
        num = num * 10 + (*nptr - '0');
        nptr++;
    }
    return num * sign;
}    
