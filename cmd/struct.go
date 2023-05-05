package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	id      string
	maxConn int
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		id:      "default",
		maxConn: 10,
		tls:     false,
	}
}

type Server struct {
	Opts
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func withMaxConn(maxConn int) OptFunc {

	return func(o *Opts) {
		o.maxConn = maxConn
	}

}

func withID(id string) OptFunc {
	return func(o *Opts) {
		o.id = id
	}
}

func NewServer(opts ...OptFunc) *Server {
	o := defaultOpts()

	for _, opt := range opts {
		opt(&o)
	}

	return &Server{
		Opts: o,
	}
}

func mainStruct() {
	s := NewServer(withID("server1"), withMaxConn(100), withTLS)
	fmt.Printf("%+v\n", s)
}
