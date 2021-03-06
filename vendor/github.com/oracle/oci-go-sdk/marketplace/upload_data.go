// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UploadData The model for upload data for images and icons.
type UploadData struct {

	// The name used to refer to the upload data.
	Name *string `mandatory:"false" json:"name"`

	// The content URL of the upload data.
	ContentUrl *string `mandatory:"false" json:"contentUrl"`

	// The MIME type of the upload data.
	MimeType *string `mandatory:"false" json:"mimeType"`

	// The file extension of the upload data.
	FileExtension *string `mandatory:"false" json:"fileExtension"`
}

func (m UploadData) String() string {
	return common.PointerString(m)
}
