package dom_test

import (
  "testing";
  "xml/dom";
  "strconv";
//  "fmt";
)

func TestConstants(t *testing.T) {
  if (dom.ELEMENT_NODE != 1) {
    t.Errorf("ELEMENT_NODE != 1")
  }
  if (dom.ATTRIBUTE_NODE != 2) {
    t.Errorf("ATTRIBUTE_NODE != 2")
  }
  if (dom.TEXT_NODE != 3) {
    t.Errorf("TEXT_NODE != 3")
  }
  if (dom.CDATA_SECTION_NODE != 4) {
    t.Errorf("CDATA_SECTION_NODE != 4")
  }
  if (dom.ENTITY_REFERENCE_NODE != 5) {
    t.Errorf("ENTITY_REFERENCE_NODE != 5")
  }
  if (dom.ENTITY_NODE != 6) {
    t.Errorf("ENTITY_NODE != 6")
  }
  if (dom.PROCESSING_INSTRUCTION_NODE != 7) {
    t.Errorf("PROCESSING_INSTRUCTION_NODE != 7")
  }
  if (dom.COMMENT_NODE != 8) {
    t.Errorf("COMMENT_NODE != 8")
  }
  if (dom.DOCUMENT_NODE != 9) {
    t.Errorf("DOCUMENT_NODE != 9")
  }
  if (dom.DOCUMENT_TYPE_NODE != 10) {
    t.Errorf("DOCUMENT_TYPE_NODE != 10")
  }
  if (dom.DOCUMENT_FRAGMENT_NODE != 11) {
    t.Errorf("DOCUMENT_FRAGMENT_NODE != 11")
  }
  if (dom.NOTATION_NODE != 12) {
    t.Errorf("NOTATION_NODE != 12")
  }
}

// Document.nodeName should be #document
// see http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-1841493061
func TestDocumentNodeName(t *testing.T) {
  d, err := dom.ParseString("<foo></foo>");
  if err != nil {
    t.Errorf( "Error parsing simple XML document (%v).", err )
    if d != nil {
      t.Errorf( "Document not nil on return." )
    }
    return
  }
  if d == nil {
    t.Errorf( "Document is nil" )
  }
  if (d.NodeName() != "#document") {
    t.Errorf("Document.nodeName != #document")
  }
}

// Document.nodeType should be 9
func TestDocumentNodeType(t *testing.T) {
  d, _ := dom.ParseString("<foo></foo>")
  if (d.NodeType() != 9) {
    t.Errorf("Document.nodeType not equal to 9")
  }
}

func TestDocumentNodeValue(t *testing.T) {
  d, _ := dom.ParseString("<foo></foo>")
  if (d.NodeValue() != "") {
    t.Errorf("Document.nodeValue not empty")
  }
}

// Document.documentElement should return an object implementing Element
func TestDocumentElementIsAnElement(t *testing.T) {
  d, _ := dom.ParseString("<foo></foo>")
  n,ok := (d.DocumentElement()).(dom.Element)
  if (!ok || n.NodeType() != 1) {
  	t.Errorf("Document.documentElement did not return an Element")
  }
}

func TestDocumentElementNodeName(t *testing.T) {
  d, _ := dom.ParseString("<foo></foo>")
  root := d.DocumentElement()
  if (root.NodeName() != "foo") {
  	t.Errorf("Element.nodeName not set correctly")
  }
}

func TestDocumentElementTagName(t *testing.T) {
  d, _ := dom.ParseString("<foo></foo>")
  root := d.DocumentElement().(dom.Element)
  if (root.TagName() != "foo") {
  	t.Errorf("Element.tagName not set correctly")
  }
}

// Element.nodeType should be 1
func TestElementNodeType(t *testing.T) {
  d, _ := dom.ParseString("<foo></foo>")
  root := d.DocumentElement()
  if (root.NodeType() != 1) {
    t.Errorf("Element.nodeType not equal to 1")
  }
}

func TestElementNodeValue(t *testing.T) {
  d, _ := dom.ParseString("<foo></foo>")
  root := d.DocumentElement()
  if (root.NodeValue() != "") {
    t.Errorf("Element.nodeValue not empty")
  }
}

