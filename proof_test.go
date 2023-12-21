package proof_test

import (
	"testing"

	"github.com/walvisk/proof"
)

func TestProve_Int(t *testing.T) {
	proof.New(t).Prove(2).Equal(2)
	proof.New(t).Prove(3).Equal(2)
}
