package triblab
import (
       "log"
       "net/rpc"
       "trib"
)
type client struct {
     Addr string
}

func (self *client) Get(key string, value *string) error {
     	conn, e := rpc.DialHTTP("tcp", self.Addr)
     	if e != nil {
     		return e
	}
	
     	// perform the call
	e = conn.Call("Storage.Get", key, value)
     	if e != nil {
     		conn.Close()
		return e
	}
	
	log.Println(self, ".Get:", "key:", key, "value:", *value)

	// close the connection
	return conn.Close()
	
}

func (self *client) Clock(atLeast uint64, ret *uint64) error {
	//connect to the server
	conn, e := rpc.DialHTTP("tcp", self.Addr)
	if e != nil {
		return e
	}
	
	//perform the call
	conn.Call("Storage.Clock", atLeast, ret)
	if e!=nil {
		conn.Close()
		return e
	}
	
	log.Println(self, ".Clock:", "atLeast:", atLeast, "ret:", *ret)

	// Close the connection
	return conn.Close()
}

func (self *client) Set(kv *trib.KeyValue, succ *bool) error {
	//connect to the server
	conn, e:= rpc.DialHTTP("tcp", self.Addr)
	if e!=nil {
		return e
	}
	
	// perform the call
	e = conn.Call("Storage.Set", kv, succ)
	if e!= nil{
		conn.Close()
		return e
	}
	
	log.Println(self, ".Set:", "kv:", kv, "succ:", *succ)
	// close the connection
	return conn.Close()
}

func (self *client) Keys(p *trib.Pattern, r *trib.List) error {
	// Connect to the server
	conn, e := rpc.DialHTTP("tcp", self.Addr)
	if e!=nil {
		return e
	}
	r.L = nil

	// perform the call
	e = conn.Call("Storage.Keys", p, r)
	if e!= nil {
		conn.Close()
		return e
	}

	if r.L == nil {
		r.L = []string{}
	}
	
	log.Println(self, ".Keys:", "pattern:", *p, "ret:", *r)
	// close the connection
	return conn.Close()

}

func (self *client) ListKeys(p* trib.Pattern, r *trib.List) error {
	// connect to the server
	conn, e := rpc.DialHTTP("tcp", self.Addr)
	if e!= nil {
		return e
	}
	
	// perform the call
	e = conn.Call("Storage.ListKeys", p, r)
	if e != nil {
		conn.Close()
		return e
	}

	if r.L == nil {
		r.L = []string{}
	}

	log.Println(self, ".ListKeys:", "p:", *p, "r:", *r)
	
	// close the connection
	return conn.Close()
}

func (self *client) ListGet(key string, ret *trib.List) error {
	// connect the server
	conn, e := rpc.DialHTTP("tcp", self.Addr)
	if e != nil {
		return e
	}
	
	ret.L = nil

	// perform the call
	e = conn.Call("Storage.ListGet", key, ret)
	if e != nil {
		conn.Close()
		return e
	}
	// close the connection
	return conn.Close()
}

func (self *client) ListAppend(kv *trib.KeyValue, succ *bool) error {
	//connect the server
	conn, e := rpc.DialHTTP("tcp", self.Addr)
	if e != nil {
		return e
	}
	
	// perform the call
	e = conn.Call("Storage.ListAppend", kv, succ)
	if e!= nil {
		conn.Close()
		return e
	}
	
	log.Println(self, ".ListAppend:", "kv:", kv.Key, kv.Value, "succ:", *succ)

	// close the connection
	return conn.Close()
}

func (self *client) ListRemove(kv *trib.KeyValue, n *int) error {
	//connect to the server 
	conn, e := rpc.DialHTTP("tcp", self.Addr)
	if e != nil {
		return e
	}
	
	// perform the call
	e = conn.Call("Storage.ListRemove", kv, n)
	if e != nil {
		conn.Close()
		return e
	}
	log.Println(self, ".ListRemove:", "kv:",kv, "n:", *n)
	
	// close the connection
	return conn.Close()
}

var _ trib.Storage = new(client)
