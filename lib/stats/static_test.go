package stats

import(
	"testing"
)

var (
	st statsM
)

func init(){
	st=NewStats()
	st.AddKey("testkey")
}

func TestAdd(t *testing.T){
	Increment(st["testkey"])		
	st.PrintStats()
}

func TestAssign(t *testing.T){
	Assign(st["testkey"],199)		
	st.PrintStats()
}
