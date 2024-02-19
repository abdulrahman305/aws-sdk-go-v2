// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CertificateType string

// Enum values for CertificateType
const (
	CertificateTypeAmplifyManaged CertificateType = "AMPLIFY_MANAGED"
	CertificateTypeCustom         CertificateType = "CUSTOM"
)

// Values returns all known values for CertificateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CertificateType) Values() []CertificateType {
	return []CertificateType{
		"AMPLIFY_MANAGED",
		"CUSTOM",
	}
}

type DomainStatus string

// Enum values for DomainStatus
const (
	DomainStatusPendingVerification        DomainStatus = "PENDING_VERIFICATION"
	DomainStatusInProgress                 DomainStatus = "IN_PROGRESS"
	DomainStatusAvailable                  DomainStatus = "AVAILABLE"
	DomainStatusImportingCustomCertificate DomainStatus = "IMPORTING_CUSTOM_CERTIFICATE"
	DomainStatusPendingDeployment          DomainStatus = "PENDING_DEPLOYMENT"
	DomainStatusAwaitingAppCname           DomainStatus = "AWAITING_APP_CNAME"
	DomainStatusFailed                     DomainStatus = "FAILED"
	DomainStatusCreating                   DomainStatus = "CREATING"
	DomainStatusRequestingCertificate      DomainStatus = "REQUESTING_CERTIFICATE"
	DomainStatusUpdating                   DomainStatus = "UPDATING"
)

// Values returns all known values for DomainStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DomainStatus) Values() []DomainStatus {
	return []DomainStatus{
		"PENDING_VERIFICATION",
		"IN_PROGRESS",
		"AVAILABLE",
		"IMPORTING_CUSTOM_CERTIFICATE",
		"PENDING_DEPLOYMENT",
		"AWAITING_APP_CNAME",
		"FAILED",
		"CREATING",
		"REQUESTING_CERTIFICATE",
		"UPDATING",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusPending      JobStatus = "PENDING"
	JobStatusProvisioning JobStatus = "PROVISIONING"
	JobStatusRunning      JobStatus = "RUNNING"
	JobStatusFailed       JobStatus = "FAILED"
	JobStatusSucceed      JobStatus = "SUCCEED"
	JobStatusCancelling   JobStatus = "CANCELLING"
	JobStatusCancelled    JobStatus = "CANCELLED"
)

// Values returns all known values for JobStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"PENDING",
		"PROVISIONING",
		"RUNNING",
		"FAILED",
		"SUCCEED",
		"CANCELLING",
		"CANCELLED",
	}
}

type JobType string

// Enum values for JobType
const (
	JobTypeRelease JobType = "RELEASE"
	JobTypeRetry   JobType = "RETRY"
	JobTypeManual  JobType = "MANUAL"
	JobTypeWebHook JobType = "WEB_HOOK"
)

// Values returns all known values for JobType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobType) Values() []JobType {
	return []JobType{
		"RELEASE",
		"RETRY",
		"MANUAL",
		"WEB_HOOK",
	}
}

type Platform string

// Enum values for Platform
const (
	PlatformWeb        Platform = "WEB"
	PlatformWebDynamic Platform = "WEB_DYNAMIC"
	PlatformWebCompute Platform = "WEB_COMPUTE"
)

// Values returns all known values for Platform. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Platform) Values() []Platform {
	return []Platform{
		"WEB",
		"WEB_DYNAMIC",
		"WEB_COMPUTE",
	}
}

type RepositoryCloneMethod string

// Enum values for RepositoryCloneMethod
const (
	RepositoryCloneMethodSsh   RepositoryCloneMethod = "SSH"
	RepositoryCloneMethodToken RepositoryCloneMethod = "TOKEN"
	RepositoryCloneMethodSigv4 RepositoryCloneMethod = "SIGV4"
)

// Values returns all known values for RepositoryCloneMethod. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RepositoryCloneMethod) Values() []RepositoryCloneMethod {
	return []RepositoryCloneMethod{
		"SSH",
		"TOKEN",
		"SIGV4",
	}
}

type Stage string

// Enum values for Stage
const (
	StageProduction   Stage = "PRODUCTION"
	StageBeta         Stage = "BETA"
	StageDevelopment  Stage = "DEVELOPMENT"
	StageExperimental Stage = "EXPERIMENTAL"
	StagePullRequest  Stage = "PULL_REQUEST"
)

// Values returns all known values for Stage. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Stage) Values() []Stage {
	return []Stage{
		"PRODUCTION",
		"BETA",
		"DEVELOPMENT",
		"EXPERIMENTAL",
		"PULL_REQUEST",
	}
}

type UpdateStatus string

// Enum values for UpdateStatus
const (
	UpdateStatusRequestingCertificate      UpdateStatus = "REQUESTING_CERTIFICATE"
	UpdateStatusPendingVerification        UpdateStatus = "PENDING_VERIFICATION"
	UpdateStatusImportingCustomCertificate UpdateStatus = "IMPORTING_CUSTOM_CERTIFICATE"
	UpdateStatusPendingDeployment          UpdateStatus = "PENDING_DEPLOYMENT"
	UpdateStatusAwaitingAppCname           UpdateStatus = "AWAITING_APP_CNAME"
	UpdateStatusUpdateComplete             UpdateStatus = "UPDATE_COMPLETE"
	UpdateStatusUpdateFailed               UpdateStatus = "UPDATE_FAILED"
)

// Values returns all known values for UpdateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UpdateStatus) Values() []UpdateStatus {
	return []UpdateStatus{
		"REQUESTING_CERTIFICATE",
		"PENDING_VERIFICATION",
		"IMPORTING_CUSTOM_CERTIFICATE",
		"PENDING_DEPLOYMENT",
		"AWAITING_APP_CNAME",
		"UPDATE_COMPLETE",
		"UPDATE_FAILED",
	}
}
