package filters

import (
	"sync"

	"github.com/cosiner/zerver"
)

// RequestId is a simple filter prevent application/user from overlap request
// the request id is generated by client itself or other server components.
type RequestId struct {
	requests     map[string]map[string]bool
	HeaderName   string
	RejectOnNoId bool
	Error        string
	ErrorOverlap string
	lock         sync.RWMutex
}

func (ri *RequestId) Init(zerver.Enviroment) error {
	if ri.HeaderName == "" {
		ri.HeaderName = "Request-Id"
	}
	if ri.Error == "" {
		ri.Error = "header value Request-Id can't be empty"
	}
	if ri.ErrorOverlap == "" {
		ri.ErrorOverlap = "request already accepted before, please wait"
	}
	ri.requests = make(map[string]map[string]bool)
	return nil
}

func (ri *RequestId) Filter(req zerver.Request, resp zerver.Response, chain zerver.FilterChain) {
	if req.Method() == "GET" {
		chain(req, resp)
		return
	}
	reqId := req.Header(ri.HeaderName)
	if reqId == "" {
		if ri.RejectOnNoId {
			resp.ReportBadRequest()
			resp.Send("error", ri.Error)
		} else {
			chain(req, resp)
		}
	} else {
		ip := req.RemoteIP()
		ri.lock.RLock()
		ipReqs := ri.requests[ip]
		ri.lock.Unlock()
		if ipReqs == nil {
			ipReqs = make(map[string]bool)
			ipReqs[reqId] = true
			ri.lock.Lock()
			ri.requests[ip] = ipReqs
			ri.lock.Unlock()
		} else {
			if _, has := ipReqs[reqId]; has {
				resp.ReportForbidden()
				resp.Send("error", ri.ErrorOverlap)
				return
			}
			ipReqs[reqId] = true
		}
		chain(req, resp)
		delete(ipReqs, reqId)
	}
}

func (ri *RequestId) Destroy() {
	ri.requests = nil
}
