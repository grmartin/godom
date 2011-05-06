package dom

/*
 * NodeList implementations
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 */

// A _childNodelist only stores a reference to its parent node.
// This way the list can be live, each time Length() or Item is
// called, fresh results are returned.
type _childNodelist struct {
  p *_node;
}

func (nl *_childNodelist) Length() uint {
  return uint(nl.p.c.Len());
}
func (nl *_childNodelist) Item(index uint) Node {
  n := Node(nil);
  if index < uint(nl.p.c.Len()) {
    // TODO: what if index == nl.p.c.Len() -1 and a node is deleted right now?
    n = nl.p.c.At(int(index)).(Node);
  }
  return n;
}

func newChildNodelist(p *_node) (*_childNodelist) {
  nl := new(_childNodelist);
  nl.p = p;
  return nl;
}

// TODO: Find a home for this function.  It operates only on interface types.
/**
 * Walks the tree of nodes in a depth-first manner, calling the
 * function f on each of the children of the passed in node.
 */
func walkTreeDepthFirst(n Node, f func(Node)) {
  childNodes := n.ChildNodes();
  numChildren := childNodes.Length();
  var ix uint
  for ix = 0; ix < numChildren; ix++ {
    child := childNodes.Item(ix)
    f(child)
    walkTreeDepthFirst(child, f)
  }
}

// A _tagNodeList only stores a reference to the node and the tagname 
// on which getElementsByTagName() was called so that the list can be 
// live.  TODO: Do we really query every time or can we cache the results
// somehow?
type _tagNodeList struct {
  e *_elem;
  tag string
}

func (nl *_tagNodeList) Length() uint {
  parentElement := nl.e
  var count uint = 0
  walkTreeDepthFirst(parentElement, func(n Node) {
    if n.NodeType() == 1 {
      if nl.tag == "*" || nl.tag == n.(Element).TagName() {
        count++;
      }
    }
  })
  return count;
}

func (nl *_tagNodeList) Item(index uint) Node {
  var count uint = 0
  parentElement := nl.e
  foundNode := Node(nil)

  walkTreeDepthFirst(parentElement, func(n Node) {
    if n.NodeType() == 1 {
      if nl.tag == "*" || nl.tag == n.(Element).TagName() {
        if count == index {
          foundNode = n
        }
        count++
      }
    }
  })

  return foundNode;
}

func newTagNodeList(p *_elem, t string) (*_tagNodeList) {
  nl := new(_tagNodeList);
  nl.e = p;
  nl.tag = t
  return nl;
}