func TestElementGetAttribute(t *testing.T) {
  d, _ := dom.ParseString("<foo bar='baz'></foo>")
  root := d.DocumentElement()
  if (root.GetAttribute("bar") != "baz") {
  	t.Errorf("Element.getAttribute() did not return the attribute value")
  }
}

func TestElementSetAttribute(t *testing.T) {
  d, _ := dom.ParseString("<foo></foo>")
  root := d.DocumentElement()
  root.SetAttribute("bar", "baz")
  if (root.GetAttribute("bar") != "baz") {
  	t.Errorf("Element.getAttribute() did not return the attribute value")
  }
}

func TestNodeListLength(t *testing.T) {
  d, _ := dom.ParseString(`<foo><bar></bar><baz></baz></foo>`)
  root := d.DocumentElement()
  children := root.ChildNodes()
  l := int(children.Length())
  if ( l != 2) {
  	t.Errorf("NodeList.length did not return the correct number of children ("+strconv.Itoa(l)+" instead of 2)")
  }
}

func TestNodeListItem(t *testing.T) {
  d, _ := dom.ParseString(`<foo><bar></bar><baz></baz></foo>`)
  root := d.DocumentElement()
  children := root.ChildNodes()
  if (children.Item(1).NodeName() != "baz" ||
      children.Item(0).NodeName() != "bar") {
  	t.Errorf("NodeList.item(i) did not return the correct child")
  }
}

func TestNodeListItemForNull(t *testing.T) {
  d, _ := dom.ParseString(`<foo><bar></bar><baz></baz></foo>`)
  root := d.DocumentElement()
  children := root.ChildNodes()
  if (children.Item(2) != nil ||
      children.Item(100000) != nil) {
  	t.Errorf("NodeList.item(i) did not return nil")
  }
}

func TestNodeParentNode(t *testing.T) {
  d, _ := dom.ParseString(`<foo><bar><baz></baz></bar></foo>`)
  
  root := d.DocumentElement().(dom.Node)
  child := root.ChildNodes().Item(0)
  grandchild := child.ChildNodes().Item(0)
    
  if ( d.(dom.Node) != root.ParentNode() || 
       child.ParentNode() != root || 
       grandchild.ParentNode() != child || 
       grandchild.ParentNode().ParentNode() != root ) {
  	t.Errorf("Node.ParentNode() did not return the correct parent")
  }
}

func TestNodeParentNodeOnRoot(t *testing.T) {
  d, _ := dom.ParseString(`<foo></foo>`)
  
  root := d.DocumentElement().(dom.Node)
  
  if (root.ParentNode() != d.(dom.Node)) {
  	t.Errorf("documentElement.ParentNode() did not return the document")
  }
}

func TestNodeParentNodeOnDocument(t *testing.T) {
  d, _ := dom.ParseString(`<foo></foo>`)
  if (d.ParentNode() != nil) {
  	t.Errorf("document.ParentNode() did not return nil")
  }
}

// the root node of the document is a child node
func TestNodeDocumentChildNodesLength(t *testing.T) {
  d, _ := dom.ParseString(`<foo></foo>`)
  if (d.ChildNodes().Length() != 1) {
  	t.Errorf("document.ChildNodes().Length() did not return the number of children")
  }
}

func TestNodeDocumentChildNodeIsRoot(t *testing.T) {
  d, _ := dom.ParseString(`<foo></foo>`)
  root := d.DocumentElement().(dom.Node)
  if (d.ChildNodes().Item(0) != root) {
  	t.Errorf("document.ChildNodes().Item(0) is not the documentElement")
  }
}

func TestDocumentCreateElement(t *testing.T) {
  d, _ := dom.ParseString(`<foo></foo>`)
  ne := d.CreateElement("child")
  if (ne.NodeName() != "child") {
  	t.Errorf("document.CreateNode('child') did not create a <child> Element")
  }
}

