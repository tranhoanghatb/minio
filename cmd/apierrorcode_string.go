// Code generated by "stringer -type=APIErrorCode -trimprefix=Err api-errors.go"; DO NOT EDIT.

package cmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrNone-0]
	_ = x[ErrAccessDenied-1]
	_ = x[ErrBadDigest-2]
	_ = x[ErrEntityTooSmall-3]
	_ = x[ErrEntityTooLarge-4]
	_ = x[ErrPolicyTooLarge-5]
	_ = x[ErrIncompleteBody-6]
	_ = x[ErrInternalError-7]
	_ = x[ErrInvalidAccessKeyID-8]
	_ = x[ErrAccessKeyDisabled-9]
	_ = x[ErrInvalidBucketName-10]
	_ = x[ErrInvalidDigest-11]
	_ = x[ErrInvalidRange-12]
	_ = x[ErrInvalidRangePartNumber-13]
	_ = x[ErrInvalidCopyPartRange-14]
	_ = x[ErrInvalidCopyPartRangeSource-15]
	_ = x[ErrInvalidMaxKeys-16]
	_ = x[ErrInvalidEncodingMethod-17]
	_ = x[ErrInvalidMaxUploads-18]
	_ = x[ErrInvalidMaxParts-19]
	_ = x[ErrInvalidPartNumberMarker-20]
	_ = x[ErrInvalidPartNumber-21]
	_ = x[ErrInvalidRequestBody-22]
	_ = x[ErrInvalidCopySource-23]
	_ = x[ErrInvalidMetadataDirective-24]
	_ = x[ErrInvalidCopyDest-25]
	_ = x[ErrInvalidPolicyDocument-26]
	_ = x[ErrInvalidObjectState-27]
	_ = x[ErrMalformedXML-28]
	_ = x[ErrMissingContentLength-29]
	_ = x[ErrMissingContentMD5-30]
	_ = x[ErrMissingRequestBodyError-31]
	_ = x[ErrMissingSecurityHeader-32]
	_ = x[ErrNoSuchBucket-33]
	_ = x[ErrNoSuchBucketPolicy-34]
	_ = x[ErrNoSuchBucketLifecycle-35]
	_ = x[ErrNoSuchLifecycleConfiguration-36]
	_ = x[ErrInvalidLifecycleWithObjectLock-37]
	_ = x[ErrNoSuchBucketSSEConfig-38]
	_ = x[ErrNoSuchCORSConfiguration-39]
	_ = x[ErrNoSuchWebsiteConfiguration-40]
	_ = x[ErrReplicationConfigurationNotFoundError-41]
	_ = x[ErrRemoteDestinationNotFoundError-42]
	_ = x[ErrReplicationDestinationMissingLock-43]
	_ = x[ErrRemoteTargetNotFoundError-44]
	_ = x[ErrReplicationRemoteConnectionError-45]
	_ = x[ErrReplicationBandwidthLimitError-46]
	_ = x[ErrBucketRemoteIdenticalToSource-47]
	_ = x[ErrBucketRemoteAlreadyExists-48]
	_ = x[ErrBucketRemoteLabelInUse-49]
	_ = x[ErrBucketRemoteArnTypeInvalid-50]
	_ = x[ErrBucketRemoteArnInvalid-51]
	_ = x[ErrBucketRemoteRemoveDisallowed-52]
	_ = x[ErrRemoteTargetNotVersionedError-53]
	_ = x[ErrReplicationSourceNotVersionedError-54]
	_ = x[ErrReplicationNeedsVersioningError-55]
	_ = x[ErrReplicationBucketNeedsVersioningError-56]
	_ = x[ErrReplicationDenyEditError-57]
	_ = x[ErrReplicationNoExistingObjects-58]
	_ = x[ErrObjectRestoreAlreadyInProgress-59]
	_ = x[ErrNoSuchKey-60]
	_ = x[ErrNoSuchUpload-61]
	_ = x[ErrInvalidVersionID-62]
	_ = x[ErrNoSuchVersion-63]
	_ = x[ErrNotImplemented-64]
	_ = x[ErrPreconditionFailed-65]
	_ = x[ErrRequestTimeTooSkewed-66]
	_ = x[ErrSignatureDoesNotMatch-67]
	_ = x[ErrMethodNotAllowed-68]
	_ = x[ErrInvalidPart-69]
	_ = x[ErrInvalidPartOrder-70]
	_ = x[ErrAuthorizationHeaderMalformed-71]
	_ = x[ErrMalformedPOSTRequest-72]
	_ = x[ErrPOSTFileRequired-73]
	_ = x[ErrSignatureVersionNotSupported-74]
	_ = x[ErrBucketNotEmpty-75]
	_ = x[ErrAllAccessDisabled-76]
	_ = x[ErrMalformedPolicy-77]
	_ = x[ErrMissingFields-78]
	_ = x[ErrMissingCredTag-79]
	_ = x[ErrCredMalformed-80]
	_ = x[ErrInvalidRegion-81]
	_ = x[ErrInvalidServiceS3-82]
	_ = x[ErrInvalidServiceSTS-83]
	_ = x[ErrInvalidRequestVersion-84]
	_ = x[ErrMissingSignTag-85]
	_ = x[ErrMissingSignHeadersTag-86]
	_ = x[ErrMalformedDate-87]
	_ = x[ErrMalformedPresignedDate-88]
	_ = x[ErrMalformedCredentialDate-89]
	_ = x[ErrMalformedCredentialRegion-90]
	_ = x[ErrMalformedExpires-91]
	_ = x[ErrNegativeExpires-92]
	_ = x[ErrAuthHeaderEmpty-93]
	_ = x[ErrExpiredPresignRequest-94]
	_ = x[ErrRequestNotReadyYet-95]
	_ = x[ErrUnsignedHeaders-96]
	_ = x[ErrMissingDateHeader-97]
	_ = x[ErrInvalidQuerySignatureAlgo-98]
	_ = x[ErrInvalidQueryParams-99]
	_ = x[ErrBucketAlreadyOwnedByYou-100]
	_ = x[ErrInvalidDuration-101]
	_ = x[ErrBucketAlreadyExists-102]
	_ = x[ErrTooManyBuckets-103]
	_ = x[ErrMetadataTooLarge-104]
	_ = x[ErrUnsupportedMetadata-105]
	_ = x[ErrMaximumExpires-106]
	_ = x[ErrSlowDown-107]
	_ = x[ErrInvalidPrefixMarker-108]
	_ = x[ErrBadRequest-109]
	_ = x[ErrKeyTooLongError-110]
	_ = x[ErrInvalidBucketObjectLockConfiguration-111]
	_ = x[ErrObjectLockConfigurationNotFound-112]
	_ = x[ErrObjectLockConfigurationNotAllowed-113]
	_ = x[ErrNoSuchObjectLockConfiguration-114]
	_ = x[ErrObjectLocked-115]
	_ = x[ErrInvalidRetentionDate-116]
	_ = x[ErrPastObjectLockRetainDate-117]
	_ = x[ErrUnknownWORMModeDirective-118]
	_ = x[ErrBucketTaggingNotFound-119]
	_ = x[ErrObjectLockInvalidHeaders-120]
	_ = x[ErrInvalidTagDirective-121]
	_ = x[ErrInvalidEncryptionMethod-122]
	_ = x[ErrInvalidEncryptionKeyID-123]
	_ = x[ErrInsecureSSECustomerRequest-124]
	_ = x[ErrSSEMultipartEncrypted-125]
	_ = x[ErrSSEEncryptedObject-126]
	_ = x[ErrInvalidEncryptionParameters-127]
	_ = x[ErrInvalidSSECustomerAlgorithm-128]
	_ = x[ErrInvalidSSECustomerKey-129]
	_ = x[ErrMissingSSECustomerKey-130]
	_ = x[ErrMissingSSECustomerKeyMD5-131]
	_ = x[ErrSSECustomerKeyMD5Mismatch-132]
	_ = x[ErrInvalidSSECustomerParameters-133]
	_ = x[ErrIncompatibleEncryptionMethod-134]
	_ = x[ErrKMSNotConfigured-135]
	_ = x[ErrKMSKeyNotFoundException-136]
	_ = x[ErrNoAccessKey-137]
	_ = x[ErrInvalidToken-138]
	_ = x[ErrEventNotification-139]
	_ = x[ErrARNNotification-140]
	_ = x[ErrRegionNotification-141]
	_ = x[ErrOverlappingFilterNotification-142]
	_ = x[ErrFilterNameInvalid-143]
	_ = x[ErrFilterNamePrefix-144]
	_ = x[ErrFilterNameSuffix-145]
	_ = x[ErrFilterValueInvalid-146]
	_ = x[ErrOverlappingConfigs-147]
	_ = x[ErrUnsupportedNotification-148]
	_ = x[ErrContentSHA256Mismatch-149]
	_ = x[ErrContentChecksumMismatch-150]
	_ = x[ErrReadQuorum-151]
	_ = x[ErrWriteQuorum-152]
	_ = x[ErrStorageFull-153]
	_ = x[ErrRequestBodyParse-154]
	_ = x[ErrObjectExistsAsDirectory-155]
	_ = x[ErrInvalidObjectName-156]
	_ = x[ErrInvalidObjectNamePrefixSlash-157]
	_ = x[ErrInvalidResourceName-158]
	_ = x[ErrServerNotInitialized-159]
	_ = x[ErrOperationTimedOut-160]
	_ = x[ErrClientDisconnected-161]
	_ = x[ErrOperationMaxedOut-162]
	_ = x[ErrInvalidRequest-163]
	_ = x[ErrTransitionStorageClassNotFoundError-164]
	_ = x[ErrInvalidStorageClass-165]
	_ = x[ErrBackendDown-166]
	_ = x[ErrMalformedJSON-167]
	_ = x[ErrAdminNoSuchUser-168]
	_ = x[ErrAdminNoSuchGroup-169]
	_ = x[ErrAdminGroupNotEmpty-170]
	_ = x[ErrAdminNoSuchJob-171]
	_ = x[ErrAdminNoSuchPolicy-172]
	_ = x[ErrAdminInvalidArgument-173]
	_ = x[ErrAdminInvalidAccessKey-174]
	_ = x[ErrAdminInvalidSecretKey-175]
	_ = x[ErrAdminConfigNoQuorum-176]
	_ = x[ErrAdminConfigTooLarge-177]
	_ = x[ErrAdminConfigBadJSON-178]
	_ = x[ErrAdminNoSuchConfigTarget-179]
	_ = x[ErrAdminConfigEnvOverridden-180]
	_ = x[ErrAdminConfigDuplicateKeys-181]
	_ = x[ErrAdminCredentialsMismatch-182]
	_ = x[ErrInsecureClientRequest-183]
	_ = x[ErrObjectTampered-184]
	_ = x[ErrSiteReplicationInvalidRequest-185]
	_ = x[ErrSiteReplicationPeerResp-186]
	_ = x[ErrSiteReplicationBackendIssue-187]
	_ = x[ErrSiteReplicationServiceAccountError-188]
	_ = x[ErrSiteReplicationBucketConfigError-189]
	_ = x[ErrSiteReplicationBucketMetaError-190]
	_ = x[ErrSiteReplicationIAMError-191]
	_ = x[ErrSiteReplicationConfigMissing-192]
	_ = x[ErrAdminBucketQuotaExceeded-193]
	_ = x[ErrAdminNoSuchQuotaConfiguration-194]
	_ = x[ErrHealNotImplemented-195]
	_ = x[ErrHealNoSuchProcess-196]
	_ = x[ErrHealInvalidClientToken-197]
	_ = x[ErrHealMissingBucket-198]
	_ = x[ErrHealAlreadyRunning-199]
	_ = x[ErrHealOverlappingPaths-200]
	_ = x[ErrIncorrectContinuationToken-201]
	_ = x[ErrEmptyRequestBody-202]
	_ = x[ErrUnsupportedFunction-203]
	_ = x[ErrInvalidExpressionType-204]
	_ = x[ErrBusy-205]
	_ = x[ErrUnauthorizedAccess-206]
	_ = x[ErrExpressionTooLong-207]
	_ = x[ErrIllegalSQLFunctionArgument-208]
	_ = x[ErrInvalidKeyPath-209]
	_ = x[ErrInvalidCompressionFormat-210]
	_ = x[ErrInvalidFileHeaderInfo-211]
	_ = x[ErrInvalidJSONType-212]
	_ = x[ErrInvalidQuoteFields-213]
	_ = x[ErrInvalidRequestParameter-214]
	_ = x[ErrInvalidDataType-215]
	_ = x[ErrInvalidTextEncoding-216]
	_ = x[ErrInvalidDataSource-217]
	_ = x[ErrInvalidTableAlias-218]
	_ = x[ErrMissingRequiredParameter-219]
	_ = x[ErrObjectSerializationConflict-220]
	_ = x[ErrUnsupportedSQLOperation-221]
	_ = x[ErrUnsupportedSQLStructure-222]
	_ = x[ErrUnsupportedSyntax-223]
	_ = x[ErrUnsupportedRangeHeader-224]
	_ = x[ErrLexerInvalidChar-225]
	_ = x[ErrLexerInvalidOperator-226]
	_ = x[ErrLexerInvalidLiteral-227]
	_ = x[ErrLexerInvalidIONLiteral-228]
	_ = x[ErrParseExpectedDatePart-229]
	_ = x[ErrParseExpectedKeyword-230]
	_ = x[ErrParseExpectedTokenType-231]
	_ = x[ErrParseExpected2TokenTypes-232]
	_ = x[ErrParseExpectedNumber-233]
	_ = x[ErrParseExpectedRightParenBuiltinFunctionCall-234]
	_ = x[ErrParseExpectedTypeName-235]
	_ = x[ErrParseExpectedWhenClause-236]
	_ = x[ErrParseUnsupportedToken-237]
	_ = x[ErrParseUnsupportedLiteralsGroupBy-238]
	_ = x[ErrParseExpectedMember-239]
	_ = x[ErrParseUnsupportedSelect-240]
	_ = x[ErrParseUnsupportedCase-241]
	_ = x[ErrParseUnsupportedCaseClause-242]
	_ = x[ErrParseUnsupportedAlias-243]
	_ = x[ErrParseUnsupportedSyntax-244]
	_ = x[ErrParseUnknownOperator-245]
	_ = x[ErrParseMissingIdentAfterAt-246]
	_ = x[ErrParseUnexpectedOperator-247]
	_ = x[ErrParseUnexpectedTerm-248]
	_ = x[ErrParseUnexpectedToken-249]
	_ = x[ErrParseUnexpectedKeyword-250]
	_ = x[ErrParseExpectedExpression-251]
	_ = x[ErrParseExpectedLeftParenAfterCast-252]
	_ = x[ErrParseExpectedLeftParenValueConstructor-253]
	_ = x[ErrParseExpectedLeftParenBuiltinFunctionCall-254]
	_ = x[ErrParseExpectedArgumentDelimiter-255]
	_ = x[ErrParseCastArity-256]
	_ = x[ErrParseInvalidTypeParam-257]
	_ = x[ErrParseEmptySelect-258]
	_ = x[ErrParseSelectMissingFrom-259]
	_ = x[ErrParseExpectedIdentForGroupName-260]
	_ = x[ErrParseExpectedIdentForAlias-261]
	_ = x[ErrParseUnsupportedCallWithStar-262]
	_ = x[ErrParseNonUnaryAgregateFunctionCall-263]
	_ = x[ErrParseMalformedJoin-264]
	_ = x[ErrParseExpectedIdentForAt-265]
	_ = x[ErrParseAsteriskIsNotAloneInSelectList-266]
	_ = x[ErrParseCannotMixSqbAndWildcardInSelectList-267]
	_ = x[ErrParseInvalidContextForWildcardInSelectList-268]
	_ = x[ErrIncorrectSQLFunctionArgumentType-269]
	_ = x[ErrValueParseFailure-270]
	_ = x[ErrEvaluatorInvalidArguments-271]
	_ = x[ErrIntegerOverflow-272]
	_ = x[ErrLikeInvalidInputs-273]
	_ = x[ErrCastFailed-274]
	_ = x[ErrInvalidCast-275]
	_ = x[ErrEvaluatorInvalidTimestampFormatPattern-276]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbolForParsing-277]
	_ = x[ErrEvaluatorTimestampFormatPatternDuplicateFields-278]
	_ = x[ErrEvaluatorTimestampFormatPatternHourClockAmPmMismatch-279]
	_ = x[ErrEvaluatorUnterminatedTimestampFormatPatternToken-280]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternToken-281]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbol-282]
	_ = x[ErrEvaluatorBindingDoesNotExist-283]
	_ = x[ErrMissingHeaders-284]
	_ = x[ErrInvalidColumnIndex-285]
	_ = x[ErrAdminConfigNotificationTargetsFailed-286]
	_ = x[ErrAdminProfilerNotEnabled-287]
	_ = x[ErrInvalidDecompressedSize-288]
	_ = x[ErrAddUserInvalidArgument-289]
	_ = x[ErrAdminResourceInvalidArgument-290]
	_ = x[ErrAdminAccountNotEligible-291]
	_ = x[ErrAccountNotEligible-292]
	_ = x[ErrAdminServiceAccountNotFound-293]
	_ = x[ErrPostPolicyConditionInvalidFormat-294]
	_ = x[ErrInvalidChecksum-295]
	_ = x[ErrPanFSBucketPahtNotFound-296]
}

