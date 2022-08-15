CREATE TABLE IF NOT EXISTS answers(
    id varchar(255) PRIMARY KEY,
    question_id varchar(255) NULL,
    answer_desc varchar(1000) NULL,
    answer_score integer NULL
);