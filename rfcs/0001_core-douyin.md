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
create database tiktok;
use tiktok;

create table user
(
    id          int auto_increment                  primary key,
    username    varchar(32)                         not null,
    password    varchar(32)                         not null,
    enable      tinyint   default 1                 null,
    login_time  datetime  default CURRENT_TIMESTAMP null,
    create_time timestamp default CURRENT_TIMESTAMP null
);

create table chat
(
    id         int auto_increment                 primary key,
    msg        text                               not null,
    sender     int                                not null,
    receiver   int                                not null,
    createtime datetime default CURRENT_TIMESTAMP not null,
    constraint chat_user_id_fk                    foreign key (sender) references user (id),
    constraint chat_user_id_fk_2                  foreign key (receiver) references user (id)
);

create index chat_receiver_sender_index
    on chat (receiver, sender);

create index chat_sender_receiver_index
    on chat (sender, receiver);

create table follow
(
    id      int auto_increment            primary key,
    user_id int               null,
    fun_id  int               not null,
    removed tinyint default 0 not null,
    msg     text              null,
    constraint follow_user_id2fun_fk_2    foreign key (fun_id) references user (id),
    constraint follow_user_id2user_fk     foreign key (user_id) references user (id)
);

create index follow_fun_id_removed_index
    on follow (fun_id, removed);

create index follow_user_id_removed_index
    on follow (user_id, removed);

create index user_username_enable_index
    on user (username, enable);

create table video
(
    id        int auto_increment primary key,
    author_id int                not null,
    play_url  varchar(32)        not null,
    cover_url varchar(32)        not null,
    time      int                not null,
    title     varchar(128)       not null,
    removed   tinyint default 0  not null,
    constraint video_user_id_fk  foreign key (author_id) references user (id)
);

create table comment
(
    id          int auto_increment                 primary key,
    user_id     int                                not null,
    video_id    int                                not null,
    create_time datetime default CURRENT_TIMESTAMP not null,
    removed     tinyint  default 0                 not null,
    deleted     tinyint  default 0                 not null,
    content     text                               not null,
    constraint comment_user_id_fk                  foreign key (user_id) references user (id),
    constraint comment_video_id_fk                 foreign key (video_id) references video (id)
);

create index comment_video_id_removed_create_time_index
    on comment (video_id, removed, create_time);

create table favorite
(
    id       int auto_increment      primary key,
    video_id int                     not null,
    user_id  int                     not null,
    removed  tinyint default 0       not null,
    constraint favorite_user_id_fk   foreign key (user_id) references user (id),
    constraint favorite_video_id_fk  foreign key (video_id) references video (id)
);

create index favorite_user_id_removed_index
    on favorite (user_id, removed);

create index favorite_video_id_removed_index
    on favorite (video_id, removed);

create index video_author_id_removed_index
    on video (author_id, removed);

create index video_time_removed_index
    on video (time, removed);


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
  