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
	buffer := bytes.NewBufferString(`{
		"apiVersion":"networking.istio.io/v1alpha3",
		"kind":"Sidecar",
		"metadata":{
			"name":"test-sidecar",
			"namespace":"istio-system"
		},
		"spec":{
			"workload_selector":{
				"labels":{
					"foo":"bar"
				}
			},
			"ingress":[
				{
					"port":{
						"number": 123
					}
				}
			],
			"egress":[
				{
					"port":{
						"number": 456
					}
				}
			]
		}
	}`)

	sidecar := Sidecar{}
	err := json.Unmarshal(buffer.Bytes(), &sidecar)
	assert.Equal(t, nil, err, "Could not unmarshal message")
	vss := sidecar.GetSpecMessage().(*istiov1alpha3.Sidecar)
	log.WithFields(log.Fields{
		"obj":  fmt.Sprintf("%+v", sidecar),
		"spec": vss.String(),
	}).Info("Unmarshalled message")

	assert.Equal(t, "networking.istio.io/v1alpha3", sidecar.TypeMeta.APIVersion)
	assert.Equal(t, "Sidecar", sidecar.TypeMeta.Kind)
	assert.Equal(t, "test-sidecar", sidecar.GetObjectMeta().GetName())
	assert.Equal(t, "bar", sidecar.Spec.WorkloadSelector.Labels["foo"])
	assert.Equal(t, uint32(123), sidecar.Spec.Ingress[0].Port.Number)
	assert.Equal(t, uint32(456), sidecar.Spec.Egress[0].Port.Number)
}
