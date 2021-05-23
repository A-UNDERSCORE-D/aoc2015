package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2015/util"
)

func main() {
	input := util.ReadLines("input.txt")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input, res)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

type op struct {
	op     string
	left   string
	right  string
	target string
	unary  bool
}

func (o *op) String() string {
	return fmt.Sprintf("%s: %q %q %q %t", o.op, o.left, o.right, o.target, o.unary)
}

func decodeOp(opLine string) *op {
	split := strings.Split(opLine, " ")
	if split[0] == "NOT" {
		// Special case for unary NOT
		return &op{
			op:     split[0],
			left:   split[1],
			target: split[3],
			unary:  true,
		}
	}
	if len(split) == 3 {
		// Unary assign
		return &op{
			op:     "SET",
			left:   split[0],
			target: split[2],
		}
	}
	// Not Unary, therefore the op is in pos 1
	return &op{
		op:     split[1],
		left:   split[0],
		right:  split[2],
		target: split[4],
	}
}

func decodeOps(ops []string) []*op {
	out := make([]*op, len(ops))
	for i, v := range ops {
		out[i] = decodeOp(v)
	}

	return out
}

type valuer interface {
	value() uint16
}

type source uint16

func (s source) value() uint16 { return uint16(s) }

type wire struct {
	name   string
	source valuer
	_value *uint16
}

func (w *wire) value() uint16 {
	if w._value != nil {
		return *w._value
	}
	v := w.source.value()
	w._value = &v
	return v
}

func isNumber(s string) (bool, uint16) {
	res, err := strconv.ParseUint(s, 10, 16)
	if err == nil {
		return true, uint16(res)
	}

	return false, 0
}

func resolve(ops []*op, name string, wires map[string]*wire) valuer {
	if res, ok := wires[name]; ok {
		return res
	}

	if ok, res := isNumber(name); ok {
		return source(res)
	}

	out := &wire{name: name}
	var sourceOp *op
	for _, op := range ops {
		if op.target == name {
			sourceOp = op
			break
		}
	}

	if sourceOp == nil {
		panic("Nil sourceop")
	}

	qResolve := func(name string) valuer {
		return resolve(ops, name, wires)
	}

	left := sourceOp.left
	right := sourceOp.right
	// target := sourceOp.target
	op := sourceOp.op

	switch op {
	case "SET":
		out.source = qResolve(left)

	case "NOT":
		if ok, res := isNumber(sourceOp.left); ok {
			out.source = ^source(res)
		} else {
			out.source = NewGate(sourceOp.op, qResolve(sourceOp.left), nil)
		}

	case "AND", "OR", "RSHIFT", "LSHIFT":
		out.source = NewGate(op, qResolve(left), qResolve(right))
	}
	wires[out.name] = out
	return out
}

type gate struct {
	typ   string
	left  valuer
	right valuer
}

func NewGate(typ string, left, right valuer) *gate {
	return &gate{
		typ:   typ,
		left:  left,
		right: right,
	}
}

func (g *gate) value() uint16 {
	switch g.typ {
	case "NOT":
		return ^g.left.value()
	case "AND":
		return g.left.value() & g.right.value()
	case "OR":
		return g.left.value() | g.right.value()
	case "LSHIFT":
		return g.left.value() << g.right.value()
	case "RSHIFT":
		return g.left.value() >> g.right.value()
	default:
		panic(g.typ)
	}
}

func part1(input []string) string {
	wires := map[string]*wire{}
	ops := decodeOps(input)

	a := resolve(ops, "a", wires)

	return fmt.Sprint(a.value())
}

func part2(input []string, origValue string) string {
	wires := map[string]*wire{}
	ops := decodeOps(input)
	for i, v := range ops {
		if v.target == "b" {
			ops[i] = &op{left: origValue, op: "SET", target: "b"}
		}
	}
	return fmt.Sprint(resolve(ops, "a", wires).value())
}
