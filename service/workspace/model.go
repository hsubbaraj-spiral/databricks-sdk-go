// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import "fmt"

// all definitions in this file are in alphabetical order

type Delete struct {
	// The absolute path of the notebook or directory.
	Path string `json:"path"`
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	Recursive bool `json:"recursive,omitempty"`
}

// Export a notebook
type Export struct {
	// Flag to enable direct download. If it is `true`, the response will be the
	// exported file itself. Otherwise, the response contains content as base64
	// encoded string.
	DirectDownload bool `json:"-" url:"direct_download,omitempty"`
	// This specifies the format of the exported file. By default, this is
	// `SOURCE`. However it may be one of: `SOURCE`, `HTML`, `JUPYTER`, `DBC`.
	//
	// The value is case sensitive.
	Format ExportFormat `json:"-" url:"format,omitempty"`
	// The absolute path of the notebook or directory. Exporting directory is
	// only support for `DBC` format.
	Path string `json:"-" url:"path"`
}

// This specifies the format of the file to be imported. By default, this is
// `SOURCE`. However it may be one of: `SOURCE`, `HTML`, `JUPYTER`, `DBC`. The
// value is case sensitive.
type ExportFormat string

const ExportFormatDbc ExportFormat = `DBC`

const ExportFormatHtml ExportFormat = `HTML`

const ExportFormatJupyter ExportFormat = `JUPYTER`

const ExportFormatRMarkdown ExportFormat = `R_MARKDOWN`

const ExportFormatSource ExportFormat = `SOURCE`

// String representation for [fmt.Print]
func (ef *ExportFormat) String() string {
	return string(*ef)
}

// Set raw string value and validate it against allowed values
func (ef *ExportFormat) Set(v string) error {
	switch v {
	case `DBC`, `HTML`, `JUPYTER`, `R_MARKDOWN`, `SOURCE`:
		*ef = ExportFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DBC", "HTML", "JUPYTER", "R_MARKDOWN", "SOURCE"`, v)
	}
}

// Type always returns ExportFormat to satisfy [pflag.Value] interface
func (ef *ExportFormat) Type() string {
	return "ExportFormat"
}

type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** will be thrown.
	Content string `json:"content,omitempty"`
}

// Get status
type GetStatus struct {
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path"`
}

type Import struct {
	// The base64-encoded content. This has a limit of 10 MB.
	//
	// If the limit (10MB) is exceeded, exception with error code
	// **MAX_NOTEBOOK_SIZE_EXCEEDED** will be thrown. This parameter might be
	// absent, and instead a posted file will be used.
	Content string `json:"content,omitempty"`
	// This specifies the format of the file to be imported. By default, this is
	// `SOURCE`. However it may be one of: `SOURCE`, `HTML`, `JUPYTER`, `DBC`.
	// The value is case sensitive.
	Format ExportFormat `json:"format,omitempty"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language Language `json:"language,omitempty"`
	// The flag that specifies whether to overwrite existing object. It is
	// `false` by default. For `DBC` format, `overwrite` is not supported since
	// it may contain a directory.
	Overwrite bool `json:"overwrite,omitempty"`
	// The absolute path of the notebook or directory. Importing directory is
	// only support for `DBC` format.
	Path string `json:"path"`
}

// The language of the object. This value is set only if the object type is
// `NOTEBOOK`.
type Language string

const LanguagePython Language = `PYTHON`

const LanguageR Language = `R`

const LanguageScala Language = `SCALA`

const LanguageSql Language = `SQL`

// String representation for [fmt.Print]
func (l *Language) String() string {
	return string(*l)
}

// Set raw string value and validate it against allowed values
func (l *Language) Set(v string) error {
	switch v {
	case `PYTHON`, `R`, `SCALA`, `SQL`:
		*l = Language(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PYTHON", "R", "SCALA", "SQL"`, v)
	}
}

// Type always returns Language to satisfy [pflag.Value] interface
func (l *Language) Type() string {
	return "Language"
}

// List contents
type List struct {
	// <content needed>
	NotebooksModifiedAfter int `json:"-" url:"notebooks_modified_after,omitempty"`
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path"`
}

type ListResponse struct {
	// List of objects.
	Objects []ObjectInfo `json:"objects,omitempty"`
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	Path string `json:"path"`
}

type ObjectInfo struct {
	// <content needed>
	CreatedAt int64 `json:"created_at,omitempty"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language Language `json:"language,omitempty"`
	// <content needed>
	ModifiedAt int64 `json:"modified_at,omitempty"`
	// <content needed>
	ObjectId int64 `json:"object_id,omitempty"`
	// The type of the object in workspace.
	ObjectType ObjectType `json:"object_type,omitempty"`
	// The absolute path of the object.
	Path string `json:"path,omitempty"`
	// <content needed>
	Size int64 `json:"size,omitempty"`
}

// The type of the object in workspace.
type ObjectType string

const ObjectTypeDirectory ObjectType = `DIRECTORY`

const ObjectTypeFile ObjectType = `FILE`

const ObjectTypeLibrary ObjectType = `LIBRARY`

const ObjectTypeNotebook ObjectType = `NOTEBOOK`

const ObjectTypeRepo ObjectType = `REPO`

// String representation for [fmt.Print]
func (ot *ObjectType) String() string {
	return string(*ot)
}

// Set raw string value and validate it against allowed values
func (ot *ObjectType) Set(v string) error {
	switch v {
	case `DIRECTORY`, `FILE`, `LIBRARY`, `NOTEBOOK`, `REPO`:
		*ot = ObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DIRECTORY", "FILE", "LIBRARY", "NOTEBOOK", "REPO"`, v)
	}
}

// Type always returns ObjectType to satisfy [pflag.Value] interface
func (ot *ObjectType) Type() string {
	return "ObjectType"
}
