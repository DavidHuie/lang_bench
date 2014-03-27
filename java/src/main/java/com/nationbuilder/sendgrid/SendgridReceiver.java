package com.nationbuilder.sendgrid;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;

public class SendgridReceiver extends HttpServlet {

	private static final long serialVersionUID = 4997832136717624098L;

	private static final String SENDGRID_QUEUE = "test_list";
	private static final String PARAMETER = "test";

	private static final JedisPool REDIS = new JedisPool("localhost", 6379);

	@Override
	protected void doGet(HttpServletRequest req, HttpServletResponse resp)
			throws ServletException, IOException {
		String data = req.getParameter(PARAMETER);
		Jedis redis = null;
		try {
			redis = REDIS.getResource();
			redis.lpush(SENDGRID_QUEUE, data);
		} finally {
			if (redis != null)
				REDIS.returnResource(redis);
		}
	}

}

