package gohaml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoctypeXmlSpecWithXhtmlFormat(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	nodes := []Node{&DoctypeNode{"XML"}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<?xml version='1.0' encoding='utf-8' ?>")
}

func TestDoctypeEmptySpecWithXhtmlFormat(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	nodes := []Node{&DoctypeNode{""}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">")
}

func TestDoctype1Point1SpecWithXhtmlFormat(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	nodes := []Node{&DoctypeNode{"1.1"}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\" \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\">")
}

func TestDoctypeMobileSpecWithXhtmlFormat(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	nodes := []Node{&DoctypeNode{"mobile"}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<!DOCTYPE html PUBLIC \"-//WAPFORUM//DTD XHTML Mobile 1.2//EN\" \"http://www.openmobilealliance.org/tech/DTD/xhtml-mobile12.dtd\">")
}

func TestDoctypeBasicSpecWithXhtmlFormat(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	nodes := []Node{&DoctypeNode{"basic"}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML Basic 1.1//EN\" \"http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd\">")
}

func TestDoctypeFramesetSpecWithXhtmlFormat(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	nodes := []Node{&DoctypeNode{"frameset"}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Frameset//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd\">")
}

func TestDoctype5SpecWithXhtmlFormat(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "xhtml"
	nodes := []Node{&DoctypeNode{"5"}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<!DOCTYPE html>")
}

func TestDoctypeXmlSpecWithHtml5Format(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html5"
	nodes := []Node{&DoctypeNode{"XML"}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "")
}

func TestDoctypeEmptySpecWithHtml5Format(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html5"
	nodes := []Node{&DoctypeNode{""}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<!DOCTYPE html>")
}

func TestDoctypeXmlSpecWithHtml4Format(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html4"
	nodes := []Node{&DoctypeNode{"XML"}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "")
}

func TestDoctypeEmptySpecWithHtml4Format(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html4"
	nodes := []Node{&DoctypeNode{""}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.01 Transitional//EN\" \"http://www.w3.org/TR/html4/loose.dtd\">")
}

func TestDoctypeFramesetSpecWithHtml4Format(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html4"
	nodes := []Node{&DoctypeNode{"frameset"}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.01 Frameset//EN\" \"http://www.w3.org/TR/html4/frameset.dtd\">")
}

func TestDoctypeStrictSpecWithHtml4Format(t *testing.T) {
	opts := CompilerOpts{}
	opts.Format = "html4"
	nodes := []Node{&DoctypeNode{"strict"}}
	pdoc := ParsedDoc{}
	pdoc.Nodes = nodes
	compiler := DefaultCompiler{}

	cdoc, e := compiler.Compile(pdoc, opts)

	assert.Nil(t, e)
	assert.NotNil(t, cdoc)
	assert.Equal(t, len(cdoc.Outputs), 1)

	output := cdoc.Outputs[0].(*StaticOutput)
	assert.Equal(t, output.Content, "<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.01//EN\" \"http://www.w3.org/TR/html4/strict.dtd\">")
}