const _APIErrorCode_name = "NoneAccessDeniedBadDigestEntityTooSmallEntityTooLargePolicyTooLargeIncompleteBodyInternalErrorInvalidAccessKeyIDAccessKeyDisabledInvalidBucketNameInvalidDigestInvalidRangeInvalidRangePartNumberInvalidCopyPartRangeInvalidCopyPartRangeSourceInvalidMaxKeysInvalidEncodingMethodInvalidMaxUploadsInvalidMaxPartsInvalidPartNumberMarkerInvalidPartNumberInvalidRequestBodyInvalidCopySourceInvalidMetadataDirectiveInvalidCopyDestInvalidPolicyDocumentInvalidObjectStateMalformedXMLMissingContentLengthMissingContentMD5MissingRequestBodyErrorMissingSecurityHeaderNoSuchBucketNoSuchBucketPolicyNoSuchBucketLifecycleNoSuchLifecycleConfigurationInvalidLifecycleWithObjectLockNoSuchBucketSSEConfigNoSuchCORSConfigurationNoSuchWebsiteConfigurationReplicationConfigurationNotFoundErrorRemoteDestinationNotFoundErrorReplicationDestinationMissingLockRemoteTargetNotFoundErrorReplicationRemoteConnectionErrorReplicationBandwidthLimitErrorBucketRemoteIdenticalToSourceBucketRemoteAlreadyExistsBucketRemoteLabelInUseBucketRemoteArnTypeInvalidBucketRemoteArnInvalidBucketRemoteRemoveDisallowedRemoteTargetNotVersionedErrorReplicationSourceNotVersionedErrorReplicationNeedsVersioningErrorReplicationBucketNeedsVersioningErrorReplicationDenyEditErrorReplicationNoExistingObjectsObjectRestoreAlreadyInProgressNoSuchKeyNoSuchUploadInvalidVersionIDNoSuchVersionNotImplementedPreconditionFailedRequestTimeTooSkewedSignatureDoesNotMatchMethodNotAllowedInvalidPartInvalidPartOrderAuthorizationHeaderMalformedMalformedPOSTRequestPOSTFileRequiredSignatureVersionNotSupportedBucketNotEmptyAllAccessDisabledMalformedPolicyMissingFieldsMissingCredTagCredMalformedInvalidRegionInvalidServiceS3InvalidServiceSTSInvalidRequestVersionMissingSignTagMissingSignHeadersTagMalformedDateMalformedPresignedDateMalformedCredentialDateMalformedCredentialRegionMalformedExpiresNegativeExpiresAuthHeaderEmptyExpiredPresignRequestRequestNotReadyYetUnsignedHeadersMissingDateHeaderInvalidQuerySignatureAlgoInvalidQueryParamsBucketAlreadyOwnedByYouInvalidDurationBucketAlreadyExistsTooManyBucketsMetadataTooLargeUnsupportedMetadataMaximumExpiresSlowDownInvalidPrefixMarkerBadRequestKeyTooLongErrorInvalidBucketObjectLockConfigurationObjectLockConfigurationNotFoundObjectLockConfigurationNotAllowedNoSuchObjectLockConfigurationObjectLockedInvalidRetentionDatePastObjectLockRetainDateUnknownWORMModeDirectiveBucketTaggingNotFoundObjectLockInvalidHeadersInvalidTagDirectiveInvalidEncryptionMethodInvalidEncryptionKeyIDInsecureSSECustomerRequestSSEMultipartEncryptedSSEEncryptedObjectInvalidEncryptionParametersInvalidSSECustomerAlgorithmInvalidSSECustomerKeyMissingSSECustomerKeyMissingSSECustomerKeyMD5SSECustomerKeyMD5MismatchInvalidSSECustomerParametersIncompatibleEncryptionMethodKMSNotConfiguredKMSKeyNotFoundExceptionNoAccessKeyInvalidTokenEventNotificationARNNotificationRegionNotificationOverlappingFilterNotificationFilterNameInvalidFilterNamePrefixFilterNameSuffixFilterValueInvalidOverlappingConfigsUnsupportedNotificationContentSHA256MismatchContentChecksumMismatchReadQuorumWriteQuorumStorageFullRequestBodyParseObjectExistsAsDirectoryInvalidObjectNameInvalidObjectNamePrefixSlashInvalidResourceNameServerNotInitializedOperationTimedOutClientDisconnectedOperationMaxedOutInvalidRequestTransitionStorageClassNotFoundErrorInvalidStorageClassBackendDownMalformedJSONAdminNoSuchUserAdminNoSuchGroupAdminGroupNotEmptyAdminNoSuchJobAdminNoSuchPolicyAdminInvalidArgumentAdminInvalidAccessKeyAdminInvalidSecretKeyAdminConfigNoQuorumAdminConfigTooLargeAdminConfigBadJSONAdminNoSuchConfigTargetAdminConfigEnvOverriddenAdminConfigDuplicateKeysAdminCredentialsMismatchInsecureClientRequestObjectTamperedSiteReplicationInvalidRequestSiteReplicationPeerRespSiteReplicationBackendIssueSiteReplicationServiceAccountErrorSiteReplicationBucketConfigErrorSiteReplicationBucketMetaErrorSiteReplicationIAMErrorSiteReplicationConfigMissingAdminBucketQuotaExceededAdminNoSuchQuotaConfigurationHealNotImplementedHealNoSuchProcessHealInvalidClientTokenHealMissingBucketHealAlreadyRunningHealOverlappingPathsIncorrectContinuationTokenEmptyRequestBodyUnsupportedFunctionInvalidExpressionTypeBusyUnauthorizedAccessExpressionTooLongIllegalSQLFunctionArgumentInvalidKeyPathInvalidCompressionFormatInvalidFileHeaderInfoInvalidJSONTypeInvalidQuoteFieldsInvalidRequestParameterInvalidDataTypeInvalidTextEncodingInvalidDataSourceInvalidTableAliasMissingRequiredParameterObjectSerializationConflictUnsupportedSQLOperationUnsupportedSQLStructureUnsupportedSyntaxUnsupportedRangeHeaderLexerInvalidCharLexerInvalidOperatorLexerInvalidLiteralLexerInvalidIONLiteralParseExpectedDatePartParseExpectedKeywordParseExpectedTokenTypeParseExpected2TokenTypesParseExpectedNumberParseExpectedRightParenBuiltinFunctionCallParseExpectedTypeNameParseExpectedWhenClauseParseUnsupportedTokenParseUnsupportedLiteralsGroupByParseExpectedMemberParseUnsupportedSelectParseUnsupportedCaseParseUnsupportedCaseClauseParseUnsupportedAliasParseUnsupportedSyntaxParseUnknownOperatorParseMissingIdentAfterAtParseUnexpectedOperatorParseUnexpectedTermParseUnexpectedTokenParseUnexpectedKeywordParseExpectedExpressionParseExpectedLeftParenAfterCastParseExpectedLeftParenValueConstructorParseExpectedLeftParenBuiltinFunctionCallParseExpectedArgumentDelimiterParseCastArityParseInvalidTypeParamParseEmptySelectParseSelectMissingFromParseExpectedIdentForGroupNameParseExpectedIdentForAliasParseUnsupportedCallWithStarParseNonUnaryAgregateFunctionCallParseMalformedJoinParseExpectedIdentForAtParseAsteriskIsNotAloneInSelectListParseCannotMixSqbAndWildcardInSelectListParseInvalidContextForWildcardInSelectListIncorrectSQLFunctionArgumentTypeValueParseFailureEvaluatorInvalidArgumentsIntegerOverflowLikeInvalidInputsCastFailedInvalidCastEvaluatorInvalidTimestampFormatPatternEvaluatorInvalidTimestampFormatPatternSymbolForParsingEvaluatorTimestampFormatPatternDuplicateFieldsEvaluatorTimestampFormatPatternHourClockAmPmMismatchEvaluatorUnterminatedTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternSymbolEvaluatorBindingDoesNotExistMissingHeadersInvalidColumnIndexAdminConfigNotificationTargetsFailedAdminProfilerNotEnabledInvalidDecompressedSizeAddUserInvalidArgumentAdminResourceInvalidArgumentAdminAccountNotEligibleAccountNotEligibleAdminServiceAccountNotFoundPostPolicyConditionInvalidFormatInvalidChecksumPanFSBucketPahtNotFound"

