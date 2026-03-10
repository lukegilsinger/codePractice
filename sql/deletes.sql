-- Created Table for the cleanup 
SELECT * 
FROM utilities.gi_cleanup_2025_12_08 
WHERE 1=1 



-- Backup BQ 
CREATE TABLE `bcs-customer360-prod.event_store_backup.SFDC_GI_account-default_bu_2025_12_08` COPY `bcs-customer360-prod.event_store_default.SFDC_GI_account-default`; 
CREATE TABLE `bcs-customer360-prod.event_store_backup.opl_account-global-object-default_bu_2025_12_08` COPY `bcs-customer360-prod.event_store_default.opl_account-global-object-default`; 
CREATE TABLE `bcs-customer360-prod.event_store_backup.RELTIO_crm_account-default_bu_2025_12_08` COPY `bcs-customer360-prod.event_store_default.RELTIO_crm_account-default`; 
CREATE TABLE `bcs-customer360-prod.event_store_backup.crm_account-global-object-default_bu_2025_12_08` COPY `bcs-customer360-prod.event_store_default.crm_account-global-object-default`; 

-- Backup PG 
CREATE TABLE SFDC_GI.ACCOUNT_bu_2025_12_08 as table SFDC_GI.ACCOUNT; 
CREATE TABLE interim.opl_account_interim_bu_2025_12_08 AS TABLE interim.opl_account_interim; 
CREATE TABLE canonical.opl_account_bu_2025_12_08 AS TABLE canonical.opl_account; 
CREATE TABLE reltio.crm_account_bu_2025_12_08 as table reltio.crm_account; 
CREATE TABLE interim.crm_account_interim_bu_2025_12_08 as table interim.crm_account_interim; 
CREATE TABLE canonical.crm_account_bu_2025_12_08 as table canonical.crm_account; 




-- Delete BQ 
--delete from `bcs-customer360-prod.event_store_default.SFDC_GI_account-default` 
--where message_id in  
( 
select a.message_id 
from `bcs-customer360-prod.event_store_default.SFDC_GI_account-default` as a 
inner join `bcs-customer360-prod.support.gi_cleanup_2025_12_08` as b 
on JSON_EXTRACT_SCALAR(data, '$.id') = b.opl_id 
WHERE 1=1 
) 
; 
--delete FROM `bcs-customer360-prod.event_store_default.opl_account-global-object-default` 
--where message_id in  
( 
SELECT a.message_id 
FROM `bcs-customer360-prod.event_store_default.opl_account-global-object-default` 
inner join `bcs-customer360-prod.support.gi_cleanup_2025_12_08` as b 
on JSON_EXTRACT_SCALAR(data, '$.account_id') = b.opl_id 
WHERE 1=1 
) 
; 




-- Delete PG 
--delete from sfdc_gi.account 
--where id in  
( 
SELECT a.id 
FROM sfdc_gi.account as a 
inner join utilities.gi_cleanup_2025_12_08 as b 
on a.id = b.opl_id 
WHERE 1=1 
) 
; 
--delete from interim.opl_account_interim 
--where account_id in  
( 
SELECT a.account_id 
FROM interim.opl_account_interim as a 
inner join utilities.gi_cleanup_2025_12_08 as b 
on a.account_id = b.opl_id 
) 
; 
--delete from canonical.opl_account 
--where account_id in  
( 
SELECT a.account_id 
FROM canonical.opl_account as a 
inner join utilities.gi_cleanup_2025_12_08 as b 
on a.account_id = b.opl_id 
WHERE 1=1 
) 
; 




-- PG Stat Activity SQL 
--active session list 
SELECT  
pid, usename, application_name 
, wait_event_type, wait_event, state 
, pg_blocking_pids(pid) AS blocked_by_pids 
, query, query_start, state_change 
--, pg_terminate_backend(pid)  
--, * 
FROM pg_stat_activity 
WHERE 1=1 
AND DATNAME = 'staging' 
and pid <> pg_backend_pid() 
and state <> 'idle' 
--and application_name = 'MetricSampler_opl_account_stale_account_active' 
--and pid = 4036773 
--and application_name ='MetricSampler_opl_account_stale_account_active' 
order by wait_event_type, wait_event, application_name 
; 