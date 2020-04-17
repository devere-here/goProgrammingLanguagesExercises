package weightconv

import "fmt"

// Exported weight values
type Pound float64
type Kilogram float64

const ratio float64 = 2.2046

func (p Pound) String() string    { return fmt.Sprintf("%g lbs", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }

// LbToKg converts pounds to kilograms.
func LbToKg(p Pound) Kilogram { return Kilogram(float64(p) / ratio) }

// KgToLb converts kilograms to pounds.
func KgToLb(k Kilogram) Pound { return Pound(float64(k) * ratio) }
