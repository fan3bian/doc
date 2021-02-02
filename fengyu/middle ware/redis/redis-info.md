# Server
redis_version:5.0.7
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:331b3b4f519a5508
redis_mode:standalone
os:Linux 4.18.16-1.el7.elrepo.x86_64 x86_64
arch_bits:64
multiplexing_api:epoll
atomicvar_api:atomic-builtin
gcc_version:8.3.0
process_id:1
run_id:0876048968fad7581f09318f0df263cdb5e6772f
tcp_port:7000
uptime_in_seconds:5480
uptime_in_days:0
hz:10
configured_hz:10
lru_clock:1564633
executable:/opt/bitnami/redis/bin/redis-server
config_file:/opt/bitnami/redis/etc/redis.conf

# Clients
connected_clients:336
client_recent_max_input_buffer:4
client_recent_max_output_buffer:0
blocked_clients:0

# Memory
used_memory:266290136
used_memory_human:253.95M
used_memory_rss:270553088
used_memory_rss_human:258.02M
used_memory_peak:278452224
used_memory_peak_human:265.55M
used_memory_peak_perc:95.63%
used_memory_overhead:18029424
used_memory_startup:791520
used_memory_dataset:248260712
used_memory_dataset_perc:93.51%
allocator_allocated:266413488
allocator_active:268201984
allocator_resident:279355392
total_system_memory:67530428416
total_system_memory_human:62.89G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.01
allocator_frag_bytes:1788496
allocator_rss_ratio:1.04
allocator_rss_bytes:11153408
rss_overhead_ratio:0.97
rss_overhead_bytes:-8802304
mem_fragmentation_ratio:1.02
mem_fragmentation_bytes:4305120
mem_not_counted_for_evict:134
mem_replication_backlog:1048576
mem_clients_slaves:99388
mem_clients_normal:6144766
mem_aof_buffer:134
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0

# Persistence
loading:0
rdb_changes_since_last_save:1899563
rdb_bgsave_in_progress:0
rdb_last_save_time:1612171911
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_last_cow_size:593920
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:3
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_last_write_status:ok
aof_last_cow_size:4591616
aof_current_size:115286274
aof_base_size:113527035
aof_pending_rewrite:0
aof_buffer_length:0
aof_rewrite_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0

# Stats
total_connections_received:34004
total_commands_processed:5948071
instantaneous_ops_per_sec:768
total_net_input_bytes:430399916
total_net_output_bytes:854517436
instantaneous_input_kbps:34.08
instantaneous_output_kbps:61.79
rejected_connections:0
sync_full:2
sync_partial_ok:0
sync_partial_err:0
expired_keys:246
expired_stale_perc:0.00
expired_time_cap_reached_count:0
evicted_keys:0
keyspace_hits:26592
keyspace_misses:15244
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:12645
migrate_cached_sockets:0
slave_expires_tracked_keys:0
active_defrag_hits:0
active_defrag_misses:0
active_defrag_key_hits:0
active_defrag_key_misses:0

# Replication
role:master
connected_slaves:2
slave0:ip=172.20.105.254,port=7000,state=online,offset=370915759,lag=1
slave1:ip=172.20.125.203,port=7000,state=online,offset=370920764,lag=0
master_replid:9ee0ffaf06cebfd27b9e89fd9f5ce524de7db5ff
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:370937891
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:369889316
repl_backlog_histlen:1048576

# CPU
used_cpu_sys:597.583288
used_cpu_user:111.092458
used_cpu_sys_children:1.205525
used_cpu_user_children:5.714404

# Cluster
cluster_enabled:0

# Keyspace
db0:keys=151416,expires=52790,avg_ttl=256616346
