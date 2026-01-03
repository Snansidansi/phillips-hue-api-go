package models

type MetadataPut struct {
	Name      *string `json:"name,omitempty"`
	Archetype *string `json:"archetype,omitempty"`
}

type ResourceIdentifier struct {
	Rid   string `json:"rid"`
	Rtype string `json:"rtype"`
}
