package lengthconv

import "fmt"

// Exported length values
type Meter float64
type Foot float64

const ratio float64 = 3.28084

func (m Meter) String() string { return fmt.Sprintf("%g m", m) }
func (ft Foot) String() string { return fmt.Sprintf("%g ft", ft) }

// MToFt converts Meters to Feet.
func MToFt(m Meter) Foot { return Foot(float64(m) * ratio) }

// FtToM converts Feet to Meters.
func FtToM(ft Foot) Meter { return Meter(float64(ft) / ratio) }