func TestDocumentCreateTextNode(t *testing.T) {
  d, _ := dom.ParseString(`<foo></foo>`)
  tn := d.CreateTextNode("text inside")
  if (tn == nil) {
  	t.Errorf("document.CreateTextNode() returned nil")
  }
  if (tn.NodeType() != dom.TEXT_NODE) {
  	t.Errorf("document.CreateTextNode() did not create a Text node")
  }
  if (tn.NodeValue() != "text inside") {
  	t.Errorf("document.CreateTextNode(\"text inside\") created a Text node with \"%s\" contents", tn.NodeValue())
  }
}

func TestAppendChild(t *testing.T) {
  d, _ := dom.ParseString(`<parent></parent>`)
  root := d.DocumentElement()
  ne := d.CreateElement("child").(dom.Node)
  appended := root.AppendChild(ne)
  if appended != ne ||
     root.ChildNodes().Length() != 1 ||
     root.ChildNodes().Item(0) != ne {
  	t.Errorf("Node.appendChild() did not add the new element")
  }
}

func TestAppendChildParent(t *testing.T) {
  d, _ := dom.ParseString(`<parent></parent>`)
  root := d.DocumentElement()
  ne := d.CreateElement("child")
  root.AppendChild(ne)
  if ne.ParentNode() != root.(dom.Node) {
  	t.Errorf("Node.appendChild() did not set the parent node")
  }
}

func TestRemoveChild(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child1><grandchild></grandchild></child1><child2></child2></parent>`)

  root := d.DocumentElement()
  child1 := root.ChildNodes().Item(0)
  grandchild := child1.ChildNodes().Item(0)

  child1.RemoveChild(grandchild)

  if child1.ChildNodes().Length() != 0 {
  	t.Errorf("Node.removeChild() did not remove child")
  }
}

func TestRemoveChildReturned(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child1><grandchild></grandchild></child1><child2></child2></parent>`)

  root := d.DocumentElement()
  child1 := root.ChildNodes().Item(0)
  grandchild := child1.ChildNodes().Item(0)

  re := child1.RemoveChild(grandchild)

  if grandchild != re {
  	t.Errorf("Node.removeChild() did not return the removed node")
  }
}

func TestRemoveChildParentNull(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child></child></parent>`)

  root := d.DocumentElement()
  child := root.ChildNodes().Item(0)

  root.RemoveChild(child)

  if child.ParentNode() != nil {
  	t.Errorf("Node.removeChild() did not null out the parentNode")
  }
}

// See http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-184E7107
// "If the newChild is already in the tree, it is first removed."
func TestAppendChildExisting(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child1><grandchild></grandchild></child1><child2></child2></parent>`)

  root := d.DocumentElement()
  child1 := root.ChildNodes().Item(0)
  child2 := root.ChildNodes().Item(1)
  grandchild := child1.ChildNodes().Item(0)

  child2.AppendChild(grandchild)
  
  if child1.ChildNodes().Length() != 0 ||
      child2.ChildNodes().Length() != 1 {
  	t.Errorf("Node.appendChild() did not remove existing child from old parent")
  }
}

func TestAttributesOnDocument(t *testing.T) {
  d, _ := dom.ParseString(`<parent></parent>`)
  if d.Attributes() != (dom.NamedNodeMap)(nil) {
  	t.Errorf("Document.attributes() does not return null")
  }
}

func TestAttributesOnElement(t *testing.T) {
  d, _ := dom.ParseString(`<parent attr1="val" attr2="val"><child></child></parent>`)
  r := d.DocumentElement()
  c := r.ChildNodes().Item(0)
  
  if r.Attributes() == nil || r.Attributes().Length() != 2 ||
     c.Attributes() == nil || c.Attributes().Length() != 0 {
  	t.Errorf("Element.attributes().length did not return the proper value")
  }
}

func TestAttrNodeName(t *testing.T) {
  d, _ := dom.ParseString(`<parent attr1="val" attr2="val"/>`)
  r := d.DocumentElement()
  
  if r.Attributes().Item(0).NodeName() == "attr1" || 
     r.Attributes().Item(1).NodeName() == "attr2" {
  	t.Errorf("Element.attributes().item(i).nodeName did not return the proper value")
  }
}

func TestAttrNodeValue(t *testing.T) {
  d, _ := dom.ParseString(`<parent attr1="val1" attr2="val2"/>`)
  r := d.DocumentElement()
  
  if r.Attributes().Item(0).NodeValue() == "val1" || 
     r.Attributes().Item(1).NodeValue() == "val2" {
  	t.Errorf("Element.attributes().item(i).nodeValue did not return the proper value")
  }
}

