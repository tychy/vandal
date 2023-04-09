package toukibo

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	ZenkakuZero          = '０'
	ZenkakuNine          = '９'
	ZenkakuA             = 'Ａ'
	ZenkakuZ             = 'Ｚ'
	ZenkakuSmallA        = 'ａ'
	ZenkakuSmallZ        = 'ｚ'
	ZenkakuSpace         = '　'
	ZenkakuColon         = '：'
	ZenkakuSlash         = '／'
	ZenkakuHyphen        = '－'
	ZenkakuStringPattern = `\p{Han}\p{Hiragana}\p{Katakana}Ａ-Ｚａ-ｚ０-９A-Za-z0-9＆’，‐．・ー\s　。－`
)

func zenkakuToHankaku(s string) string {
	var result string
	for _, r := range s {
		if r >= ZenkakuZero && r <= ZenkakuNine {
			result += string(r - ZenkakuZero + '0')
		} else if r >= ZenkakuA && r <= ZenkakuZ {
			result += string(r - ZenkakuA + 'A')
		} else if r >= ZenkakuSmallA && r <= ZenkakuSmallZ {
			result += string(r - ZenkakuSmallA + 'a')
		} else if r == ZenkakuSlash {
			result += "/"
		} else if r == ZenkakuColon {
			result += ":"
		} else if r == ZenkakuSpace {
			result += " "
		} else if r == ZenkakuHyphen {
			result += "-"
		} else {
			result += string(r)
		}
	}
	return result
}

type HoujinkakuType string

const (
	HoujinKakuKabusiki     HoujinkakuType = "株式会社"
	HoujinKakuYugen        HoujinkakuType = "有限会社"
	HoujinKakuGoudou       HoujinkakuType = "合同会社"
	HoujinKakuGousi        HoujinkakuType = "合資会社"
	HoujinKakuGoumei       HoujinkakuType = "合名会社"
	HoujinKakuTokutei      HoujinkakuType = "特定目的会社"
	HoujinKakuKyodou       HoujinkakuType = "協同組合"
	HoujinKakuRoudou       HoujinkakuType = "労働組合"
	HoujinKakuSinrin       HoujinkakuType = "森林組合"
	HoujinKakuSeikatuEisei HoujinkakuType = "生活衛生同業組合"
	HoujinKakuSinyou       HoujinkakuType = "信用金庫"
	HoujinKakuShokoukai    HoujinkakuType = "商工会"
	HoujinKakuKoueki       HoujinkakuType = "公益財団法人"
	HoujinKakuNouji        HoujinkakuType = "農事組合"
	HoujinKakuKanriKumiai  HoujinkakuType = "管理組合法人"
	HoujinKakuIryo         HoujinkakuType = "医療法人"
	HoujinKakuSihoshosi    HoujinkakuType = "司法書士法人"
	HoujinKakuZeirishi     HoujinkakuType = "税理士法人"
	HoujinKakuShakaifukusi HoujinkakuType = "社会福祉法人"
	HoujinKakuIppanShadan  HoujinkakuType = "一般社団法人"
	HoujinKakuIppanZaisan  HoujinkakuType = "一般財産法人"
	HoujinKakuIppanZaidan  HoujinkakuType = "一般財団法人"
	HoujinKakuNPO          HoujinkakuType = "NPO法人"
	HoujinKakuHieiri       HoujinkakuType = "特定非営利活動法人"
)

func FindHoujinKaku(s string) HoujinkakuType {
	if strings.Contains(s, "株式会社") {
		return HoujinKakuKabusiki
	} else if strings.Contains(s, "有限会社") {
		return HoujinKakuYugen
	} else if strings.Contains(s, "合同会社") {
		return HoujinKakuGoudou
	} else if strings.Contains(s, "合資会社") {
		return HoujinKakuGousi
	} else if strings.Contains(s, "合名会社") {
		return HoujinKakuGoumei
	} else if strings.Contains(s, "特定目的会社") {
		return HoujinKakuTokutei
	} else if strings.Contains(s, "協同組合") {
		return HoujinKakuKyodou
	} else if strings.Contains(s, "労働組合") {
		return HoujinKakuRoudou
	} else if strings.Contains(s, "森林組合") {
		return HoujinKakuSinrin
	} else if strings.Contains(s, "生活衛生同業組合") {
		return HoujinKakuSeikatuEisei
	} else if strings.Contains(s, "信用金庫") {
		return HoujinKakuSinyou
	} else if strings.Contains(s, "商工会") {
		return HoujinKakuShokoukai
	} else if strings.Contains(s, "公益財団法人") {
		return HoujinKakuKoueki
	} else if strings.Contains(s, "農事組合") {
		return HoujinKakuNouji
	} else if strings.Contains(s, "管理組合法人") {
		return HoujinKakuKanriKumiai
	} else if strings.Contains(s, "医療法人") {
		return HoujinKakuIryo
	} else if strings.Contains(s, "司法書士法人") {
		return HoujinKakuSihoshosi
	} else if strings.Contains(s, "税理士法人") {
		return HoujinKakuZeirishi
	} else if strings.Contains(s, "社会福祉法人") {
		return HoujinKakuShakaifukusi
	} else if strings.Contains(s, "一般社団法人") {
		return HoujinKakuIppanShadan
	} else if strings.Contains(s, "一般財産法人") {
		return HoujinKakuIppanZaisan
	} else if strings.Contains(s, "一般財団法人") {
		return HoujinKakuIppanZaidan
	} else if strings.Contains(s, "NPO法人") {
		return HoujinKakuNPO
	} else if strings.Contains(s, "特定非営利活動法人") {
		return HoujinKakuHieiri
	}

	panic("法人格が見つかりませんでした")
}

