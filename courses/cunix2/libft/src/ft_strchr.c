char *ft_strchr(const char *s, int c)
{  
    c %= 256;

    while (*s != '\0')
    {   
        if (*s == c)
        {
            return (char *) s;
        }
        s++;
    }

    return c == 0 ? (char *) s : 0;
}

