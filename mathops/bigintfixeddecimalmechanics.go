package mathops

import "sync"

type bigIntFixedDecMechanics struct {
  lock *sync.Mutex
}