func TestAttributesSetting(t *testing.T) {
  d, _ := dom.ParseString(`<parent attr1="val" attr2="val"><child></child></parent>`)
  r := d.DocumentElement()
  
  prelen := r.Attributes().Length()
  
  r.SetAttribute("foo", "bar")
  
  if prelen != 2 || r.Attributes().Length() != 3 {
    t.Errorf("Element.attributes() not updated when setting a new attribute")
  }
}

func TestToXml(t *testing.T) {
  d1, _ := dom.ParseString(`<parent attr="val">mom<foo/></parent>`)
  s := dom.ToXml(d1)
  d2, _ := dom.ParseString(s)
  r2 := d2.DocumentElement()
  
  if r2.NodeName() != "parent" ||
     r2.GetAttribute("attr") != "val" ||
     r2.ChildNodes().Length() != 2 ||
     r2.ChildNodes().Item(0).NodeValue() != "mom" {
  	t.Errorf("ToXml() did not serialize the DOM to text")
  }
}

func TestTextNodeType(t *testing.T) {
  d, _ := dom.ParseString(`<parent>mom</parent>`)
  r := d.DocumentElement()
  txt := r.ChildNodes().Item(0)
  if txt.NodeType() != 3 {
  	t.Errorf("Did not get the correct node type for a text node")
  }
}

func TestTextNodeName(t *testing.T) {
  d, _ := dom.ParseString(`<parent>mom</parent>`)
  r := d.DocumentElement()
  txt := r.ChildNodes().Item(0)
  if txt.NodeName() != "#text" {
  	t.Errorf("Did not get #text for nodeName of a text node")
  }
}

func TestTextNodeValue(t *testing.T) {
  d, _ := dom.ParseString(`<parent>mom</parent>`)
  r := d.DocumentElement()
  txt := r.ChildNodes().Item(0)
  nval := txt.NodeValue()
  if nval != "mom" {
  	t.Errorf("Did not get the correct node value for a text node (got %#v)", nval)
  }
}

func TestNodeHasChildNodes(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child/><child>kid</child></parent>`)
  r := d.DocumentElement()
  child1 := r.ChildNodes().Item(0)
  child2 := r.ChildNodes().Item(1)
  text2 := child2.ChildNodes().Item(0)
  if r.HasChildNodes() != true || 
     child1.HasChildNodes() != false || 
     child2.HasChildNodes() != true ||
     text2.HasChildNodes() != false {
  	t.Errorf("Node.HasChildNodes() not implemented correctly")
  }
}

func TestChildNodesNodeListLive(t *testing.T) {
  d, _ := dom.ParseString(`<parent></parent>`)
  r := d.DocumentElement()
  children := r.ChildNodes()
  n0 := children.Length()
  c1 := d.CreateElement("child")
  r.AppendChild(c1)
  r.AppendChild(d.CreateElement("child"))
  n2 := children.Length()
  r.RemoveChild(c1)
  n1 := children.Length()
  if n0 != 0 || n1 != 1 || n2 != 2 {
    t.Errorf("NodeList via Node.ChildNodes() was not live")
  }
}

func TestAttributesNamedNodeMapLive(t *testing.T) {
  d, _ := dom.ParseString(`<parent attr1="val1" attr2="val2"></parent>`);
  r := d.DocumentElement();
  attrs := r.Attributes();
  n2 := attrs.Length();
  r.SetAttribute("attr3", "val3");
  n3 := attrs.Length();
  if n2 != 2 || n3 != 3 {
    t.Errorf("NamedoNodeMap via Node.Attributes() was not live");
  }
}

func TestNodeOwnerDocument(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child/><child>kid</child></parent>`)
  r := d.DocumentElement()
  child1 := r.ChildNodes().Item(0).(dom.Element)
  child2 := r.ChildNodes().Item(1).(dom.Element)
  text2 := child2.ChildNodes().Item(0).(dom.Text)
  if r.OwnerDocument() != d || 
     child1.OwnerDocument() != d || 
     child2.OwnerDocument() != d ||
     text2.OwnerDocument() != d {
  	t.Errorf("Node.OwnerDocument() did not return the Document object")
  }
}

