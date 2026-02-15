package templates

const PythonTemplate = `# syntax=docker/dockerfile:1
FROM {{ .BaseImageBuilder }} AS builder
WORKDIR {{ .WorkDir }}
COPY {{ .Requirements }} ./
RUN pip install --no-cache-dir --prefix=/install -r {{ .Requirements }}

FROM {{ .BaseImageRuntime }} AS runner
WORKDIR {{ .WorkDir }}
COPY --from=builder /install /usr/local
COPY . .
{{- if .ExposePort }}
EXPOSE {{ .Port }}
{{- end }}
CMD {{ .StartJSON }}
`
