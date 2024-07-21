package taxes

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestTaxes(t *testing.T) {
	gomega.RegisterFailHandler(Fail)

	RunBracketTests()
	RunMedicareTests()

	RunSpecs(t, "taxes suite")
}
