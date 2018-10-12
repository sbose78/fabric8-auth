ALTER TABLE RESOURCE_TYPE ADD COLUMN DEFAULT_ROLE_ID uuid NULL references role(role_id);


UPDATE RESOURCE_TYPE SET DEFAULT_ROLE_ID = a.ROLE_ID FROM ( SELECT ROLE_ID FROM ROLE WHERE NAME = 'admin' AND RESOURCE_TYPE_ID IN ( SELECT RESOURCE_TYPE_ID FROM RESOURCE_TYPE WHERE NAME = 'identity/organization')) a WHERE NAME = 'identity/organization';
UPDATE RESOURCE_TYPE SET DEFAULT_ROLE_ID = a.ROLE_ID FROM ( SELECT ROLE_ID FROM ROLE WHERE NAME = 'admin' AND RESOURCE_TYPE_ID IN ( SELECT RESOURCE_TYPE_ID FROM RESOURCE_TYPE WHERE NAME = 'openshift.io/resource/system')) a WHERE NAME = 'openshift.io/resource/system';
UPDATE RESOURCE_TYPE SET DEFAULT_ROLE_ID = a.ROLE_ID FROM ( SELECT ROLE_ID FROM ROLE WHERE NAME = 'admin' AND RESOURCE_TYPE_ID IN ( SELECT RESOURCE_TYPE_ID FROM RESOURCE_TYPE WHERE NAME = 'openshift.io/resource/space')) a WHERE NAME = 'openshift.io/resource/space';

