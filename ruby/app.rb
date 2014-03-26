require 'json'
require 'redis'
require 'sinatra'

set :port, 3001

redis = Redis.new

REDIS_LIST = 'test_list'

get '/' do
  length = redis.lpush(REDIS_LIST, params['test'])
  length.to_s
end
