// +build !ignore_autogenerated

/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"encoding/json"
	"github.com/gardener/landscaper/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationConfig) DeepCopyInto(out *ApplicationConfig) {
	*out = *in
	in.TypeSpecificData.DeepCopyInto(&out.TypeSpecificData)
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretValues != nil {
		in, out := &in.SecretValues, &out.SecretValues
		*out = new(SecretValues)
		(*in).DeepCopyInto(*out)
	}
	if in.NamedSecretValues != nil {
		in, out := &in.NamedSecretValues, &out.NamedSecretValues
		*out = make(map[string]NamedSecretValues, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.ReadyRequirements.DeepCopyInto(&out.ReadyRequirements)
	in.InternalImportParameters.DeepCopyInto(&out.InternalImportParameters)
	if in.ImportParameters != nil {
		in, out := &in.ImportParameters, &out.ImportParameters
		*out = make([]ImportParameter, len(*in))
		copy(*out, *in)
	}
	in.ExportParameters.DeepCopyInto(&out.ExportParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationConfig.
func (in *ApplicationConfig) DeepCopy() *ApplicationConfig {
	if in == nil {
		return nil
	}
	out := new(ApplicationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationState) DeepCopyInto(out *ApplicationState) {
	*out = *in
	in.DetailedState.DeepCopyInto(&out.DetailedState)
	if in.InstallationState != nil {
		in, out := &in.InstallationState, &out.InstallationState
		*out = new(InstallationState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationState.
func (in *ApplicationState) DeepCopy() *ApplicationState {
	if in == nil {
		return nil
	}
	out := new(ApplicationState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoDelete) DeepCopyInto(out *AutoDelete) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoDelete.
func (in *AutoDelete) DeepCopy() *AutoDelete {
	if in == nil {
		return nil
	}
	out := new(AutoDelete)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBom) DeepCopyInto(out *ClusterBom) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBom.
func (in *ClusterBom) DeepCopy() *ClusterBom {
	if in == nil {
		return nil
	}
	out := new(ClusterBom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBom) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBomCondition) DeepCopyInto(out *ClusterBomCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBomCondition.
func (in *ClusterBomCondition) DeepCopy() *ClusterBomCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterBomCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBomList) DeepCopyInto(out *ClusterBomList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterBom, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBomList.
func (in *ClusterBomList) DeepCopy() *ClusterBomList {
	if in == nil {
		return nil
	}
	out := new(ClusterBomList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBomList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBomSpec) DeepCopyInto(out *ClusterBomSpec) {
	*out = *in
	if in.ApplicationConfigs != nil {
		in, out := &in.ApplicationConfigs, &out.ApplicationConfigs
		*out = make([]ApplicationConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoDelete != nil {
		in, out := &in.AutoDelete, &out.AutoDelete
		*out = new(AutoDelete)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBomSpec.
func (in *ClusterBomSpec) DeepCopy() *ClusterBomSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterBomSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBomStatus) DeepCopyInto(out *ClusterBomStatus) {
	*out = *in
	if in.ApplicationStates != nil {
		in, out := &in.ApplicationStates, &out.ApplicationStates
		*out = make([]ApplicationState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.OverallTime.DeepCopyInto(&out.OverallTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterBomCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBomStatus.
func (in *ClusterBomStatus) DeepCopy() *ClusterBomStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterBomStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBomSync) DeepCopyInto(out *ClusterBomSync) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBomSync.
func (in *ClusterBomSync) DeepCopy() *ClusterBomSync {
	if in == nil {
		return nil
	}
	out := new(ClusterBomSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBomSync) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBomSyncList) DeepCopyInto(out *ClusterBomSyncList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterBomSync, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBomSyncList.
func (in *ClusterBomSyncList) DeepCopy() *ClusterBomSyncList {
	if in == nil {
		return nil
	}
	out := new(ClusterBomSyncList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBomSyncList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBomSyncSpec) DeepCopyInto(out *ClusterBomSyncSpec) {
	*out = *in
	in.Timestamp.DeepCopyInto(&out.Timestamp)
	in.Until.DeepCopyInto(&out.Until)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBomSyncSpec.
func (in *ClusterBomSyncSpec) DeepCopy() *ClusterBomSyncSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterBomSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBomSyncStatus) DeepCopyInto(out *ClusterBomSyncStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBomSyncStatus.
func (in *ClusterBomSyncStatus) DeepCopy() *ClusterBomSyncStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterBomSyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurrentOperation) DeepCopyInto(out *CurrentOperation) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurrentOperation.
func (in *CurrentOperation) DeepCopy() *CurrentOperation {
	if in == nil {
		return nil
	}
	out := new(CurrentOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfig) DeepCopyInto(out *DeploymentConfig) {
	*out = *in
	in.TypeSpecificData.DeepCopyInto(&out.TypeSpecificData)
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.NamedInternalSecretNames != nil {
		in, out := &in.NamedInternalSecretNames, &out.NamedInternalSecretNames
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.ReconcileTime.DeepCopyInto(&out.ReconcileTime)
	in.ReadyRequirements.DeepCopyInto(&out.ReadyRequirements)
	in.InternalImportParameters.DeepCopyInto(&out.InternalImportParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfig.
func (in *DeploymentConfig) DeepCopy() *DeploymentConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetailedState) DeepCopyInto(out *DetailedState) {
	*out = *in
	if in.DeletionTimestamp != nil {
		in, out := &in.DeletionTimestamp, &out.DeletionTimestamp
		*out = (*in).DeepCopy()
	}
	in.CurrentOperation.DeepCopyInto(&out.CurrentOperation)
	in.LastOperation.DeepCopyInto(&out.LastOperation)
	if in.Reachability != nil {
		in, out := &in.Reachability, &out.Reachability
		*out = new(Reachability)
		(*in).DeepCopyInto(*out)
	}
	if in.Readiness != nil {
		in, out := &in.Readiness, &out.Readiness
		*out = new(Readiness)
		(*in).DeepCopyInto(*out)
	}
	if in.HdcConditions != nil {
		in, out := &in.HdcConditions, &out.HdcConditions
		*out = make([]HubDeploymentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TypeSpecificStatus != nil {
		in, out := &in.TypeSpecificStatus, &out.TypeSpecificStatus
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetailedState.
func (in *DetailedState) DeepCopy() *DetailedState {
	if in == nil {
		return nil
	}
	out := new(DetailedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorEntry) DeepCopyInto(out *ErrorEntry) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorEntry.
func (in *ErrorEntry) DeepCopy() *ErrorEntry {
	if in == nil {
		return nil
	}
	out := new(ErrorEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorHistory) DeepCopyInto(out *ErrorHistory) {
	*out = *in
	if in.ErrorEntries != nil {
		in, out := &in.ErrorEntries, &out.ErrorEntries
		*out = make([]ErrorEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorHistory.
func (in *ErrorHistory) DeepCopy() *ErrorHistory {
	if in == nil {
		return nil
	}
	out := new(ErrorHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportParameters) DeepCopyInto(out *ExportParameters) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]json.RawMessage, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(json.RawMessage, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportParameters.
func (in *ExportParameters) DeepCopy() *ExportParameters {
	if in == nil {
		return nil
	}
	out := new(ExportParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubDeployItemConfiguration) DeepCopyInto(out *HubDeployItemConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.DeploymentConfig.DeepCopyInto(&out.DeploymentConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubDeployItemConfiguration.
func (in *HubDeployItemConfiguration) DeepCopy() *HubDeployItemConfiguration {
	if in == nil {
		return nil
	}
	out := new(HubDeployItemConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubDeployItemProviderStatus) DeepCopyInto(out *HubDeployItemProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.LastOperation.DeepCopyInto(&out.LastOperation)
	if in.Reachability != nil {
		in, out := &in.Reachability, &out.Reachability
		*out = new(Reachability)
		(*in).DeepCopyInto(*out)
	}
	if in.Readiness != nil {
		in, out := &in.Readiness, &out.Readiness
		*out = new(Readiness)
		(*in).DeepCopyInto(*out)
	}
	if in.TypeSpecificStatus != nil {
		in, out := &in.TypeSpecificStatus, &out.TypeSpecificStatus
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubDeployItemProviderStatus.
func (in *HubDeployItemProviderStatus) DeepCopy() *HubDeployItemProviderStatus {
	if in == nil {
		return nil
	}
	out := new(HubDeployItemProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubDeploymentCondition) DeepCopyInto(out *HubDeploymentCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubDeploymentCondition.
func (in *HubDeploymentCondition) DeepCopy() *HubDeploymentCondition {
	if in == nil {
		return nil
	}
	out := new(HubDeploymentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportParameter) DeepCopyInto(out *ImportParameter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportParameter.
func (in *ImportParameter) DeepCopy() *ImportParameter {
	if in == nil {
		return nil
	}
	out := new(ImportParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationState) DeepCopyInto(out *InstallationState) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastError != nil {
		in, out := &in.LastError, &out.LastError
		*out = new(v1alpha1.Error)
		(*in).DeepCopyInto(*out)
	}
	if in.Imports != nil {
		in, out := &in.Imports, &out.Imports
		*out = make([]v1alpha1.ImportStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationState.
func (in *InstallationState) DeepCopy() *InstallationState {
	if in == nil {
		return nil
	}
	out := new(InstallationState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalImportParameters) DeepCopyInto(out *InternalImportParameters) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]json.RawMessage, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(json.RawMessage, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalImportParameters.
func (in *InternalImportParameters) DeepCopy() *InternalImportParameters {
	if in == nil {
		return nil
	}
	out := new(InternalImportParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Job) DeepCopyInto(out *Job) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Job.
func (in *Job) DeepCopy() *Job {
	if in == nil {
		return nil
	}
	out := new(Job)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastOperation) DeepCopyInto(out *LastOperation) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
	if in.ErrorHistory != nil {
		in, out := &in.ErrorHistory, &out.ErrorHistory
		*out = new(ErrorHistory)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastOperation.
func (in *LastOperation) DeepCopy() *LastOperation {
	if in == nil {
		return nil
	}
	out := new(LastOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedSecretValues) DeepCopyInto(out *NamedSecretValues) {
	*out = *in
	if in.StringData != nil {
		in, out := &in.StringData, &out.StringData
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedSecretValues.
func (in *NamedSecretValues) DeepCopy() *NamedSecretValues {
	if in == nil {
		return nil
	}
	out := new(NamedSecretValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reachability) DeepCopyInto(out *Reachability) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reachability.
func (in *Reachability) DeepCopy() *Reachability {
	if in == nil {
		return nil
	}
	out := new(Reachability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Readiness) DeepCopyInto(out *Readiness) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Readiness.
func (in *Readiness) DeepCopy() *Readiness {
	if in == nil {
		return nil
	}
	out := new(Readiness)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadyRequirements) DeepCopyInto(out *ReadyRequirements) {
	*out = *in
	if in.Jobs != nil {
		in, out := &in.Jobs, &out.Jobs
		*out = make([]Job, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]Resource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadyRequirements.
func (in *ReadyRequirements) DeepCopy() *ReadyRequirements {
	if in == nil {
		return nil
	}
	out := new(ReadyRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	if in.SuccessValues != nil {
		in, out := &in.SuccessValues, &out.SuccessValues
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretValues) DeepCopyInto(out *SecretValues) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretValues.
func (in *SecretValues) DeepCopy() *SecretValues {
	if in == nil {
		return nil
	}
	out := new(SecretValues)
	in.DeepCopyInto(out)
	return out
}
