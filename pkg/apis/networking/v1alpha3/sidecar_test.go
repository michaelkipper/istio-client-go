package v1alpha3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	istiov1alpha3 "istio.io/api/networking/v1alpha3"

	log "github.com/sirupsen/logrus"
)

func Test_Sidecar(t *testing.T) {
	// This was extracted from a call in another test to Kubernetes.
	buffer := bytes.NewBufferString(`{
		"apiVersion":"networking.istio.io/v1alpha3",
		"kind":"Sidecar",
		"metadata":{
			"name":"test-sidecar",
			"namespace":"istio-system"
		},
		"spec":{
		}
	}`)

	vs := Sidecar{}
	err := json.Unmarshal(buffer.Bytes(), &vs)
	assert.Equal(t, nil, err, "Could not unmarshal message")
	vss := vs.GetSpecMessage().(*istiov1alpha3.Sidecar)
	log.WithFields(log.Fields{
		"obj":  fmt.Sprintf("%+v", vs),
		"spec": vss.String(),
	}).Info("Unmarshalled message")

	assert.Equal(t, "networking.istio.io/v1alpha3", vs.TypeMeta.APIVersion)
	assert.Equal(t, "Sidecar", vs.TypeMeta.Kind)
	assert.Equal(t, "test-sidecar", vs.GetObjectMeta().GetName())
}
