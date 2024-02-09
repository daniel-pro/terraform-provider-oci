// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AnalyzeImageResult The image analysis results.
type AnalyzeImageResult struct {

	// The detected objects.
	ImageObjects []ImageObject `mandatory:"false" json:"imageObjects"`

	// The image classification labels.
	Labels []Label `mandatory:"false" json:"labels"`

	// The ontologyClasses of image labels.
	OntologyClasses []OntologyClass `mandatory:"false" json:"ontologyClasses"`

	ImageText *ImageText `mandatory:"false" json:"imageText"`

	// The detected object proposals.
	ObjectProposals []ObjectProposal `mandatory:"false" json:"objectProposals"`

	// The detected faces.
	DetectedFaces []Face `mandatory:"false" json:"detectedFaces"`

	// The image classification model version.
	ImageClassificationModelVersion *string `mandatory:"false" json:"imageClassificationModelVersion"`

	// The object detection model version.
	ObjectDetectionModelVersion *string `mandatory:"false" json:"objectDetectionModelVersion"`

	// The text detection model version.
	TextDetectionModelVersion *string `mandatory:"false" json:"textDetectionModelVersion"`

	// The object proposal model version.
	ObjectProposalModelVersion *string `mandatory:"false" json:"objectProposalModelVersion"`

	// The face detection model version.
	FaceDetectionModelVersion *string `mandatory:"false" json:"faceDetectionModelVersion"`

	// The errors encountered during image analysis.
	Errors []ProcessingError `mandatory:"false" json:"errors"`
}

func (m AnalyzeImageResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnalyzeImageResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
