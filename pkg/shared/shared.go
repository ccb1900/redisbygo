package shared

import (
	"fmt"
	"github.com/ccb1900/redisbygo/pkg"
	"strconv"
	"sync"
)

type Shared struct {
	CRLF             *pkg.RedisObject
	Ok               *pkg.RedisObject
	EmptyBulk        *pkg.RedisObject
	CZero            *pkg.RedisObject
	Cone             *pkg.RedisObject
	CNeGone          *pkg.RedisObject
	Pong             *pkg.RedisObject
	Space            *pkg.RedisObject
	Colon            *pkg.RedisObject
	NullBulk         *pkg.RedisObject
	NullMultiBulk    *pkg.RedisObject
	Queued           *pkg.RedisObject
	EmptyMultiBulk   *pkg.RedisObject
	WrongTypeErr     *pkg.RedisObject
	NoKeyErr         *pkg.RedisObject
	SyntaxErr        *pkg.RedisObject
	SameObjectErr    *pkg.RedisObject
	OutOfRangeErr    *pkg.RedisObject
	NoScriptErr      *pkg.RedisObject
	LoadingErr       *pkg.RedisObject
	SlowScriptErr    *pkg.RedisObject
	BgSaveErr        *pkg.RedisObject
	MasterDownErr    *pkg.RedisObject
	RoSalveErr       *pkg.RedisObject
	ExecAbortErr     *pkg.RedisObject
	NoAuthErr        *pkg.RedisObject
	NoReplicasErr    *pkg.RedisObject
	BusyKeyErr       *pkg.RedisObject
	OOMErr           *pkg.RedisObject
	Plus             *pkg.RedisObject
	MessageBulk      *pkg.RedisObject
	PMessageBulk     *pkg.RedisObject
	SubscribeBulk    *pkg.RedisObject
	UnSubscribeBulk  *pkg.RedisObject
	PSubscribeBulk   *pkg.RedisObject
	PUnSubscribeBulk *pkg.RedisObject
	Del              *pkg.RedisObject
	Unlink           *pkg.RedisObject
	RPop             *pkg.RedisObject
	LPop             *pkg.RedisObject
	LPush            *pkg.RedisObject
	RPopLPush        *pkg.RedisObject
	ZPopMin          *pkg.RedisObject
	ZPopMax          *pkg.RedisObject
	EmptyScan        *pkg.RedisObject
	Selects          [pkg.SharedSelectCmds]*pkg.RedisObject
	Integers         [pkg.ObjSharedIntegers]*pkg.RedisObject
	MBulkHdr         [pkg.ObjSharedBulkhdrLen]*pkg.RedisObject
	BulkHdr          [pkg.ObjSharedBulkhdrLen]*pkg.RedisObject
	Err              *pkg.RedisObject
	MinString        string
	MaxString        string
}

var once sync.Once
var shared *Shared

