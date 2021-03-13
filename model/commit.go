package model

type ModelData struct {
	Id string `json:"id"`
	Stzkm string `json:"stzkm"`
	Dqszd string `json:"dqszd"`
	Hwgj string `json:"hwgj"`
	Hwcs string `json:"hwcs"`
	Hwxxdz string `json:"hwxxdz"`
	Dqszdsheng string `json:"dqszdsheng"`
	Dqszdshi string `json:"dqszdshi"`
	Dqszdqu string `json:"dqszdqu"`
	Gnxxdz string `json:"gnxxdz"`
	Dqztm string `json:"dqztm"`
	Dqztbz string `json:"dqztbz"`
	Brfsgktt string `json:"brfsgktt"`
	Brzgtw string `json:"brzgtw"`
	Brsfjy string `json:"brsfjy"`
	Brjyyymc string `json:"brjyyymc"`
	Brzdjlm string `json:"brzdjlm"`
	Brzdjlbz string `json:"brzdjlbz"`
	Qtbgsx string `json:"qtbgsx"`
	Sffwwhhb string `json:"sffwwhhb"`
	Sftjwhjhb string `json:"sftjwhjhb"`
	Tcyhbwhrysfjc string `json:"tcyhbwhrysfjc"`
	Sftzrychbwhhl string `json:"sftzrychbwhhl"`
	Sfjdwhhbry string `json:"sfjdwhhbry"`
	Tcjtfs string `json:"tcjtfs"`
	Tchbcc string `json:"tchbcc"`
	Tccx string `json:"tccx"`
	Tczwh string `json:"tczwh"`
	Tcjcms string `json:"tcjcms"`
	Gpsxx string `json:"gpsxx"`
	Sfjcqthbwhry string `json:"sfjcqthbwhry"`
	Sfjcqthbwhrybz string `json:"sfjcqthbwhrybz"`
	Tcjtfsbz string `json:"tcjtfsbz"`
}

type ID struct {
	Id string `json:"id"`
}

type CommitResult struct {
	IsSuccess bool `json:"isSuccess"`
	Module string `json:"module"`
	Msg string `json:"msg"`
}

type CommitData struct {
	Model ModelData `json:"model"`
}
