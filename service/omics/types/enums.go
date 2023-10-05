// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Accelerators string

// Enum values for Accelerators
const (
	AcceleratorsGpu Accelerators = "GPU"
)

// Values returns all known values for Accelerators. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (Accelerators) Values() []Accelerators {
	return []Accelerators{
		"GPU",
	}
}

type AnnotationType string

// Enum values for AnnotationType
const (
	// Generic text file. No genomic information
	AnnotationTypeGeneric AnnotationType = "GENERIC"
	// Contains contig and 1-base position
	AnnotationTypeChrPos AnnotationType = "CHR_POS"
	// Contains contig, 1-base position, ref and alt allele information
	AnnotationTypeChrPosRefAlt AnnotationType = "CHR_POS_REF_ALT"
	// Contains contig, start, and end positions. Coordinates are 1-based
	AnnotationTypeChrStartEndOneBase AnnotationType = "CHR_START_END_ONE_BASE"
	// Contains contig, start, end, ref and alt allele information. Coordinates are
	// 1-based
	AnnotationTypeChrStartEndRefAltOneBase AnnotationType = "CHR_START_END_REF_ALT_ONE_BASE"
	// Contains contig, start, and end positions. Coordinates are 0-based
	AnnotationTypeChrStartEndZeroBase AnnotationType = "CHR_START_END_ZERO_BASE"
	// Contains contig, start, end, ref and alt allele information. Coordinates are
	// 0-based
	AnnotationTypeChrStartEndRefAltZeroBase AnnotationType = "CHR_START_END_REF_ALT_ZERO_BASE"
)

// Values returns all known values for AnnotationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AnnotationType) Values() []AnnotationType {
	return []AnnotationType{
		"GENERIC",
		"CHR_POS",
		"CHR_POS_REF_ALT",
		"CHR_START_END_ONE_BASE",
		"CHR_START_END_REF_ALT_ONE_BASE",
		"CHR_START_END_ZERO_BASE",
		"CHR_START_END_REF_ALT_ZERO_BASE",
	}
}

type CreationType string

// Enum values for CreationType
const (
	CreationTypeImport CreationType = "IMPORT"
	CreationTypeUpload CreationType = "UPLOAD"
)

// Values returns all known values for CreationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CreationType) Values() []CreationType {
	return []CreationType{
		"IMPORT",
		"UPLOAD",
	}
}

type EncryptionType string

// Enum values for EncryptionType
const (
	// KMS
	EncryptionTypeKms EncryptionType = "KMS"
)

// Values returns all known values for EncryptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionType) Values() []EncryptionType {
	return []EncryptionType{
		"KMS",
	}
}

type ETagAlgorithm string

// Enum values for ETagAlgorithm
const (
	ETagAlgorithmFastqMd5up ETagAlgorithm = "FASTQ_MD5up"
	ETagAlgorithmBamMd5up   ETagAlgorithm = "BAM_MD5up"
	ETagAlgorithmCramMd5up  ETagAlgorithm = "CRAM_MD5up"
)

// Values returns all known values for ETagAlgorithm. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ETagAlgorithm) Values() []ETagAlgorithm {
	return []ETagAlgorithm{
		"FASTQ_MD5up",
		"BAM_MD5up",
		"CRAM_MD5up",
	}
}

type FileType string

// Enum values for FileType
const (
	FileTypeFastq FileType = "FASTQ"
	FileTypeBam   FileType = "BAM"
	FileTypeCram  FileType = "CRAM"
)

// Values returns all known values for FileType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (FileType) Values() []FileType {
	return []FileType{
		"FASTQ",
		"BAM",
		"CRAM",
	}
}

type FormatToHeaderKey string

// Enum values for FormatToHeaderKey
const (
	FormatToHeaderKeyChr   FormatToHeaderKey = "CHR"
	FormatToHeaderKeyStart FormatToHeaderKey = "START"
	FormatToHeaderKeyEnd   FormatToHeaderKey = "END"
	FormatToHeaderKeyRef   FormatToHeaderKey = "REF"
	FormatToHeaderKeyAlt   FormatToHeaderKey = "ALT"
	FormatToHeaderKeyPos   FormatToHeaderKey = "POS"
)

