//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.10.2, generator: @autorest/go@4.0.0-preview.63)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ServerFactory is a fake server for instances of the generated.ClientFactory type.
type ServerFactory struct {
	HcpClusterVersionsServer   HcpClusterVersionsServer
	HcpOpenShiftClustersServer HcpOpenShiftClustersServer
	NodePoolsServer            NodePoolsServer
	OperationsServer           OperationsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of generated.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of generated.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                          *ServerFactory
	trMu                         sync.Mutex
	trHcpClusterVersionsServer   *HcpClusterVersionsServerTransport
	trHcpOpenShiftClustersServer *HcpOpenShiftClustersServerTransport
	trNodePoolsServer            *NodePoolsServerTransport
	trOperationsServer           *OperationsServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "HcpClusterVersionsClient":
		initServer(s, &s.trHcpClusterVersionsServer, func() *HcpClusterVersionsServerTransport {
			return NewHcpClusterVersionsServerTransport(&s.srv.HcpClusterVersionsServer)
		})
		resp, err = s.trHcpClusterVersionsServer.Do(req)
	case "HcpOpenShiftClustersClient":
		initServer(s, &s.trHcpOpenShiftClustersServer, func() *HcpOpenShiftClustersServerTransport {
			return NewHcpOpenShiftClustersServerTransport(&s.srv.HcpOpenShiftClustersServer)
		})
		resp, err = s.trHcpOpenShiftClustersServer.Do(req)
	case "NodePoolsClient":
		initServer(s, &s.trNodePoolsServer, func() *NodePoolsServerTransport { return NewNodePoolsServerTransport(&s.srv.NodePoolsServer) })
		resp, err = s.trNodePoolsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
