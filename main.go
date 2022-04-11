package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <mysql.h>

static int is_arg_string(UDF_ARGS *args,int arg_num) {
	if (args->arg_count > arg_num && args->arg_type[arg_num] == STRING_RESULT) {
		return 1;
	}
	return 0;
}
static int is_arg_int(UDF_ARGS *args,int arg_num) {
	if (args->arg_count > arg_num && args->arg_type[arg_num] == INT_RESULT) {
		return 1;
	}
	return 0;
}

static long long get_int_val(UDF_ARGS *args, int arg_num) {
	long long int_val;
	if (args->arg_count > arg_num) {
		int_val = *((long long*) args->args[arg_num]);
	}
	return int_val;
}
*/
import "C"

import (
	"github.com/bwmarrin/snowflake"
	"unicode/utf8"
	"unsafe"
)

//export DSID_init
func DSID_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.my_bool {
	if args.arg_count != 1 {
		msg := C.CString("DSID(VARCHAR) requires one VARCHAR argument\n")
		C.strcpy(message, msg)
		C.free(unsafe.Pointer(msg))
		return 1
	}
	if C.is_arg_string(args, 0) == 0 {
		msg := C.CString("DSID(VARCHAR) requires one VARCHAR argument\n")
		C.strcpy(message, msg)
		C.free(unsafe.Pointer(msg))
		return 1
	}
	return 0
}

//export DSID_deinit
func DSID_deinit(initid *C.UDF_INIT) {
	C.free(unsafe.Pointer(initid.ptr))
}

// DSID 解码字符串为bigint数字id
//export DSID
func DSID(initid *C.UDF_INIT, args *C.UDF_ARGS, isNull *C.char, error *C.char) C.longlong {
	gArg_count := uint(args.arg_count)
	if gArg_count != 1 {
		return 0
	}

	//字符串id
	baseStr := C.GoString(*args.args)
	if baseStr == "" {
		return 0
	}
	sID, err := snowflake.ParseBase58([]byte(baseStr))
	if err != nil {
		return 0
	}
	return C.longlong(sID.Int64())
}

//export ESID_init
func ESID_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.my_bool {
	if args.arg_count != 1 {
		msg := C.CString("ESID(INTEGER) requires one INTEGER argument\n")
		C.strcpy(message, msg)
		C.free(unsafe.Pointer(msg))
		return 1
	}
	if C.is_arg_int(args, 0) == 0 {
		msg := C.CString("ESID(INTEGER) requires one INTEGER argument\n")
		C.strcpy(message, msg)
		C.free(unsafe.Pointer(msg))
		return 1
	}
	return 0
}

//export ESID_deinit
func ESID_deinit(initid *C.UDF_INIT) {
	C.free(unsafe.Pointer(initid.ptr))
}

// ESID 编码bigint为字符串id
//export ESID
func ESID(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *uint64, isNull *C.char, error *C.char) *C.char {
	gArg_count := uint(args.arg_count)
	if gArg_count != 1 {
		*length = 0
		s := C.CString("")
		initid.ptr = s
		return s
	}
	//bigint的id
	i64 := int64(C.longlong(C.get_int_val(args, 0)))
	sId := snowflake.ParseInt64(i64)
	baseStr := sId.Base58()
	*length = uint64(utf8.RuneCountInString(baseStr))
	s := C.CString(baseStr)
	initid.ptr = s
	return s
}

func main() {}
