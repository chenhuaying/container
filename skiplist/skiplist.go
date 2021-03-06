package skiplist

import (
	"math/rand"
	"time"

	"github.com/chenhuaying/container"
)

type SkipListLevel struct {
	forward *SkipListNode
}

type SkipListNode struct {
	key      container.Comparer
	value    interface{}
	backword *SkipListNode
	level    []SkipListLevel
}

type SkipList struct {
	header *SkipListNode
	tail   *SkipListNode
	length int
	level  int
}

const (
	MaxLevel  = 32
	SkipListP = 0.25
)

func init() {
	rand.Seed(time.Now().Unix())
}

func NewSkipListNode(level int, key container.Comparer, value interface{}) *SkipListNode {
	node := &SkipListNode{key: key, value: value, level: make([]SkipListLevel, level)}
	return node
}

func NewSkipList() *SkipList {
	header := NewSkipListNode(MaxLevel, nil, nil)
	skiplist := &SkipList{header: header, tail: nil, length: 0, level: 1}
	if skiplist != nil {
		for i := 0; i < MaxLevel; i++ {
			skiplist.header.level[i].forward = nil
		}
	}
	return skiplist
}

// Just for test
func randomLevel() int {
	level := 1

	for rand.Float32() < 0.25 && level < MaxLevel {
		level += 1
	}
	return level
}

func (l *SkipList) randomLevel() int {
	level := 1

	for rand.Float32() < 0.25 && level < MaxLevel {
		level += 1
	}
	return level
}

func (l *SkipList) Insert(key container.Comparer, value interface{}) {
	update := [MaxLevel]*SkipListNode{}
	x := l.header
	for i := l.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && x.level[i].forward.key.Less(key) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward
	if x != nil && x.key == key {
		x.value = value
	} else {
		level := l.randomLevel()
		if level > l.level {
			for i := l.level; i < level; i++ {
				update[i] = l.header
			}
			// update skiplist max level
			l.level = level
		}

		x = NewSkipListNode(level, key, value)
		for i := 0; i < level; i++ {
			x.level[i].forward = update[i].level[i].forward
			update[i].level[i].forward = x
		}
		l.length += 1
	}

	if update[0] == l.header {
		x.backword = nil
	} else {
		x.backword = update[0]
	}

	if x.level[0].forward != nil {
		x.level[0].forward.backword = x
	} else {
		// insert to the last position of the list
		l.tail = x
	}
}

func (l *SkipList) SearchNode(key container.Comparer) *SkipListNode {
	x := l.header
	for i := l.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && x.level[i].forward.key.Less(key) {
			x = x.level[i].forward
		}
	}

	x = x.level[0].forward
	if x != nil && x.key == key {
		return x
	}

	return nil
}

func (l *SkipList) Search(key container.Comparer) interface{} {
	node := l.SearchNode(key)
	if node != nil {
		return node.value
	}
	return nil
}

func (l *SkipList) LowerBoundNode(key container.Comparer) *SkipListNode {
	x := l.header
	for i := l.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && x.level[i].forward.key.Less(key) {
			x = x.level[i].forward
		}
	}
	x = x.level[0].forward
	return x
}

func (l *SkipList) LowerBoundNodeFn(key container.Comparer, fn func(x, y container.Comparer) bool) *SkipListNode {
	x := l.header
	for i := l.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && fn(x.level[i].forward.key, key) {
			x = x.level[i].forward
		}
	}
	x = x.level[0].forward
	return x
}

func (l *SkipList) DeleteNode(key container.Comparer) *SkipListNode {
	update := [MaxLevel]*SkipListNode{}
	x := l.header
	for i := l.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && x.level[i].forward.key.Less(key) {
			x = x.level[i].forward
		}
		update[i] = x
	}
	x = x.level[0].forward
	if x != nil && x.key == key {
		for i := 0; i < l.level; i++ {
			if update[i].level[i].forward != x {
				break
			} else {
				update[i].level[i].forward = x.level[i].forward
			}
		}

		if x.level[0].forward == nil {
			l.tail = x.backword
		} else {
			x.level[0].forward.backword = x.backword
		}

		// if x is the hightest node
		for l.level > 1 && l.header.level[l.level-1].forward == nil {
			l.level -= 1
		}
		l.length -= 1

		return x
	} else {
		return nil
	}
}

// ===========================functional utils=====================
func (l *SkipList) IsEmpty() bool {
	return l.length == 0
}

func (l *SkipList) First() *SkipListNode {
	return l.header.level[0].forward
}

func (n *SkipListNode) Next() *SkipListNode {
	return n.level[0].forward
}

func (n *SkipListNode) Prev() *SkipListNode {
	return n.backword
}

func (n *SkipListNode) Key() container.Comparer {
	return n.key
}

func (n *SkipListNode) Value() interface{} {
	return n.value
}

func (l *SkipList) Length() int {
	return l.length
}
