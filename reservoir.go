/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package metrics

// A statistically representative reservoir of a data stream.
type Reservoir interface {
	// Returns the number of values recorded.
	Size() int
	// Adds a new recorded value to the reservoir.
	UpdateN(value int64)
	// Returns a snapshot of the reservoir's values.
	GetSnapshot() Snapshot
}

type ReservoirType int

const (
	BucketReservoirType ReservoirType = iota
	UniformReservoirType
	ExponentiallyDecayingReservoirType
)
