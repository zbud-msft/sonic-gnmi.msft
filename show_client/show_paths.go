package show_client

import (
	sdc "github.com/sonic-net/sonic-gnmi/sonic_data_client"
)

// All SHOW path and getters are defined here
func init() {
	sdc.RegisterCliPath(
		[]string{"SHOW", "buffer_pool", "persistent-watermark"},
		getBufferPoolPersistentWatermark,
		"SHOW/buffer_pool/persistent-watermark[OPTIONS]: Show persistent WM for buffer pools",
		0,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "buffer_pool", "watermark"},
		getBufferPoolWatermark,
		"SHOW/buffer_pool/watermark[OPTIONS]: Show user WM for buffer pools",
		0,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
	)

	sdc.RegisterCliPath(
		[]string{"SHOW", "clock"},
		getDate,
		"SHOW/clock[OPTIONS]: Show date and time",
		0,
		map[string]string{
			"timezones": "show/clock/timezones: List of available timezones",
		},
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "clock", "timezones"},
		getDateTimezone,
		"SHOW/clock/timezones[OPTIONS]: List of available timezones",
		0,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "dropcounters", "capabilities"},
		getDropcountersCapabilities,
		"SHOW/dropcounters/capabilities[OPTIONS]: Show device drop counters capabilities",
		0,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "dropcounters", "counts"},
		getDropCounters,
		"SHOW/dropcounters/counts[OPTIONS]: Show drop counts",
		0,
		nil,
		showCmdOptionGroup,
		showCmdOptionCounterType,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "headroom-pool", "persistent-watermark"},
		getHeadroomPoolPersistentWatermark,
		"SHOW/headroom-pool/persistent-watermark[OPTIONS]: Show persistent WM for headroom pool",
		0,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "headroom-pool", "watermark"},
		getHeadroomPoolWatermark,
		"SHOW/headroom-pool/watermark[OPTIONS]: Show user WM for headroom pool",
		0,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interfaces", "alias"},
		getInterfaceAlias,
		"SHOW/interfaces/alias/{INTERFACENAME}[OPTIONS]: Show Interface Name/Alias Mapping",
		1,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionDisplay,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interfaces", "counters"},
		getInterfaceCounters,
		"SHOW/interfaces/counters[OPTIONS]: Show interface counters",
		0,
		map[string]string{
			"detailed":      "show/interfaces/counters/detailed: Show interface counters detailed",
			"errors":        "show/interfaces/counters/errors: Show interface counters errors",
			"fec-histogram": "show/interfaces/counters/fec-histogram: Show interface counters fec-histogram",
			"fec-stats":     "show/interfaces/counters/fec-stats: Show interface counters rates",
			"rates":         "show/interfaces/counters/rates: Show interface counters rates",
			"rif":           "show/interfaces/counters/rif: Show interface counters rif",
			"trim":          "show/interfaces/counters/trim: Show interface counters trim",
		},
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionDisplay,
		showCmdOptionInterfaces,
		showCmdOptionPeriod,
		showCmdOptionJson,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interfaces", "errors"},
		getInterfaceErrors,
		"SHOW/interfaces/errors/INTERFACENAME[OPTIONS]: Show Interface Errors <interfacename>",
		1,
		nil,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interfaces", "fec", "status"},
		getInterfaceFecStatus,
		"SHOW/interfaces/fec/status/{INTERFACENAME}[OPTIONS]: Show interface fec status",
		1,
		nil,
		showCmdOptionInterface,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionDisplay,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interfaces", "flap"},
		getInterfaceFlap,
		"SHOW/interfaces/flap/{INTERFACENAME}[OPTIONS]: Show Interface Flap Information",
		1,
		nil,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interfaces", "status"},
		getInterfaceStatus,
		"SHOW/interfaces/status/{INTERFACENAME}[OPTIONS]: Show Interface status information",
		1,
		nil,
		showCmdOptionInterface,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionDisplay,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interfaces", "switchport", "config"},
		getInterfaceSwitchportConfig,
		"SHOW/interfaces/switchport/config[OPTIONS]: Show interface switchport config information",
		0,
		nil,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interfaces", "switchport", "status"},
		getInterfaceSwitchportStatus,
		"SHOW/interfaces/switchport/status[OPTIONS]: Show interface switchport status information",
		0,
		nil,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interfaces", "transceiver", "error-status"},
		getTransceiverErrorStatus,
		"SHOW/interfaces/transceiver/error-status/{INTERFACENAME}[OPTIONS]: Show transceiver error-status",
		1,
		nil,
		showCmdOptionVerbose,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		sdc.UnimplementedOption(showCmdOptionFetchFromHW),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "interfaces", "transceiver", "presence"},
		getInterfaceTransceiverPresence,
		"SHOW/interfaces/transceiver/presence/{INTERFACENAME}[OPTIONS]: Show interface transceiver presence",
		1,
		nil,
		showCmdOptionVerbose,
		sdc.UnimplementedOption(showCmdOptionNamespace),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "ipv6", "bgp", "neighbors"},
		getIPv6BGPNeighborsHandler,
		"SHOW/ipv6/bgp/neighbors/{IPADDRESS}/{routes|advertised-routes|received-routes}[OPTIONS]: Show IPv6 BGP neighbors",
		2,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "ipv6", "bgp", "network"},
		getIPv6BGPNetwork,
		"SHOW/ipv6/bgp/network/{ipv6-address|ipv6-prefix}/{bestpath|json|longer-prefixes|multipath}[OPTIONS]: Show BGP ipv6 network",
		2,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "ipv6", "bgp", "summary"},
		getIPv6BGPSummary,
		"SHOW/ipv6/bgp/summary[OPTIONS]: Show summarized information of IPv6 BGP state",
		0,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionDisplay,
	)

	sdc.RegisterCliPath(
		[]string{"SHOW", "ipv6", "fib"},
		getIPv6Fib,
		"SHOW/ipv6/fib/{IPADDRESS}[OPTIONS]: Show IP FIB table",
		1,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "ipv6", "interfaces"},
		getIPv6Interfaces,
		"SHOW/ipv6/interfaces[OPTIONS]: Show ipv6 interfaces",
		0,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionDisplay,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "ipv6", "link-local-mode"},
		getPortsIpv6LinkLocalMode,
		"SHOW/ipv6/link-local-mode[OPTIONS]: Show ipv6 link-local-mode",
		0,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "ipv6", "protocol"},
		getIPv6Protocol,
		"SHOW/ipv6/protocol[OPTIONS]: Show IPv6 protocol information",
		0,
		nil,
		showCmdOptionVerbose,
	)

	sdc.RegisterCliPath(
		[]string{"SHOW", "lldp", "neighbors"},
		getLLDPNeighbors,
		"SHOW/lldp/neighbors/{INTERFACENAME}[OPTIONS]: Show LLDP neighbors",
		1,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "lldp", "table"},
		getLLDPTable,
		"SHOW/lldp/table[OPTIONS]: Show LLDP neighbors in a tabular format",
		0,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "mac"},
		getMacTable,
		"SHOW/mac[OPTIONS]: Show MAC (FDB) entries",
		0,
		map[string]string{
			"aging-time": "show/mac/aging-time",
		},
		showCmdOptionVlan,
		showCmdOptionPort,
		showCmdOptionAddress,
		showCmdOptionType,
		showCmdOptionCount,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "mac", "aging-time"},
		getMacAgingTime,
		"SHOW/mac/aging-time[OPTIONS]: Show mac aging-time",
		0,
		nil,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "mmu"},
		getMmuConfig,
		"SHOW/mmu[OPTIONS]: Show mmu configuration",
		0,
		nil,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "processes"},
		getProcessesRoot,
		"SHOW/processes/COMMAND[OPTIONS]: Show process information",
		0,
		map[string]string{
			"summary": "show/processes/summary: Show processses info",
			"cpu":     "show/processes/cpu: Show processes CPU info",
			"mem":     "show/processes/mem: Show processes memory info",
		},
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "processes", "summary"},
		getProcessesSummary,
		"SHOW/processes/summary[OPTIONS]: Show processes info",
		0,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "queue", "counters"},
		getQueueCounters,
		"SHOW/queue/counters/{INTERFACENAME}[OPTIONS]: Show queue counters",
		1,
		nil,
		showCmdOptionDisplay,
		showCmdOptionNonzero,
		showCmdOptionTrim,
		sdc.UnimplementedOption(showCmdOptionNamespace),
		showCmdOptionVerbose,
		// Add all opton, voq unimplemeneted, json, 
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "reboot-cause"},
		getPreviousRebootCause,
		"SHOW/reboot-cause[OPTIONS]: Show cause of most recent reboot",
		0,
		map[string]string{
			"history": "show/reboot-cause/history: Show history of reboot-cause",
		},
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "reboot-cause", "history"},
		getRebootCauseHistory,
		"SHOW/reboot-cause/history[OPTIONS]: Show history of reboot-cause",
		0,
		nil,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "srv6", "stats"},
		getSRv6Stats,
		"SHOW/srv6/stats/{SID}[OPTIONS]: Show SRv6 counters statistics",
		1,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "system-memory"},
		getSystemMemory,
		"SHOW/system-memory[OPTIONS]: Show memory information",
		0,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "uptime"},
		getUptime,
		"SHOW/uptime[OPTIONS]: Show system uptime",
		0,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "version"},
		getVersion,
		"SHOW/version[OPTIONS]: Show version information",
		0,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "vlan", "brief"},
		getVlanBrief,
		"SHOW/vlan/brief[OPTIONS]: Show all bridge information",
		0,
		nil,
		showCmdOptionVerbose,
	)
	sdc.RegisterCliPath(
		[]string{"SHOW", "watermark", "telemetry", "interval"},
		getWatermarkTelemetryInterval,
		"SHOW/watermark/telemetry/interval[OPTIONS]: Show telemetry interval",
		0,
		nil,
	)
}
