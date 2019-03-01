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

func Test_DestinationRule(t *testing.T) {
	buffer := bytes.NewBufferString(`{
		"apiVersion":"networking.istio.io/v1alpha3",
		"kind":"DestinationRule",
		"metadata":{
			"name":"test-destination-rule",
			"namespace":"istio-system"
		},
		"spec":{
			"host":"test-host",
			"traffic_policy":{
				"load_balancer":{
					"simple":"RANDOM"
				}
			}
		}
	}`)

	destinationRule := DestinationRule{}
	err := json.Unmarshal(buffer.Bytes(), &destinationRule)
	assert.Equal(t, nil, err, "Could not unmarshal message")
	vss := destinationRule.GetSpecMessage().(*istiov1alpha3.DestinationRule)
	log.WithFields(log.Fields{
		"obj":  fmt.Sprintf("%+v", destinationRule),
		"spec": vss.String(),
	}).Info("Unmarshalled message")

	assert.Equal(t, "networking.istio.io/v1alpha3", destinationRule.TypeMeta.APIVersion)
	assert.Equal(t, "DestinationRule", destinationRule.TypeMeta.Kind)
	assert.Equal(t, "test-destination-rule", destinationRule.GetObjectMeta().GetName())
	assert.Equal(t, "test-host", destinationRule.Spec.Host)
	assert.Equal(t, "RANDOM", destinationRule.Spec.TrafficPolicy.LoadBalancer.GetSimple().String())
}
