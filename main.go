package main

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

type Header struct {
	Address int
	isFree bool
	size int
	nextHeader *Header
}

type Allocator struct {
	heapSize int
	freeSize int
	heap []byte
	head *Header
}

func (al *Allocator)init(heapSize int)(err error){
	al.heapSize = heapSize
	al.freeSize = heapSize
	al.heap,err = syscall.Mmap(-1, 0, al.heapSize, syscall.PROT_READ | syscall.PROT_WRITE, syscall.MAP_PRIVATE | syscall.MAP_ANONYMOUS)
	if err !=nil{
		return err
	}
	al.head = &Header{
		isFree:     true,
		size:      heapSize,
		nextHeader: nil,
	}
	return nil
}

func (al *Allocator)mem_alloc(size int)(*interface{}){
	fmt.Printf("Allocated %v bytes\n",size)
	currentHeader := al.head
	adressNum := 0
	for currentHeader !=nil{
		if currentHeader.isFree && currentHeader.size >= size{
			break
		}
		adressNum+=currentHeader.size
		currentHeader = currentHeader.nextHeader
	}
	if currentHeader == nil{
		return nil
	}
	currentHeader.size = size
	currentHeader.Address = adressNum
	currentHeader.isFree = false
	al.freeSize-=size
	if currentHeader.nextHeader == nil{
		currentHeader.nextHeader = &Header{
			isFree:     true,
			size:  al.freeSize,
			nextHeader: nil,
		}
	}
	return (*interface{})(unsafe.Pointer(&al.heap[adressNum]))
}


func (al *Allocator)mem_realloc(address *interface{},size int)(*interface{}){
	currentHeader,err := al.findHeader(address)
	if err !=nil{
		return nil
	}
	fmt.Printf("Realocated %v bytes from %v\n",size,address)
	newAddress := al.mem_alloc(size)
	al.mem_free((*interface{})(unsafe.Pointer(&al.heap[currentHeader.Address])))
	return newAddress

}

func (al *Allocator)mem_free(address *interface{}){
	fmt.Printf("Free block at %v\n",address)
	wantedAddr := (*byte)(unsafe.Pointer(address))
	currentHeader := al.head
	for currentHeader != nil{
		if &al.heap[currentHeader.Address] == wantedAddr{
			break
		}
		currentHeader=currentHeader.nextHeader
	}
	if currentHeader != nil{
		currentHeader.isFree = true
	}

}


func (al *Allocator)findHeader(address *interface{})(header *Header,err error){
	wantedAddr := (*byte)(unsafe.Pointer(address))
	currentHeader := al.head
	for currentHeader != nil{
		if &al.heap[currentHeader.Address] == wantedAddr{
			break
		}
		currentHeader=currentHeader.nextHeader
	}
	if currentHeader == nil{
		return nil,errors.New("There is no block with such address")
	}
	return currentHeader,nil
}

func (al *Allocator)mem_dump(){
	fmt.Printf("\n_____HEAP(%v)_____\n",al.heapSize)
	currentHeader := al.head
	blockNum := 0
	for currentHeader != nil{
		fmt.Printf("BLOCK#%v:Free-%v,Size-%v;\n",blockNum,currentHeader.isFree,currentHeader.size)
		currentHeader = currentHeader.nextHeader
		blockNum++
	}
	fmt.Printf("\n")
}
func main(){
	var all Allocator
	var err error
	err = all.init(64)
	if err!= nil{
		panic(err)
	}
	all.mem_dump()
	adr16:=all.mem_alloc(16)
	all.mem_dump()
	adr8:=all.mem_alloc(8)
	all.mem_dump()
	all.mem_alloc(4)
	all.mem_dump()
	all.mem_free(adr8)
	all.mem_dump()
	all.mem_alloc(2)
	all.mem_dump()
	all.mem_realloc(adr16,32)
	all.mem_dump()
}
