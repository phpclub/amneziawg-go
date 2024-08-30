//go:build !linux

package device

import (
	"github.com/phpclub/amneziawg-go/conn"
	"github.com/phpclub/amneziawg-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
