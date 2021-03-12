package model

type ModelData struct {
	id string `json:"id"`
	stzkm string `json:"stzkm"`
	dqszd string `json:"dqszd"`
	hwgj string `json:"hwgj"`
	hwcs string `json:"hwcs"`
	hwxxdz string `json:"hwxxdz"`
	dqszdsheng string `json:"dqszdsheng"`
	dqszdshi string `json:"dqszdshi"`
	dqszdqu string `json:"dqszdqu"`
	gnxxdz string `json:"gnxxdz"`
	dqztm string `json:"dqztm"`
	dqztbz string `json:"dqztbz"`
	brfsgktt string `json:"brfsgktt"`
	brzgtw string `json:"brzgtw"`
	brsfjy string `json:"brsfjy"`
	brjyyymc string `json:"brjyyymc"`
	brzdjlm string `json:"brzdjlm"`
	brzdjlbz string `json:"brzdjlbz"`
	qtbgsx string `json:"qtbgsx"`
	sffwwhhb string `json:"sffwwhhb"`
	sftjwhjhb string `json:"sftjwhjhb"`
	tcyhbwhrysfjc string `json:"tcyhbwhrysfjc"`
	sftzrychbwhhl string `json:"sftzrychbwhhl"`
	sfjdwhhbry string `json:"sfjdwhhbry"`
	tcjtfs string `json:"tcjtfs"`
	tchbcc string `json:"tchbcc"`
	tccx string `json:"tccx"`
	tczwh string `json:"tczwh"`
	tcjcms string `json:"tcjcms"`
	gpsxx string `json:"gpsxx"`
	sfjcqthbwhry string `json:"sfjcqthbwhry"`
	sfjcqthbwhrybz string `json:"sfjcqthbwhrybz"`
	tcjtfsbz string `json:"tcjtfsbz"`
}

type CommitData struct {
	Model ModelData `json:"model"`
}
