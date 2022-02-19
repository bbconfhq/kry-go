insert into kry.problems
(created_at, updated_at, title, content, note, time_limit, memory_limit)
values
    (now(), now(), 'A+B', 'A + B를 리턴하시오.', '', 0.125, 128),
    (now(), now(), 'A-B', 'A - B를 리턴하시오.', '', 0.125, 128),
    (now(), now(), 'A/B', 'A / B를 리턴하시오.', '재미없다.', 0.125, 128),
    (now(), now(), 'A*B', 'A * B를 리턴하시오.', '재미있다.', 0.125, 128),
    (now(), now(), '01234567890123456789012345678901234567890123456789', 'Test Problem', '풀 수 없다.', 0.125, 128);

insert into kry.contests
(created_at, updated_at, title, content, started_at, ended_at)
values
    (now(), now(), '제 1회 완장전', '커동빵디', DATE_SUB(CURDATE(),INTERVAL 7 DAY), DATE_ADD(CURDATE(),INTERVAL 7 DAY)),
    (now(), now(), '제 2회 완장전', '커동빵디', DATE_SUB(CURDATE(),INTERVAL 7 DAY), DATE_SUB(CURDATE(),INTERVAL 3 DAY)),
    (now(), now(), '제 3회 완장전', '커동빵디', DATE_SUB(CURDATE(),INTERVAL 6 DAY), DATE_SUB(CURDATE(),INTERVAL 2 DAY)),
    (now(), now(), '제 4회 완장전', '커동빵디', DATE_SUB(CURDATE(),INTERVAL 5 DAY), DATE_ADD(CURDATE(),INTERVAL 2 DAY)),
    (now(), now(), '제 5회 완장전', '커동빵디', DATE_ADD(CURDATE(),INTERVAL 4 DAY), DATE_ADD(CURDATE(),INTERVAL 5 DAY)),
    (now(), now(), '제 6회 완장전', '커동빵디', DATE_ADD(CURDATE(),INTERVAL 3 DAY), DATE_ADD(CURDATE(),INTERVAL 6 DAY)),
    (now(), now(), '제 7회 완장전', '커동빵디', DATE_ADD(CURDATE(),INTERVAL 2 DAY), DATE_ADD(CURDATE(),INTERVAL 7 DAY));

insert into kry.tags
(created_at, updated_at, name)
values
    (now(), now(), '연산'),
    (now(), now(), 'DP'),
    (now(), now(), '브루트포스');

insert into kry.testcases
(created_at, updated_at, visible, input, output, problem_id)
values
    (now(), now(), true, '1 1', '2', 1),
    (now(), now(), false, '2 2', '4', 1),
    (now(), now(), true, '3 3', '6', 1),
    (now(), now(), false, '4 4', '8', 1),
    (now(), now(), true, '5 5', '10', 1),
    (now(), now(), false, '6 6', '12', 1),
    (now(), now(), false, '7 7', '14', 1),
    (now(), now(), false, '8 8', '16', 1),
    (now(), now(), false, '9 9', '18', 1),
    (now(), now(), true, '1000000000 1234567890', '2234567890', 1),
    (now(), now(), true, '1 1', '0', 2),
    (now(), now(), true, '2 2', '0', 2),
    (now(), now(), false, '3 3', '0', 2),
    (now(), now(), false, '4 4', '0', 2),
    (now(), now(), false, '5 5', '0', 2),
    (now(), now(), false, '6 6', '0', 2),
    (now(), now(), false, '7 7', '0', 2),
    (now(), now(), false, '8 8', '0', 2),
    (now(), now(), false, '9 9', '0', 2),
    (now(), now(), true, '1234567890 1000000000', '234567890', 2),
    (now(), now(), true, '1 1', '1', 3),
    (now(), now(), true, '2 2', '1', 3),
    (now(), now(), false, '3 3', '1', 3),
    (now(), now(), false, '4 4', '1', 3),
    (now(), now(), true, '5 5', '1', 3),
    (now(), now(), false, '6 6', '1', 3),
    (now(), now(), false, '7 7', '1', 3),
    (now(), now(), false, '8 8', '1', 3),
    (now(), now(), false, '9 9', '1', 3),
    (now(), now(), true, '1234567890 1000000000', '1', 2),
    (now(), now(), true, '1 1', '1', 4),
    (now(), now(), true, '2 2', '4', 4),
    (now(), now(), true, '3 3', '9', 4),
    (now(), now(), true, '4 4', '16', 4),
    (now(), now(), true, '5 5', '25', 4),
    (now(), now(), true, '6 6', '36', 4),
    (now(), now(), true, '7 7', '49', 4),
    (now(), now(), true, '8 8', '64', 4),
    (now(), now(), true, '9 9', '81', 4),
    (now(), now(), true, '1 1', '', 5),
    (now(), now(), false, '2 2', '', 5),
    (now(), now(), false, '3 3', '', 5),
    (now(), now(), false, '4 4', '', 5),
    (now(), now(), false, '5 5', '', 5),
    (now(), now(), false, '6 6', '', 5),
    (now(), now(), false, '7 7', '', 5),
    (now(), now(), false, '8 8', '', 5),
    (now(), now(), false, '9 9', '', 5);

insert into kry.problem_tags
(problem_id, tag_id)
values
    (1, 1),
    (2, 1),
    (3, 1),
    (4, 1),
    (5, 1),
    (5, 2),
    (5, 3);

insert into kry.contest_problems
(contest_id, problem_id)
values
    (1, 1),
    (1, 2),
    (1, 3),
    (1, 4),
    (1, 5),
    (2, 1),
    (3, 1),
    (3, 2),
    (3, 3),
    (3, 4),
    (4, 1),
    (4, 2),
    (4, 4),
    (5, 1),
    (5, 3),
    (5, 5),
    (6, 1),
    (6, 5),
    (7, 1),
    (7, 2),
    (7, 3),
    (7, 4),
    (7, 5);

select * from kry.contests;
select * from kry.problems;
select * from kry.tags;
select * from kry.testcases;
select * from kry.problem_tags;
select * from kry.contest_problems;