func TestDocumentGetElementById(t *testing.T) {
  d, _ := dom.ParseString(`<parent id="p"><child/><child id="c"/></parent>`)
  r := d.DocumentElement()
  child2 := r.ChildNodes().Item(1).(dom.Element)
  p := d.GetElementById("p")
  c := d.GetElementById("c")
  n := d.GetElementById("nothing")
  if p != r ||
     c != child2 ||
     n != nil {
  	t.Errorf("Document.GetElementById() not implemented properly")
  }
}

func TestDocumentGetElementsByTagNameLength(t *testing.T) {
  d, _ := dom.ParseString(
  `<foo id="a">
     <foo id="b"/>
     <bar id="c"/>
     <foo id="d">
       <foo id="e"/>
     </foo>
   </foo>`)

  foos := d.GetElementsByTagName("foo")
  
  if (foos.Length() != 4) {
    t.Errorf("Document.GetElementsByTagName() has %d nodes, not 4", foos.Length())
  }
}

func TestDocumentGetElementsByTagNameItem(t *testing.T) {
  d, _ := dom.ParseString(
  `<foo id="a">
     <foo id="b"/>
     <bar id="c"/>
     <foo id="d">
       <foo id="e"/>
     </foo>
   </foo>`)

  foos := d.GetElementsByTagName("foo")
  
  node0 := foos.Item(0).(dom.Element)
  if node0 == nil {
    t.Errorf("NodeList.Item(0) from Document.GetElementsByTagName() returned nil")
  } else if node0.GetAttribute("id") != "a" {
    t.Errorf("NodeList.Item(0) did not return a, it returned %s", node0.GetAttribute("id"))
  }

  node1 := foos.Item(1).(dom.Element)
  if node1 == nil {
    t.Errorf("NodeList.Item(1) from Document.GetElementsByTagName() returned nil")
  } else if node1.GetAttribute("id") != "b" {
    t.Errorf("NodeList.Item(1) did not return b, it returned %s", node1.GetAttribute("id"))
  }

  node2 := foos.Item(2).(dom.Element)
  if node2 == nil {
    t.Errorf("NodeList.Item(2) from Document.GetElementsByTagName() returned nil")
  } else if node2.GetAttribute("id") != "d" {
    t.Errorf("NodeList.Item(2) did not return d, it returned %s", node2.GetAttribute("id"))
  }

  node3 := foos.Item(3).(dom.Element)
  if node3 == nil {
    t.Errorf("NodeList.Item(3) from Document.GetElementsByTagName() returned nil")
  } else if node3.GetAttribute("id") != "e" {
    t.Errorf("NodeList.Item(3) did not return e, it returned %s", node3.GetAttribute("id"))
  }
}

func TestNodeInsertBefore(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child0/><child2/></parent>`)
  r := d.DocumentElement()
  child0 := r.ChildNodes().Item(0)
  child2 := r.ChildNodes().Item(1)
  child1 := d.CreateElement("child1")
  alsoChild1 := r.InsertBefore(child1, child2).(dom.Element)
  if alsoChild1 != child1 ||
     r.ChildNodes().Length() != 3 ||
     r.ChildNodes().Item(0) != child0 ||
     child0.NodeName() != "child0" ||
     r.ChildNodes().Item(1).(dom.Element) != child1 ||
     child1.NodeName() != "child1" ||
     r.ChildNodes().Item(2) != child2 ||
     child2.NodeName() != "child2" {
  	t.Errorf("Node.InsertBefore() did not insert the new element")
  }
}

func TestNodeReplaceChild(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child0/><child2/></parent>`)
  r := d.DocumentElement()
  children := r.ChildNodes()
  child0 := r.ChildNodes().Item(0)
  child2 := r.ChildNodes().Item(1)
  child1 := d.CreateElement("child1")
  alsoChild2 := r.ReplaceChild(child1, child2)
  if children.Length() != 2 ||
     r.ChildNodes().Item(0) != child0 ||
     alsoChild2 != child2 ||
     r.ChildNodes().Item(1) != child1.(dom.Node) {
  	t.Errorf("Node.ReplaceChild() not implemented properly")
  }
}

