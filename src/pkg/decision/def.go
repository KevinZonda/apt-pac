package decision

import (
	"errors"
)

type CheckTree struct {
	Key      string
	Data     *string
	Children map[string]*CheckTree
}

func (t *CheckTree) AddKey(key string) error {
	return t.Add(&CheckTree{Key: key})
}

func (t *CheckTree) AddKeyAndValue(key, value string) error {
	return t.Add(&CheckTree{Key: key, Data: &value})
}

func (t *CheckTree) Add(v *CheckTree) error {
	if t == nil {
		return errors.New("empty tree")
	}
	if t.Children == nil {
		t.Children = make(map[string]*CheckTree)
	}
	t.Children[v.Key] = v
	return nil
}

func (t *CheckTree) AddToSub(keys []string, data string) {
	if len(keys) == 1 {
		n := t.GetNode(keys[0])
		if n == nil {
			_ = t.AddKeyAndValue(keys[0], data)
			return
		}
		n.Data = &data
		return
	}
	curr := t
	index := 0
	for i, key := range keys {
		g := curr.GetNode(key)
		// fmt.Printf("GET: %s -> %s -> %v\n", curr.Key, key, g == nil)
		index = i
		if g != nil {
			curr = g
			continue
		}
		break
	}
	toAdd := keys[index:]
	for _, key := range toAdd {
		// fmt.Printf("ADD: %s -> %s\n", curr.Key, key)
		_ = curr.AddKey(key)
		curr = curr.GetNode(key)
	}
	curr.Data = &data
}

func (r *CheckTree) GetNode(key string) *CheckTree {
	if r == nil {
		return nil
	}
	c, ok := r.Children[key]
	if !ok {
		return nil
	}
	return c
}

func (r *CheckTree) Next(key string) (next *CheckTree, isMatchGeneric bool) {
	c, ok := r.Children[key]
	if ok {
		return c, false
	}
	if r.Key == "$$" {
		return r, true
	}
	c, ok = r.Children["$$"]
	if ok {
		return c, true
	}
	c, ok = r.Children["$"]
	return c, true
}
