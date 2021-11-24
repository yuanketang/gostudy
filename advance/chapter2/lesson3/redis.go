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

	_, err := db.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	defer db.Close()

	// 字符串
	log.Println("STRING测试..............")
	db.Set("KEY_CACHE", "test", time.Second*60)

	data, err := db.Get("KEY_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("data:", data)

	// Hash
	log.Println("Hash测试..............")
	db.HSet("HASH_CACHE", "name", "zhangsan")
	db.HSet("HASH_CACHE", "age", "20")
	total, err := db.HLen("HASH_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("total:", total)
	values, err := db.HGetAll("HASH_CACHE").Result()
	if err != nil {
		log.Println(err)
	}
	for key, val := range values {
		log.Println(key, val)
	}

	// List
	log.Println("LIST测试..............")
	db.LPush("LIST_CACHE", 1, 2, 3, 4, 5, 6)
	total, err = db.LLen("LIST_CACHE").Result()
	if err != nil {
		log.Println(err)
	}
	log.Println("total:", total)

	log.Println("模拟先进后出........")
	for i := 0; i < 6; i++ {
		data, err := db.RPop("LIST_CACHE").Result()
		if err != nil {
			log.Println(err)
		}
		log.Println("data:", data)
	}

	db.LPush("LIST_CACHE", 1, 2, 3, 4, 5, 6)
	log.Println("模拟后进先出........")
	for i := 0; i < 6; i++ {
		data, err := db.LPop("LIST_CACHE").Result()
		if err != nil {
			log.Println(err)
		}
		log.Println("data:", data)
	}

	// Set
	log.Println("SET测试..............")

	db.SAdd("SET_LIST", 1, 2, 3, 4, 5, 6)
	total, err = db.SCard("SET_LIST").Result()
	if err != nil {
		log.Println(err)
	}
	log.Println("total:", total)

	members, err := db.SMembers("SET_CACHE").Result()
	if err != nil {
		log.Println(err)
	}
	for _, val := range members {
		log.Println("data:", val)
	}

	// Ordered Set
	log.Println("ORDERED SET测试..............")
	scoresOfStudents := []redis.Z{
		{Score: 60.0, Member: "xiaoming"},
		{Score: 75.0, Member: "xiaohong"},
		{Score: 82.0, Member: "xiaoli"},
		{Score: 99.0, Member: "xiaosun"},
	}
	_, err = db.ZAdd("ORDERED_SET_CACHE", scoresOfStudents...).Result()
	if err != nil {
		log.Fatal(err)
	}
	total, err = db.ZCount("ORDERED_SET_CACHE", "-inf", "+inf").Result()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("total:", total)

	log.Println("按分数由低到高排序")
	result, err := db.ZRangeByScore("ORDERED_SET_CACHE", redis.ZRangeBy{
		Min: "70",
		Max: "100",
	}).Result()
	if err != nil {
		log.Fatal(err)
	}
	for _, ele := range result {
		log.Println("ele:", ele)
	}

	log.Println("按分数由高到低排序")
	result, err = db.ZRevRangeByScore("ORDERED_SET_CACHE", redis.ZRangeBy{
		Min: "70",
		Max: "100",
	}).Result()
	if err != nil {
		log.Fatal(err)
	}
	for _, ele := range result {
		log.Println("ele:", ele)
	}

	// Pipeline 管道
	log.Println("PIPELINE测试...........")
	var cmd *redis.IntCmd
	_, err = db.Pipelined(func(pipe redis.Pipeliner) error {
		cmd = pipe.Incr("INCR_CACHE")
		pipe.Expire("INCR_CACHE", time.Second*10)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("PIPELINE:", cmd.Val())

	// 事务
	log.Println("事务测试...........")
	_, err = db.TxPipelined(func(pipe redis.Pipeliner) error {
		cmd = pipe.Incr("TX_INCR_CACHE")
		pipe.Expire("TX_INCR_CACHE", time.Second*10)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TX PIPELINE:", cmd.Val())

	// 发布/订阅者模式
	log.Println("事务测试...........")
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
