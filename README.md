# ATLANtis

> **DB used** : **Elasticsearch** as a value can be searched and indexed on the basis of mathching as well(we might have to find a word in a paragraph, or a word closest to a particular word). 
>> * Couldn't use SQL databases as the best case would be to traverse throught all entry and search for a word in a column.
>> * Didn't use MongoDB as search a value in multi_match is easy in elasticsearch and we get a ranking of data closest to the query.

> **Language Used** : **Golang** (highly scalable language,easy to learn and a fast and easy to learn language.

> **Error Tracing, Logs, track performance, Error alerts** : **Sentry** because of easy GoLang integration, complete context, real time updates.

> **Queue Service** : Used Kafka Services, because of producer and consumer configuration, partition and task that make it scalable.
>> * Didn't use RabbitMQ as it doesn't support multiple consumers.


## Overall Structure

![System Flow (14)](https://user-images.githubusercontent.com/60891544/171741548-84a2dc2f-ac7f-4dfd-a5f5-ff2ba92c79fd.png)

