DO $$
DECLARE
     ques_id varchar  := gen_random_uuid()::varchar;
   
BEGIN  
    INSERT INTO public.questions 
    (
        id,
        question_desc
    )VALUES(
         ques_id,
        'MBTITYPE'
    );
    INSERT INTO public.answers 
    (
        id,
        question_id,
        answer_desc,
        answer_score
    )VALUES(
        gen_random_uuid()::varchar,
        ques_id ,
        'INFJ',
        5
    );

    INSERT INTO public.answers 
    (
        id,
        question_id,
        answer_desc,
        answer_score
    )VALUES(
        gen_random_uuid()::varchar,
        ques_id ,
        'ENFJ',
        8
    );

    INSERT INTO public.answers 
    (
        id,
        question_id,
        answer_desc,
        answer_score
    ) VALUES(
        gen_random_uuid()::varchar,
        ques_id ,
        'INFP',
        0
    );
    INSERT INTO public.answers 
    (
        id,
        question_id,
        answer_desc,
        answer_score
    ) VALUES(
        gen_random_uuid()::varchar,
        ques_id ,
        'ENFP',
        1
    );

    ques_id   := gen_random_uuid()::varchar;

    INSERT INTO public.questions 
    (
        id,
        question_desc
    )VALUES(
         ques_id,
        'Blood Group'
    );
    INSERT INTO public.answers 
    (
        id,
        question_id,
        answer_desc,
        answer_score
    )VALUES(
        gen_random_uuid()::varchar,
        ques_id ,
        'A',
        5
    );

    INSERT INTO public.answers 
    (
        id,
        question_id,
        answer_desc,
        answer_score
    )VALUES(
        gen_random_uuid()::varchar,
        ques_id ,
        'B',
        8
    );

    INSERT INTO public.answers 
    (
        id,
        question_id,
        answer_desc,
        answer_score
    ) VALUES(
        gen_random_uuid()::varchar,
        ques_id ,
        'O',
        0
    );
    INSERT INTO public.answers 
    (
        id,
        question_id,
        answer_desc,
        answer_score
    ) VALUES(
        gen_random_uuid()::varchar,
        ques_id ,
        'AB',
        1
    );

END $$;


