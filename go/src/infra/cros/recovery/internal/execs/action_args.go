// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package execs

import (
	"context"
	"strconv"
	"strings"

	"infra/cros/recovery/internal/log"
)

const (
	// This character separates the name and values for extra
	// arguments defined for actions.
	DefaultSplitter = ":"

	// This character demarcates the individual values among
	// multi-valued extra arguments defined for actions.
	MultiValueSplitter = ","
)

// The map representing key-value pairs parsed from extra args in the
// configuration.
type ParsedArgs map[string]string

// AsBool returns the value for the passed key as a boolean. If the
// key does not exist in the parsed arguments, a default value of
// false is returned.
func (parsedArgs ParsedArgs) AsBool(ctx context.Context, key string) bool {
	defaultValue := false
	if value, ok := parsedArgs[key]; ok {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
		log.Debug(ctx, "Parsed Args As Bool: value %q for key %q is not a valid boolean, returning default value %t.", value, key, defaultValue)
	} else {
		log.Debug(ctx, "Parsed Args As Bool: key %q does not exist in the parsed arguments, returning default value %t.", key, defaultValue)
	}
	return defaultValue
}

// AsString returns the value for the passed key as a string. If the
// key does not exist in the parsed arguments, a default value of
// empty string is returned.
func (parsedArgs ParsedArgs) AsString(ctx context.Context, key string) string {
	defaultValue := ""
	if value, ok := parsedArgs[key]; ok {
		log.Debug(ctx, "Parsed Args As String: value %q found for key %q", value, key)
		return value
	}
	log.Debug(ctx, "Parsed Args As String: key %q not found, default value of empty string returned")
	return defaultValue
}

// AsStringSlice returns the value for the passed key as a slice of
// string. If the key does not exist in the parsed arguments, an empty
// slice is returned.
func (parsedArgs ParsedArgs) AsStringSlice(ctx context.Context, key string) []string {
	value := parsedArgs.AsString(ctx, key)
	if len(value) > 0 {
		return strings.Split(value, MultiValueSplitter)
	}
	return make([]string, 0)
}

// AsInt returns the value for the passed key as a int.
// @params defaultValue: if the value cannot be interpreted as int, then the passed in defaultValue is being returned.
func (parsedArgs ParsedArgs) AsInt(ctx context.Context, key string, defaultValue int) int {
	if value, ok := parsedArgs[key]; ok {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
		log.Debug(ctx, "Parsed Args As int: value %q for key %q is not a valid integer, returning default value %t.", value, key, defaultValue)
	} else {
		log.Debug(ctx, "Parsed Args As int: key %q does not exist in the parsed arguments, returning default value %t.", key, defaultValue)
	}
	return defaultValue
}

// ParseActionArgs parses the action arguments using the splitter, and
// returns ParsedArgs object containing key and values in the action
// arguments. If any mal-formed action arguments are found their value
// is set to empty string.
func ParseActionArgs(ctx context.Context, actionArgs []string, splitter string) ParsedArgs {
	parsedArgs := ParsedArgs(make(map[string]string))
	for _, a := range actionArgs {
		a := strings.TrimSpace(a)
		if a == "" {
			continue
		}
		log.Debug(ctx, "Parse Action Args: action arg %q", a)
		i := strings.Index(a, splitter)
		// Separator has to be at least second letter in the string to provide one letter key.
		if i < 1 {
			log.Debug(ctx, "Parse Action Args: malformed action arg %q", a)
			parsedArgs[a] = ""
		} else {
			k := strings.TrimSpace(a[:i])
			v := strings.TrimSpace(a[i+1:])
			log.Debug(ctx, "Parse Action Args: k: %q, v: %q", k, v)
			parsedArgs[k] = v
		}
	}
	return parsedArgs
}