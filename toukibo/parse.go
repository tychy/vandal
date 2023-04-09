package toukibo

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	beginContent = "┏━━━━━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓"
	endContent   = "┗━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛"
	separator1   = "┠────────┼─────────────────────────────────────┨"
	separator2   = "┣━━━━━━━━┿━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫"
	separator3   = "┣━━━━━━━━┿━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━┫"
	separator4   = "┣━━━━━━━━┿━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━┫"
	separator5   = "┠────────┼───────────────────────┴─────────────┨"
)

type ToukiboHeader struct {
	CreatedAt      time.Time
	CompanyAddress string
	CompanyName    string
}

func ReadCreatedAt(s string) (time.Time, error) {
	// 正規表現パターン: 全角数字で構成された日付と時刻
	pattern := "([０-９]{2,4}／[０-９]{1,2}／[０-９]{1,2})　*([０-９]{1,2}：[０-９]{1,2})"
	regex := regexp.MustCompile(pattern)

	// 抽出された日付と時刻を表示
	matches := regex.FindStringSubmatch(s)
	if len(matches) > 0 {
		// 全角数字を半角数字に変換
		dateStr := zenkakuToHankaku(matches[1])
		timeStr := zenkakuToHankaku(matches[2])

		// 日付と時刻を time.Time 型に変換
		layout := "2006/01/02 15:04"
		dt, err := time.Parse(layout, fmt.Sprintf("%s %s", dateStr, timeStr))
		if err != nil {
			return time.Time{}, fmt.Errorf("日付と時刻の変換に失敗しました: %w", err)
		}
		return dt, nil
	} else {
		return time.Time{}, fmt.Errorf("日付と時刻が見つかりませんでした")
	}
}

type ToukiboContent struct {
	Header *ToukiboHeader

	HeaderString string
	Content      string
	Parts        []string
}

func findBeginContent(content string) (int, error) {
	index := strings.Index(content, beginContent)
	if index == -1 {
		return 0, fmt.Errorf("not found begin content")
	}
	return index, nil
}

func findEndContent(content string) (int, error) {
	index := strings.Index(content, endContent)
	if index == -1 {
		return 0, fmt.Errorf("not found end content")
	}
	return index, nil
}

func GetHeaderAndContent(content string) (string, string, error) {
	beginContentIdx, err := findBeginContent(content)
	if err != nil {
		return "", "", err
	}

	endContentIdx, err := findEndContent(content)
	if err != nil {
		return "", "", err
	}

	return content[:beginContentIdx], content[beginContentIdx+len(beginContent) : endContentIdx], nil
}

func DivideToukiboContent(input string) (ToukiboContent, error) {
	tc := ToukiboContent{}
	header, content, err := GetHeaderAndContent(input)
	if err != nil {
		return tc, err
	}
	tc.HeaderString = header
	tc.Content = content

	separatorPattern := fmt.Sprintf("%s|%s|%s|%s|%s", separator1, separator2, separator3, separator4, separator5)
	re := regexp.MustCompile(separatorPattern)

	parts := re.Split(tc.Content, -1)

	//for i, part := range parts {
	//	fmt.Printf("Part %d: \n%s\n", i, part)
	//}
	tc.Parts = parts

	return tc, nil
}

func trimPattern(s, pattern string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(s, "")
}

func CleanHoujinNamePart(s string) string {
	// 変更履歴削除
	pattern := "　*│(大正|昭和|平成|令和)[０-９]{1,2}年　+[０-９]{1,2}月[０-９]{1,2}日変更┃ ┃　*│　*"
	cleanedText := trimPattern(s, pattern)

	// 商号が2行にまたがっている場合
	pattern = "　*┃ ┃　*│　*"
	cleanedText = trimPattern(cleanedText, pattern)
	return cleanedText
}

func ParseHeader(s string) (*ToukiboHeader, error) {
	arr := strings.Split(s, " 　")
	header := ToukiboHeader{}
	createdAt, err := ReadCreatedAt(arr[0])
	if err != nil {
		return nil, err
	}
	header.CreatedAt = createdAt
	header.CompanyAddress = arr[2]
	header.CompanyName = arr[3]
	return &header, nil
}

func Parse(input string) (ToukiboContent, error) {
	tc, err := DivideToukiboContent(input)
	if err != nil {
		return tc, err
	}
	header, err := ParseHeader(tc.HeaderString)
	if err != nil {
		return tc, err
	}
	tc.Header = header

	return tc, nil
}
