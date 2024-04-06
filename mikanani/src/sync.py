import re
import os
import traceback
from bson.int64 import Int64
from db_helper import get_mongo_client, get_mysql_conn
from configs import QbitConfig, MongoDBConfig
from logger import LOGGER
from utils import numset2bitmap

class MikanamiAnimeSync:
    def sync(self) -> None:
        # Get expected-list in mongodb
        anime_infos = dict()
        
        conn = get_mysql_conn()
        sql = f"SELECT uid, download_bitmap FROM `mikanani`.`anime_meta` WHERE is_active IS TRUE;"
        cursor = conn.cursor()
        cursor.execute(sql)
        result = cursor.fetchall()
        if not result:
            return

        for uid, download_bitmap in result:
            anime_infos.setdefault(uid, {'downloaded_bitmap': download_bitmap})

        mongo_client = get_mongo_client()
        mongo_db = mongo_client[MongoDBConfig['mikandb']]
        mongo_col = mongo_db[MongoDBConfig['mikancollection']]

        query = { "uid": {"$in": [ Int64(uid) for uid in anime_infos.keys() ]} }

        doc_data = [x for x in mongo_col.find(query, {"_id": 0})]
        for doc in doc_data:
            anime_infos[int(doc["uid"])]["regex"] = doc.get("regex")
        
        to_update_animes = dict()
        for uid, info in anime_infos.items():
            target_dir = os.path.join(QbitConfig["nfs_media_file_dir"], f"medias/{uid}")
            num_set = set()
            
            files = os.listdir(target_dir)
            files = [entry for entry in files if os.path.isfile(os.path.join(target_dir, entry))]
            regex_pattern = re.compile(info.get("regex"))
            for file in files:
                if match_obj := regex_pattern.match(file):
                    num_set.add(int(match_obj.groups(1)[0]))
            curr_bitmap = numset2bitmap(num_set)
            if curr_bitmap != info.get("download_bitmap"):
                to_update_animes[uid] = curr_bitmap
        
        if to_update_animes:
            LOGGER.info(f"[Sync]need to update bitmap today count: {len(to_update_animes)}")
            try:
                conn = get_mysql_conn()
                for uid, bitmap in to_update_animes.items():
                    usql = ("UPDATE `mikanani`.`anime_meta` "
                            f"SET download_bitmap = {bitmap} "
                            f"WHERE uid = {uid};")
                    cursor = conn.cursor()
                    cursor.execute(usql)
                    conn.commit()
                    LOGGER.info(f"[Sync][SUCCESS]: uid:{uid} bitmap to {bitmap}.")
            except Exception as e:
                LOGGER.error(f"[Sync][FAILED] Exception {e}: {traceback.format_exc()}")