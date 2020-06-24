package main

import (
	"fmt"
	"net/http"
	"strings"
)

type (
	//Route ...
	Route struct {
		Method string `json:"method"`
		Path   string `json:"path"`
		Name   string `json:"name"`
	}

	// Router is the registry of all registered routes for an `Echo` instance for
	// request matching and URL path parameter parsing.
	Router struct {
		tree   *node
		routes map[string]*Route
		echo   *Echo
	}

	// Echo ...
	Echo struct {
		maxParam *int
	}
	node struct {
		kind          kind
		label         byte
		prefix        string
		parent        *node
		children      children
		ppath         string
		pnames        []string
		methodHandler *methodHandler
	}
	kind          uint8
	children      []*node
	methodHandler struct {
		connect  HandlerFunc
		delete   HandlerFunc
		get      HandlerFunc
		head     HandlerFunc
		options  HandlerFunc
		patch    HandlerFunc
		post     HandlerFunc
		propfind HandlerFunc
		put      HandlerFunc
		trace    HandlerFunc
		report   HandlerFunc
	}
)

// HandlerFunc ...
type HandlerFunc func() error

const (
	skind kind = iota
	pkind
	akind
	leafkind
	bkind
)

// NewRouter returns a new Router instance.
func NewRouter(e *Echo) *Router {
	return &Router{
		tree: &node{
			methodHandler: new(methodHandler),
		},
		routes: map[string]*Route{},
		echo:   e,
	}
}

func hello() error {
	return nil
}

func main() {
	/*
		e := &Echo{
			maxParam: new(int),
		}
		r := NewRouter(e)
		f := hello

		r.Add("get", "/a/b", f)
		r.show()
		r.Add("get", "/a/c", f)
		r.show()
		r.Add("get", "/aa/dd", f)
		r.show()
		r.Add("post", "/aa/c", f)
		//r.Add("get", "/v1/api/check", f)

		r.show()
	*/
	t := Constructor()
	pt := &t
	pt.Insert("abcd")
	pt.Insert("abef")
	pt.show()
	fmt.Println(pt.Search("abef"))
}

var rawSpace = "    "

func (n *node) show(level int) {
	space := strings.Repeat(rawSpace, level)
	fmt.Printf("%s %d level path:%s\n", space, level, n.prefix)
	if len(n.children) != 0 {
		for _, v := range n.children {
			v.show(level + 1)
		}
	}
}

func (r *Router) show() {
	fmt.Printf("root path:%s\n", r.tree.prefix)
	for _, v := range r.tree.children {
		v.show(1)
	}
}

// Add registers a new route for method and path with matching handler.
func (r *Router) Add(method, path string, h HandlerFunc) {
	// Validate path
	if path == "" {
		path = "/"
	}
	if path[0] != '/' {
		path = "/" + path
	}
	pnames := []string{} // Param names
	ppath := path        // Pristine path

	//滑动扫描path
	//如果有:,分开path和参数
	for i, l := 0, len(path); i < l; i++ {
		if path[i] == ':' {
			j := i + 1

			r.insert(method, path[:i], nil, skind, "", nil)
			for ; i < l && path[i] != '/'; i++ {
			} //i移动到参数后，比如url为/api/user/:id/profile则移动到profile前

			//取:目标名
			pnames = append(pnames, path[j:i])
			path = path[:j] + path[i:] //忽略:标的参数组装
			i, l = j, len(path)

			if i == l { //如果没有/profile
				r.insert(method, path[:i], h, pkind, ppath, pnames)
			} else { //如果有/profile
				r.insert(method, path[:i], nil, pkind, "", nil)
			}
		} else if path[i] == '*' {
			r.insert(method, path[:i], nil, skind, "", nil)
			pnames = append(pnames, "*")
			r.insert(method, path[:i+1], h, akind, ppath, pnames)
		}
	}

	r.insert(method, path, h, skind, ppath, pnames)
}

