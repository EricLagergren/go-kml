// generated by stringer -type=viewRefreshEnum; DO NOT EDIT

package kml

import "fmt"

const _viewRefreshEnum_name = "NeverOnRequestOnStopOnRegion"

var _viewRefreshEnum_index = [...]uint8{0, 5, 14, 20, 28}

func (i viewRefreshEnum) String() string {
	if i < 0 || i+1 >= viewRefreshEnum(len(_viewRefreshEnum_index)) {
		return fmt.Sprintf("viewRefreshEnum(%d)", i)
	}
	return _viewRefreshEnum_name[_viewRefreshEnum_index[i]:_viewRefreshEnum_index[i+1]]
}
