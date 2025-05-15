package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type bigIntFixedDecElectron struct {
	lock *sync.Mutex
}
