/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"encoding/json"
	"os"
)

// checks if the DebPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebPackage{}

// DebPackage A Serializer for Package.
type DebPackage struct {
	// A URI of a repository the new content unit should be associated with.
	Repository *string `json:"repository,omitempty"`
	// Artifact file representing the physical content
	Artifact *string `json:"artifact,omitempty"`
	// Path where the artifact is located relative to distributions base_path
	RelativePath *string `json:"relative_path,omitempty"`
	// An uploaded file that may be turned into the artifact of the content unit.
	File **os.File `json:"file,omitempty"`
	// An uncommitted upload that may be turned into the artifact of the content unit.
	Upload *string `json:"upload,omitempty"`
	// Name of the distribution.
	Distribution *string `json:"distribution,omitempty"`
	// Name of the component.
	Component *string `json:"component,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DebPackage DebPackage

// NewDebPackage instantiates a new DebPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebPackage() *DebPackage {
	this := DebPackage{}
	return &this
}

// NewDebPackageWithDefaults instantiates a new DebPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebPackageWithDefaults() *DebPackage {
	this := DebPackage{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *DebPackage) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackage) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *DebPackage) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *DebPackage) SetRepository(v string) {
	o.Repository = &v
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *DebPackage) GetArtifact() string {
	if o == nil || IsNil(o.Artifact) {
		var ret string
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackage) GetArtifactOk() (*string, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *DebPackage) HasArtifact() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given string and assigns it to the Artifact field.
func (o *DebPackage) SetArtifact(v string) {
	o.Artifact = &v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise.
func (o *DebPackage) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath) {
		var ret string
		return ret
	}
	return *o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackage) GetRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.RelativePath) {
		return nil, false
	}
	return o.RelativePath, true
}

// HasRelativePath returns a boolean if a field has been set.
func (o *DebPackage) HasRelativePath() bool {
	if o != nil && !IsNil(o.RelativePath) {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given string and assigns it to the RelativePath field.
func (o *DebPackage) SetRelativePath(v string) {
	o.RelativePath = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *DebPackage) GetFile() *os.File {
	if o == nil || IsNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackage) GetFileOk() (**os.File, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *DebPackage) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *DebPackage) SetFile(v *os.File) {
	o.File = &v
}

// GetUpload returns the Upload field value if set, zero value otherwise.
func (o *DebPackage) GetUpload() string {
	if o == nil || IsNil(o.Upload) {
		var ret string
		return ret
	}
	return *o.Upload
}

// GetUploadOk returns a tuple with the Upload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackage) GetUploadOk() (*string, bool) {
	if o == nil || IsNil(o.Upload) {
		return nil, false
	}
	return o.Upload, true
}

// HasUpload returns a boolean if a field has been set.
func (o *DebPackage) HasUpload() bool {
	if o != nil && !IsNil(o.Upload) {
		return true
	}

	return false
}

// SetUpload gets a reference to the given string and assigns it to the Upload field.
func (o *DebPackage) SetUpload(v string) {
	o.Upload = &v
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *DebPackage) GetDistribution() string {
	if o == nil || IsNil(o.Distribution) {
		var ret string
		return ret
	}
	return *o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackage) GetDistributionOk() (*string, bool) {
	if o == nil || IsNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *DebPackage) HasDistribution() bool {
	if o != nil && !IsNil(o.Distribution) {
		return true
	}

	return false
}

// SetDistribution gets a reference to the given string and assigns it to the Distribution field.
func (o *DebPackage) SetDistribution(v string) {
	o.Distribution = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *DebPackage) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackage) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *DebPackage) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *DebPackage) SetComponent(v string) {
	o.Component = &v
}

func (o DebPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	if !IsNil(o.RelativePath) {
		toSerialize["relative_path"] = o.RelativePath
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.Upload) {
		toSerialize["upload"] = o.Upload
	}
	if !IsNil(o.Distribution) {
		toSerialize["distribution"] = o.Distribution
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DebPackage) UnmarshalJSON(bytes []byte) (err error) {
	varDebPackage := _DebPackage{}

	if err = json.Unmarshal(bytes, &varDebPackage); err == nil {
		*o = DebPackage(varDebPackage)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "repository")
		delete(additionalProperties, "artifact")
		delete(additionalProperties, "relative_path")
		delete(additionalProperties, "file")
		delete(additionalProperties, "upload")
		delete(additionalProperties, "distribution")
		delete(additionalProperties, "component")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDebPackage struct {
	value *DebPackage
	isSet bool
}

func (v NullableDebPackage) Get() *DebPackage {
	return v.value
}

func (v *NullableDebPackage) Set(val *DebPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableDebPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableDebPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebPackage(val *DebPackage) *NullableDebPackage {
	return &NullableDebPackage{value: val, isSet: true}
}

func (v NullableDebPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

