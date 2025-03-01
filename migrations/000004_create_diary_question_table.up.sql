CREATE TABLE diary_question
(
    diary_id    int references diaries (id) on update cascade on delete cascade,
    question_id int references questions (id) on update cascade on delete cascade,
    constraint diary_question_pkey primary key (question_id, diary_id)
)