package fuzz

import (
	"bytes"

	"github.com/pkg/errors"
	capnp "zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/pogs"

	"github.com/Oncilla/go-capnproto2-fuzz/proto"
)

func Fuzz(data []byte) int {
	if _, err := Parse(data); err != nil {
		return 0
	}
	return 1
}

type List struct {
	Entries []Entry
}

type Entry struct {
	Data []byte
}

func Parse(data []byte) (*List, error) {
	list := &List{}
	if err := parse(list, data); err != nil {
		return nil, err
	}
	return list, nil
}

func parse(l *List, data []byte) error {
	msg, err := capnp.NewPackedDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return errors.New("Failed to decode capnp message")
	}
	rootPtr, err := msg.RootPtr()
	if err != nil {
		return errors.New("Failed to get root pointer from capnp message")
	}
	if err := pogs.Extract(l, proto.L_TypeID, rootPtr.Struct()); err != nil {
		return errors.New("Failed to extract struct from capnp message")
	}
	return nil
}

func (l *List) Pack() ([]byte, error) {
	msg, arena, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, errors.New("Failed to create new capnp message")
	}
	s, err := proto.NewRootL(arena)
	if err != nil {
		return nil, err
	}
	if err := pogs.Insert(proto.L_TypeID, s.Struct, l); err != nil {
		return nil, errors.New(err.Error())
	}
	raw, err := msg.MarshalPacked()
	if err != nil {
		return nil, errors.New("Failed to marshal capnp struct")
	}
	return raw, nil
}
