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



![System Flow (20)](https://user-images.githubusercontent.com/60891544/171760652-6ff90978-ea63-41a2-a8bc-54c8832d3bf4.png)
![System Flow (19)](https://user-images.githubusercontent.com/60891544/171760656-d4346741-e0cb-42c7-9323-73d21a264d96.png)
![System Flow (18)](https://user-images.githubusercontent.com/60891544/171760657-c3274cf5-f636-48a8-a0ba-c48430b14971.png)
![System Flow (17)](https://user-images.githubusercontent.com/60891544/171760660-c74bec5b-ee93-4998-bd1c-fd488e293587.png)
![System Flow (16)](https://user-images.githubusercontent.com/60891544/171760665-a9999a37-6464-48f8-8856-5ec9d483a64a.png)
