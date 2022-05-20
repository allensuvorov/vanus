// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package offset

import (
	"context"
	"testing"
	"time"

	"github.com/linkall-labs/vanus/internal/controller/trigger/storage"
	"github.com/linkall-labs/vanus/internal/primitive/info"
	"github.com/linkall-labs/vanus/internal/primitive/vanus"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetOffset(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storage := storage.NewMockOffsetStorage(ctrl)
	storage.EXPECT().DeleteOffset(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	m := NewOffsetManager(storage, 10*time.Microsecond).(*manager)
	subId := vanus.ID(1)
	eventLogID := vanus.ID(1)
	offset := uint64(1)

	Convey("get offset storage is empty", t, func() {
		storage.EXPECT().GetOffsets(gomock.Any(), subId).Return(info.ListOffsetInfo{}, nil)
		offsets, _ := m.GetOffset(ctx, subId)
		So(len(offsets), ShouldEqual, 0)
		subOffset, exist := m.subOffset.Load(subId)
		So(exist, ShouldBeTrue)
		So(subOffset, ShouldNotBeNil)
		m.RemoveRegisterSubscription(ctx, subId)
	})

	Convey("get offset storage has", t, func() {
		storage.EXPECT().GetOffsets(gomock.Any(), subId).Return(info.ListOffsetInfo{info.OffsetInfo{
			EventLogID: eventLogID,
			Offset:     offset,
		}}, nil)
		offsets, _ := m.GetOffset(ctx, subId)
		So(len(offsets), ShouldEqual, 1)
		So(offsets[0].Offset, ShouldEqual, offset)
		subOffset, exist := m.subOffset.Load(subId)
		So(exist, ShouldBeTrue)
		So(subOffset, ShouldNotBeNil)
	})
}

func TestSetOffset(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storage := storage.NewMockOffsetStorage(ctrl)
	storage.EXPECT().DeleteOffset(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	m := NewOffsetManager(storage, 10*time.Microsecond)
	subId := vanus.ID(1)
	eventLogID := vanus.ID(1)
	offset := uint64(1)

	Convey("set offset", t, func() {
		storage.EXPECT().GetOffsets(gomock.Any(), subId).Return(info.ListOffsetInfo{}, nil)
		m.Offset(ctx, subId, []info.OffsetInfo{{EventLogID: eventLogID, Offset: offset}})
		offsets, _ := m.GetOffset(ctx, subId)
		So(len(offsets), ShouldEqual, 1)
		So(offsets[0].Offset, ShouldEqual, offset)
		m.RemoveRegisterSubscription(ctx, subId)
	})
}

func TestCommit(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storage := storage.NewMockOffsetStorage(ctrl)
	storage.EXPECT().DeleteOffset(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	m := NewOffsetManager(storage, 10*time.Microsecond).(*manager)
	subId := vanus.ID(1)
	eventLogID := vanus.ID(1)
	offset := uint64(1)

	Convey("commit", t, func() {
		Convey("commit with storage create", func() {
			storage.EXPECT().GetOffsets(gomock.Any(), subId).Return(info.ListOffsetInfo{}, nil)
			storage.EXPECT().CreateOffset(gomock.Any(), subId, gomock.Any()).Return(nil)
			m.Offset(ctx, subId, []info.OffsetInfo{{EventLogID: eventLogID, Offset: offset}})
			offsets, _ := m.GetOffset(ctx, subId)
			So(len(offsets), ShouldEqual, 1)
			So(offsets[0].Offset, ShouldEqual, offset)
			m.commit(ctx)
			Convey("commit with storage update", func() {
				offset++
				m.Offset(ctx, subId, []info.OffsetInfo{{EventLogID: eventLogID, Offset: offset}})
				storage.EXPECT().UpdateOffset(gomock.Any(), subId, gomock.Any()).Return(nil)
				m.commit(ctx)
			})
		})
	})

}

func TestStart(t *testing.T) {
	ctx := context.Background()
	storage := storage.NewFakeStorage()
	commitInterval := 10 * time.Millisecond
	m := NewOffsetManager(storage, commitInterval)
	Convey("commit", t, func() {
		subId := vanus.ID(1)
		eventLogID := vanus.ID(1)
		offset := uint64(1)
		m.Offset(ctx, subId, []info.OffsetInfo{{EventLogID: eventLogID, Offset: offset}})
		Convey("commit storage created", func() {
			offsets, _ := m.GetOffset(ctx, subId)
			So(len(offsets), ShouldEqual, 1)
			So(offsets[0].Offset, ShouldEqual, offset)
			offsets, _ = storage.GetOffsets(ctx, subId)
			So(len(offsets), ShouldEqual, 0)
			m.Start()
			time.Sleep(2 * commitInterval)
			m.Stop()
			offsets, _ = storage.GetOffsets(ctx, subId)
			So(len(offsets), ShouldEqual, 1)
			So(offsets[0].Offset, ShouldEqual, offset)

			offset++
			m.Offset(ctx, subId, []info.OffsetInfo{{EventLogID: eventLogID, Offset: offset}})
			Convey("commit storage update", func() {
				offsets, _ = m.GetOffset(ctx, subId)
				So(len(offsets), ShouldEqual, 1)
				So(offsets[0].Offset, ShouldEqual, offset)
				offsets, _ = storage.GetOffsets(ctx, subId)
				So(len(offsets), ShouldEqual, 1)
				m.Start()
				time.Sleep(2 * commitInterval)
				m.Stop()
				offsets, _ = storage.GetOffsets(ctx, subId)
				So(len(offsets), ShouldEqual, 1)
				So(offsets[0].Offset, ShouldEqual, offset)
			})
		})

	})
}