//
func (r *Router) insert(method, path string, h HandlerFunc, t kind, ppath string, pnames []string) {
	fmt.Printf("router insert path:%s,ppath:%s,pnames:%v\n", path, ppath, pnames)
	// Adjust max param
	l := len(pnames)
	if *r.echo.maxParam < l {
		*r.echo.maxParam = l
	}

	cn := r.tree // Current node as root
	if cn == nil {
		panic("echo: invalid method")
	}
	search := path

	for {
		sl := len(search)
		pl := len(cn.prefix)
		l := 0

		// LCP
		max := pl
		if sl < max {
			max = sl
		}
		for ; l < max && search[l] == cn.prefix[l]; l++ {
		}

		if l == 0 {
			// At root node
			cn.label = search[0]
			cn.prefix = search
			if h != nil {
				cn.kind = t
				cn.addHandler(method, h)
				cn.ppath = ppath
				cn.pnames = pnames
			}
		} else if l < pl {
			// Split node
			n := newNode(cn.kind, cn.prefix[l:], cn, cn.children, cn.methodHandler, cn.ppath, cn.pnames)

			// Reset parent node
			cn.kind = skind
			cn.label = cn.prefix[0]
			cn.prefix = cn.prefix[:l]
			cn.children = nil
			cn.methodHandler = new(methodHandler)
			cn.ppath = ""
			cn.pnames = nil

			cn.addChild(n)

			if l == sl {
				// At parent node
				cn.kind = t
				cn.addHandler(method, h)
				cn.ppath = ppath
				cn.pnames = pnames
			} else {
				// Create child node
				n = newNode(t, search[l:], cn, nil, new(methodHandler), ppath, pnames)
				n.addHandler(method, h)
				cn.addChild(n)
			}
		} else if l < sl {
			search = search[l:]
			c := cn.findChildWithLabel(search[0])
			if c != nil {
				// Go deeper
				cn = c
				continue
			}
			// Create child node
			n := newNode(t, search, cn, nil, new(methodHandler), ppath, pnames)
			n.addHandler(method, h)
			cn.addChild(n)
		} else {
			// Node already exists
			if h != nil {
				cn.addHandler(method, h)
				cn.ppath = ppath
				if len(cn.pnames) == 0 { // Issue #729
					cn.pnames = pnames
				}
			}
		}
		return
	}
}

func newNode(t kind, pre string, p *node, c children, mh *methodHandler, ppath string, pnames []string) *node {
	return &node{
		kind:          t,
		label:         pre[0],
		prefix:        pre,
		parent:        p,
		children:      c,
		ppath:         ppath,
		pnames:        pnames,
		methodHandler: mh,
	}
}

func (n *node) addChild(c *node) {
	n.children = append(n.children, c)
}

func (n *node) findChild(l byte, t kind) *node {
	for _, c := range n.children {
		if c.label == l && c.kind == t {
			return c
		}
	}
	return nil
}

func (n *node) findChildWithLabel(l byte) *node {
	for _, c := range n.children {
		if c.label == l {
			return c
		}
	}
	return nil
}

func (n *node) findChildByKind(t kind) *node {
	for _, c := range n.children {
		if c.kind == t {
			return c
		}
	}
	return nil
}

const (
	// REPORT Method can be used to get information about a resource, see rfc 3253
	REPORT = "REPORT"
)

func (n *node) addHandler(method string, h HandlerFunc) {
	switch method {
	case http.MethodConnect:
		n.methodHandler.connect = h
	case http.MethodDelete:
		n.methodHandler.delete = h
	case http.MethodGet:
		n.methodHandler.get = h
	case http.MethodHead:
		n.methodHandler.head = h
	case http.MethodOptions:
		n.methodHandler.options = h
	case http.MethodPatch:
		n.methodHandler.patch = h
	case http.MethodPost:
		n.methodHandler.post = h
	case http.MethodPut:
		n.methodHandler.put = h
	case http.MethodTrace:
		n.methodHandler.trace = h
	case REPORT:
		n.methodHandler.report = h
	}
}

func (n *node) findHandler(method string) HandlerFunc {
	switch method {
	case http.MethodConnect:
		return n.methodHandler.connect
	case http.MethodDelete:
		return n.methodHandler.delete
	case http.MethodGet:
		return n.methodHandler.get
	case http.MethodHead:
		return n.methodHandler.head
	case http.MethodOptions:
		return n.methodHandler.options
	case http.MethodPatch:
		return n.methodHandler.patch
	case http.MethodPost:
		return n.methodHandler.post
	case http.MethodPut:
		return n.methodHandler.put
	case http.MethodTrace:
		return n.methodHandler.trace
	case REPORT:
		return n.methodHandler.report
	default:
		return nil
	}
}

