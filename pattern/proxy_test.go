package pattern

import "testing"

func TestProxy(t *testing.T) {
	article := "baidu_weekly.docs"
	r := &ReportProxy{userid: 10, name: article}
	r.showReport()
}
