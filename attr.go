package dom

/*
* Attr implementation
*
* Copyright (c) 2009, Jeff Schiller
*/

import "xml";

type _attr struct {
  *_node;
  v string; // value (for attr)
}

func (a *_attr) NodeValue() string { return a.v; }
func (a *_attr) AppendChild(n Node) Node { return n; }
func (a *_attr) RemoveChild(n Node) Node { return n; }
func (a *_attr) ParentNode() Node { return Node(nil); }
func (a *_attr) ChildNodes() NodeList { return NodeList(nil); }
func (a *_attr) Attributes() NamedNodeMap { return NamedNodeMap(nil); }

func newAttr(name string, val string) (*_attr) {
  node := newNode(2);
  node.n = xml.Name{"", name};
  return &_attr { node, val };
}