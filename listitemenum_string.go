// generated by stringer -type=listItemEnum; DO NOT EDIT

package kml

import "fmt"

const _listItemEnum_name = "CheckRadioFolderCheckOffOnlyCheckHideChildren"

var _listItemEnum_index = [...]uint8{0, 5, 16, 28, 45}

func (i listItemEnum) String() string {
	if i < 0 || i+1 >= listItemEnum(len(_listItemEnum_index)) {
		return fmt.Sprintf("listItemEnum(%d)", i)
	}
	return _listItemEnum_name[_listItemEnum_index[i]:_listItemEnum_index[i+1]]
}
