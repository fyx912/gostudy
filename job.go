package main 
import(
	"github.com/robfig/cron"
	"log"
)

type TestJob struct{}

func (this TestJob)Run() {
	log.Println("test job...")
}
func main(){
	i := 0
	c := cron.New()
	
	spec := "*/5 * * * * ?"
	c.AddFunc(spec,	func ()  {
		i++
		log.Println("cron runing:",i)
	})
	//添加job
	c.AddJob(spec,TestJob{})

	c.Start()
	defer c.Stop()

	select{}
}