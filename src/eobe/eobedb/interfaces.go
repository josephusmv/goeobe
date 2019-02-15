package eobedb

//DBLoggerInf Let User to implement this interface.
// If no concrete implementation provided, will use stdout(fmt.Print) as default.
type DBLoggerInf interface {
	TraceError(format string, a ...interface{}) error
	TraceDev(format string, a ...interface{})
	TraceInfo(format string, a ...interface{})
}

//DBQueryInf Interface for running DB, actually the same instance as DBServerInf, indenpendt for protecting DBServer
type DBQueryInf interface {
	ExecDBAction(QueryData, DBLoggerInf) (QueryResult, error)
}

//DBServerInf Represented the DB server, open for start and stop
type DBServerInf interface {
	//Start Start DB server.
	Init() error
	//SetDBOptions Set DB Server options, for supported options and its types, see eobedb.OptionDB* constants
	//	Recommand to call this method Before calling DBServerInf.Init()
	SetDBOptions(option int, value interface{}) error
	//StopDBServer Set DB Server options, for supported options and its types, see eobedb.OptionDB* constants
	//	Recommand to call this method after DB initated.
	StopDBServer()
}
