package userdata

type Meta struct {
	ID  uint64 `bson:"_id"`
	Tag *Tag   `bson:"-"`
}

func New(id uint64) *Meta {
	return &Meta{
		ID: id,
	}
}

type Tag struct {
	IsOnline bool  // 是否在线
	GwID     int32 // 网关id
}