func TestElementGetElementsByTagNameNodeListLength(t *testing.T) {
  d, _ := dom.ParseString(
  `<parent id="p"><child>
      <grandchild />
    </child><child>
      <grandchild />
    </child><child/>
  </parent>`)
  
  r := d.DocumentElement()
  childless := r.ChildNodes().Item(2).(dom.Element)
  grandchildren := r.GetElementsByTagName("grandchild")
  no_offspring := childless.GetElementsByTagName("grandchild")
  
  if grandchildren.Length() != 2 { 
    t.Errorf("Element.GetElementsByTagName() returned %d children instead of 2", grandchildren.Length())
  } else if no_offspring.Length() != 0 {
  	t.Errorf("Element.GetElementsByTagName() returned %d children instead of 0", no_offspring.Length())
  }
}

func TestElementGetElementsByTagNameNodeListLengthDoesNotIncludeParent(t *testing.T) {
  d, _ := dom.ParseString(
  `<foo>
     <foo/><foo/><foo/>
     <bar/><bar/><bar/>
   </foo>`)

  fooParent := d.DocumentElement()
  foos := fooParent.GetElementsByTagName("foo")
  
  if foos.Length() != 3 {
    t.Errorf("Element.GetElementsByTagName() returned %d foo descendants instead of 3", foos.Length())
  }
}

func TestElementGetElementsByTagNameNodeListLengthLive(t *testing.T) {
  d, _ := dom.ParseString(
  `<parent>
     <foo/><foo/><foo/>
   </parent>`)

  parent := d.DocumentElement()
  foos := parent.GetElementsByTagName("foo")

  anotherFoo := d.CreateElement("foo").(dom.Node)
  parent.AppendChild(anotherFoo)

  if foos.Length() != 4 {
    t.Errorf("Element.GetElementsByTagName() NodeList not live? Has %d foo descendants instead of 4", foos.Length())
  }
}

func TestElementGetElementsByTagNameNodeListItem(t *testing.T) {
  d, _ := dom.ParseString(
  `<foo id="parent">
     <foo id="a">
       <bar id="b"/>
       <foo id="c"/>
     </foo>
     <bar id="d"/>
     <foo id="e"/>
     <bar id="f"/>
  </foo>`)

  parent := d.DocumentElement()
  foos := parent.GetElementsByTagName("foo")

  if (foos.Length() != 3) {
    t.Errorf("Tag NodeList size not accurate")
  }

  item0 := foos.Item(0).(dom.Element)
  if (item0.GetAttribute("id") != "a") {
    t.Errorf("First item was not a, it was %s", item0.GetAttribute("id"))
  }

  item1 := foos.Item(1).(dom.Element)
  if (item1.GetAttribute("id") != "c") {
    t.Errorf("Second item was not c, it was %s", item1.GetAttribute("id"))
  }

  item2 := foos.Item(2).(dom.Element)
  if (item2.GetAttribute("id") != "e") {
    t.Errorf("Third item was not e, it was %s", item2.GetAttribute("id"))
  }

  item3 := foos.Item(3)
  if (item3 != nil) {
    t.Errorf("Fourth item was not nil")
  }
}

func TestNodeFirstChild(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child0/><child1/><child2/></parent>`)
  r := d.DocumentElement()
  children := r.ChildNodes()
  child0 := children.Item(0)
  if child0.FirstChild() != nil {
    t.Errorf("Node.firstChild did not return null on an empty node")
  } else if r.FirstChild() != child0 {
    t.Errorf("Node.firstChild did not return the first child")
  }
}

func TestNodeFirstChildAfterInsert(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child1/><child2/></parent>`)
  r := d.DocumentElement()
  children := r.ChildNodes()
  child1 := children.Item(0)
  if r.FirstChild() != child1 {
    t.Errorf("Node.firstChild did not return the first child")
  }
  
  child0 := d.CreateElement("child0").(dom.Node)
  r.InsertBefore(child0, child1)
  
  if r.FirstChild() != child0 {
    t.Errorf("Node.firstChild did not return the first child after inserting a new element")
  }
}

