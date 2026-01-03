package builders

import "github.com/snansidansi/hueapi/models"

type RoomBuilder struct {
	put  models.RoomPut
	post models.RoomPost
}

func NewRoomBuilder() *RoomBuilder {
	return &RoomBuilder{
		put:  models.RoomPut{},
		post: models.RoomPost{},
	}
}

func (b *RoomBuilder) WithName(name string) *RoomBuilder {
	if b.put.Metadata == nil {
		b.put.Metadata = &models.MetadataPut{}
	}
	b.put.Metadata.Name = &name

	if b.post.Metadata == nil {
		b.post.Metadata = &models.MetadataPut{}
	}
	b.post.Metadata.Name = &name
	return b
}

func (b *RoomBuilder) WithArchetype(archetype string) *RoomBuilder {
	if b.put.Metadata == nil {
		b.put.Metadata = &models.MetadataPut{}
	}
	b.put.Metadata.Archetype = &archetype

	if b.post.Metadata == nil {
		b.post.Metadata = &models.MetadataPut{}
	}
	b.post.Metadata.Archetype = &archetype
	return b
}

func (b *RoomBuilder) WithChildren(children []models.ResourceIdentifier) *RoomBuilder {
	b.put.Children = &children
	b.post.Children = &children
	return b
}

func (b *RoomBuilder) BuildPut() models.RoomPut {
	return b.put
}

func (b *RoomBuilder) BuildPost() models.RoomPost {
	b.post.Type = "room"
	return b.post
}