var (
	methods = [...]string{
		http.MethodConnect,
		http.MethodDelete,
		http.MethodGet,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		http.MethodPut,
		http.MethodTrace,
		REPORT,
	}
	NotFoundHandler = func() error {
		return fmt.Errorf("not found")
	}

	MethodNotAllowedHandler = func() error {
		return fmt.Errorf("method not allowed")
	}
)

func (n *node) checkMethodNotAllowed() HandlerFunc {
	for _, m := range methods {
		if h := n.findHandler(m); h != nil {
			return MethodNotAllowedHandler
		}
	}
	return NotFoundHandler
}

// Find lookup a handler registered for method and path. It also parses URL for path
// parameters and load them into context.
//
// For performance:
//
// - Get context from `Echo#AcquireContext()`
// - Reset it `Context#Reset()`
// - Return it `Echo#ReleaseContext()`.
/*
func (r *Router) Find(method, path string, c Context) {
	ctx := c.(*context)
	ctx.path = path
	cn := r.tree // Current node as root

	var (
		search  = path
		child   *node         // Child node
		n       int           // Param counter
		nk      kind          // Next kind
		nn      *node         // Next node
		ns      string        // Next search
		pvalues = ctx.pvalues // Use the internal slice so the interface can keep the illusion of a dynamic slice
	)

	// Search order static > param > any
	for {
		if search == "" {
			break
		}

		pl := 0 // Prefix length
		l := 0  // LCP length

		if cn.label != ':' {
			sl := len(search)
			pl = len(cn.prefix)

			// LCP
			max := pl
			if sl < max {
				max = sl
			}
			for ; l < max && search[l] == cn.prefix[l]; l++ {
			}
		}

		if l == pl {
			// Continue search
			search = search[l:]
		} else {
			cn = nn
			search = ns
			if nk == pkind {
				goto Param
			} else if nk == akind {
				goto Any
			}
			// Not found
			return
		}

		if search == "" {
			break
		}

		// Static node
		if child = cn.findChild(search[0], skind); child != nil {
			// Save next
			if cn.prefix[len(cn.prefix)-1] == '/' { // Issue #623
				nk = pkind
				nn = cn
				ns = search
			}
			cn = child
			continue
		}

		// Param node
	Param:
		if child = cn.findChildByKind(pkind); child != nil {
			// Issue #378
			if len(pvalues) == n {
				continue
			}

			// Save next
			if cn.prefix[len(cn.prefix)-1] == '/' { // Issue #623
				nk = akind
				nn = cn
				ns = search
			}

			cn = child
			i, l := 0, len(search)
			for ; i < l && search[i] != '/'; i++ {
			}
			pvalues[n] = search[:i]
			n++
			search = search[i:]
			continue
		}

		// Any node
	Any:
		if cn = cn.findChildByKind(akind); cn == nil {
			if nn != nil {
				cn = nn
				nn = cn.parent // Next (Issue #954)
				search = ns
				if nk == pkind {
					goto Param
				} else if nk == akind {
					goto Any
				}
			}
			// Not found
			return
		}
		pvalues[len(cn.pnames)-1] = search
		break
	}

	ctx.handler = cn.findHandler(method)
	ctx.path = cn.ppath
	ctx.pnames = cn.pnames

	// NOTE: Slow zone...
	if ctx.handler == nil {
		ctx.handler = cn.checkMethodNotAllowed()

		// Dig further for any, might have an empty value for *, e.g.
		// serving a directory. Issue #207.
		if cn = cn.findChildByKind(akind); cn == nil {
			return
		}
		if h := cn.findHandler(method); h != nil {
			ctx.handler = h
		} else {
			ctx.handler = cn.checkMethodNotAllowed()
		}
		ctx.path = cn.ppath
		ctx.pnames = cn.pnames
		pvalues[len(cn.pnames)-1] = ""
	}

	return
}*/
