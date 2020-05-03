package svgengine

type SvgRequest struct {
	Maps []IMap `json:"maps"`
}

type IMap struct {
	Breakpoint string `json:"breakpoint"`
	ImageUrl   string `json:"imageUrl"`
	ImageFile  string `json:"imageFile"`
	ImageType  string `json:"imageType"`
}
