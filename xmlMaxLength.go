// Copyright 2020 - 2024 The xgen Authors. All rights reserved. Use of this
// source code is governed by a BSD-style license that can be found in the
// LICENSE file.
//
// Package xgen written in pure Go providing a set of functions that allow you
// to parse XSD (XML schema files). This library needs Go version 1.10 or
// later.

package xgen

import (
	"encoding/xml"
	"strconv"
)

func (opt *Options) OnMaxLength(ele xml.StartElement, protoTree []interface{}) (err error) {
	for _, attr := range ele.Attr {
		if attr.Name.Local == "value" {
			if opt.SimpleType.Peek() != nil {
				opt.SimpleType.Peek().(*SimpleType).Restriction.MaxLength, _ = strconv.Atoi(attr.Value)
				opt.SimpleType.Peek().(*SimpleType).Restriction.hasMaxLength = true
			}
		}
	}

	return
}

// EndMaxLength handles parsing event on the maxLength end elements. MaxLength
// specifies the maximum number of characters or list items allowed. Must be
// equal to or greater than zero.
func (opt *Options) EndMaxLength(ele xml.EndElement, protoTree []interface{}) (err error) {
	if opt.SimpleType.Len() > 0 && opt.Element.Len() > 0 {
		if opt.Element.Peek().(*Element).Type, err = opt.GetValueType(opt.SimpleType.Peek().(*SimpleType).Base, opt.ProtoTree); err != nil {
			return
		}
		opt.CurrentEle = ""
	}
	return
}
