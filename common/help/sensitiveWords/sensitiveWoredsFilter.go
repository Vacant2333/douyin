package sensitiveWords

import (
	"github.com/zeromicro/go-zero/core/stringx"
)

func SensitiveWordsFliter(sensitiveWords []string, context string, replaceMask rune) string {
	//敏感词过滤
	filter := stringx.NewTrie(sensitiveWords, stringx.WithMask(replaceMask)) // 默认替换为*
	safe, _, _ := filter.Filter(context)
	return safe
}
