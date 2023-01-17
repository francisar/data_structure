package skiplist

import (
	"errors"
	"fmt"
	"github.com/francisar/data_structure"
	"math/rand"
	"sync"
)

type SkipList struct {
	MaxLevel int
	mutex      sync.RWMutex
	Skip float32
	Count uint
	Header *Node
	update []*Node
}

func (s *SkipList)Find(item data_structure.OPItem) (data_structure.OPItem, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	current := s.Header
	last := current
	for l := current.Level; l >0; l-- {
		for current=last;current!=nil; {
			if current.Item.Equal(item) {
				return current.Item, nil
			} else if current.Item.LessThan(item) {
				last = current
				current = current.Next[l]
			} else {
				break
			}
		}
	}
	return nil, errors.New("can not find")
}

func (s *SkipList)newNode(item data_structure.OPItem) *Node {
	return &Node{
		Next: make([]*Node, s.MaxLevel),
		Item: item,
		Level: 0,
		Pre: nil,
	}
}

func (s *SkipList)Insert(item data_structure.OPItem)  {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	last := s.Header
	currentLevel := 0
	if last != nil {
		currentLevel = last.Level
	}
	for l := currentLevel; l >=0; l-- {
		for current:=last;current!=nil; {
			if current.Item.Equal(item) {
				current.Item.DeepCopy(item)

				return
			} else if current.Item.LessThan(item) {
				last = current
				current = current.Next[l]
			} else {
				break
			}
		}
		s.update[l] = last
	}
	node := s.newNode(item)
	level := s.randomLevel()
	if s.Header != nil && level > s.Header.Level {
		level = s.Header.Level +1
	}
	node.Level = level
	s.Count ++
	if last == nil {
		s.Header = node
	} else {
		node.Next[0] = last.Next[0]
		last.Next[0] = node
		node.Pre = last
		if node.Next[0] != nil {
			node.Next[0].Pre = node
		}
	}
	if last == s.Header {
		if last.Item.MoreThan(item) {
			node.Item = last.Item
			last.Item = item
		}
	}
	currentIndex := s.Header
	for i:=level;i > 0; i-- {
		if s.update[i] != nil {
			currentIndex = s.update[i]
		}
		node.Next[i] = currentIndex.Next[i]
		currentIndex.Next[i] = node
		if currentIndex.Level < i {
			currentIndex.Level = i
		}
		s.update[i] = nil
	}
}

func (s *SkipList)Delete(item data_structure.OPItem) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	last := s.Header
	currentLevel := 0
	if last != nil {
		currentLevel = last.Level
	}
	var deleteNode *Node = nil
	for l := currentLevel; l >=0; l-- {
		for current:=last;current!=nil; {
			if current.Item.Equal(item) {

				deleteNode = current
				break
			} else if current.Item.LessThan(item) {
				last = current
				current = current.Next[l]
			} else {
				break
			}
		}
		s.update[l] = last
	}
	if deleteNode == nil {
		return errors.New("can not find item")
	}
	if deleteNode == s.Header {
		if s.Header.Next[0] != nil {
			deleteNode = s.Header.Next[0]
			s.Header.Item = deleteNode.Item
			last = s.Header
			last.Pre = nil
		} else {
			s.Header = nil
			return nil
		}
	}
	deleteNode.Pre.Next[0] = deleteNode.Next[0]
	deleteNode.Next[0].Pre = deleteNode.Pre
	s.Count --
	level := s.Header.Level
	for i:=level;i > 0; i-- {
		if s.update[i] != nil && s.update[i].Next[i] == deleteNode {
			s.update[i].Next[i] = deleteNode.Next[i]
		}
		s.update[i] = nil
	}
	return nil
}

func (s *SkipList)randomLevel() int {
	level := 0
	for rand.Float32() < s.Skip && level < s.MaxLevel {
		level++
	}
	return level
}



func (s *SkipList)PrintSkipList() {
	printArray := make([]string, s.MaxLevel)
	for current := s.Header; current != nil; current = current.Next[0]  {
		for l := s.MaxLevel-1; l>=0; l-- {
			var str string
			 if l == 0 {
				str =  fmt.Sprintf("%s ",current.Item.String())
			} else if current.Next[l] == nil {
				str = "---"
			}  else {
				str = fmt.Sprintf(">%s", current.Next[l].Item.String())
			}
			printArray[l] += str
		}
	}
	for l := s.MaxLevel-1; l>=0; l-- {
		fmt.Println(printArray[l])
	}
}


func (s *SkipList)PrintList() {

	for currentNode := s.Header; currentNode !=nil; currentNode = currentNode.Next[0] {
		msg := fmt.Sprintf("Current:%s", currentNode.Item.String())
		if currentNode.Pre != nil {
			msg = fmt.Sprintf("%s Pre:%s", msg, currentNode.Pre.Item.String())
		}
		println(msg)
	}
}

func NewSkipList(maxLevel int, skip float32) *SkipList {
	return &SkipList{
		MaxLevel: maxLevel,
		Skip: skip,
		Count: 0,
		Header: nil,
		update: make([]*Node, maxLevel),
	}
}