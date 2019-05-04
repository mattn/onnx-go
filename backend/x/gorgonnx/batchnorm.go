package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
)

// SPEC: https://github.com/onnx/onnx/blob/master/docs/Operators.md#BatchNormalization
// Gorgonia implem: https://godoc.org/gorgonia.org/gorgonia#BatchNorm

type batchnorm struct {
	epsilon  float64
	momentum float64
}

func init() {
	register("BatchNormalization", &batchnorm{})
}

func (b *batchnorm) apply(g *Graph, n *Node) error {
	return &onnx.ErrNotImplemented{
		Operator: "BatchNormalization",
	}
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}
	n.gorgoniaNode = children[0].gorgoniaNode
	return err
}

func (b *batchnorm) init(o onnx.Operation) error {
	b.epsilon = 1e-5
	b.momentum = 0.9
	if e, ok := o.Attributes["epsilon"]; ok {
		if v, ok := e.(float32); ok {
			b.epsilon = float64(v)
		} else {
			return errors.New("epsilon is not a float64")
		}
	}
	if e, ok := o.Attributes["momentum"]; ok {
		if v, ok := e.(float32); ok {
			b.momentum = float64(v)
		} else {
			return errors.New("momentum is not a float64")
		}
	}
	return nil
}
