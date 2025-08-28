package proofs

import . "github.com/mattfenwick/algorithms/pkg/logic/firstorder"

var (
	P = Prop("P")
	Q = Prop("Q")
	R = Prop("R")
	S = Prop("S")
)

type ProofsSection struct {
	Name   string
	Proofs []*Proof
}

func NewProofsSection(name string, proofs ...*Proof) *ProofsSection {
	return &ProofsSection{Name: name, Proofs: proofs}
}

var proofs = []*ProofsSection{
	NewProofsSection("theorems",
		NewRootProof("P v ~ P",
			NewProofContradiction(
				Not(Or(P, Not(P))),
				NewProofContradiction(
					P,
					IOr(P, Not(P), true),
					&Reiterate{Term: Not(Or(P, Not(P)))},
				),
				IOr(P, Not(P), false),
			),
		),
	),
}
