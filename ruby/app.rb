require 'json'
require 'redis'
require 'sinatra'

redis = Redis.new

REDIS_LIST = 'test_list'

get '/' do
  payload = JSON(params)
  redis.lpush(REDIS_LIST, payload)
  redis.llen(REDIS_LIST).to_s
end
