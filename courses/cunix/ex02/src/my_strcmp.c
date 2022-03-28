int my_strcmp(char *str1, char *str2) 
{
    while (*str1 && (*str1 == *str2))
    {
        str1++;
        str2++;
    }

    if ((*str1 - *str2) > 0) 
    {
        return 1;
    }
    else if ((*str1 - *str2) < 0)
    {
        return -1;
    }
    else
    {
        return 0;
    }
}
