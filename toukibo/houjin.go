package toukibo

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type HoujinkakuType string

const (
	HoujinKakuUnknown      HoujinkakuType = "不明"
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
	HoujinKakuShukyo       HoujinkakuType = "宗教法人"
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
	return HoujinKakuUnknown
}

type Houjin struct {
	Content            string
	Parts              []string
	CreatedAt          time.Time
	HoujinNumber       string
	HoujinType         HoujinkakuType
	CompanyName        string
	CompanyAddress     string
	Sihonkin           string
	Koukoku            string
	CompanyCreatedDate string
	ToukiJiko          string
}

func NewHoujinFromToukibo(tc ToukiboContent) *Houjin {
	return &Houjin{
		Content:        tc.Content,
		Parts:          tc.Parts,
		CreatedAt:      tc.Header.CreatedAt,
		CompanyName:    tc.Header.CompanyName,
		CompanyAddress: tc.Header.CompanyAddress,
	}
}

func (h *Houjin) String() string {
	return fmt.Sprintf("法人番号: %s\n法人格: %s\n商号: %s\n住所: %s\n資本金: %s\n設立日: %s\n登記事項: %s\n",
		h.HoujinNumber,
		h.HoujinType,
		h.CompanyName,
		h.CompanyAddress,
		h.Sihonkin,
		h.CompanyCreatedDate,
		h.ToukiJiko)
}

func (h *Houjin) ReadHoujinNumber() error {
	// 正規表現パターン: 全角数字で構成された法人番号
	pattern := "([０-９]{1,4}－[０-９]{1,2}－[０-９]{1,6})"
	regex := regexp.MustCompile(pattern)

	matches := regex.FindStringSubmatch(h.Parts[0])
	if len(matches) > 0 {
		h.HoujinNumber = zenkakuToHankaku(matches[1])
	} else {
		return fmt.Errorf("法人番号が見つかりませんでした")
	}
	return nil
}

func contains(target HoujinkakuType, list []HoujinkakuType) bool {
	for _, t := range list {
		if t == target {
			return true
		}
	}
	return false
}
func (h *Houjin) ReadKoukoku() error {
	koukokuList := []HoujinkakuType{
		HoujinKakuKabusiki,
		HoujinKakuYugen,
		HoujinKakuGoudou,
		HoujinKakuGousi,
		HoujinKakuGoumei,
	}
	if !contains(h.HoujinType, koukokuList) {
		return nil
	}

	pattern := fmt.Sprintf(`公告をする方法　*│　*([%s]+)`, ZenkakuStringPattern)
	regex := regexp.MustCompile(pattern)

	matches := regex.FindStringSubmatch(h.Content)
	if len(matches) > 0 {
		h.Koukoku = strings.TrimSpace(matches[1])
	} else {
		return fmt.Errorf("公告をする方法が見つかりませんでした。")
	}
	return nil
}

func (h *Houjin) ReadCompanyCreatedDate() error {
	pattern := fmt.Sprintf(`(会社|法人)成立の年月日　*│　*([%s]+)`, ZenkakuStringPattern)
	regex := regexp.MustCompile(pattern)

	matches := regex.FindStringSubmatch(h.Content)
	if len(matches) > 0 {
		h.CompanyCreatedDate = zenkakuToHankaku(strings.TrimSpace(matches[2]))
	} else {
		return fmt.Errorf("法人成立の年月日が見つかりませんでした。")
	}
	return nil
}

func (h *Houjin) CheckShukyoHoujin() bool {
	if h.HoujinType == HoujinKakuUnknown && strings.Contains(h.Content, "宗教法人") {
		h.HoujinType = HoujinKakuShukyo
		return true
	}
	return false
}

func (h *Houjin) ReadToukiJikou() error {
	s := h.Parts[len(h.Parts)-1]
	pattern := "┃登記記録に関する│　*"
	cleanedText := trimPattern(s, pattern)

	pattern = "　*┃ ┃事項　　　　　　│　*"
	cleanedText = trimPattern(cleanedText, pattern)

	//　改行に対応
	pattern = "　*┃ ┃　　　　　　　　│　*"
	cleanedText = trimPattern(cleanedText, pattern)

	// 登記事項の変更に対応
	pattern = "┃ ┃　　　　　　　　├─────────────────────────────────────┨ ┃　　　　　　　　│"
	cleanedText = trimPattern(cleanedText, pattern)

	// 末尾の記号を削除
	cleanedText = strings.Replace(cleanedText, "┃", "", -1)
	h.ToukiJiko = zenkakuToHankaku(cleanedText)
	return nil
}

func (h *Houjin) Extract() error {
	err := h.ReadHoujinNumber()
	if err != nil {
		panic(err)
	}

	h.HoujinType = FindHoujinKaku(h.CompanyName)
	if h.HoujinType == HoujinKakuUnknown {
		check := h.CheckShukyoHoujin()
		if !check {
			return fmt.Errorf("法人格が不明です")
		}
	}

	err = h.ReadKoukoku()
	if err != nil {
		panic(err)
	}

	err = h.ReadCompanyCreatedDate()
	if err != nil {
		panic(err)
	}

	err = h.ReadToukiJikou()
	if err != nil {
		panic(err)
	}

	return nil
}
