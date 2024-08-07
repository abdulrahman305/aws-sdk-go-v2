// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CaptureTool string

// Enum values for CaptureTool
const (
	CaptureToolPrecisely CaptureTool = "Precisely"
	CaptureToolAwsDms    CaptureTool = "AWS DMS"
)

// Values returns all known values for CaptureTool. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CaptureTool) Values() []CaptureTool {
	return []CaptureTool{
		"Precisely",
		"AWS DMS",
	}
}

type CloudFormationActionType string

// Enum values for CloudFormationActionType
const (
	CloudFormationActionTypeCreate CloudFormationActionType = "Create"
	CloudFormationActionTypeDelete CloudFormationActionType = "Delete"
)

// Values returns all known values for CloudFormationActionType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CloudFormationActionType) Values() []CloudFormationActionType {
	return []CloudFormationActionType{
		"Create",
		"Delete",
	}
}

type ComparisonStatusEnum string

// Enum values for ComparisonStatusEnum
const (
	ComparisonStatusEnumDifferent  ComparisonStatusEnum = "Different"
	ComparisonStatusEnumEquivalent ComparisonStatusEnum = "Equivalent"
	ComparisonStatusEnumEqual      ComparisonStatusEnum = "Equal"
)

// Values returns all known values for ComparisonStatusEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ComparisonStatusEnum) Values() []ComparisonStatusEnum {
	return []ComparisonStatusEnum{
		"Different",
		"Equivalent",
		"Equal",
	}
}

type DataSetType string

// Enum values for DataSetType
const (
	DataSetTypePs DataSetType = "PS"
)

// Values returns all known values for DataSetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataSetType) Values() []DataSetType {
	return []DataSetType{
		"PS",
	}
}

type Format string

// Enum values for Format
const (
	FormatFixed          Format = "FIXED"
	FormatVariable       Format = "VARIABLE"
	FormatLineSequential Format = "LINE_SEQUENTIAL"
)

// Values returns all known values for Format. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Format) Values() []Format {
	return []Format{
		"FIXED",
		"VARIABLE",
		"LINE_SEQUENTIAL",
	}
}

type M2ManagedActionType string

// Enum values for M2ManagedActionType
const (
	M2ManagedActionTypeConfigure   M2ManagedActionType = "Configure"
	M2ManagedActionTypeDeconfigure M2ManagedActionType = "Deconfigure"
)

// Values returns all known values for M2ManagedActionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (M2ManagedActionType) Values() []M2ManagedActionType {
	return []M2ManagedActionType{
		"Configure",
		"Deconfigure",
	}
}

type M2ManagedRuntime string

// Enum values for M2ManagedRuntime
const (
	M2ManagedRuntimeMicrofocus M2ManagedRuntime = "MicroFocus"
)

// Values returns all known values for M2ManagedRuntime. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (M2ManagedRuntime) Values() []M2ManagedRuntime {
	return []M2ManagedRuntime{
		"MicroFocus",
	}
}

type M2NonManagedActionType string

// Enum values for M2NonManagedActionType
const (
	M2NonManagedActionTypeConfigure   M2NonManagedActionType = "Configure"
	M2NonManagedActionTypeDeconfigure M2NonManagedActionType = "Deconfigure"
)

// Values returns all known values for M2NonManagedActionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (M2NonManagedActionType) Values() []M2NonManagedActionType {
	return []M2NonManagedActionType{
		"Configure",
		"Deconfigure",
	}
}

type M2NonManagedRuntime string

// Enum values for M2NonManagedRuntime
const (
	M2NonManagedRuntimeBluage M2NonManagedRuntime = "BluAge"
)

// Values returns all known values for M2NonManagedRuntime. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (M2NonManagedRuntime) Values() []M2NonManagedRuntime {
	return []M2NonManagedRuntime{
		"BluAge",
	}
}

type ScriptType string

// Enum values for ScriptType
const (
	ScriptTypeSelenium ScriptType = "Selenium"
)

// Values returns all known values for ScriptType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScriptType) Values() []ScriptType {
	return []ScriptType{
		"Selenium",
	}
}

