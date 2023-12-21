package utils

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hexops/valast"
)

func PrettyJson(data interface{}, debug bool) string {
	empty := ""
	tab := "\t"
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	if debug {
		print(valast.String(data))
	}

	print(buffer.String())
	return buffer.String()
}
