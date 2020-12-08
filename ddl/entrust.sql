CREATE ROLE entrust_role;

CREATE SCHEMA IF NOT EXISTS entrust

    -- 委托记录
    CREATE TABLE IF NOT EXISTS entrust
    (
        id                   INTEGER PRIMARY KEY DEFAULT NEXTVAL("entrust_id"), -- 委托ID
        market_name          VARCHAR(50)     NOT NULL,                          -- 市场名称
        market_id            INT             NOT NULL,                          -- 市场ID
        user_id              INT             NOT NULL,
        unit_price           NUMERIC(30, 15) NOT NULL,                          -- 下单价格
        numbers              NUMERIC(30, 15) NOT NULL,
        total_money          NUMERIC(30, 15) NOT NULL,
        complete_numbers     NUMERIC(30, 15) NOT NULL,
        complete_total_money NUMERIC(30, 15) NOT NULL,
        types                INT                 DEFAULT 0,
        status               INT                 DEFAULT 3,
        acct_type            INT                 DEFAULT 1,
        create_at            TIMESTAMP           DEFAULT CURRENT_TIMESTAMP AT TIME ZONE DEFAULT,
        cancel_at            TIMESTAMP           DEFAULT NULL AT TIME ZONE DEFAULT,
        update_at            TIMESTAMP           DEFAULT NULL AT TIME ZONE DEFAULT
    ) TABLESPACE "entrust";

CREATE SEQUENCE "entrust_id" OWNED BY entrust.entrust.id; -- entrustId 序列

ALTER TABLE entrust.entrust
    OWNER TO entrust_role;
;
