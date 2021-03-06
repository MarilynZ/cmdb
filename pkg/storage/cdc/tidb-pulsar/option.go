// Copyright 2020 Zhizhesihai (Beijing) Technology Limited.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package tidb_pulsar

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func parseOptions(u *url.URL) (opt *pulsar.ClientOptions, err error) {
	switch u.Scheme {
	case "pulsar", "pulsar+ssl":
	default:
		return nil, fmt.Errorf("unsupported pulsar scheme: %s", u.Scheme)
	}
	return parseClientOption(u)
}

func parseClientOption(u *url.URL) (opt *pulsar.ClientOptions, err error) {
	vs := values(u.Query())
	opt = &pulsar.ClientOptions{
		URL:                        (&url.URL{Scheme: u.Scheme, Host: u.Host}).String(),
		ConnectionTimeout:          vs.Duration("connectionTimeout"),
		OperationTimeout:           vs.Duration("operationTimeout"),
		TLSTrustCertsFilePath:      vs.Str("tlsTrustCertsFilePath"),
		TLSAllowInsecureConnection: vs.Bool("tlsAllowInsecureConnection"),
		TLSValidateHostname:        vs.Bool("tlsValidateHostname"),
		MaxConnectionsPerBroker:    vs.Int("maxConnectionsPerBroker"),
	}
	auth := vs.Str("auth")
	if auth == "" {
		if u.User.Username() == "" {
			// no auth
			return opt, nil
		}
		// use token provider by default
		opt.Authentication = pulsar.NewAuthenticationToken(u.User.Username())
		return opt, nil
	}
	param := jsonStr(vs.SubPathKV("auth"))
	opt.Authentication, err = pulsar.NewAuthentication(auth, param)
	return opt, nil
}

type values url.Values

func (vs values) Int(name string) int {
	value, ok := vs[name]
	if !ok {
		return 0
	}
	if len(value) == 0 {
		return 0
	}
	v, _ := strconv.Atoi(value[0])
	return v
}

func (vs values) Duration(name string) time.Duration {
	value, ok := vs[name]
	if !ok {
		return 0
	}
	if len(value) == 0 {
		return 0
	}
	v, _ := time.ParseDuration(value[0])
	return v
}

func (vs values) Bool(name string) bool {
	value, ok := vs[name]
	if !ok {
		return false
	}
	if len(value) == 0 {
		return true
	}
	v, _ := strconv.ParseBool(value[0])
	return v
}

func (vs values) Str(name string) string {
	value, ok := vs[name]
	if !ok {
		return ""
	}
	if len(value) == 0 {
		return ""
	}
	return value[0]
}

func (vs values) SubPathKV(prefix string) map[string]string {
	prefix = prefix + "."
	var m = map[string]string{}
	for name, value := range vs {
		if !strings.HasPrefix(name, prefix) {
			continue
		}
		var v string
		if len(value) != 0 {
			v = value[0]
		}
		m[name[len(prefix):]] = v
	}
	return m
}

func jsonStr(m interface{}) string {
	data, _ := json.Marshal(m)
	return string(data)
}
