// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package apache

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "apache", asset.ModuleFieldsPri, AssetApache); err != nil {
		panic(err)
	}
}

// AssetApache returns asset data.
// This is the base64 encoded gzipped contents of module/apache.
func AssetApache() string {
	return "eJysmE2PszYQx+/5FKPnvj70mEOlqlLVQytV2r0jx0zAXfDQ8ZAV377ibZ+8GGJDOEUm8///7BkMnjf4xO4IutGmxAOAWKnwCD9+GwZ+HABy9IZtI5bcEX49AACMN+FvytuqD/IlsWSG3NkWRxBu+8GzxSr3xyHgDZyucbb5ZRgDkK7BIxRMbTONBLwGv8pqjx7OxHDS5vNLcw6G6kaLPdnKSgdfVkqgKp8tJvtJ4prlhscY9P57OIQUCr+WYKxJMLPNzd1ZSffkd3caLeURPLVsUOk851uE/qptwXpchWk1H429r1TDJGSoSvQe10iNs1crQrEcxjYl8m6KoEwMw4nyLvPoRJ06wXu7ZyClSKMYfUPOo+q1gjIxIK1HzvqfiQh9nArExXjWKCXl2+b8X4teVFAharqcWnotV4rYFtbpLcXWY2cXZG/JbZlxODTGea6PzFCemt3bAvOipfUhnTiOMzInP2w3+V7QiLHXBTrZUNzZEBib+uVna9n/ftOGhY37WjLHizX3WVifUnBao07oAQ7Nbomm1v/SfU42wCzJRGNY9xKMBZlYjEaLKfdjLMnEYizkNJFiZ2VQyCQRgbw6t1UVejuloWQvqlPyu0u1h3lNtfYwOwuW/NLaprPsTVJ4m93CsqwU87ookBa+jLds14acWIdO9iz09OFdIKmnerHrbah1wl1mPYXe65vQnirGwlVkhr/th1pRioVhLCy5F+VvXSw6eVa6VxXUilTiCr2ulJ4LLqHNSMh8szOmn5IrvGDqUaGiQoXi4s5E3usi9Ts9HBXj19jUA1jDNJx7HyNj/GSzn5SMOlebXOu54ZNiPB3zhxpSQYWHr/8DPPSMDtdmMS2jsUU11qSahne2gRas+ut3cqKt85PF0KmSEmeMPz8+/oF35AvyZNZX9zdXiA1SWj2f2H0R36d0hbe/3t//glkVpoPxT6JFkJVez2aMUXMwURt3nxWXP6acMNXXSRl0k/KwUv+bpv5R4qTZJyJvjXXFQFhRUWA+70fq8H8AAAD//18WdVc="
}
