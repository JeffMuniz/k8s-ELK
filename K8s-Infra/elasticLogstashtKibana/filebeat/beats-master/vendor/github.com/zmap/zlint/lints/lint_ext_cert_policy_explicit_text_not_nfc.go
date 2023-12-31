package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

/************************************************
  When the UTF8String encoding is used, all character sequences SHOULD be
  normalized according to Unicode normalization form C (NFC) [NFC].
************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"golang.org/x/net v0.7.0x/text/unicode/norm"
)

type ExtCertPolicyExplicitTextNotNFC struct{}

func (l *ExtCertPolicyExplicitTextNotNFC) Initialize() error {
	return nil
}

func (l *ExtCertPolicyExplicitTextNotNFC) CheckApplies(c *x509.Certificate) bool {
	for _, text := range c.ExplicitTexts {
		if text != nil {
			return true
		}
	}
	return false
}

func (l *ExtCertPolicyExplicitTextNotNFC) Execute(c *x509.Certificate) *LintResult {
	for _, firstLvl := range c.ExplicitTexts {
		for _, text := range firstLvl {
			if text.Tag == 12 || text.Tag == 30 {
				if !norm.NFC.IsNormal(text.Bytes) {
					return &LintResult{Status: Warn}
				}
			}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_cert_policy_explicit_text_not_nfc",
		Description:   "When utf8string or bmpstring encoding is used for explicitText field in certificate policy, it SHOULD be normalized by NFC format",
		Citation:      "RFC6181 3",
		Source:        RFC5280,
		EffectiveDate: util.RFC6818Date,
		Lint:          &ExtCertPolicyExplicitTextNotNFC{},
	})
}
