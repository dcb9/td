// Code generated by gotdgen, DO NOT EDIT.

package e2e

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// TypesMap returns mapping from type ids to TL type names.
func TypesMap() map[uint32]string {
	return map[uint32]string{
		IntTypeID:                                      "int#a8509bda",
		LongTypeID:                                     "long#22076cba",
		DoubleTypeID:                                   "double#2210c154",
		StringTypeID:                                   "string#b5286e24",
		BytesTypeID:                                    "bytes#e937bb82",
		BoolFalseTypeID:                                "boolFalse#bc799737",
		BoolTrueTypeID:                                 "boolTrue#997275b5",
		TrueTypeID:                                     "true#3fedd339",
		DecryptedMessage8TypeID:                        "decryptedMessage8#1f814f1f",
		DecryptedMessageService8TypeID:                 "decryptedMessageService8#aa48327d",
		DecryptedMessageMediaEmptyTypeID:               "decryptedMessageMediaEmpty#89f5c4a",
		DecryptedMessageMediaPhoto23TypeID:             "decryptedMessageMediaPhoto23#32798a8c",
		DecryptedMessageMediaVideo8TypeID:              "decryptedMessageMediaVideo8#4cee6ef3",
		DecryptedMessageMediaGeoPointTypeID:            "decryptedMessageMediaGeoPoint#35480a59",
		DecryptedMessageMediaContactTypeID:             "decryptedMessageMediaContact#588a0a97",
		DecryptedMessageActionSetMessageTTLTypeID:      "decryptedMessageActionSetMessageTTL#a1733aec",
		DecryptedMessageMediaDocument23TypeID:          "decryptedMessageMediaDocument23#b095434b",
		DecryptedMessageMediaAudio8TypeID:              "decryptedMessageMediaAudio8#6080758f",
		DecryptedMessageActionReadMessagesTypeID:       "decryptedMessageActionReadMessages#c4f40be",
		DecryptedMessageActionDeleteMessagesTypeID:     "decryptedMessageActionDeleteMessages#65614304",
		DecryptedMessageActionScreenshotMessagesTypeID: "decryptedMessageActionScreenshotMessages#8ac1f475",
		DecryptedMessageActionFlushHistoryTypeID:       "decryptedMessageActionFlushHistory#6719e45c",
		DecryptedMessage23TypeID:                       "decryptedMessage23#204d3878",
		DecryptedMessageServiceTypeID:                  "decryptedMessageService#73164160",
		DecryptedMessageMediaVideo23TypeID:             "decryptedMessageMediaVideo23#524a415d",
		DecryptedMessageMediaAudioTypeID:               "decryptedMessageMediaAudio#57e0a9cb",
		DecryptedMessageLayerTypeID:                    "decryptedMessageLayer#1be31789",
		SendMessageTypingActionTypeID:                  "sendMessageTypingAction#16bf744e",
		SendMessageCancelActionTypeID:                  "sendMessageCancelAction#fd5ec8f5",
		SendMessageRecordVideoActionTypeID:             "sendMessageRecordVideoAction#a187d66f",
		SendMessageUploadVideoActionTypeID:             "sendMessageUploadVideoAction#92042ff7",
		SendMessageRecordAudioActionTypeID:             "sendMessageRecordAudioAction#d52f73f7",
		SendMessageUploadAudioActionTypeID:             "sendMessageUploadAudioAction#e6ac8a6f",
		SendMessageUploadPhotoActionTypeID:             "sendMessageUploadPhotoAction#990a3c1a",
		SendMessageUploadDocumentActionTypeID:          "sendMessageUploadDocumentAction#8faee98e",
		SendMessageGeoLocationActionTypeID:             "sendMessageGeoLocationAction#176f8ba1",
		SendMessageChooseContactActionTypeID:           "sendMessageChooseContactAction#628cbc6f",
		DecryptedMessageActionResendTypeID:             "decryptedMessageActionResend#511110b0",
		DecryptedMessageActionNotifyLayerTypeID:        "decryptedMessageActionNotifyLayer#f3048883",
		DecryptedMessageActionTypingTypeID:             "decryptedMessageActionTyping#ccb27641",
		DecryptedMessageActionRequestKeyTypeID:         "decryptedMessageActionRequestKey#f3c9611b",
		DecryptedMessageActionAcceptKeyTypeID:          "decryptedMessageActionAcceptKey#6fe1735b",
		DecryptedMessageActionAbortKeyTypeID:           "decryptedMessageActionAbortKey#dd05ec6b",
		DecryptedMessageActionCommitKeyTypeID:          "decryptedMessageActionCommitKey#ec2e0b9b",
		DecryptedMessageActionNoopTypeID:               "decryptedMessageActionNoop#a82fdd63",
		DocumentAttributeImageSizeTypeID:               "documentAttributeImageSize#6c37c15c",
		DocumentAttributeAnimatedTypeID:                "documentAttributeAnimated#11b58939",
		DocumentAttributeSticker23TypeID:               "documentAttributeSticker23#fb0a5727",
		DocumentAttributeVideoTypeID:                   "documentAttributeVideo#5910cccb",
		DocumentAttributeAudio23TypeID:                 "documentAttributeAudio23#51448e5",
		DocumentAttributeFilenameTypeID:                "documentAttributeFilename#15590068",
		PhotoSizeEmptyTypeID:                           "photoSizeEmpty#e17e23c",
		PhotoSizeTypeID:                                "photoSize#77bfb61b",
		PhotoCachedSizeTypeID:                          "photoCachedSize#e9a734fa",
		FileLocationUnavailableTypeID:                  "fileLocationUnavailable#7c596b46",
		FileLocationTypeID:                             "fileLocation#53d69076",
		DecryptedMessageMediaExternalDocumentTypeID:    "decryptedMessageMediaExternalDocument#fa95b0dd",
		DocumentAttributeAudio45TypeID:                 "documentAttributeAudio45#ded218e0",
		DecryptedMessage46TypeID:                       "decryptedMessage46#36b091de",
		DecryptedMessageMediaPhotoTypeID:               "decryptedMessageMediaPhoto#f1fa8d78",
		DecryptedMessageMediaVideoTypeID:               "decryptedMessageMediaVideo#970c8c0e",
		DecryptedMessageMediaDocument46TypeID:          "decryptedMessageMediaDocument46#7afe8ae2",
		DocumentAttributeStickerTypeID:                 "documentAttributeSticker#3a556302",
		DocumentAttributeAudioTypeID:                   "documentAttributeAudio#9852f9c6",
		MessageEntityUnknownTypeID:                     "messageEntityUnknown#bb92ba95",
		MessageEntityMentionTypeID:                     "messageEntityMention#fa04579d",
		MessageEntityHashtagTypeID:                     "messageEntityHashtag#6f635b0d",
		MessageEntityBotCommandTypeID:                  "messageEntityBotCommand#6cef8ac7",
		MessageEntityURLTypeID:                         "messageEntityUrl#6ed02538",
		MessageEntityEmailTypeID:                       "messageEntityEmail#64e475c2",
		MessageEntityBoldTypeID:                        "messageEntityBold#bd610bc9",
		MessageEntityItalicTypeID:                      "messageEntityItalic#826f8b60",
		MessageEntityCodeTypeID:                        "messageEntityCode#28a20571",
		MessageEntityPreTypeID:                         "messageEntityPre#73924be0",
		MessageEntityTextURLTypeID:                     "messageEntityTextUrl#76a6d327",
		MessageEntityMentionNameTypeID:                 "messageEntityMentionName#352dca58",
		MessageEntityPhoneTypeID:                       "messageEntityPhone#9b69e34b",
		MessageEntityCashtagTypeID:                     "messageEntityCashtag#4c4e743f",
		MessageEntityBankCardTypeID:                    "messageEntityBankCard#761e6af4",
		InputStickerSetShortNameTypeID:                 "inputStickerSetShortName#861cc8a0",
		InputStickerSetEmptyTypeID:                     "inputStickerSetEmpty#ffb62b95",
		DecryptedMessageMediaVenueTypeID:               "decryptedMessageMediaVenue#8a0df56f",
		DecryptedMessageMediaWebPageTypeID:             "decryptedMessageMediaWebPage#e50511d8",
		SendMessageRecordRoundActionTypeID:             "sendMessageRecordRoundAction#88f27fbc",
		SendMessageUploadRoundActionTypeID:             "sendMessageUploadRoundAction#bb718624",
		DocumentAttributeVideo66TypeID:                 "documentAttributeVideo66#ef02ce6",
		DecryptedMessageTypeID:                         "decryptedMessage#91cc4674",
		MessageEntityUnderlineTypeID:                   "messageEntityUnderline#9c4e7e8b",
		MessageEntityStrikeTypeID:                      "messageEntityStrike#bf0693d4",
		MessageEntityBlockquoteTypeID:                  "messageEntityBlockquote#20df5d0",
		DecryptedMessageMediaDocumentTypeID:            "decryptedMessageMediaDocument#6abd9782",
		TestDummyFunctionRequestTypeID:                 "test.dummyFunction#c8357709",
	}
}

