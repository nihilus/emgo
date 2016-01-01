package main

import (
	"go/ast"
	"strings"
	"text/template"
)

type instance struct {
	Name string
	Addr string
}

type multiCtx struct {
	Pkg       string
	Periph    string
	Instances []instance
	Regs      []*reg
	Import    []string
}

func instances(f string, lines []string) ([]instance, []string) {
	var insts []instance
loop:
	for len(lines) > 0 {
		line := strings.TrimSpace(lines[0])
		switch line {
		case "Registers:", "Import:":
			break loop
		}
		lines = lines[1:]
		if line == "" {
			continue
		}
		name, line := split(line)
		addr, _ := split(line)
		insts = append(insts, instance{Name: name, Addr: addr})
	}
	return insts, lines
}

func uses(lines []string) ([]string, []string) {
	var imports []string
loop:
	for len(lines) > 0 {
		line := strings.TrimSpace(lines[0])
		switch line {
		case "Instances:", "Registers:":
			break loop
		}
		lines = lines[1:]
		if line == "" {
			continue
		}
		imports = append(imports, line)
	}
	return imports, lines
}

func multi(pkg, f, txt string, decls []ast.Decl) {
	lines := strings.Split(txt, "\n")
	periph, _ := split(lines[0][len("Peripheral:"):])
	lines = lines[1:]
	ctx := &multiCtx{
		Pkg:    pkg,
		Periph: periph,
	}
	for len(lines) > 0 {
		line := strings.TrimSpace(lines[0])
		lines = lines[1:]
		switch line {
		case "Instances:":
			var insts []instance
			insts, lines = instances(f, lines)
			ctx.Instances = append(ctx.Instances, insts...)
		case "Registers:":
			var regs []*reg
			regs, lines = registers(f, lines, decls)
			ctx.Regs = append(ctx.Regs, regs...)
		case "Import:":
			var imports []string
			imports, lines = uses(lines)
			ctx.Import = append(ctx.Import, imports...)
		}
	}
	save(f, multiTmpl, ctx)
}

const multiText = `package {{.Pkg}}

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"
	
	{{range .Import}}"{{.}}"
	{{end}}
)
{{$p := .Periph}}
type {{$p}} struct {
	{{range .Regs}}{{if .Name}} {{.Name}} {{if .Len}}[{{.Len}}]{{end}}{{.Name}} {{else}} _ {{if .Len}}[{{.Len}}]{{end}}uint{{.Size}} {{end}}
{{end}}}

func (p *{{$p}}) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

{{range .Instances}}
var {{.Name}} = (*{{$p}})(unsafe.Pointer(uintptr({{.Addr}})))
{{end}}

{{range .Regs}}{{if .Name}}

{{$len  := .Len}}
{{$uint := print "uint" .Size}}
{{$u    := print "U" .Size}}
{{$mu   := print "mmio." $u}}
{{$ru   := print "r." $u}}
{{$um   := print "UM" .Size }}
{{$mum  := print "mmio." $um}}
{{$rmum := print "rm." $um}}
{{$po   := print "unsafe.Pointer(uintptr(unsafe.Pointer(p))+" .Offset ")"}}
{{$bits := print .Name "_Bits" }}
{{$rm   := print .Name "_Mask" }}

type {{$bits}} {{$uint}}

func (b {{$bits}}) Field(mask {{$bits}}) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask {{$bits}}) J(v int) {{$bits}} {
	return {{$bits}}(bits.Make32(v, uint32(mask)))
}

type {{.Name}} struct { {{$mu}} }

func (r *{{.Name}}) Bits(mask {{$bits}}) {{$bits}} {return {{$bits}}({{$ru}}.Bits({{$uint}}(mask))) }
func (r *{{.Name}}) StoreBits(mask, b {{$bits}})   { {{$ru}}.StoreBits({{$uint}}(mask), {{$uint}}(b)) }
func (r *{{.Name}}) SetBits(mask {{$bits}})        { {{$ru}}.SetBits({{$uint}}(mask)) }
func (r *{{.Name}}) ClearBits(mask {{$bits}})      { {{$ru}}.ClearBits({{$uint}}(mask)) }
func (r *{{.Name}}) Load() {{$bits}}               { return {{$bits}}({{$ru}}.Load()) }
func (r *{{.Name}}) Store(b {{$bits}})             { {{$ru}}.Store({{$uint}}(b)) }

type {{$rm}} struct { {{$mum}} }

func (rm {{$rm}}) Load() {{$bits}}   { return {{$bits}}({{$rmum}}.Load()) }
func (rm {{$rm}}) Store(b {{$bits}}) { {{$rmum}}.Store({{$uint}}(b)) }

{{range .Bits}}
{{if $len}}
func (p *{{$p}}) {{.}}(n int) {{$rm}} {
	return {{print $rm "{" $mum "{&(*[" $len "]" $mu ")(" $po ")[n]," $uint "(" . ")}}"}}
}
{{else}}
func (p *{{$p}}) {{.}}() {{$rm}} {
	return {{print $rm "{" $mum "{(*" $mu ")(" $po ")," $uint "(" . ")}}"}}
}
{{end}}
{{end}}

{{end}}{{end}}
`

var multiTmpl = template.Must(template.New("multi").Parse(multiText))
