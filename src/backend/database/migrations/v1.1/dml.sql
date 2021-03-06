INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('122', 'INGRESS_CREATE', '', now(), now());
INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('123', 'INGRESS_UPDATE', '', now(), now());
INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('124', 'INGRESS_READ', '', now(), now());
INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('125', 'INGRESS_DELETE', '', now(), now());
INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('126', 'INGRESS_DEPLOY', '', now(), now());

INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '95');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '96');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '97');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '98');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '99');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '100');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '101');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '102');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '103');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '122');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '123');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '124');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '125');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('10', '126');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('11', '96');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('11', '97');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('11', '99');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('11', '100');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('11', '123');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('11', '124');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('11', '126');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('12', '124');
INSERT INTO  group_permissions  ( group_id,  permission_id ) VALUES ('12', '126');
