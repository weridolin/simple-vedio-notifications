package main

import (
	"fmt"

	// schedulers "github.com/weridolin/simple-vedio-notifications/schedulers"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/schedulers"
	// httpConsumers "github.com/weridolin/simple-vedio-notifications/servers/http/consumers"
)

// type Context struct {
// 	tp     *schedulers.TickerPool
// 	config *configs.Config
// }

// func NewContext(tp *schedulers.TickerPool, config *configs.Config) *Context {
// 	return &Context{
// 		tp:     tp,
// 		config: config,
// 	}
// }

func TestCron() {
	c := cron.New()
	i := 1

	c.Start()

	EntryID, err := c.AddFunc("*/1 * * * *", func() {
		fmt.Println(time.Now(), "每分钟执行一次", i)
		i++
	})
	fmt.Println(time.Now(), EntryID, err)

	c.Stop()
	time.Sleep(time.Minute * 1)
	time.Sleep(time.Minute * 2)
}

func Setup() {
	//  DB迁移
	DB := database.GetDB()
	// AutoMigrate : 1 添加字段会自定添加  2.删除字段原来的会保存，新的记录会默认删除的字段为空
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&database.User{},
		&database.Scheduler{},
		&database.Task{},
		&database.EmailNotifier{},
		&database.EmailNotifierTask{},
		&database.SchedulerTask{},
	)
}

func main() {
	Setup()
	// ctx := context.Background()
	// storage := storage.NewStorage(ctx)
	// var info = []interface{}{{"name": "11"}}
	// storage.Save(info)
	sync := schedulers.NewSynchronizer()
	go sync.Start()
	// config := config.GetAppConfig()
	// for i := 0; i < config.EmailConsumerCount; i++ {
	// 	emailConsumer := consumers.NewEmailConsumer(tools.GetUUID())
	// 	go emailConsumer.Start()
	// }

	// http.Start()
	time.Sleep(time.Minute * 10)
}
