{{range .configList}}
- type: log
  enabled: true
  paths:
      - {{ .HostDir }}/{{ .File }}
  scan_frequency: 10s
  harvester_buffer_size: 1048576
  fields_under_root: true
  {{if .Stdout}}
  docker-json: true
  {{end}}
  {{if eq .Format "json"}}
  json.keys_under_root: true
  {{end}}
  fields:
      {{range $key, $value := .Tags}}
      {{ $key }}: {{ $value }}
      {{end}}
      {{range $key, $value := $.container}}
      {{ $key }}: {{ $value }}
      {{end}}
  tail_files: false
  close_inactive: 2h
  close_eof: false
  close_removed: true
  clean_removed: true
  close_renamed: false
  multiline.pattern: '^[[:space:]]+(at|\.{3})\b|^Caused by:'
  multiline.negate: false
  multiline.match: after
{{end}}
