insert into sys_users(created_at, updated_at, deleted_at, uuid, username, password,role)
values (now(), now(), null, gen_random_uuid(), 'root', '21232f297a57a5a743894a0e4a801fc3','super');
insert into sys_users(created_at, updated_at, deleted_at, uuid, username, password,role)
values (now(), now(), null, gen_random_uuid(), 'admin', '21232f297a57a5a743894a0e4a801fc3','admin');
insert into sys_users(created_at, updated_at, deleted_at, uuid, username, password,role)
values (now(), now(), null, gen_random_uuid(), 'noob', '21232f297a57a5a743894a0e4a801fc3','user');
commit;