package models

type Groups struct {
	Groups *[]Group `json:"groups" yaml:"groups"`
}

type Group struct {
	Caption string `json:"caption" yaml:"caption"`
	Tags    *[]Tag `json:"tags" yaml:"tags"`
}

type Tag struct {
	Caption string      `json:"caption" yaml:"caption"`
	Links   *[]TagLinks `json:"links" yaml:"links"`
}

type TagLinks struct {
	Caption string `json:"caption" yaml:"caption"`
	Link    string `json:"link" yaml:"link"`
}
