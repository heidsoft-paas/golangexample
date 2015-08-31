//http://www.alexonlinux.com/how-debugger-works
//objdump -d sleeper
#include <stdio.h>

int main()
{
        printf( "~~~~~~~~~~~~> Before breakpoint\n" );
        // The breakpoint
        printf( "~~~~~~~~~~~~> After breakpoint\n" );

        return 0;
}
