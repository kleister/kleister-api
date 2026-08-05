{{- define "kleister-api.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "kleister-api.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "kleister-api.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "kleister-api.labels" -}}
helm.sh/chart: "{{ include "kleister-api.chart" . }}"
app.kubernetes.io/name: "kleister-api"
app.kubernetes.io/instance: "{{ .Release.Name }}"
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service | quote }}
{{- with .Values.labels }}
{{ toYaml . }}
{{- end }}
{{- end -}}

{{- define "kleister-api.server.labels" -}}
{{- include "kleister-api.labels" . }}
app.kubernetes.io/component: server
{{- end -}}

{{- define "kleister-api.cleanup.labels" -}}
{{- include "kleister-api.labels" . }}
app.kubernetes.io/component: cleanup
{{- end -}}

{{- define "kleister-api.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default "kleister-api" .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{- define "kleister-api.server.selectorLabels" -}}
app.kubernetes.io/name: kleister-api
app.kubernetes.io/component: server
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "kleister-api.database.secretName" -}}
{{- .Values.config.database.existingSecret | default (printf "%s-database" (include "kleister-api.fullname" .)) -}}
{{- end -}}

{{- define "kleister-api.token.secretName" -}}
{{- .Values.config.token.existingSecret | default (printf "%s-token" (include "kleister-api.fullname" .)) -}}
{{- end -}}

{{- define "kleister-api.admin.secretName" -}}
{{- .Values.config.admin.existingSecret | default (printf "%s-admin" (include "kleister-api.fullname" .)) -}}
{{- end -}}

{{- define "kleister-api.scim.secretName" -}}
{{- .Values.config.scim.existingSecret | default (printf "%s-scim" (include "kleister-api.fullname" .)) -}}
{{- end -}}

{{- define "kleister-api.shared.environment" -}}
- name: KLEISTER_API_LOG_LEVEL
  value: "{{ .Values.config.log.level }}"
- name: KLEISTER_API_LOG_PRETTY
  value: "false"
- name: KLEISTER_API_LOG_COLOR
  value: "false"
{{- if eq .Values.config.database.driver "sqlite3" }}
- name: KLEISTER_API_DATABASE_DRIVER
  value: "{{ .Values.config.database.driver }}"
- name: KLEISTER_API_DATABASE_NAME
  value: "{{ .Values.config.database.name }}"
{{- else }}
- name: KLEISTER_API_DATABASE_DRIVER
  value: "{{ .Values.config.database.driver }}"
- name: KLEISTER_API_DATABASE_ADDRESS
  value: "{{ .Values.config.database.address }}"
- name: KLEISTER_API_DATABASE_PORT
  value: "{{ .Values.config.database.port }}"
- name: KLEISTER_API_DATABASE_NAME
  value: "{{ .Values.config.database.name }}"
- name: KLEISTER_API_DATABASE_USERNAME
  value: "{{ .Values.config.database.username }}"
- name: KLEISTER_API_DATABASE_PASSWORD
  valueFrom:
    secretKeyRef:
      name: "{{ include "kleister-api.database.secretName" . }}"
      key: "{{ .Values.config.database.passwordKey }}"
{{- end }}
{{- if eq .Values.config.upload.driver "file" }}
- name: KLEISTER_API_UPLOAD_DRIVER
  value: "{{ .Values.config.upload.driver }}"
- name: KLEISTER_API_UPLOAD_PATH
  value: "{{ .Values.config.upload.path }}"
- name: KLEISTER_API_UPLOAD_PERMS
  value: "{{ .Values.config.upload.perms }}"
{{- end }}
{{- if eq .Values.config.upload.driver "s3" }}
- name: KLEISTER_API_UPLOAD_DRIVER
  value: "{{ .Values.config.upload.driver }}"
- name: KLEISTER_API_UPLOAD_ENDPOINT
  value: "{{ .Values.config.upload.endpoint }}"
- name: KLEISTER_API_UPLOAD_BUCKET
  value: "{{ .Values.config.upload.bucket }}"
- name: KLEISTER_API_UPLOAD_REGION
  value: "{{ .Values.config.upload.region }}"
- name: KLEISTER_API_UPLOAD_PATHSTYLE
  value: "{{ .Values.config.upload.pathstyle }}"
- name: KLEISTER_API_UPLOAD_PATH
  value: "{{ .Values.config.upload.path }}"
- name: KLEISTER_API_UPLOAD_ACCESS
  valueFrom:
    secretKeyRef:
      name: "{{ include "kleister-api.upload.secretName" . }}"
      key: "{{ .Values.config.upload.accessKey }}"
- name: KLEISTER_API_UPLOAD_SECRET
  valueFrom:
    secretKeyRef:
      name: "{{ include "kleister-api.upload.secretName" . }}"
      key: "{{ .Values.config.upload.secretKey }}"
- name: KLEISTER_API_UPLOAD_PROXY
  value: "{{ .Values.config.upload.proxy }}"
- name: KLEISTER_API_UPLOAD_PERMS
  value: "{{ .Values.config.upload.perms }}"
{{- end }}
{{- end -}}

{{- define "kleister-api.server.environment" -}}
{{- include "kleister-api.shared.environment" . }}
- name: KLEISTER_API_SERVER_HOST
  value: "{{ .Values.config.server.host }}"
- name: KLEISTER_API_SERVER_ROOT
  value: "{{ .Values.config.server.root }}"
- name: KLEISTER_API_SERVER_DOCS
  value: "{{ .Values.config.server.docs }}"
- name: KLEISTER_API_TOKEN_EXPIRE
  value: "{{ .Values.config.token.expire }}"
- name: KLEISTER_API_TOKEN_SECRET
  valueFrom:
    secretKeyRef:
      name: "{{ include "kleister-api.token.secretName" . }}"
      key: "{{ .Values.config.token.secretKey }}"
- name: KLEISTER_API_ADMIN_CREATE
  value: "{{ .Values.config.admin.create }}"
{{- if .Values.config.admin.create }}
- name: KLEISTER_API_ADMIN_USERNAME
  valueFrom:
    secretKeyRef:
      name: "{{ include "kleister-api.admin.secretName" . }}"
      key: "{{ .Values.config.admin.usernameKey }}"
- name: KLEISTER_API_ADMIN_PASSWORD
  valueFrom:
    secretKeyRef:
      name: "{{ include "kleister-api.admin.secretName" . }}"
      key: "{{ .Values.config.admin.passwordKey }}"
- name: KLEISTER_API_ADMIN_EMAIL
  valueFrom:
    secretKeyRef:
      name: "{{ include "kleister-api.admin.secretName" . }}"
      key: "{{ .Values.config.admin.emailKey }}"
{{- end }}
- name: KLEISTER_API_SCIM_ENABLED
  value: "{{ .Values.config.scim.enabled }}"
{{- if .Values.config.scim.enabled }}
- name: KLEISTER_API_SCIM_TOKEN
  valueFrom:
    secretKeyRef:
      name: "{{ include "kleister-api.scim.secretName" . }}"
      key: "{{ .Values.config.scim.tokenKey }}"
{{- end }}
- name: KLEISTER_API_AUTH_CONFIG
  value: "/etc/kleister-api/auth/config.yaml"
{{- end -}}

{{- define "kleister-api.cleanup.environment" -}}
{{- include "kleister-api.shared.environment" . }}
{{- end -}}