// NamesMap returns mapping from type names to TL type ids.
func NamesMap() map[string]uint32 {
	return map[string]uint32{
		"int":                                      IntTypeID,
		"long":                                     LongTypeID,
		"double":                                   DoubleTypeID,
		"string":                                   StringTypeID,
		"bytes":                                    BytesTypeID,
		"boolFalse":                                BoolFalseTypeID,
		"boolTrue":                                 BoolTrueTypeID,
		"true":                                     TrueTypeID,
		"decryptedMessage8":                        DecryptedMessage8TypeID,
		"decryptedMessageService8":                 DecryptedMessageService8TypeID,
		"decryptedMessageMediaEmpty":               DecryptedMessageMediaEmptyTypeID,
		"decryptedMessageMediaPhoto23":             DecryptedMessageMediaPhoto23TypeID,
		"decryptedMessageMediaVideo8":              DecryptedMessageMediaVideo8TypeID,
		"decryptedMessageMediaGeoPoint":            DecryptedMessageMediaGeoPointTypeID,
		"decryptedMessageMediaContact":             DecryptedMessageMediaContactTypeID,
		"decryptedMessageActionSetMessageTTL":      DecryptedMessageActionSetMessageTTLTypeID,
		"decryptedMessageMediaDocument23":          DecryptedMessageMediaDocument23TypeID,
		"decryptedMessageMediaAudio8":              DecryptedMessageMediaAudio8TypeID,
		"decryptedMessageActionReadMessages":       DecryptedMessageActionReadMessagesTypeID,
		"decryptedMessageActionDeleteMessages":     DecryptedMessageActionDeleteMessagesTypeID,
		"decryptedMessageActionScreenshotMessages": DecryptedMessageActionScreenshotMessagesTypeID,
		"decryptedMessageActionFlushHistory":       DecryptedMessageActionFlushHistoryTypeID,
		"decryptedMessage23":                       DecryptedMessage23TypeID,
		"decryptedMessageService":                  DecryptedMessageServiceTypeID,
		"decryptedMessageMediaVideo23":             DecryptedMessageMediaVideo23TypeID,
		"decryptedMessageMediaAudio":               DecryptedMessageMediaAudioTypeID,
		"decryptedMessageLayer":                    DecryptedMessageLayerTypeID,
		"sendMessageTypingAction":                  SendMessageTypingActionTypeID,
		"sendMessageCancelAction":                  SendMessageCancelActionTypeID,
		"sendMessageRecordVideoAction":             SendMessageRecordVideoActionTypeID,
		"sendMessageUploadVideoAction":             SendMessageUploadVideoActionTypeID,
		"sendMessageRecordAudioAction":             SendMessageRecordAudioActionTypeID,
		"sendMessageUploadAudioAction":             SendMessageUploadAudioActionTypeID,
		"sendMessageUploadPhotoAction":             SendMessageUploadPhotoActionTypeID,
		"sendMessageUploadDocumentAction":          SendMessageUploadDocumentActionTypeID,
		"sendMessageGeoLocationAction":             SendMessageGeoLocationActionTypeID,
		"sendMessageChooseContactAction":           SendMessageChooseContactActionTypeID,
		"decryptedMessageActionResend":             DecryptedMessageActionResendTypeID,
		"decryptedMessageActionNotifyLayer":        DecryptedMessageActionNotifyLayerTypeID,
		"decryptedMessageActionTyping":             DecryptedMessageActionTypingTypeID,
		"decryptedMessageActionRequestKey":         DecryptedMessageActionRequestKeyTypeID,
		"decryptedMessageActionAcceptKey":          DecryptedMessageActionAcceptKeyTypeID,
		"decryptedMessageActionAbortKey":           DecryptedMessageActionAbortKeyTypeID,
		"decryptedMessageActionCommitKey":          DecryptedMessageActionCommitKeyTypeID,
		"decryptedMessageActionNoop":               DecryptedMessageActionNoopTypeID,
		"documentAttributeImageSize":               DocumentAttributeImageSizeTypeID,
		"documentAttributeAnimated":                DocumentAttributeAnimatedTypeID,
		"documentAttributeSticker23":               DocumentAttributeSticker23TypeID,
		"documentAttributeVideo":                   DocumentAttributeVideoTypeID,
		"documentAttributeAudio23":                 DocumentAttributeAudio23TypeID,
		"documentAttributeFilename":                DocumentAttributeFilenameTypeID,
		"photoSizeEmpty":                           PhotoSizeEmptyTypeID,
		"photoSize":                                PhotoSizeTypeID,
		"photoCachedSize":                          PhotoCachedSizeTypeID,
		"fileLocationUnavailable":                  FileLocationUnavailableTypeID,
		"fileLocation":                             FileLocationTypeID,
		"decryptedMessageMediaExternalDocument":    DecryptedMessageMediaExternalDocumentTypeID,
		"documentAttributeAudio45":                 DocumentAttributeAudio45TypeID,
		"decryptedMessage46":                       DecryptedMessage46TypeID,
		"decryptedMessageMediaPhoto":               DecryptedMessageMediaPhotoTypeID,
		"decryptedMessageMediaVideo":               DecryptedMessageMediaVideoTypeID,
		"decryptedMessageMediaDocument46":          DecryptedMessageMediaDocument46TypeID,
		"documentAttributeSticker":                 DocumentAttributeStickerTypeID,
		"documentAttributeAudio":                   DocumentAttributeAudioTypeID,
		"messageEntityUnknown":                     MessageEntityUnknownTypeID,
		"messageEntityMention":                     MessageEntityMentionTypeID,
		"messageEntityHashtag":                     MessageEntityHashtagTypeID,
		"messageEntityBotCommand":                  MessageEntityBotCommandTypeID,
		"messageEntityUrl":                         MessageEntityURLTypeID,
		"messageEntityEmail":                       MessageEntityEmailTypeID,
		"messageEntityBold":                        MessageEntityBoldTypeID,
		"messageEntityItalic":                      MessageEntityItalicTypeID,
		"messageEntityCode":                        MessageEntityCodeTypeID,
		"messageEntityPre":                         MessageEntityPreTypeID,
		"messageEntityTextUrl":                     MessageEntityTextURLTypeID,
		"messageEntityMentionName":                 MessageEntityMentionNameTypeID,
		"messageEntityPhone":                       MessageEntityPhoneTypeID,
		"messageEntityCashtag":                     MessageEntityCashtagTypeID,
		"messageEntityBankCard":                    MessageEntityBankCardTypeID,
		"inputStickerSetShortName":                 InputStickerSetShortNameTypeID,
		"inputStickerSetEmpty":                     InputStickerSetEmptyTypeID,
		"decryptedMessageMediaVenue":               DecryptedMessageMediaVenueTypeID,
		"decryptedMessageMediaWebPage":             DecryptedMessageMediaWebPageTypeID,
		"sendMessageRecordRoundAction":             SendMessageRecordRoundActionTypeID,
		"sendMessageUploadRoundAction":             SendMessageUploadRoundActionTypeID,
		"documentAttributeVideo66":                 DocumentAttributeVideo66TypeID,
		"decryptedMessage":                         DecryptedMessageTypeID,
		"messageEntityUnderline":                   MessageEntityUnderlineTypeID,
		"messageEntityStrike":                      MessageEntityStrikeTypeID,
		"messageEntityBlockquote":                  MessageEntityBlockquoteTypeID,
		"decryptedMessageMediaDocument":            DecryptedMessageMediaDocumentTypeID,
		"test.dummyFunction":                       TestDummyFunctionRequestTypeID,
	}
}

