package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const (
	BINARY_TYPE uint8 = iota + 1
	STRING_TYPE

	MAX_PAYLOAD_SIZE uint32 = 10 << 20 // 10MB
)

var ErrMaxPayloadSize = errors.New("maximum payload size exceeded")

type Payload interface {
	fmt.Stringer
	io.ReaderFrom
	io.WriterTo
	Bytes() []byte
}

type Binary []byte

func (m Binary) Bytes() []byte {
	return m
}

func (m Binary) String() string {
	return string(m)
}

func (m Binary) WriteTo(w io.Writer) (int64, error) {
	// Type
	err := binary.Write(w, binary.BigEndian, BINARY_TYPE)
	if err != nil {
		return 0, err
	}
	// 기록한 바이트수
	var n int64 = 1

	err = binary.Write(w, binary.BigEndian, uint32(len(m)))
	if err != nil {
		return n, err
	}
	n += 4

	// 페이로드 기록
	o, err := w.Write(m)

	return n + int64(o), err
}

func (m *Binary) ReadFrom(r io.Reader) (int64, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		return 0, err
	}

	// 읽은 바이트수
	var n int64 = 1
	if typ != BINARY_TYPE {
		return n, errors.New("invalid Binary")
	}

	var size uint32
	// 4바이트
	err = binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		return n, err
	}
	n += 4

	if size > MAX_PAYLOAD_SIZE {
		return n, ErrMaxPayloadSize
	}

	*m = make([]byte, size)
	o, err := r.Read(*m)
	return n + int64(o), err
}

type String string

func (m String) Bytes() []byte {
	return []byte(m)
}

func (m String) String() string {
	return string(m)
}

func (m String) WriteTo(w io.Writer) (int64, error) {
	// Type
	err := binary.Write(w, binary.BigEndian, STRING_TYPE)
	if err != nil {
		return 0, err
	}
	// 기록한 바이트수
	var n int64 = 1

	// 길이
	err = binary.Write(w, binary.BigEndian, uint32(len(m)))
	if err != nil {
		return n, err
	}
	n += 4

	// 페이로드 기록
	o, err := w.Write([]byte(m))

	return n + int64(o), err
}

func (m *String) ReadFrom(r io.Reader) (int64, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		return 0, err
	}
	// 읽은 바이트수
	var n int64 = 1
	if typ != STRING_TYPE {
		return n, errors.New("invalid String")
	}

	var size uint32
	// 4바이트
	err = binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		return n, err
	}
	n += 4

	if size > MAX_PAYLOAD_SIZE {
		return n, ErrMaxPayloadSize
	}

	buf := make([]byte, size)
	o, err := r.Read(buf)
	if err != nil {
		return n, err
	}
	*m = String(buf)

	return n + int64(o), nil
}

func decode(r io.Reader) (Payload, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		return nil, err
	}
	var payload Payload
	switch typ {
	case BINARY_TYPE:
		payload = new(Binary)
	case STRING_TYPE:
		payload = new(String)
	default:
		return nil, errors.New("unknown type")
	}

	_, err = payload.ReadFrom(io.MultiReader(bytes.NewReader([]byte{typ}), r))
	if err != nil {
		return nil, err
	}
	return payload, nil
}