var _APIErrorCode_index = [...]uint16{0, 4, 16, 25, 39, 53, 67, 81, 94, 112, 129, 146, 159, 171, 193, 213, 239, 253, 274, 291, 306, 329, 346, 364, 381, 405, 420, 441, 459, 471, 491, 508, 531, 552, 564, 582, 603, 631, 661, 682, 705, 731, 768, 798, 831, 856, 888, 918, 947, 972, 994, 1020, 1042, 1070, 1099, 1133, 1164, 1201, 1225, 1253, 1283, 1292, 1304, 1320, 1333, 1347, 1365, 1385, 1406, 1422, 1433, 1449, 1477, 1497, 1513, 1541, 1555, 1572, 1587, 1600, 1614, 1627, 1640, 1656, 1673, 1694, 1708, 1729, 1742, 1764, 1787, 1812, 1828, 1843, 1858, 1879, 1897, 1912, 1929, 1954, 1972, 1995, 2010, 2029, 2043, 2059, 2078, 2092, 2100, 2119, 2129, 2144, 2180, 2211, 2244, 2273, 2285, 2305, 2329, 2353, 2374, 2398, 2417, 2440, 2462, 2488, 2509, 2527, 2554, 2581, 2602, 2623, 2647, 2672, 2700, 2728, 2744, 2767, 2778, 2790, 2807, 2822, 2840, 2869, 2886, 2902, 2918, 2936, 2954, 2977, 2998, 3021, 3031, 3042, 3053, 3069, 3092, 3109, 3137, 3156, 3176, 3193, 3211, 3228, 3242, 3277, 3296, 3307, 3320, 3335, 3351, 3369, 3383, 3400, 3420, 3441, 3462, 3481, 3500, 3518, 3541, 3565, 3589, 3613, 3634, 3648, 3677, 3700, 3727, 3761, 3793, 3823, 3846, 3874, 3898, 3927, 3945, 3962, 3984, 4001, 4019, 4039, 4065, 4081, 4100, 4121, 4125, 4143, 4160, 4186, 4200, 4224, 4245, 4260, 4278, 4301, 4316, 4335, 4352, 4369, 4393, 4420, 4443, 4466, 4483, 4505, 4521, 4541, 4560, 4582, 4603, 4623, 4645, 4669, 4688, 4730, 4751, 4774, 4795, 4826, 4845, 4867, 4887, 4913, 4934, 4956, 4976, 5000, 5023, 5042, 5062, 5084, 5107, 5138, 5176, 5217, 5247, 5261, 5282, 5298, 5320, 5350, 5376, 5404, 5437, 5455, 5478, 5513, 5553, 5595, 5627, 5644, 5669, 5684, 5701, 5711, 5722, 5760, 5814, 5860, 5912, 5960, 6003, 6047, 6075, 6089, 6107, 6143, 6166, 6189, 6211, 6239, 6262, 6280, 6307, 6339, 6354, 6377}

func (i APIErrorCode) String() string {
	if i < 0 || i >= APIErrorCode(len(_APIErrorCode_index)-1) {
		return "APIErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _APIErrorCode_name[_APIErrorCode_index[i]:_APIErrorCode_index[i+1]]
}
