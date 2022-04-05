#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

char *_numtostr(unsigned int num, char *str)
{
    if (num != 0)
    {
        str = (char *) _numtostr(num / 10, str);
        *str++ = '0' + (num % 10);
    }
    return str;
}

char *my_itoa(int num)
{
    if (num == -2147483648)
    {
        char *str = "-2147483648";
        return str;
    }
    if (num == 0)
    {
        char *str = (char *) malloc(sizeof(char) * 2);
        *str = '0';
        *(str + 1) = '\0';
        return str;
    }

    unsigned int is_negative = (num < 0);
    num *= 1 - 2 * is_negative;

    unsigned int num_len = 1 + is_negative;
    int _num = num;

    while (_num > 9)
    {
        _num /= 10;
        num_len++;
    }

    char *str = (char *) malloc(sizeof(char) * (num_len + 1));
    if (is_negative)
    {
        *str++ = '-';
    }

    str = (char *) _numtostr(num, str);
    *str = '\0';
    return str - num_len;
}

void buff_print(int len, char space, int minus, int plus, int space_count)
{
    char *s = &space;
    char *dec = "-";
    char *inc = "+";
    char *spa = " ";
    minus == 1 && len != 0 ? len-- : 0;
    minus == 1 && space_count != 0 ? space_count-- : 0;
    plus == 1 && len != 0 && minus != 1 ? len-- : 0;

    if (space == '0' && minus == 1) write(1, dec, 1);
    if (space == '0' && minus != 1 && plus == 1) write(1, inc, 1); 

    while (space_count != 0)
    {
        write(1, spa, 1);
        space_count--;
        len > 0 ? len-- : 0;
    }

    while (len != 0)
    {
        write(1, s, 1);
        len--;
    }

    if (space == ' ' && minus == 1) write(1, dec, 1);
    if (space == ' ' && minus != 1 && plus == 1) write(1, inc, 1);
}

void print(char *str)
{
    while (*str != '\0')
    {
        write(1, str, 1);
        str++;
    }
    return;
}

void ft_free(char *str)
{
    if (*str == '\0') free(str);
}

int length(char *str, int str_len)
{
    int len = 0;
    while (*str != '\0')
    {
        len++;
        str++;
    }
    str_len -= len;
    str_len = str_len < 0 ? 0 : str_len;
    return str_len;
}

int ft_printf(const char *format, ...)
{
    va_list ap;
    va_start(ap, format);
    const char *p;

    int count, len, plus, space_count;
    char space;

    for (p = format; *p; p++)
    {
        if (*p != '%')
        {
            write(1, p, 1);
            continue;
        }
        p++;

        len = 0;
        count = 1;
        space_count = 0;

        while (*p == ' ')
        {
            space_count++;
            p++;
        }

        plus = *p == '+' ? 1 : 0;
        *p == '+' ? p++ : p;
        space = *p == '0' ? '0' : ' ';
        *p == '0' ? p++ : p;

        while (*p >= '0' && *p <= '9')
        {
            len *= count;
            count *= 10;
            len += (*p - '0');
            p++;
        }
        len = len < 0 ? 0 : len;
        switch (*p)
        { 
            case 'c':
            {
                char str = va_arg(ap, int);
                char *s = &str;
                len = len < 1 ? 1 : len;
                buff_print(len - 1, space, 0, 0, 0);
                write(1, s, 1);
                break;
            }
            case 's':
            {
                char *str = va_arg(ap, char *);
                if (str == NULL) str = "(null)";
                buff_print(length(str, len), space, 0, 0, 0);
                print(str);
                break;
            }
            case 'i': 
            case 'd':
            {
                int num = va_arg(ap, int);
                char *str = my_itoa(num);
                int minus = num < 0 ? 1 : 0;
                num < 0 ? str++ : 0;
                buff_print(length(str, len), space, minus, plus, space_count);
                print(str);
                num < 0 ? free(--str) : free(str);

                break;
            }
            default:
            {
                write(1, p, 1);
                break;
            }
        }
    }
    return 0;
}
