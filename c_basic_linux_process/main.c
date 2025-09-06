#include <stdio.h>
#include <stdlib.h>

int A = 356;

int main() {
  int b = 105;
  int *c = malloc(sizeof(int));
  *c = A + b;
  printf("'c' is %d", *c);
  free(c);
  return 0;
}
