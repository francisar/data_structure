package main

import (
	"fmt"
	rb_tree "github.com/francisar/data_structure/tree/rb-tree"
	"math/rand"
	"strconv"
	"time"
)

type Int struct {
	value int
}

func NewInt(value int) rb_tree.RBItem {
	i := Int{
		value: value,
	}
	return &i
}
func (i *Int) LessThan(v rb_tree.RBItem) bool {
	return i.value < (v.(*Int)).value
}

func (i *Int) Equal(v rb_tree.RBItem) bool {
	return i.value == v.(*Int).value
}

func (i *Int) MoreThan(v rb_tree.RBItem) bool {
	return i.value > (v.(*Int)).value
}

func (i *Int) DeepCopy(v rb_tree.RBItem) {
	i.value = (v.(*Int)).value
}

func (i *Int) Marshal() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *Int) String() string {
	return fmt.Sprintf("%d", i.value)
}

func (i *Int) UnMarshal(str string) error {
	value, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	i.value = value
	return nil
}

func main() {
	rbTree := rb_tree.NewRBTree()
	totalNode := 25
	var nodes = make([]int, totalNode)
	for i := 0; i < totalNode; i++ {
		value := rand.Intn(100)
		err := rbTree.Insert(NewInt(value))
		if err != nil {

			panic(err)
		}
		nodes[i] = value
	}
	rbTree.PrintTree()
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		index := rand.Intn(totalNode)
		value := nodes[index]
		removeValue := fmt.Sprintf("remove %d", value)
		println(removeValue)
		err := rbTree.Delete(NewInt(value))
		if err != nil {
			continue
		} else {
			rbTree.PrintTree()
		}

	}
}
