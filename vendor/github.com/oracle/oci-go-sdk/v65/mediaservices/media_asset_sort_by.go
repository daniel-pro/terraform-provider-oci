// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"strings"
)

// MediaAssetSortByEnum Enum with underlying type: string
type MediaAssetSortByEnum string

// Set of constants representing the allowable values for MediaAssetSortByEnum
const (
	MediaAssetSortByCompartmentId      MediaAssetSortByEnum = "compartmentId"
	MediaAssetSortByType               MediaAssetSortByEnum = "type"
	MediaAssetSortByLifecycleState     MediaAssetSortByEnum = "lifecycleState"
	MediaAssetSortByParentMediaAssetId MediaAssetSortByEnum = "parentMediaAssetId"
	MediaAssetSortByMasterMediaAssetId MediaAssetSortByEnum = "masterMediaAssetId"
	MediaAssetSortByDisplayName        MediaAssetSortByEnum = "displayName"
	MediaAssetSortByTimeCreated        MediaAssetSortByEnum = "timeCreated"
	MediaAssetSortByTimeUpdated        MediaAssetSortByEnum = "timeUpdated"
)

var mappingMediaAssetSortByEnum = map[string]MediaAssetSortByEnum{
	"compartmentId":      MediaAssetSortByCompartmentId,
	"type":               MediaAssetSortByType,
	"lifecycleState":     MediaAssetSortByLifecycleState,
	"parentMediaAssetId": MediaAssetSortByParentMediaAssetId,
	"masterMediaAssetId": MediaAssetSortByMasterMediaAssetId,
	"displayName":        MediaAssetSortByDisplayName,
	"timeCreated":        MediaAssetSortByTimeCreated,
	"timeUpdated":        MediaAssetSortByTimeUpdated,
}

var mappingMediaAssetSortByEnumLowerCase = map[string]MediaAssetSortByEnum{
	"compartmentid":      MediaAssetSortByCompartmentId,
	"type":               MediaAssetSortByType,
	"lifecyclestate":     MediaAssetSortByLifecycleState,
	"parentmediaassetid": MediaAssetSortByParentMediaAssetId,
	"mastermediaassetid": MediaAssetSortByMasterMediaAssetId,
	"displayname":        MediaAssetSortByDisplayName,
	"timecreated":        MediaAssetSortByTimeCreated,
	"timeupdated":        MediaAssetSortByTimeUpdated,
}

// GetMediaAssetSortByEnumValues Enumerates the set of values for MediaAssetSortByEnum
func GetMediaAssetSortByEnumValues() []MediaAssetSortByEnum {
	values := make([]MediaAssetSortByEnum, 0)
	for _, v := range mappingMediaAssetSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetMediaAssetSortByEnumStringValues Enumerates the set of values in String for MediaAssetSortByEnum
func GetMediaAssetSortByEnumStringValues() []string {
	return []string{
		"compartmentId",
		"type",
		"lifecycleState",
		"parentMediaAssetId",
		"masterMediaAssetId",
		"displayName",
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingMediaAssetSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMediaAssetSortByEnum(val string) (MediaAssetSortByEnum, bool) {
	enum, ok := mappingMediaAssetSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
