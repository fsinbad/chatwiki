// Copyright © 2016- 2024 Sesame Network Technology all right reserved

package define

import (
	"strings"
)

const LocalUploadPrefix = `/upload/`

const ImageLimitSize = 100 * 1024          //100KB
const ImageAvatarLimitSize = 1024 * 1024   //1m
const LibFileLimitSize = 100 * 1024 * 1024 //100MB
const LibImageLimitSize = 2 * 1024 * 1024  // 2M

var ImageAllowExt = []string{`heic`, `gif`, `jpg`, `jpeg`, `png`, `swf`, `bmp`, `webp`}
var LibFileAllowExt = []string{`pdf`, `docx`, `txt`, `md`, `xlsx`, `csv`, `html`}
var FormFileAllowExt = []string{`json`, `xlsx`, `csv`}

func IsTableFile(ext string) bool {
	ext = strings.ToLower(ext)
	return ext == `xlsx` || ext == `csv`
}
