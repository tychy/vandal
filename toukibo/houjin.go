package toukibo

import (
	"fmt"
	"regexp"
	"time"
)

const ZenkakuZero = '０'
const ZenkakuNine = '９'
const ZenkakuSpace = '　'
const ZenkakuColon = '：'
const ZenkakuSlash = '／'

func zenkakuToHankaku(s string) string {
	var result string
	for _, r := range s {
		if r >= ZenkakuZero && r <= ZenkakuNine {
			result += string(r - ZenkakuZero + '0')
		} else if r == ZenkakuSlash {
			result += "/"
		} else if r == ZenkakuColon {
			result += ":"
		} else if r == ZenkakuSpace {
			result += " "
		} else {
			result += string(r)
		}
	}
	return result
}

type Houjin struct {
	content   string
	CreatedAt time.Time
}

func (h Houjin) GetContent() string {
	return h.content
}

func (h *Houjin) ReadCreatedAt() error {
	// 正規表現パターン: 全角数字で構成された日付と時刻
	pattern := "([０-９]{2,4}／[０-９]{1,2}／[０-９]{1,2})　*([０-９]{1,2}：[０-９]{1,2})"
	regex := regexp.MustCompile(pattern)

	// 抽出された日付と時刻を表示
	matches := regex.FindStringSubmatch(h.content)
	if len(matches) > 0 {
		// 全角数字を半角数字に変換
		dateStr := zenkakuToHankaku(matches[1])
		timeStr := zenkakuToHankaku(matches[2])

		// 日付と時刻を time.Time 型に変換
		layout := "2006/01/02 15:04"
		dt, err := time.Parse(layout, fmt.Sprintf("%s %s", dateStr, timeStr))
		if err != nil {
			return fmt.Errorf("日付と時刻の変換に失敗しました: %w", err)
		}
		h.CreatedAt = dt
	} else {
		return fmt.Errorf("日付と時刻が見つかりませんでした")
	}
	return nil
}

func Extract(content string) (Houjin, error) {
	houjin := Houjin{content: content}
	houjin.ReadCreatedAt()
	return houjin, nil
}
