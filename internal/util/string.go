package util

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/hcd233/aris-api-tmpl/internal/common/constant"
)

// ToDataURL 将文件转换为 data URL
//
//	@param contentType 文件类型
//	@param bytes
//	@return string
//	@author centonhuang
//	@update 2025-11-13 17:49:49
func ToDataURL(contentType string, bytes []byte) string {
	base64Data := base64.StdEncoding.EncodeToString(bytes)
	return fmt.Sprintf("data:%s;base64,%s", contentType, base64Data)
}

// MaskSecret 掩码敏感信息，保留前 4 和后 4 个字符。
func MaskSecret(key string) string {
	if len(key) <= 8 {
		return constant.MaskSecretPlaceholder
	}
	return fmt.Sprintf("%s***%s", key[:4], key[len(key)-4:])
}

// TruncateFieldValue 截断过长的字符串值。
func TruncateFieldValue(val string, maxLen int) string {
	if len(val) <= maxLen {
		return val
	}

	var builder strings.Builder
	builder.WriteString(val[:maxLen])
	builder.WriteString("...(truncated, total ")
	fmt.Fprintf(&builder, "%d", len(val))
	builder.WriteString(" chars)")
	return builder.String()
}

// TruncateMapValues 递归截断 map 中过长的字符串值。
func TruncateMapValues(input map[string]any, maxLen int) map[string]any {
	result := make(map[string]any, len(input))
	for key, value := range input {
		result[key] = truncateValue(value, maxLen)
	}
	return result
}

func truncateValue(value any, maxLen int) any {
	switch typedValue := value.(type) {
	case string:
		return TruncateFieldValue(typedValue, maxLen)
	case map[string]any:
		return TruncateMapValues(typedValue, maxLen)
	case []any:
		result := make([]any, len(typedValue))
		for i, item := range typedValue {
			result[i] = truncateValue(item, maxLen)
		}
		return result
	default:
		return value
	}
}
