// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package pb

import "storj.io/storj/pkg/storj"

// NodeIDsToLookupRequests ...
func NodeIDsToLookupRequests(nodeIDs storj.NodeIDList) *LookupRequests {
	var rq []*LookupRequest
	for _, v := range nodeIDs {
		r := &LookupRequest{NodeId: v}
		rq = append(rq, r)
	}
	return &LookupRequests{LookupRequest: rq}
}

// LookupResponsesToNodes ...
func LookupResponsesToNodes(responses *LookupResponses) []*Node {
	var nodes []*Node
	for _, v := range responses.LookupResponse {
		n := v.Node
		nodes = append(nodes, n)
	}
	return nodes
}

// CopyNode returns a deep copy of a node
// It would be better to use `proto.Clone` but it is curently incompatible
// with gogo's customtype extension.
// (see https://github.com/gogo/protobuf/issues/147)
func CopyNode(src *Node) (dst *Node) {
	node := Node{Id: storj.NodeID{}}
	copy(node.Id[:], src.Id[:])
	if src.Address != nil {
		node.Address = &NodeAddress{
			Transport: src.Address.Transport,
			Address:   src.Address.Address,
		}
	}
	if src.Metadata != nil {
		node.Metadata = &NodeMetadata{
			Email:  src.Metadata.Email,
			Wallet: src.Metadata.Wallet,
		}
	}
	if src.Restrictions != nil {
		node.Restrictions = &NodeRestrictions{
			FreeBandwidth: src.Restrictions.FreeBandwidth,
			FreeDisk:      src.Restrictions.FreeDisk,
		}
	}

	node.AuditSuccess = src.AuditSuccess
	node.IsUp = src.IsUp
	node.LatencyList = src.LatencyList
	node.Type = src.Type
	node.UpdateAuditSuccess = src.UpdateAuditSuccess
	node.UpdateLatency = src.UpdateLatency
	node.UpdateUptime = src.UpdateUptime

	return &node
}
