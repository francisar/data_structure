package main

import (
	"bytes"
	"github.com/francisar/data_structure"
	"strconv"
)

type term struct {
	Key []byte
	value uint64
}

func (t *term) LessThan(v data_structure.OPItem) bool {
	item := v.(*term)
	return bytes.Compare(t.Key, item.Key) < 0
}

func (t *term) Equal(v data_structure.OPItem) bool {
	item := v.(*term)
	return bytes.Compare(t.Key, item.Key) == 0
}

func (t *term) MoreThan(v data_structure.OPItem) bool {
	item := v.(*term)
	return bytes.Compare(t.Key, item.Key) > 0
}

func (t *term) DeepCopy(v data_structure.OPItem) {
	return
}

func (t *term) Marshal() ([]byte, error) {
	return t.Key, nil
}

func (t *term) UnMarshal(str string) error {
	t.Key = []byte(str)
	return nil
}

func (t *term) String() string {
	return strconv.FormatUint(t.value, 10)
}

func newTerm(key string, value uint64) data_structure.OPItem {
	t := &term{
		Key: []byte(key),
		value: value,
	}
	return t
}