// TypesConstructorMap maps type ids to constructors.
func TypesConstructorMap() map[uint32]func() bin.Object {
	return map[uint32]func() bin.Object{
		IntTypeID:                                      func() bin.Object { return &Int{} },
		LongTypeID:                                     func() bin.Object { return &Long{} },
		DoubleTypeID:                                   func() bin.Object { return &Double{} },
		StringTypeID:                                   func() bin.Object { return &String{} },
		BytesTypeID:                                    func() bin.Object { return &Bytes{} },
		BoolFalseTypeID:                                func() bin.Object { return &BoolFalse{} },
		BoolTrueTypeID:                                 func() bin.Object { return &BoolTrue{} },
		TrueTypeID:                                     func() bin.Object { return &True{} },
		DecryptedMessage8TypeID:                        func() bin.Object { return &DecryptedMessage8{} },
		DecryptedMessageService8TypeID:                 func() bin.Object { return &DecryptedMessageService8{} },
		DecryptedMessageMediaEmptyTypeID:               func() bin.Object { return &DecryptedMessageMediaEmpty{} },
		DecryptedMessageMediaPhoto23TypeID:             func() bin.Object { return &DecryptedMessageMediaPhoto23{} },
		DecryptedMessageMediaVideo8TypeID:              func() bin.Object { return &DecryptedMessageMediaVideo8{} },
		DecryptedMessageMediaGeoPointTypeID:            func() bin.Object { return &DecryptedMessageMediaGeoPoint{} },
		DecryptedMessageMediaContactTypeID:             func() bin.Object { return &DecryptedMessageMediaContact{} },
		DecryptedMessageActionSetMessageTTLTypeID:      func() bin.Object { return &DecryptedMessageActionSetMessageTTL{} },
		DecryptedMessageMediaDocument23TypeID:          func() bin.Object { return &DecryptedMessageMediaDocument23{} },
		DecryptedMessageMediaAudio8TypeID:              func() bin.Object { return &DecryptedMessageMediaAudio8{} },
		DecryptedMessageActionReadMessagesTypeID:       func() bin.Object { return &DecryptedMessageActionReadMessages{} },
		DecryptedMessageActionDeleteMessagesTypeID:     func() bin.Object { return &DecryptedMessageActionDeleteMessages{} },
		DecryptedMessageActionScreenshotMessagesTypeID: func() bin.Object { return &DecryptedMessageActionScreenshotMessages{} },
		DecryptedMessageActionFlushHistoryTypeID:       func() bin.Object { return &DecryptedMessageActionFlushHistory{} },
		DecryptedMessage23TypeID:                       func() bin.Object { return &DecryptedMessage23{} },
		DecryptedMessageServiceTypeID:                  func() bin.Object { return &DecryptedMessageService{} },
		DecryptedMessageMediaVideo23TypeID:             func() bin.Object { return &DecryptedMessageMediaVideo23{} },
		DecryptedMessageMediaAudioTypeID:               func() bin.Object { return &DecryptedMessageMediaAudio{} },
		DecryptedMessageLayerTypeID:                    func() bin.Object { return &DecryptedMessageLayer{} },
		SendMessageTypingActionTypeID:                  func() bin.Object { return &SendMessageTypingAction{} },
		SendMessageCancelActionTypeID:                  func() bin.Object { return &SendMessageCancelAction{} },
		SendMessageRecordVideoActionTypeID:             func() bin.Object { return &SendMessageRecordVideoAction{} },
		SendMessageUploadVideoActionTypeID:             func() bin.Object { return &SendMessageUploadVideoAction{} },
		SendMessageRecordAudioActionTypeID:             func() bin.Object { return &SendMessageRecordAudioAction{} },
		SendMessageUploadAudioActionTypeID:             func() bin.Object { return &SendMessageUploadAudioAction{} },
		SendMessageUploadPhotoActionTypeID:             func() bin.Object { return &SendMessageUploadPhotoAction{} },
		SendMessageUploadDocumentActionTypeID:          func() bin.Object { return &SendMessageUploadDocumentAction{} },
		SendMessageGeoLocationActionTypeID:             func() bin.Object { return &SendMessageGeoLocationAction{} },
		SendMessageChooseContactActionTypeID:           func() bin.Object { return &SendMessageChooseContactAction{} },
		DecryptedMessageActionResendTypeID:             func() bin.Object { return &DecryptedMessageActionResend{} },
		DecryptedMessageActionNotifyLayerTypeID:        func() bin.Object { return &DecryptedMessageActionNotifyLayer{} },
		DecryptedMessageActionTypingTypeID:             func() bin.Object { return &DecryptedMessageActionTyping{} },
		DecryptedMessageActionRequestKeyTypeID:         func() bin.Object { return &DecryptedMessageActionRequestKey{} },
		DecryptedMessageActionAcceptKeyTypeID:          func() bin.Object { return &DecryptedMessageActionAcceptKey{} },
		DecryptedMessageActionAbortKeyTypeID:           func() bin.Object { return &DecryptedMessageActionAbortKey{} },
		DecryptedMessageActionCommitKeyTypeID:          func() bin.Object { return &DecryptedMessageActionCommitKey{} },
		DecryptedMessageActionNoopTypeID:               func() bin.Object { return &DecryptedMessageActionNoop{} },
		DocumentAttributeImageSizeTypeID:               func() bin.Object { return &DocumentAttributeImageSize{} },
		DocumentAttributeAnimatedTypeID:                func() bin.Object { return &DocumentAttributeAnimated{} },
		DocumentAttributeSticker23TypeID:               func() bin.Object { return &DocumentAttributeSticker23{} },
		DocumentAttributeVideoTypeID:                   func() bin.Object { return &DocumentAttributeVideo{} },
		DocumentAttributeAudio23TypeID:                 func() bin.Object { return &DocumentAttributeAudio23{} },
		DocumentAttributeFilenameTypeID:                func() bin.Object { return &DocumentAttributeFilename{} },
		PhotoSizeEmptyTypeID:                           func() bin.Object { return &PhotoSizeEmpty{} },
		PhotoSizeTypeID:                                func() bin.Object { return &PhotoSize{} },
		PhotoCachedSizeTypeID:                          func() bin.Object { return &PhotoCachedSize{} },
		FileLocationUnavailableTypeID:                  func() bin.Object { return &FileLocationUnavailable{} },
		FileLocationTypeID:                             func() bin.Object { return &FileLocation{} },
		DecryptedMessageMediaExternalDocumentTypeID:    func() bin.Object { return &DecryptedMessageMediaExternalDocument{} },
		DocumentAttributeAudio45TypeID:                 func() bin.Object { return &DocumentAttributeAudio45{} },
		DecryptedMessage46TypeID:                       func() bin.Object { return &DecryptedMessage46{} },
		DecryptedMessageMediaPhotoTypeID:               func() bin.Object { return &DecryptedMessageMediaPhoto{} },
		DecryptedMessageMediaVideoTypeID:               func() bin.Object { return &DecryptedMessageMediaVideo{} },
		DecryptedMessageMediaDocument46TypeID:          func() bin.Object { return &DecryptedMessageMediaDocument46{} },
		DocumentAttributeStickerTypeID:                 func() bin.Object { return &DocumentAttributeSticker{} },
		DocumentAttributeAudioTypeID:                   func() bin.Object { return &DocumentAttributeAudio{} },
		MessageEntityUnknownTypeID:                     func() bin.Object { return &MessageEntityUnknown{} },
		MessageEntityMentionTypeID:                     func() bin.Object { return &MessageEntityMention{} },
		MessageEntityHashtagTypeID:                     func() bin.Object { return &MessageEntityHashtag{} },
		MessageEntityBotCommandTypeID:                  func() bin.Object { return &MessageEntityBotCommand{} },
		MessageEntityURLTypeID:                         func() bin.Object { return &MessageEntityURL{} },
		MessageEntityEmailTypeID:                       func() bin.Object { return &MessageEntityEmail{} },
		MessageEntityBoldTypeID:                        func() bin.Object { return &MessageEntityBold{} },
		MessageEntityItalicTypeID:                      func() bin.Object { return &MessageEntityItalic{} },
		MessageEntityCodeTypeID:                        func() bin.Object { return &MessageEntityCode{} },
		MessageEntityPreTypeID:                         func() bin.Object { return &MessageEntityPre{} },
		MessageEntityTextURLTypeID:                     func() bin.Object { return &MessageEntityTextURL{} },
		MessageEntityMentionNameTypeID:                 func() bin.Object { return &MessageEntityMentionName{} },
		MessageEntityPhoneTypeID:                       func() bin.Object { return &MessageEntityPhone{} },
		MessageEntityCashtagTypeID:                     func() bin.Object { return &MessageEntityCashtag{} },
		MessageEntityBankCardTypeID:                    func() bin.Object { return &MessageEntityBankCard{} },
		InputStickerSetShortNameTypeID:                 func() bin.Object { return &InputStickerSetShortName{} },
		InputStickerSetEmptyTypeID:                     func() bin.Object { return &InputStickerSetEmpty{} },
		DecryptedMessageMediaVenueTypeID:               func() bin.Object { return &DecryptedMessageMediaVenue{} },
		DecryptedMessageMediaWebPageTypeID:             func() bin.Object { return &DecryptedMessageMediaWebPage{} },
		SendMessageRecordRoundActionTypeID:             func() bin.Object { return &SendMessageRecordRoundAction{} },
		SendMessageUploadRoundActionTypeID:             func() bin.Object { return &SendMessageUploadRoundAction{} },
		DocumentAttributeVideo66TypeID:                 func() bin.Object { return &DocumentAttributeVideo66{} },
		DecryptedMessageTypeID:                         func() bin.Object { return &DecryptedMessage{} },
		MessageEntityUnderlineTypeID:                   func() bin.Object { return &MessageEntityUnderline{} },
		MessageEntityStrikeTypeID:                      func() bin.Object { return &MessageEntityStrike{} },
		MessageEntityBlockquoteTypeID:                  func() bin.Object { return &MessageEntityBlockquote{} },
		DecryptedMessageMediaDocumentTypeID:            func() bin.Object { return &DecryptedMessageMediaDocument{} },
		TestDummyFunctionRequestTypeID:                 func() bin.Object { return &TestDummyFunctionRequest{} },
	}
}

