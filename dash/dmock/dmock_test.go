package dmock_test

import (
	"context"
	"testing"

	"aduu.dev/utils/dash"
	"aduu.dev/utils/dash/dmock"
	"github.com/stretchr/testify/assert"
)

func TestMock(t *testing.T) {
	m := dmock.New()

	want := `/dev/disk1s1 on / (apfs, NFS exported, local, read-only, journaled)
	devfs on /dev (devfs, local, nobrowse)
	/dev/disk1s5 on /System/Volumes/Data (apfs, local, journaled, nobrowse)
	/dev/disk1s4 on /private/var/vm (apfs, local, journaled, nobrowse)
	map auto_home on /home (autofs, automounted, nobrowse)
	/dev/disk2s2 on /Volumes/Backup Drive (hfs, local, nodev, nosuid, journaled)
	/dev/disk3 on /Volumes/DerivedData-ramdisk (hfs, local, nodev, nosuid, noowners, mounted by fabiansturm)
	/dev/disk15 on /Volumes/DerivedData-ramdisk 1 (hfs, local, nodev, nosuid, noowners, mounted by fabiansturm)`

	m.On("RunWithOutputE",
		context.Background(),
		&dash.SplitResult{
			Name: "mount",
			Args: []string{},
		},
		[]dash.SettingsFunc(nil),
	).Return(
		want,
		nil,
	)

	got, err := m.RunWithOutputE(
		context.Background(),
		dash.TemplateSplitExpand(`mount`, ""))
	if err != nil {
		return
	}

	if !assert.Equal(t, want, got) {
		t.Fail()
	}

	m.AssertExpectations(t)
}
