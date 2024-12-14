package userops

import (
	"context"
	"eastv2/game/modules/play/userops/userdata"
	"eastv2/pb"
	"reflect"
	"time"

	"github.com/tnnmigga/corev2/conc"
	"github.com/tnnmigga/corev2/infra/mgdb"
	"github.com/tnnmigga/corev2/log"
	"github.com/tnnmigga/corev2/message"
	"github.com/tnnmigga/corev2/message/codec"
	"github.com/tnnmigga/corev2/utils"
)

func (m *Manager) regHandle() {
	message.Handle(m, onC2SPackageMsg)
	message.Response(m, onCreateNewUserReq)
}

func onC2SPackageMsg(req *pb.C2SPackageMsg) {
	msg, err := codec.Decode(req.Body)
	if err != nil {
		log.Errorf("onC2SPackage decode error %s", err)
		return
	}
	LoadAsync(req.UserID, func(m *userdata.Meta, err error) {
		if err != nil {
			log.Errorf("onC2SPackage LoadAsync error %v", err)
			return
		}
		h, ok := msgHandler[reflect.TypeOf(msg)]
		if !ok {
			log.Error("onC2SPackage msg handler not found %s", utils.TypeName(msg))
			return
		}
		h(m, msg)
	})
}

func onCreateNewUserReq(req *pb.CreateNewUserReq, response func(*pb.CreateNewUserResp, error)) {
	conc.Async(manager, func() (*userdata.Meta, error) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		user := userdata.New(req.UserID)
		_, err := mgdb.Default().Collection("userdata").InsertOne(ctx, user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}, func(user *userdata.Meta, err error) {
		if err != nil {
			log.Errorf("LoadAsync load error %v", err)
		} else {
			manager.cache[req.UserID] = user
		}
		log.Debugf("create user %d %v", req.UserID, err)
	}, groupKey(req.UserID))
}
