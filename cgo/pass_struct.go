package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct string_t {
	char* buf;
    int len;
} string_t;

void free2(void* s) {
	printf("[C freeStr] ptr param s: %p\n", (void *)s);
}

void handle1(string_t* s) {
	printf("[C handle1] ptr param s: %p\n", (void *)s);
	char *v = "aaa";
	char *vPtr = strdup(v);
	s->buf = vPtr;
	printf("[C handle1] ptr s.buf: %p\n", (void *)s->buf);
}

int Print(string_t* s) {
	//puts(s->buf);
	printf("string param: %s |||||\n", s->buf);
	char *strVal = "hello cgoooooo";
	char *strPtr = strdup(strVal);
	// s->buf = strPtr;
    s->len = 2;
	// printf("string param: %s |||||\n", s->buf);
	return s->len;
}

void Print2(char* bArray) {
	int size = sizeof(bArray);
	printf("byte array: ");
    for (int i = 0; i < size; i++) {
        printf("%c,", bArray[i]);
    }
    printf("\n");
}

// unsigned char* BArray() {
//     unsigned char b[] = {1, 2, 3};
// 	return b;
// }

*/
import "C"
import (
	"fmt"
	"time"

	"unsafe"
)

// import "C"更像是一个关键字，CGO工具在预处理时会删掉这一行

type str struct {
	buf string
	len int
}

func free(strC *C.struct_string_t) {
	fmt.Printf("[go free] address of strC: %p", unsafe.Pointer(strC))
}

func main() {
	strC := C.struct_string_t{}
	fmt.Printf("[go] address of strC: %p\n", &strC)

	// defer free(&strC)

	// err: cannot use _cgo0 (variable of type *_Ctype_struct_string_t) as type unsafe.Pointer in argument to _Cfunc_free2
	// defer C.free2(&strC)

	// defer C.free2(unsafe.Pointer(&strC))

	defer C.free2(unsafe.Pointer(strC.buf))

	time.Sleep(3 * time.Second)
	C.handle1(&strC)
	fmt.Printf("strC: %+v\n", strC)
	fmt.Printf("strC: {buf:%v, len:%d}\n", C.GoString(strC.buf), int(strC.len))
	// free(&strC)
}

func main1() {
	s := C.CString("123")
	defer C.Print2(s)
	s2 := C.CString("456")
	defer C.Print2(s2)
	s3 := C.CString("789")
	defer C.Print2(s3)
	C.puts(C.CString("Hello, Cgo\n"))
	time.Sleep(10 * time.Second)

	// s := &str{
	// 	// buf: "hello, XXXXXXXXXXXXXXX",
	// 	// len: 10,
	// }
	// cs := (*C.struct_string_t)(unsafe.Pointer(s))
	// // cs := (*C.string_t)(unsafe.Pointer(s))
	// C.Print(cs)
	// fmt.Printf("\n return value: %+v\n\n", s)
	// fmt.Printf("len(s.buf): %d\n", len(s.buf))
	// C.free(unsafe.Pointer(cs.buf))
	// C.free(unsafe.Pointer(&s.buf))
	// fmt.Printf("\n return value: %+v\n\n", s.buf)
	// fmt.Printf("\n str.buf len: %d\n\n", len(s.buf))

	// cString := C.struct_string_t{}
	// C.Print(&cString)
	// // cannot use _cgo0 (variable of type *_Ctype_struct_string_t) as type unsafe.Pointer in argument to _Cfunc_free
	// // C.free(&cString)
	//
	// fmt.Printf("cStr{buf:%v, len:%d}\n", C.GoString(cString.buf), int(cString.len))
	//
	// // error: malloc: *** error for object 0x...: pointer being freed was not allocated
	// // C.free(unsafe.Pointer(&cString))
	//
	// C.free(unsafe.Pointer(cString.buf))
	// fmt.Printf("cString.buf is released\n")
	// error: malloc: *** error for object 0x14000012248: pointer being freed was not allocated
	// C.free(unsafe.Pointer(&cString.len))
	// fmt.Printf("cString.len is released\n")

	// error: cannot use _cgo0 (variable of type *_Ctype_struct_string_t) as type
	// unsafe.Pointer in argument to _Cfunc_free
	// C.free((*C.struct_string_t)(unsafe.Pointer(&cString)))

	// goSlice := []byte{1, 2, 3}
	// cArray := C.CBytes(goSlice)
	// goSlice[0] = 100
	// fmt.Printf("goSlice: %+v\n", goSlice)
	// C.Print2((*C.uchar)(cArray))

	// // cArray2 := C.BArray()
	// C.Print2((*C.uchar)(cArray))
	// slice2 := C.GoBytes(cArray, C.int(len(goSlice)))
	// fmt.Printf("goSlice2: %+v\n", slice2)
	// slice2[0] = 200
	// fmt.Printf("goSlice2: %+v\n", slice2)
	// C.Print2((*C.uchar)(cArray))
}
