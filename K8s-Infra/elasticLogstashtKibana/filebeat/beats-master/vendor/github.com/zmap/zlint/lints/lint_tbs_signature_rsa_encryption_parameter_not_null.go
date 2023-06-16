package lints

/*
 * ZLint Copyright 2019 Regents of the University of Michigan
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

/*******************************************************************************************************
"RFC5280: RFC 4055, Section 5"
RSA: Encoded algorithm identifier MUST have NULL parameters.
*******************************************************************************************************/

import (
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"golang.org/x/net v0.7.0x/crypto/cryptobyte"
	cryptobyte_asn1 "golang.org/x/net v0.7.0x/crypto/cryptobyte/asn1"
)

type rsaTBSSignatureEncryptionParamNotNULL struct{}

func (l *rsaTBSSignatureEncryptionParamNotNULL) Initialize() error {
	return nil
}

func (l *rsaTBSSignatureEncryptionParamNotNULL) CheckApplies(c *x509.Certificate) bool {
	_, ok := util.RSAAlgorithmIDToDER[c.SignatureAlgorithmOID.String()]
	return ok
}

func (l *rsaTBSSignatureEncryptionParamNotNULL) Execute(c *x509.Certificate) *LintResult {
	input := cryptobyte.String(c.RawTBSCertificate)

	var tbsCert cryptobyte.String
	if !input.ReadASN1(&tbsCert, cryptobyte_asn1.SEQUENCE) {
		return &LintResult{Status: Fatal, Details: "error reading tbsCertificate"}
	}

	if !tbsCert.SkipOptionalASN1(cryptobyte_asn1.Tag(0).Constructed().ContextSpecific()) {
		return &LintResult{Status: Fatal, Details: "error reading tbsCertificate.version"}
	}

	if !tbsCert.SkipASN1(cryptobyte_asn1.INTEGER) {
		return &LintResult{Status: Fatal, Details: "error reading tbsCertificate.serialNumber"}
	}

	var signatureAlgoID cryptobyte.String
	var tag cryptobyte_asn1.Tag
	// use ReadAnyElement to preserve tag and length octets
	if !tbsCert.ReadAnyASN1Element(&signatureAlgoID, &tag) {
		return &LintResult{Status: Fatal, Details: "error reading tbsCertificate.signature"}
	}

	if err := util.CheckAlgorithmIDParamNotNULL(signatureAlgoID, c.SignatureAlgorithmOID); err != nil {
		return &LintResult{Status: Error, Details: fmt.Sprintf("certificate tbsCertificate.signature %s", err.Error())}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_tbs_signature_rsa_encryption_parameter_not_null",
		Description:   "RSA: Encoded signature algorithm identifier MUST have NULL parameters",
		Citation:      "RFC 4055, Section 5",
		Source:        RFC5280, // RFC4055 is referenced in RFC5280, Section 1
		EffectiveDate: util.RFC5280Date,
		Lint:          &rsaTBSSignatureEncryptionParamNotNULL{},
	})
}
