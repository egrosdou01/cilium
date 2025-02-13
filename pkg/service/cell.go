// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package service

import (
	"github.com/cilium/cilium/pkg/datapath/types"
	"github.com/cilium/cilium/pkg/envoy"
	"github.com/cilium/cilium/pkg/hive/cell"
	monitorAgent "github.com/cilium/cilium/pkg/monitor/agent"
)

// Cell provides access to the Service Manager.
var Cell = cell.Module(
	"service-manager",
	"Service Manager",

	cell.Provide(newServiceManager),
)

type serviceManagerParams struct {
	cell.In

	EnvoyXdsServer envoy.XDSServer
	Datapath       types.Datapath
	MonitorAgent   monitorAgent.Agent
}

func newServiceManager(params serviceManagerParams) ServiceManager {
	return NewService(params.MonitorAgent, params.EnvoyXdsServer, params.Datapath.LBMap())
}
