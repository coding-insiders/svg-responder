package svgengine

import (
	"fmt"
)

func Run() {
	out, err := Render(SvgRequest{
		Maps: []IMap{
			IMap{
				Breakpoint: "0",
				ImageUrl:   "http://via.placeholder.com/450x250",
			},
			IMap{
				Breakpoint: "650",
				ImageUrl:   "http://via.placeholder.com/650x350",
			},
			IMap{
				Breakpoint: "750",
				ImageUrl:   "http://via.placeholder.com/750x450",
			},
			IMap{
				Breakpoint: "950",
				ImageUrl:   "http://via.placeholder.com/950x550",
			},
			IMap{
				Breakpoint: "1950",
				ImageUrl:   "http://via.placeholder.com/1950x850",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	o := fmt.Sprint(out)
	println(o)
}