// Values returns all known values for FormatToHeaderKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FormatToHeaderKey) Values() []FormatToHeaderKey {
	return []FormatToHeaderKey{
		"CHR",
		"START",
		"END",
		"REF",
		"ALT",
		"POS",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	// The Job has been submitted to run
	JobStatusSubmitted JobStatus = "SUBMITTED"
	// The Job is executing
	JobStatusInProgress JobStatus = "IN_PROGRESS"
	// The Job was cancelled
	JobStatusCancelled JobStatus = "CANCELLED"
	// The Job has completed
	JobStatusCompleted JobStatus = "COMPLETED"
	// The Job failed
	JobStatusFailed JobStatus = "FAILED"
	// The Job completed with failed runs
	JobStatusCompletedWithFailures JobStatus = "COMPLETED_WITH_FAILURES"
)

// Values returns all known values for JobStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"CANCELLED",
		"COMPLETED",
		"FAILED",
		"COMPLETED_WITH_FAILURES",
	}
}

type ReadSetActivationJobItemStatus string

// Enum values for ReadSetActivationJobItemStatus
const (
	ReadSetActivationJobItemStatusNotStarted ReadSetActivationJobItemStatus = "NOT_STARTED"
	ReadSetActivationJobItemStatusInProgress ReadSetActivationJobItemStatus = "IN_PROGRESS"
	ReadSetActivationJobItemStatusFinished   ReadSetActivationJobItemStatus = "FINISHED"
	ReadSetActivationJobItemStatusFailed     ReadSetActivationJobItemStatus = "FAILED"
)

// Values returns all known values for ReadSetActivationJobItemStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ReadSetActivationJobItemStatus) Values() []ReadSetActivationJobItemStatus {
	return []ReadSetActivationJobItemStatus{
		"NOT_STARTED",
		"IN_PROGRESS",
		"FINISHED",
		"FAILED",
	}
}

type ReadSetActivationJobStatus string

// Enum values for ReadSetActivationJobStatus
const (
	ReadSetActivationJobStatusSubmitted             ReadSetActivationJobStatus = "SUBMITTED"
	ReadSetActivationJobStatusInProgress            ReadSetActivationJobStatus = "IN_PROGRESS"
	ReadSetActivationJobStatusCancelling            ReadSetActivationJobStatus = "CANCELLING"
	ReadSetActivationJobStatusCancelled             ReadSetActivationJobStatus = "CANCELLED"
	ReadSetActivationJobStatusFailed                ReadSetActivationJobStatus = "FAILED"
	ReadSetActivationJobStatusCompleted             ReadSetActivationJobStatus = "COMPLETED"
	ReadSetActivationJobStatusCompletedWithFailures ReadSetActivationJobStatus = "COMPLETED_WITH_FAILURES"
)

// Values returns all known values for ReadSetActivationJobStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReadSetActivationJobStatus) Values() []ReadSetActivationJobStatus {
	return []ReadSetActivationJobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"CANCELLING",
		"CANCELLED",
		"FAILED",
		"COMPLETED",
		"COMPLETED_WITH_FAILURES",
	}
}

type ReadSetExportJobItemStatus string

// Enum values for ReadSetExportJobItemStatus
const (
	ReadSetExportJobItemStatusNotStarted ReadSetExportJobItemStatus = "NOT_STARTED"
	ReadSetExportJobItemStatusInProgress ReadSetExportJobItemStatus = "IN_PROGRESS"
	ReadSetExportJobItemStatusFinished   ReadSetExportJobItemStatus = "FINISHED"
	ReadSetExportJobItemStatusFailed     ReadSetExportJobItemStatus = "FAILED"
)

// Values returns all known values for ReadSetExportJobItemStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReadSetExportJobItemStatus) Values() []ReadSetExportJobItemStatus {
	return []ReadSetExportJobItemStatus{
		"NOT_STARTED",
		"IN_PROGRESS",
		"FINISHED",
		"FAILED",
	}
}

type ReadSetExportJobStatus string