func TestNodeLastChildAfterAppend(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child0/><child1/></parent>`)
  r := d.DocumentElement()
  children := r.ChildNodes()
  child1 := children.Item(1)
  if r.LastChild() != child1 {
    t.Errorf("Node.lasstChild did not return the last child")
  }
  
  child2 := d.CreateElement("child2").(dom.Node)
  r.AppendChild(child2)
  
  if r.LastChild() != child2 {
    t.Errorf("Node.lastChild did not return the last child after appending a new element")
  }
}

func TestNodeFirstChildAfterRemove(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child1/><child2/></parent>`)
  r := d.DocumentElement()
  children := r.ChildNodes()
  child1 := children.Item(0)
  child2 := children.Item(1)
  
  if r.FirstChild() != child1 {
    t.Errorf("Node.firstChild did not return the first child")
  }
  
  r.RemoveChild(r.FirstChild())
  
  if r.FirstChild() != child2 {
    t.Errorf("Node.firstChild did not return the first child after removing an element")
  }
}

func TestNodeLastChild(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child0/><child1/><child2/></parent>`)
  r := d.DocumentElement()
  children := r.ChildNodes()
  child2 := children.Item(2)
  if child2.LastChild() != nil {
    t.Errorf("Node.lastChild did not return null on an empty node")
  } else if r.LastChild() != child2 {
    t.Errorf("Node.lastChild did not return the last child")
  }
}

func TestNodePreviousSibling(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child0/><child1/><child2/></parent>`);
  r := d.DocumentElement()
  children := r.ChildNodes()
  child0 := children.Item(0)
  child1 := children.Item(1)
  child2 := children.Item(2)
  if child0.PreviousSibling() != nil {
    t.Errorf("Node.previousSibling did not return null on the first child")
  } else if child1.PreviousSibling() != child0 {
    t.Errorf("Node.previousSibling did not return the previous sibling")
  } else if child2.PreviousSibling().PreviousSibling() != child0 {
    t.Errorf("child2.previousSibling.previousSibling did not return child0")
  }
}

func TestNodeNextSibling(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child0/><child1/><child2/></parent>`)
  r := d.DocumentElement()
  children := r.ChildNodes()
  child0 := children.Item(0)
  child1 := children.Item(1)
  child2 := children.Item(2)
  if child2.NextSibling() != nil {
    t.Errorf("Node.nextSibling did not return null on the last child")
  } else if child1.NextSibling() != child2 {
    t.Errorf("Node.nextSibling did not return the next sibling")
  } else if child0.NextSibling().NextSibling() != child2 {
    t.Errorf("child0.nextSibling.nextSibling did not return child2")
  }
}

func TestNodeNextPrevSibling(t *testing.T) {
  d, _ := dom.ParseString(`<parent><child0/><child1/></parent>`)
  r := d.DocumentElement()
  children := r.ChildNodes()
  child0 := children.Item(0)
  child1 := children.Item(1)
  if child0.NextSibling().PreviousSibling() != child0 {
    t.Errorf("Node.nextSibling.previousSibling did not return itself")
  } else if child1.PreviousSibling().NextSibling() != child1 {
    t.Errorf("Node.previousSibling.nextSibling did not return itself")
  }
}

func TestElementRemoveAttribute(t *testing.T) {
  d, _ := dom.ParseString(`<parent attr="val"/>`)
  r := d.DocumentElement()
  r.RemoveAttribute("attr")
  if r.GetAttribute("attr") != "" {
    t.Errorf("Element.RemoveAttribute() did not remove the attribute, GetAttribute() returns '%s'", r.GetAttribute("attr"))
  }
}

func TestElementHasAttribute(t *testing.T) {
  d, _ := dom.ParseString(`<parent attr="val"/>`)
  r := d.DocumentElement()
  yes := r.HasAttribute("attr")
  r.RemoveAttribute("attr")
  no := r.HasAttribute("attr")
  if yes != true {
    t.Errorf("Element.HasAttribute() returned false when an attribute was present")
  } else if no != false {
    t.Errorf("Element.HasAttribute() returned true after removing an attribute")
  }
}
