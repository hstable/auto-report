package model

type ReportData struct {
	Bgsj           string `json:"bgsj"`
	Bj             string `json:"bj"`
	Brfsgktt       string `json:"brfsgktt"`
	Brfsgkttjssj   string `json:"brfsgkttjssj"`
	Brfsgkttsfjs   string `json:"brfsgkttsfjs"`
	Brfsgkttsj     string `json:"brfsgkttsj"`
	Brjyyymc       string `json:"brjyyymc"`
	Brsfjy         string `json:"brsfjy"`
	Brzdjlbz       string `json:"brzdjlbz"`
	Brzdjlm        string `json:"brzdjlm"`
	Brzdjlmc       string `json:"brzdjlmc"`
	Brzgtw         string `json:"brzgtw"`
	Bzrdh          string `json:"bzrdh"`
	Bzrxm          string `json:"bzrxm"`
	Cjsj           string `json:"cjsj"`
	Dqszd          string `json:"dqszd"`
	Dqszdqu        string `json:"dqszdqu"`
	Dqszdqumc      string `json:"dqszdqumc"`
	Dqszdsheng     string `json:"dqszdsheng"`
	Dqszdshengmc   string `json:"dqszdshengmc"`
	Dqszdshi       string `json:"dqszdshi"`
	Dqszdshimc     string `json:"dqszdshimc"`
	Dqztbz         string `json:"dqztbz"`
	Dqztm          string `json:"dqztm"`
	Dqztmc         string `json:"dqztmc"`
	Dsdh           string `json:"dsdh"`
	Dsxm           string `json:"dsxm"`
	Dtid           string `json:"dtid"`
	Fdyxgr         string `json:"fdyxgr"`
	Fdyxgsj        string `json:"fdyxgsj"`
	Gnxxdz         string `json:"gnxxdz"`
	Gpsxx          string `json:"gpsxx"`
	Gpsxxip        string `json:"gpsxxip"`
	Hwcs           string `json:"hwcs"`
	Hwgj           string `json:"hwgj"`
	Hwxxdz         string `json:"hwxxdz"`
	Id             string `json:"id"`
	Jsfsgktt       string `json:"jsfsgktt"`
	Jsfsgkttjssj   string `json:"jsfsgkttjssj"`
	Jsfsgkttsfjs   string `json:"jsfsgkttsfjs"`
	Jsfsgkttsj     string `json:"jsfsgkttsj"`
	Jsjyyymc       string `json:"jsjyyymc"`
	Jssfjy         string `json:"jssfjy"`
	Jszdjlbz       string `json:"jszdjlbz"`
	Jszdjlm        string `json:"jszdjlm"`
	Jszdjlmc       string `json:"jszdjlmc"`
	Jszgtw         string `json:"jszgtw"`
	Lat            string `json:"lat"`
	Lng            string `json:"lng"`
	Nj             string `json:"nj"`
	Qrr            string `json:"qrr"`
	Qrsj           string `json:"qrsj"`
	Qtbgsx         string `json:"qtbgsx"`
	Rq             string `json:"rq"`
	Sffwwhhb       string `json:"sffwwhhb"`
	Sfjcqthbwhry   string `json:"sfjcqthbwhry"`
	Sfjcqthbwhrybz string `json:"sfjcqthbwhrybz"`
	Sfjdwhhbry     string `json:"sfjdwhhbry"`
	Sfkxg          string `json:"sfkxg"`
	Sftjwhjhb      string `json:"sftjwhjhb"`
	Sftzrychbwhhl  string `json:"sftzrychbwhhl"`
	Ssfdy          string `json:"ssfdy"`
	Stzkm          string `json:"stzkm"`
	Stzkmc         string `json:"stzkmc"`
	Tccx           string `json:"tccx"`
	Tchbcc         string `json:"tchbcc"`
	Tcjcms         string `json:"tcjcms"`
	Tcjtfs         string `json:"tcjtfs"`
	Tcjtfsbz       string `json:"tcjtfsbz"`
	Tcjtfsmc       string `json:"tcjtfsmc"`
	Tcyhbwhrysfjc  string `json:"tcyhbwhrysfjc"`
	Tczwh          string `json:"tczwh"`
	Tjsj           string `json:"tjsj"`
	Xdm            string `json:"xdm"`
	Xgmsxgr        string `json:"xgmsxgr"`
	Xgmsxgsj       string `json:"xgmsxgsj"`
	Xh             string `json:"xh"`
	Xslb           string `json:"xslb"`
	Xydm           string `json:"xydm"`
	Zt             string `json:"zt"`
	Zydm           string `json:"zydm"`
}

type Data struct {
	Data []ReportData `json:"data"`
}

type ResultData struct {
	IsSuccess bool   `json:"isSuccess"`
	Module    Data   `json:"module"`
	Msg       string `json:"msg"`
}