type Houjin struct {
	content            string
	CreatedAt          time.Time
	HoujinNumber       string
	HoujinType         HoujinkakuType
	CompanyName        string
	CompanyAddress     string
	Koukoku            string
	CompanyCreatedDate string
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
		fmt.Printf("日時: %v\n", h.CreatedAt)
	} else {
		return fmt.Errorf("日付と時刻が見つかりませんでした")
	}
	return nil
}

func (h *Houjin) ReadHoujinNumber() error {
	// 正規表現パターン: 全角数字で構成された法人番号
	pattern := "([０-９]{1,4}－[０-９]{1,2}－[０-９]{1,6})"
	regex := regexp.MustCompile(pattern)

	matches := regex.FindStringSubmatch(h.content)
	if len(matches) > 0 {
		h.HoujinNumber = zenkakuToHankaku(matches[1])
		fmt.Printf("法人番号: %s\n", h.HoujinNumber)
	} else {
		return fmt.Errorf("法人番号が見つかりませんでした")
	}
	return nil
}

func (h *Houjin) ReadCompanyName() error {
	// 商号に利用できる文字
	// https://www.moj.go.jp/MINJI/minji44.html
	pattern := fmt.Sprintf(`(商　*号|名　*称)　*│　*([%s]+)`, ZenkakuStringPattern)
	regex := regexp.MustCompile(pattern)

	// 抽出された会社名を表示
	matches := regex.FindStringSubmatch(h.content)
	if len(matches) > 0 {
		h.CompanyName = zenkakuToHankaku(strings.TrimSpace(matches[2]))
		fmt.Printf("会社名: %s\n", h.CompanyName)
		h.HoujinType = FindHoujinKaku(h.CompanyName)
	} else {
		return fmt.Errorf("会社名が見つかりませんでした。")
	}
	return nil
}

func (h *Houjin) ReadCompanyAddress() error {
	pattern := fmt.Sprintf(`(本　*店|主たる事務所)　*│　*([%s]+)`, ZenkakuStringPattern)
	regex := regexp.MustCompile(pattern)

	matches := regex.FindStringSubmatch(h.content)
	if len(matches) > 0 {
		h.CompanyAddress = zenkakuToHankaku(strings.TrimSpace(matches[2]))
		fmt.Printf("会社住所: %s\n", h.CompanyAddress)
	} else {
		return fmt.Errorf("会社住所が見つかりませんでした。")
	}
	return nil
}

func (h *Houjin) ReadKoukoku() error {
	noKoukokuList := []HoujinkakuType{
		HoujinKakuKyodou,
		HoujinKakuRoudou,
		HoujinKakuNPO,
		HoujinKakuSihoshosi,
		HoujinKakuHieiri,
		HoujinKakuSinrin,
		HoujinKakuIryo,
		HoujinKakuShokoukai,
		HoujinKakuZeirishi,
		HoujinKakuIppanShadan,
		HoujinKakuShakaifukusi,
		HoujinKakuTokutei,
		HoujinKakuSinyou,
		HoujinKakuIppanZaisan,
		HoujinKakuIppanZaidan,
		HoujinKakuNouji,
		HoujinKakuSeikatuEisei,
	}

	for _, v := range noKoukokuList {
		if h.HoujinType == v {
			return nil
		}
	}

	pattern := fmt.Sprintf(`公告をする方法　*│　*([%s]+)`, ZenkakuStringPattern)
	regex := regexp.MustCompile(pattern)

	matches := regex.FindStringSubmatch(h.content)
	if len(matches) > 0 {
		h.Koukoku = strings.TrimSpace(matches[1])
		fmt.Printf("公告をする方法: %s\n", h.Koukoku)
	} else {
		return fmt.Errorf("公告をする方法が見つかりませんでした。")
	}
	return nil
}

func (h *Houjin) ReadCompanyCreatedDate() error {
	pattern := fmt.Sprintf(`(会社|法人)成立の年月日　*│　*([%s]+)`, ZenkakuStringPattern)
	regex := regexp.MustCompile(pattern)

	matches := regex.FindStringSubmatch(h.content)
	if len(matches) > 0 {
		h.CompanyCreatedDate = zenkakuToHankaku(strings.TrimSpace(matches[2]))
		fmt.Printf("法人成立の年月日: %s\n", h.CompanyCreatedDate)
	} else {
		return fmt.Errorf("法人成立の年月日が見つかりませんでした。")
	}
	return nil
}

func Extract(content string) (Houjin, error) {
	houjin := Houjin{content: content}
	err := houjin.ReadCreatedAt()
	if err != nil {
		panic(err)
	}
	err = houjin.ReadHoujinNumber()
	if err != nil {
		panic(err)
	}
	err = houjin.ReadCompanyName()
	if err != nil {
		panic(err)
	}

	err = houjin.ReadCompanyAddress()
	if err != nil {
		panic(err)
	}

	err = houjin.ReadKoukoku()
	if err != nil {
		panic(err)
	}
	err = houjin.ReadCompanyCreatedDate()
	if err != nil {
		panic(err)
	}
	return houjin, nil
}
