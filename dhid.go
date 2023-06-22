package dhid

import "strings"

// DHID is the Deutsche Haltestellen-ID. Defined by VDV, based on IFOPT.
type DHID string

// New returns a new DHID by given string
func New(s string) DHID {
	return DHID(s)
}

func (dhid DHID) String() string {
	return string(dhid)
}

func (dhid DHID) ToSlug() string {
	return strings.ReplaceAll(string(dhid), ":", "-")
}

func (dhid DHID) IsValid() bool {
	return len(dhid) > 0
}

// GetStopPart returns the part of the stop itself without "Bereich" (area) and "Steig" (platform)
func (dhid DHID) GetStopPart() DHID {
	var partsArray [3]string
	parts := strings.Split(dhid.String(), ":")
	// Use an array to get only the first (minimum) three parts of the slice
	copy(partsArray[:], parts)
	// Join the three parts in one IFOPT string again
	ret := strings.Join(partsArray[:], ":")
	//fmt.Println(partsArray)
	return DHID(ret)
}
