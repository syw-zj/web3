package jsonrpc

import (
	"github.com/laizy/web3/jsonrpc/transport"
)

// Client is the jsonrpc client
type Client struct {
	transport transport.Transport
	endpoints endpoints

	GasLimitFactor func(gasLimit uint64) uint64
}

func DefaultGasFactor(i uint64) uint64 {
	return i*130/100 + 500000
}

type endpoints struct {
	w *Web3
	e *Eth
	n *Net
}

func NewClientWithTransport(trans transport.Transport) *Client {
	c := &Client{GasLimitFactor: DefaultGasFactor}
	c.endpoints.w = &Web3{c}
	c.endpoints.e = &Eth{c}
	c.endpoints.n = &Net{c}

	c.transport = trans
	return c
}

// NewClient creates a new client
func NewClient(addr string) (*Client, error) {
	t, err := transport.NewTransport(addr)
	if err != nil {
		return nil, err
	}

	return NewClientWithTransport(t), nil
}

// Close closes the tranport
func (c *Client) Close() error {
	return c.transport.Close()
}

// Call makes a jsonrpc call
func (c *Client) Call(method string, out interface{}, params ...interface{}) error {
	return c.transport.Call(method, out, params...)
}
