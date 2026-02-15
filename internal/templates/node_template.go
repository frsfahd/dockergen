package templates

const NodeTemplate = `# syntax=docker/dockerfile:1
FROM {{ .BaseImageBuilder }} AS deps
WORKDIR {{ .WorkDir }}
COPY package*.json ./
RUN {{ .InstallCommand }}

FROM {{ .BaseImageBuilder }} AS build
WORKDIR {{ .WorkDir }}
{{ .CopyDepsCommand }}
COPY . .
{{- if .BuildCommand }}
RUN {{ .BuildCommand }}
{{- end }}

FROM {{ .BaseImageRuntime }} AS runner
WORKDIR {{ .WorkDir }}
ENV NODE_ENV=production
COPY --from=build {{ .WorkDir }} {{ .WorkDir }}
{{- if .ExposePort }}
EXPOSE {{ .Port }}
{{- end }}
CMD {{ .StartJSON }}
`
