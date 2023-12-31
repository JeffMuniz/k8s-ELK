package ca

import (
	"context"

	"google.golang.org/x/net v0.7.0grpc/metadata"
	"google.golang.org/x/net v0.7.0grpc/peer"
)

const (
	certForwardedKey = "forwarded_cert"
	certCNKey        = "forwarded_cert_cn"
	certOUKey        = "forwarded_cert_ou"
	certOrgKey       = "forwarded_cert_org"
	remoteAddrKey    = "remote_addr"
)

// forwardedTLSInfoFromContext obtains forwarded TLS CN/OU from the grpc.MD
// object in ctx.
func forwardedTLSInfoFromContext(ctx context.Context) (remoteAddr string, cn string, org string, ous []string) {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md[remoteAddrKey]) != 0 {
		remoteAddr = md[remoteAddrKey][0]
	}
	if len(md[certCNKey]) != 0 {
		cn = md[certCNKey][0]
	}
	if len(md[certOrgKey]) != 0 {
		org = md[certOrgKey][0]
	}
	ous = md[certOUKey]
	return
}

func isForwardedRequest(ctx context.Context) bool {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md[certForwardedKey]) != 1 {
		return false
	}
	return md[certForwardedKey][0] == "true"
}

// WithMetadataForwardTLSInfo reads certificate from context and returns context where
// ForwardCert is set based on original certificate.
func WithMetadataForwardTLSInfo(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}

	ous := []string{}
	org := ""
	cn := ""

	certSubj, err := certSubjectFromContext(ctx)
	if err == nil {
		cn = certSubj.CommonName
		ous = certSubj.OrganizationalUnit
		if len(certSubj.Organization) > 0 {
			org = certSubj.Organization[0]
		}
	}

	// If there's no TLS cert, forward with blank TLS metadata.
	// Note that the presence of this blank metadata is extremely
	// important. Without it, it would look like manager is making
	// the request directly.
	md[certForwardedKey] = []string{"true"}
	md[certCNKey] = []string{cn}
	md[certOrgKey] = []string{org}
	md[certOUKey] = ous
	peer, ok := peer.FromContext(ctx)
	if ok {
		md[remoteAddrKey] = []string{peer.Addr.String()}
	}

	return metadata.NewOutgoingContext(ctx, md), nil
}
