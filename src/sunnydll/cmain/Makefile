#g++ compiler
CC = gcc
CCFLAGS = -I. -Wall -fmessage-length=0 -fPIC
LDFLAGS = -Wl,-rpath,'./'

#Debug or Release, Debug:-g Release:-O3
ifdef GDB
	OPTS = -g -rdynamic
else
	OPTS = -O2 -rdynamic
endif

#Link the .so or .a library
LIBS = -L./ -lsunny

#Target file
TARGET = sunny.exe

#Gernation the target file
all:$(TARGET) 
.PHONY : clean

objects = main.o

$(TARGET):$(objects)
	$(CC) -o $@ $(objects) $(LIBS) $(LDFLAGS)

%.o : %.c
	$(CC) $(CFLAGS) $(OPTS) -c $<

clean:
	-rm -f $(objects) sunny.exe

