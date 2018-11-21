package v1alpha3

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Simple function to strip off the first '{' and last '}'
func inlineJSON(val []byte) string {
	length := len(val)
	return string(val[1 : length-1])
}

func MarshalJSON(typeMeta metav1.TypeMeta, objectMeta metav1.ObjectMeta, spec proto.Message) ([]byte, error) {
	buffer := bytes.NewBufferString("{")

	jsonValue, err := json.Marshal(typeMeta)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(fmt.Sprintf(inlineJSON(jsonValue)))
	buffer.WriteString(",")

	jsonValue, err = json.Marshal(objectMeta)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(fmt.Sprintf("\"%s\":%s", "metadata", string(jsonValue)))
	buffer.WriteString(",")

	encoder := jsonpb.Marshaler{OrigName: false}
	jsonBytes, err := encoder.MarshalToString(spec)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(fmt.Sprintf("\"%s\":%s", "spec", string(jsonBytes)))

	buffer.WriteString("}")

	return buffer.Bytes(), nil
}