// Enum values for ReadSetExportJobStatus
const (
	ReadSetExportJobStatusSubmitted             ReadSetExportJobStatus = "SUBMITTED"
	ReadSetExportJobStatusInProgress            ReadSetExportJobStatus = "IN_PROGRESS"
	ReadSetExportJobStatusCancelling            ReadSetExportJobStatus = "CANCELLING"
	ReadSetExportJobStatusCancelled             ReadSetExportJobStatus = "CANCELLED"
	ReadSetExportJobStatusFailed                ReadSetExportJobStatus = "FAILED"
	ReadSetExportJobStatusCompleted             ReadSetExportJobStatus = "COMPLETED"
	ReadSetExportJobStatusCompletedWithFailures ReadSetExportJobStatus = "COMPLETED_WITH_FAILURES"
)

// Values returns all known values for ReadSetExportJobStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReadSetExportJobStatus) Values() []ReadSetExportJobStatus {
	return []ReadSetExportJobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"CANCELLING",
		"CANCELLED",
		"FAILED",
		"COMPLETED",
		"COMPLETED_WITH_FAILURES",
	}
}

type ReadSetFile string

// Enum values for ReadSetFile
const (
	ReadSetFileSource1 ReadSetFile = "SOURCE1"
	ReadSetFileSource2 ReadSetFile = "SOURCE2"
	ReadSetFileIndex   ReadSetFile = "INDEX"
)

// Values returns all known values for ReadSetFile. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReadSetFile) Values() []ReadSetFile {
	return []ReadSetFile{
		"SOURCE1",
		"SOURCE2",
		"INDEX",
	}
}

type ReadSetImportJobItemStatus string

// Enum values for ReadSetImportJobItemStatus
const (
	ReadSetImportJobItemStatusNotStarted ReadSetImportJobItemStatus = "NOT_STARTED"
	ReadSetImportJobItemStatusInProgress ReadSetImportJobItemStatus = "IN_PROGRESS"
	ReadSetImportJobItemStatusFinished   ReadSetImportJobItemStatus = "FINISHED"
	ReadSetImportJobItemStatusFailed     ReadSetImportJobItemStatus = "FAILED"
)

// Values returns all known values for ReadSetImportJobItemStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReadSetImportJobItemStatus) Values() []ReadSetImportJobItemStatus {
	return []ReadSetImportJobItemStatus{
		"NOT_STARTED",
		"IN_PROGRESS",
		"FINISHED",
		"FAILED",
	}
}

type ReadSetImportJobStatus string

// Enum values for ReadSetImportJobStatus
const (
	ReadSetImportJobStatusSubmitted             ReadSetImportJobStatus = "SUBMITTED"
	ReadSetImportJobStatusInProgress            ReadSetImportJobStatus = "IN_PROGRESS"
	ReadSetImportJobStatusCancelling            ReadSetImportJobStatus = "CANCELLING"
	ReadSetImportJobStatusCancelled             ReadSetImportJobStatus = "CANCELLED"
	ReadSetImportJobStatusFailed                ReadSetImportJobStatus = "FAILED"
	ReadSetImportJobStatusCompleted             ReadSetImportJobStatus = "COMPLETED"
	ReadSetImportJobStatusCompletedWithFailures ReadSetImportJobStatus = "COMPLETED_WITH_FAILURES"
)

// Values returns all known values for ReadSetImportJobStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReadSetImportJobStatus) Values() []ReadSetImportJobStatus {
	return []ReadSetImportJobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"CANCELLING",
		"CANCELLED",
		"FAILED",
		"COMPLETED",
		"COMPLETED_WITH_FAILURES",
	}
}

type ReadSetPartSource string

// Enum values for ReadSetPartSource
const (
	ReadSetPartSourceSource1 ReadSetPartSource = "SOURCE1"
	ReadSetPartSourceSource2 ReadSetPartSource = "SOURCE2"
)

// Values returns all known values for ReadSetPartSource. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReadSetPartSource) Values() []ReadSetPartSource {
	return []ReadSetPartSource{
		"SOURCE1",
		"SOURCE2",
	}
}

type ReadSetStatus string

