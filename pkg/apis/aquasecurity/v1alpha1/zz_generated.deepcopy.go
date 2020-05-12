// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Artifact) DeepCopyInto(out *Artifact) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Artifact.
func (in *Artifact) DeepCopy() *Artifact {
	if in == nil {
		return nil
	}
	out := new(Artifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubernetesBenchmark) DeepCopyInto(out *CISKubernetesBenchmark) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubernetesBenchmark.
func (in *CISKubernetesBenchmark) DeepCopy() *CISKubernetesBenchmark {
	if in == nil {
		return nil
	}
	out := new(CISKubernetesBenchmark)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CISKubernetesBenchmark) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubernetesBenchmarkList) DeepCopyInto(out *CISKubernetesBenchmarkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CISKubernetesBenchmark, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubernetesBenchmarkList.
func (in *CISKubernetesBenchmarkList) DeepCopy() *CISKubernetesBenchmarkList {
	if in == nil {
		return nil
	}
	out := new(CISKubernetesBenchmarkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CISKubernetesBenchmarkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubernetesBenchmarkReport) DeepCopyInto(out *CISKubernetesBenchmarkReport) {
	*out = *in
	in.GeneratedAt.DeepCopyInto(&out.GeneratedAt)
	out.Scanner = in.Scanner
	if in.Sections != nil {
		in, out := &in.Sections, &out.Sections
		*out = make([]CISKubernetesBenchmarkSection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubernetesBenchmarkReport.
func (in *CISKubernetesBenchmarkReport) DeepCopy() *CISKubernetesBenchmarkReport {
	if in == nil {
		return nil
	}
	out := new(CISKubernetesBenchmarkReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubernetesBenchmarkResult) DeepCopyInto(out *CISKubernetesBenchmarkResult) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubernetesBenchmarkResult.
func (in *CISKubernetesBenchmarkResult) DeepCopy() *CISKubernetesBenchmarkResult {
	if in == nil {
		return nil
	}
	out := new(CISKubernetesBenchmarkResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubernetesBenchmarkSection) DeepCopyInto(out *CISKubernetesBenchmarkSection) {
	*out = *in
	if in.Tests != nil {
		in, out := &in.Tests, &out.Tests
		*out = make([]CISKubernetesBenchmarkTests, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubernetesBenchmarkSection.
func (in *CISKubernetesBenchmarkSection) DeepCopy() *CISKubernetesBenchmarkSection {
	if in == nil {
		return nil
	}
	out := new(CISKubernetesBenchmarkSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubernetesBenchmarkTests) DeepCopyInto(out *CISKubernetesBenchmarkTests) {
	*out = *in
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]CISKubernetesBenchmarkResult, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubernetesBenchmarkTests.
func (in *CISKubernetesBenchmarkTests) DeepCopy() *CISKubernetesBenchmarkTests {
	if in == nil {
		return nil
	}
	out := new(CISKubernetesBenchmarkTests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Check) DeepCopyInto(out *Check) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Check.
func (in *Check) DeepCopy() *Check {
	if in == nil {
		return nil
	}
	out := new(Check)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigAudit) DeepCopyInto(out *ConfigAudit) {
	*out = *in
	in.GeneratedAt.DeepCopyInto(&out.GeneratedAt)
	out.Scanner = in.Scanner
	out.Resource = in.Resource
	if in.PodChecks != nil {
		in, out := &in.PodChecks, &out.PodChecks
		*out = make([]Check, len(*in))
		copy(*out, *in)
	}
	if in.ContainerChecks != nil {
		in, out := &in.ContainerChecks, &out.ContainerChecks
		*out = make(map[string][]Check, len(*in))
		for key, val := range *in {
			var outVal []Check
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]Check, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigAudit.
func (in *ConfigAudit) DeepCopy() *ConfigAudit {
	if in == nil {
		return nil
	}
	out := new(ConfigAudit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigAuditReport) DeepCopyInto(out *ConfigAuditReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigAuditReport.
func (in *ConfigAuditReport) DeepCopy() *ConfigAuditReport {
	if in == nil {
		return nil
	}
	out := new(ConfigAuditReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigAuditReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigAuditReportList) DeepCopyInto(out *ConfigAuditReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigAuditReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigAuditReportList.
func (in *ConfigAuditReportList) DeepCopy() *ConfigAuditReportList {
	if in == nil {
		return nil
	}
	out := new(ConfigAuditReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigAuditReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeHunterOutput) DeepCopyInto(out *KubeHunterOutput) {
	*out = *in
	in.GeneratedAt.DeepCopyInto(&out.GeneratedAt)
	out.Scanner = in.Scanner
	if in.Vulnerabilities != nil {
		in, out := &in.Vulnerabilities, &out.Vulnerabilities
		*out = make([]KubeHunterVulnerability, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeHunterOutput.
func (in *KubeHunterOutput) DeepCopy() *KubeHunterOutput {
	if in == nil {
		return nil
	}
	out := new(KubeHunterOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeHunterReport) DeepCopyInto(out *KubeHunterReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeHunterReport.
func (in *KubeHunterReport) DeepCopy() *KubeHunterReport {
	if in == nil {
		return nil
	}
	out := new(KubeHunterReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeHunterReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeHunterReportList) DeepCopyInto(out *KubeHunterReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeHunterReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeHunterReportList.
func (in *KubeHunterReportList) DeepCopy() *KubeHunterReportList {
	if in == nil {
		return nil
	}
	out := new(KubeHunterReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeHunterReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeHunterVulnerability) DeepCopyInto(out *KubeHunterVulnerability) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeHunterVulnerability.
func (in *KubeHunterVulnerability) DeepCopy() *KubeHunterVulnerability {
	if in == nil {
		return nil
	}
	out := new(KubeHunterVulnerability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesNamespacedResource) DeepCopyInto(out *KubernetesNamespacedResource) {
	*out = *in
	out.KubernetesResource = in.KubernetesResource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesNamespacedResource.
func (in *KubernetesNamespacedResource) DeepCopy() *KubernetesNamespacedResource {
	if in == nil {
		return nil
	}
	out := new(KubernetesNamespacedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesResource) DeepCopyInto(out *KubernetesResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesResource.
func (in *KubernetesResource) DeepCopy() *KubernetesResource {
	if in == nil {
		return nil
	}
	out := new(KubernetesResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scanner) DeepCopyInto(out *Scanner) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scanner.
func (in *Scanner) DeepCopy() *Scanner {
	if in == nil {
		return nil
	}
	out := new(Scanner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vulnerability) DeepCopyInto(out *Vulnerability) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vulnerability.
func (in *Vulnerability) DeepCopy() *Vulnerability {
	if in == nil {
		return nil
	}
	out := new(Vulnerability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vulnerability) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityItem) DeepCopyInto(out *VulnerabilityItem) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityItem.
func (in *VulnerabilityItem) DeepCopy() *VulnerabilityItem {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityList) DeepCopyInto(out *VulnerabilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vulnerability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityList.
func (in *VulnerabilityList) DeepCopy() *VulnerabilityList {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VulnerabilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityReport) DeepCopyInto(out *VulnerabilityReport) {
	*out = *in
	in.GeneratedAt.DeepCopyInto(&out.GeneratedAt)
	out.Scanner = in.Scanner
	out.Registry = in.Registry
	out.Artifact = in.Artifact
	out.Summary = in.Summary
	if in.Vulnerabilities != nil {
		in, out := &in.Vulnerabilities, &out.Vulnerabilities
		*out = make([]VulnerabilityItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityReport.
func (in *VulnerabilityReport) DeepCopy() *VulnerabilityReport {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilitySummary) DeepCopyInto(out *VulnerabilitySummary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilitySummary.
func (in *VulnerabilitySummary) DeepCopy() *VulnerabilitySummary {
	if in == nil {
		return nil
	}
	out := new(VulnerabilitySummary)
	in.DeepCopyInto(out)
	return out
}