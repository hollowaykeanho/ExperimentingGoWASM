// Copyright 2022 "Holloway" Chew, Kean Ho <hollowaykeanho@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package hestiaError

const (
	OK       = 0
	ERROR_OK = ""

	EPERM       = 1
	ERROR_EPERM = "operation not permitted"

	ENOENT       = 2
	ERROR_ENOENT = "no such entity like data, file, or directory"

	ESRCH       = 3
	ERROR_ESRCH = "no such process"

	EINTR       = 4
	ERROR_EINTR = "interrupted system call"

	EIO       = 5
	ERROR_EIO = "I/O error"

	ENXIO       = 6
	ERROR_ENXIO = "no such device or address"

	E2BIG       = 7
	ERROR_E2BIG = "argument list too long"

	ENOEXEC       = 8
	ERROR_ENOEXEC = "exec format error"

	EBADF       = 9
	ERROR_EBADF = "bad file number"

	ECHILD       = 10
	ERROR_ECHILD = "no child processes"

	EAGAIN       = 11
	ERROR_EAGAIN = "try again"

	ENOMEM       = 12
	ERROR_ENOMEM = "out of memory"

	EACCES       = 13
	ERROR_EACCES = "permission denied"

	EFAULT       = 14
	ERROR_EFAULT = "bad address"

	ENOTBLK       = 15
	ERROR_ENOTBLK = "block device required"

	EBUSY       = 16
	ERROR_EBUSY = "device or resource busy"

	EEXIST       = 17
	ERROR_EEXIST = "target exists"

	EXDEV       = 18
	ERROR_EXDEV = "cross-device link"

	ENODEV       = 19
	ERROR_ENODEV = "no such device"

	ENOTDIR       = 20
	ERROR_ENOTDIR = "not a directory"

	EISDIR       = 21
	ERROR_EISDIR = "is a directory"

	EINVAL       = 22
	ERROR_EINVAL = "invalid argument"

	ENFILE       = 23
	ERROR_ENFILE = "file table overflow"

	EMFILE       = 24
	ERROR_EMFILE = "too many open file"

	ENOTTY       = 25
	ERROR_ENOTTY = "not a typewriter"

	ETXTBSY       = 26
	ERROR_ETXTBSY = "text file busy"

	EFBIG       = 27
	ERROR_EFBIG = "file too large"

	ENOSPC       = 28
	ERROR_ENOSPC = "no space left"

	ESPIPE       = 29
	ERROR_ESPIPE = "illegal seek"

	EROFS       = 30
	ERROR_EROFS = "read-only filesystem"

	EMLINK       = 31
	ERROR_EMLINK = "too many links"

	EPIPE       = 32
	ERROR_EPIPE = "broken pipe"

	EDOM       = 33
	ERROR_EDOM = "math argument out of domain of function"

	ERANGE       = 34
	ERROR_ERANGE = "math result not representable"

	EDEADLK       = 35
	ERROR_EDEADLK = "deadlock occured"

	ENAMETOOLONG       = 36
	ERROR_ENAMETOOLONG = "filename too long"

	ENOLOCK       = 37
	ERROR_ENOLOCK = "no such lock"

	ENOSYS       = 38
	ERROR_ENOSYS = "invalid system call number"

	ENOTEMPTY       = 39
	ERROR_ENOTEMPTY = "directory not empty"

	ELOOP       = 40
	ERROR_ELOOP = "too many symbolic link encountered"

	EWOULDBLOCK       = 41
	ERROR_EWOULDBLOCK = "operation would block"

	ENOMSG       = 42
	ERROR_ENOMSG = "no message found"

	EIDRM       = 43
	ERROR_EIDRM = "identifier removed"

	ECHRNG       = 44
	ERROR_ECHRNG = "channel out of range"

	// Skip unrelated Linux-specific error codes

	EBFONT       = 59
	ERROR_EBFONT = "error bad font"

	ENOSTR       = 60
	ERROR_ENOSTR = "operation would block"

	ENODATA       = 61
	ERROR_ENODATA = "no data available"

	ETIME       = 62
	ERROR_ETIME = "time expired"

	ENOSR       = 63
	ERROR_ENOSR = "out of stream resources"

	ENONET       = 64
	ERROR_ENONET = "not on network"

	ENOPKG       = 65
	ERROR_ENOPKG = "package unavailable or not installed"

	EREMOTE       = 66
	ERROR_EREMOTE = "object is remote"

	ENOLINK       = 67
	ERROR_ENOLINK = "link unavailable or was severed"

	EADV       = 68
	ERROR_EADV = "advertise error"

	ESRMNT       = 69
	ERROR_ESRMNT = "surmount error"

	ECOMM       = 70
	ERROR_ECOMM = "error on send communication"

	EPROTO       = 71
	ERROR_EPROTO = "protocol error"

	EMULTIHOP       = 72
	ERROR_EMULTIHOP = "multihop attempted"

	EDOTDOT       = 73
	ERROR_EDOTDOT = "RFS specific error"

	EBADMSG       = 74
	ERROR_EBADMSG = "not a data message"

	EOVERFLOW       = 75
	ERROR_EOVERFLOW = "value too large for defined data type"

	ENOTUNIQ       = 76
	ERROR_ENOTUNIQ = "value is not unique"

	EBADFD       = 77
	ERROR_EBADFD = "file descriptor in bad state"

	EREMCHG       = 78
	ERROR_EREMCHG = "remote address changed"

	ELIBACC       = 79
	ERROR_ELIBACC = "cannot access shared library"

	ELIBBAD       = 80
	ERROR_ELIBBAD = "bad shared library"

	ELIBSCN       = 81
	ERROR_ELIBSCN = "shared library corrupted"

	ELIBMAX       = 82
	ERROR_ELIBMAX = "too many shared libraries"

	ELIBEXEC       = 83
	ERROR_ELIBEXEC = "cannot exec shared library"

	EILSEQ       = 84
	ERROR_EILSEQ = "Illegal byte sequence"

	ERESTART       = 85
	ERROR_ERESTART = "should be restarted"

	ESTRPIPE       = 86
	ERROR_ESTRPIPE = "stream pipe error"

	EUSERS       = 87
	ERROR_EUSERS = "too many users"

	ENOTSOCK       = 88
	ERROR_ENOTSOCK = "socket operations on non-socket"

	EDESTADDRREQ       = 89
	ERROR_EDESTADDRREQ = "destination address required"

	EMSGSIZE       = 90
	ERROR_EMSGSIZE = "message too long"

	EPROTOTYPE       = 91
	ERROR_EPROTOTYPE = "wrong protocol"

	ENOPROTOOPT       = 92
	ERROR_ENOPROTOOPT = "unavailable protocol"

	EPROTONOSUPPORT       = 93
	ERROR_EPROTONOSUPPORT = "unsupported protocol"

	ESOCKTNOSUPPORT       = 94
	ERROR_ESOCKTNOSUPPORT = "unsupported socket"

	EOPNOTSUPP       = 95
	ERROR_EOPNOTSUPP = "unsupported operation"

	EPFNOSUPPORT       = 96
	ERROR_EPFNOSUPPORT = "unsupported protocol suite/family"

	EAFNOSUPPORT       = 97
	ERROR_EAFNOSUPPORT = "address unsupported by protocol suite/family"

	EADDRINUSE       = 98
	ERROR_EADDRINUSE = "address in use"

	EADDRNOTAVAIL       = 99
	ERROR_EADDRNOTAVAIL = "cannot assign requested address"

	ENETDOWN       = 100
	ERROR_ENETDOWN = "network down"

	ENETUNREACH       = 101
	ERROR_ENETUNREACH = "network unreachable"

	ENETRESET       = 102
	ERROR_ENETRESET = "network reset"

	ECONNABORTED       = 103
	ERROR_ECONNABORTED = "network aborted"

	ECONNRESET       = 104
	ERROR_ECONNRESET = "connection reset by peer"

	ENOBUFS       = 105
	ERROR_ENOBUFS = "not buffer space available"

	EISCONN       = 106
	ERROR_EISCONN = "is already connected"

	ENOTCONN       = 107
	ERROR_ENOTCONN = "not connected"

	ESHUTDOWN       = 108
	ERROR_ESHUTDOWN = "already shutdown"

	ETOOMANYREFS       = 109
	ERROR_ETOOMANYREFS = "too many references"

	ETIMEDOUT       = 110
	ERROR_ETIMEDOUT = "timeout"

	ECONNREFUSED       = 111
	ERROR_ECONNREFUSED = "connection refused"

	EHOSTDOWN       = 112
	ERROR_EHOSTDOWN = "host is down"

	EHOSTUNREACH       = 113
	ERROR_EHOSTUNREACH = "host is unreachable"

	EALREADY       = 114
	ERROR_EALREADY = "operations is already in progress"

	EINPROGRESS       = 115
	ERROR_EINPROGRESS = "operations is now in progress"

	ESTALE       = 116
	ERROR_ESTALE = "operations is stalled"

	EUCLEAN       = 117
	ERROR_EUCLEAN = "cleaning is required"

	// skip linux specific error codes

	EREMOTEIO       = 121
	ERROR_EREMOTEIO = "remote I/O error"

	EDQUOT       = 122
	ERROR_EDQUOT = "quota exceeded"

	// skip linux specific error codes

	ECANCELED       = 125
	ERROR_ECANCELED = "opreation cancelled"

	ENOKEY       = 126
	ERROR_ENOKEY = "required key not available"

	EKEYEXPIRED       = 127
	ERROR_EKEYEXPIRED = "required key has expired"

	EKEYREVOKED       = 128
	ERROR_EKEYREVOKED = "required key has been revoked"

	EKEYREJECTED       = 129
	ERROR_EKEYREJECTED = "required key has been rejected"

	EOWNERDEAD       = 130
	ERROR_EOWNERDEAD = "owner died"

	ENOTRECOVERABLE       = 131
	ERROR_ENOTRECOVERABLE = "current state is unrecoverable"

	ERFKILL       = 132
	ERROR_ERFKILL = "operations not possible due to RF-kill"

	EHWPOISON       = 133
	ERROR_EHWPOISON = "hardware error"
)
