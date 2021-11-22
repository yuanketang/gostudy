package main

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

var db *redis.Client

func init() {
	// 单机模式
	db = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// 主从模式
	//redis.NewFailoverClient(&redis.FailoverOptions{
	//	MasterName: "master",
	//	SentinelAddrs: []string{"192.168.1.2:6379", "192.168.1.3:6379"},
	//})

	// 集群模式
	//redis.NewClusterClient(&redis.ClusterOptions{
	//	Addrs: []string{"192.168.1.2:6379", "192.168.1.3:6379"},
	//})

	// 连接测试
	_, err := db.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	defer db.Close()

	// 字符串
	log.Println("KEY测试..............")
	db.Set("KEY_CACHE", "123", time.Second*10)
	data, err := db.Get("KEY_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)

	// hash
	log.Println("HASH测试..............")
	db.HSet("HASH_CACHE", "name", "zhangsan")
	db.HSet("HASH_CACHE", "age", "20")
	if err != nil {
		log.Fatal(err)
	}
	total, err := db.HLen("HASH_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("total", total)
	values, err := db.HGetAll("HASH_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range values {
		log.Println(key, value)
	}

	// list
	log.Println("LIST测试..............")
	db.LPush("LIST_CACHE", 1, 2, 3, 4, 5, 6)
	total, err = db.LLen("LIST_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("total", total)
	log.Println("模拟后进先出..........")
	for i := 0; i < 6; i++ {
		data, err = db.LPop("LIST_CACHE").Result()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("data", data)
	}

	db.LPush("LIST_CACHE", 1, 2, 3, 4, 5, 6)
	log.Println("模拟先进后出..........")
	for i := 0; i < 6; i++ {
		data, err = db.RPop("LIST_CACHE").Result()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("data", data)
	}

	// set
	log.Println("SET测试..............")
	_, err = db.SAdd("SET_CACHE", 1, 2, 3, 4, 5, 6).Result()
	if err != nil {
		log.Fatal(err)
	}
	total, err = db.SCard("SET_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("total", total)
	result, err := db.SMembers("SET_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	for ele := range result {
		log.Println(ele)
	}

	// ordered set
	log.Println("ORDERED SET测试..............")
	scoreOfStudents := []redis.Z{
		{Score: 65.0, Member: "xiaoming"},
		{Score: 75.0, Member: "xiaohong"},
		{Score: 85.0, Member: "xiaoli"},
		{Score: 95.0, Member: "xiaosun"},
	}
	_, err = db.ZAdd("ORDERED_SET_CACHE", scoreOfStudents...).Result()
	if err != nil {
		log.Fatal(err)
	}
	total, err = db.ZCount("ORDERED_SET_CACHE", "-inf", "+inf").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("total", total)
	log.Println("按分数由低到高")
	result, err = db.ZRangeByScore("ORDERED_SET_CACHE", redis.ZRangeBy{
		Min: "70",
		Max: "90",
	}).Result()
	if err != nil {
		log.Fatal(err)
	}
	for _, ele := range result {
		log.Println(ele)
	}
	log.Println("按分数由高到低")
	result, err = db.ZRevRangeByScore("ORDERED_SET_CACHE", redis.ZRangeBy{
		Min: "70",
		Max: "90",
	}).Result()
	if err != nil {
		log.Fatal(err)
	}
	for _, ele := range result {
		log.Println(ele)
	}

	// 使用管道一次发送多条命令
	log.Println("PIPELINE测试..............")
	var cmd *redis.IntCmd
	_, err = db.Pipelined(func(pipe redis.Pipeliner) error {
		cmd = pipe.Incr("INCR_CACHE")
		pipe.Expire("INCR_CACHE", time.Second*10)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("PIPELINE", cmd.Val())

	// 事务
	log.Println("事务测试..............")
	_, err = db.TxPipelined(func(pipe redis.Pipeliner) error {
		cmd = pipe.Incr("TX_INCR_CACHE")
		pipe.Expire("TX_INCR_CACHE", time.Second*10)
		return nil
	})
	log.Println("TX PIPELINE", cmd.Val())

	// 发布订阅模式
	log.Println("发布订阅模式............")
	sub := db.Subscribe("chatroom")
	for {
		msg, err := sub.ReceiveMessage()
		if err != nil {
			log.Fatal(err)
		}
		if msg != nil {
			log.Printf("从 [%s] 收到消息： %s\n", msg.Channel, msg.Payload)
		}
	}
}
