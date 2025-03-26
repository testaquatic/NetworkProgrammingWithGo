package protobuf

import (
	"io"

	"github.com/testaquatic/NetworkProgrammingWithGo/ch12/housework/v1"
	"google.golang.org/protobuf/proto"
)

func Load(r io.Reader) ([]*housework.Chore, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var chores housework.Chores

	return chores.Chores, proto.Unmarshal(b, &chores)
}

func Flush(w io.Writer, chores []*housework.Chore) error {
	b, err := proto.Marshal(&housework.Chores{Chores: chores})
	if err != nil {
		return err
	}
	_, err = w.Write(b)

	return err
}
