package main

type State struct {
	requestId                any
	IsCancelled              bool
	IsRedirected             bool
	IsRequestHeaderModified  bool
	IsResponseHeaderModified bool
	IsRequestBodyModified    bool
	IsResponseBodyModified   bool
	DelayedBy                int
	RequestBody              string
}
