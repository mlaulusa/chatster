package main

type Broadcast struct {
	from *Client
	message []byte
}