CREATE TABLE IF NOT EXISTS user_questions(
    id varchar(255) PRIMARY KEY,
    user_id varchar(255) NULL,
    question_id varchar(255) NULL,
    status varchar(20) NULL
);