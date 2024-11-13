package path_update

import (
	"github.com/stretchr/testify/require"
	go_serializer "gitlab.com/pietroski-software-company/devex/golang/serializer/pkg/tools/serializer"
	"testing"
)

func TestPathUpdate(t *testing.T) {

	type (
		SubStructExample struct {
			FieldStr  string `json:"field_str,omitempty"`
			FieldStr2 string `json:"field_str_2,omitempty"`
		}

		StructExample struct {
			FieldStr    string           `json:"field_str,omitempty"`
			FieldInt    int              `json:"field_int,omitempty"`
			FieldStruct SubStructExample `json:"field_struct"`
		}
	)

	t.Run(
		"path deserialiser",
		func(t *testing.T) {
			// "field_struct/field_str_2": "any new string"

			serializer := go_serializer.NewJsonSerializer()

			payload := StructExample{
				FieldStr: "value FieldStr",
				FieldInt: 99,
				FieldStruct: SubStructExample{
					FieldStr:  "another FieldStr value",
					FieldStr2: "another FieldStr2 value",
				},
			}
			bs, err := serializer.Serialize(payload)
			require.NoError(t, err)

			var DeserializerType map[string]interface{}
			err = serializer.Deserialize(bs, &DeserializerType)
			require.NoError(t, err)

			t.Logf("DeserializerType - %v", DeserializerType)

			value1, ok := DeserializerType["field_struct"]
			require.True(t, ok)

			value2, ok := value1.(map[string]interface{})
			require.True(t, ok)

			t.Logf("value2 - %v", value2)

			_, ok = value2["field_str_2"]
			require.True(t, ok)

			value2["field_str_2"] = "any new string"
		},
	)
}
