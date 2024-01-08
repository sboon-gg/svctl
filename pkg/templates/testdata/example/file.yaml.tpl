pyBool: {{ pyBool .Values.boolTest }}
another: {{ .Values.test }}
quoted: {{ .Values.quoted | quote }}
negativeBool: {{ .Values.negativeBool }}
