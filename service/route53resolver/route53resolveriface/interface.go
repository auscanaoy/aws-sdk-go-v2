// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package route53resolveriface provides an interface to enable mocking the Amazon Route 53 Resolver service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package route53resolveriface

import (
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

// Route53ResolverAPI provides an interface to enable mocking the
// route53resolver.Route53Resolver service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Route 53 Resolver.
//    func myFunc(svc route53resolveriface.Route53ResolverAPI) bool {
//        // Make svc.AssociateResolverEndpointIpAddress request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := route53resolver.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRoute53ResolverClient struct {
//        route53resolveriface.Route53ResolverAPI
//    }
//    func (m *mockRoute53ResolverClient) AssociateResolverEndpointIpAddress(input *route53resolver.AssociateResolverEndpointIpAddressInput) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRoute53ResolverClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type Route53ResolverAPI interface {
	AssociateResolverEndpointIpAddressRequest(*route53resolver.AssociateResolverEndpointIpAddressInput) route53resolver.AssociateResolverEndpointIpAddressRequest

	AssociateResolverRuleRequest(*route53resolver.AssociateResolverRuleInput) route53resolver.AssociateResolverRuleRequest

	CreateResolverEndpointRequest(*route53resolver.CreateResolverEndpointInput) route53resolver.CreateResolverEndpointRequest

	CreateResolverRuleRequest(*route53resolver.CreateResolverRuleInput) route53resolver.CreateResolverRuleRequest

	DeleteResolverEndpointRequest(*route53resolver.DeleteResolverEndpointInput) route53resolver.DeleteResolverEndpointRequest

	DeleteResolverRuleRequest(*route53resolver.DeleteResolverRuleInput) route53resolver.DeleteResolverRuleRequest

	DisassociateResolverEndpointIpAddressRequest(*route53resolver.DisassociateResolverEndpointIpAddressInput) route53resolver.DisassociateResolverEndpointIpAddressRequest

	DisassociateResolverRuleRequest(*route53resolver.DisassociateResolverRuleInput) route53resolver.DisassociateResolverRuleRequest

	GetResolverEndpointRequest(*route53resolver.GetResolverEndpointInput) route53resolver.GetResolverEndpointRequest

	GetResolverRuleRequest(*route53resolver.GetResolverRuleInput) route53resolver.GetResolverRuleRequest

	GetResolverRuleAssociationRequest(*route53resolver.GetResolverRuleAssociationInput) route53resolver.GetResolverRuleAssociationRequest

	GetResolverRulePolicyRequest(*route53resolver.GetResolverRulePolicyInput) route53resolver.GetResolverRulePolicyRequest

	ListResolverEndpointIpAddressesRequest(*route53resolver.ListResolverEndpointIpAddressesInput) route53resolver.ListResolverEndpointIpAddressesRequest

	ListResolverEndpointsRequest(*route53resolver.ListResolverEndpointsInput) route53resolver.ListResolverEndpointsRequest

	ListResolverRuleAssociationsRequest(*route53resolver.ListResolverRuleAssociationsInput) route53resolver.ListResolverRuleAssociationsRequest

	ListResolverRulesRequest(*route53resolver.ListResolverRulesInput) route53resolver.ListResolverRulesRequest

	ListTagsForResourceRequest(*route53resolver.ListTagsForResourceInput) route53resolver.ListTagsForResourceRequest

	PutResolverRulePolicyRequest(*route53resolver.PutResolverRulePolicyInput) route53resolver.PutResolverRulePolicyRequest

	TagResourceRequest(*route53resolver.TagResourceInput) route53resolver.TagResourceRequest

	UntagResourceRequest(*route53resolver.UntagResourceInput) route53resolver.UntagResourceRequest

	UpdateResolverEndpointRequest(*route53resolver.UpdateResolverEndpointInput) route53resolver.UpdateResolverEndpointRequest

	UpdateResolverRuleRequest(*route53resolver.UpdateResolverRuleInput) route53resolver.UpdateResolverRuleRequest
}

var _ Route53ResolverAPI = (*route53resolver.Route53Resolver)(nil)
