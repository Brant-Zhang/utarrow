package redislock
import(
	"testing"
	"time"
)

var redylock *RedisLockEx

func init(){
	err := LoadLockRedis("")
	if err != nil {
		panic(err)
	}
	redylock = NewRedisLockEx("xxx.dislock",10)
}

func TestLock(t *testing.T){
	err:=redylock.LockOnce()
	if err!=nil{
		t.Fatal(err)
	}
	time.Sleep(1e9*60)

	//err=redylock.Unlock()
	//if err!=nil{
	//	t.Fatal(err)
	//}
	
}

