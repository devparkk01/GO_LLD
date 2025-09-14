package main 

type RequestContext struct {
	id int 
	userId string 
	ip string 
	endpoint string 
}

func NewRequestContext(userId , ip , endpoint string) *RequestContext {
	return &RequestContext{
		userId: userId,
		ip: ip,
		endpoint: endpoint,
	}
}
