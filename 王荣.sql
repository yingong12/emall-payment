select userid, tag_ids, update_time, workid,category ,whole_score, opt_time, songid from zt_audit.audit_result where `opt_time` > '2022-08-20' and  `final_audit_status` = 1 and `category`!=7 and `whole_score` >= 3.0 and `status` = 0 order by update_time desc