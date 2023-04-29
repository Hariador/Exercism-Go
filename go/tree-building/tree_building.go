package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

type recordList []Record

func (r recordList) Less(i, j int) bool {
	return r[i].ID < r[j].ID
}
func (r recordList) Len() int {
	return len(r)
}
func (r recordList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func Build(records recordList) (*Node, error) {
	sort.Sort(records)
	if len(records) == 0 {
		return nil, nil
	}
	if records[0].ID != records[0].Parent {
		return nil, errors.New("no root node")
	}
	root := Node{records[0].ID, []*Node{}}
	for _, r := range records[1:] {
		if r.ID >= len(records) {
			return nil, errors.New("stupid restriction ahoy")
		}
		if r.ID == r.Parent {
			return nil, errors.New("cannot refer to self")
		}
		if r.Parent > r.ID {
			return nil, errors.New("invalid record")
		}
		t, err := root.getNodeById(r.Parent, 0)
		if err != nil {
			return nil, err
		}
		if t.ID == r.ID {
			return nil, errors.New("duplicate node")
		}
		err = t.addChild(&Node{r.ID, []*Node{}})
		if err != nil {
			return nil, err
		}
	}

	return &root, nil

}

func (n *Node) getNodeById(id, depth int) (*Node, error) {
	if depth > 5 {
		return nil, errors.New("max depth exceeded")
	}
	if n.ID == id {
		return n, nil
	}
	for _, child := range n.Children {
		res, err := child.getNodeById(id, depth+1)
		if err != nil {
			return nil, err
		}
		if res != nil {
			return res, nil
		}
	}
	return nil, nil
}

func (n *Node) addChild(child *Node) error {
	t, err := n.getNodeById(child.ID, 1)
	if err != nil {
		return err
	}
	if t != nil {
		return errors.New("duplicate node")
	}
	n.Children = append(n.Children, child)
	return nil
}
