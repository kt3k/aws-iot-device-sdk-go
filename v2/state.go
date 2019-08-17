// Copyright 2018 SEQSENSE, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package awsiotdev

type deviceState int

const (
	inactive deviceState = iota
	established
	terminating
)

func (s deviceState) String() string {
	switch s {
	case inactive:
		return "inactive"
	case established:
		return "established"
	case terminating:
		return "terminating"
	default:
		return "unknown"
	}
}

func (s deviceState) isActive() bool {
	return s == established
}