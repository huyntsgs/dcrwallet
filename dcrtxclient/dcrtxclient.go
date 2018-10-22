package dcrtxclient

import (
	//"fmt"
	"sync"

	"github.com/decred/dcrwallet/dcrtxclient/service"
	"google.golang.org/grpc"
)

func init() {
	grpc.EnableTracing = false
}

type (
	Config struct {
		Enable  bool
		Address string
		Timeout uint32
	}

	Client struct {
		sync.Mutex
		Cfg  *Config
		conn *grpc.ClientConn
		*service.TransactionService
		IsShutdown bool
	}
)

func SetConfig(cfg *Config) *Client {
	client := &Client{
		Cfg: cfg,
	}
	return client
}

func (c *Client) StartSession() (*Client, error) {
	if c.Cfg.Enable {

		// connect to dcrtxmatcher server if enable
		conn, err := c.Connect()
		if err != nil {
			return nil, err
		}

		// somehow conn object is still nil
		if conn == nil {
			return nil, ErrCannotConnect
		}

		c.conn = conn

		// register services
		c.registerServices()
	}

	return c, nil
}

// connect attempts to connect to our dcrtxmatcher server
func (c *Client) Connect() (*grpc.ClientConn, error) {
	c.Lock()
	defer c.Unlock()

	conn, err := grpc.Dial(c.Cfg.Address, grpc.WithInsecure())
	if err != nil {
		log.Warn("Unable to connect to dcrtxmatcher server")
		return nil, err
	}

	return conn, nil

}

// Disconnect disconnects client from server
// returns error if client is not connected
func (c *Client) Disconnect() error {
	if c.isConnected() {
		c.conn.Close()
		log.Info("dcrTxClient disconnected")
		return nil
	}
	return ErrNotConnected
}

// isConnected checks if client is connected to server
// returns appropriate boolen depending on result
func (c *Client) isConnected() bool {
	if c.conn != nil {
		return true
	}

	return false
}

func (c *Client) registerServices() error {
	if !c.isConnected() {
		return ErrNotConnected
	}

	c.TransactionService = service.NewTransactionService(c.conn)

	return nil
}
