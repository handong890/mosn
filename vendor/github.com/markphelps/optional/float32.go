// Code generated by go generate
// This file was generated by robots at 2018-03-27 01:14:04.975426789 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Float32 is an optional float32
type Float32 struct {
	value *float32
}

// NewFloat32 creates a optional.Float32 from a float32
func NewFloat32(v float32) Float32 {
	return Float32{&v}
}

// Set sets the float32 value
func (f Float32) Set(v float32) {
	f.value = &v
}

// Get returns the float32 value or an error if not present
func (f Float32) Get() (float32, error) {
	if !f.Present() {
		return *f.value, errors.New("value not present")
	}
	return *f.value, nil
}

// Present returns whether or not the value is present
func (f Float32) Present() bool {
	return f.value != nil
}

// OrElse returns the float32 value or a default value if the value is not present
func (f Float32) OrElse(v float32) float32 {
	if f.Present() {
		return *f.value
	}
	return v
}

// If calls the function f with the value if the value is present
func (f Float32) If(fn func(float32)) {
	if f.Present() {
		fn(*f.value)
	}
}

func (f Float32) MarshalJSON() ([]byte, error) {
	if f.Present() {
		return json.Marshal(f.value)
	}
	var zero float32
	return json.Marshal(zero)
}

func (f *Float32) UnmarshalJSON(data []byte) error {
	var value float32

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	f.value = &value
	return nil
}
