//http://www.alexonlinux.com/how-debugger-works
#include <stdio.h>
#include <signal.h>
#include <sys/ptrace.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <sys/user.h>
#include <time.h>

void signal_handler( int sig )
{
	printf( "Process %ld received signal %d\n", (long)getpid(), sig );
}

void do_debugger( pid_t child )
{
	int status = 0;
	long data;
	long orig_data;
	unsigned long addr;

	struct user_regs_struct regs;

	printf( "In debugger process %ld\n", (long)getpid() );

	if (signal( SIGCHLD, signal_handler ) == SIG_ERR) 
	{
		perror( "signal" );
		exit( -1 );
	}

	// Waiting for child process to stop...
	wait( &status );

	// Placing breakpoint...
	//addr = 0x400516;
	addr = 0x4004d2;

	data = ptrace( PTRACE_PEEKTEXT, child, (void *)addr, NULL );
	orig_data = data;
	data = (data & ~0xff) | 0xcc;
	ptrace( PTRACE_POKETEXT, child, (void *)addr, data );

	// Breakpoint is ready. Telling child to continue running...
	ptrace( PTRACE_CONT, child, NULL, NULL );
	child = wait( &status );

	// Restoring original data...
	ptrace( PTRACE_POKETEXT, child, (void *)addr, orig_data );

	// Changing RIP register so that it will point to the right address...
	memset( &regs, 0, sizeof( regs ) );
	ptrace( PTRACE_GETREGS, child, NULL, &regs );
	printf( "RIP before resuming child is %lx\n", regs.rip );
	regs.rip = addr;
	ptrace( PTRACE_SETREGS, child, NULL, &regs );

	// Debuggie is now ready to get resumed... Waiting ten seconds...
	printf( "Time before debugger falling asleep: %ld\n", (long)time( NULL ) );
	sleep( 5 );
	printf( "Time after debugger falling asleep: %ld. Resuming debuggie...\n", (long)time( NULL ) );

	ptrace( PTRACE_CONT, child, NULL, NULL );

	child = wait( &status );
	if (WIFSTOPPED( status ))
		printf( "Debuggie stopped %d\n", WSTOPSIG( status ) );
	if (WIFEXITED( status ))
		printf( "Debuggie exited...\n" );

	printf( "Debugger exiting...\n" );
}

void do_debuggie( void )
{
	char* argv[] = { NULL };
	char* envp[] = { NULL };
	
	printf( "In debuggie process %ld\n", (long)getpid() );

	if (ptrace( PTRACE_TRACEME, 0, NULL, NULL ))
	{
		perror( "ptrace" );
		return;
	}

	execve( "sleeper", argv, envp );
}

int main()
{
	pid_t child;

	// Creating child process. It will execute do_debuggie().
	// Parent process will continue to do_debugger().
	child = fork();
	if (child == 0)
		do_debuggie();
	else if (child > 0)
		do_debugger( child ); 
	else
	{
		perror( "fork" );
		return -1;
	}

	return 0;
}
