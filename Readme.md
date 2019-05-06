YOUR CONTRIBUTION IS VERY WELCOMED!!
PLEASE GIVE ME PULL REQUEST.

It is a grpc server demo.

this server contains APIs to set contacter info, and also to get, modify and delete etc.

mysql DB is used to store contacter data.
Redis will be used to store index.

TechDesign 1:
1)create a map and load contact info to map;
2)provide reload api to manually reload data from db to map. auto-reload every day;
pros: data can realtime and gurranteed.
cons: call db more often; slow; unsafe;

TechDesign 2: 
1) indexing; DB binlog change-->notify to index into redis. 
2) Search: request to search from redis-->get ids and then get contacter from db. 
pros: 1)faster, do not call db directly; 2)after get id, can rank with the id according to the request params; 
cons: 1)still neeed to call db if need extra info; 2) data is not garrateed to be updated.