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

// Handler hanldes ...
type Handler struct {
	Transformers []Transformer
	Senders      []Sender
}

// Sender is anything that can send a Container
type Sender interface {
	Send(c *Container) error
}

// Transformer is anything that can transform an Container
type Transformer interface {
	Transform(c *Container) error
}

// Receiver is anything that can start
type Receiver interface {
	Start() error
}
