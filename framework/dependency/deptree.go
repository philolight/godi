package dependency

import (
	"reflect"
	"fmt"
)

type node struct{
	name string
	bean interface{}
	kids map[string]*node
}

func newNode(name string, bean interface{}) *node{
	return &node{
		name : name,
		bean : bean,
		kids : make(map[string]*node, 0),
	}
}

func (n *node) Add(kid *node, field string){
	n.kids[field] = kid
}

func (n *node) Call(funcName string, called *map[string]bool){
	for _, kid := range n.kids {
		kid.Call(funcName, called)
	}

	f := reflect.ValueOf(n.bean).MethodByName(funcName)
	if f.Kind() == reflect.Func && !(*called)[n.name]{
		f.Call(nil)
		(*called)[n.name] = true
	}
}

func (n *node) Find(name string) *node{
	if n.name == name {
		return n
	}

	for _, kid := range n.kids{
		if ret := kid.Find(name); ret != nil {
			return ret
		}
	}

	return nil
}

func (n *node) Print(fieldName string, indent string, unitIndent string){
	if fieldName == "" {
		fmt.Printf("%s-%s\n", indent, n.name)
	} else {
		fmt.Printf("%s-%s=%s\n", indent, fieldName, n.name)
	}
	for key, kid := range n.kids{
		kid.Print(key, indent + unitIndent, unitIndent)
	}
}

type treeContainer struct {
	roots map[string]*node
}

func newTreeContainer() *treeContainer{
	return &treeContainer{
		roots: make(map[string]*node, 0),
	}
}

func (t *treeContainer) Add(pname string, pbean interface{}, kname string, kbean interface{}, field string) {
	pnode := t.roots[pname]
	if pnode == nil {
		pnode = newNode(pname, pbean)
		t.roots[pname] = pnode
	}
	knode := t.roots[kname]
	if knode == nil {
		knode = newNode(kname, kbean)
	}
	pnode.Add(knode, field)
	delete(t.roots, kname)
}

func (t *treeContainer) Call(funcName string){
	called := make(map[string]bool, 0)
	for _, root := range t.roots{
		root.Call(funcName, &called)
	}
}

func (t *treeContainer) Print(){
	fmt.Println("<<<<<<<< Object Diagram >>>>>>>")

	for _, root := range t.roots{
		root.Print("","", "   ")
	}
}