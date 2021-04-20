package redis

//可靠队列
type ReliableQueue struct {
	source, dest string
}

func NewReliableQueue(key string) ReliableQueue {
	new := ReliableQueue{
		source: key,
		dest:   key + ":bak",
	}
	//初始化时检测是否有备份消息
	if count := global.REDIS.LLen(new.dest).Val(); count > 0 {
		for {
			result := global.REDIS.RPopLPush(new.dest, new.source)
			if len(result.Val()) < 1 || result.Err() != nil {
				break
			}
		}
	}
	return new
}

func (r ReliableQueue) Push(msg ...interface{}) error {
	return global.REDIS.LPush(r.source, msg...).Err()
}

func (r ReliableQueue) Subscribe(f func(val string)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				global.LOG.Error("ReliableQueue Subscribe", zap.Any("key", r.source), zap.Any("panic", err))
			}
		}()
		errCount := 0
		for {
			result := global.REDIS.BRPopLPush(r.source, r.dest, 0)
			err := result.Err()
			if err != nil {
				global.LOG.Error("ReliableQueue Subscribe", zap.Any("key", r.source), zap.Error(err))
				errCount++
				if errCount > 3 {
					return
				}
				continue
			}
			f(result.Val())
			//删除备份
			global.REDIS.LRem(r.dest, 1, result.Val())
			if err != nil {
				global.LOG.Error("ReliableQueue Subscribe", zap.Any("key", r.source), zap.Error(err))
				errCount++
				if errCount > 3 {
					return
				}
				continue
			}
		}
	}()
}
