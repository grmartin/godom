package dom_test

/*
 * Part of the xml/dom Go package
 *
 * Tests some of the constants used to identify node types.
 *
 * Copyright (c) 2011, Robert Johnstone
 */ 

import (
  "testing"
  "xml/dom"
)

func TestConst(t *testing.T) {
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

