FROM redis

COPY init/redis.conf /usr/local/etc/redis/redis.conf
EXPOSE 6397
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]


