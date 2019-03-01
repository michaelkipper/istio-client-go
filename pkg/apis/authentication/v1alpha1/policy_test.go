package v1alpha1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	istiov1alpha1 "istio.io/api/authentication/v1alpha1"

	log "github.com/sirupsen/logrus"
)

func Test_Policy(t *testing.T) {
	buffer := bytes.NewBufferString(`{
		"apiVersion":"networking.istio.io/v1alpha3",
		"kind":"Policy",
		"metadata":{
			"name":"test-policy",
			"namespace":"istio-system"
		},
		"spec":{
			"targets":[]
		}
	}`)

	policy := Policy{}
	err := json.Unmarshal(buffer.Bytes(), &policy)
	assert.Equal(t, nil, err, "Could not unmarshal message")
	vss := policy.GetSpecMessage().(*istiov1alpha1.Policy)
	log.WithFields(log.Fields{
		"obj":  fmt.Sprintf("%+v", policy),
		"spec": vss.String(),
	}).Info("Unmarshalled message")

	assert.Equal(t, "networking.istio.io/v1alpha3", policy.TypeMeta.APIVersion)
	assert.Equal(t, "Policy", policy.TypeMeta.Kind)
	assert.Equal(t, "test-policy", policy.GetObjectMeta().GetName())
}
