//echo.h
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <errno.h>

#define SERV_IP "172.16.213.128"
#define SERV_PORT     9877
#define MAXLINE        4096
#define LISTENQ        5

void
err_sys(const char *fmt, ...);

ssize_t
readn(int fd, void *vptr, size_t n);
