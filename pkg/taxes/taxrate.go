package taxes

import "fmt"

var (
	Rate_0Percent = &TaxRate{Rate: 0, Divisor: 100}
	Rate15Percent = &TaxRate{Rate: 15, Divisor: 100}
	Rate20Percent = &TaxRate{Rate: 20, Divisor: 100}

	Rate10Percent = &TaxRate{Rate: 10, Divisor: 100}
	Rate12Percent = &TaxRate{Rate: 12, Divisor: 100}
	Rate22Percent = &TaxRate{Rate: 22, Divisor: 100}
	Rate24Percent = &TaxRate{Rate: 24, Divisor: 100}
	Rate32Percent = &TaxRate{Rate: 32, Divisor: 100}
	Rate35Percent = &TaxRate{Rate: 35, Divisor: 100}
	Rate37Percent = &TaxRate{Rate: 37, Divisor: 100}
)

type TaxRate struct {
	Rate    int64
	Divisor int64
}

func (t *TaxRate) DebugString() string {
	return fmt.Sprintf("%d / %d", t.Rate, t.Divisor)
}

func (t *TaxRate) ToDebugFloat() float32 {
	return float32(t.Rate) / float32(t.Divisor)
}

func (t *TaxRate) ToDebugPercentage() float32 {
	return float32(t.Rate) / float32(t.Divisor) * 100.0
}

func (t *TaxRate) GetTax(taxableAmount int64) int64 {
	return t.Rate * taxableAmount / t.Divisor
}
