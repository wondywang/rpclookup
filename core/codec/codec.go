package codec

import (
	"io"
	"net"
	"bufio"
	"bytes"
	"encoding/gob"
)

type Header struct {
	T byte
	Seq int64
	Size int32
}

type Request struct {
	Cmd  string
	Body []byte
}

type Response struct {
	Cmd  string
	Body []byte
}

func EncodeHeartBeat(conn net.Conn, seq int64, obj interface{}) error {
	var header Header
	header.T = 'H'
	header.Seq = seq

	err := Encode(conn, header)
	if err != nil {
		return err
	}

	err = Encode(conn, obj)
	if err != nil {
		return err
	}

	buf := bufio.NewWriter(conn)
	buf.Flush()
	return nil
}

func HeaderSize() int {
	var buff bytes.Buffer
	var header Header
	header.T = 'H'
	header.Seq = int64(1542852055751046236)
	header.Size = 9
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(header)
    if err != nil {
		return -1
	}
	return len(buff.Bytes())
}

func Decode(body []byte, obj interface{}) error {
	buff := bytes.NewBuffer(body)
	dec := gob.NewDecoder(buff)
	err := dec.Decode(obj)
	if err != nil {
		return err
	}
	return nil
}

func DecodeHeader(conn io.ReadWriteCloser, header interface{}) error {
	dec := gob.NewDecoder(conn)
	err := dec.Decode(header)
	if err != nil {
		return err
	}
	return nil
}

func DecodeBody(conn io.ReadWriteCloser,body interface{}) error {
	dec := gob.NewDecoder(conn)
	err := dec.Decode(body)
	if err != nil {
		return err
	}
	return nil
}

func Encode(conn io.ReadWriteCloser, obj interface{}) error {
	buf := bufio.NewWriter(conn)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(obj)
    if err != nil {
		return err
	}
	return nil
}