// Enum values for ReadSetStatus
const (
	ReadSetStatusArchived         ReadSetStatus = "ARCHIVED"
	ReadSetStatusActivating       ReadSetStatus = "ACTIVATING"
	ReadSetStatusActive           ReadSetStatus = "ACTIVE"
	ReadSetStatusDeleting         ReadSetStatus = "DELETING"
	ReadSetStatusDeleted          ReadSetStatus = "DELETED"
	ReadSetStatusProcessingUpload ReadSetStatus = "PROCESSING_UPLOAD"
	ReadSetStatusUploadFailed     ReadSetStatus = "UPLOAD_FAILED"
)

// Values returns all known values for ReadSetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReadSetStatus) Values() []ReadSetStatus {
	return []ReadSetStatus{
		"ARCHIVED",
		"ACTIVATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"PROCESSING_UPLOAD",
		"UPLOAD_FAILED",
	}
}

type ReferenceFile string

// Enum values for ReferenceFile
const (
	ReferenceFileSource ReferenceFile = "SOURCE"
	ReferenceFileIndex  ReferenceFile = "INDEX"
)

// Values returns all known values for ReferenceFile. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReferenceFile) Values() []ReferenceFile {
	return []ReferenceFile{
		"SOURCE",
		"INDEX",
	}
}

type ReferenceImportJobItemStatus string

// Enum values for ReferenceImportJobItemStatus
const (
	ReferenceImportJobItemStatusNotStarted ReferenceImportJobItemStatus = "NOT_STARTED"
	ReferenceImportJobItemStatusInProgress ReferenceImportJobItemStatus = "IN_PROGRESS"
	ReferenceImportJobItemStatusFinished   ReferenceImportJobItemStatus = "FINISHED"
	ReferenceImportJobItemStatusFailed     ReferenceImportJobItemStatus = "FAILED"
)

// Values returns all known values for ReferenceImportJobItemStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ReferenceImportJobItemStatus) Values() []ReferenceImportJobItemStatus {
	return []ReferenceImportJobItemStatus{
		"NOT_STARTED",
		"IN_PROGRESS",
		"FINISHED",
		"FAILED",
	}
}

type ReferenceImportJobStatus string

// Enum values for ReferenceImportJobStatus
const (
	ReferenceImportJobStatusSubmitted             ReferenceImportJobStatus = "SUBMITTED"
	ReferenceImportJobStatusInProgress            ReferenceImportJobStatus = "IN_PROGRESS"
	ReferenceImportJobStatusCancelling            ReferenceImportJobStatus = "CANCELLING"
	ReferenceImportJobStatusCancelled             ReferenceImportJobStatus = "CANCELLED"
	ReferenceImportJobStatusFailed                ReferenceImportJobStatus = "FAILED"
	ReferenceImportJobStatusCompleted             ReferenceImportJobStatus = "COMPLETED"
	ReferenceImportJobStatusCompletedWithFailures ReferenceImportJobStatus = "COMPLETED_WITH_FAILURES"
)

// Values returns all known values for ReferenceImportJobStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReferenceImportJobStatus) Values() []ReferenceImportJobStatus {
	return []ReferenceImportJobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"CANCELLING",
		"CANCELLED",
		"FAILED",
		"COMPLETED",
		"COMPLETED_WITH_FAILURES",
	}
}

type ReferenceStatus string

// Enum values for ReferenceStatus
const (
	ReferenceStatusActive   ReferenceStatus = "ACTIVE"
	ReferenceStatusDeleting ReferenceStatus = "DELETING"
	ReferenceStatusDeleted  ReferenceStatus = "DELETED"
)

// Values returns all known values for ReferenceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReferenceStatus) Values() []ReferenceStatus {
	return []ReferenceStatus{
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

type ResourceOwner string

// Enum values for ResourceOwner
const (
	// The resource owner is the calling account
	ResourceOwnerSelf ResourceOwner = "SELF"
	// The resource owner is an account other than the caller
	ResourceOwnerOther ResourceOwner = "OTHER"
)

// Values returns all known values for ResourceOwner. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceOwner) Values() []ResourceOwner {
	return []ResourceOwner{
		"SELF",
		"OTHER",
	}
}

type RunExport string

