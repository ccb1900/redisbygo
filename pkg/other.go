package pkg

const RedisVersion = "5.0.5"
const CRLF = "\r\n"

var CommandGroups = [...]string{
	"generic",
	"string",
	"list",
	"set",
	"sorted_set",
	"hash",
	"pubsub",
	"transactions",
	"connection",
	"server",
	"scripting",
	"hyperloglog",
	"cluster",
	"geo",
	"stream",
}

const ProtoMaxQueryBufLen = 1024 * 1024 * 1024 /* 1GB max query buffer. */
const ProtoIoBufLen = 1024                     /* Generic I/O buffer size */
//const ProtoIoBufLen = 1024 * 16        /* Generic I/O buffer size */
const ProtoReplyChunkBytes = 16 * 1024 /* 16k output buffer */
const ProtoInlineMaxSize = 1024 * 64   /* Max size of inline reads */
const ProtoMBulkBigArg = 1024 * 32
const LongStrSize = 21                      /* Bytes needed for long -> str + '\0' */
const RedisAutoSyncBytes = 1024 * 1024 * 32 /* fdatasync every 32MB */

const (
	ObjEncodingRaw        = iota // raw
	ObjEncodingInt        = iota // int
	ObjEncodingHt         = iota //hash
	ObjEncodingZipMap     = iota // zipmap
	ObjEncodingLinkedList = iota // No longer used: old list encoding
	ObjEncodingZipList    = iota // ziplist
	ObjEncodingIntSet     = iota // intset
	ObjEncodingSkipList   = iota // skiplist
	ObjEncodingEmbStr     = iota // Embedded sds string encoding
	ObjEncodingQuickList  = iota // linked list of ziplists
	ObjEncodingStream     = iota // radix tree of listpacks
)

const (
	ObjString = iota
	ObjList   = iota
	ObjSet    = iota
	ObjZSet   = iota
	ObjHash   = iota
	ObjModule = iota
	ObjStream = iota
)

const LookupNone = 0
const CErr = 0
const COk = 1

const SharedSelectCmds = 10
const ObjSharedIntegers = 10000
const ObjSharedBulkhdrLen = 32
const OBJ_SHARED_REFCOUNT = 2147483647
