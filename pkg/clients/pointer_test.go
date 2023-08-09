/*
Copyright 2019 The Crossplane Authors.

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

package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/go-cmp/cmp"
)

func TestLateInitStringPtrSlice(t *testing.T) {
	type args struct {
		in   []*string
		from []*string
	}

	cases := map[string]struct {
		args args
		want []*string
	}{
		"BothNil": {
			args: args{},
			want: nil,
		},
		"BothEmpty": {
			args: args{
				in:   []*string{},
				from: []*string{},
			},
			want: []*string{},
		},
		"FromNil": {
			args: args{
				in:   aws.StringSlice([]string{"hi!"}),
				from: nil,
			},
			want: aws.StringSlice([]string{"hi!"}),
		},
		"InNil": {
			args: args{
				in:   nil,
				from: aws.StringSlice([]string{"hi!"}),
			},
			want: aws.StringSlice([]string{"hi!"}),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := LateInitializeStringPtrSlice(tc.args.in, tc.args.from)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("\nLateInitializeStringPtrSlice(...): -got, +want:\n%s", diff)
			}
		})
	}
}

func TestLateInitInt64PtrSlice(t *testing.T) {
	type args struct {
		in   []*int64
		from []*int64
	}

	cases := map[string]struct {
		args args
		want []*int64
	}{
		"BothNil": {
			args: args{},
			want: nil,
		},
		"BothEmpty": {
			args: args{
				in:   []*int64{},
				from: []*int64{},
			},
			want: []*int64{},
		},
		"FromNil": {
			args: args{
				in:   aws.Int64Slice([]int64{1}),
				from: nil,
			},
			want: aws.Int64Slice([]int64{1}),
		},
		"InNil": {
			args: args{
				in:   nil,
				from: aws.Int64Slice([]int64{1}),
			},
			want: aws.Int64Slice([]int64{1}),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := LateInitializeInt64PtrSlice(tc.args.in, tc.args.from)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("\nLateInitializeInt64PtrSlice(...): -got, +want:\n%s", diff)
			}
		})
	}
}
