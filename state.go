package main

type State struct {
	requestId    any
	IsCancelled  bool
	IsRedirected bool
}
