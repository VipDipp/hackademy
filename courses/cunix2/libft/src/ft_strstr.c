char *ft_strstr(const char *str, const char *substr)
{
    if (*str == '\0')
    {
        return (char *) substr;
    }
    if (*substr == '\0')
    {
        return (char *) str;
    }
    char *str_pos;
    char *substr_pos;
    while (*str != '\0')
    {
        str_pos = (char *) str;
        substr_pos = (char *) substr;
        if (*str == *substr)
        {
            while (*str_pos != '\0')
            {    
                if (*str_pos != *substr_pos)
                {
                    break;
                }
                str_pos++;
                substr_pos++;
                if (*substr_pos == '\0')
                {
                    return (char *) str;
                }
            }
        }
        str++;
    }
    return 0;
}
