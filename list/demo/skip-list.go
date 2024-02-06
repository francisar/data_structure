package main

import (
	"fmt"
	"github.com/francisar/data_structure"
	"github.com/francisar/data_structure/list/skiplist"
	"strconv"
)

type Int struct {
	value int
}

func NewInt(value int) data_structure.OPItem {
	i := Int{
		value: value,
	}
	return &i
}
func (i *Int) LessThan(v data_structure.OPItem) bool {
	return i.value < (v.(*Int)).value
}

func (i *Int) Equal(v data_structure.OPItem) bool {
	return i.value == v.(*Int).value
}

func (i *Int) MoreThan(v data_structure.OPItem) bool {
	return i.value > (v.(*Int)).value
}

func (i *Int) DeepCopy(v data_structure.OPItem) {
	i.value = (v.(*Int)).value
}

func (i *Int) Marshal() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *Int) String() string {
	return fmt.Sprintf("%2d", i.value)
}

func (i *Int) UnMarshal(str string) error {
	value, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	i.value = value
	return nil
}

func main()  {
	skipList := skiplist.NewSkipList(6,0.5)
	/*
	for i, data := range datas {
		t := newTerm(data, uint64(i))
		skipList.Insert(t)
	}
	
	 */
	skipList.PrintSkipList()


}