package memo

import "fmt"

//!+Func

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res      result
	ready    chan struct{} // closed when res is ready
	canceled <-chan struct{}
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
	done     <-chan struct{}
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

//!-get

//!+monitor

var ErrCanceled = fmt.Errorf("Get: canceled")

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{}), canceled: req.done}
			cache[req.key] = e

			go func(req request) {
				go e.call(f, req.key)

				select {
				case <-e.ready:
					req.response <- e.res
				case <-e.canceled:
					delete(cache, req.key)
					req.response <- result{nil, ErrCanceled}
				}
			}(req)
		} else {
			select {
			case <-e.ready:
				req.response <- e.res
			case <-e.canceled:
				req.response <- result{nil, ErrCanceled}
			}
		}
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	fDone := make(chan struct{})
	go func() {
		e.res.value, e.res.err = f(key)
		close(fDone)
	}()

	select {
	case <-fDone:
		// Broadcast the ready condition.
		close(e.ready)
	case <-e.canceled:
	}
}

//!-monitor
