# 03-01-redis-cache

В предыдущем задании мы с вами практиковались с миграциями и БД. Сейчас попробуем написать простой кеш для наших данных.

В этом задании нам нужно будет поработать с mongodb и redis.
В этом шаблоне уже есть необходимые интерфейсы. Нужно сделать так, чтобы кеш принимал в себя экземпляр нашей базы данных и прозрачно кешировал данные. 
При вызове метода FindByCommand мы сначала должны найти данные в кеше и если их нет то забрать их их базы данных. 
Если их нет в базе данных то вернуть ошибку.
