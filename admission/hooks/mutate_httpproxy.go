package hooks

import (
	"context"
	"encoding/json"
	"net/http"

	contourv1 "github.com/projectcontour/contour/apis/projectcontour/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

const (
	annotationKubernetesIngressClass = "kubernetes.io/ingress.class"
	annotationContourIngressClass    = "projectcontour.io/ingress.class"
)

// +kubebuilder:webhook:verbs=create,path=/mutate-projectcontour-io-httpproxy,mutating=true,failurePolicy=fail,groups=projectcontour.io,resources=httpproxies,versions=v1,name=mhttpproxy.kb.io

type contourHTTPProxyMutator struct {
	client       client.Client
	decoder      *admission.Decoder
	defaultClass string
}

// NewContourHTTPProxyMutator creates a webhook handler for Contour HTTPProxy.
func NewContourHTTPProxyMutator(c client.Client, dec *admission.Decoder, defaultClass string) http.Handler {
	return &webhook.Admission{Handler: &contourHTTPProxyMutator{c, dec, defaultClass}}
}

func (m *contourHTTPProxyMutator) Handle(ctx context.Context, req admission.Request) admission.Response {
	hp := &contourv1.HTTPProxy{}
	err := m.decoder.Decode(req, hp)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	if _, ok := hp.Annotations[annotationKubernetesIngressClass]; ok {
		return admission.Allowed("ok")
	}
	if _, ok := hp.Annotations[annotationContourIngressClass]; ok {
		return admission.Allowed("ok")
	}

	hpPatched := hp.DeepCopy()
	if hpPatched.Annotations == nil {
		hpPatched.Annotations = make(map[string]string)
	}
	hpPatched.Annotations[annotationKubernetesIngressClass] = m.defaultClass

	marshaled, err := json.Marshal(hpPatched)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.Object.Raw, marshaled)
}
