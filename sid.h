/* Code generated by cmd/cgo; DO NOT EDIT. */

/* package udf-sid */


#line 1 "cgo-builtin-export-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#endif

#endif

/* Start of preamble from import "C" comments.  */


#line 3 "main.go"

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

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef _GoString_ GoString;
#endif
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif

extern my_bool DSID_init(UDF_INIT* initid, UDF_ARGS* args, char* message);
extern void DSID_deinit(UDF_INIT* initid);

// DSID ??????base58????????????bigint??????id
extern long long int DSID(UDF_INIT* initid, UDF_ARGS* args, char* isNull, char* error);
extern my_bool ESID_init(UDF_INIT* initid, UDF_ARGS* args, char* message);
extern void ESID_deinit(UDF_INIT* initid);

// ESID ??????bigint???base64?????????id
extern char* ESID(UDF_INIT* initid, UDF_ARGS* args, char* result, GoUint64* length, char* isNull, char* error);

#ifdef __cplusplus
}
#endif
