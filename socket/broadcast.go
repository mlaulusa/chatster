package socket

type Broadcast struct {
	from *Client
	message []byte
}