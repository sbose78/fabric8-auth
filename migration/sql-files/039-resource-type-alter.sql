ALTER TABLE RESOURCE_TYPE ADD COLUMN DEFAULT_ROLE_ID uuid NULL references role(role_id);
UPDATE RESOURCE_TYPE SET DEFAULT_ROLE_ID = '2d993cbd-83f5-4e8c-858f-ca11bcf718b0' where NAME = 'openshift.io/resource/space';