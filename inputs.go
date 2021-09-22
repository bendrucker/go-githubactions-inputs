package inputs

import "os"

var defaultDecoder = &decoder{
	getenv:    os.Getenv,
	lookupenv: os.LookupEnv,
}

type decoder struct {
	getenv    getenvFunc
	lookupenv lookupEnvFunc
}

type getenvFunc func(string) string
type lookupEnvFunc func(string) (string, bool)
