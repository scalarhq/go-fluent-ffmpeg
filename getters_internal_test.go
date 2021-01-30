package fluentffmpeg

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/fatih/structs"
)

func TestFieldsHaveGetters(t *testing.T) {
	cmd := NewCommand("")

	if err := testFieldsHaveGetters(cmd.Args.input, cmd); err != nil {
		t.Error(err.Error())
		return
	}

	if err := testFieldsHaveGetters(cmd.Args.output, cmd); err != nil {
		t.Error(err.Error())
	}
}

func testFieldsHaveGetters(argType interface{}, cmd *Command) error {

	for _, field := range structs.Names(argType) {
		sf, _ := reflect.TypeOf(argType).FieldByName(field)
		if sf.Tag.Get("getter") != "none" {
			method := reflect.ValueOf(cmd.Args).MethodByName("Get" + strings.Title(field))
			if (method == reflect.Value{}) {
				return fmt.Errorf("field: (%s) has no getter", field)
			}
		}
	}

	return nil
}
