char *my_strcpy(char *dest, const char *src)
{
    char *buffer = dest;
    while (*src) 
    {
        *dest++ = *src ++;
    }
    *dest = '\0';
    return buffer;
}