// ClassConstructorsMap maps class schema name to constructors type ids.
func ClassConstructorsMap() map[string][]uint32 {
	return map[string][]uint32{
		BoolClassName: {
			BoolFalseTypeID,
			BoolTrueTypeID,
		},
		DecryptedMessageActionClassName: {
			DecryptedMessageActionSetMessageTTLTypeID,
			DecryptedMessageActionReadMessagesTypeID,
			DecryptedMessageActionDeleteMessagesTypeID,
			DecryptedMessageActionScreenshotMessagesTypeID,
			DecryptedMessageActionFlushHistoryTypeID,
			DecryptedMessageActionResendTypeID,
			DecryptedMessageActionNotifyLayerTypeID,
			DecryptedMessageActionTypingTypeID,
			DecryptedMessageActionRequestKeyTypeID,
			DecryptedMessageActionAcceptKeyTypeID,
			DecryptedMessageActionAbortKeyTypeID,
			DecryptedMessageActionCommitKeyTypeID,
			DecryptedMessageActionNoopTypeID,
		},
		DecryptedMessageClassName: {
			DecryptedMessage8TypeID,
			DecryptedMessageService8TypeID,
			DecryptedMessage23TypeID,
			DecryptedMessageServiceTypeID,
			DecryptedMessage46TypeID,
			DecryptedMessageTypeID,
		},
		DecryptedMessageMediaClassName: {
			DecryptedMessageMediaEmptyTypeID,
			DecryptedMessageMediaPhoto23TypeID,
			DecryptedMessageMediaVideo8TypeID,
			DecryptedMessageMediaGeoPointTypeID,
			DecryptedMessageMediaContactTypeID,
			DecryptedMessageMediaDocument23TypeID,
			DecryptedMessageMediaAudio8TypeID,
			DecryptedMessageMediaVideo23TypeID,
			DecryptedMessageMediaAudioTypeID,
			DecryptedMessageMediaExternalDocumentTypeID,
			DecryptedMessageMediaPhotoTypeID,
			DecryptedMessageMediaVideoTypeID,
			DecryptedMessageMediaDocument46TypeID,
			DecryptedMessageMediaVenueTypeID,
			DecryptedMessageMediaWebPageTypeID,
			DecryptedMessageMediaDocumentTypeID,
		},
		DocumentAttributeClassName: {
			DocumentAttributeImageSizeTypeID,
			DocumentAttributeAnimatedTypeID,
			DocumentAttributeSticker23TypeID,
			DocumentAttributeVideoTypeID,
			DocumentAttributeAudio23TypeID,
			DocumentAttributeFilenameTypeID,
			DocumentAttributeAudio45TypeID,
			DocumentAttributeStickerTypeID,
			DocumentAttributeAudioTypeID,
			DocumentAttributeVideo66TypeID,
		},
		FileLocationClassName: {
			FileLocationUnavailableTypeID,
			FileLocationTypeID,
		},
		InputStickerSetClassName: {
			InputStickerSetShortNameTypeID,
			InputStickerSetEmptyTypeID,
		},
		MessageEntityClassName: {
			MessageEntityUnknownTypeID,
			MessageEntityMentionTypeID,
			MessageEntityHashtagTypeID,
			MessageEntityBotCommandTypeID,
			MessageEntityURLTypeID,
			MessageEntityEmailTypeID,
			MessageEntityBoldTypeID,
			MessageEntityItalicTypeID,
			MessageEntityCodeTypeID,
			MessageEntityPreTypeID,
			MessageEntityTextURLTypeID,
			MessageEntityMentionNameTypeID,
			MessageEntityPhoneTypeID,
			MessageEntityCashtagTypeID,
			MessageEntityBankCardTypeID,
			MessageEntityUnderlineTypeID,
			MessageEntityStrikeTypeID,
			MessageEntityBlockquoteTypeID,
		},
		PhotoSizeClassName: {
			PhotoSizeEmptyTypeID,
			PhotoSizeTypeID,
			PhotoCachedSizeTypeID,
		},
		SendMessageActionClassName: {
			SendMessageTypingActionTypeID,
			SendMessageCancelActionTypeID,
			SendMessageRecordVideoActionTypeID,
			SendMessageUploadVideoActionTypeID,
			SendMessageRecordAudioActionTypeID,
			SendMessageUploadAudioActionTypeID,
			SendMessageUploadPhotoActionTypeID,
			SendMessageUploadDocumentActionTypeID,
			SendMessageGeoLocationActionTypeID,
			SendMessageChooseContactActionTypeID,
			SendMessageRecordRoundActionTypeID,
			SendMessageUploadRoundActionTypeID,
		},
	}
}