// Enum values for RunExport
const (
	RunExportDefinition RunExport = "DEFINITION"
)

// Values returns all known values for RunExport. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RunExport) Values() []RunExport {
	return []RunExport{
		"DEFINITION",
	}
}

type RunLogLevel string

// Enum values for RunLogLevel
const (
	RunLogLevelOff   RunLogLevel = "OFF"
	RunLogLevelFatal RunLogLevel = "FATAL"
	RunLogLevelError RunLogLevel = "ERROR"
	RunLogLevelAll   RunLogLevel = "ALL"
)

// Values returns all known values for RunLogLevel. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RunLogLevel) Values() []RunLogLevel {
	return []RunLogLevel{
		"OFF",
		"FATAL",
		"ERROR",
		"ALL",
	}
}

type RunRetentionMode string

// Enum values for RunRetentionMode
const (
	RunRetentionModeRetain RunRetentionMode = "RETAIN"
	RunRetentionModeRemove RunRetentionMode = "REMOVE"
)

// Values returns all known values for RunRetentionMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RunRetentionMode) Values() []RunRetentionMode {
	return []RunRetentionMode{
		"RETAIN",
		"REMOVE",
	}
}

type RunStatus string

// Enum values for RunStatus
const (
	RunStatusPending   RunStatus = "PENDING"
	RunStatusStarting  RunStatus = "STARTING"
	RunStatusRunning   RunStatus = "RUNNING"
	RunStatusStopping  RunStatus = "STOPPING"
	RunStatusCompleted RunStatus = "COMPLETED"
	RunStatusDeleted   RunStatus = "DELETED"
	RunStatusCancelled RunStatus = "CANCELLED"
	RunStatusFailed    RunStatus = "FAILED"
)

// Values returns all known values for RunStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RunStatus) Values() []RunStatus {
	return []RunStatus{
		"PENDING",
		"STARTING",
		"RUNNING",
		"STOPPING",
		"COMPLETED",
		"DELETED",
		"CANCELLED",
		"FAILED",
	}
}

type SchemaValueType string

// Enum values for SchemaValueType
const (
	// LONG type
	SchemaValueTypeLong SchemaValueType = "LONG"
	// INT type
	SchemaValueTypeInt SchemaValueType = "INT"
	// STRING type
	SchemaValueTypeString SchemaValueType = "STRING"
	// FLOAT type
	SchemaValueTypeFloat SchemaValueType = "FLOAT"
	// DOUBLE type
	SchemaValueTypeDouble SchemaValueType = "DOUBLE"
	// BOOLEAN type
	SchemaValueTypeBoolean SchemaValueType = "BOOLEAN"
)

// Values returns all known values for SchemaValueType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SchemaValueType) Values() []SchemaValueType {
	return []SchemaValueType{
		"LONG",
		"INT",
		"STRING",
		"FLOAT",
		"DOUBLE",
		"BOOLEAN",
	}
}

type ShareStatus string

// Enum values for ShareStatus
const (
	// The share has been created but is not yet active
	ShareStatusPending ShareStatus = "PENDING"
	// The share is activated
	ShareStatusActivating ShareStatus = "ACTIVATING"
	// The share is active and can be used
	ShareStatusActive ShareStatus = "ACTIVE"
	// The share is being deleted
	ShareStatusDeleting ShareStatus = "DELETING"
	// The share has been deleted
	ShareStatusDeleted ShareStatus = "DELETED"
	// The share has failed to activate or delete
	ShareStatusFailed ShareStatus = "FAILED"
)

// Values returns all known values for ShareStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ShareStatus) Values() []ShareStatus {
	return []ShareStatus{
		"PENDING",
		"ACTIVATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

type StoreFormat string

// Enum values for StoreFormat
const (
	// GFF3 Format
	StoreFormatGff StoreFormat = "GFF"
	// TSV Format
	StoreFormatTsv StoreFormat = "TSV"
	// VCF Format
	StoreFormatVcf StoreFormat = "VCF"
)

// Values returns all known values for StoreFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StoreFormat) Values() []StoreFormat {
	return []StoreFormat{
		"GFF",
		"TSV",
		"VCF",
	}
}

type StoreStatus string

