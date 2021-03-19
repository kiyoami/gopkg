package tree

import (
	"encoding/json"
	"errors"
	"fmt"
)

// LoadJson : Json decodes and stores to object.
func (tree *Tree) LoadJson(jsonData []byte) error {
	var root interface{} // The memory for stored for json data prepares.

	if err := json.Unmarshal(jsonData, &root); err != nil {
		tree.root = nil
		return errors.New(fmt.Sprintf(`Error on parsing JSON, %s`, err))
	}
	tree.root = &root
	return nil
}
