// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	s.cfg.NotFound(w, r)
}

func (s *Server) notAllowed(w http.ResponseWriter, r *http.Request, allowed string) {
	s.cfg.MethodNotAllowed(w, r, allowed)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch r.Method {
				case "GET":
					s.handleDescribeInstanceRequest([0]string{}, w, r)
				default:
					s.notAllowed(w, r, "GET")
				}

				return
			}
			switch elem[0] {
			case 'a': // Prefix: "actions"
				if l := len("actions"); len(elem) >= l && elem[0:l] == "actions" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "PUT":
						s.handleCreateSyncActionRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "PUT")
					}

					return
				}
			case 'b': // Prefix: "b"
				if l := len("b"); len(elem) >= l && elem[0:l] == "b" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "alloon"
					if l := len("alloon"); len(elem) >= l && elem[0:l] == "alloon" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleDescribeBalloonConfigRequest([0]string{}, w, r)
						case "PATCH":
							s.handlePatchBalloonRequest([0]string{}, w, r)
						case "PUT":
							s.handlePutBalloonRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET,PATCH,PUT")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/statistics"
						if l := len("/statistics"); len(elem) >= l && elem[0:l] == "/statistics" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleDescribeBalloonStatsRequest([0]string{}, w, r)
							case "PATCH":
								s.handlePatchBalloonStatsIntervalRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "GET,PATCH")
							}

							return
						}
					}
				case 'o': // Prefix: "oot-source"
					if l := len("oot-source"); len(elem) >= l && elem[0:l] == "oot-source" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "PUT":
							s.handlePutGuestBootSourceRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "PUT")
						}

						return
					}
				}
			case 'd': // Prefix: "drives/"
				if l := len("drives/"); len(elem) >= l && elem[0:l] == "drives/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "drive_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "PATCH":
						s.handlePatchGuestDriveByIDRequest([1]string{
							args[0],
						}, w, r)
					case "PUT":
						s.handlePutGuestDriveByIDRequest([1]string{
							args[0],
						}, w, r)
					default:
						s.notAllowed(w, r, "PATCH,PUT")
					}

					return
				}
			case 'l': // Prefix: "logger"
				if l := len("logger"); len(elem) >= l && elem[0:l] == "logger" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "PUT":
						s.handlePutLoggerRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "PUT")
					}

					return
				}
			case 'm': // Prefix: "m"
				if l := len("m"); len(elem) >= l && elem[0:l] == "m" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "achine-config"
					if l := len("achine-config"); len(elem) >= l && elem[0:l] == "achine-config" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetMachineConfigurationRequest([0]string{}, w, r)
						case "PATCH":
							s.handlePatchMachineConfigurationRequest([0]string{}, w, r)
						case "PUT":
							s.handlePutMachineConfigurationRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET,PATCH,PUT")
						}

						return
					}
				case 'e': // Prefix: "etrics"
					if l := len("etrics"); len(elem) >= l && elem[0:l] == "etrics" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "PUT":
							s.handlePutMetricsRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "PUT")
						}

						return
					}
				case 'm': // Prefix: "mds"
					if l := len("mds"); len(elem) >= l && elem[0:l] == "mds" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleMmdsGetRequest([0]string{}, w, r)
						case "PATCH":
							s.handleMmdsPatchRequest([0]string{}, w, r)
						case "PUT":
							s.handleMmdsPutRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET,PATCH,PUT")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/config"
						if l := len("/config"); len(elem) >= l && elem[0:l] == "/config" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "PUT":
								s.handleMmdsConfigPutRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "PUT")
							}

							return
						}
					}
				}
			case 'n': // Prefix: "network-interfaces/"
				if l := len("network-interfaces/"); len(elem) >= l && elem[0:l] == "network-interfaces/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "iface_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "PATCH":
						s.handlePatchGuestNetworkInterfaceByIDRequest([1]string{
							args[0],
						}, w, r)
					case "PUT":
						s.handlePutGuestNetworkInterfaceByIDRequest([1]string{
							args[0],
						}, w, r)
					default:
						s.notAllowed(w, r, "PATCH,PUT")
					}

					return
				}
			case 's': // Prefix: "snapshot/"
				if l := len("snapshot/"); len(elem) >= l && elem[0:l] == "snapshot/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "create"
					if l := len("create"); len(elem) >= l && elem[0:l] == "create" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "PUT":
							s.handleCreateSnapshotRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "PUT")
						}

						return
					}
				case 'l': // Prefix: "load"
					if l := len("load"); len(elem) >= l && elem[0:l] == "load" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "PUT":
							s.handleLoadSnapshotRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "PUT")
						}

						return
					}
				}
			case 'v': // Prefix: "v"
				if l := len("v"); len(elem) >= l && elem[0:l] == "v" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'm': // Prefix: "m"
					if l := len("m"); len(elem) >= l && elem[0:l] == "m" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "PATCH":
							s.handlePatchVmRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/config"
						if l := len("/config"); len(elem) >= l && elem[0:l] == "/config" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetExportVmConfigRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				case 's': // Prefix: "sock"
					if l := len("sock"); len(elem) >= l && elem[0:l] == "sock" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "PUT":
							s.handlePutGuestVsockRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "PUT")
						}

						return
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [1]string{}
		elem = path
	)
	r.args = args
	if elem == "" {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch method {
				case "GET":
					r.name = "DescribeInstance"
					r.operationID = "describeInstance"
					r.args = args
					r.count = 0
					return r, true
				default:
					return
				}
			}
			switch elem[0] {
			case 'a': // Prefix: "actions"
				if l := len("actions"); len(elem) >= l && elem[0:l] == "actions" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "PUT":
						// Leaf: CreateSyncAction
						r.name = "CreateSyncAction"
						r.operationID = "createSyncAction"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'b': // Prefix: "b"
				if l := len("b"); len(elem) >= l && elem[0:l] == "b" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "alloon"
					if l := len("alloon"); len(elem) >= l && elem[0:l] == "alloon" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "DescribeBalloonConfig"
							r.operationID = "describeBalloonConfig"
							r.args = args
							r.count = 0
							return r, true
						case "PATCH":
							r.name = "PatchBalloon"
							r.operationID = "patchBalloon"
							r.args = args
							r.count = 0
							return r, true
						case "PUT":
							r.name = "PutBalloon"
							r.operationID = "putBalloon"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/statistics"
						if l := len("/statistics"); len(elem) >= l && elem[0:l] == "/statistics" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: DescribeBalloonStats
								r.name = "DescribeBalloonStats"
								r.operationID = "describeBalloonStats"
								r.args = args
								r.count = 0
								return r, true
							case "PATCH":
								// Leaf: PatchBalloonStatsInterval
								r.name = "PatchBalloonStatsInterval"
								r.operationID = "patchBalloonStatsInterval"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				case 'o': // Prefix: "oot-source"
					if l := len("oot-source"); len(elem) >= l && elem[0:l] == "oot-source" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "PUT":
							// Leaf: PutGuestBootSource
							r.name = "PutGuestBootSource"
							r.operationID = "putGuestBootSource"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'd': // Prefix: "drives/"
				if l := len("drives/"); len(elem) >= l && elem[0:l] == "drives/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "drive_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					switch method {
					case "PATCH":
						// Leaf: PatchGuestDriveByID
						r.name = "PatchGuestDriveByID"
						r.operationID = "patchGuestDriveByID"
						r.args = args
						r.count = 1
						return r, true
					case "PUT":
						// Leaf: PutGuestDriveByID
						r.name = "PutGuestDriveByID"
						r.operationID = "putGuestDriveByID"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
			case 'l': // Prefix: "logger"
				if l := len("logger"); len(elem) >= l && elem[0:l] == "logger" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "PUT":
						// Leaf: PutLogger
						r.name = "PutLogger"
						r.operationID = "putLogger"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'm': // Prefix: "m"
				if l := len("m"); len(elem) >= l && elem[0:l] == "m" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "achine-config"
					if l := len("achine-config"); len(elem) >= l && elem[0:l] == "achine-config" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: GetMachineConfiguration
							r.name = "GetMachineConfiguration"
							r.operationID = "getMachineConfiguration"
							r.args = args
							r.count = 0
							return r, true
						case "PATCH":
							// Leaf: PatchMachineConfiguration
							r.name = "PatchMachineConfiguration"
							r.operationID = "patchMachineConfiguration"
							r.args = args
							r.count = 0
							return r, true
						case "PUT":
							// Leaf: PutMachineConfiguration
							r.name = "PutMachineConfiguration"
							r.operationID = "putMachineConfiguration"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'e': // Prefix: "etrics"
					if l := len("etrics"); len(elem) >= l && elem[0:l] == "etrics" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "PUT":
							// Leaf: PutMetrics
							r.name = "PutMetrics"
							r.operationID = "putMetrics"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'm': // Prefix: "mds"
					if l := len("mds"); len(elem) >= l && elem[0:l] == "mds" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "MmdsGet"
							r.operationID = ""
							r.args = args
							r.count = 0
							return r, true
						case "PATCH":
							r.name = "MmdsPatch"
							r.operationID = ""
							r.args = args
							r.count = 0
							return r, true
						case "PUT":
							r.name = "MmdsPut"
							r.operationID = ""
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/config"
						if l := len("/config"); len(elem) >= l && elem[0:l] == "/config" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "PUT":
								// Leaf: MmdsConfigPut
								r.name = "MmdsConfigPut"
								r.operationID = ""
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'n': // Prefix: "network-interfaces/"
				if l := len("network-interfaces/"); len(elem) >= l && elem[0:l] == "network-interfaces/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "iface_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					switch method {
					case "PATCH":
						// Leaf: PatchGuestNetworkInterfaceByID
						r.name = "PatchGuestNetworkInterfaceByID"
						r.operationID = "patchGuestNetworkInterfaceByID"
						r.args = args
						r.count = 1
						return r, true
					case "PUT":
						// Leaf: PutGuestNetworkInterfaceByID
						r.name = "PutGuestNetworkInterfaceByID"
						r.operationID = "putGuestNetworkInterfaceByID"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
			case 's': // Prefix: "snapshot/"
				if l := len("snapshot/"); len(elem) >= l && elem[0:l] == "snapshot/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "create"
					if l := len("create"); len(elem) >= l && elem[0:l] == "create" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "PUT":
							// Leaf: CreateSnapshot
							r.name = "CreateSnapshot"
							r.operationID = "createSnapshot"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'l': // Prefix: "load"
					if l := len("load"); len(elem) >= l && elem[0:l] == "load" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "PUT":
							// Leaf: LoadSnapshot
							r.name = "LoadSnapshot"
							r.operationID = "loadSnapshot"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'v': // Prefix: "v"
				if l := len("v"); len(elem) >= l && elem[0:l] == "v" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'm': // Prefix: "m"
					if l := len("m"); len(elem) >= l && elem[0:l] == "m" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "PATCH":
							r.name = "PatchVm"
							r.operationID = "patchVm"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/config"
						if l := len("/config"); len(elem) >= l && elem[0:l] == "/config" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetExportVmConfig
								r.name = "GetExportVmConfig"
								r.operationID = "getExportVmConfig"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				case 's': // Prefix: "sock"
					if l := len("sock"); len(elem) >= l && elem[0:l] == "sock" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "PUT":
							// Leaf: PutGuestVsock
							r.name = "PutGuestVsock"
							r.operationID = "putGuestVsock"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			}
		}
	}
	return r, false
}
