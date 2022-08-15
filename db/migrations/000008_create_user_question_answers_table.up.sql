CREATE TABLE IF NOT EXISTS user_question_answers(
    id varchar(255) PRIMARY KEY,
    user_question_id varchar(255) NULL,
    answer_id varchar(255) NULL
);