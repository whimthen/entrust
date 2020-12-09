CREATE ROLE entrust_role;

CREATE SCHEMA IF NOT EXISTS entrust_schema

    -- 委托记录
    CREATE TABLE IF NOT EXISTS entrust
    (
        id                   SERIAL8 PRIMARY KEY,                 -- 委托ID
        market_name          VARCHAR(50)     NOT NULL,            -- 市场名称
        market_id            INT             NOT NULL,            -- 市场ID
        user_id              INT             NOT NULL,            -- 用户ID
        price                NUMERIC(30, 15) NOT NULL,            -- 下单价格
        numbers              NUMERIC(30, 15) NOT NULL,            -- 下单数量
        total_money          NUMERIC(30, 15) NOT NULL,            -- 下单总额(价格 * 数量)
        complete_numbers     NUMERIC(30, 15) NOT NULL,            -- 完成数量
        complete_total_money NUMERIC(30, 15) NOT NULL,            -- 完成总额
        types                INT       DEFAULT 0,                 -- 下单类型
        status               INT       DEFAULT 3,                 -- 委托单状态
        acct_type            INT       DEFAULT 1,                 -- 账户类型
        create_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间
        cancel_at            TIMESTAMP DEFAULT NULL,              -- 取消时间
        update_at            TIMESTAMP DEFAULT NULL,              -- 更新时间
        complete_at          TIMESTAMP DEFAULT NULL               -- 完成时间
    );

CREATE INDEX idx_user_id ON entrust_schema.entrust (user_id);
CREATE INDEX idx_types ON entrust_schema.entrust (types);
CREATE INDEX idx_create_at ON entrust_schema.entrust (create_at);
CREATE INDEX idx_status ON entrust_schema.entrust (status);
CREATE INDEX idx_userId_status_createAt ON entrust_schema.entrust (user_id, status, create_at);

COMMENT ON TABLE entrust_schema.entrust IS '委托记录';

ALTER TABLE entrust_schema.entrust
    OWNER TO entrust_role;
;
