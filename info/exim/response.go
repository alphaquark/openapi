package exim

type Response struct {
	RESULT          int    `json:"result"`          // 조회 결과 (1 : 성공, 2 : DATA코드 오류, 3 : 인증코드 오류, 4 : 일일제한횟수 마감)
	CUR_UNIT        string `json:"cur_unit"`        // 통화 코드
	CUR_NM          string `json:"cur_nm"`          // 국가/통화명
	TTB             string `json:"ttb"`             // 전신환(송금) 받을 때
	TTS             string `json:"tts"`             // 전신환(송금) 보낼 때
	DEAL_BAS_R      string `json:"deal_bas_r"`      // 매매 기준율
	BKPR            string `json:"dkpr"`            // 장부가격
	YY_EFEE_R       string `json:"yy_efee_r"`       // 년환가료율
	TEN_DD_EFEE_R   string `json:"ten_dd_efee_r"`   // 10일환가료율
	KFTC_DEAL_BAS_R string `json:"kftc_deal_bas_r"` // 서울외국환중개 매매 기준율
	KFTC_BKPR       string `json:"kftc_bkpr"`       // 서울외국환중개 장부가격
}
