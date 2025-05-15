package mathops

import "sync"

type bigIntNumMechanics struct {
  lock *sync.Mutex
}
