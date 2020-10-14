package channel

import (
	"net"

	"github.com/renproject/aw/experiment/codec"
	"github.com/renproject/id"
)

type Channel struct {
	self    id.Signatory
	remote  id.Signatory
	conn    net.Conn
	encoder codec.Encoder
	decoder codec.Decoder
}

func (ch Channel) Write(buf []byte) (int, error) {
	return ch.encoder(ch.conn, buf)
}

func (ch Channel) Read(buf []byte) (int, error) {
	return ch.decoder(ch.conn, buf)
}

func (ch Channel) Self() id.Signatory {
	return ch.self
}

func (ch Channel) Remote() id.Signatory {
	return ch.remote
}

func (ch Channel) Connection() net.Conn {
	return ch.conn
}
