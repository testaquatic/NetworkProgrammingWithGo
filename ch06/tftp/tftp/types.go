package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strings"
)

const (
	// 최대 지원하는 데이터그램 크기
	DATAGRAM_SIZE = 516
	// DATAGRAM_SIZE - 4바이트 헤더
	BLOCK_SIZE = DATAGRAM_SIZE - 4
)

type OpCode uint16

const (
	// 읽기 요청
	OP_RRQ OpCode = iota + 1
	// WRQ는 정의하지 않음
	_
	// 데이터 작업
	OP_DATA
	// 메시지 승인
	OP_ACK
	// 오류
	OP_ERR
)

type ErrCode uint16

const (
	ERR_UNKNOWN ErrCode = iota
	ERR_NOTFOUND
	ERR_ACCESSVIOLATION
	ERR_DISKFULL
	ERR_ILLEGALOP
	ERR_UNKNOWNID
	ERR_FILEEXISTS
	ERR_NOUSER
)

type ReadReq struct {
	Filename string
	Mode     string
}

func (q ReadReq) MarshalBinary() ([]byte, error) {
	mode := "octet"
	if q.Mode != "" {
		mode = q.Mode
	}

	cap := 2 + 2 + len(q.Filename) + 1 + len(mode) + 1

	b := new(bytes.Buffer)
	b.Grow(cap)

	err := binary.Write(b, binary.BigEndian, OP_RRQ)
	if err != nil {
		return nil, err
	}

	_, err = b.WriteString(q.Filename)
	if err != nil {
		return nil, err
	}
	err = b.WriteByte(0)
	if err != nil {
		return nil, err
	}

	_, err = b.WriteString(mode)
	if err != nil {
		return nil, err
	}

	err = b.WriteByte(0)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (q *ReadReq) UnmarshalBinary(p []byte) error {
	r := bytes.NewBuffer(p)

	var code OpCode
	// OpCode를 읽는다
	err := binary.Read(r, binary.BigEndian, &code)
	if err != nil {
		return err
	}

	if code != OP_RRQ {
		return errors.New("invalid RRQ")
	}

	// 파일명을 읽는다.
	q.Filename, err = r.ReadString(0)
	if err != nil {
		return errors.New("invalid RRQ")
	}

	// 0바이트 제거
	q.Filename = strings.TrimRight(q.Filename, "\x00")
	if len(q.Filename) == 0 {
		return errors.New("invalid RRQ")
	}

	// 모드 정보 읽기
	q.Mode, err = r.ReadString(0)
	if err != nil {
		return errors.New("invalid RRQ")
	}

	// 0바이트 제거
	q.Mode = strings.TrimRight(q.Mode, "\x00")
	if len(q.Mode) == 0 {
		return errors.New("invalid RRQ")
	}

	actual := strings.ToLower(q.Mode)
	if actual != "octet" {
		return errors.New("only binary transfers supported")
	}

	return nil
}

// | OpCode(2B) | Block #(2B) | Data |
type Data struct {
	Block   uint16
	Payload io.Reader
}

func (d *Data) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	b.Grow(DATAGRAM_SIZE)

	d.Block++
	// OpCode 쓰기
	err := binary.Write(b, binary.BigEndian, OP_DATA)
	if err != nil {
		return nil, err
	}

	// 블록 번호 쓰기
	err = binary.Write(b, binary.BigEndian, d.Block)
	if err != nil {
		return nil, err
	}

	// BlockSize 크기만큼 쓰기
	_, err = io.CopyN(b, d.Payload, BLOCK_SIZE)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return b.Bytes(), nil
}

func (d *Data) UnmarshalBinary(p []byte) error {
	if l := len(p); l < 4 || l > DATAGRAM_SIZE {
		return errors.New("invalid DATA")
	}

	var opcode OpCode

	err := binary.Read(bytes.NewBuffer(p[:2]), binary.BigEndian, &opcode)
	if err != nil || opcode != OP_DATA {
		return errors.New("invalid DATA")
	}

	err = binary.Read(bytes.NewBuffer(p[2:4]), binary.BigEndian, &d.Block)
	if err != nil {
		return errors.New("invalid DATA")
	}

	d.Payload = bytes.NewBuffer(p[4:])

	return nil
}

// OpCode(2B) | Block #(2B)
type Ack uint16

func (a Ack) MarshalBinary() ([]byte, error) {
	cap := 2 + 2

	b := new(bytes.Buffer)
	b.Grow(cap)

	err := binary.Write(b, binary.BigEndian, OP_ACK)
	if err != nil {
		return nil, err
	}

	err = binary.Write(b, binary.BigEndian, a)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (a *Ack) UnmarshalBinary(p []byte) error {
	var code OpCode

	r := bytes.NewReader(p)

	// OpCode 읽기
	err := binary.Read(r, binary.BigEndian, &code)
	if err != nil {
		return err
	}

	if code != OP_ACK {
		return errors.New("invalid ACK")
	}

	// 블록번호 읽기
	return binary.Read(r, binary.BigEndian, a)
}

// | OpCode(2B) | ErrorCode(2B) | ErrMsg | 0 |
type Err struct {
	Error   ErrCode
	Message string
}

func (e Err) MarshalBinary() ([]byte, error) {
	cap := 2 + 2 + len(e.Message) + 1

	b := new(bytes.Buffer)
	b.Grow(cap)

	// OpCode 쓰기
	err := binary.Write(b, binary.BigEndian, OP_ERR)
	if err != nil {
		return nil, err
	}

	// 에러코드 쓰기
	err = binary.Write(b, binary.BigEndian, e.Error)
	if err != nil {
		return nil, err
	}

	// 메시지 쓰기
	_, err = b.WriteString(e.Message)
	if err != nil {
		return nil, err
	}

	// 0 바이트 쓰기
	err = b.WriteByte(0)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (e *Err) UnmarshalBinary(p []byte) error {
	r := bytes.NewBuffer(p)

	var code OpCode
	// OpCode 읽기
	err := binary.Read(r, binary.BigEndian, &code)
	if err != nil {
		return err
	}

	if code != OP_ERR {
		return errors.New("invalid ERR")
	}

	// 에러 메시지 읽기
	err = binary.Read(r, binary.BigEndian, &e.Error)
	if err != nil {
		return err
	}

	e.Message, err = r.ReadString(0)
	if err != nil {
		return err
	}

	// 0바이트 제거
	e.Message = strings.TrimRight(e.Message, "\x00")

	return nil
}
