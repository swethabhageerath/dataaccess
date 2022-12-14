package client

type Sqler interface {
	Executor
	Querier
}
