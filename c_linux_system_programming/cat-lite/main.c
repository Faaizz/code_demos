#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <stdio.h>
#include <errno.h>

int main(int argc, char *argv[]){
  if (argc > 2) {
    errno = E2BIG;
    perror("main");
    return -1;
  }

  char *fname = argv[1];  
  
  int fd;
  fd = open(fname, O_RDONLY);
  if (fd == -1) perror("open");
}
