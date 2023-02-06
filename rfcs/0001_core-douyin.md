# Features

- Feature Name: core-reservation

- Start Date: 2023-2-4 14:34:05  

## Sumary

`#Seperate #Design`

- A project which based on **micro-service** architechture and with the goal of implenting the service of douyin/titok

- Currently with these 3 main content, but a lot things have not been thought enough:
    1. User infomation management and permission validation
    2. Video stream, feed service
    3. Relations between users (followers/followes)

- Likese
    1. API intergation
        - ~~Do we need a APIs gateway?~~
        - If we use `kerbernets`, it seems that no need service-registration.  Kubernets have done these things by `coredns` for us;  
        - The services could be call together by the `service-name`, like a `register` service call database with a domain of `database.k8s.svc`(unclearly now);
        - For now, all services will be design to deploy in one kubernets cluster, if some of services **not in one cluster**, still need registeration-center between servies and cluster, clusters and clusters;

    2. Devide services into more **layers and slices**
        - 3 modules of design are simple (for archtiecture), but also too bulky for development, debugging, etc;
        - We need to make each single service as  simple and single responsibilty as possible;
        - To split `Permission Validation` out as a single **Micro** service

    3. What about the middleware?
        > Middleware means the software that make connection of sofwares between system and user.
        - JWT
        - Access flow limitaion
        - MQ (Kafka)
        - Service real-time monitor (grafana, metrics, prometheus)
        - Cache (Redis)

## Motivation

- Our goal is to get the better(or the best) rankings.
- Get the greate experiences about golang, micro-scervice and kubernets , etc.

## Guide-level explannation

- The message protolcal will be `gRPC`. (Or zRPC?) So we will use the IDL to define our interfaces.
- The database we will use is `MySQL`. It need a good schema for the subsequent development.
- With the `interfaces` that provide by ByteYouthCamp, we don't need to define more IDL  files. (maybe need add cutom interface)
- `Redis` that make cache
- `Kafka` as message queue

>
### DataBase schema

```sql
create database titok;

create table titok.user{
    id int(64) AUTO_INCREMENT PRIMARY KEY,
    username varchar(32) NOT NULL,
    password varchar(32) NOT NULL,
    resource_id varchar(64) NOT NULL,
    login_time timestamp NOT NULL,
    create_time timestamp NOT NULL,
    removed tinyint(4)  NOT NULL,
    deleted  tinyint(4) NOT NULL,
}

create table titok.follow{
    ...
}
...

```

### Service interface
>
> The interface and IDL example refs: <https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof>

~~#### User~~

### Caches

1. ~~Follower and followee list~~
2. Like ~~list~~ count of video
3. ~~User info~~
4. ~~Published list (refence)~~
5. Comment list of video

### Message Queue

#### TOPICS

- On give/post a:
  - Like
  - Chat
  - ~~Publish~~
  - Comment
  - Follow/Unfollow
  