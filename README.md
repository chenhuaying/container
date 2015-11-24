Container for Go
================
## Key interface
```
type Comparer interface {
	Less(x Comparer) bool
}

```

define int key like this
```
type KeyInt int

func (k KeyInt) Less(x Comparer) bool {
	return k < x.(KeyInt)
}
```

## SkipList feature
### SkipList
1. NewSkipList
2. DeleteNode
3. First, get the key like cpp
4. Second, get the value like cpp
5. Insert
6. IsEmpty
7. LowerBoundNode, find the lower bound node, like cpp map, key >= find key
8. Search, SearchNode

### SkipListNode
1. Key, get the key
2. Value, get the value
3. Next, the next node
4. Prev, the prev node
5. NewSkipListNode

### Others
1. MaxLevel 32
2. *level i+1*/*level i* = 0.25
