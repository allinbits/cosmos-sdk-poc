# ORM
ORM brings basic functionalities over objects saved into KVStores.

It's divided in two different stores.

# Objects

Objects store is used to save objects.

Objects are identified by a primary key and a type.

The type is represented by the protobuf message type. 

Primary key is marked during object registration by the json name of the field.

The object is then stored in the following way:

```
saveKey := protobuf.MessageName(object) + / + schema.GetPrimaryKey(object)
```

# Indexes
Indexes store is used to save protobuf message fields (as protobuf json field name) as secondary keys.

The following is the secondary key structure:

```
indexerKey := <index_prefix><object_type_name_length as 8bytes little endian bytes><object_type_name>
```