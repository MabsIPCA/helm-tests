{{/*
Helper templates for testing runtime failures in _helpers.tpl
These tests focus on errors that occur during template execution,
not parse-time errors (which would break all templates).
*/}}

{{/*
Helper with nil pointer access - accessing nested field on nil
*/}}
{{- define "helpers.nilPointer" -}}
{{ .Values.nilHelper.nested.value }}
{{- end -}}

{{/*
Helper with required function failure - missing required value
*/}}
{{- define "helpers.required" -}}
{{ required "missingRequiredField is required but was not provided" .Values.missingRequiredField }}
{{- end -}}

{{/*
Helper with type mismatch - trying to range over a number
*/}}
{{- define "helpers.rangeOverScalar" -}}
{{- range .Values.port }}
item: {{ . }}
{{- end }}
{{- end -}}

{{/*
Helper with division by zero
*/}}
{{- define "helpers.divZero" -}}
result: {{ div 100 0 }}
{{- end -}}

{{/*
Helper calling a non-existent template via include
*/}}
{{- define "helpers.missingInclude" -}}
{{ include "non.existent.helper.template" . }}
{{- end -}}

{{/*
Helper with fail function - explicit failure
*/}}
{{- define "helpers.explicitFail" -}}
{{ fail "This helper explicitly fails with a custom message" }}
{{- end -}}

