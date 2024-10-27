package main

type State struct {
	requestId    int
	IsCancelled  bool
	IsRedirected bool
}