type SourceDatabase string

// Enum values for SourceDatabase
const (
	SourceDatabaseZOsDb2 SourceDatabase = "z/OS-DB2"
)

// Values returns all known values for SourceDatabase. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SourceDatabase) Values() []SourceDatabase {
	return []SourceDatabase{
		"z/OS-DB2",
	}
}

type StepRunStatus string

// Enum values for StepRunStatus
const (
	StepRunStatusSuccess StepRunStatus = "Success"
	StepRunStatusFailed  StepRunStatus = "Failed"
	StepRunStatusRunning StepRunStatus = "Running"
)

// Values returns all known values for StepRunStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StepRunStatus) Values() []StepRunStatus {
	return []StepRunStatus{
		"Success",
		"Failed",
		"Running",
	}
}

type TargetDatabase string

// Enum values for TargetDatabase
const (
	TargetDatabasePostgresql TargetDatabase = "PostgreSQL"
)

// Values returns all known values for TargetDatabase. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetDatabase) Values() []TargetDatabase {
	return []TargetDatabase{
		"PostgreSQL",
	}
}

type TestCaseLifecycle string

// Enum values for TestCaseLifecycle
const (
	TestCaseLifecycleActive   TestCaseLifecycle = "Active"
	TestCaseLifecycleDeleting TestCaseLifecycle = "Deleting"
)

// Values returns all known values for TestCaseLifecycle. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TestCaseLifecycle) Values() []TestCaseLifecycle {
	return []TestCaseLifecycle{
		"Active",
		"Deleting",
	}
}

type TestCaseRunStatus string

// Enum values for TestCaseRunStatus
const (
	TestCaseRunStatusSuccess TestCaseRunStatus = "Success"
	TestCaseRunStatusRunning TestCaseRunStatus = "Running"
	TestCaseRunStatusFailed  TestCaseRunStatus = "Failed"
)

// Values returns all known values for TestCaseRunStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TestCaseRunStatus) Values() []TestCaseRunStatus {
	return []TestCaseRunStatus{
		"Success",
		"Running",
		"Failed",
	}
}

type TestConfigurationLifecycle string

// Enum values for TestConfigurationLifecycle
const (
	TestConfigurationLifecycleActive   TestConfigurationLifecycle = "Active"
	TestConfigurationLifecycleDeleting TestConfigurationLifecycle = "Deleting"
)

// Values returns all known values for TestConfigurationLifecycle. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TestConfigurationLifecycle) Values() []TestConfigurationLifecycle {
	return []TestConfigurationLifecycle{
		"Active",
		"Deleting",
	}
}

type TestRunStatus string

// Enum values for TestRunStatus
const (
	TestRunStatusSuccess  TestRunStatus = "Success"
	TestRunStatusRunning  TestRunStatus = "Running"
	TestRunStatusFailed   TestRunStatus = "Failed"
	TestRunStatusDeleting TestRunStatus = "Deleting"
)

// Values returns all known values for TestRunStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TestRunStatus) Values() []TestRunStatus {
	return []TestRunStatus{
		"Success",
		"Running",
		"Failed",
		"Deleting",
	}
}

type TestSuiteLifecycle string

// Enum values for TestSuiteLifecycle
const (
	TestSuiteLifecycleCreating TestSuiteLifecycle = "Creating"
	TestSuiteLifecycleUpdating TestSuiteLifecycle = "Updating"
	TestSuiteLifecycleActive   TestSuiteLifecycle = "Active"
	TestSuiteLifecycleFailed   TestSuiteLifecycle = "Failed"
	TestSuiteLifecycleDeleting TestSuiteLifecycle = "Deleting"
)

// Values returns all known values for TestSuiteLifecycle. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TestSuiteLifecycle) Values() []TestSuiteLifecycle {
	return []TestSuiteLifecycle{
		"Creating",
		"Updating",
		"Active",
		"Failed",
		"Deleting",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "other"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"unknownOperation",
		"cannotParse",
		"fieldValidationFailed",
		"other",
	}
}
