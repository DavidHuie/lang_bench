var redis	= require("redis"),
    express	= require("express");

var app		= express(),
    redisClient	= redis.createClient();

var redisKey = 'test_list';

app.get('/', function(req, res){
  var value = req.query.test;
  redisClient.lpush(redisKey, value, function(err, pushResponse) {
    res.send(String(pushResponse));
  });
});

app.listen(3001);
