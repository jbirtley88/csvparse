package csvparse

import "errors"

var (
	ErrHeaderCountDoesNotMatchRowCount  = errors.New("header and row count must be identical")
	ErrRowCountIsSmallerThanTargetCount = errors.New("no more targets can be given than the available columns")
)
