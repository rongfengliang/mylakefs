package catalog

import (
	"github.com/treeverse/lakefs/pkg/graveler"
	"github.com/treeverse/lakefs/pkg/ident"
	"google.golang.org/protobuf/proto"
)

func ValueToEntry(value *graveler.Value) (*Entry, error) {
	if value == nil {
		return nil, nil
	}
	var ent Entry
	err := proto.Unmarshal(value.Data, &ent)
	if err != nil {
		return nil, err
	}
	return &ent, nil
}

func EntryToValue(entry *Entry) (*graveler.Value, error) {
	// marshal data using pb
	data, err := proto.Marshal(entry)
	if err != nil {
		return nil, err
	}
	// calculate entry identity
	checksum := ident.NewAddressWriter().
		MarshalInt64(entry.Size).
		MarshalString(entry.ETag).
		MarshalStringMap(entry.Metadata).
		MarshalStringOpt(entry.ContentType). // optional in order to keep identity of old entries without content-type
		Identity()
	return &graveler.Value{
		Identity: checksum,
		Data:     data,
	}, nil
}

func MustEntryToValue(entry *Entry) *graveler.Value {
	if entry == nil {
		return nil
	}
	v, err := EntryToValue(entry)
	if err != nil {
		panic(err)
	}
	return v
}
