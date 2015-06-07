//
// Copyright (c) 2014 The heketi Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package mock

type MockDB struct {
	nodes      map[uint64]*Node
	volumes    map[uint64]*Volume
	current_id uint64
}

type MockPlugin struct {
	db MockDB
}

func NewMockPlugin() *MockPlugin {
	m := &MockPlugin{}
	m.db.nodes = make(map[uint64]*Node)

	return m
}