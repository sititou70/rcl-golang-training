package lengthconv

import "fmt"

type Metre float64
type Feet float64

func (m Metre) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }

func MToF(m Metre) Feet { return Feet(m / 0.3048) }

func FToM(f Feet) Metre { return Metre(f * 0.3048) }