// Enum values for StoreStatus
const (
	// The Store is being created
	StoreStatusCreating StoreStatus = "CREATING"
	// The Store is updating
	StoreStatusUpdating StoreStatus = "UPDATING"
	// The Store is deleting
	StoreStatusDeleting StoreStatus = "DELETING"
	// The Store is active
	StoreStatusActive StoreStatus = "ACTIVE"
	// The Store creation failed
	StoreStatusFailed StoreStatus = "FAILED"
)

// Values returns all known values for StoreStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StoreStatus) Values() []StoreStatus {
	return []StoreStatus{
		"CREATING",
		"UPDATING",
		"DELETING",
		"ACTIVE",
		"FAILED",
	}
}

type TaskStatus string

// Enum values for TaskStatus
const (
	TaskStatusPending   TaskStatus = "PENDING"
	TaskStatusStarting  TaskStatus = "STARTING"
	TaskStatusRunning   TaskStatus = "RUNNING"
	TaskStatusStopping  TaskStatus = "STOPPING"
	TaskStatusCompleted TaskStatus = "COMPLETED"
	TaskStatusCancelled TaskStatus = "CANCELLED"
	TaskStatusFailed    TaskStatus = "FAILED"
)

// Values returns all known values for TaskStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TaskStatus) Values() []TaskStatus {
	return []TaskStatus{
		"PENDING",
		"STARTING",
		"RUNNING",
		"STOPPING",
		"COMPLETED",
		"CANCELLED",
		"FAILED",
	}
}

type VersionStatus string

// Enum values for VersionStatus
const (
	// The Version is being created
	VersionStatusCreating VersionStatus = "CREATING"
	// The Version is updating
	VersionStatusUpdating VersionStatus = "UPDATING"
	// The Version is deleting
	VersionStatusDeleting VersionStatus = "DELETING"
	// The Version is active
	VersionStatusActive VersionStatus = "ACTIVE"
	// The Version creation failed
	VersionStatusFailed VersionStatus = "FAILED"
)

// Values returns all known values for VersionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VersionStatus) Values() []VersionStatus {
	return []VersionStatus{
		"CREATING",
		"UPDATING",
		"DELETING",
		"ACTIVE",
		"FAILED",
	}
}

type WorkflowEngine string

// Enum values for WorkflowEngine
const (
	WorkflowEngineWdl      WorkflowEngine = "WDL"
	WorkflowEngineNextflow WorkflowEngine = "NEXTFLOW"
	WorkflowEngineCwl      WorkflowEngine = "CWL"
)

// Values returns all known values for WorkflowEngine. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowEngine) Values() []WorkflowEngine {
	return []WorkflowEngine{
		"WDL",
		"NEXTFLOW",
		"CWL",
	}
}

type WorkflowExport string

// Enum values for WorkflowExport
const (
	WorkflowExportDefinition WorkflowExport = "DEFINITION"
)

// Values returns all known values for WorkflowExport. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowExport) Values() []WorkflowExport {
	return []WorkflowExport{
		"DEFINITION",
	}
}

type WorkflowStatus string

// Enum values for WorkflowStatus
const (
	WorkflowStatusCreating WorkflowStatus = "CREATING"
	WorkflowStatusActive   WorkflowStatus = "ACTIVE"
	WorkflowStatusUpdating WorkflowStatus = "UPDATING"
	WorkflowStatusDeleted  WorkflowStatus = "DELETED"
	WorkflowStatusFailed   WorkflowStatus = "FAILED"
	WorkflowStatusInactive WorkflowStatus = "INACTIVE"
)

// Values returns all known values for WorkflowStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowStatus) Values() []WorkflowStatus {
	return []WorkflowStatus{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETED",
		"FAILED",
		"INACTIVE",
	}
}

type WorkflowType string

// Enum values for WorkflowType
const (
	WorkflowTypePrivate   WorkflowType = "PRIVATE"
	WorkflowTypeReady2run WorkflowType = "READY2RUN"
)

// Values returns all known values for WorkflowType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowType) Values() []WorkflowType {
	return []WorkflowType{
		"PRIVATE",
		"READY2RUN",
	}
}
