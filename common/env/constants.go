package env

const (
	EnvConfigHome    = "CONFIG_HOME"
	EnvLogHome       = "LOG_HOME"
	EnvListenAddress = "LISTEN_ADDR"
)

const (
	LogLocal = "log.local"
	LogHome  = "log.home"
	LogFile  = "log.file"
	LogLevel = "log.level"
	LogColor = "log.color"
)

const (
	DbType = "db.type"
	DbUri  = "db.uri"
	DbUser = "db.user"
	DbPass = "db.pass"
	DbHost = "db.host"
	DbPort = "db.port"
	DbName = "db.name"
)

const (
	RedisAddress  = "redis.address"
	RedisPassword = "redis.password"
	RedisPrefix   = "redis.prefix"
)

const (
	GrpcAddress = "grpc.address"
)
