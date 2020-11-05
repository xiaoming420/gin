package redis

import (
	"fmt"
	"gin/pkg/setting"
	"github.com/go-redis/redis" 					// 实现了redis连接池
	"log"
	"time"
)

// 定义redis链接池
var client *redis.Client

// 初始化redis链接池
func init(){
	sec, err := setting.Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis': %v", err)
	}
	size,_ := sec.Key("RedisPoolSize").Int()
	client = redis.NewClient(&redis.Options{
		Addr:     sec.Key("Host").String(), 		// Redis地址
		Password: sec.Key("Password").String(),  	// Redis账号
		PoolSize: size,  		    				   	// Redis连接池大小
		MaxRetries: 3,              					// 最大重试次数
		IdleTimeout: 10 * time.Second,            		// 空闲链接超时时间
	})
	pong, err := client.Ping().Result()
	if err == redis.Nil {
		log.Fatal("Redis异常")
	} else if err != nil {
		log.Fatal("失败:", err)
	} else {
		log.Println("成功连接"+pong)
	}
}

// 向key的hash中添加元素field的值
func HashSet(key, field string, data interface{},exp int) {
	err := client.HSet(key, field, data)
	client.Expire(key,time.Duration(exp)*time.Second)
	if err != nil {

	}
}




// 通过key获取hash的元素值
func HashGet(key, field string) string {
	result := ""
	val, err := client.HGet(key, field).Result()
	if err == redis.Nil {
		return result
	}else if err != nil {

		return result
	}
	return val
}

// 通过key获取hash的元素值
func HashGetall(key string) map[string]string {
	result := make(map[string]string)
	val, err := client.HGetAll(key).Result()
	if err == redis.Nil {

		return result
	}else if err != nil {
		return result
	}
	return val
}


// 批量获取key的hash中对应多元素值
func BatchHashGet(key string, fields ...string) map[string]interface{} {
	resMap := make(map[string]interface{})
	for _, field := range fields {
		var result interface{}
		val, err := client.HGet(key, fmt.Sprintf("%s", field)).Result()
		if err == redis.Nil {

			resMap[field] = result
		}else if err != nil {

			resMap[field] = result
		}
		if val != "" {
			resMap[field] = val
		}else {
			resMap[field] = result
		}
	}
	return resMap
}

// 获取自增唯一ID
func Incr(key string) int {
	val, err := client.Incr(key).Result()
	if err != nil {

	}
	return int(val)
}

// 添加集合数据
func SetAdd(key, val string,exp int){
	client.SAdd(key, val)
	client.Expire(key,time.Duration(exp)*time.Second)
}

// 从集合中获取数据
func SetGet(key string)[]string{
	val, err := client.SMembers(key).Result()
	if err != nil{

	}
	return val
}

func Listlpush(key string,vals string,exp int)(int64) {
	val, err := client.LPush(key,vals).Result()
	client.Expire(key,time.Duration(exp)*time.Second)
	if err!=nil {

	}
	return val
}

func Listlrange(key string,exp int)([]string) {
	val, err := client.LRange(key,0,-1).Result()
	client.Expire(key,time.Duration(exp)*time.Second)
	if err!=nil {

	}
	return val
}

func HashDel(key string,data string) int64  {
	val, err := client.HDel(key,data).Result()
	if err!=nil {

	}
	return val
}

func ListPop(key string) string{
	val, err := client.LPop(key).Result()
	if err!=nil {

	}
	return val
}

func ListDel(uid,token_key string)int64{
	val,err:=client.LRem("tw_user_token:"+uid,0,token_key).Result()
	if err!=nil {

	}
	return val
}

func KeySet(key string,value interface{}, exp time.Duration) (string ,error){
	val,err := client.Set(key,value,exp).Result()
	if err != nil {

	}
	return val, err
}

//redis.KeySet("test",111112,time.Minute * 5)
func KeyGet(key string) (string ,error){
	val,err := client.Get(key).Result()
	if err != nil {
		return "",err
	}
	return val, err
}

func KeySetNX(key string,value interface{}, exp time.Duration) (bool ,error){
	val,err := client.SetNX(key,value,exp).Result()
	if err != nil {

	}
	return val, err
}