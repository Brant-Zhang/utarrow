package pattern

import "fmt"

type WeekReport interface {
	showReport()
}

type RealReport struct {
	name    string
	content string
}

func (this *RealReport) showReport() {
	//load the content based on name
	this.content = "hello world again!"
	fmt.Println(this.content)
}

type ReportProxy struct {
	userid int
	name   string
}

func (this *ReportProxy) showReport() {
	if !this.checkAccess() {
		return
	}
	r := &RealReport{name: this.name}
	r.showReport()
}

func (this *ReportProxy) checkAccess() bool {
	fmt.Println("user:", this.userid)
	return true
}
