-- テーブル作成
DROP DATABASE IF EXISTS new_sample_db;
CREATE DATABASE sample_db;

-- データベースを選択
USE sample_db;

CREATE TABLE `category` (
    id INT NOT NULL AUTO_INCREMENT,
    obj_id VARCHAR(36) NOT NULL,
    name VARCHAR(20) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY idx_obj_id (obj_id)
);

-- 商品
CREATE TABLE `product` (
    id INT NOT NULL AUTO_INCREMENT,
    obj_id VARCHAR(36) NOT NULL,
    name VARCHAR(30) NOT NULL,
    price INT NOT NULL,
    category_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY idx_obj_id (obj_id),
    FOREIGN KEY category_fk (category_id) REFERENCES `category` (obj_id)
);

-- 商品カテゴリ
INSERT INTO `category` (obj_id,name) VALUES ('f52h3wf3-g467-jfy5-fg47-gf55gd37jjvs','文房具');
INSERT INTO `category` (obj_id,name) VALUES ('254gfdgh-bgdw-4678-scrr-grgs45gj3efg','雑貨');
INSERT INTO `category` (obj_id,name) VALUES ('jfg46hf3-f6u7-scgt-35uj-cervvrdg367b','パソコン周辺機器');
-- 商品
INSERT INTO `product` (obj_id,name,price,category_id) VALUES ('f52h3wf3-g467-jfy5-fg47-gf55gd37jjvs','水性ボールペン（黒）',120,'f52h3wf3-g467-jfy5-fg47-gf55gd37jjvs');
INSERT INTO `product` (obj_id,name,price,category_id) VALUES ('gdgerhfh-h7h4-fdv3-gfge-47g3y5hfh3g','観葉植物',120,'254gfdgh-bgdw-4678-scrr-grgs45gj3efg');
INSERT INTO `product` (obj_id,name,price,category_id) VALUES ('sd34yh5r-sfef-3257-sc35-ftgd35y6htg4','有線マウス',120,'jfg46hf3-f6u7-scgt-35uj-cervvrdg367b');
INSERT INTO `product` (obj_id,name,price,category_id) VALUES ('sd2g3gdg-g477-ff3g-df35-f46hfg4tegfg','無線キーボード',120,'jfg46hf3-f6u7-scgt-35uj-cervvrdg367b');
