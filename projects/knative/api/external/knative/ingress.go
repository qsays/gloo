package knative

import (
	"reflect"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	"knative.dev/serving/pkg/apis/networking/v1alpha1"
)

type Ingress v1alpha1.Ingress

func (p *Ingress) GetMetadata() core.Metadata {
	return kubeutils.FromKubeMeta(p.ObjectMeta)
}

func (p *Ingress) SetMetadata(meta core.Metadata) {
	p.ObjectMeta = kubeutils.ToKubeMeta(meta)
}

func (p *Ingress) Equal(that interface{}) bool {
	return reflect.DeepEqual(p, that)
}

func (p *Ingress) Clone() *Ingress {
	ing := v1alpha1.Ingress(*p)
	copy := ing.DeepCopy()
	newIng := Ingress(*copy)
	return &newIng
}

func (p *Ingress) IsPublic() bool {
	ing := v1alpha1.Ingress(*p)
	return ing.IsPublic()
}
