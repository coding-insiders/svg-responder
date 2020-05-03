package svgengine


var TEMPLATE = `
<svg xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink"
     viewBox="0 0 400 400" enable-background="new 0 0 400 400">
    <defs>
        <style>
        .shape {
            width: 100%;
        }
{{ $mpse := .Maps -}}
{{range $index1, $element := .Maps -}}
	{{if gt $index1 0 -}}
		@media (min-width: {{.Breakpoint}}px) {
	{{- end }}
	{{range $index2, $ele := $mpse -}}
			{{if eq $index2 $index1 -}}
			.shape.num-{{$index2}} {
				display: block;
			}
			{{- end}}
			{{if ne $index2 $index1 }}
			.shape.num-{{ $index2}} {
				display: none;
			}
			{{ end}}
	{{- end}}
	{{if gt $index1 0 -}}
		}
	{{- end }}
{{- end}}
        </style>
    </defs>
    {{range $index1, $element := .Maps -}}
    <g class="shape num-{{$index1}}">
	{{if ne .ImageUrl "" -}}
		<image xlink:href="{{ .ImageUrl }}" x="0" y="0" height="100%" width="100%"/>
	{{- else }}
		<image xlink:href="data:{{.ImageType}};base64,{{ .ImageFile }}" x="0" y="0" height="100%" width="100%"/>
	{{- end }}
    </g>
    {{- end }}
</svg>
`
