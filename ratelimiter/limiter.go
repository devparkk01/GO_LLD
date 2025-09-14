package main 


type Limiter interface {
	Allow(req *RequestContext) bool 
	Name() string 
}

