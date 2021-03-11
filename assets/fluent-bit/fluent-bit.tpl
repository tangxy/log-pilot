{{range .configList}}
[INPUT]
    Name           tail
    Tag            log{{replace .HostDir "/" "."}}
    Path           {{ .HostDir }}/{{ .File }}
    DB             /run/fluent-bit/input-storage.db
    Mem_Buf_Limit  32MB
[FILTER]
    Name modify
    Match log{{replace .HostDir "/" "."}}
    {{range $key, $value := .Tags}}
    {{ $key }} {{ $value }}
    {{end}}
    {{range $key, $value := $.container}}
    {{ $key }} {{ $value }}
    {{end}}
{{end}}