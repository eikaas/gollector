/* Package gollector is a collector of various data
 *
 * Copyright (c) 2019 Telenor Norge AS
 * Author(s):
 *  - Kristian Lyngstøl <kly@kly.no>
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
 * 02110-1301  USA
 */

package gollector

import (
	"fmt"
	"time"
)

// Metric is a metric
type Metric struct {
	// Time is the timestamp of the Metric
	Time *time.Time `json:"timestamp,omitempty"`
	// Metadata contains Metric metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Data contains the Metric  data
	Data map[string]interface{} `json:"data,omitempty"`
}

// Validate returns an error if the metric is invalid
func (m Metric) Validate() error {
	if m.Data == nil {
		return fmt.Errorf("Missing data for metric")
	}
	return nil
}
