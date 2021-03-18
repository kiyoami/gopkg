package tree

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Get : This tool can easily access to the node on nested data.
func (tree *Tree) Get(path string) (interface{}, error) {
	current := tree.root

	for level, key := range strings.Split(path, ".") {
		switch data := (*current).(type) {
		case map[string]interface{}:
			// if map type, an element select by string key
			if val, ok := data[key]; ok {
				current = &val
			} else {
				return nil, errors.New(fmt.Sprintf(`Requested key cannot find in map. "%s(%s)"`, path, key))
			}
		case []interface{}:
			// if slice type, an element select by numeric index
			i, err := strconv.Atoi(key)
			if err != nil {
				return nil, errors.New(fmt.Sprintf(`Index of slice cannot convert to integer. "%s(%s)"`, path, key))
			}
			if l := len(data); i < l {
				val := data[i]
				current = &val
			} else {
				return nil, errors.New(fmt.Sprintf(`Index of slice is invalid. Available index [0..%d], but referred "%s(%d)".`, l-1, path, i))
			}
		default:
			return nil, errors.New(fmt.Sprintf(`node type error %d "%s"`, level, key))
		}
	}

	return *current, nil
}
