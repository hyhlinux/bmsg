create table `message`
    -- --------------------------------------------------
    --  Table Structure for `bmsg/models.Message`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS "message" (
        "id" serial NOT NULL PRIMARY KEY,
        "from_user_id" text NOT NULL DEFAULT '' ,
        "to_user_id" text NOT NULL DEFAULT '' ,
        "created_at" timestamp with time zone,
        "update_at" timestamp with time zone,
        "title" text,
        "text" text,
        "is_delete" bool NOT NULL DEFAULT FALSE ,
        "status" integer
    );
    CREATE INDEX "message_from_user_id_to_user_id_status_is_delete" ON "message" ("from_user_id", "to_user_id", "status", "is_delete");


--- 用户站内信表
-- fromUserId int
-- toUserId
-- title
-- message
-- isDelete
-- status:READ UNREA  str
-- updateDate
-- createDate