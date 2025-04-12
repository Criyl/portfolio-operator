package v1

import "k8s.io/apimachinery/pkg/runtime"

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *Portfolio) DeepCopyInto(out *Portfolio) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = PortfolioSpec{
		Name:        in.Spec.Name,
		Url:         in.Spec.Url,
		Blog:        in.Spec.Blog,
		Icon:        in.Spec.Icon,
		Healthcheck: in.Spec.Healthcheck,
		Tags:        in.Spec.Tags,
	}
}

// DeepCopyObject returns a generically typed copy of an object
func (in *Portfolio) DeepCopyObject() runtime.Object {
	out := Portfolio{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject returns a generically typed copy of an object
func (in *PortfolioList) DeepCopyObject() runtime.Object {
	out := PortfolioList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Portfolio, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
