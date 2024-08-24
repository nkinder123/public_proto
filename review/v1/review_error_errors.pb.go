// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsOrderHasReview(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ORDER_HAS_REVIEW.String() && e.Code == 301
}

func ErrorOrderHasReview(format string, args ...interface{}) *errors.Error {
	return errors.New(301, ErrorReason_ORDER_HAS_REVIEW.String(), fmt.Sprintf(format, args...))
}

func IsCreateReviewHasError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CREATE_REVIEW_HAS_ERROR.String() && e.Code == 302
}

func ErrorCreateReviewHasError(format string, args ...interface{}) *errors.Error {
	return errors.New(302, ErrorReason_CREATE_REVIEW_HAS_ERROR.String(), fmt.Sprintf(format, args...))
}
