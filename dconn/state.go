package dconn

type stateFn func(*conn) stateFn

type stateSig uint