func NewShared() *Shared {
	once.Do(func() {
		shared = new(Shared)
		crlf := "\r\n"
		shared.CRLF = pkg.NewStringRedisObject(crlf)
		shared.Ok = pkg.NewStringRedisObject("+OK" + crlf)
		shared.Err = pkg.NewStringRedisObject("-ERR" + crlf)
		shared.EmptyBulk = pkg.NewStringRedisObject("$0" + crlf + crlf)
		shared.CZero = pkg.NewStringRedisObject(":0" + crlf)
		shared.Cone = pkg.NewStringRedisObject(":1" + crlf)
		shared.CNeGone = pkg.NewStringRedisObject(":-1" + crlf)
		shared.NullBulk = pkg.NewStringRedisObject("$-1" + crlf)
		shared.NullMultiBulk = pkg.NewStringRedisObject("*-1" + crlf)
		shared.EmptyMultiBulk = pkg.NewStringRedisObject("*0" + crlf)
		shared.Pong = pkg.NewStringRedisObject("+PONG" + crlf)
		shared.Queued = pkg.NewStringRedisObject("+QUEUED" + crlf)
		shared.EmptyScan = pkg.NewStringRedisObject("*2" + crlf + "$1" + crlf + "0" + crlf + "*0" + crlf)
		shared.WrongTypeErr = pkg.NewStringRedisObject("-WRONGTYPE Operation against a key holding the wrong kind of value" + crlf)
		shared.NoKeyErr = pkg.NewStringRedisObject("-ERR no such key" + crlf)
		shared.SyntaxErr = pkg.NewStringRedisObject("-ERR syntax error" + crlf)
		shared.SameObjectErr = pkg.NewStringRedisObject("-ERR source and destination objects are the same" + crlf)
		shared.OutOfRangeErr = pkg.NewStringRedisObject("-ERR index out of range" + crlf)
		shared.NoScriptErr = pkg.NewStringRedisObject("-NOSCRIPT No matching script. Please use EVAL." + crlf)
		shared.LoadingErr = pkg.NewStringRedisObject("-LOADING Redis is loading the dataset in memory" + crlf)
		shared.SlowScriptErr = pkg.NewStringRedisObject("-BUSY Redis is busy running a script. You can only call SCRIPT KILL or SHUTDOWN NOSAVE." + crlf)
		shared.MasterDownErr = pkg.NewStringRedisObject("-MASTERDOWN Link with MASTER is down and replica-serve-stale-data is set to 'no'." + crlf)
		shared.BgSaveErr = pkg.NewStringRedisObject("-MISCONF Redis is configured to save RDB snapshots, but it is currently not able to persist on disk. Commands that may modify the data set are disabled, because this instance is configured to report errors during writes if RDB snapshotting fails (stop-writes-on-bgsave-error option). Please check the Redis logs for details about the RDB error." + crlf)
		shared.RoSalveErr = pkg.NewStringRedisObject("-READONLY You can't write against a read only replica." + crlf)
		shared.NoAuthErr = pkg.NewStringRedisObject("-NOAUTH Authentication required." + crlf)
		shared.OOMErr = pkg.NewStringRedisObject("-OOM command not allowed when used memory > 'maxmemory'." + crlf)
		shared.ExecAbortErr = pkg.NewStringRedisObject("-EXECABORT Transaction discarded because of previous errors." + crlf)
		shared.NoReplicasErr = pkg.NewStringRedisObject("-NOREPLICAS Not enough good replicas to write." + crlf)
		shared.BusyKeyErr = pkg.NewStringRedisObject("-BUSYKEY Target key name already exists." + crlf)

		shared.Space = pkg.NewStringRedisObject(" ")
		shared.Plus = pkg.NewStringRedisObject("+")
		shared.Colon = pkg.NewStringRedisObject(":")

		shared.MessageBulk = pkg.NewStringRedisObject("$7" + crlf + "message" + crlf)
		shared.PMessageBulk = pkg.NewStringRedisObject("$8" + crlf + "pmessage" + crlf)
		shared.SubscribeBulk = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)
		shared.UnSubscribeBulk = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)
		shared.PSubscribeBulk = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)
		shared.PUnSubscribeBulk = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)
		shared.Del = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)
		shared.Unlink = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)
		shared.RPop = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)
		shared.LPop = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)
		shared.LPush = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)
		shared.RPopLPush = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)
		shared.ZPopMin = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)
		shared.ZPopMax = pkg.NewStringRedisObject("$9" + crlf + "subscribe" + crlf)

		shared.MaxString = "maxstring"
		shared.MinString = "minstring"

		for i := 0; i < pkg.SharedSelectCmds; i++ {
			shared.Selects[i] = pkg.NewStringRedisObject(fmt.Sprintf("*2"+crlf+"$6"+crlf+"SELECT"+crlf+"$%d"+crlf+"%s"+crlf,
				1, "1"))
		}
		for i := 0; i < pkg.ObjSharedBulkhdrLen; i++ {
			shared.BulkHdr[i] = pkg.NewStringRedisObject(fmt.Sprintf("$%d"+crlf, i))
			shared.MBulkHdr[i] = pkg.NewStringRedisObject(fmt.Sprintf("*%d"+crlf, i))
		}
		for i := 0; i < pkg.ObjSharedIntegers; i++ {
			shared.Integers[i] = pkg.NewStringRedisObject(strconv.Itoa(i)).MakeShared()
			shared.Integers[i].Encoding = pkg.ObjEncodingInt
		}

	})
	return shared
}
