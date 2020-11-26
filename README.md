# os_labs
In this realisation of Memory Allocator we got 2 structures and 6 functions.

```
type Header struct { 
	Address    uint 
	isFree     bool
	size       uint
	nextHeader *Header
}
```
Struct Header consists of Address(Adress of head of the block in memory heap),isFree(boolean showing if block is free),size(size of memory block),nextHeader(pointer to next header) so whole struct takes **32 bytes.**

```
type Allocator struct {
	heapSize uint
	freeSize uint
	heap     []byte
	head     *Header
}
```
Struct Allocator consists of heapSize(number of bytes in whole memore heap we work with),freeSize(Number of free bytes in memory heap),heap([]bytes we work with),head(pointer on start Header).In this lab we work with heap with size of **1024 bytes**

Lets see console output with the result of work functions mem_alloc,mem_free,mem_realloc,mem_dump.

```

_____HEAP(1024)_____
BLOCK#0:Free-true,Size-1024;

Allocated 216+32 bytes

_____HEAP(1024)_____
BLOCK#0:Free-false,Size-248;
BLOCK#1:Free-true,Size-776;

Allocated 124+32 bytes

_____HEAP(1024)_____
BLOCK#0:Free-false,Size-248;
BLOCK#1:Free-false,Size-156;
BLOCK#2:Free-true,Size-620;

Allocated 64+32 bytes

_____HEAP(1024)_____
BLOCK#0:Free-false,Size-248;
BLOCK#1:Free-false,Size-156;
BLOCK#2:Free-false,Size-96;
BLOCK#3:Free-true,Size-524;

Free block at 0x7f4caf1240f8

_____HEAP(1024)_____
BLOCK#0:Free-false,Size-248;
BLOCK#1:Free-true,Size-156;
BLOCK#2:Free-false,Size-96;
BLOCK#3:Free-true,Size-524;

Allocated 32+32 bytes

_____HEAP(1024)_____
BLOCK#0:Free-false,Size-248;
BLOCK#1:Free-false,Size-64;
BLOCK#2:Free-false,Size-96;
BLOCK#3:Free-true,Size-524;

Realocated 108 bytes from 0x7f4caf124000
Allocated 108+32 bytes
Free block at 0x7f4caf124000

_____HEAP(1024)_____
BLOCK#0:Free-true,Size-248;
BLOCK#1:Free-false,Size-64;
BLOCK#2:Free-false,Size-96;
BLOCK#3:Free-false,Size-140;
BLOCK#4:Free-true,Size-320;

```
