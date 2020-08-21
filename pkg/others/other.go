package others

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
