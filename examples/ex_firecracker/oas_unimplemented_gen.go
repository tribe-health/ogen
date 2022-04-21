// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

var _ Handler = UnimplementedHandler{}

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

// CreateSnapshot implements createSnapshot operation.
//
// Creates a snapshot of the microVM state. The microVM should be in the `Paused` state.
//
// PUT /snapshot/create
func (UnimplementedHandler) CreateSnapshot(ctx context.Context, req SnapshotCreateParams) (r CreateSnapshotRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateSyncAction implements createSyncAction operation.
//
// PUT /actions
func (UnimplementedHandler) CreateSyncAction(ctx context.Context, req InstanceActionInfo) (r CreateSyncActionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DescribeBalloonConfig implements describeBalloonConfig operation.
//
// GET /balloon
func (UnimplementedHandler) DescribeBalloonConfig(ctx context.Context) (r DescribeBalloonConfigRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DescribeBalloonStats implements describeBalloonStats operation.
//
// GET /balloon/statistics
func (UnimplementedHandler) DescribeBalloonStats(ctx context.Context) (r DescribeBalloonStatsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DescribeInstance implements describeInstance operation.
//
// GET /
func (UnimplementedHandler) DescribeInstance(ctx context.Context) (r DescribeInstanceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetExportVmConfig implements getExportVmConfig operation.
//
// Gets configuration for all VM resources.
//
// GET /vm/config
func (UnimplementedHandler) GetExportVmConfig(ctx context.Context) (r GetExportVmConfigRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetMachineConfiguration implements getMachineConfiguration operation.
//
// Gets the machine configuration of the VM. When called before the PUT operation, it will return the
// default values for the vCPU count (=1), memory size (=128 MiB). By default Hyperthreading is
// disabled and there is no CPU Template.
//
// GET /machine-config
func (UnimplementedHandler) GetMachineConfiguration(ctx context.Context) (r GetMachineConfigurationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// LoadSnapshot implements loadSnapshot operation.
//
// Loads the microVM state from a snapshot. Only accepted on a fresh Firecracker process (before
// configuring any resource other than the Logger and Metrics).
//
// PUT /snapshot/load
func (UnimplementedHandler) LoadSnapshot(ctx context.Context, req SnapshotLoadParams) (r LoadSnapshotRes, _ error) {
	return r, ht.ErrNotImplemented
}

// MmdsConfigPut implements  operation.
//
// Creates MMDS configuration to be used by the MMDS network stack.
//
// PUT /mmds/config
func (UnimplementedHandler) MmdsConfigPut(ctx context.Context, req MmdsConfig) (r MmdsConfigPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

// MmdsGet implements  operation.
//
// GET /mmds
func (UnimplementedHandler) MmdsGet(ctx context.Context) (r MmdsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// MmdsPatch implements  operation.
//
// PATCH /mmds
func (UnimplementedHandler) MmdsPatch(ctx context.Context, req *MmdsPatchReq) (r MmdsPatchRes, _ error) {
	return r, ht.ErrNotImplemented
}

// MmdsPut implements  operation.
//
// PUT /mmds
func (UnimplementedHandler) MmdsPut(ctx context.Context, req *MmdsPutReq) (r MmdsPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PatchBalloon implements patchBalloon operation.
//
// Updates an existing balloon device, before or after machine startup. Will fail if update is not
// possible.
//
// PATCH /balloon
func (UnimplementedHandler) PatchBalloon(ctx context.Context, req BalloonUpdate) (r PatchBalloonRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PatchBalloonStatsInterval implements patchBalloonStatsInterval operation.
//
// Updates an existing balloon device statistics interval, before or after machine startup. Will fail
// if update is not possible.
//
// PATCH /balloon/statistics
func (UnimplementedHandler) PatchBalloonStatsInterval(ctx context.Context, req BalloonStatsUpdate) (r PatchBalloonStatsIntervalRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PatchGuestDriveByID implements patchGuestDriveByID operation.
//
// Updates the properties of the drive with the ID specified by drive_id path parameter. Will fail if
// update is not possible.
//
// PATCH /drives/{drive_id}
func (UnimplementedHandler) PatchGuestDriveByID(ctx context.Context, req PartialDrive, params PatchGuestDriveByIDParams) (r PatchGuestDriveByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PatchGuestNetworkInterfaceByID implements patchGuestNetworkInterfaceByID operation.
//
// Updates the rate limiters applied to a network interface.
//
// PATCH /network-interfaces/{iface_id}
func (UnimplementedHandler) PatchGuestNetworkInterfaceByID(ctx context.Context, req PartialNetworkInterface, params PatchGuestNetworkInterfaceByIDParams) (r PatchGuestNetworkInterfaceByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PatchMachineConfiguration implements patchMachineConfiguration operation.
//
// Partially updates the Virtual Machine Configuration with the specified input. If any of the
// parameters has an incorrect value, the whole update fails.
//
// PATCH /machine-config
func (UnimplementedHandler) PatchMachineConfiguration(ctx context.Context, req OptMachineConfiguration) (r PatchMachineConfigurationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PatchVm implements patchVm operation.
//
// Sets the desired state (Paused or Resumed) for the microVM.
//
// PATCH /vm
func (UnimplementedHandler) PatchVm(ctx context.Context, req VM) (r PatchVmRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PutBalloon implements putBalloon operation.
//
// Creates a new balloon device if one does not already exist, otherwise updates it, before machine
// startup. This will fail after machine startup. Will fail if update is not possible.
//
// PUT /balloon
func (UnimplementedHandler) PutBalloon(ctx context.Context, req Balloon) (r PutBalloonRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PutGuestBootSource implements putGuestBootSource operation.
//
// Creates new boot source if one does not already exist, otherwise updates it. Will fail if update
// is not possible.
//
// PUT /boot-source
func (UnimplementedHandler) PutGuestBootSource(ctx context.Context, req BootSource) (r PutGuestBootSourceRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PutGuestDriveByID implements putGuestDriveByID operation.
//
// Creates new drive with ID specified by drive_id path parameter. If a drive with the specified ID
// already exists, updates its state based on new input. Will fail if update is not possible.
//
// PUT /drives/{drive_id}
func (UnimplementedHandler) PutGuestDriveByID(ctx context.Context, req Drive, params PutGuestDriveByIDParams) (r PutGuestDriveByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PutGuestNetworkInterfaceByID implements putGuestNetworkInterfaceByID operation.
//
// Creates new network interface with ID specified by iface_id path parameter.
//
// PUT /network-interfaces/{iface_id}
func (UnimplementedHandler) PutGuestNetworkInterfaceByID(ctx context.Context, req NetworkInterface, params PutGuestNetworkInterfaceByIDParams) (r PutGuestNetworkInterfaceByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PutGuestVsock implements putGuestVsock operation.
//
// The first call creates the device with the configuration specified in body. Subsequent calls will
// update the device configuration. May fail if update is not possible.
//
// PUT /vsock
func (UnimplementedHandler) PutGuestVsock(ctx context.Context, req Vsock) (r PutGuestVsockRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PutLogger implements putLogger operation.
//
// PUT /logger
func (UnimplementedHandler) PutLogger(ctx context.Context, req Logger) (r PutLoggerRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PutMachineConfiguration implements putMachineConfiguration operation.
//
// Updates the Virtual Machine Configuration with the specified input. Firecracker starts with
// default values for vCPU count (=1) and memory size (=128 MiB). With Hyperthreading enabled, the
// vCPU count is restricted to be 1 or an even number, otherwise there are no restrictions regarding
// the vCPU count. If any of the parameters has an incorrect value, the whole update fails.
//
// PUT /machine-config
func (UnimplementedHandler) PutMachineConfiguration(ctx context.Context, req OptMachineConfiguration) (r PutMachineConfigurationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PutMetrics implements putMetrics operation.
//
// PUT /metrics
func (UnimplementedHandler) PutMetrics(ctx context.Context, req Metrics) (r PutMetricsRes, _ error) {
	return r, ht.ErrNotImplemented
}
