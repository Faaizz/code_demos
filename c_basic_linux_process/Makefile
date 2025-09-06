default: main.s

main: main.o
	$(CC) -o main main.o

main.s: main.c
	$(CC) -S main.c

debug: main.o
	$(CC) -g -o main main.o

main.o: main.c
	$(CC) -c main.c

clean:
	rm -f *.s
