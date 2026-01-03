package models

type RoomPost struct {
	Children *[]ResourceIdentifier `json:"children,omitempty"`
	Metadata *MetadataPut          `json:"metadata,omitempty"`
	Type     string                `json:"type,omitempty"`
}

type RoomPut struct {
	Children *[]ResourceIdentifier `json:"children,omitempty"`
	Metadata *MetadataPut          `json:"metadata,omitempty"`
}

type Room struct {
	ID       string `json:"id"`
	IDV1     string `json:"id_v1"`
	Children []struct {
		Rid   string `json:"rid"`
		Rtype string `json:"rtype"`
	} `json:"children"`
	Metadata struct {
		Name      string `json:"name"`
		Archetype string `json:"archetype"`
	} `json:"metadata"`
	Services []struct {
		Rid   string `json:"rid"`
		Rtype string `json:"rtype"`
	} `json:"services"`
	Type string `json:"type"`
}
