// author: asydevc <asydev@163.com>
// date: 2021-08-04

// Package application.
package app

import (
	"sync"
)

func init() {
	new(sync.Once).Do(func() {
		Config = (&configuration{}).init()
		With = (&with{}).init()
		Validate = (&validation{}).init()
	})
}
